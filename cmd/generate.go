// File: cmd/generate.go
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/manyfacedqod/smartcommit/config"
	"github.com/manyfacedqod/smartcommit/diff"
	"github.com/manyfacedqod/smartcommit/llm"
)

var yesFlag bool

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a commit message using AI",
	Long: `Generate a commit message from staged Git changes using a local or remote LLM.

By default, it launches an interactive flow.
Use --yes or -y to skip the prompt and commit directly.
Flags after -- will be passed directly to git. `,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadOrDefault()

		diffText, err := diff.GetStagedDiff(args)
		if err != nil || diffText == "" {
			fmt.Println("❌ No staged changes found.")
			return
		}

		promptText := cfg.SystemPrompt + "\n\nDiff:\n" + diffText

		provider, err := llm.GetProvider(cfg)
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		message, err := provider.Generate(promptText)
		if err != nil {
			fmt.Println("❌ Generation failed:", err)
			return
		}
		message = strings.TrimSpace(message)

		if yesFlag {
			fmt.Println("💡 Generated Commit Message:")
			fmt.Println("----------------------------------")
			fmt.Println(message)
			fmt.Println("----------------------------------")
			runGitCommit(args, message)
			return
		}

		for {
			fmt.Println("\n💡 Generated Commit Message:")
			fmt.Println("----------------------------------")
			fmt.Println(message)
			fmt.Println("----------------------------------")
			fmt.Print("Choose [c]ommit, [e]dit, [r]egen, [q]uit: ")

			if err := keyboard.Open(); err != nil {
				fmt.Println("❌ Keyboard input error:", err)
				return
			}
			char, _, err := keyboard.GetSingleKey()
			keyboard.Close()
			if err != nil {
				fmt.Println("❌ Failed to read key:", err)
				return
			}

			switch strings.ToLower(string(char)) {
			case "c":
				runGitCommit(args, message)
				return
			case "e":
				message = editMessage(message)
			case "r":
				fmt.Print("\n🔄 Regenerating")
				for i := 0; i < 3; i++ {
					fmt.Print(".")
					time.Sleep(200 * time.Millisecond)
				}
				message, err = provider.Generate(promptText)
				if err != nil {
					fmt.Println("❌ Regeneration failed:", err)
					continue
				}
				message = strings.TrimSpace(message)
			case "q":
				fmt.Println("👋 Aborted.")
				return
			default:
				fmt.Println("❓ Invalid choice.")
			}
		}
	},
}

func init() {
	generateCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Autogenerate and commit without prompting")
	rootCmd.AddCommand(generateCmd)
}

func runGitCommit(options []string, msg string) {
	fmt.Println("✅ Committing with:")
	fmt.Println(msg)
	options = append(options, "commit", "-m", msg);
	cmd := exec.Command("git", options...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Git commit failed:", err)
	}
}

func editMessage(current string) string {
	prompt := promptui.Prompt{
		Label:     "Edit Commit Message",
		Default:   current,
		AllowEdit: true,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println("⚠️ Edit cancelled.")
		return current
	}
	return strings.TrimSpace(result)
}

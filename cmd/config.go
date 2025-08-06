package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manyfacedqod/smartcommit/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View or manage SmartCommit configuration",
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadOrDefault()
		fmt.Println("📁 Current Config:")
		fmt.Println("  Provider:     ", cfg.Provider)
		fmt.Println("  Model:        ", cfg.Model)
		fmt.Println("  Base URL:     ", cfg.BaseURL)
		if cfg.APIKey != "" {
			fmt.Println("  API Key:      ", mask(cfg.APIKey))
		}
		fmt.Println("  SystemPrompt: ", cfg.SystemPrompt)
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the system prompt used for generation (in Vim)",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadOrDefault()

		tmpFile, err := os.CreateTemp("", "smartcommit-prompt-*.txt")
		if err != nil {
			fmt.Println("❌ Could not create temp file:", err)
			return
		}
		defer os.Remove(tmpFile.Name())

		// Write current prompt
		_, _ = tmpFile.WriteString(cfg.SystemPrompt)
		tmpFile.Close()

		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = os.Getenv("VISUAL")
		}
		if editor == "" {
			editor = "vi"
		}

		editCmd := exec.Command(editor, tmpFile.Name())
		editCmd.Stdin = os.Stdin
		editCmd.Stdout = os.Stdout
		editCmd.Stderr = os.Stderr
		if err := editCmd.Run(); err != nil {
			fmt.Println("❌ Editor closed with error:", err)
			return
		}

		// Read new prompt
		newPromptBytes, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			fmt.Println("❌ Could not read edited file:", err)
			return
		}
		newPrompt := strings.TrimSpace(string(newPromptBytes))
		if newPrompt == "" {
			fmt.Println("⚠️  Prompt is empty — edit cancelled.")
			return
		}

		cfg.SystemPrompt = newPrompt
		if err := config.Save(cfg); err != nil {
			fmt.Println("❌ Failed to save config:", err)
			return
		}
		fmt.Println("✅ System prompt updated.")
	},
}


func init() {
	configCmd.AddCommand(showCmd)
	configCmd.AddCommand(editCmd)
	rootCmd.AddCommand(configCmd)
}

func mask(s string) string {
	if len(s) <= 6 {
		return "******"
	}
	return s[:3] + "..." + s[len(s)-3:]
}

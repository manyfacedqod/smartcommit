# SmartCommit

**SmartCommit** is a blazing-fast CLI tool that uses AI to generate meaningful Git commit messages from your staged changes — powered by local or remote LLMs (Ollama, OpenAI, or any HTTP-based LLM).

---

## 🚀 Features

* 🔍 Reads your staged Git diff  
* 🤖 Supports multiple LLM providers: Ollama, OpenAI, HTTP endpoints (Claude, Gemini, etc.)  
* ✏️ Customizable system prompt (edit in Vim)  
* ⚡ Interactive flow with options to commit, edit, regenerate, or quit  
* ⚡ One‐shot `--yes` flag for auto‐commit  
* 📦 Single static binaries for Linux, macOS (Intel/ARM), and Windows  
* 🔖 `--version` flag and `version` subcommand  

---

## 📥 Installation

### Using Go

```bash
go install github.com/manyfacedqod/smartcommit@latest
```

### Download Prebuilt Binary

1. Visit the [Releases page](https://github.com/manyfacedqod/smartcommit/releases).
2. Download the `.tar.gz` for your OS and extract:

   ```bash
   tar -xzf smartcommit-<os>.tar.gz
   ```

3. Move the binary into your `PATH`:

   ```bash
   mv smartcommit-<os> /usr/local/bin/smartcommit  # or equivalent
   chmod +x /usr/local/bin/smartcommit
   ```

---

## ⚙️ First‑Time Setup

```bash
smartcommit setup
```

Configure in an interactive prompt:

1. **Provider**: `ollama`, `openai`, or `http`  
2. **Model name**: e.g., `llama3`, `gpt-4`, `gemini`  
3. **API Key** (HTTP/OpenAI) or **Base URL** (HTTP/Ollama)  
4. **System prompt** saved to `~/.config/smartcommit/config.yaml`  

---

## 💡 Usage Examples

### Interactive Commit Flow

```bash
git add .
smartcommit generate
```

```
💡 Generated Commit Message:
fix: handle empty username in login flow

Choose [c]ommit, [e]dit, [r]egenerate, [q]uit:
```

* **c**: commit  
* **e**: inline edit via PromptUI  
* **r**: regenerate  
* **q**: quit  

### Auto‑Commit Without Prompt

```bash
smartcommit generate --yes
```

### Config Commands (the system prompt for your commits)

```bash
smartcommit config show   # display current config
smartcommit config edit   # edit system prompt in $EDITOR
```


### Version

```bash
smartcommit --version
smartcommit version
```

---

## 📄 Command Reference

| Command                      | Description                                 |
| ---------------------------- | ------------------------------------------- |
| `smartcommit setup`          | Configure your LLM provider and settings    |
| `smartcommit generate`       | Generate commit message from staged diff    |
| `smartcommit generate --yes` | Auto‑generate & commit without prompt       |
| `smartcommit config show`    | Display current configuration               |
| `smartcommit config edit`    | Edit the system prompt using your `$EDITOR` |
| `smartcommit --version`      | Show the CLI version                        |
| `smartcommit version`        | Same as `--version`                         |

---

## 👨‍💻 Contributing

We welcome your contributions! 🚀

1. **Fork** the repo and **clone** your fork:

   ```bash
   git clone https://github.com/manyfacedqod/smartcommit.git
   cd smartcommit
   ```

2. **Create** a feature branch:

   ```bash
   git checkout -b feat/your-feature
   ```

3. **Make changes**, **add tests**, and **update docs**.

4. **Run** tests & build:

   ```bash
   go test ./...
   go build
   ```

5. **Commit**, **push**, and open a **Pull Request**.

### Guidelines

- Follow existing code style and idioms.  
- Use [Conventional Commits](https://www.conventionalcommits.org/) (`feat:`, `fix:`, `docs:`, etc.).  
- Write clear tests for new features or fixes.  
- Update this README when adding or changing commands.  

---
(this one kinda ai generated, will update later)
## 📄 License

MIT License © manyfacedqod

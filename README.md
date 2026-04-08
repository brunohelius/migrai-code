# MigrAI Code

AI coding agent for the terminal.

MigrAI Code is a fork of [OpenCode](https://github.com/opencode-ai/opencode)
(MIT License), customized for the MigrAI platform. OpenCode was created by
[@kujtimiihoxha](https://github.com/kujtimiihoxha) and the project has since
been archived upstream, continuing as
[Crush](https://github.com/charmbracelet/crush) under the Charm team.

---

## Features

- Interactive TUI built with [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- Multiple AI providers: OpenAI, Anthropic, Google Gemini, AWS Bedrock, Groq, Azure OpenAI, GitHub Copilot, OpenRouter
- Session management with persistent SQLite storage
- Tool integration: shell commands, file search, code editing, LSP diagnostics
- MCP (Model Context Protocol) support for external tool servers
- Non-interactive prompt mode for scripting and automation

---

## Installation

### From Source (Go)

```bash
go install github.com/brunohelius/migrai-code/cmd/...@latest
```

### Build Locally

```bash
git clone https://github.com/brunohelius/migrai-code.git
cd migrai-code
go build -o migrai-code ./cmd/...
./migrai-code
```

---

## Quick Start

Set your API key for the provider you want to use:

```bash
# Anthropic (Claude)
export ANTHROPIC_API_KEY="your-key"

# OpenAI
export OPENAI_API_KEY="your-key"

# Google Gemini
export GEMINI_API_KEY="your-key"

# MigrAI platform key (when available)
export MIGRAI_API_KEY="your-key"
```

Then launch the agent:

```bash
migrai-code
```

Or run a single prompt non-interactively:

```bash
migrai-code -p "Explain the use of context in Go"
```

---

## Configuration

MigrAI Code looks for configuration in:

- `$HOME/.migrai-code.json`
- `$XDG_CONFIG_HOME/migrai-code/.migrai-code.json`
- `./.migrai-code.json` (project-local)

See the full configuration reference in the upstream
[OpenCode documentation](https://github.com/opencode-ai/opencode#configuration).

---

## Command-Line Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--help` | `-h` | Display help |
| `--version` | `-v` | Print version |
| `--debug` | `-d` | Enable debug mode |
| `--cwd` | `-c` | Set working directory |
| `--prompt` | `-p` | Run a single prompt (non-interactive) |
| `--output-format` | `-f` | Output format: `text` (default) or `json` |
| `--quiet` | `-q` | Hide spinner in non-interactive mode |

---

## Links

- [MigrAI](https://migrai.com) -- MigrAI platform
- [OpenCode](https://github.com/opencode-ai/opencode) -- upstream (archived, MIT)
- [Crush](https://github.com/charmbracelet/crush) -- upstream continuation by Charm

---

## License

MIT License. See [LICENSE](LICENSE) for details.

Portions of this software are derived from OpenCode, copyright its original
authors, used under the MIT License.

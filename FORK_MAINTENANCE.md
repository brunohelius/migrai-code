# Fork Maintenance Guide

This document describes how to maintain the MigrAI Code fork relative to the
upstream OpenCode repository.

---

## Fork Origin

| | |
|---|---|
| **Upstream** | https://github.com/opencode-ai/opencode (MIT License, archived) |
| **Fork** | https://github.com/brunohelius/migrai-code |
| **License** | MIT (inherited from upstream) |

OpenCode is currently archived. The upstream project continued under the name
[Crush](https://github.com/charmbracelet/crush) by the original author and the
Charm team. If the community creates patches against the original OpenCode repo,
or if Crush diverges in useful ways, we can evaluate pulling changes.

---

## Remote Setup

```
git remote -v
# origin    https://github.com/brunohelius/migrai-code.git (fetch/push)
# upstream  https://github.com/opencode-ai/opencode.git (fetch/push)
```

If `upstream` is not configured:

```bash
git remote add upstream https://github.com/opencode-ai/opencode.git
```

---

## How to Sync with Upstream

```bash
git fetch upstream
git checkout main
git merge upstream/main --no-edit
# Resolve conflicts (mainly in rebranded files — see sections below)
# Re-apply MigrAI branding patches if needed
git push origin main
```

After merging, always verify the build:

```bash
go build ./cmd/...
```

---

## MigrAI-Specific Files

These files do not exist in upstream and will never produce merge conflicts:

- `internal/llm/models/migrai.go` -- MigrAI model catalog (when created)
- `internal/llm/provider/migrai.go` -- MigrAI provider (when created)
- Any files under `internal/migrai/` directory
- `FORK_MAINTENANCE.md` -- this file

---

## Files with MigrAI Modifications

These files have been modified from upstream and are the most likely sources of
merge conflicts during a sync:

| File | What Changed |
|------|-------------|
| `go.mod` | Module path changed to `github.com/brunohelius/migrai-code` |
| `cmd/root.go` | Binary name, CLI descriptions |
| `internal/config/config.go` | Config paths, environment variables |
| `internal/llm/provider/provider.go` | MigrAI case in provider switch |
| `internal/llm/models/models.go` | Model registration, provider list |
| `internal/llm/prompt/coder.go` | System prompt adjustments |
| `internal/tui/theme/` | Theme rename / branding |
| `README.md` | Full rebrand |
| `.goreleaser.yml` | Binary name, project name, archive templates |

**Note:** As of the initial fork, imports across ~84 files were bulk-renamed
from `github.com/opencode-ai/opencode` to `github.com/brunohelius/migrai-code`.
A future upstream merge will conflict on every import line. Use a script to
re-apply the module rename after merge:

```bash
# After merging upstream, fix all import paths
find . -name '*.go' -exec sed -i '' \
  's|github.com/opencode-ai/opencode|github.com/brunohelius/migrai-code|g' {} +
# Also fix go.mod
sed -i '' 's|module github.com/opencode-ai/opencode|module github.com/brunohelius/migrai-code|g' go.mod
# Verify
go build ./cmd/...
```

---

## Conflict Resolution Strategy

1. **Always keep MigrAI branding changes** -- module path, binary name, CLI
   descriptions, README, system prompts.
2. **Pull upstream bug fixes and new features** -- provider improvements, tool
   fixes, TUI enhancements, new model support.
3. **Test after every merge** -- `go build ./cmd/...` at minimum; run the binary
   to verify TUI launches.
4. **Re-run the import rename script** if upstream touched many files.

---

## Release Process

### Build for All Platforms

```bash
# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o migrai-code-darwin-arm64 ./cmd/...

# macOS x86_64
GOOS=darwin GOARCH=amd64 go build -o migrai-code-darwin-amd64 ./cmd/...

# Linux x86_64
GOOS=linux GOARCH=amd64 go build -o migrai-code-linux-amd64 ./cmd/...

# Windows x86_64
GOOS=windows GOARCH=amd64 go build -o migrai-code-windows-amd64.exe ./cmd/...
```

### Using GoReleaser

```bash
goreleaser release --clean
```

Make sure `.goreleaser.yml` references `migrai-code` as the project name and
binary output before running a release.

---

## Checklist: After Upstream Sync

- [ ] `git fetch upstream && git merge upstream/main --no-edit`
- [ ] Resolve all merge conflicts
- [ ] Re-apply import path rename (`opencode-ai/opencode` -> `brunohelius/migrai-code`)
- [ ] Verify `go.mod` module path is `github.com/brunohelius/migrai-code`
- [ ] Verify `cmd/root.go` still shows MigrAI branding
- [ ] Verify `.goreleaser.yml` project name is `migrai-code`
- [ ] Run `go build ./cmd/...`
- [ ] Smoke-test the binary: `./migrai-code`
- [ ] Push: `git push origin main`

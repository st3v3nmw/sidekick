# Sidekick

A CLI tool that understands your natural language commands and executes them for you.

It's simple, just ask your sidekick to do something 😊:

```console
$ sk push to gh
Step #1
Command: git status
Check repo status first (read-only)

On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   README.md

no changes added to commit (use "git add" and/or "git commit -a")

Step #2
Command: git add README.md
Stage modified README.md file (local change only)

Step #3
Command: git commit -m "Update README.md"
Commit staged changes (local change only)

[main b4fa01c] Update README.md
 1 file changed, 43 insertions(+), 1 deletion(-)

Step #4
Command: git push
Push commits to remote (network operation)
```

A special command `sk --fix` exists that checks `tmux`'s pane output for errors and attempts to fix them.

```console
$ sk --fix
```

## Installation

Install `sidekick`:

```console
$ go install github.com/st3v3nmw/sidekick/cmd/sidekick
```

Add an alias to your shell profile (`.bashrc`, `.zshrc`, etc.):

```bash
alias sk="sidekick"
```

Make sure these environment variables are set:

```bash
export SK_PROVIDER="openrouter"
export SK_MODEL="anthropic/claude-3.5-sonnet"
export SK_API_KEY="your-api-key"
```

Currently, only [OpenRouter](https://openrouter.ai/) is supported.

### Upgrading

Remove the old version:

```console
$ rm $(which sidekick)
```

Install the new version:

```console
$ go install github.com/st3v3nmw/sidekick/cmd/sidekick
```

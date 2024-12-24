# Sidekick

A CLI tool that understands your natural language commands and executes them for you.

It's simple, just ask your sidekick to do something 😊:

```console
$ sk push to gh
Step #1
Command: git status
Check repo status first (1 - read-only)

On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   README.md

no changes added to commit (use "git add" and/or "git commit -a")

Step #2
Command: git add README.md
Stage modified README.md file (2 - local change only)

Step #3
Command: git commit -m "Update README.md"
Commit staged changes (2 - local change only)

[main b4fa01c] Update README.md
 1 file changed, 43 insertions(+), 1 deletion(-)

Step #4
Command: git push
Push commits to remote (3 - network operation)
```

A special command `sk --fix` exists that checks `tmux`'s pane output for errors and attempts to fix them.

```console
$ python avg.py
Traceback (most recent call last):
  File "/home/stephen/Projects/sidekick/avg.py", line 7, in <module>
    print("Average:", avg(numbers))
                      ~~~^^^^^^^^^
  File "/home/stephen/Projects/sidekick/avg.py", line 3, in avg
    average = total / len(numbers)
              ~~~~~~^~~~~~~~~~~~~~
ZeroDivisionError: division by zero

$ sk --fix
Step #1
Command: cat avg.py
Need to see code content first to fix division by zero (1 - read only)

def avg(numbers):
    total = sum(numbers)
    average = total / len(numbers)
    return average

numbers = []
print("Average:", avg(numbers))

Step #2
Command: echo 'def avg(numbers):
    if not numbers:
        return 0
    total = sum(numbers)
    average = total / len(numbers)
    return average

numbers = []
print("Average:", avg(numbers))' > avg.py
Fix division by zero by checking empty list (3 - modifies file)
```

## Installation

### Homebrew

Install `sidekick`:

```console
$ brew tap st3v3nmw/tap
$ brew install st3v3nmw/tap/sidekick
```

### Go

Install `sidekick`:

```console
$ go install github.com/st3v3nmw/sidekick/cmd/sidekick
```

## Setup

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

## Upgrading

### Homebrew

```console
$ brew upgrade sidekick
```

### Go

Remove the old version:

```console
$ rm $(which sidekick)
```

Install the new version:

```console
$ go install github.com/st3v3nmw/sidekick/cmd/sidekick
```

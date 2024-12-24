# Sidekick

A command execution assistant that helps users execute commands safely and effectively.

## Features

- Natural language command interpretation
- Risk assessment and safety checks
- Step-by-step task breakdown
- Clear execution reasoning

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

Currently, only OpenRouter is supported.

### Upgrading

Remove the old version:

```console
$ rm $(which sidekick)
```

Install the new version:

```console
$ go install github.com/st3v3nmw/sidekick/cmd/sidekick
```

## Usage

Ask your sidekick to do something:

```console
$ sk
```

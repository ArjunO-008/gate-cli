# Gate CLI

A lightweight automation CLI for running sequential command workflows defined in JSON config files.

---

## Installation

Download the latest binary for your platform from the [GitHub Releases](../../releases) page.

### Linux / macOS

```bash
# Download the zip for your platform, extract, and move to PATH
unzip gate-linux-amd64.zip
chmod +x gate
sudo mv gate /usr/local/bin/
```

### Windows

Download `gate-windows-amd64.zip` from the releases page, extract it, and add the folder to your system `PATH`.

---

## Quick Start

```bash
# 1. Initialize gate (creates config directory and a sample config)
gate init

# 2. Add your own config file
gate add my-workflow.json

# 3. Run it
gate run my-workflow
```

---

## Commands

| Command | Description |
|---|---|
| `gate help` | Show help message |
| `gate version` | Show gate version |
| `gate init` | Initialize the gate config directory |
| `gate add <file...>` | Add one or more JSON config files |
| `gate run <name>` | Run a config by name |
| `gate delete <name...>` | Delete a config by name |
| `gate edit <name>` | Open a config in your default editor |
| `gate config path` | Show the config directory path |
| `gate config list` | List all config files |

---

## Config File Format

Gate uses JSON config files to define automation workflows. Each config has a name, optional settings, and a list of steps to execute sequentially.

```json
{
  "version": "0.1.0",
  "name": "sampleconfig",
  "settings": {
    "workingDirectory": "."
  },
  "steps": [
    {
      "type": "executable",
      "command": "gate",
      "args": "version",
      "dir": "."
    },
    {
      "type": "shellexecutable",
      "command": "echo Build completed successfully"
    }
  ]
}
```

### Top-level Fields

| Field | Type | Required | Description |
|---|---|---|---|
| `version` | string | Yes | Config format version (currently `"0.1.0"`) |
| `name` | string | Yes | Unique identifier used to reference this config |
| `settings` | object | No | Global settings applied to all steps |
| `steps` | array | Yes | List of steps to execute in order |

### Settings

| Field | Type | Description |
|---|---|---|
| `workingDirectory` | string | Base directory for all steps. Defaults to `"."` |

### Step Fields

| Field | Type | Required | Description |
|---|---|---|---|
| `type` | string | Yes | Step type — `"executable"` or `"shellexecutable"` |
| `command` | string | Yes | The command to run |
| `args` | string | No | Arguments to pass (space-separated string, `executable` only) |
| `dir` | string | No | Subdirectory to run the step in, relative to `workingDirectory` |

### Step Types

**`executable`** — Runs a system binary directly without a shell. Best for calling programs like `git`, `go`, `gate`, etc.

```json
{
  "type": "executable",
  "command": "go",
  "args": "build ./...",
  "dir": "."
}
```

**`shellexecutable`** — Runs a command through the system shell (`sh -c` on Linux/macOS, `cmd /C` on Windows). Useful for shell built-ins, scripts, pipes, and chained commands.

```json
{
  "type": "shellexecutable",
  "command": "echo Done && mkdir -p dist"
}
```

---

## Example Workflows

### Build and notify

```json
{
  "version": "0.1.0",
  "name": "build",
  "settings": {
    "workingDirectory": "/path/to/project"
  },
  "steps": [
    {
      "type": "executable",
      "command": "go",
      "args": "build ./...",
      "dir": "."
    },
    {
      "type": "shellexecutable",
      "command": "echo Build finished successfully"
    }
  ]
}
```

Save it and run with:

```bash
gate add build.json
gate run build
```

---

## Config Directory

Gate stores all configs in your OS user config directory:

| Platform | Path |
|---|---|
| Linux | `~/.config/gate-cli/` |
| macOS | `~/Library/Application Support/gate-cli/` |
| Windows | `%APPDATA%\gate-cli\` |

An `.index.json` file inside this directory maps config names to their file paths. It is managed automatically by gate.

Run `gate config path` to print the exact path on your machine.

---

## License

MIT — see [LICENSE](LICENSE)

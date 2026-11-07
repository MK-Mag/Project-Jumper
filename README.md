# Project Jumper

A lightweight Windows CLI that remembers where your projects are and opens them in your configured editor with a single command.

Instead of typing a long directory path and then launching your editor, register a project once with a short name:

```powershell
pj add api here
pj jump api
```

## Features

- **Named project registry** — save projects under short, memorable names.
- **Current-folder shortcut** — use `here` to register the directory you're currently in.
- **Quick project launch** — run `pj jump <project_name>` to open a project in your configured editor.
- **Project listing** — use `pj list` to view saved project names, paths, and creation times.
- **Safe removal** — remove a project from PJ's registry without deleting its actual files.
- **Configurable editor** — use an editor command such as `code`, `idea64`, or another executable available on your `PATH`.
- **Automatic setup** — PJ creates its application-data directory and database table on first run.

## Requirements

- Windows
- Go 1.27 or newer, only if building from source

## Installation

### Option 1: Download a release

1. Visit the [GitHub Releases](https://github.com/MK-Mag/Project-Jumper/releases) page.
2. Download `pj.exe` from the latest release, if a Windows executable is provided.
3. Open PowerShell in the directory where you downloaded the file.
4. Move the executable into your local Programs directory:

```powershell
New-Item -ItemType Directory -Force "$env:LOCALAPPDATA\Programs" | Out-Null
Move-Item .\pj.exe "$env:LOCALAPPDATA\Programs\pj.exe"
```

### Option 2: Build from source

Install Go 1.27 or newer, then run these commands in PowerShell:

```powershell
git clone https://github.com/MK-Mag/Project-Jumper.git
Set-Location Project-Jumper
go build -o pj.exe .
```

Move the generated executable into your local Programs directory:

```powershell
New-Item -ItemType Directory -Force "$env:LOCALAPPDATA\Programs" | Out-Null
Move-Item .\pj.exe "$env:LOCALAPPDATA\Programs\pj.exe"
```

### Add PJ to your PATH

Putting `pj.exe` in `%LOCALAPPDATA%\Programs` does not automatically make the `pj` command available in every terminal. Add that directory to your **user PATH** once by running:

```powershell
$programs = "$env:LOCALAPPDATA\Programs"
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")

if (($userPath -split ";") -notcontains $programs) {
    $newPath = if ([string]::IsNullOrWhiteSpace($userPath)) {
        $programs
    } else {
        "$userPath;$programs"
    }

    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
}
```

Close and reopen PowerShell (or your terminal) so it loads the updated PATH.

Check that PJ is available:

```powershell
pj --help
pj list
```

If you prefer not to modify PATH, run the executable using its full path:

```powershell
& "$env:LOCALAPPDATA\Programs\pj.exe" --help
```

## Usage

### Add a project

```text
pj add <project_name> <project_path>
```

Register an existing directory. PJ stores its absolute path; relative paths are resolved from the current directory.

```powershell
# Register a project using its path
pj add api C:\Users\you\Documents\work\api

# Register the directory you're currently in
Set-Location C:\Users\you\Documents\work\api
pj add api here
```

Project names must be unique. PJ reports an error if you try to add a name that is already registered.

### Jump to a project

```text
pj jump <project_name>
```

Finds the saved project and launches your configured editor with the project's path:

```powershell
pj jump api
```

A CLI child process cannot change the working directory of the terminal that launched it. For that reason, `pj jump` opens the project in your editor rather than changing your shell's directory.

### List projects

```text
pj list
```

Displays all registered projects, sorted by name, along with their paths and creation times.

### Remove a project

```text
pj remove <project_name>
```

Removes a project from Project Jumper's registry:

```powershell
pj remove api
```

**This does not delete the project directory or any of its files.** It only removes the saved entry from PJ's database. You can register the project again later with `pj add`.

### Configure your editor

Show or change the editor used by `pj jump`:

```text
pj config ide [name]
```

With no name, PJ displays the current editor. Provide a name to change it. The default editor is `code`.

```powershell
# Display the current editor
pj config ide

# Set the editor command
pj config ide idea64
```

The editor command must be available on your `PATH`, or otherwise be specified in a way Windows can launch.

## Where PJ stores data

PJ stores its data in:

```text
%LOCALAPPDATA%\PJ\
```

The directory is created automatically when PJ first runs.

| File | Purpose |
| --- | --- |
| `PJ.db` | SQLite-compatible database containing registered projects |
| `config.json` | Configuration settings, including the editor command |

The database contains a `projects` table with the project ID, name, path, and creation time. The table is created automatically if it does not exist.

The configuration file is plain JSON. For example:

```json
{
  "IDE": "code"
}
```

Deleting the `%LOCALAPPDATA%\PJ\` directory removes PJ's local registry and settings. It does not remove the project directories registered in PJ.

## Project structure

```text
.
├── cmd/                  # Cobra commands
│   ├── root.go           # Root command
│   ├── add.go            # pj add
│   ├── jump.go           # pj jump
│   ├── list.go           # pj list
│   ├── remove.go         # pj remove
│   └── config.go         # pj config
├── internal/
│   ├── Config/           # Load and save configuration
│   └── store/            # Database setup and project queries
├── main.go
└── go.mod
```

The `cmd` package handles CLI commands and user-facing output. Database operations live in `internal/store`, while editor settings are handled by `internal/Config`.

## Tech stack

- [Cobra](https://github.com/spf13/cobra) — CLI commands and argument handling.
- [Turso Go driver](https://turso.tech) — embedded SQLite-compatible database through Go's `database/sql` package.

## Contributing

Issues and pull requests are welcome. When adding a command, follow the existing structure: create a command file in `cmd/`, define a `cobra.Command`, and attach it to the root command.

## License

Project Jumper is licensed under the [Apache License, Version 2.0](LICENSE).

You may use, modify, and redistribute the project, including in commercial software, subject to the terms of the license. See the license file for the full terms.

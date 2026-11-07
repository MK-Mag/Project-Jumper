# Project Jumper

A tiny Windows CLI that remembers where your projects live and opens them in your configured editor with one command.

Instead of typing \`cd C:\Users\you\Documents\work\some-project\` and then \`code .\`, you register the folder once under a short name and jump straight into it afterwards.

\`\`\`
> pj add api here
Saved project "api"
Path: C:\Users\you\Documents\work\api

> pj jump api
\`\`\`

## Features

- **Named project registry** — give every project a short, memorable alias.
- **\`here\` shortcut** — register the directory you are currently standing in without pasting its path.
- **One-word launch** — \`pj jump <name>\` opens the project in your configured editor.
- **Tab-aligned listing** — \`pj list\` shows name, path and creation time for everything you saved.
- **Project removal** — \`pj remove <name>\` removes a saved project from the registry without deleting its folder.
- **Configurable editor** — point it at \`code\`, \`idea64\`, \`subl\`, or any executable on your \`PATH\`.
- **Zero setup** — the data directory and database schema are created automatically on first use.

## Requirements

- Go 1.27 or newer (only needed to build from source)
- Windows (paths and application data are handled through \`%LOCALAPPDATA%\`)

## Installation

### Option 1: Download a release (recommended)

1. Open the [GitHub Releases](https://github.com/MK-Mag/Project-Jumper/releases) page.
2. Download \`pj.exe\` from the latest release, if a Windows executable is available.
3. Open PowerShell in the folder where you downloaded the executable and move it into a tools folder:

\`\`\`powershell
New-Item -ItemType Directory -Force "$env:LOCALAPPDATA\Programs" | Out-Null
Move-Item .\pj.exe "$env:LOCALAPPDATA\Programs\pj.exe"
\`\`\`

### Option 2: Build from source

Install Go 1.27 or newer, then clone and build the project:

\`\`\`powershell
git clone https://github.com/MK-Mag/Project-Jumper.git
cd Project-Jumper
go build -o pj.exe .
\`\`\`

Move the generated \`pj.exe\` into the tools folder:

\`\`\`powershell
New-Item -ItemType Directory -Force "$env:LOCALAPPDATA\Programs" | Out-Null
Move-Item .\pj.exe "$env:LOCALAPPDATA\Programs\pj.exe"
\`\`\`

### Add PJ to PATH

Moving the executable into \`%LOCALAPPDATA%\Programs\` does **not** automatically make the \`pj\` command available. Add that folder to your user PATH once.

Run this in PowerShell:

\`\`\`powershell
$programs = "$env:LOCALAPPDATA\Programs"
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")

if (($userPath -split ";") -notcontains $programs) {
    $newPath = if ([string]::IsNullOrWhiteSpace($userPath)) { $programs } else { "$userPath;$programs" }
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
}
\`\`\`

Close and reopen your terminal so it picks up the updated PATH.

Verify the installation:

\`\`\`powershell
pj --help
pj list
\`\`\`

If you do not want to edit PATH, run PJ using its full path, for example:

\`\`\`powershell
& "$env:LOCALAPPDATA\Programs\pj.exe" --help
\`\`\`

## Usage

### \`pj add <project_name> <project_path>\`

Registers a project. The path must already exist and must be a directory; it is stored as an absolute path, so relative paths are resolved for you.

\`\`\`bat
:: Register an existing folder by path
pj add api C:\Users\you\Documents\work\api

:: Register the folder you are currently in
cd C:\Users\you\Documents\work\api
pj add api here
\`\`\`

Project names are unique. Adding a name that already exists reports an error rather than overwriting the old entry.

### \`pj jump <project_name>\`

Looks up the project and launches your editor with its path as the working folder argument — roughly \`code C:\Users\you\Documents\work\api\`.

\`\`\`bat
pj jump api
\`\`\`

A child process cannot change your shell's working directory, so \`jump\` hands the folder to your editor instead of trying to \`cd\`. If you want the path for your own \`cd\`, read it from \`pj list\`.

### \`pj list\`

Prints every saved project, ordered by name:

\`\`\`
NAME  PATH                                    CREATED AT
api   C:\Users\you\Documents\work\api         2026-03-14 09:41:22
pj    C:\Users\you\Documents\my projects\...  2026-03-14 09:38:05
\`\`\`

### \`pj remove <project_name>\`

Removes a project from Project Jumper's registry.

**Important:** this only removes the saved project entry from PJ's database. It does **not** delete the actual project directory or any of its files.

\`\`\`bat
pj remove api
Removed project "api"
\`\`\`

If you want to use the project with PJ again, register it with \`pj add\`.

### \`pj config ide [name]\`

Shows or changes the editor used by \`pj jump\`. With no argument it prints the current value; with one it saves the new value. Defaults to \`code\`.

\`\`\`bat
:: Show the current editor
pj config ide
Current editor: code

:: Switch to another editor on your PATH
pj config ide idea64
\`\`\`

## Where things are stored

Everything lives under \`%LOCALAPPDATA%\PJ\\\`, which is created on first run:

| Path | Purpose |
| --- | --- |
| \`PJ\PJ.db\` | SQLite database holding the registered projects |
| \`PJ\config.json\` | Settings, currently the editor name |

The database contains a single \`projects\` table (\`id\`, \`name\`, \`path\`, \`created_at\`) and is created with \`CREATE TABLE IF NOT EXISTS\`, so it upgrades in place. Deleting the folder resets the tool completely.

\`config.json\` is plain JSON and safe to edit by hand:

\`\`\`json
{
  "IDE": "code"
}
\`\`\`

## Project layout

\`\`\`
.
├── cmd/                  Cobra commands
│   ├── root.go           Root command and entry point for Execute
│   ├── add.go            pj add
│   ├── jump.go           pj jump
│   ├── list.go           pj list
│   ├── remove.go         pj remove
│   └── config.go         pj config ide
├── internal/
│   ├── Config/           Load and save config.json
│   └── store/            Database opening, schema, and project queries
├── main.go
└── go.mod
\`\`\`

Commands only handle arguments and output; all persistence lives in \`internal/store\` and all settings access in \`internal/Config\`.

## Tech stack

- [Cobra](https://github.com/spf13/cobra) — command and flag structure
- [Turso Go driver](https://turso.tech) — embedded SQLite-compatible database via \`database/sql\`

## Contributing

Issues and pull requests are welcome. If you plan a change that adds a command, follow the existing pattern: one file in \`cmd/\`, a package-level \`cobra.Command\`, and an \`init()\` that attaches it to \`rootCmd\`.

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for the full text.

In short: you may use, modify and redistribute this project, including in commercial closed-source software, provided you keep a copy of the license with it, retain the existing attribution notices, and mark any files you change.

You can obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.

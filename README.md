# pinit
A personal note-taking app.

> [!Note]
> Work in progress.. `pinit` is currently under development.

## Usage
```console
$ pinit --help
A personal note-taking app

Usage:
  pinit [command]

Available Commands:
  init        Initialize pinit. Run database migration and setup pinit
  serve       Serve pinit

Flags:
      --help      Show help information
      --version   Show version

Use "pinit [command] --help" for more information about a command.
```

## How to develop
```bash
go generate ./...
go run . serve
```

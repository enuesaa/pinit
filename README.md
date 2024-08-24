# pinit
A personal note-taking app.

> [!Note]
> Work in progress.. `pinit` is currently under development.

## Usage
```console
$ pinit --help
A personal note-taking app

Usage:
  pinit [flags]

Flags:
      --help       Show help information
      --port int   port. Default: 3000 (default 3000)
      --serve      Start pinit app
      --version    Show version

$ pinit --serve
 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.52.4                   │
 │               http://127.0.0.1:3000               │
 │                                                   │
 │ Handlers ............ 13  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............. 58300 │
 └───────────────────────────────────────────────────┘
```

## How to develop
```bash
go generate ./...
go run . serve
```

## TODO
- [app] set openai api token on web
- [note] enable to add multiple notes to binder.

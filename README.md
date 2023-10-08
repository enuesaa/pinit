**Work in progress..**
# pinit
A personal note-taking app.

> **Note**
> Work in progress.. `pinit` is currently under development.

## Install
```bash
git clone https://github.com/enuesaa/pinit.git --depth 1
cd pinit
go install
```

## Usage
```console
$ pinit --help
A personal note-taking app.

Usage:
  pinit [command]

Available Commands:
  add         create a note
  edit        edit a note.
  lookup      lookup a note.
  ls          list notes.
  rm          remove a note
  setup       setup to take notes.

Flags:
      --help      Show help information
  -v, --version   version for pinit

Use "pinit [command] --help" for more information about a command.
```

### Setup database
For first time use, it is necessary to set up the database.
```bash
pinit setup configure --database-dsn <DATABASE_DSN>
pinit setup migration
```

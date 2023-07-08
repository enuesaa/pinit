# designdoc
## Commands

### List
```bash
$ pinit

FILENAME TAG
.gitignore <main>
.prettierrc.cjs <node>

$ pinit --tag node
```

### Register
```bash
$ pinit .gitignore
Not found.

To register, run following command:
  # create from local file.
  pinit .gitignore --tag main --register ./.gitignore

  # create from remote file
  pinit .gitignore --tag main --register https://raw.githubusercontent.com/example/example/blob/main/.gitignore
```

### Remove
```bash
$ pinit .gitignore --tag main --remove
```

### Apply
```bash
$ pinit .gitignore --tag main
$ pinit .gitignore --tag node
Not found.

Other tags are found.
FILENAME TAG
.gitignore <main>
.gitignore <go>

$ pinit .gitignore --tag main
Already Exists.

To override .gitignore, run following command:
  rm .gitignore
  pinit .gitignore --tag main
```

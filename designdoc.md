# designdoc

## コンセプト
- prettier とか eslint だとかの設定をローカル(もしくはリモート)に保持して、いつでもコピペできるようにしたい
- ユースケース(というか使うタイミング)は gh repo create に近い
- ただし必ずしもアプリケーションの作成時ではない

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
  pinit .gitignore --tag main --register
  pinit /something/path/.gitignore --tag main --register
  pinit https://raw.githubusercontent.com/example/example/blob/main/.gitignore --tag main --register
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

## 保存場所
```
~/.pinit/{tag}-{filename}
```

### Example
```
~/.pinit/main-.gitignore
~/.pinit/dev-index.js
```
### Memo
これテンプレートとして sh とかバイナリ置かれたらそのまま実行でき困るので、別のフォーマットにしたい

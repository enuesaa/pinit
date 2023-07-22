# designdoc

## コンセプト
- prettier とか eslint だとかの設定をローカル(もしくはリモート)に保持して、いつでもコピペできるようにしたい
- ユースケース(というか使うタイミング)は gh repo create に近い
- ただし必ずしもアプリケーションの作成時ではない

## Commands

### List
```bash
$ pinit list

FILENAME TAG
.gitignore <main>
.prettierrc.cjs <node>

$ pinit --tag node
```

### Register
```bash
$ pinit apply .gitignore
Not found.

To register, run following command:
  pinit register .gitignore --tag main
  pinit register /something/path/.gitignore --tag main
  pinit register https://raw.githubusercontent.com/example/example/blob/main/.gitignore --tag main
```

### Remove
```bash
$ pinit remove .gitignore --tag main
```

### Apply
```bash
$ pinit apply .gitignore --tag main
$ pinit apply .gitignore --tag node
Not found.

Other tags are found.
FILENAME TAG
.gitignore <main>
.gitignore <go>

$ pinit apply .gitignore --tag main
Already Exists.

To override .gitignore, run following command:
  rm .gitignore
  pinit apply .gitignore --tag main
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

# designdoc
## development plan
- [cli] 実装
- [web] ウェブアプリで作成・更新できるようにする
- [text] chatgptへの入出力をテンプレート化したい。タブを開いたら自動でAPIが呼ばれるイメージ
- [database] support sqlite
- [text] whisper api の呼び出しを go 経由にする
- [text] markdown
- [others] 可能であればスライドとか作成したい
- [others] 図

## dev command
### run dev server
```bash
go run -tags dev . serve
```

### build
```bash
go generate ./...
go build
```
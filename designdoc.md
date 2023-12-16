# designdoc
## development plan
- [web] ウェブアプリで作成・更新できるようにする
- [text] chatgptへの入出力をテンプレート化したい。タブを開いたら自動でAPIが呼ばれるイメージ
- [database] support sqlite
- [text] markdown
- [others] 可能であればスライドとか作成したい

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

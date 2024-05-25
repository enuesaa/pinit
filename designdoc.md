# designdoc

## 目的
- サーバにおき運用し始めたい
- 現状、ファイルシステムに保存しているデータがあり、冗長化したとき同期できなく問題となるのでなんとかしたい

## 改善案
- 環境変数にDBへの接続情報を入れる
  - PINIT_DB_HOST
  - PINIT_DB_USER
  - PINIT_DB_PASSWORD
- cofig は DB へ格納する

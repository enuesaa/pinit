# pinit
A personal note-taking app

## Stacks
- DynamoDB

## 構成
- バイナリにフロントエンドの dist を embed
- コマンドラインからそのまま立ち上げ可能

## データ設計
- 検索機能なし。データ設計する上で気になる要件はない
- Primary Key を Binder Name へ Sort Key を Note Name にすれば、ある程度耐えられると思う

## Future plan
- コマンド的なのを増やす
- DynamoDB のテーブル名を環境変数で指定できるよう
- DynamoDB のテーブルを作成できるよう (マイグレーション)
  - DynamoDB だけに依存する
  - 個人的には Terraform で問題ないがテーブル設定とアプリケーションが分散してしまい、環境の立ち上げが面倒

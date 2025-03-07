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

# pinit
A personal note-taking app

## Refactor Plan
- cloudnative に組み直す
- メモ帳 + 音声認識 (Whisper API) + LLM (ChatGPT) のようなアプリへ

## Stacks
- DynamoDB
- CloudFront
- Lambda

## 構成
- バイナリにフロントエンドの dist を embed しておりコマンドラインからそのまま立ち上げ可能
- バイナリを Lambda へデプロイし Web Adapter を介してホストする
- 認証機構はスコープ外

![構成図](./architecture.png)


## データ設計
- 検索機能なし。データ設計する上で気になる要件はない
- Primary Key を Binder Name へ Sort Key を Note Name にすれば、ある程度耐えられると思う

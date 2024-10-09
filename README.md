# Sample Go Application

## 環境構築方法

1. リポジトリをクローンします。
```sh
git clone <リポジトリのURL>
cd <リポジトリ名>
```

2. 必要な依存関係をインストールします。
```sh
go mod tidy
```

3. Dockerを使用してデータベースをセットアップします。
```sh
docker-compose up --build -d
```

## ディレクトリ構成
```
Sample Go Application
├── .gitignore
├── cmd/
│   └── server/
│       └── main.go  # アプリケーションのエントリーポイント
├── db/
│   └── migrations/
│       ├── 000001_create_samples_table.down.sql  # サンプルテーブルを削除するSQLスクリプト
│       └── 000001_create_samples_table.up.sql    # サンプルテーブルを作成するSQLスクリプト
├── docker-compose.yml  # Docker Compose設定ファイル
├── Dockerfile  # Dockerビルド設定ファイル
├── go.mod  # Goモジュール設定ファイル
├── go.sum  # Goモジュールの依存関係リスト
├── internal/
│   ├── infrastructure/
│   │   └── router/
│   │       └── router.go  # ルーティング設定
│   ├── sample/
│   │   ├── domain/
│   │   │   └── model/
│   │   │       └── sample.go  # サンプルドメインモデル
│   │   ├── infrastructure/
│   │   │   └── db/
│   │   │       └── sample_mysql.go  # MySQLに関するインフラストラクチャコード
│   │   ├── interface/
│   │   │   └── handler/
│   │   │       └── sample_handler.go  # HTTPハンドラー
│   │   ├── tests/
│   │   │   └── sample_test.go  # テストコード
│   │   └── usecase/
│   │       └── sample_usecase.go  # ユースケースロジック
│   └── todo/  # 例。他のモジュールがこのレイヤーに入る
└── README
```
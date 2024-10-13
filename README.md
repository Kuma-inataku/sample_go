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
/gin-api-sample                          # ルートディレクトリ
│
├── .gitignore
│
├── cmd/
│   └── server/
│       └── main.go  # アプリケーションのエントリーポイント
│
├── /src
│   ├── /sample                     # サンプルモジュール
│   │   ├── /sampleA                # sampleAに関するドメイン
│   │   │   ├── /domain             # ドメイン層
│   │   │   │   ├── /model          # エンティティ、値オブジェクト、集約
│   │   │   │   ├── /repository     # ドメインリポジトリのインターフェース
│   │   │   │   ├── /service        # ドメインサービス
│   │   │   │   └── /event          # ドメインイベント
│   │   │   ├── /application        # アプリケーション層
│   │   │   │   ├── /service        # アプリケーションサービス
│   │   │   │   └── /dto            # データ転送オブジェクト（DTO）
│   │   │   ├── /infrastructure     # インフラ層
│   │   │   │   ├── /repository     # ドメインリポジトリの実装
│   │   │   │   ├── /api            # 外部APIとの連携
│   │   │   │   └── /persistence    # 永続化層（DB接続やORM設定）
│   │   │   └── /interface          # インターフェース層
│   │   │       ├── /controller     # APIのエントリーポイント（REST/GraphQLなど）
│   │   │       └── /viewmodel      # ビューモデル
│   │   │
│   │   ├── /sampleB                # sampleBに関するドメイン
│   │   │   ├── /domain
│   │   │   ├── /application
│   │   │   ├── /infrastructure
│   │   │   └── /interface
│   │
│   ├── /task-manager                # タスク管理モジュール
│   │   ├── /worker                  # workerに関するドメイン
│   │   │   ├── /domain
│   │   │   ├── /application
│   │   │   ├── /infrastructure
│   │   │   └── /interface
│   │   ├── /manager                 # managerに関するドメイン
│   │   │   ├── /domain
│   │   │   ├── /application
│   │   │   ├── /infrastructure
│   │   │   └── /interface
│   │   ├── /task                    # taskに関するドメイン
│   │   │   ├── /domain
│   │   │   ├── /application
│   │   │   ├── /infrastructure
│   │   │   └── /interface
│   │
│   ├── /report-comment-fixer        # レポートコメント修正モジュール
│   │   ├── /worker                  # workerに関するドメイン
│   │   │   ├── /domain
│   │   │   ├── /application
│   │   │   ├── /infrastructure
│   │   │   └── /interface
│   │   ├── /manager                 # managerに関するドメイン
│   │   │   ├── /domain
│   │   │   ├── /application
│   │   │   ├── /infrastructure
│   │   │   └── /interface
│   │
│   ├── /shared                      # 共通モジュール（Shared）
│   │   ├── authentication/  # 認証関連の共通機能
│   │   ├── authorization/  # 認可関連の共通機能
│   │   ├── cache/  # キャッシュ関連の共通機能
│   │   ├── helper/  # ヘルパー関数
│   │   ├── logging/  # ロギング関連の共通機能
│   │   ├── notification/  # 通知関連の共通機能
│   │   └── payment/  # 支払い関連の共通機能
│   │
│   └── /infrastructure              # 全体に関わるインフラ層
│       ├── /docker                  # Docker関連ファイル
│       └── /database                # データベース関連（マイグレーションや設定）
│
├── /scripts                         # CI/CD用のスクリプト
│
├── /docs                            # ドキュメント
│
├── docker-compose.yml  # Docker Compose設定ファイル
│
├── Dockerfile  # Dockerビルド設定ファイル
│
├── go.mod  # Goモジュール設定ファイル
│
├── go.sum  # Goモジュールの依存関係リスト
│
└── README
```

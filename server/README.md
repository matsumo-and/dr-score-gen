# Dr Score Gen Server

Goのベストプラクティスに従ったクリーンアーキテクチャで構築されたAPIサーバー

## プロジェクト構造

```
server/
├── cmd/
│   └── api/
│       └── main.go              # アプリケーションのエントリーポイント
├── internal/                    # 外部に公開しない内部パッケージ
│   ├── domain/                  # ドメイン層
│   │   ├── entity/             # ビジネスエンティティ
│   │   └── repository/         # リポジトリインターフェース
│   ├── usecase/                # ユースケース層（ビジネスロジック）
│   ├── adapter/                # インターフェースアダプター層
│   │   ├── controller/         # HTTPコントローラー
│   │   ├── presenter/          # レスポンスプレゼンター（今後実装）
│   │   └── gateway/            # 外部サービスゲートウェイ（今後実装）
│   └── infrastructure/         # インフラストラクチャ層
│       ├── persistence/        # データ永続化の実装
│       │   └── memory/         # インメモリ実装
│       └── router/             # HTTPルーティング
└── pkg/                        # 外部に公開可能なパッケージ（今後実装）
```

## アーキテクチャの特徴

### 1. **依存性の方向**
- 外側から内側への依存のみ許可
- ドメイン層は他の層に依存しない
- インターフェースを使用した依存性の逆転

### 2. **各層の責務**

#### Domain層
- ビジネスエンティティとルールを定義
- 他の層に依存しない純粋なビジネスロジック
- `entity/`: エンティティの定義
- `repository/`: データアクセスのインターフェース

#### Usecase層
- アプリケーション固有のビジネスルール
- ドメイン層のエンティティとリポジトリを使用
- 外部の詳細（DB、Web）から独立

#### Adapter層
- 外部とのインターフェース変換
- `controller/`: HTTPリクエスト/レスポンスの処理
- `presenter/`: 出力データの整形
- `gateway/`: 外部APIとの通信

#### Infrastructure層
- フレームワークやツールの実装詳細
- `persistence/`: データベースアクセスの実装
- `router/`: HTTPルーティングの設定

## 実行方法

```bash
# サーバーの起動
go run cmd/api/main.go

# または
go build -o server cmd/api/main.go
./server
```

## API エンドポイント

### ヘルスチェック
- `GET /health` - サーバーの状態確認

### スコア管理
- `POST /api/v1/scores` - 新規スコア作成
- `GET /api/v1/scores` - 全スコア取得
- `GET /api/v1/scores/:id` - 特定のスコア取得
- `PUT /api/v1/scores/:id` - スコア更新
- `DELETE /api/v1/scores/:id` - スコア削除

## 今後の拡張

1. データベース実装の追加（PostgreSQL、MySQL等）
2. 認証・認可機能
3. ロギング・モニタリング
4. エラーハンドリングの改善
5. バリデーションの強化

# tcm-platform

東京音楽大学 - 練習室予約サービス

### 概要

Web ブラウザから利用可能な gRPC の技術検証と、Go におけるクリーンアーキテクチャ＋モジュラーモノリス設計の実装アウトプットを目的として開発。

東京音楽大学公式サイトの UX 向上を目的に、友人グループ内での運用を想定。

```mermaid
sequenceDiagram
    box 東京音楽大学公式サイト
        participant OfficialSite as 公式サイト
    end

    box tcm-platform
        participant User as ユーザー
        participant TCMApp as tcm-platform(アプリ)
        participant TCMDB as tcm-platform DB
        participant Batch as バッチ処理
    end

    %% ログイン・練習室一覧取得
    User->>TCMApp: ログインリクエスト
    TCMApp->>OfficialSite: スクレイピングでログイン認証
    OfficialSite-->>TCMApp: 認証結果
    TCMApp-->>User: ログイン成功

    User->>TCMApp: 練習室一覧取得
    TCMApp->>OfficialSite: スクレイピングで練習室一覧取得
    OfficialSite-->>TCMApp: 練習室一覧
    TCMApp-->>User: 練習室一覧表示

    %% 予約情報一覧取得
    User->>TCMApp: 予約情報一覧取得
    TCMApp->>TCMDB: 登録済み予約一覧取得
    TCMDB-->>TCMApp: 予約一覧
    TCMApp-->>User: 予約一覧表示

    %% 仮予約登録
    User->>TCMApp: 仮予約登録
    TCMApp->>TCMDB: 仮予約データ登録

    %% 予約削除
    User->>TCMApp: 予約削除
    TCMApp->>TCMDB: 予約データ削除

    %% 毎日12:00 バッチ処理
    Note over Batch: 毎日12:00実行
    Batch->>TCMDB: 当日の予約一覧取得
    loop 予約ごとに非同期登録
        Batch->>OfficialSite: スクレイピングで本予約登録
        OfficialSite-->>Batch: 登録結果
    end
    Batch->>TCMDB: 処理結果まとめて登録
```

###### アーキテクチャ

クリーンアーキテクチャ + モジュラーモノリス

###### 通信

gRPC（Connect）

###### フロントエンド

TypeScript / React Router v7 / React Aria / Tailwind CSS

###### バックエンド

Go / sqlc / PostgreSQL

###### スクレイピング

[ekkx/tcmrsv](https://github.com/ekkx/tcmrsv) を自作後、外部パッケージとして利用

###### デプロイ環境

自宅サーバー上の Docker コンテナ

# GEMINI.md

## 概要

- プロジェクト名: **noq**
- 目的: **ローカルの Markdown を唯一のソース**として、Notion(将来的に Qiita)へ安全に同期するための CLI ツール。

## スコープ

ブランチごとにタスクをtask/{コマンド名}.mdとして作成する。
各ブランチごとに作成されタスクを確認して、作業を行う

完了条件

- タスクの先頭に`[x]`があること
- `- [ ]`がある場合は未完了である

## コマンド(想定)

- init
- config: list / set
- status
- memo: new / list / add / reset / move-db / pull / push
- article: new / list / add / reset / pull / push(Qiitaは後続詳細)

## データ/ファイル仕様

### ローカルディレクトリ

- memo: `memo/yyyymmdd-{num}/yyyymmdd-{num}.md` + `assets/`
- article: memo と同型

### mdファイル内のYAML部分

- id
- kind
- title
- created_at
- updated_at
- destinations:
  - type: notion | qiita
  - db / tags / page_id / url など
- 方針: `destinations.properties` は **push で更新するプロパティのパッチのみ**を記載し、未記載は触らない。

## 設定ファイル

- パス: `~/.config/noq/config.yaml`(または `.yml`、最終確定要)
- 主なキー
  - memo_root(default: `~/noq/memo`)
  - article_root(default: `~/noq/article`)
  - notion:
    - token
    - default_memo_db
    - databases(alias -> DB ID マップ)
  - qiita:
    - token(未確定事項あり)
- 既存 config がある場合の挙動: 上書き禁止 or 確認 or マージ(要決定)
- 秘密情報の保存方針: 平文保存＋注意書き寄り(暗号化の是非は検討中)

## アーキテクチャ

- 採用: [Onion Architecture](docs\Architecture.md)
- 層構造
  - Domain: Entity / ValueObject / Domain Serlice
  - Application: Usecase / Repository IF / External IF
  - Presentation: Cobra CLI(cmd/args/dto)
  - Infrastructure: Repository/External の実装(FS/NotionAPI/QiitaAPI)
- 依存方向: 外 → 内(Presentation → Application → Domain)
- Infrastructure は IF 実装として内向き依存(Application/Domain の型・IFに依存)

## ディレクトリ構造

```
main.go
├─cmd
├─docs
├─infrastructure
│ └─repository
└─internal
    ├─application
    │ ├─repository
    │ ├─usecase
    │ └─validation
    ├─domain
    │ ├─entity
    │ ├─service
    │ └─valueobject
    └─presentation
    ├─dto
    ├─mapper
    └─prompt
```

## 依存性注入(Composition Root)

- Infrastructure の具体実装生成と Usecase への注入は **main.go を最外周(composition root)**として扱い、そこで組み立てる。
- Presentation(cmd 配下)は Infrastructure を import しない。
- Usecase は interface(例: `ConfigUseCase`)として Presentation に渡す。

## 実装技術

- Go
- Cobra(CLI)
- liper(設定読み込み)
- yaml.v3(frontmatter/設定の YAML)
- 標準ライブラリ(HTTP/JSON/FS)

## 実装メモ(現在の論点)

- ValueObject の値型/ポインタ型の統一方針
  - map の key にポインタを使うと同値でも引けない等の問題が出るため、辞書用途(alias->id)は値型/文字列寄せを検討
- コンストラクタで error を捨てない(panic 回避のため `(..., error)` を返す)
- CLI 入力: 任意型を更新したい場合は `parse func(string) (T, error)` を渡すジェネリクス関数で対応

## 注意

ファイルを生成する時に文字化け(�)することが多発するので出力時に再度確認してください。

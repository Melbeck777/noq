- [x] `noq init` の完了条件を定義する（ローカル運用優先）
  - memo_root があれば最低限動く
  - Notion設定は任意（未設定でもOK）

- [x] 設定ファイル名を `config.yml` に統一する
  - 生成先を `config.yml` に固定
  - 読み込みも `config.yml` を正とする

- [x] `init` 実行時の既存 `config.yml` 読み込みを実装する
  - 既存が無ければ、デフォルト値を読み込むようにする

- [ ] `init` の保存ロジックを「マージ（追加・更新）」にする
  - 未入力の既存値は保持
  - 未設定項目は出力しない（キー自体を書かない）

- [x] `notion.databases`（alias -> db_id）入力ループを実装する
  - alias / db_id を1件ずつ追加・更新
  - 削除は `init` ではしない

- [ ] `default_memo_db` を alias でのみ設定できるようにする
  - `notion.databases` の alias 一覧（ソート済み）から選択
  - `databases` が0件なら default_memo_db は未設定のまま

- [ ] YAML出力の整形（未設定は書かない）を実装する
  - `notion:` 自体も未設定なら出さない
  - `databases:` が空なら出さない
  - `default_memo_db` 未設定なら出さない
  - `qiita:` も未設定なら出さない

- [ ] `init` 実行後のメッセージ整備
  - 作成/更新したファイルパス
  - 次に必要な設定があれば案内（例: Notion連携は config set で追加できる、など）

- [ ] 最低限のテスト観点を用意する
  - 既存configなしで生成できる
  - 既存configありでマージできる（databasesが保持・追加・更新される）
  - databases=0件のとき default_memo_db を書かない

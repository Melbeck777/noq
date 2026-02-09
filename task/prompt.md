- [x] SelectInput を作る
  - [x] 候補（[]T）を受け取り、番号選択で 1 件返す
  - [x] 表示用フォーマット（例: "0 : xxx"）を統一する
  - [x] 不正入力（数値以外、範囲外、空入力）の再入力制御を入れる
  - [x] キャンセル手段を用意する（例: 0未満 でキャンセル、または空でキャンセル）
  - [x] 戻り値（selected T, cancelled bool, err error）を決めて実装する

- [x] MapKeySelect を作る
  - [x] map[string]string から keys を抽出して []string 化する
  - [x] keys を sort して表示順を安定させる
  - [x] MapKeySelect -> SelectInput の依存で選択を実装する
  - [x] 選択された key を返す（キャンセル時の扱いも含む）

- [x] テストを実装する（必須）
  - [x] SelectInput: 正常系（選択できる）
  - [x] SelectInput: 異常系（数値以外 / 範囲外 / 空入力の再入力）
  - [x] SelectInput: キャンセル系（キャンセル時の戻り値が期待通り）
  - [x] MapKeySelect: keys 抽出と sort が期待通り
  - [x] MapKeySelect: SelectInput への委譲で期待した key が返る

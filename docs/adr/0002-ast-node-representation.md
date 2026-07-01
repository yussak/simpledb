# 0002. AST ノードの表現方式

- ステータス: 採用
- 日付: 2026-06-27

## コンテキスト

SQL をパーサが読み取ると、「SELECT文で、カラムは name、テーブルは users」のような構造化されたデータ（AST: 抽象構文木）に変換される。実行器はこの AST を見て処理を行う。

SELECT 文・INSERT 文・CREATE TABLE 文はそれぞれ持つ情報が異なるため、Go の型システムでどう表現するかに設計の選択肢がある。この表現は全 Phase に波及するため、Phase 0 の着手前に決める。

## 検討した案

### 案A: interface + 具象構造体

全ノード共通のインターフェースを定義し、SQL 文ごとに別の構造体を作る方式。

```go
type Node interface {
    nodeType() string
}

type SelectStatement struct {
    Table   string
    Columns []string
}

type InsertStatement struct {
    Table   string
    Values  [][]Value
}
```

実行器では `switch node.(type)` で文の種類を判別して処理を分岐する。

- 利点:
  - Go で最も標準的なパターン。`go/ast`（Go 自身のパーサ）もこの方式。
  - 構造体ごとにフィールドが異なるので、SELECT に不要な Values フィールドが生えるといった問題が起きない。
  - WHERE や新しい文を追加するとき、既存コードを壊さず構造体を追加するだけで済む。
- 欠点:
  - 構造体の数が増える（ただし Phase 0 では 3 つのみ）。

### 案B: 種別フィールド付きの単一構造体

1 つの構造体に全 SQL 文のフィールドをまとめ、Kind フィールドで種類を区別する方式。

```go
type Statement struct {
    Kind    string   // "SELECT", "INSERT", "CREATE"
    Table   string
    Columns []string
    Values  [][]Value
}
```

- 利点:
  - 構造体が 1 つなので最初は楽。
- 欠点:
  - 文の種類によって「意味のないフィールド」が大量に生まれる（SELECT では Values を使わない等）。
  - Phase 1 で WHERE を足すと条件フィールドも増え、どのフィールドがどの Kind で有効かがコードから読み取れなくなる。
  - 規模が大きくなると破綻しやすい。

### 案C: map ベース（非構造化）

`map[string]interface{}` で自由にフィールドを持たせる方式。

```go
type Node map[string]interface{}
// node["kind"] = "SELECT"
// node["table"] = "users"
```

- 利点:
  - 最も柔軟。フィールドを自由に追加できる。
- 欠点:
  - 型安全性がゼロ。タイポしてもコンパイル時に気づけない。
  - 構造が見えづらく、RDBMS の内部構造を理解する学習目的には向かない。

## 決定

案A（interface + 具象構造体）を採用する。

主な判断:

- Go の標準ライブラリ自身がこの方式を採っており、Go らしい設計を学べる。
- 各 SQL 文が「何の情報を持つか」が構造体の定義を見れば分かり、RDBMS の構造を理解しやすい。
- Phase 1 以降で WHERE や DELETE を追加するとき、既存コードを壊さず拡張できる。

## 結果

- Phase 0 では `SelectStatement`・`InsertStatement`・`CreateTableStatement` の 3 構造体から始める。
- 共通インターフェース `Node` を定義し、実行器は型スイッチで分岐する。
- Phase 1 以降で文や式が増えた場合は、インターフェースを満たす構造体を追加する。

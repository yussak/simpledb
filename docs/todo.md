# TODO

## SELECT 文（縦に貫通させる）

- [x] レキサー: SELECT 文のトークン分解
- [x] パーサ: トークン列 → AST（SelectStatement）
- [x] メモリテーブル: テーブル定義とデータの保持
- [x] 実行器: AST を受け取り、メモリテーブルから結果を返す
- [x] 未実装文のスキップテスト追加（CREATE, INSERT）

## CREATE TABLE 文

- [ ] レキサー: CREATE TABLE 用トークン追加
- [ ] パーサ: CreateTableStatement
- [ ] 実行器: テーブル作成

## INSERT 文

- [ ] レキサー: INSERT 用トークン追加（INT・STRING リテラル含む）
- [ ] パーサ: InsertStatement
- [ ] 実行器: テーブルへの行追加
- [ ] ADR 0003: 内部値の表現方式を検討

## end-to-end 検証（Phase 0 ゴール）

- [ ] CREATE → INSERT → SELECT の一連の流れがテストで通る

package parser

import (
	"simpledb/lexer"
	"testing"
)

func TestSelectAsterisk(t *testing.T) {
	tokens := lexer.New("SELECT * FROM users;").Tokenize()
	node, err := New(tokens).Parse()
	if err != nil {
		t.Fatalf("パースエラー: %v", err)
	}

	stmt, ok := node.(*SelectStatement)
	if !ok {
		t.Fatalf("SelectStatement ではない: %T", node)
	}

	if len(stmt.Columns) != 1 || stmt.Columns[0] != "*" {
		t.Errorf("Columns: got=%v, want=[*]", stmt.Columns)
	}
	if stmt.Table != "users" {
		t.Errorf("Table: got=%q, want=%q", stmt.Table, "users")
	}
}

func TestSelectColumns(t *testing.T) {
	tokens := lexer.New("SELECT id, name FROM users").Tokenize()
	node, err := New(tokens).Parse()
	if err != nil {
		t.Fatalf("パースエラー: %v", err)
	}

	stmt, ok := node.(*SelectStatement)
	if !ok {
		t.Fatalf("SelectStatement ではない: %T", node)
	}

	if len(stmt.Columns) != 2 || stmt.Columns[0] != "id" || stmt.Columns[1] != "name" {
		t.Errorf("Columns: got=%v, want=[id, name]", stmt.Columns)
	}
	if stmt.Table != "users" {
		t.Errorf("Table: got=%q, want=%q", stmt.Table, "users")
	}
}

func TestSelectNoFrom(t *testing.T) {
	tokens := lexer.New("SELECT id").Tokenize()
	_, err := New(tokens).Parse()
	if err == nil {
		t.Fatal("FROM なしでエラーにならなかった")
	}
}

func TestCreateTableNotImplemented(t *testing.T) {
	t.Skip("未実装: CREATE TABLE のパース")
}

func TestInsertNotImplemented(t *testing.T) {
	t.Skip("未実装: INSERT のパース")
}

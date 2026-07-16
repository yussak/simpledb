package lexer

import "testing"

func TestSelectAsterisk(t *testing.T) {
	tokens := New("SELECT * FROM users;").Tokenize()

	expected := []Token{
		{Type: TOKEN_SELECT, Literal: "SELECT"},
		{Type: TOKEN_ASTERISK, Literal: "*"},
		{Type: TOKEN_FROM, Literal: "FROM"},
		{Type: TOKEN_IDENT, Literal: "users"},
		{Type: TOKEN_SEMICOLON, Literal: ";"},
		{Type: TOKEN_EOF, Literal: ""},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("トークン数が違う: got=%d, want=%d", len(tokens), len(expected))
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.Type || tokens[i].Literal != exp.Literal {
			t.Errorf("tokens[%d]: got={%d, %q}, want={%d, %q}",
				i, tokens[i].Type, tokens[i].Literal, exp.Type, exp.Literal)
		}
	}
}

func TestSelectColumns(t *testing.T) {
	tokens := New("SELECT id, name FROM users").Tokenize()

	expected := []Token{
		{Type: TOKEN_SELECT, Literal: "SELECT"},
		{Type: TOKEN_IDENT, Literal: "id"},
		{Type: TOKEN_COMMA, Literal: ","},
		{Type: TOKEN_IDENT, Literal: "name"},
		{Type: TOKEN_FROM, Literal: "FROM"},
		{Type: TOKEN_IDENT, Literal: "users"},
		{Type: TOKEN_EOF, Literal: ""},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("トークン数が違う: got=%d, want=%d", len(tokens), len(expected))
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.Type || tokens[i].Literal != exp.Literal {
			t.Errorf("tokens[%d]: got={%d, %q}, want={%d, %q}",
				i, tokens[i].Type, tokens[i].Literal, exp.Type, exp.Literal)
		}
	}
}

func TestKeywordCaseInsensitive(t *testing.T) {
	tokens := New("select * from users").Tokenize()

	if tokens[0].Type != TOKEN_SELECT {
		t.Errorf("小文字の select がキーワードとして認識されない: got=%d", tokens[0].Type)
	}
	if tokens[2].Type != TOKEN_FROM {
		t.Errorf("小文字の from がキーワードとして認識されない: got=%d", tokens[2].Type)
	}
}

func TestCreateTable(t *testing.T) {
	tokens := New("CREATE TABLE users (id INT, name TEXT);").Tokenize()

	expected := []Token{
		{Type: TOKEN_CREATE, Literal: "CREATE"},
		{Type: TOKEN_TABLE, Literal: "TABLE"},
		{Type: TOKEN_IDENT, Literal: "users"},
		{Type: TOKEN_OPEN, Literal: "("},
		{Type: TOKEN_IDENT, Literal: "id"},
		{Type: TOKEN_INT, Literal: "INT"},
		{Type: TOKEN_COMMA, Literal: ","},
		{Type: TOKEN_IDENT, Literal: "name"},
		{Type: TOKEN_TEXT, Literal: "TEXT"},
		{Type: TOKEN_CLOSE, Literal: ")"},
		{Type: TOKEN_SEMICOLON, Literal: ";"},
		{Type: TOKEN_EOF, Literal: ""},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("トークン数が違う: got=%d, want=%d", len(tokens), len(expected))
	}

	for i, exp := range expected {
		if tokens[i].Type != exp.Type || tokens[i].Literal != exp.Literal {
			t.Errorf("tokens[%d]: got={%d, %q}, want={%d, %q}",
				i, tokens[i].Type, tokens[i].Literal, exp.Type, exp.Literal)
		}
	}
}

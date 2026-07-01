package lexer

type TokenType int

const (
	TOKEN_ILLEGAL TokenType = iota
	TOKEN_EOF

	// SELECT 文で使うキーワード
	TOKEN_SELECT
	TOKEN_FROM

	// 識別子（テーブル名、カラム名）
	TOKEN_IDENT

	// 記号
	TOKEN_COMMA    // ,
	TOKEN_ASTERISK // *
	TOKEN_SEMICOLON // ;
)

type Token struct {
	Type    TokenType
	Literal string
}

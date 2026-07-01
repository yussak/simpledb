package parser

import (
	"fmt"
	"simpledb/lexer"
)

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() (Node, error) {
	tok := p.current()
	switch tok.Type {
	case lexer.TOKEN_SELECT:
		return p.parseSelect()
	default:
		return nil, fmt.Errorf("予期しないトークン: %q", tok.Literal)
	}
}

func (p *Parser) parseSelect() (*SelectStatement, error) {
	p.advance() // SELECT を消費

	columns, err := p.parseColumns()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.TOKEN_FROM); err != nil {
		return nil, err
	}

	table, err := p.expectIdent()
	if err != nil {
		return nil, err
	}

	// セミコロンがあれば消費する（なくてもよい）
	if p.current().Type == lexer.TOKEN_SEMICOLON {
		p.advance()
	}

	return &SelectStatement{Columns: columns, Table: table}, nil
}

func (p *Parser) parseColumns() ([]string, error) {
	if p.current().Type == lexer.TOKEN_ASTERISK {
		p.advance()
		return []string{"*"}, nil
	}

	var columns []string
	name, err := p.expectIdent()
	if err != nil {
		return nil, fmt.Errorf("カラム名が必要です: %w", err)
	}
	columns = append(columns, name)

	for p.current().Type == lexer.TOKEN_COMMA {
		p.advance() // , を消費
		name, err := p.expectIdent()
		if err != nil {
			return nil, fmt.Errorf("カラム名が必要です: %w", err)
		}
		columns = append(columns, name)
	}

	return columns, nil
}

func (p *Parser) current() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{Type: lexer.TOKEN_EOF}
	}
	return p.tokens[p.pos]
}

func (p *Parser) advance() {
	p.pos++
}

func (p *Parser) expect(expected lexer.TokenType) error {
	if p.current().Type != expected {
		return fmt.Errorf("期待: %d, 実際: %q", expected, p.current().Literal)
	}
	p.advance()
	return nil
}

func (p *Parser) expectIdent() (string, error) {
	if p.current().Type != lexer.TOKEN_IDENT {
		return "", fmt.Errorf("識別子が必要ですが %q が来ました", p.current().Literal)
	}
	lit := p.current().Literal
	p.advance()
	return lit, nil
}

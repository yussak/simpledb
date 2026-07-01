package parser

// ADR-0002: interface + 具象構造体方式

type Node interface {
	nodeType() string
}

type SelectStatement struct {
	Columns []string // カラム名のリスト。"*" の場合は ["*"]
	Table   string
}

func (s *SelectStatement) nodeType() string { return "SELECT" }

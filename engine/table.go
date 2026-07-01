package engine

import "fmt"

type Table struct {
	Columns []string
	Rows    [][]string
}

type Database struct {
	tables map[string]*Table
}

func NewDatabase() *Database {
	return &Database{tables: make(map[string]*Table)}
}

func (db *Database) AddTable(name string, table *Table) {
	db.tables[name] = table
}

func (db *Database) GetTable(name string) (*Table, error) {
	t, ok := db.tables[name]
	if !ok {
		return nil, fmt.Errorf("テーブル %q が見つかりません", name)
	}
	return t, nil
}

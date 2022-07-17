package builder

import (
	"fmt"
	"strings"
)

type SQLKeyword string

const (
	Select SQLKeyword = `SELECT`
	From   SQLKeyword = `FROM`
	Where  SQLKeyword = `WHERE`
)

var mapPatterns = map[SQLKeyword]string{
	Select: "SELECT %s", // SELECT [column,...]
	From:   "FROM %s",   // FROM table
	Where:  "WHERE %s",  // WHERE predicates
}

type Builder struct {
	statment string
	table    *string
	where    *string
}

func (b *Builder) Select(c ...string) *table {
	columns := "*"

	if formatedColumns := strings.Join(c, ", "); len(formatedColumns) > 0 {
		columns = formatedColumns
	}

	b.statment = fmt.Sprintf(mapPatterns[Select], columns)
	return &table{b}
}

type table struct {
	*Builder
}

func (t *table) From(name string) *where {
	if len(name) > 0 {
		table := fmt.Sprintf(mapPatterns[From], name)
		t.table = &table
	}

	return &where{t.Builder}
}

type where struct {
	*Builder
}

type predicateGen interface {
	Sql() string
}

func (w *where) Where(p predicateGen) *Builder {
	sqlPredictions := p.Sql()
	if sqlPredictions != "" {
		raw := fmt.Sprintf(mapPatterns[Where], sqlPredictions)
		w.where = &raw
	}
	return w.Builder
}

func (b *Builder) Build() string {
	var opt string
	if b.table != nil {
		opt += " " + *b.table
	}
	if b.where != nil {
		opt += " " + *b.where
	}
	return b.statment + opt
}

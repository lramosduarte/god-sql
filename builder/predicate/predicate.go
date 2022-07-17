package predicate

import "fmt"

type keyCombine string

const (
	and  keyCombine = "AND"
	or   keyCombine = "OR"
	none keyCombine = ""
)

type Operator string

const (
	Equal        Operator = "="
	NotEqual     Operator = "<>"
	Greater      Operator = ">"
	GreaterEqual Operator = ">="
	Less         Operator = "<"
	LessEqual    Operator = "<="
	In           Operator = "IN"
	Between      Operator = "BETWEEN"
	Like         Operator = "LIKE"
	IsNull       Operator = "IS NULL"
	IsNotNull    Operator = "IS NOT NULL"
)

type predicate struct {
	keyCombine keyCombine
	column     string
	operator   Operator
	value      string
}

func (p *predicate) Sql() string {
	var sqlCombiner string
	if p.keyCombine != none {
		sqlCombiner = fmt.Sprintf("%s ", p.keyCombine)
	}
	return fmt.Sprintf("%s%s %s %s", sqlCombiner, p.column, p.operator, p.value)
}

type Predicates struct {
	predicates []*predicate
}

func (ps *Predicates) And(c string, o Operator, v string) *Predicates {
	p := &predicate{
		keyCombine: and,
		column:     c,
		operator:   o,
		value:      v,
	}
	ps.add(p)
	return ps
}

func (ps *Predicates) Or(c string, o Operator, v string) *Predicates {
	p := &predicate{
		keyCombine: or,
		column:     c,
		operator:   o,
		value:      v,
	}
	ps.add(p)
	return ps
}

func (ps *Predicates) add(p *predicate) {
	if len(ps.predicates) == 0 {
		p.keyCombine = none
	}
	ps.predicates = append(ps.predicates, p)
}

func (ps Predicates) Sql() string {
	var sqlPredicates string
	for i, p := range ps.predicates {
		if i > 0 {
			sqlPredicates += " "
		}
		sqlPredicates += p.Sql()
	}
	return sqlPredicates
}

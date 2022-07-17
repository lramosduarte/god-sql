package builder

import (
	"testing"

	"github.com/lramosduarte/god-sql/builder/predicate"
)

func TestBuilder_SelectColumnsValues(t *testing.T) {
	columnsTestCases := []struct {
		Description string
		Want        string
		Params      []string
	}{
		{
			Description: "build a query without set columns, returing a string with `*` for all columns",
			Want:        "SELECT *",
			Params:      nil,
		},
		{
			Description: "build a query with string of colums, returing a string with all columns separed by `,`",
			Want:        "SELECT c1, c2, c3",
			Params:      []string{"c1", "c2", "c3"},
		},
	}

	for _, testcase := range columnsTestCases {
		b := Builder{}
		got := b.Select(testcase.Params...).Build()
		if testcase.Want != got {
			t.Error("diff values", testcase.Want, got)
		}
	}
}

func TestTable_From(t *testing.T) {
	columnsTestCases := []struct {
		Description string
		Want        string
		TableName   string
	}{
		{
			Description: "build a query without set tablename, returing a empty string",
			Want:        "",
			TableName:   "",
		},
		{
			Description: "build a query passing a table name, returing a string with `from` clause",
			Want:        " FROM foo_bar",
			TableName:   "foo_bar",
		},
	}

	for _, testcase := range columnsTestCases {
		tb := table{Builder: &Builder{}}
		got := tb.From(testcase.TableName).Build()
		if testcase.Want != got {
			t.Error("diff values", testcase.Want, got)
		}
	}
}

func TestTable_Where(t *testing.T) {
	columnsTestCases := []struct {
		Description string
		Want        string
		Predicates  predicate.Predicates
	}{
		{
			Description: "build a query without set predicate, returing a empty string",
			Want:        "",
			Predicates:  predicate.Predicates{},
		},
		{
			Description: "build a query passing a list of predicates name, return a string with a `where` clause",
			Want:        " WHERE c1 = $1",
			Predicates: func() predicate.Predicates {
				ps := predicate.Predicates{}
				ps.And("c1", predicate.Equal, "$1")
				return ps
			}(),
		},
	}

	for _, testcase := range columnsTestCases {
		tb := where{Builder: &Builder{}}
		got := tb.Where(testcase.Predicates).Build()
		if testcase.Want != got {
			t.Error("diff values", testcase.Want, got)
		}
	}
}

func TestTable_Build(t *testing.T) {
	columnsTestCases := []struct {
		Description string
		Want        string
		fn          func() string
	}{
		{
			Description: "build a query without table and where clause, returning an select ",
			Want:        "SELECT now()",
			fn: func() string {
				b := Builder{}
				return b.Select("now()").Build()
			},
		},
		{
			Description: "build a query without where clause, returning an select without filters",
			Want:        "SELECT * FROM foo_bar",
			fn: func() string {
				b := Builder{}
				return b.Select().From("foo_bar").Build()
			},
		},
		{
			Description: "build a common query with filter, returning an common sql",
			Want:        "SELECT * FROM foo_bar WHERE c1 = $1 AND c2 > $2",
			fn: func() string {
				b := Builder{}
				ps := predicate.Predicates{}
				ps.And("c1", predicate.Equal, "$1").And("c2", predicate.Greater, "$2")
				return b.Select().From("foo_bar").Where(ps).Build()
			},
		},
	}

	for _, testcase := range columnsTestCases {
		got := testcase.fn()
		if testcase.Want != got {
			t.Error("diff values", testcase.Want, got)
		}
	}
}

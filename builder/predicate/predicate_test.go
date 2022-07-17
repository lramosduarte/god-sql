package predicate

import (
	"testing"
)

func TestPredicateIndividual_Sql(t *testing.T) {
	testCases := []struct {
		Description string
		Want        string
		P           predicate
	}{
		{
			Description: "generating equal predicate",
			Want:        "c1 = $1",
			P:           predicate{keyCombine: none, column: "c1", operator: Equal, value: "$1"},
		},
		{
			Description: "generating not equal predicate",
			Want:        "c1 <> $1",
			P:           predicate{keyCombine: none, column: "c1", operator: NotEqual, value: "$1"},
		},
		{
			Description: "generating less predicate",
			Want:        "c1 < $1",
			P:           predicate{keyCombine: none, column: "c1", operator: Less, value: "$1"},
		},
		{
			Description: "generating less equal predicate",
			Want:        "c1 <= $1",
			P:           predicate{keyCombine: none, column: "c1", operator: LessEqual, value: "$1"},
		},
		{
			Description: "generating greater predicate",
			Want:        "c1 > $1",
			P:           predicate{keyCombine: none, column: "c1", operator: Greater, value: "$1"},
		},
		{
			Description: "generating greater equal predicate",
			Want:        "c1 >= $1",
			P:           predicate{keyCombine: none, column: "c1", operator: GreaterEqual, value: "$1"},
		},
		{
			Description: "generating in predicate",
			Want:        "c1 IN $1",
			P:           predicate{keyCombine: none, column: "c1", operator: In, value: "$1"},
		},
		{
			Description: "generating between predicate",
			Want:        "c1 BETWEEN $1",
			P:           predicate{keyCombine: none, column: "c1", operator: Between, value: "$1"},
		},
		{
			Description: "generating is_null predicate",
			Want:        "c1 IS NULL $1",
			P:           predicate{keyCombine: none, column: "c1", operator: IsNull, value: "$1"},
		},
		{
			Description: "generating is_not_null predicate",
			Want:        "c1 IS NOT NULL $1",
			P:           predicate{keyCombine: none, column: "c1", operator: IsNotNull, value: "$1"},
		},
		{
			Description: "generating like predicate",
			Want:        "c1 LIKE $1",
			P:           predicate{keyCombine: none, column: "c1", operator: Like, value: "$1"},
		},
		{
			Description: "generating like predicate with combiner keyword",
			Want:        "AND c1 LIKE $1",
			P:           predicate{keyCombine: and, column: "c1", operator: Like, value: "$1"},
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.Description, func(t *testing.T) {
			got := testcase.P.Sql()
			if testcase.Want != got {
				t.Error("diff sql values", testcase.Want, got)
			}
		})
	}
}

func TestPredicates_addFirstElement(t *testing.T) {
	t.Run("add only one predication, value of key combine is removed", func(t *testing.T) {
		ps := Predicates{}
		want := &predicate{keyCombine: and}
		ps.add(want)
		if ps.predicates[0].keyCombine != none {
			t.Errorf("the first object should be none in combine key, but given %v", ps.predicates[0])
		}
	})
}

func stubPredicate() *predicate {
	return &predicate{keyCombine: none}
}

func TestPredicates_And(t *testing.T) {
	t.Run("And prediction add in list of prediction, check second element because the if clause", func(t *testing.T) {
		want := predicate{keyCombine: and, column: "c1", operator: Equal, value: "$1"}
		ps := Predicates{}
		ps.add(stubPredicate())
		ps.And(want.column, want.operator, want.value)
		got := *ps.predicates[1]
		if want.keyCombine != got.keyCombine ||
			want.column != got.column ||
			want.operator != got.operator ||
			want.value != got.value {
			t.Errorf("the first object should be none in combine key, but given %v", ps.predicates[0])
		}
	})
}

func TestPredicates_Or(t *testing.T) {
	t.Run("OR prediction add in list of prediction, check second element because the if clause", func(t *testing.T) {
		want := predicate{keyCombine: or, column: "c1", operator: Equal, value: "$1"}
		ps := Predicates{}
		ps.add(stubPredicate())
		ps.Or(want.column, want.operator, want.value)
		got := *ps.predicates[1]
		if want.keyCombine != got.keyCombine ||
			want.column != got.column ||
			want.operator != got.operator ||
			want.value != got.value {
			t.Errorf("the first object should be none in combine key, but given %v", ps.predicates[0])
		}
	})
}

func TestPredicates_add(t *testing.T) {
	testcases := []struct {
		Description string
		Want        string
		ps          []*predicate
	}{
		{
			Description: "add nothing and returing empty string",
			Want:        "",
			ps:          []*predicate{},
		},
		{
			Description: "add 2 `and` predicates and returing sql",
			Want:        "c1 = $1 AND c2 = $2",
			ps: []*predicate{
				{keyCombine: and, column: "c1", operator: Equal, value: "$1"},
				{keyCombine: and, column: "c2", operator: Equal, value: "$2"},
			},
		},
		{
			Description: "add 2 `and` with `or` predicates and returing sql",
			Want:        "c1 = $1 OR c2 = $2 AND c3 = $3",
			ps: []*predicate{
				{keyCombine: and, column: "c1", operator: Equal, value: "$1"},
				{keyCombine: or, column: "c2", operator: Equal, value: "$2"},
				{keyCombine: and, column: "c3", operator: Equal, value: "$3"},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Description, func(t *testing.T) {
			ps := Predicates{}
			for _, p := range testcase.ps {
				ps.add(p)
			}
			got := ps.Sql()
			if testcase.Want != got {
				t.Error("diff values", testcase.Want, got)
			}
		})
	}
}

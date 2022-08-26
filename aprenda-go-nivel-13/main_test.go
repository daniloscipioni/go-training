package main

import "testing"

type testsoma struct {
	values []int
	soma   int
}

type testmult struct {
	values []int
	mult   int
}

var testsSoma = []testsoma{
	{[]int{1, 3, 5}, 9},
	{[]int{2, 2, 2}, 6},
	{[]int{6, 1, 3}, 10},
}

var testsMult = []testmult{
	{[]int{1, 1, 1}, 1},
	{[]int{1, 2, 5}, 10},
	{[]int{2, 3, 7}, 42},
}

func TestSoma(t *testing.T) {
	for _, v := range testsSoma {
		value := Soma(v.values...)

		if value != v.soma {
			t.Error("For ", v.values,
				"expected", v.soma,
				"Got", value,
			)
		}
	}
}

func TestMultiplicacao(t *testing.T) {
	var value int

	for _, v := range testsMult {
		value = Multiplicacao(v.values...)

		if value != v.mult {
			t.Error("For", v.values,
				"Expected", v.mult,
				"Got:", value)
		}

	}

	value = Multiplicacao(1, 3)

	if value != 3 {
		t.Error("Expected 3, got ", value)
	}
}

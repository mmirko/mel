package m3number

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3numberEvaluator(t *testing.T) {

	a := new(M3numberMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = false
	a.MelInit(c, ep)

	tests := []string{"m3numberconst(45.2)", "m3numberconst(45.2)"}
	tests = append(tests, "add(m3numberconst(1E+1),m3numberconst(2))", "m3numberconst(1.2E+01)")
	tests = append(tests, "mult(m3numberconst(3.2),m3numberconst(5))", "m3numberconst(1.6E+01)")

	for i, iString := range tests {

		if i%2 == 1 {
			continue
		}

		a.MelStringImport(iString)
		a.Compute()
		if a.Inspect() != tests[i+1] {
			t.Errorf("Expected %s, got %s", tests[i+1], a.Inspect())
		}

	}

	b := new(M3numberMe3li)
	b.MelInit(c, ep)
	b.MelStringImport("div(m3numberconst(1),m3numberconst(0))")
	if err := b.Compute(); err == nil || err.Error() != "division by zero" {
		t.Errorf("Expected division by zero, got %v", err)
	}
}

func TestM3numberOperations(t *testing.T) {
	var ep *mel.EvolutionParameters
	config := &mel.MelConfig{Debug: false}

	operations := []struct {
		name      string
		programID uint16
		calculate func(float32, float32) float32
	}{
		{name: "add", programID: ADD, calculate: func(left, right float32) float32 { return left + right }},
		{name: "sub", programID: SUB, calculate: func(left, right float32) float32 { return left - right }},
		{name: "mult", programID: MULT, calculate: func(left, right float32) float32 { return left * right }},
		{name: "div", programID: DIV, calculate: func(left, right float32) float32 { return left / right }},
	}

	for _, operation := range operations {
		for index := 1; index <= 100; index++ {
			t.Run(fmt.Sprintf("%s/%03d", operation.name, index), func(t *testing.T) {
				left := float32(index)
				right := float32(index%9 + 1)
				expression := fmt.Sprintf("%s(m3numberconst(%g),m3numberconst(%g))", operation.name, left, right)

				prog := new(M3numberMe3li)
				prog.MelInit(config, ep)
				if err := prog.MelStringImport(expression); err != nil {
					t.Fatalf("MelStringImport(%q): %v", expression, err)
				}
				if err := prog.Compute(); err != nil {
					t.Fatalf("Compute(%q): %v", expression, err)
				}

				want := "m3numberconst(" + strconv.FormatFloat(float64(operation.calculate(left, right)), 'E', -1, 32) + ")"
				if got := prog.Inspect(); got != want {
					t.Errorf("%s: expected %s, got %s", expression, want, got)
				}
			})
		}
	}
}

func TestM3numberChainedOperations(t *testing.T) {
	var ep *mel.EvolutionParameters
	config := &mel.MelConfig{Debug: false}
	tests := []struct {
		name       string
		expression string
		want       float32
	}{
		{
			name:       "level-1/add",
			expression: "add(m3numberconst(2),m3numberconst(3))",
			want:       5,
		},
		{
			name:       "level-2/sub-mult",
			expression: "sub(add(m3numberconst(9),m3numberconst(6)),mult(m3numberconst(2),m3numberconst(3)))",
			want:       9,
		},
		{
			name:       "level-2/mult-add-sub",
			expression: "mult(add(m3numberconst(1),m3numberconst(2)),sub(m3numberconst(8),m3numberconst(5)))",
			want:       9,
		},
		{
			name:       "level-2/div-add-sub",
			expression: "div(add(m3numberconst(6),m3numberconst(6)),sub(m3numberconst(5),m3numberconst(2)))",
			want:       4,
		},
		{
			name:       "level-3/add",
			expression: "add(mult(add(m3numberconst(1),m3numberconst(2)),sub(m3numberconst(8),m3numberconst(5))),div(m3numberconst(20),m3numberconst(4)))",
			want:       14,
		},
		{
			name:       "level-3/sub",
			expression: "sub(div(mult(m3numberconst(3),m3numberconst(4)),add(m3numberconst(2),m3numberconst(2))),m3numberconst(1))",
			want:       2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prog := new(M3numberMe3li)
			prog.MelInit(config, ep)
			if err := prog.MelStringImport(test.expression); err != nil {
				t.Fatalf("MelStringImport(%q): %v", test.expression, err)
			}
			if err := prog.Compute(); err != nil {
				t.Fatalf("Compute(%q): %v", test.expression, err)
			}

			want := "m3numberconst(" + strconv.FormatFloat(float64(test.want), 'E', -1, 32) + ")"
			if got := prog.Inspect(); got != want {
				t.Errorf("expected %s, got %s", want, got)
			}
		})
	}
}

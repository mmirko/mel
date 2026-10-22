package m3numberext

import (
	"strconv"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3numberExtOperations(t *testing.T) {
	var parameters *mel.EvolutionParameters
	config := &mel.MelConfig{Debug: false}
	tests := []struct {
		name       string
		expression string
		want       float32
	}{
		{name: "sqrt", expression: "sqrt(m3numberconst(9))", want: 3},
		{name: "pow", expression: "pow(m3numberconst(2),m3numberconst(3))", want: 8},
		{name: "exp", expression: "exp(m3numberconst(0))", want: 1},
		{name: "log", expression: "log(m3numberconst(2.7182817))", want: 0.99999994},
		{name: "log2", expression: "log2(m3numberconst(8))", want: 3},
		{name: "log10", expression: "log10(m3numberconst(1000))", want: 3},
		{name: "abs", expression: "abs(m3numberconst(-2.5))", want: 2.5},
		{name: "ceil", expression: "ceil(m3numberconst(1.2))", want: 2},
		{name: "floor", expression: "floor(m3numberconst(1.8))", want: 1},
		{name: "round", expression: "round(m3numberconst(1.236),m3uintconst(2))", want: 1.24},
		{name: "trunc", expression: "trunc(m3numberconst(-1.8))", want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			program := new(M3numberExtMe3li)
			program.MelInit(config, parameters)
			if err := program.MelStringImport(test.expression); err != nil {
				t.Fatalf("MelStringImport(%q): %v", test.expression, err)
			}
			if err := program.Compute(); err != nil {
				t.Fatalf("Compute(%q): %v", test.expression, err)
			}

			want := "m3numberconst(" + strconv.FormatFloat(float64(test.want), 'E', -1, 32) + ")"
			if got := program.Inspect(); got != want {
				t.Errorf("expected %s, got %s", want, got)
			}
		})
	}
}

func TestM3numberExtDomains(t *testing.T) {
	for _, expression := range []string{"sqrt(m3numberconst(-1))", "log(m3numberconst(0))"} {
		t.Run(expression, func(t *testing.T) {
			program := new(M3numberExtMe3li)
			program.MelInit(&mel.MelConfig{Debug: false}, nil)
			if err := program.MelStringImport(expression); err != nil {
				t.Fatalf("MelStringImport(%q): %v", expression, err)
			}
			if err := program.Compute(); err == nil || err.Error() != "input outside function domain" {
				t.Fatalf("expected domain error, got %v", err)
			}
		})
	}
}

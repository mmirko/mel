package m3numbertriginv

import (
	"math"
	"strconv"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3numberTrigInverseOperations(t *testing.T) {
	var parameters *mel.EvolutionParameters
	config := &mel.MelConfig{Debug: false}
	tests := []struct {
		name       string
		expression string
		want       float32
	}{
		{name: "asin", expression: "asin(m3numberconst(0.5))", want: float32(math.Pi / 6)},
		{name: "acos", expression: "acos(m3numberconst(0.5))", want: float32(math.Pi / 3)},
		{name: "atan", expression: "atan(m3numberconst(1))", want: float32(math.Pi / 4)},
		{name: "acsc", expression: "acsc(m3numberconst(2))", want: float32(math.Pi / 6)},
		{name: "asec", expression: "asec(m3numberconst(2))", want: float32(math.Pi / 3)},
		{name: "acot-positive", expression: "acot(m3numberconst(1))", want: float32(math.Pi / 4)},
		{name: "acot-negative", expression: "acot(m3numberconst(-1))", want: float32(3 * math.Pi / 4)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			program := new(M3numberTrigInvMe3li)
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

func TestM3numberTrigInverseDomain(t *testing.T) {
	tests := []string{
		"asin(m3numberconst(2))",
		"acos(m3numberconst(-2))",
		"acsc(m3numberconst(0.5))",
		"asec(m3numberconst(-0.5))",
	}
	for _, expression := range tests {
		t.Run(expression, func(t *testing.T) {
			program := new(M3numberTrigInvMe3li)
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

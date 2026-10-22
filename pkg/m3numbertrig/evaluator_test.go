package m3numbertrig

import (
	"strconv"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3numberTrigOperations(t *testing.T) {
	var parameters *mel.EvolutionParameters
	config := &mel.MelConfig{Debug: false}
	tests := []struct {
		name       string
		expression string
		want       float32
	}{
		{name: "sin", expression: "sin(m3numberconst(0))", want: 0},
		{name: "cos", expression: "cos(m3numberconst(0))", want: 1},
		{name: "tan", expression: "tan(m3numberconst(0))", want: 0},
		{name: "csc", expression: "csc(m3numberconst(1.5707963267948966))", want: 1},
		{name: "sec", expression: "sec(m3numberconst(0))", want: 1},
		{name: "cot", expression: "cot(m3numberconst(0.7853981633974483))", want: 0.99999994},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			program := new(M3numberTrigMe3li)
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

func TestM3numberTrigReciprocalZero(t *testing.T) {
	var parameters *mel.EvolutionParameters
	program := new(M3numberTrigMe3li)
	program.MelInit(&mel.MelConfig{Debug: false}, parameters)
	if err := program.MelStringImport("csc(m3numberconst(0))"); err != nil {
		t.Fatalf("MelStringImport: %v", err)
	}
	if err := program.Compute(); err == nil || err.Error() != "division by zero" {
		t.Fatalf("expected division by zero, got %v", err)
	}
}

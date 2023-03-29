package envfloat

import "errors"

type EnvFloat struct {
	inVars   []float32
	keepVars []float32
	outVars  []float32
}

func (env *EnvFloat) Init(inputs []float32, outputs int) {
	copy(env.inVars, inputs)
	env.keepVars = make([]float32, 0)
	env.outVars = make([]float32, outputs)
}

func (env *EnvFloat) ReadInput(num int) (float32, error) {
	if num < len(env.inVars) {
		return env.inVars[num], nil
	}
	return 0, errors.New("input index out of range")
}

func (env *EnvFloat) WriteOutput(num int, value float32) error {
	if num < len(env.outVars) {
		env.outVars[num] = value
		return nil
	}
	return errors.New("output index out of range")
}

func (env *EnvFloat) ReadVar(num int) (float32, error) {
	if num < len(env.keepVars) {
		return env.keepVars[num], nil
	}
	return 0, errors.New("var index out of range")
}

func (env *EnvFloat) WriteVar(num int, value float32) error {
	if num < len(env.keepVars) {
		env.keepVars[num] = value
		return nil
	}
	return errors.New("var index out of range")
}

func (env *EnvFloat) PushVar(value float32) {
	env.keepVars = append(env.keepVars, value)
}

func (env *EnvFloat) PopVar() (float32, error) {
	if len(env.keepVars) > 0 {
		value := env.keepVars[len(env.keepVars)-1]
		env.keepVars = env.keepVars[:len(env.keepVars)-1]
		return value, nil
	}
	return 0, errors.New("no vars to pop")
}

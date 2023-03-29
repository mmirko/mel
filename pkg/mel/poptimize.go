package mel

import (
	"fmt"
	"sort"
)

type OptimizerValues []float32

// Info on each generation
type OptimizerInfo map[string]OptimizerValues

func (ri *OptimizerInfo) dumpOptimizerInfo() string {
	result := ""

	// Order the keys
	keys := make([]string, 0)
	for k := range *ri {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := (*ri)[key]
		result += fmt.Sprintf(" %s: [%f]", key, value[len(value)-1])
	}
	return result
}

func (ri *OptimizerInfo) InsertOptimizerInfo(key string, value float32) {
	var r OptimizerInfo
	if ri != nil {
		r = *ri
	}
	if r == nil {
		r = make(map[string]OptimizerValues)
	}
	if slice, ok := r[key]; ok {
		r[key] = append(slice, value)
	} else {
		newOptimizerInfoValues := make(OptimizerValues, 1)
		newOptimizerInfoValues[0] = value
		r[key] = newOptimizerInfoValues
	}
	*ri = r
}

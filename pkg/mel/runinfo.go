package mel

import (
	"fmt"
	"sort"
)

type RunInfoValues []float32

// Info on each generation
type RunInfo map[string]RunInfoValues

func (ri *RunInfo) dumpRunInfoLatest() string {
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

func (ri *RunInfo) InsertRunInfo(key string, value float32) {
	var r RunInfo
	if ri != nil {
		r = *ri
	}
	if r == nil {
		r = make(map[string]RunInfoValues)
	}
	if slice, ok := r[key]; ok {
		r[key] = append(slice, value)
	} else {
		newRunInfoValues := make(RunInfoValues, 1)
		newRunInfoValues[0] = value
		r[key] = newRunInfoValues
	}
	*ri = r
}

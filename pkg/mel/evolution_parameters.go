package mel

import (
	"errors"
	"strconv"
	"strings"
)

type EvolutionParameters struct {
	Pars map[string]string
}

func (ep *EvolutionParameters) GetMatchingList(match string) (map[string]string, bool) {

	result := make(map[string]string)

	for key, value := range ep.Pars {
		if strings.HasPrefix(key, match) {
			result[key[len(match):]] = value
		}
	}

	if len(result) > 0 {
		return result, true
	}

	return result, false
}

func (ep *EvolutionParameters) GetValue(param string) (string, bool) {
	if result, ok := ep.Pars[param]; ok {
		return result, true
	}
	return "", false
}

func (ep *EvolutionParameters) GetInt(param string) (int, bool) {
	if result, ok := ep.Pars[param]; ok {

		// Convert result to int
		if result_int, ok := strconv.Atoi(result); ok == nil {
			return result_int, true
		}
	}
	return 0, false
}

func (ep *EvolutionParameters) SetValue(param string, value string) error {
	if ep != nil {
		if ep.Pars == nil {
			ep.Pars = make(map[string]string)
		}
		ep.Pars[param] = value
	}

	return errors.New("uninitialized ep")
}

func GetNthParamsInt(param string, n int) (int, bool) {
	splitted := strings.Split(param, ":")
	if n < len(splitted) {
		result_str := splitted[n]
		if result, ok := strconv.Atoi(result_str); ok == nil {
			return result, true
		}
	}
	return 0, false
}

func GetNthParamsString(param string, n int) (string, bool) {
	splitted := strings.Split(param, ":")
	if n < len(splitted) {
		result_str := splitted[n]
		return result_str, true
	}
	return "", false
}

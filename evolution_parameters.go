package mel

import (
	"strconv"
	"strings"
)

type EvolutionParameters struct {
	Pars map[string]string
}

func (ep *EvolutionParameters) Get_matching_list(match string) (map[string]string, bool) {

	result := make(map[string]string)

	for param, value := range ep.Pars {
		if strings.HasPrefix(param, match) {
			result[param[len(match):len(param)]] = value
		}
	}

	if len(result) > 0 {
		return result, true
	}

	return result, false
}

func (ep *EvolutionParameters) Get_value(param string) (string, bool) {
	if result, ok := ep.Pars[param]; ok {
		return result, true
	}
	return "", false
}

func Get_nth_params_int(param string, n int) (int, bool) {
	splitted := strings.Split(param, ":")
	if n < len(splitted) {
		result_str := splitted[n]
		if result, ok := strconv.Atoi(result_str); ok == nil {
			return result, true
		}
	}
	return 0, false
}

func Get_nth_params_string(param string, n int) (string, bool) {
	splitted := strings.Split(param, ":")
	if n < len(splitted) {
		result_str := splitted[n]
		return result_str, true
	}
	return "", false
}

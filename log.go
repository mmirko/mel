package mel

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type log_entry struct {
	message      string
	message_prio int
}

func logger(logger_id int, ep *EvolutionParameters, logchan <-chan log_entry) {

	log_targets, _ := ep.Get_matching_list("log_target:")

	target := "stdout"
	w := bufio.NewWriter(os.Stdout)

	for log_values, _ := range log_targets {
		if log_target_id_str, ok := Get_nth_params_string(log_values, 0); ok {
			if log_target_verbosity, ok := Get_nth_params_string(log_values, 1); ok {
				log_target_id, _ := strconv.Atoi(log_target_id_str)
				if log_target_id == logger_id {

					if targett, ok := ep.Get_value("log_target:" + log_target_id_str + ":" + log_target_verbosity); ok {

						if targett != "stdout" {

							target = targett

							fi, err := os.Create(target)
							if err != nil {
								panic(err)
							}

							defer func() {
								if err := fi.Close(); err != nil {
									panic(err)
								}
							}()

							w = bufio.NewWriter(fi)
						}
					}
				}
			}
		}
	}

	for {
		entry := <-logchan
		fmt.Fprintf(w, "%-40s - %s\n", time.Now(), entry.message)
		w.Flush()
	}
}

func logi(num int) string {
	result := strconv.Itoa(num)
	return result
}

func logf(num float32) string {
	result := strconv.FormatFloat(float64(num), 'f', 10, 64)
	return result
}

func logit(data log_entry, logverbs []int, logchans []chan log_entry) {
	for chanid, logchan := range logchans {
		if data.message_prio <= logverbs[chanid] {
			logchan <- data
		}
	}
}

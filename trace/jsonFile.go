package trace

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ms-henglu/pal/rawlog"
	"github.com/ms-henglu/pal/types"
)

func requestTracesFromJsonFile(input string) ([]types.RequestTrace, error) {
	fileData, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %v", err)
	}

	defer fileData.Close()

	scanner := bufio.NewScanner(fileData)

	var jsonLine map[string]interface{}

	traces := make([]types.RequestTrace, 0)

	traceLines, requestCount, responseCount := 0, 0, 0
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &jsonLine); err != nil {
			return nil, fmt.Errorf("could not unmarhal text into json %v", err)
		}

		l, err := rawlog.NewRawLogJson(jsonLine)
		if err != nil {
			log.Printf("[WARN] failed to parse log: %v", err)
		}

		if l == nil {
			continue
		}

		t, err := newRequestTrace(*l)
		if err == nil {
			traces = append(traces, *t)

			if t.Request != nil {
				requestCount++
			}
			if t.Response != nil {
				responseCount++
			}
		}

		traceLines++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input file: %v", err)
	}

	log.Printf("[INFO] total traces: %d", traceLines)
	log.Printf("[INFO] request count: %d", requestCount)
	log.Printf("[INFO] response count: %d", responseCount)

	return mergeTraces(traces), nil
}

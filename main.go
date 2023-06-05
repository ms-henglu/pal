package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/ms-henglu/pal/formatter"
	"github.com/ms-henglu/pal/provider"
	"github.com/ms-henglu/pal/types"
	"github.com/ms-henglu/pal/utils"
)

var rawLogRegex = regexp.MustCompile(`([\d+.:T\-]{28})\s\[([A-Z]+)]`)

var providers = []types.Provider{
	provider.AzureADProvider{},
	provider.AzureRMProvider{},
	provider.AzAPIProvider{},
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <input file>", os.Args[0])
	}
	input := os.Args[1]
	log.Printf("[INFO] input file: %s", input)

	data, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	logRegex := regexp.MustCompile(`[\d+.:T\-]{28}\s\[[A-Z]+]`)
	lines := utils.SplitBefore(string(data), logRegex)
	log.Printf("[INFO] total lines: %d", len(lines))

	traces := make([]types.RequestTrace, 0)
	for _, line := range lines {
		l, err := NewRawLog(line)
		if err != nil {
			log.Printf("[WARN] failed to parse log: %v", err)
		}
		if l == nil {
			continue
		}
		t, err := NewRequestTrace(*l)
		if err == nil {
			traces = append(traces, *t)
		}
	}
	requestCount, responseCount := 0, 0
	for _, t := range traces {
		if t.Request != nil {
			requestCount++
		}
		if t.Response != nil {
			responseCount++
		}
	}
	log.Printf("[INFO] total traces: %d", len(traces))
	log.Printf("[INFO] request count: %d", requestCount)
	log.Printf("[INFO] response count: %d", responseCount)

	mergedTraces := make([]types.RequestTrace, 0)
	for i := 0; i < len(traces); i++ {
		if traces[i].Request != nil {
			found := false
			for j := i + 1; j < len(traces); j++ {
				if traces[j].Response == nil || traces[i].Url != traces[j].Url || traces[i].Host != traces[j].Host {
					continue
				}
				found = true
				mergedTraces = append(mergedTraces, types.RequestTrace{
					TimeStamp:  traces[i].TimeStamp,
					Url:        traces[i].Url,
					Method:     traces[i].Method,
					Host:       traces[i].Host,
					StatusCode: traces[j].StatusCode,
					Request:    traces[i].Request,
					Response:   traces[j].Response,
				})
				break
			}
			if !found {
				log.Printf("[WARN] failed to find response for request: %v", traces[i])
				mergedTraces = append(mergedTraces, traces[i])
			}
		}
	}
	log.Printf("[INFO] merged traces: %d", len(mergedTraces))

	output := tips
	format := formatter.MarkdownFormatter{}
	for _, t := range mergedTraces {
		output += format.Format(t)
	}
	if err := os.WriteFile(path.Join(input, "..", "output.md"), []byte(output), 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
	log.Printf("[INFO] output file: %s", path.Clean(path.Join(input, "..", "output.md")))
}

const tips = `<!--
Tips:

1. Use Markdown preview mode to get a better reading experience.
2. If you want to select some of the request traces, in VSCode, use shortcut "Ctrl + K, 0" to fold all blocks.

-->

`

func NewRawLog(message string) (*types.RawLog, error) {
	matches := rawLogRegex.FindAllStringSubmatch(message, -1)
	if len(matches) == 0 || len(matches[0]) != 3 {
		return nil, fmt.Errorf("failed to parse log message: %s", message)
	}
	t, err := time.Parse("2006-01-02T15:04:05.999-0700", matches[0][1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse timestamp: %v", err)
	}
	m := message[len(matches[0][0]):]
	m = strings.Trim(m, " \n")
	return &types.RawLog{
		TimeStamp: t,
		Level:     matches[0][2],
		Message:   m,
	}, nil
}

func NewRequestTrace(l types.RawLog) (*types.RequestTrace, error) {
	for _, p := range providers {
		if p.IsRequestTrace(l) {
			return p.ParseRequest(l)
		}
		if p.IsResponseTrace(l) {
			return p.ParseResponse(l)
		}
	}
	return nil, fmt.Errorf("TODO: implement other providers")
}

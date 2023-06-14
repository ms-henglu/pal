package rawlog

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type RawLog struct {
	TimeStamp time.Time
	Level     string
	Message   string
}

var rawLogRegex = regexp.MustCompile(`([\d+.:T\-]{28})\s\[([A-Z]+)]`)

func NewRawLog(message string) (*RawLog, error) {
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
	return &RawLog{
		TimeStamp: t,
		Level:     matches[0][2],
		Message:   m,
	}, nil
}

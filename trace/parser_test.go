package trace

import (
	"bytes"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestTracesFromMixedFile(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	inputFile := filepath.Join("..", "testdata", "input_mixed.txt")

	mergedTraces, err := NewRequestTraceParser().ParseFromFile(inputFile)
	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, 15, len(mergedTraces)) // 15 merged traces

	assert.Contains(t, logBuffer.String(), "total traces: 29191", "there must be 29191 traces inside the json log")
	assert.Contains(t, logBuffer.String(), "total plain log lines: 21373", "there must be 21373 traces in plain format inside the json log")
	assert.Contains(t, logBuffer.String(), "total json log lines: 7818", "there must be 7818 traces in json forma inside the json log")
	assert.Contains(t, logBuffer.String(), "request count: 20", "there must be 20 requests inside the json log")
	assert.Contains(t, logBuffer.String(), "response count: 20", "there must be 20 responses inside the json log")
}

package trace

import (
	"bytes"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestTracesFromPlainTextFile(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	inputFile := filepath.Join("..", "testdata", "input.txt")

	mergedTraces, err := NewRequestTraceParser().ParseFromFile(inputFile)
	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, 0, len(mergedTraces)) // we are skipping request that just use the url "/" and have a status code of 0

	assert.Contains(t, logBuffer.String(), "total traces: 21373", "there must be 21373 traces inside the json log")
	assert.Contains(t, logBuffer.String(), "request count: 5", "there must be 5 requests inside the json log")
	assert.Contains(t, logBuffer.String(), "response count: 5", "there must be 5 responses inside the json log")
}

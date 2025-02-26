package trace

import (
	"bytes"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestTracesFromJsonFile(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	inputFile := filepath.Join("..", "testdata", "input_json.txt")

	mergedTraces, err := NewRequestTraceParser().ParseFromFile(inputFile)
	assert.NoError(t, err, "There should no error")
	assert.Equal(t, 15, len(mergedTraces))

	assert.Contains(t, logBuffer.String(), "total traces: 7818", "there must be 7818 traces inside the json log")
	assert.Contains(t, logBuffer.String(), "request count: 15", "there must be 15 requests inside the json log")
	assert.Contains(t, logBuffer.String(), "response count: 15", "there must be 15 responses inside the json log")
}

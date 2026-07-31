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

func TestRequestTracesFromUTCTimestampFile(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	inputFile := filepath.Join("..", "testdata", "input_utc.txt")

	mergedTraces, err := NewRequestTraceParser().ParseFromFile(inputFile)
	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, 4, len(mergedTraces)) // resourceGroup create, routeTable create, route existence check, route create failure

	assert.Equal(t, "PUT", mergedTraces[0].Method)
	assert.Equal(t, 201, mergedTraces[0].StatusCode)
	assert.Equal(t, "PUT", mergedTraces[1].Method)
	assert.Equal(t, 201, mergedTraces[1].StatusCode)
	assert.Equal(t, "GET", mergedTraces[2].Method)
	assert.Equal(t, 404, mergedTraces[2].StatusCode)
	assert.Equal(t, "PUT", mergedTraces[3].Method)
	assert.Equal(t, 400, mergedTraces[3].StatusCode)

	assert.Contains(t, logBuffer.String(), "request count: 4")
	assert.Contains(t, logBuffer.String(), "response count: 4")
}

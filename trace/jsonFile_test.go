package trace

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestTracesFromJsonFile(t *testing.T) {
	inputFile := filepath.Join("..", "testdata", "input_json.txt")

	mergedTraces, err := NewRequestTraceParser().ParseFromFile(inputFile)
	assert.NoError(t, err, "There should no error")
	assert.Equal(t, 30, len(mergedTraces))
}

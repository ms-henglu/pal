package trace

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestTracesFromPlainTextFile(t *testing.T) {
	inputFile := filepath.Join("..", "testdata", "input.txt")

	mergedTraces, err := NewRequestTraceParser().ParseFromFile(inputFile)
	assert.NoError(t, err, "There should no error")
	assert.Equal(t, 10, len(mergedTraces))
}

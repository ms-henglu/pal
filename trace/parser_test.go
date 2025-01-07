package trace

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestTracesFromMixedFile(t *testing.T) {
	inputFile := filepath.Join("..", "testdata", "input_mixed.txt")

	mergedTraces, err := NewRequestTraceParser().ParseFromFile(inputFile)
	assert.NoError(t, err, "There should no error")
	assert.Equal(t, 40, len(mergedTraces)) //30 from json 10 from plain
}

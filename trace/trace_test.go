package trace

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyRequresTraceJsonInput(t *testing.T) {
	inputFile := filepath.Join("..", "testdata", "input_json.txt")

	mergedTraces, err := requestTracesFromJsonFile(inputFile)

	assert.NoError(t, err, "There should no error")
	assert.Equal(t, 15, len(mergedTraces))

	for i := range mergedTraces {
		out := VerifyRequestTrace(mergedTraces[i])
		assert.Equal(t, 0, len(out), "verification should not fail")
	}
}

func TestVerifyRequresTracePlainInput(t *testing.T) {
	inputFile := filepath.Join("..", "testdata", "input.txt")

	mergedTraces, err := requestTracesFromFile(inputFile)

	assert.NoError(t, err, "There should no error")
	assert.Equal(t, 5, len(mergedTraces))

	for i := range mergedTraces {
		out := VerifyRequestTrace(mergedTraces[i])
		assert.Equal(t, 0, len(out), "verification should not fail")
	}
}

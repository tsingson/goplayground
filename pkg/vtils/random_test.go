package vtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandInt(t *testing.T) {
	i := GenerateRandInt(1, 2)
	assert.Equal(t, i, 1)
}

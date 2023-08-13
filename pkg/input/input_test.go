package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveWithPattern(t *testing.T) {
	strWithoutSpace := RemoveWithPattern("9 87458-854214 741", "[^0-9]+")
	assert.Equal(t, "987458854214741", strWithoutSpace)
}

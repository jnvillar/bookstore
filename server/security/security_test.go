package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndSalt(t *testing.T) {
	hash1, err1 := HashAndSalt("admin")
	assert.Nil(t, err1)
	result := ComparePassWords("admin", hash1)
	assert.True(t, result)
}

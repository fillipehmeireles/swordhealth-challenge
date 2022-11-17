package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashPass(t *testing.T) {
	password := "sword_health"

	hashedPassword, err := HashPass(password)

	assert.NoError(t, err)
	assert.Equal(t, true, CheckPass(password, hashedPassword))
}

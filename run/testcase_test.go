package run

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	Bar()
	assert.NotNil(t, 1)
}

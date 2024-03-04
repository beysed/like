package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	r, e := Parse("example.like", []byte("hell"))
	t.Log(r)
	assert.Equal(t, true, e == nil)
}

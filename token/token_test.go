package token

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNew(t *testing.T) {
	tok := New("TT_INT", "100")
	assert.Equal(t, "TT_INT", tok.name)
	assert.Equal(t, "100", tok.value)
}

func TestNewTokenList(t *testing.T) {
	list := NewTokenList()
	assert.NotEqual(t, nil, &list)
}

func TestToString(t *testing.T) {
	tok := New("TT_INT", "100")
	str := tok.ToString()
	assert.Equal(t, "TT_INT: 100", str)

	tok = New("TT_PLUS", "")
	str = tok.ToString()
	assert.Equal(t, "TT_PLUS", str)
}

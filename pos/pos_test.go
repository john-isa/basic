package pos

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNew(t *testing.T) {
	new_pos := New("1+1")

	assert.NotEqual(t, nil, &new_pos)    // Position object exists.
	assert.Equal(t, "1+1", new_pos.text) // The text that is stored
	assert.Equal(t, -1, new_pos.index)   // Pointing to just before the first character in the text.
	assert.Equal(t, -1, new_pos.column)  // Pointing to just before the first column in the line of text.
	assert.Equal(t, 0, new_pos.line)     // Pointing to the first line of text.
}

func TestAdvance(t *testing.T) {
	new_pos := New("1+1")

	letter := new_pos.Advance()
	assert.Equal(t, "1", letter)

	letter = new_pos.Advance()
	assert.Equal(t, "+", letter)

	letter = new_pos.Advance()
	assert.Equal(t, "1", letter)

	letter = new_pos.Advance()
	assert.Equal(t, "", letter)
}

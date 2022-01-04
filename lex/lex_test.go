package lex

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNew(t *testing.T) {
	lexer := New("1 + 1")

	assert.Equal(t, "1", lexer.current_char)
	assert.NotEqual(t, nil, &lexer.position, "Starts just before the begining of Text")

	assert.NotEqual(t, nil, &lexer.tok, "A token is born!")
	assert.NotEqual(t, nil, &lexer.tokens, "List of tokens")

	assert.Equal(t, "", lexer.details, "Contains INFO/ERROR deatils")
}

func TestAdvance(t *testing.T) {
	lexer := New("1 + 1")

	assert.Equal(t, "1", lexer.current_char)

	lexer.Advance()
	assert.Equal(t, " ", lexer.current_char)

	lexer.Advance()
	assert.Equal(t, "+", lexer.current_char)

	lexer.Advance()
	assert.Equal(t, " ", lexer.current_char)

	lexer.Advance()
	assert.Equal(t, "1", lexer.current_char)

	lexer.Advance()
	assert.Equal(t, "", lexer.current_char)

	lexer.Advance()
	assert.Equal(t, "", lexer.current_char)
}

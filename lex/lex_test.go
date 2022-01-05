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

func TestMakeTokens(t *testing.T) {
	lexer := New("1 + 1 - 4.2")

	tokenList := lexer.MakeTokens()

	assert.Equal(t, 5, len(tokenList), "We have the correct number of tokens")
	assert.Equal(t, "TT_INT: 1", tokenList[0].ToString(), "The first number")
	assert.Equal(t, "TT_PLUS", tokenList[1].ToString(), "The 'PLUS' operator")
	assert.Equal(t, "TT_INT: 1", tokenList[2].ToString(), "The second number")
	assert.Equal(t, "TT_MINUS", tokenList[3].ToString(), "The 'MINUS' operator")
	assert.Equal(t, "TT_FLOAT: 4.2", tokenList[4].ToString(), "The third number")
}

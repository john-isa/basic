package lex

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNew(t *testing.T) {
	lexer := New("1 + 1")

	assert.Equal(t, "1", lexer.current_char)
	assert.NotEqual(t, nil, &lexer.position, "Starts just before the beginning of Text")

	assert.NotEqual(t, nil, &lexer.tok, "A token is born!")
	assert.NotEqual(t, nil, &lexer.tokenList, "List of tokens")

	assert.Equal(t, "", lexer.details, "Contains INFO/ERROR details")
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
func TestMake_number(t *testing.T) {
	lexer := New("1")
	tokenList := lexer.MakeTokens()
	assert.Equal(t, "TT_INT: 1", tokenList[0].ToString(), "Correct INTEGER")

	lexer = New(".1")
	tokenList = lexer.MakeTokens()
	assert.Equal(t, "TT_FLOAT: 0.1", tokenList[0].ToString(), "Correct FLOAT")

	lexer = New("0.1")
	tokenList = lexer.MakeTokens()
	assert.Equal(t, "TT_FLOAT: 0.1", tokenList[0].ToString(), "Correct FLOAT")

	lexer = New("1.")
	tokenList = lexer.MakeTokens()
	assert.Equal(t, "TT_FLOAT: 1.", tokenList[0].ToString(), "Correct FLOAT")

	lexer = New("1.0")
	tokenList = lexer.MakeTokens()
	assert.Equal(t, "TT_FLOAT: 1.0", tokenList[0].ToString(), "Correct FLOAT")

	lexer = New("1..")
	tokenList = lexer.MakeTokens()
	assert.Equal(t, "ILLEGAL_CHAR: ILLEGAL_CHAR", tokenList[0].ToString(), "The first number")
	assert.Equal(t, "Too many decimal points '.' to make a number!", lexer.details)

	lexer = New("1a")
	tokenList = lexer.MakeTokens()
	assert.Equal(t, "ILLEGAL_CHAR: ILLEGAL_CHAR", tokenList[0].ToString(), "The first number")
	assert.Equal(t, "The character 'a' cannot be part of a number!", lexer.details)

	lexer = New("1!")
	tokenList = lexer.MakeTokens()
	assert.Equal(t, "ILLEGAL_CHAR: ILLEGAL_CHAR", tokenList[0].ToString(), "The first number")
	assert.Equal(t, "The character '!' cannot be part of a number!", lexer.details)
}

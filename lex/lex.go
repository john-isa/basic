package lex

import (
	"fmt"

	"github.com/john-isa/basic/pos"
	"github.com/john-isa/basic/token"
)

type Lexer struct {
	details      string
	position     pos.Position
	current_char string
	tokens       []token.Token
}

func New(text string) Lexer {
	fmt.Println("Creating the Lexer")
	l := Lexer{}

	l.details = ""
	l.position = pos.New(text)
	l.current_char = ""
	l.tokens = token.NewTokenList()
	l.Advance()

	return l
}

func (l *Lexer) Advance() {
	l.current_char = l.position.Advance(l.current_char)
}

func (l *Lexer) MakeTokens() []token.Token {
	fmt.Println("creating a token list")

	return l.tokens
}

func (l *Lexer) make_number() token.Token {
	fmt.Println("Making a number")
	dot_count := 0
	dot_count++

	char := l.current_char

	if char != "" {
		tok := token.New("TT_INT", "1")
		return tok
	}

	tok := token.New("TT_INT", "2")
	return tok
}

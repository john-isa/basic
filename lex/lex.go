package lex

import (
	"fmt"

	"github.com/john-isa/basic/pos"
)

type Lexer struct {
	details      string
	position     Position
	current_char string
	tokens       []Token
}

func New(text string) Lexer {
	fmt.Println("Creating the Lexer")
	l := Lexer{}

	l.details = ""
	l.position = pos.New(text)
	l.current_char = ""
	l.tokens = []Token{}
	l.Advance()

	return l
}

func (l *Lexer) Advance() {
	l.current_char = l.position.advance(l.current_char)
}

func (l *Lexer) MakeTokens() []Token {
	fmt.Println("creating a token list")

	return l.tokens
}

func (l *Lexer) make_number() Token {
	fmt.Println("Making a number")
	dot_count := 0
	dot_count++

	char := l.current_char

	if char != "" {
		token := Token{"TT_INT", "1"}
		return token
	}

	return Token{"TT_INT", "2"}
}

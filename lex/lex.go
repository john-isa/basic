package lex

import (
	"fmt"
	"strings"

	"github.com/john-isa/basic/constant"
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

	l.make_number()

	return l.tokens
}

func (l *Lexer) make_number() token.Token {
	fmt.Println("Making a number")
	number := ""
	dot_count := 0

	// check that the current character is a digit or a decimal point (dot)
	for strings.Contains(constant.DIGITS, l.current_char) || l.current_char == constant.DECIMAL_POINT {
		if l.current_char == constant.DECIMAL_POINT {
			if dot_count == 1 {
				return token.Token{}
			}
			dot_count++
			l.Advance()
		} else {
			number = number + l.current_char
			l.Advance()
		}
	}

	if dot_count == 0 {
		tok := token.New("TT_INT", number)
		return tok
	} else {
		tok := token.New("TT_FLOAT", number)
		return tok
	}
}

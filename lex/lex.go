package lex

import (
	"fmt"
	"strings"

	"github.com/john-isa/basic/constant"
	"github.com/john-isa/basic/pos"
	"github.com/john-isa/basic/token"
)

//=============================================================================
// Lexer object
//=============================================================================
type Lexer struct {
	details      string
	position     pos.Position
	current_char string
	tok          token.Token
	tokens       []token.Token
}

//=============================================================================
// New        : Creates a new Lexer bject for use
// PARAMETERS : The text that will be parsed
// RETURNS    : A lexer object
//=============================================================================
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

//=============================================================================
// Advance    : Fetches the next valid character in the text
// PARAMETERS : none
// RETURNS    : none
//=============================================================================
func (l *Lexer) Advance() {
	l.current_char = l.position.Advance(l.current_char)
}

//=============================================================================
// MakeTokens : Scans the text and parses it into a series of tokens in a list
// PARAMETERS : none
// RETURNS    : a Token array (slice)
//=============================================================================
func (l *Lexer) MakeTokens() []token.Token {
	fmt.Println("creating a token list")

	for l.current_char != "" {
		switch l.current_char {
		case " ", "\t":
			l.Advance()

		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			l.make_number()

		case "+":
			l.tok = token.New(constant.TT_PLUS, "")
			l.Advance()
		case "-":
			l.tok = token.New(constant.TT_MINUS, "")
			l.Advance()
		case "/":
			l.tok = token.New(constant.TT_DIV, "")
			l.Advance()
		case "*":
			l.tok = token.New(constant.TT_MUL, "")
			l.Advance()
		case "(":
			l.tok = token.New(constant.TT_LPAREN, "")
			l.Advance()
		case ")":
			l.tok = token.New(constant.TT_RPAREN, "")
			l.Advance()
		}
		l.tokens = append(l.tokens, l.tok)
	}
	return l.tokens
}

//=============================================================================
// make_number: parses any string of digits and create the correct number token
// PARAMETERS : none
// RETURNS    : none
//=============================================================================
func (l *Lexer) make_number() {
	fmt.Println("Making a number")
	number := ""
	dot_count := 0

	// check that the current character is a digit or a decimal point (dot)
	for strings.Contains(constant.DIGITS, l.current_char) || l.current_char == constant.DECIMAL_POINT {
		if l.current_char == constant.DECIMAL_POINT {
			if dot_count == 1 {
				l.tok = token.Token{}
			}
			dot_count++
		} else {
			number = number + l.current_char
		}
		l.Advance()
	}

	if dot_count == 0 {
		l.tok = token.New("TT_INT", number)
	} else {
		l.tok = token.New("TT_FLOAT", number)
	}
}

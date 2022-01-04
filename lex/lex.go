package lex

import (
	"strings"

	"github.com/john-isa/basic/constant"
	"github.com/john-isa/basic/pos"
	"github.com/john-isa/basic/token"
)

// The Lexer object contains the current position of the text being parsed, its poisution in the text, and
// also builds the list of tokens that represent the program.
//
// There is also storage for INFO/Error details.
//
// No inputs required.
type Lexer struct {
	current_char string        // The character being studied.
	position     pos.Position  // The position of the  current character being studied.
	tok          token.Token   // The token istself
	tokens       []token.Token // The list of tokens that represents the program.
	details      string        // The error or info details.
}

// New creates a new Lexer object that stores the text and sets all the parameters to point to the start.
//
// The text that will be parsed is supplied as a parameter and a Lexer object is returned.
func New(text string) Lexer {
	//
	// Create the lexer that will create and store the token list.
	//
	l := Lexer{}

	l.details = ""
	l.position = pos.New(text)
	l.current_char = ""
	l.tokens = token.NewTokenList()
	l.Advance()

	return l
}

// Advance to the next character in the text. Returns emtpy string if the EOT is reached.
func (l *Lexer) Advance() {
	l.current_char = l.position.Advance()
}

//=============================================================================.
// MakeTokens : Scans the text and parses it into a series of tokens in a list.
// PARAMETERS : none.
// RETURNS    : a Token array (slice).
//=============================================================================.
func (l *Lexer) MakeTokens() []token.Token {
	for l.current_char != "" {
		switch l.current_char {
		case " ", "\n", "\t":
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

//=============================================================================.
// make_number: Parses any string of digits to create the correct number token
//              and store it.
// PARAMETERS : none.
// RETURNS    : none.
//=============================================================================.
func (l *Lexer) make_number() {
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

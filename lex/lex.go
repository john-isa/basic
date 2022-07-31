package lex

import (
	"fmt"
	"strings"

	"basic/constant"
	"basic/pos"
	"basic/token"
)

// Lexer contains the current position of the text being parsed, its poisution in the text, and
// also builds the list of tokens that represent the program.
//   There is also storage for INFO/Error details.
//   No inputs required.
type Lexer struct {
	current_char string        // The character being studied.
	position     pos.Position  // The position of the  current character being studied.
	tok          token.Token   // The token istself
	tokens       []token.Token // The list of tokens that represents the program.
	details      string        // The error or info details.
}

// New creates a new Lexer object that stores the text and sets all the parameters to point to the start.
//  The text that will be parsed is supplied as a parameter and a Lexer object is returned.
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

// MakeTokens scans the text and parses it into a series of tokens in a list.
//  The list is returned to the caller as a slice (array) of tokens.
func (l *Lexer) MakeTokens() []token.Token {
	for l.current_char != "" {
		switch l.current_char {
		case constant.WHITESPACE:
			l.Advance()

		case constant.DIGITS:
			l.make_number()
			l.tokens = append(l.tokens, l.tok)
		case "+":
			l.tok = token.New(constant.TT_PLUS, "")
			l.tokens = append(l.tokens, l.tok)
			l.Advance()
		case "-":
			l.tok = token.New(constant.TT_MINUS, "")
			l.tokens = append(l.tokens, l.tok)
			l.Advance()
		case "/":
			l.tok = token.New(constant.TT_DIV, "")
			l.tokens = append(l.tokens, l.tok)
			l.Advance()
		case "*":
			l.tok = token.New(constant.TT_MUL, "")
			l.tokens = append(l.tokens, l.tok)
			l.Advance()
		case "(":
			l.tok = token.New(constant.TT_LPAREN, "")
			l.tokens = append(l.tokens, l.tok)
			l.Advance()
		case ")":
			l.tok = token.New(constant.TT_RPAREN, "")
			l.tokens = append(l.tokens, l.tok)
			l.Advance()
		}
	}
	return l.tokens
}

// make_number is a private function that collects a series of digits (and a decimal point if it exists) and returens an Integer or Floating-point number.
//  It keeps going until it finds a SPACE, EOT, or an illegal character (not a decimal point or a digit)
func (l *Lexer) make_number() {
	number := ""
	dot_count := 0
	firstTimeRunning := true

	// check that the current character is a digit or a decimal point (dot)
	for strings.Contains(constant.DIGITS, l.current_char) {

		if l.current_char == constant.DOT {
			if dot_count == 1 {
				l.tok = token.New(constant.TT_ILLEGAL_CHAR, constant.TT_ILLEGAL_CHAR)
				l.details = fmt.Sprintf("Too many decimal points '%s' to make a number!", l.current_char)
				l.Advance()
				return
			} else {
				dot_count++
			}
		}
		if firstTimeRunning && dot_count == 1 {
			number = number + "0" + l.current_char
		} else {
			number = number + l.current_char
		}
		firstTimeRunning = false
		l.Advance()

		if l.current_char == "" || l.current_char == " " {
			break
		}

		if !strings.Contains(constant.DIGITS, l.current_char) {
			l.tok = token.New(constant.TT_ILLEGAL_CHAR, constant.TT_ILLEGAL_CHAR)
			l.details = fmt.Sprintf("The character '%s' cannot be part of a number!", l.current_char)
			l.Advance()
			return
		}
	}

	if dot_count == 0 {
		l.tok = token.New("TT_INT", number)
	} else {
		l.tok = token.New("TT_FLOAT", number)
	}
}

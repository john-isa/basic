package lex

import (
	"fmt"
	"strings"

	"basic/constant"
	"basic/pos"
	"basic/token"
)

// Lexer contains the current position of the text being parsed, its position in the text, and
// also builds the tokenList of tokens that represent the program.
//
//	There is also storage for INFO/Error details.
//	No inputs required.
type Lexer struct {
	current_char string          // The character being studied.
	position     pos.Position    // The position of the  current character being studied.
	tok          token.Token     // The token itself
	tokenList    token.TokenList // The list of tokens that represents the program.
	details      string          // The error or info details.
}

// New creates a new Lexer object that stores the text and sets all the parameters to point to the start.
//
//	The text that will be parsed is supplied as a parameter and a Lexer object is returned.
func New(text string) Lexer {
	//
	// Create the lexer that will create and store the token list.
	//
	lex := Lexer{
		current_char: "",
		position:     pos.New(text),
		tokenList:    token.NewTokenList(),
		details:      "",
	}
	lex.Advance()

	return lex
}

// Advance to the next character in the text. Returns empty string if the EOT is reached.
func (l *Lexer) Advance() {
	l.current_char = l.position.Advance()
}

// MakeTokens scans the text and parses it into a series of tokens in a list.
//
//	The tokenList is returned to the caller as a slice (array) of tokens.
func (l *Lexer) MakeTokens() token.TokenList {
	for l.current_char != "" {
		switch l.current_char {
		case constant.SPACE, constant.NEWLINE, constant.TAB:
			l.Advance()

		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".":
			l.tokenList = append(l.tokenList, l.make_number())
		case "+":
			l.tokenList = append(l.tokenList, token.New(constant.TT_PLUS, ""))
			l.Advance()
		case "-":
			l.tokenList = append(l.tokenList, token.New(constant.TT_MINUS, ""))
			l.Advance()
		case "/":
			l.tokenList = append(l.tokenList, token.New(constant.TT_DIV, ""))
			l.Advance()
		case "*":
			l.tokenList = append(l.tokenList, token.New(constant.TT_MUL, ""))
			l.Advance()
		case "(":
			l.tokenList = append(l.tokenList, token.New(constant.TT_MUL, ""))
			l.Advance()
		case ")":
			l.tokenList = append(l.tokenList, token.New(constant.TT_RPAREN, ""))
			l.Advance()
		}
	}
	return l.tokenList
}

// make_number is a private function that collects a series of digits (and a decimal point if it exists)
// and returns an Integer or Floating-point number.
// It keeps going until it finds a SPACE, EOT, or an illegal character (not a decimal point or a digit)
func (l *Lexer) make_number() token.Token {
	dotCount := 0         // No decimal point assumed.
	tokenType := "TT_INT" // Assume it is an integer until shown otherwise
	number := ""          // No number as yet.

	// check that the current character is a digit or a decimal point (dot)
	for strings.Contains(constant.DIGITS, l.current_char) {
		switch l.current_char {
		case constant.DOT:
			switch dotCount {
			case 0:
				dotCount++
				if number == "" {
					number = number + "0" + l.current_char
				} else {
					number = number + l.current_char
				}
				tokenType = "TT_FLOAT"
				l.Advance()
			case 1:
				l.tok = token.New(constant.TT_ILLEGAL_CHAR, constant.TT_ILLEGAL_CHAR)
				l.details = fmt.Sprintf("Too many decimal points '%s' to make a number!", l.current_char)
				return l.tok
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			number = number + l.current_char
			l.Advance()
		default:
			if (len(l.current_char) == 1 && l.current_char != " ") || (len(l.current_char) == 1 && l.current_char != "\n") || (len(l.current_char) == 1 && l.current_char != "\t") {
				tokenType = constant.TT_ILLEGAL_CHAR
			}

			l.tok = token.New(tokenType, l.current_char)
			return l.tok
		}
	}

	if (len(l.current_char) == 1 && l.current_char != " ") || (len(l.current_char) == 1 && l.current_char != "\n") || (len(l.current_char) == 1 && l.current_char != "\t") {
		tokenType = constant.TT_ILLEGAL_CHAR
		number = constant.TT_ILLEGAL_CHAR
	}

	l.tok = token.New(tokenType, number)
	return l.tok
}

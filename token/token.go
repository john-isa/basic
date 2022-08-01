package token

// Token object.
//   Stores the TOKEN name and value information.
type Token struct {
	name  string
	value string
}

type TokenList []Token

// New creates a new Token for use. It contains the:
//   PARAMETERS : name  - The name of the Token, and the,
//              : value - The value of the Token
//   RETURNS    : the Token object.
func New(name string, value string) Token {
	return Token{
		name:  name,
		value: value,
	}
}

// New creates a new list of Tokens (slice).
//   PARAMETERS : none
//   RETURNS    : The Token slice (array - currently empty)
func NewTokenList() TokenList {
	return TokenList{}
}

// Add inserts the given token into the tokenList.
func (l TokenList) Add(tok Token) {
	l = append(l, tok)
}

// ToString creates a string representation of the name and value of the Token
//   PARAMETERS : none
//   RETURNS    : The string representing the Token
func (t *Token) ToString() string {
	if t.value != "" {
		return t.name + ": " + t.value
	} else {
		return t.name
	}
}

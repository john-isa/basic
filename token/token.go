package token

// Token object.
//   Stores the TOKEN name and value information.
type Token struct {
	name  string
	value string
}

// New creates a new Token for use. It contains the:
//   name  - The name of the Token, and the,
//   value - The value of the Token
// It returns the Token object.
func New(n string, v string) Token {
	t := Token{name: n, value: v}
	return t
}

// New creates a new list of Tokens (slice).
//   PARAMETERS : none
//   RETURNS    : The Token slice (array - currently empty)
func NewTokenList() []Token {
	tokens := []Token{}

	return tokens
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

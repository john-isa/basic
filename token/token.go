package token

//=============================================================================
// Token object
//=============================================================================
type Token struct {
	name  string
	value string
}

//=============================================================================
// New        : Creates a new Token for use
// PARAMETERS : name  - The name of the token
//              value - The value of the TYoken
// RETURNS    : The Token object
//=============================================================================
func New(n string, v string) Token {
	t := Token{name: n, value: v}
	return t
}

//=============================================================================
// New        : Creates a new list of Tokens (slice) for use
// PARAMETERS : none
// RETURNS    : The Token slice (array)
//=============================================================================
func NewTokenList() []Token {
	tokens := []Token{}

	return tokens
}

//=============================================================================
// ToString   : Creates a string representation of the name and value of the Token
// PARAMETERS : none
// RETURNS    : The string representing the Token
//=============================================================================
func (t *Token) ToString() string {
	if t.value != "" {
		return t.name + ": " + t.value
	} else {
		return t.name
	}
}

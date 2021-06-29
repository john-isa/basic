package token

type Token struct {
	name  string
	value string
}

func New(n string, v string) Token {
	t := Token{name: n, value: v}
	return t
}

func NewTokenList() []Token {
	tokens := []Token{}

	return tokens
}

func (t *Token) GetToken() string {
	return t.ToString()
}

func (t *Token) ToString() string {
	if t.value != "" {
		return t.name + ": " + t.value
	} else {
		return t.name
	}
}

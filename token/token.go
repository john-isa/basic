package token

type Token struct {
	name  string
	value string
}

func NewToken(n string, v string) Token {
	t := Token{name: n, value: v}
	return t
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

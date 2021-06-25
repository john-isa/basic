package pos

type Position struct {
	text   string
	index  int
	line   int
	column int
}

func New(text_ string) Position {
	p := Position{}

	p.text = text_
	p.index = -1
	p.line = 0
	p.column = -1

	return p
}

func (p *Position) Advance(current_char string) string {
	char := current_char

	p.index++
	p.column++

	if char != "" {
		if char == "\n" {
			p.line++
			p.column = 0
		}
	}

	if p.index >= len(p.text) {
		return "" // returning an empty char
	}

	return p.text[p.index : p.index+1]
}

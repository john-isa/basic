package pos

// Position object stores the text and the current posotion of the current
// character being studied.
type Position struct {
	text   string
	index  int
	line   int
	column int
}

// New creates a new Position object to keep track of the current location in the text.
//
// NOTE       : Initially points to just before the start of text.
func New(text_ string) Position {
	p := Position{text: text_, index: -1, line: 0, column: -1}

	return p
}

// Advance fetches the next character in the current text and returns it to the caller.
//
// If the current_char is valid and it is not at the end of the text then the next
// character is returned, otherwise an emtpy character is returned.
func (p *Position) Advance() string {
	//
	// Point to the next position in the text.
	//
	p.index = p.index + 1
	p.column = p.column + 1

	// Fetch the next character while we are in range of text.
	if p.index < len(p.text) {
		return p.text[p.index : p.index+1] // returning the next character.
	} else {
		return "" // We are now at EOT
	}
}

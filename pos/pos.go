package pos

//=============================================================================
// Position object
//=============================================================================
type Position struct {
	text   string
	index  int
	line   int
	column int
}

//=============================================================================
// New        : Creates a new Position for use
// PARAMETERS : The text that the position is maintained in
// RETURNS    : The Position object
//=============================================================================
func New(text_ string) Position {
	p := Position{}

	p.text = text_
	p.index = -1
	p.line = 0
	p.column = -1

	return p
}

//=============================================================================
// Advance    : gets the next character in the current text and returns it to
//              the caller.
//              If the current_char is valid and it is not at the end of the
//              text then the next character is returned, otherwise an emtpy
//              character is returned.
//
// PARAMETERS : current char
// RETURNS    : next characeter in the stream, or an empty character.
//=============================================================================
func (p *Position) Advance(current_char string) string {
	char := current_char

	p.index++
	p.column++

	if char != "" {
		if char == "\n" {
			p.line++
			p.column = 0
		}
	} else {
		return "" // returning an empty char
	}

	if p.index >= len(p.text) {
		return "" // returning an empty char
	}

	return p.text[p.index : p.index+1]
}

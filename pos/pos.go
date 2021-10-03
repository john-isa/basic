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
	p := Position{text: text_, index: -1, line: 0, column: -1}

	return p
}

//=============================================================================
// Advance    : Fetches the next character in the current text and returns it to
//              the caller.
//
//              If the current_char is valid and it is not at the end of the
//              text then the next character is returned, otherwise an emtpy
//              character is returned.
//
// PARAMETERS : current char
// RETURNS    : next characeter in the stream, or an empty character.
//=============================================================================
func (p *Position) Advance(current_char string) string {
	p.index = p.index + 1
	p.column = p.column + 1

	if current_char != "" {
		if current_char == "\n" {
			p.line = p.line + 1
			p.column = 0
		}
	} else {
		return "" // not a valid character, returning empty character
	}

	if p.index >= len(p.text) {
		return "" // reached the End-Of-Text, returning an empty character
	}

	return p.text[p.index : p.index+1] // returning the next current valid character
}

package pos

// Program - the text that is parsed and executed/compiled
type Program interface {
	Advance()
	GetPosition()
}

// Position object stores the text and the current position of the current character being studied.
type Position struct {
	text   string // The text of the code to be parsed.
	length int64  // The length of the program text.
	index  int64  // The position of the current character in the text.
	line   int64  // The current line number of the text.
	column int64  // The current column of the text.
}

// New creates a new Position object to keep track of the current location in the text.
//
// NOTE : Initially points to just before the start of text.
func New(text string) Position {
	textLength := int64(len(text))
	return Position{
		text:   text,
		length: textLength,
		index:  -1, // Point to just before the first character in the text.
		line:   0,  // Point to the very first line of the text.
		column: -1, // Point to just before the first column of the text.
	}
}

// GetPosition - returns the current cursor position (line, column).
func (p *Position) GetPosition() (int64, int64) {
	return p.line, p.column
}

// Advance fetches the next character in the current text and returns it to the caller.
//
// If the current_char is valid and it is not at the end of the text then the next
// character is returned, otherwise an empty character is returned.
func (p *Position) Advance() string {
	//
	// Point to the next position in the text.
	//
	p.index++
	p.column++

	// Fetch the next character while we are in range of text.
	if p.index < p.length {
		character := p.text[p.index : p.index+1]

		switch character {
		case "\n": // End-Of-Line character
			p.line++     // Point to the next line of code.
			p.column = 0 // Reset the column pointer.
			p.index++    // Point to the next character.
			return character
		default:
			return character
		}
	} else {
		return "" // We are now at EOT (End-Of-Text).
	}
}

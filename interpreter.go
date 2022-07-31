package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/john-isa/basic/lex"
)

func readLine() ([]byte, error) {
	fmt.Print("BASIC -%> ") // Show the prompt.

	reader := bufio.NewReader(os.Stdin)  // Create a reader to read from the command line.
	text, err := reader.ReadString('\n') // Read a line of text.
	if err != nil {
		return nil, err
	}

	text = strings.TrimSpace(text) // Trim all undesired spaces.
	bytes := []byte(text)

	return bytes, nil
}

func readProgram(name string) ([]byte, error) {
	program, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func Run(program string) {
	var (
		text []byte
		err  error
	)

	if len(program) == 0 {
		text, _ = readLine() // Read a line of text.
	} else {
		text, err = readProgram(program)
		if err != nil {
			return
		}
	}

	lexer := lex.New(text)       // Create a lexer that contains the string to be parsed
	tokens := lexer.MakeTokens() // Parse the string into an array of tokens

	fmt.Println(tokens)
}

//==========================================================================================
// The main entry point
//==========================================================================================
func main() {
	Run("") //
}

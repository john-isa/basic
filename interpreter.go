package main

import (
	"fmt"
	"os"

	"github.com/john-isa/basic/lex"
)

// func readLine() ([]byte, error) {
// 	fmt.Print("BASIC -%> ") // Show the prompt.

// 	reader := bufio.NewReader(os.Stdin)  // Create a reader to read from the command line.
// 	text, err := reader.ReadString('\n') // Read a line of text.
// 	if err != nil {
// 		return nil, err
// 	}

// 	text = strings.TrimSpace(text) // Trim all undesired spaces.
// 	bytes := []byte(text)

// 	return bytes, nil
// }

func readProgram(name string) ([]byte, error) {
	program, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func Run() {
	var (
		text []byte
		err  error
	)

	// args := os.Args

	// if len(args) == 1 {
	// 	text, _ = readLine() // Read a line of text.
	// } else {
	// 	text, err = readProgram(args[1])
	// 	if err != nil {
	// 		return
	// 	}
	// }

	text, err = readProgram("math.bas")
	if err != nil {
		return
	}

	lexer := lex.New(string(text[:])) // Create a lexer that contains the string to be parsed
	tokens := lexer.MakeTokens()      // Parse the string into an array of tokens

	fmt.Println(tokens)
}

//==========================================================================================
// The main entry point
//==========================================================================================
func main() {
	Run()
}

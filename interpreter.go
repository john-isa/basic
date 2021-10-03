package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/john-isa/basic/lex"
)

func readLine(r *bufio.Reader) (string, error) {

	fmt.Print("BASIC -%> ")
	return r.ReadString('\n')
}

func Run() {
	reader := bufio.NewReader(os.Stdin) // create a reader to read from the command line
	text, _ := readLine(reader)         // read a line of text
	text = strings.TrimSpace(text)      // Trims all undesired spaces
	lexer := lex.New(text)              // create a lexer that contains the string to be parsed
	tokens := lexer.MakeTokens()        // parse the string into an array of tokens

	fmt.Println(tokens)
}

//==========================================================================================
// The main entry point
//==========================================================================================
func main() {
	Run()
}

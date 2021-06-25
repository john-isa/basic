package main

//==========================================================================================
// The main entry point
//==========================================================================================
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/john-isa/basic/lex"
)

func Run(text string) {
	lexer := lex.New(text)       // create a lexer that contains the string to be parsed
	tokens := lexer.MakeTokens() // parse the string into an array of tokens

	fmt.Println(tokens)
}

func readLine(r *bufio.Reader) (string, error) {

	fmt.Print("BASIC-> ")
	return r.ReadString('\n')
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := readLine(reader)

	text = strings.TrimSpace(text)

	Run(text)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Number struct{ value int }
type Symbol struct{ name string }

// List(expression)
type List struct{ elements []interface{} }
type Function struct{ name string }
type Lambda struct {
	// string -> expression
	environment map[string]interface{}
	parameters  []string
	// expression
	body interface{}
}

func main() {
	tokens := popString(tokenise(readStdin()))
	fmt.Printf("%v\n", strings.Join(tokens, ", "))
}

func popString(slice []string) []string {
	return append(slice[:0], slice[1:]...)
}

func read(tokens []string) []interface{} {
	return []interface{}{Number{1}}
}

func tokenise(program string) []string {
	padded := " " + program + " "
	spaced := regexp.MustCompile("[()]").ReplaceAllString(padded, " $0 ")
	withPadding := regexp.MustCompile("[ \n]+").Split(spaced, -1)
	return withPadding[1 : len(withPadding)-1]
}

func readStdin() string {
	result := ""
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		result = result + scanner.Text() + "\n"
	}

	if scanner.Err() != nil {
		fmt.Println("Failed to read from stdin.")
	}

	return result
}

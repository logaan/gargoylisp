package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type number struct{ value int }
type symbol struct{ name string }

// List(expression)
type list struct{ elements []interface{} }
type function struct{ name string }
type lambda struct {
	// string -> expression
	environment map[string]interface{}
	parameters  []string
	// expression
	body interface{}
}

func main() {
	tokens := tokenise(readStdin())
	expression := read(tokens)
	fmt.Printf("%+v\n", expression)
}

func popString(slice []string) (string, []string) {
	head := slice[0]
	tail := append(slice[:0], slice[1:]...)
	return head, tail
}

// Should be handling unexpected ")" but exceptions
func read(tokens []string) interface{} {
	token, tokens := popString(tokens)
	if token == "(" {
		list := list{[]interface{}{}}
		for tokens[0] != ")" {
			list.elements = append(list.elements, read(tokens))
		}
		_, tokens = popString(tokens)
		return list
	} else {
		return atom(token)
	}
}

func atom(token string) interface{} {
	i, err := strconv.Atoi(token)
	if err != nil {
		return symbol{token}
	}

	return number{i}
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

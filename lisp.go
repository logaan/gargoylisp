package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Printf("%v\n", tokenise(readStdin()))
}

func tokenise(program string) []string {
	// replace ( and ) with " ( " and " ) "
	spaced := regexp.MustCompile("[()]").ReplaceAllString(program, " $0 ")
	// split on \w*, return all substrings
	return regexp.MustCompile("[ \n]+").Split(spaced, -1)
}

func readStdin() string {
	result := ""
	fmt.Println("Echoing")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result = result + scanner.Text() + "\n"
	}

	if scanner.Err() != nil {
		fmt.Println("Some error occurred.")
	}

	return result
}

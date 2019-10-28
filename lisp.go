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
	spaced := regexp.MustCompile("[()]").ReplaceAllString(program, " $0 ")
	return regexp.MustCompile("[ \n]+").Split(spaced, -1)
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

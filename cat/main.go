package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func Println(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		content, _ := io.ReadAll(os.Stdin)
		Print(string(content))
	}

	for _, filename := range args {

		content, err := os.ReadFile(filename)
		if err != nil {
			Print("ERROR: ")
			Println(err.Error())
			os.Exit(1)
		}
		Print(string(content))
	}
}

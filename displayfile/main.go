package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print("File name missing\n")
	} else if len(os.Args) == 2 {
		fmt.Print("Almost there!!\n")
	} else {
		fmt.Print("Too many arguments\n")
	}
}

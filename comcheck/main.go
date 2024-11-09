package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "01" {
			fmt.Println("Alert!!!")
			return
		}
		if arg == "galaxy" {
			fmt.Println("Alert!!!")
			return
		}
		if arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}

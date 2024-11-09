package main

import "github.com/01-edu/z01"

func PrintStrln(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	PrintStrln("Door Opening...")
	PrintStrln("is the Door closed ?")
	PrintStrln("is the Door opened ?")
	PrintStrln("Door Closing...")
}

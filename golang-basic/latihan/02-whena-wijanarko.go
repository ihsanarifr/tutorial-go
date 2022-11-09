package main

import "fmt"

func main() {
	
	isCharacter("a")
	isCharacter("b")
	isCharacter("j")
	isCharacter("k")
	isCharacter("e")
	isCharacter("f")
	isCharacter("u")
	isCharacter("g")
	isCharacter("u")
	isCharacter("l")
}

func isCharacter(character string) {
	switch character {
	case "a", "i", "u", "e", "o":
		fmt.Printf("%s ini adalah vokal\n", character)
	default:
		fmt.Printf("%s ini adalah konsonan\n", character)
	}
}
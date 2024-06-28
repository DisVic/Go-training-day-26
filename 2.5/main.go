package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	lastRune, size := utf8.DecodeLastRuneInString(text)
	if lastRune == utf8.RuneError && size == 0 {
		return
	}
	if unicode.IsUpper(rune(text[0])) && lastRune == '.' {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}

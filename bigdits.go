package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{" 000 ", "0 0", "0 0", "0 0", "0 0", "0 0", " 000 "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2 2", " 2 ", " 2 ", " 2 ", "2", "2222"},
	{" 222 ", "2 2", " 2 ", " 2 ", " 2 ", "2", "2222"}, {" 222 ", "2 2", " 2 ", " 2 ", " 2 ", "2", "2222"}, {" 222 ", "2 2", " 2 ", " 2 ", " 2 ", "2", "2222"}, {" 222 ", "2 2", " 2 ", " 2 ", " 2 ", "2", "2222"}, {" 222 ", "2 2", " 2 ", " 2 ", " 2 ", "2", "2222"}, {" 222 ", "2 2", " 2 ", " 2 ", " 2 ", "2", "2222"},
	{" 9999", "9 9", "9 9", " 9999", " 9", "    9"},
}

func main() {

	if len(os.Args) == 1 {

		fmt.Printf("Usage:%s<whole-nuymber>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""

		for column := range stringOfDigits {
			digits := stringOfDigits[column] - '0'

			if 0 <= digits && digits <= 9 {
				line += bigDigits[digits][row] + " "

			} else {
				log.Fatal("Invalid  Whole number")
			}

		}
		fmt.Println(line)
	}
}

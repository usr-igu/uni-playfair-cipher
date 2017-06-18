package main

import (
	"fmt"
)

// Playfair
func Playfair(msg, keyword string) string {
	fmt.Printf("Message to cipher: %s, Keyword: %s\n", msg, keyword)

	table := createTable(keyword)

	printTable(table)

	return ""
}

// createTable
func createTable(keyword string) [5][5]rune {

	usedLetters := make(map[rune]bool)

	table := [5][5]rune{}

	row, col := 0, 0
	for _, v := range keyword {
		if !usedLetters[v] {

			usedLetters[v] = true
			table[row][col] = v

			// Anda de acordo na matriz :)
			if col == 4 && row == 4 {
				continue
			}
			col++
			if col >= 5 {
				col = 0
				row++
			}
			if row >= 5 {
				row = 0
			}
		}

	}

	for i := 'A'; i <= 'Z'; i++ {

		c := rune(i)

		if !usedLetters[c] {


			table[row][col] = c

			// Anda de acordo na matriz :)
			if col == 4 && row == 4 {
				continue
			}
			col++
			if col >= 5 {
				col = 0
				row++
			}
			if row >= 5 {
				row = 0
			}
		}
	}

	return table
}

// printTable
func printTable(table [5][5]rune) {
	fmt.Println("Table: ")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%c ", table[i][j])
		}
		fmt.Println()
	}
}

func main() {

	keyword := "PLAYFAIR"
	msg := "HELLOWORLD"

	fmt.Printf("Encoded msg: %s\n", Playfair(msg, keyword))
}

package main

import (
	"bytes"
	"fmt"
	"math"
)

// Playfair recebe uma mensagem e uma chave e criptografa a mensagem usando o técnica de Playfair.
func Playfair(msg, keyword string) string {
	msg = prepareMsg(msg)

	encodedMsg := ""

	fmt.Printf("Message to cipher: %s (%d), Keyword: %s\n", msg, len(msg), keyword)

	table := createTable(keyword)

	for i := 0; i < len(msg); i += 2 {

		c, cn := rune(msg[i]), rune(msg[i+1])

		row, col := whereInTheTable(c, table)
		rown, coln := whereInTheTable(cn, table)

		if row == rown {

			if col < 4 {
				encodedMsg += string(table[row][col+1])
			} else {
				encodedMsg += string(table[row][4-col+1])
			}

			if coln < 4 {
				encodedMsg += string(table[rown][coln+1])
			} else {
				encodedMsg += string(table[rown][4-coln+1])
			}

		} else if col == coln {

			if row < 4 {
				encodedMsg += string(table[row+1][col])
			} else {
				encodedMsg += string(table[rown+1][coln])
			}

			if rown < 4 {
				encodedMsg += string(table[rown+1][col])
			} else {
				encodedMsg += string(table[4-rown+1][coln])
			}

		} else {
			dist := col - coln
			encodedMsg += string(table[row][abs(col-dist)])
			encodedMsg += string(table[rown][abs(coln+dist)])
		}
	}
	printTable(table)

	return encodedMsg
}

// prepareMsg prepara msg para ser utilizada em Playfair.
func prepareMsg(msg string) string {

	msgBs := []byte(msg)

	msgBs = bytes.Replace(msgBs, []byte("J"), []byte("I"), -1) // I == J
	msgBs = bytes.ToUpper(msgBs)

	msgBr := bytes.NewReader(msgBs)
	bs := make([]byte, 2) // Vamos ler do buffer dois bytes de cada vez.
	preparedMessage := make([]byte, 0, 32)
	var a, b, c byte

	for {
		n, err := msgBr.Read(bs)
		if err != nil {
			break
		}

		for _, v := range bs {
			if v == ' ' {
				continue
			}
		}

		switch n {
		case 2:
			a, b = bs[0], bs[1]
			if a == b {
				preparedMessage = append(preparedMessage, a, 'X', b)
			} else {
				preparedMessage = append(preparedMessage, a, b)
			}
		case 1:
			c = bs[0]
			preparedMessage = append(preparedMessage, c)
		}
	}

	if len(preparedMessage)%2 != 0 { // Caso o tamanho da mensagem não seja par adiciona um X no final.
		preparedMessage = append(preparedMessage, 'X')
	}

	return fmt.Sprintf("%s", preparedMessage)
}

// createTable cria e popula uma table.
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

		// Vamos supor que 'I' == 'J' xD
		// Então retiramos 'J' da tabela.
		if c == 'J' {
			continue
		}

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

// printTable imprime table na saída padrão.
func printTable(table [5][5]rune) {
	fmt.Println("Table: ")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%c ", table[i][j])
		}
		fmt.Println()
	}
}

// whereInTheTable procura por uma rune em table.
func whereInTheTable(c rune, table [5][5]rune) (x, y int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if table[i][j] == c {
				x = i
				y = j
			}
		}
	}
	return
}

// abs retorna o valor absoluto de x.
func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	keyword := "PLAYFAIREXAMPLE"
	msg := "HIDETHEGOLDINTHETREESTUMP"
	fmt.Printf("Encoded msg: %s\n", Playfair(msg, keyword))
}

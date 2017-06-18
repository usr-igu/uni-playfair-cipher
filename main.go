package main

import (
	"fmt"
	"math"
	"strings"
)

// Playfair ...
func Playfair(msg, keyword string) string {

	msg = processMsg(msg)
	encodedMsg := ""

	fmt.Printf("Message to cipher: %s (%d), Keyword: %s\n", msg, len(msg), keyword)

	table := createTable(keyword)

	// TODO: Simplificar essa função : (
	for i := 0; i < len(msg); i += 2 {

		c, cn := rune(msg[i]), rune(msg[i+1])

		row, col := whereInTheTable(c, table)
		rown, coln := whereInTheTable(cn, table)

		fmt.Printf("%d %d\n", row, col)

		// FIXME: Outro bug! Preciso checar melhor os ifs aqui xD
		if row == rown {
			if row <= 4 {
				encodedMsg += string(table[row][col+1])
				encodedMsg += string(table[rown][coln+1])
			} else {
				encodedMsg += string(table[row][5-col+1])
				encodedMsg += string(table[rown][5-coln+1])
			}
		} else if col == coln {
			if col <= 4 {
				encodedMsg += string(table[row+1][col])
				encodedMsg += string(table[rown+1][coln])
			} else {
				encodedMsg += string(table[5-row+1][col])
				encodedMsg += string(table[5-rown+1][coln])
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

// processMsg prepara msg para ser utilizada em Playfair.
//
// Converte todos as runes de msg para caixa-alta, remove todos os espaços e
// por fim adiciona um X entre caracteres repeditos.
// TODO: Essa função não está funcionando para todos os casos e pode ser melhorada.
func processMsg(msg string) string {
	// FIXME: Acho que isso aqui é meio ineficiente rsrs.
	msg = strings.Join(strings.Split(strings.ToUpper(msg), " "), "")

	processedMsg := ""

	if len(msg)%2 != 0 {
		msg = msg + "X"
	}

	// todo: Um caso de erro é se houverem mais de dois caracteres repeditos.
	for i := 0; i < len(msg); i += 2 {
		a := msg[i]
		b := msg[i+1]
		if a != b {
			processedMsg += string(a) + string(b)
		} else {
			processedMsg += string(a) + "X" + string(b)
		}
	}

	// TODO: Isso aqui é coisa de macaco : (
	if len(processedMsg)%2 != 0 {
		processedMsg = processedMsg[:len(processedMsg)-1]
	}

	return processedMsg
}

// createTable cria e popula uma table.
//
// Primeiramente cria uma table([5][5]rune) e a popula com as runes de keyword (sem repetições),
// depois preenche o que restar de espaço em table com runes não repetidas restantes (ascii, ordem alfabética)
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
//
// Caso encontre retorna a posição dela em table caso contrário retorna 0,0.
// todo: Fazer essa função retornar um erro.
func whereInTheTable(c rune, table [5][5]rune) (int, int) {
	x, y := 0, 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if table[i][j] == c {
				x = i
				y = j
			}
		}
	}
	return x, y
}

// abs retorna o valor absoluto de x.
func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	keyword := "PLAYFAIREXAMPLE"
	msg := "Hide the gold in the tree stump"
	fmt.Printf("Encoded msg: %s\n", Playfair(msg, keyword))
}

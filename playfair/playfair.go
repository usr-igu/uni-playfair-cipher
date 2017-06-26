package playfair

import (
	"bytes"
	"fmt"
)

type table [5][5]byte

// Playfair recebe uma mensagem e uma chave e criptografa a mensagem usando o técnica de playfair.
func Playfair(msg, keyword string) string {
	keywordBytes := []byte(keyword)
	preparedMsg := prepareMsg(msg)
	encodedMsg := make([]byte, 0, 32)
	fmt.Printf("Mensagem a ser cifrada: %s (%d), Chave: %s\n", preparedMsg, len(preparedMsg), keywordBytes)
	table := createTable(keywordBytes)
	for i := 0; i < len(preparedMsg); i += 2 {
		c, cn := preparedMsg[i], preparedMsg[i+1]
		row, col := whereInTheTable(c, table)
		rown, coln := whereInTheTable(cn, table)
		if row == rown {
			where := 4 - col + 1
			if col < 4 {
				where = col + 1
			}
			encodedMsg = append(encodedMsg, table[row][where])
			where = 4 - coln + 1
			if coln < 4 {
				where = coln + 1
			}
			encodedMsg = append(encodedMsg, table[rown][where])
		} else if col == coln {
			where := 4 - row + 1
			if row < 4 {
				where = row + 1
			}
			encodedMsg = append(encodedMsg, table[where][col])
			where = 4 - rown + 1
			if rown < 4 {
				where = rown + 1
			}
			encodedMsg = append(encodedMsg, table[where][coln])
		} else {
			dist := col - coln
			encodedMsg = append(encodedMsg, table[row][abs(col-dist)])
			encodedMsg = append(encodedMsg, table[rown][abs(coln+dist)])
		}
	}
	printTable(table)
	return fmt.Sprintf("%s", encodedMsg)
}

// prepareMsg prepara msg para ser utilizada em playfair.
func prepareMsg(msg string) []byte {
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
	return preparedMessage
}

// createTable cria e popula uma table.
func createTable(keyword []byte) [5][5]byte {
	usedLetters := make(map[byte]bool)
	table := [5][5]byte{}
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
		c := byte(i)
		// Vamos supor que 'Y' == 'Z' xD
		// Então retiramos 'Z' da tabela.
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
func printTable(table [5][5]byte) {
	fmt.Println("Tabela de cifragem: ")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%c ", table[i][j])
		}
		fmt.Println()
	}
}

// whereInTheTable procura por uma rune em table.
func whereInTheTable(c byte, table [5][5]byte) (x, y int) {
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
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

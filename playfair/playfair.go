package playfair

import (
	"fmt"
	"strings"
)

type table [5][5]byte

// Encrypt criptografa msg usando a cifra de playfair.
func Encrypt(msg, key string) string {

	table := makeTable(key)

	// todo: Tratar letras repetidas.
	msg = strings.Replace(msg, " ", "", -1)  // Remove todos os espaços.
	msg = strings.Replace(msg, "Z", "S", -1) // Trocando todos os Ws por Ms.
	msg = strings.ToUpper(msg)               // Tudo em caixa alta.
	if len(msg)%2 != 0 {                     // Se o número de caracteres da msg não for par fazemos ser!
		msg += "X"
	}

	encryptedMessage := make([]byte, 0, 32)

	for i := 0; i < len(msg); i += 2 {

		row1, col1 := table.where(msg[i])
		row2, col2 := table.where(msg[i+1])

		if row1 == row2 {
			if col1 < 4 {
				encryptedMessage = append(encryptedMessage, table[row1][col1+1])
			} else {
				encryptedMessage = append(encryptedMessage, table[row1][0])
			}
			if col2 < 4 {
				encryptedMessage = append(encryptedMessage, table[row2][col2+1])
			} else {
				encryptedMessage = append(encryptedMessage, table[row2][0])
			}
		} else if col1 == col2 {
			if row1 < 4 {
				encryptedMessage = append(encryptedMessage, table[row1+1][col1])
			} else {
				encryptedMessage = append(encryptedMessage, table[0][col1])
			}
			if row2 < 4 {
				encryptedMessage = append(encryptedMessage, table[row2+1][col1])
			} else {
				encryptedMessage = append(encryptedMessage, table[0][col2])
			}
		} else {
			dist := col1 - col2
			encryptedMessage = append(encryptedMessage, table[row1][abs(col1-dist)])
			encryptedMessage = append(encryptedMessage, table[row2][abs(col2+dist)])
		}
	}
	return fmt.Sprintf("%s", encryptedMessage)
}

// Decrypt descriptografa uma msg criptografada pela cifra
// de playfair usando a key.
func Decrypt(msg, key string) string {

	table := makeTable(key)
	encryptedMessage := make([]byte, 0, 32)

	for i := 0; i < len(msg); i += 2 {

		row1, col1 := table.where(msg[i])
		row2, col2 := table.where(msg[i+1])
		if row1 == row2 {
			if col1 > 0 {
				encryptedMessage = append(encryptedMessage, table[row1][col1-1])
			} else {
				encryptedMessage = append(encryptedMessage, table[row1][4])
			}
			if col2 > 0 {
				encryptedMessage = append(encryptedMessage, table[row2][col2-1])
			} else {
				encryptedMessage = append(encryptedMessage, table[row2][4])
			}
		} else if col1 == col2 {
			if row1 > 0 {
				encryptedMessage = append(encryptedMessage, table[row1-1][col1])
			} else {
				encryptedMessage = append(encryptedMessage, table[4][col1])
			}
			if row2 > 0 {
				encryptedMessage = append(encryptedMessage, table[row2-1][col1])
			} else {
				encryptedMessage = append(encryptedMessage, table[4][col2])
			}
		} else {
			dist := col1 - col2
			encryptedMessage = append(encryptedMessage, table[row1][abs(col1-dist)])
			encryptedMessage = append(encryptedMessage, table[row2][abs(col2+dist)])
		}
	}
	table.show()
	return string(encryptedMessage)
}

// makeTable cria e popula uma table.
func makeTable(key string) table {

	usedLetters := make(map[byte]bool)
	table := [5][5]byte{}

	letters := make([]byte, 0, 25)

	for _, v := range key {
		b := byte(v)
		if !usedLetters[b] {
			letters = append(letters, b)
			usedLetters[b] = true
		}
	}

	for i := 'A'; i <= 'Z'; i++ {
		b := byte(i)
		if len(letters) >= 25 {
			break
		}

		if b == 'J' {
			continue
		}

		if !usedLetters[b] {
			letters = append(letters, b)
			usedLetters[b] = true
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			table[i][j] = letters[(i*5)+j]
		}
	}

	return table
}

// show imprime table na saída padrão.
func (t *table) show() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%c ", t[i][j])
		}
		fmt.Println()
	}
}

// where procura por uma rune em table.
func (t *table) where(c byte) (int, int) {
	x, y := -1, -1
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if t[i][j] == c {
				x = i
				y = j
			}
		}
	}
	return x, y
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

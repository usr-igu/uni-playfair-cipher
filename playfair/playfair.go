package playfair

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type KeyTable [5][5]byte

func Encrypt(msg string, table KeyTable) string {

	// todo: Tratar letras repetidas.
	msg = strings.Replace(msg, " ", "", -1)  // Remove todos os espaços.
	msg = strings.Replace(msg, "Z", "K", -1) // Trocando todos os Zs por Ks.
	msg = strings.ToUpper(msg)               // Tudo em caixa alta.
	if len(msg)%2 != 0 {                     // Se  o número de caracteres da msg não for par fazemos ser!
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

func Decrypt(msg string, table KeyTable) string {

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
	return string(encryptedMessage)
}

func NewKeyTable(key string) KeyTable {

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

		if b == 'Z' {
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

func LoadKeyTableFromFile(filename string) (KeyTable, error) {

	table := KeyTable{}

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return table, err
	}

	fileFilteredBytes := fileBytes[:0]
	for _, v := range fileBytes {
		if v != ' ' && v != '\n' {
			fileFilteredBytes = append(fileFilteredBytes, v)
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			table[i][j] = fileFilteredBytes[(i*5)+j]
		}
	}

	return table, nil
}

func (t *KeyTable) where(c byte) (int, int) {
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

func (t KeyTable) String() string {
	tableStr := make([]byte, 0, 32)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j < 4 {
				tableStr = append(tableStr, t[i][j])
				tableStr = append(tableStr, ' ')
			} else {
				tableStr = append(tableStr, t[i][j])
				tableStr = append(tableStr, '\n')
			}
		}
	}
	return fmt.Sprintf("%s", tableStr)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

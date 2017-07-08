package main

import (
	"flag"
	"fmt"

	"github.com/fuzzyqu/playfair-cipher/playfair"
)

func main() {

	msg := flag.String("msg", "CRIPTOGRAFIA", "Mensagem que vai ser criptografada/descriptografada")
	key := flag.String("key", "PLAYFAIR", "Keyword usada para criptografar/descriptografar uma mensagem")
	flag.Parse()

	table := playfair.NewKeyTable(*key)
	//table, _ := playfair.LoadKeyTableFromFile("keytable.txt")

	encryptedMessage := playfair.Encrypt(*msg, table)
	decryptedMessage := playfair.Decrypt(encryptedMessage, table)

	fmt.Println("Tabela de Playfair.")
	fmt.Print(table)

	fmt.Printf("Original message: %s\n", *msg)
	fmt.Printf("Encrypted message: %s\n", encryptedMessage)
	fmt.Printf("Decrypted message: %s\n", decryptedMessage)

}

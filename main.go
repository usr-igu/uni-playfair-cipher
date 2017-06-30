package main

import (
	"flag"
	"fmt"

	"github.com/fuzzyqu/playfair-cipher/playfair"
)

func main() {

	msg := flag.String("msg", "DOTINHA DE CADA DIA", "Mensagem que vai ser criptografada/descriptografada")
	key := flag.String("key", "NARUTO", "Keyword usada para criptografar/descriptografar uma mensagem")
	flag.Parse()

	table := playfair.MakeKeyTable(*key)

	encryptedMessage := playfair.Encrypt(*msg, *key, table)
	decryptedMessage := playfair.Decrypt(encryptedMessage, *key, table)

	fmt.Println("Tabela de Playfair.")
	fmt.Print(table)

	fmt.Printf("Original message: %s\n", *msg)
	fmt.Printf("Encrypted message: %s\n", encryptedMessage)
	fmt.Printf("Decrypted message: %s\n", decryptedMessage)
}

package main

import (
	"flag"
	"fmt"

	"github.com/fuzzyqu/playfair-cipher/playfair"
)

func main() {

	msg := flag.String("msg", "HIDETHEGOLDINTHETREXESTUMP", "Mensagem que vai ser criptografada/descriptografada")
	key := flag.String("key", "NARUTO", "Keyword usada para criptografar/descriptografar uma mensagem")
	flag.Parse()

	encryptedMessage := playfair.Encrypt(*msg, *key)
	decryptedMessage := playfair.Decrypt(encryptedMessage, *key)

	fmt.Printf("Original message: %s\n", *msg)
	fmt.Printf("Encrypted message: %s\n", encryptedMessage)
	fmt.Printf("Decrypted message: %s\n", decryptedMessage)
}

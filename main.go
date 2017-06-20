package main

import (
	"flag"
	"fmt"

	pf "github.com/fuzzyquanta/br/ufmt/ic/criptografia/atividade01/playfair"
)

func main() {

	msg := flag.String("msg", "Hidethegoldinthetreestump", "Mensagem que vai ser criptografada/descriptografada")
	key := flag.String("key", "PLAYFAIREXAMPLE", "Keyword usada para criptografar/descriptografar uma mensagem")
	flag.Parse()

	encodedMsg := pf.Playfair(*msg, *key)

	fmt.Printf("Encoded msg: %s\n", encodedMsg)
}

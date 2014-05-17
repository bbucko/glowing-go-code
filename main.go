package main

import "fmt"
import "github.com/dgryski/dkeyczar"

func main() {
	fmt.Printf("Hello, world.\n")

	plaintext := []byte("hello world")

	reader := dkeyczar.NewFileReader("./")
	crypter, _ := dkeyczar.NewCrypter(reader)

	ciphertext, _ := crypter.Encrypt(plaintext)

	fmt.Print(ciphertext)
}

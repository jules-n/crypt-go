package main

import (
	"fmt"
	"strconv"
)

func main() {
	text := Read("input.txt")

	privateKey, publicKey, r, s, err := GenerateECDSAKeysAndSign(text)
	if err != nil {
		fmt.Println("Помилка при генерації ключів та підпису:", err)
		return
	}

	fmt.Println("Приватний ключ:", privateKey)
	fmt.Println("Публічний ключ:", publicKey)
	fmt.Printf("Підпис (r, s): (%x, %x)\n", r, s)

	SavePublicKeyToFile(publicKey, "ecc.pub")
	SavePrivateKeyToFile(privateKey, "ecc.pem")

	key, encryptedText, err := Encrypt(text)

	if err != nil {
		fmt.Println("Помилка при aes шифруванні:", err)
		return
	}

	Write("aes.key", key)
	Write("encrypted.bin", encryptedText)

	decrypted, err := Decrypt(encryptedText, key)

	if err != nil {
		fmt.Println("Помилка при aes дешифруванні:", err)
		return
	}

	fmt.Println(strconv.Quote(string(decrypted)))
}

package main

import (
	"fmt"
	"io/ioutil"
)

func Read(inputFilename string) []byte {
	plaintext, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		fmt.Println("Помилка при зчитуванні вхідних даних з файлу:", err)
	}
	return plaintext
}

func Write(outputFilename string, ciphertext []byte) {
	err := ioutil.WriteFile(outputFilename, ciphertext, 0644)
	if err != nil {
		fmt.Println("Помилка при записі зашифрованих даних у файл:", err)
		return
	}
	fmt.Println("Зашифровані дані успішно записані у файл", outputFilename)
}

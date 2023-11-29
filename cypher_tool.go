package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("\nWelcome to the Cypher Tool!\n")

	toEncrypt, encoding, message := getInput()

	var result string

	switch encoding {
	case "ROT13":
		if toEncrypt {
			result = encrypt_rot13(message)
		} else {
			result = decrypt_rot13(message)
		}
	case "reverse":
		if toEncrypt {
			result = encrypt_reverse(message)
		} else {
			result = decrypt_reverse(message)
		}
	case "random":
		if toEncrypt {
			result = encryptRandom(message)
		} else {
			result = decryptRandom(message)
		}
	}

	if toEncrypt {
		fmt.Printf("\nEncrypted message using %s:\n", encoding)
	} else {
		fmt.Printf("\nDecrypted message using %s:\n", encoding)
	}
	fmt.Println(result)
}

func getInput() (toEncrypt bool, encoding string, message string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select operation (1/2):")
		fmt.Println("1. Encrypt.")
		fmt.Println("2. Decrypt.")

		operation, _ := reader.ReadString('\n')
		operation = strings.TrimSpace(operation)

		if operation == "1" {
			toEncrypt = true
			break
		} else if operation == "2" {
			toEncrypt = false
			break
		} else {
			fmt.Println("\nInvalid operation. Please select 1 for Encrypt or 2 for Decrypt.\n")
		}
	}

	for {
		fmt.Println("\nSelect cypher (1/3):")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Println("3. Random.")
		encryptionType, _ := reader.ReadString('\n')
		encryptionType = strings.TrimSpace(encryptionType)

		switch encryptionType {
		case "1":
			encoding = "ROT13"
		case "2":
			encoding = "reverse"
		case "3":
			encoding = "random"
		default:
			fmt.Println("\nInvalid encryption type. Please select 1 for ROT13, 2 for Reverse or 3 for Random.")
		}
		if encoding != "" {
			break
		}
	}

	for {
		fmt.Print("\nEnter the message: \n")
		message, _ = reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if len(message) >= 1 {
			break
		} else {
			fmt.Println("Invalid message. Must contain one or more characters")
		}
	}
	return toEncrypt, encoding, message
}

func encrypt_rot13(s string) string {
	var result string

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string('a' + (char-'a'+13)%26)
		} else if 'A' <= char && char <= 'Z' {
			result += string('A' + (char-'A'+13)%26)
		} else {
			result += string(char)
		}
	}

	return result
}

func decrypt_rot13(s string) string {
	return encrypt_rot13(s)
}

func encrypt_reverse(s string) string {
	var result string

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string('z' - (char - 'a'))
		} else if char >= 'A' && char <= 'Z' {
			result += string('Z' - (char - 'A'))
		} else {
			result += string(char)
		}
	}

	return result
}

func decrypt_reverse(s string) string {
	return encrypt_reverse(s)
}

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	codemap = "UmRSqPBolEXtZYAkJvcNHDWgIFumrsqpboleXtzyaKJvCnhDwgIf"
)

func encryptRandom(s string) string {
	var result string

	for _, char := range s {
		if index := strings.Index(letters, string(char)); index != -1 {
			result += string(codemap[index])
		} else {
			result += string(char)
		}
	}

	return result
}

func decryptRandom(s string) string {
	var result string

	for _, char := range s {
		if index := strings.Index(codemap, string(char)); index != -1 {
			result += string(letters[index])
		} else {
			result += string(char)
		}
	}

	return result
}

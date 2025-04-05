package utils

import "strings"

const (
	characters string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	divisor    int    = 62
)

func Encode(plain_text int) string {
	if plain_text < 62 {
		return string(characters[plain_text])
	}

	cipher := ""
	for plain_text >= divisor {
		remainder := plain_text % divisor
		cipher = string(characters[remainder]) + cipher
		plain_text /= divisor
	}
	cipher = string(characters[plain_text]) + cipher

	return cipher
}

func Decode(cipher string) int {
	plain_text := 0
	for _, value := range cipher {
		converted_value := strings.IndexRune(characters, value)
		plain_text = plain_text*divisor + converted_value
	}
	return plain_text
}

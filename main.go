package main

import (
	"fmt"
)

func main() {
	var length, useDigits, useSymbols = getPasswordParams()

	var password, err = generatePassword(length, useDigits, useSymbols)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	showPassword(length, useDigits, useSymbols, password)
	copyPassword(password)
}

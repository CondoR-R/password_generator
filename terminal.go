package main

import (
	"fmt"
	"os"
	"strconv"

	clipboard "github.com/atotto/clipboard"
)

// Получение параметор для генерации пароля из аргументов командной строки.
// Возвращает переданные параметры для генерации пароля: длина, нужны ли цифры,
// нужны ли специальные символы.
func getPasswordParams() (int, bool, bool) {
	var args = os.Args[1:]

	var (
		length     = 16
		useDigits  = false
		useSymbols = false
	)

	for _, arg := range args {
		switch arg {
		case "-digits", "-d":
			useDigits = true
		case "-symbols", "-s":
			useSymbols = true
		default:
			if num, err := strconv.Atoi(arg); err == nil {
				length = num
			}
		}
	}

	return length, useDigits, useSymbols
}

func showPassword(length int, useDigits, useSymbols bool, password string) {
	fmt.Printf("Длина пароля: %d\n", length)
	fmt.Printf("Использование цифр ")
	if useDigits {
		fmt.Printf("вкл\n")
	} else {
		fmt.Printf("выкл\n")
	}
	fmt.Printf("Использование специальных символов ")
	if useSymbols {
		fmt.Printf("вкл\n")
	} else {
		fmt.Printf("выкл\n")
	}
	fmt.Printf("Сгенерированный пароль: %s\n", password)
}

// Копирование в буфер обмена.
// Принимает пароль, который необходимо скопировать.
// Выводит в терминал сообщение об успешном/не успешном копировании.
func copyPassword(password string) {
	err := clipboard.WriteAll(password)
	if err != nil {
		fmt.Printf("Не удалось скопировать в буфер обмена: %s\n", err.Error())
	} else {
		fmt.Printf("Пароль скопирован в буфер обмена!\n")
	}
}

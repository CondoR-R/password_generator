package main

import (
	"fmt"
	"os"
	"strconv"

	clipboard "github.com/atotto/clipboard"
)

// Выводит сообщение с описанием программы
func showHelp() {
	var message = `
Программа для генерации паролей.
По умолчанию (без указания дополнительных аргументов при запуске программы)
генерирует пароль длиной 16 символов из сточных и прописных латинских букв.

Возможные аргументы:

[длина] - целое число, длина пароля (рекомендуемое значение 16+ символов)
-d / -digits - включить в генерацию пароля цифры
-s / -symbols - включить в генерацию пароля специальные символы
-c / -copy - автоматически скопировать в буфер обмена сгенерированный пароль
`
	fmt.Println(message)
}

// Получение параметор для генерации пароля из аргументов командной строки.
// Возвращает переданные параметры для генерации пароля: длина, нужны ли цифры,
// нужны ли специальные символы.
func getPasswordParams() (length int, useDigits, useSymbols, isCopy, isHelp bool) {
	var args = os.Args[1:]

	length = 16
	useDigits = false
	useSymbols = false
	isCopy = false
	isHelp = false

	for _, arg := range args {
		switch arg {
		case "-digits", "-d":
			useDigits = true
		case "-symbols", "-s":
			useSymbols = true
		case "-copy", "-c":
			isCopy = true
		case "-help", "-h":
			isHelp = true
		default:
			if num, err := strconv.Atoi(arg); err == nil {
				length = num
			}
		}
	}

	return
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

// Вывод паузы для предотвращения мгновенного завершения программы
func showPause() {
	fmt.Println("Нажмите Enter для выхода...")
	fmt.Scanln()
}

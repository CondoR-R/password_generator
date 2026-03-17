package main

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"strings"
)

// Получение алфавита пароля исходя из условия.
// Принимает 2 булевых параметра: нужны ли цифры и специальные символы в алфавите пароля.
// Возвращает алфавит пароля.
func getAlphabet(useDigits, useSymbols bool) string {
	// алфавит из строчных и прописных латинских букв без похожих символов (l I, O 0)
	var alphabet = "qwertyuiopasdfghjkzxcvbnmQWERTYUPASDFGHJKLZXCVBNM"

	if useDigits {
		alphabet += "123456789"
	}
	if useSymbols {
		alphabet += "!@#$%^&*()_+-=[]{}|;:,.<>?"
	}

	return alphabet
}

// Генерация пароля.
// Принимает длину и алфавит пароля.
// Возвращет пароль и ошибку.
func fillPassword(length int, alphabet string) (string, error) {
	var (
		alphabetLen = uint64(len(alphabet)) // длина алфавита
		// вычисление максимального числа кратного длине алфавита для равномерного распределения
		max      = ^uint64(0) - (^uint64(0) % alphabetLen)
		password strings.Builder
	)

	// выделяем место в памяти для пароля длиной length
	password.Grow(length)

	for i := 0; i < length; {
		// Создаем массив на 8 байт, так как именно 8 байт необходимо для uint64
		var bytes [8]byte

		// Заполняем массив случаными байтами и проверяем на возможную ошибку
		if _, err := rand.Read(bytes[:]); err != nil {
			return "", err
		}

		// Преобразуем байты в 64-битное число
		var num = binary.LittleEndian.Uint64(bytes[:])

		if num < max {
			var index = num % alphabetLen
			password.WriteByte(alphabet[index])
			i++
		}
	}

	return password.String(), nil
}

// Генератор паролей.
// Принимает длину пароля типа int и bool параметры, необходимо ли включать цифры
// и спец. символы.
// Возвращает пароль и информацию об ошибке.
func generatePassword(length int, useDigits, useSymbols bool) (string, error) {
	var alphabet = getAlphabet(useDigits, useSymbols)

	// для избежания возможных ошибок при ином способе формирования алфавита в будущем
	if alphabet == "" {
		return "", errors.New("Алфавит пароля не должен быть пустым!")
	}

	var password, err = fillPassword(length, alphabet)

	if err != nil {
		return "", err
	}

	return password, nil
}

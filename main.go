package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// 1. Имя пользователя из переменной окружения
	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("USERNAME")
	}
	fmt.Printf("Пользователь: %s\n", username)

	// 2. Аргументы командной строки
	fmt.Printf("Аргументы: %v\n", os.Args)

	// 3. Версия Go
	fmt.Printf("Версия Go: %s\n", runtime.Version())
}

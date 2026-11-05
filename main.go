package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Имя пользователя: %s\n", os.Getenv("USER"))

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Аргументы CLI: (не переданы)")
	} else {
		fmt.Println("Аргументы CLI:")
		for i, arg := range args {
			fmt.Printf("  [%d] %s\n", i, arg)
		}
	}

	// 3. Версия Go
	fmt.Printf("Версия Go: %s\n", runtime.Version())
}
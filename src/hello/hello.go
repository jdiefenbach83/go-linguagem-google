package main

import "fmt"

func main() {
	name := "Jefferson"
	version := 1.2

	fmt.Println("Hello", name)
	fmt.Println("Program version", version)

	fmt.Println("1 - Inicial monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var command int
	fmt.Scan(&command)

	fmt.Println("Variable command address", &command)
	fmt.Println("Chosen command: ", command)
}

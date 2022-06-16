package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoring = 2
const delay = 5

func main() {
	showIntroduction()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Jefferson"
	version := 1.2

	fmt.Println("Hello", name)
	fmt.Println("Program version", version)
}

func showMenu() {
	fmt.Println("1 - Inicial monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var readCommand int
	fmt.Scan(&readCommand)
	fmt.Println("Chosen command: ", readCommand)
	fmt.Println("")

	return readCommand
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}

	for i := 0; i < monitoring; i++ {
		for pos, site := range sites {
			fmt.Println("Testando site", pos, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site: ", site, " was loaded with success!")
	} else {
		fmt.Println("Site: ", site, " has problems! Status code:", response.StatusCode)
	}
}

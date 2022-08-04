package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 2
const delay = 5

func main() {

	showIntro()

	for {

		showMenu()

		command := readCommand()

		switch command {
		case 1:
			starMonitoring()
		case 2:
			fmt.Println("Exibindo logs ...")
			printLog()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Leonardo"
	version := 1.1
	age := 30

	fmt.Println("Olá", name, "sua idade é", age)
	fmt.Println("Este programa está na versão", version)
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O valor da variável comando é:", command)

	return command
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func starMonitoring() {
	fmt.Println("Monitorando ...")

	sites := readFileSite()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)

			time.Sleep(delay * time.Second)
			fmt.Println("")
		}
	}
	fmt.Println("")
}

func testSite(site string) {
	result, _ := http.Get(site)

	if result.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", result.StatusCode)
		registerLog(site, false)
	}
}

func readFileSite() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		sites = append(sites, row)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2007 14:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + " ")

	file.Close()
}

func printLog() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}

package main

import (
	"bufio"
	"fmt" //impressão e formatação de msg
	"io"
	"net/http" // requisões web
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3 //constantes não podem ser alteradas
const delay = 5

func main() {

	exibeIntrodução()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) //tudo correu bem

		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) //indica que ocorreu algum erro - 255
		}

	}
}

func exibeIntrodução() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://httpbin.org/status/404", "https://httpbin.org/status/200", "https://httpbin.org/status/404"}

	sites := leSitesDoArquivo()

	// fmt.Println(sites)

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testanto site", i, ":", site)
			testaSite(site)

		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs(){

	fmt.Println("Exibindo Logs...")

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	
	fmt.Println(string(arquivo))
}

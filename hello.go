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

const monitoriamentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {

		exibeMenu()

		switch leComando() {
		case 1:
			iniciarMonitoriamento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}
	// if comando == 1 {
	// 	fmt.Println("Monitorando........")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs..............")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo..............")
	// } else {
	// 	fmt.Println("Não conheço esse comando")
	// }

}

func exibeIntroducao() {

	nome := "Allan"
	versao := 1.1
	fmt.Println("Olá, sr.,", nome)
	fmt.Println("Este programa está na versão", versao)

	// var nome string = "Allan"
	// var idade int = 36
	// var versao float32 = 1.1
	// fmt.Println("Olá, sr.,", nome, "sua idade é", idade)
	// fmt.Println("Este programa está na versão", versao)
	// fmt.Println(reflect.TypeOf(nome))

	// nome := "Allan"
	// idade := 36
	// versao := 1.1
	// fmt.Println("Olá, sr.,", nome, "sua idade é", idade)
	// fmt.Println("Este programa está na versão", versao)
	// fmt.Println(reflect.TypeOf(nome))

}

func exibeMenu() {

	fmt.Println("1- Iniciar Monitoriamento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}

func leComando() int {

	var comando int
	fmt.Scan(&comando)
	//fmt.Println("O endereço da minha variavel comando é", &comando)
	//fmt.Println("O comando escolhido foi", comando)

	return comando
}

func iniciarMonitoriamento() {

	fmt.Println("Monitorando...")
	//sites := []string{"(https://www.alura.com.br", "https://www.caelum.com.br", "https://httpbin.org/status/200"}

	sites := leSiteDoArquivo()

	for i := 0; i < monitoriamentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
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

func leSiteDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocooreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			fmt.Println("Ocooreu um erro", err)
			break
		}

	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocooreu um erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

	fmt.Println(arquivo)

}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))

}

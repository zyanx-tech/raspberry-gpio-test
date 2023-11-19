package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/stianeikeland/go-rpio/v4"
)

func processCommand(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Formato inválido. Use: -io <num1> <num2> ... -up/-down")
		return
	}

	// Encontra a posição do comando -up ou -down
	cmdPos := -1
	for i, part := range parts {
		if part == "-up" || part == "-down" {
			cmdPos = i
			break
		}
	}

	if cmdPos == -1 {
		fmt.Println("Comando -up ou -down não encontrado")
		return
	}

	// Lê os números de GPIO
	gpioNumsStr := parts[1:cmdPos]
	command := parts[cmdPos]

	for _, numStr := range gpioNumsStr {
		gpioNum, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Número de GPIO inválido: '%s'\n", numStr)
			continue
		}

		pin := rpio.Pin(gpioNum)
		switch command {
		case "-down":
			pin.Output()
			pin.High()
		case "-up":
			pin.Output()
			pin.Low()
		default:
			fmt.Println("Comando inválido. Use: -up ou -down")
			return
		}
	}
}

func showHelp() {
	fmt.Println("Uso do Programa:")
	fmt.Println("./main -io <num1> <num2> ... -up/-down : Executa o comando e sai do programa")
	fmt.Println("./main -i                              : Entra no modo interativo")
	fmt.Println("./main -h                              : Mostra esta mensagem de ajuda")
}

func main() {
	// Inicializa go-rpio
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return
	}
	defer rpio.Close()

	args := os.Args[1:]

	if len(args) > 0 {
		switch args[0] {
		case "-i":
			// Modo interativo
			// Modo interativo
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("Digite o comando (ex: -io 5 6 7 12 77 2 -up): ")
				text, _ := reader.ReadString('\n')
				text = strings.TrimSpace(text)
				parts := strings.Fields(text)

				processCommand(parts)
			}
		case "-h":
			// Mostra a ajuda
			showHelp()
			return
		default:
			// Executa o comando dos argumentos
			processCommand(args)
			return // Sai do programa após executar o comando
		}
	} else {
		fmt.Println("Nenhum argumento fornecido. Entrando no modo interativo.")
		main() // Chama main novamente para entrar no modo interativo
	}
}

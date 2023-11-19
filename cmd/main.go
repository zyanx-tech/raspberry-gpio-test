package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	// Inicializa go-rpio
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return
	}
	defer rpio.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Digite o comando (ex: -io 5 6 7 12 77 2 -up): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		parts := strings.Fields(text)

		if len(parts) < 3 {
			fmt.Println("Formato inválido. Use: -io <num1> <num2> ... -up/-down")
			continue
		}

		// Certifique-se de que a entrada começa com '-io'
		if parts[0] != "-io" {
			fmt.Println("Comando deve começar com '-io'")
			continue
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
			continue
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
				break
			}
		}
	}
}

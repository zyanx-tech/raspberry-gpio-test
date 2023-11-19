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
		fmt.Print("Digite o comando (ex: IO 5,6,7,8,12 UP): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		parts := strings.Fields(text)

		if len(parts) < 2 {
			fmt.Println("Formato inválido. Use: IO<num1,num2,...> UP/DOWN")
			continue
		}

		// Certifique-se de que a parte do número começa com 'IO'
		if !strings.HasPrefix(parts[0], "IO") {
			fmt.Println("Comando deve começar com 'IO'")
			continue
		}

		// Removendo o prefixo 'IO' e dividindo os números
		numsPart := strings.TrimPrefix(parts[0], "IO")
		gpioNumsStr := strings.Split(numsPart, ",")
		command := parts[1]

		for _, numStr := range gpioNumsStr {
			gpioNum, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				fmt.Printf("Número de GPIO inválido: %s\n", numStr)
				continue
			}

			pin := rpio.Pin(gpioNum)
			switch command {
			case "DOWN":
				pin.Output()
				pin.High()
			case "UP":
				pin.Output()
				pin.Low()
			default:
				fmt.Println("Comando inválido. Use: UP ou DOWN")
				break
			}
		}
	}
}

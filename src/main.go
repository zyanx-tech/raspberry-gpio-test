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
		fmt.Print("Digite o comando (ex: gpio5 UP): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		parts := strings.Split(text, " ")

		if len(parts) != 2 {
			fmt.Println("Formato inválido. Use: gpio<num> UP/DOWN")
			continue
		}

		gpioNum, err := strconv.Atoi(parts[0][4:])
		if err != nil {
			fmt.Println("Número de GPIO inválido")
			continue
		}

		command := parts[1]
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
		}
	}
}

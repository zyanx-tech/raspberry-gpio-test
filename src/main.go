package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	// Inicializa periph.io
	if _, err := host.Init(); err != nil {
		fmt.Println(err)
		return
	}

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

		gpioPin := parts[0]
		command := parts[1]

		pin := gpioreg.ByName(gpioPin)
		if pin == nil {
			fmt.Printf("GPIO %s não encontrada.\n", gpioPin)
			continue
		}

		switch command {
		case "UP":
			pin.Out(gpio.High)
		case "DOWN":
			pin.Out(gpio.Low)
		default:
			fmt.Println("Comando inválido. Use: UP ou DOWN")
		}
	}
}

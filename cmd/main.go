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
		fmt.Println("Invalid format. Use: -io <num1> <num2> ... -up/-down")
		return
	}

	cmdPos := -1
	for i, part := range parts {
		if part == "-up" || part == "-down" {
			cmdPos = i
			break
		}
	}

	if cmdPos == -1 {
		fmt.Println("-up or -down command not found")
		return
	}

	gpioNumsStr := parts[1:cmdPos]
	command := parts[cmdPos]

	for _, numStr := range gpioNumsStr {
		gpioNum, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Invalid GPIO number: '%s'\n", numStr)
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
			fmt.Println("Invalid command. Use: -up or -down")
			return
		}
	}
}

func showHelp() {
	fmt.Println("Use of the Program:")
	fmt.Println("./main -io <num1> <num2> ... -up/-down : Execute the command and exit the program")
	fmt.Println("./main -i                              : Enter interactive mode")
	fmt.Println("./main -h                              : Show this help message")
}

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return
	}
	defer rpio.Close()

	args := os.Args[1:]

	if len(args) > 0 {
		switch args[0] {
		case "-i":
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("Digite o comando (ex: -io 5 6 7 12 77 2 -up): ")
				text, _ := reader.ReadString('\n')
				text = strings.TrimSpace(text)
				parts := strings.Fields(text)

				processCommand(parts)
			}
		case "-h":
			showHelp()
			return
		default:
			processCommand(args)
			return
		}
	} else {
		fmt.Println("No arguments provided. Exit.")
	}
}

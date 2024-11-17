package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func validate(args []string) {
	for i := 1; i < len(args); i++ {
		if len(args[i]) > 12 && len(args[i]) < 20 {
			if len(args[i]) > 12 && lunacheckS(args[i]) {
				fmt.Println("OK")
			} else {
				fmt.Println("INCORRECT")
				os.Exit(1)
			}
		} else {
			fmt.Println("Invalid input for validate command.")
			os.Exit(1)
		}
	}
}

func validateStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		cardNumbers := strings.Fields(line)
		for _, number := range cardNumbers {
			if len(number) > 12 && len(number) < 20 {
				if len(number) > 12 && lunacheckS(number) {
					fmt.Println("OK")
				} else {
					fmt.Println("INCORRECT")
					os.Exit(1)
				}
			} else {
				fmt.Println("Invalid input for validate command.")
				os.Exit(1)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}

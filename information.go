package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	brandsFile  = "brands.txt"
	issuersFile = "issuers.txt"
)

func information(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: <command> [card number ...] [--stdin]")
		return
	}

	brands, err := readFile(brandsFile)
	if err != nil {
		fmt.Println("Error reading brands file:", err)
		return
	}
	issuers, err := readFile(issuersFile)
	if err != nil {
		fmt.Println("Error reading issuers file:", err)
		return
	}
	for idx, val := range args {
		if (idx > 0 && len(val) > 12 && len(val) < 17) || (idx > 0 && (args[idx] == "--brands=brands.txt")) || (idx > 0 && (args[idx] == "--issuers=issuers.txt")) {
			var inform string
			if len(val) > 12 && len(val) < 17 {
				for _, char := range val {
					if !(char >= '0' && char <= '9') {
						fmt.Println("Invalid input")
						os.Exit(1)
						continue
					} else if char >= '0' && char <= '9' {
						inform = AppendRune(inform, char)
					}
				}
			}

			if len(args) > 2 {
				if !((args[1] == "--brands=brands.txt") || (args[1] == "--issuers=issuers.txt")) {
					fmt.Println("Invalid input for flag command.")
					os.Exit(1)
				}
			}
			if idx > 3 {
				for _, char := range val {
					if !(char >= '0' && char <= '9') {
						fmt.Println("Invalid input")
						os.Exit(1)
					}
				}
			}
			if !(inform == "") {
				fmt.Println(inform)
			}
			if !(inform == "") {
				if lunacheckS(inform) {
					fmt.Println("Correct: yes")
					if (args[1] == "--brands=brands.txt") || (args[2] == "--brands=brands.txt") {
						brand := findBrand(inform, brands)
						if brand == "VISA" {
							if len(inform) != 16 {
								fmt.Println("Invalid card number")
								os.Exit(1)
							}
						}
						if brand == "MASTERCARD" {
							if len(inform) != 16 {
								fmt.Println("Invalid card number")
								os.Exit(1)
							}
						}
						if brand == "AMEX" {
							if len(inform) != 15 {
								fmt.Println("Invalid card number")
								os.Exit(1)
							}
						}
						fmt.Printf("Card Brand: %s\n", brand)
					}
					if (args[1] == "--issuers=issuers.txt") || (args[2] == "--issuers=issuers.txt") {
						issuer := findIssuer(inform, issuers)
						fmt.Printf("Card Issuer: %s\n", issuer)
					}
				} else {
					fmt.Println("Correct: no")
					if (args[1] == "--brands=brands.txt") || (args[2] == "--brands=brands.txt") {
						fmt.Println("Card Brand: -")
					}
					if (args[1] == "--issuers=issuers.txt") || (args[2] == "--issuers=issuers.txt") {
						fmt.Println("Card Issuer: -")
					}
				}
			}

		} else if idx > 2 && len(val) > 16 {
			fmt.Println("Invalid input")
			os.Exit(1)
		}
	}
}

func processStdin(args []string) {
	brands, err := readFile(brandsFile)
	if err != nil {
		fmt.Println("Error reading brands file:", err)
		return
	}
	issuers, err := readFile(issuersFile)
	if err != nil {
		fmt.Println("Error reading issuers file:", err)
		return
	}
	var need string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		cardNumbers := strings.Fields(line)

		for idx, val := range args {
			if len(args) < 4 {
				fmt.Println("Invalid number for command, add more command.")
				os.Exit(1)
			}
			if len(args) > 4 {
				fmt.Println("Invalid number for command, nery big amount command.")
				os.Exit(1)
			}
			if len(args) > 2 {
				if !((args[1] == "--brands=brands.txt") || (args[1] == "--issuers=issuers.txt") || (args[1] == "--stdin")) {
					fmt.Println("Invalid input for flag command.")
					os.Exit(1)
				}
				if !((args[2] == "--brands=brands.txt") || (args[2] == "--issuers=issuers.txt") || (args[2] == "--stdin")) {
					fmt.Println("Invalid input for flag command.")
					os.Exit(1)
				}
				if !((args[3] == "--brands=brands.txt") || (args[3] == "--issuers=issuers.txt") || (args[3] == "--stdin")) {
					fmt.Println("Invalid input for flag command.")
					os.Exit(1)
				}
			}
			if idx > 3 {
				for _, char := range val {
					if !(char >= '0' && char <= '9') {
						fmt.Println("Invalid input")
						os.Exit(1)
					}
				}
			}

		}

		for _, number := range cardNumbers {
			need = number
			fmt.Println(need)
			if lunacheckS(number) {

				fmt.Println("Correct: yes")
				brand := findBrand(number, brands)
				if brand == "VISA" {
					if len(need) != 16 {
						fmt.Println("Invalid card number")
						os.Exit(1)
					}
				}
				if brand == "MASTERCARD" {
					if len(need) != 16 {
						fmt.Println("Invalid card number")
						os.Exit(1)
					}
				}
				if brand == "AMEX" {
					if len(need) != 15 {
						fmt.Println("Invalid card number")
						os.Exit(1)
					}
				}
				issuer := findIssuer(number, issuers)
				fmt.Printf("Card Brand: %s\n", brand)
				fmt.Printf("Card Issuer: %s\n", issuer)
			} else {

				fmt.Println("Correct: no")
				fmt.Println("Card Brand: -")
				fmt.Println("Card Issuer: -")
				continue
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
	}
}

func findBrand(cardNumber string, brands [][2]string) string {
	for _, entry := range brands {
		prefix := entry[1]
		if startsWith(cardNumber, prefix) {
			return entry[0]
		}
	}
	return "-"
}

func findIssuer(cardNumber string, issuers [][2]string) string {
	for _, entry := range issuers {
		prefix := entry[1]
		if startsWith(cardNumber, prefix) {
			return entry[0]
		}
	}
	return "-"
}

func startsWith(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

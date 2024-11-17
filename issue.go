package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GenerateRandomCard(brand, issuer string, brands, issuers [][2]string) (string, error) {
	Error := "error"
	brandPrefix := findPrefix(brand, brands)
	if brandPrefix == "" {
		return "", fmt.Errorf("brand not found")
	}
	issuerPrefix := findPrefix(issuer, issuers)
	if issuerPrefix == "" {
		return "", fmt.Errorf("issuer not found")
	}

	if checkPrefixEquality(brandPrefix, issuerPrefix) {
		for {
			cardNumber := generateCardNumber(brandPrefix, issuerPrefix)
			if lunacheckS(cardNumber) {
				return cardNumber, nil
			}
		}
	}
	return Error, nil
}

func generateCardNumber(brandPrefix, issuerPrefix string) string {
	prefix := issuerPrefix

	cardNumber := prefix
	if brandPrefix == "34" || brandPrefix == "37" {
		for len(cardNumber) < 15 {
			cardNumber += byteToString(randomByte())
		}
		return cardNumber[:15]
	} else {
		for len(cardNumber) < 16 {
			cardNumber += byteToString(randomByte())
		}
		return cardNumber[:16]
	}
}

func issue(args []string) {
	if len(os.Args) < 5 {
		fmt.Println("Usage: <command> --brands=brands.txt --issuers=issuers.txt --brand=<brand> --issuer=<issuer>")
		os.Exit(1)
	}
	if len(os.Args) > 6 {
		fmt.Println("Usage: <command> --brands=brands.txt --issuers=issuers.txt --brand=<brand> --issuer=<issuer> lox")
		os.Exit(1)
	}
	brand := ""
	issuer := ""

	if len(args) == 5 {
		if !((args[1] == "--brands=brands.txt") || (args[1] == "--issuers=issuers.txt")) {
			fmt.Println("Invalid input for flag command.")
			os.Exit(1)
		}
		if !((args[2] == "--brands=brands.txt") || (args[2] == "--issuers=issuers.txt")) {
			fmt.Println("Invalid input for flag command.")
			os.Exit(1)
		}

	}

	for _, arg := range os.Args {
		if len(arg) > 7 && compareStrings(arg[:7], "--brand") {
			brand = arg[8:]
		}
		if len(arg) > 8 && compareStrings(arg[:8], "--issuer") {
			issuer = arg[9:]
		}
	}

	if brand == "" || issuer == "" {
		fmt.Println("Brand and issuer must be specified.")
		os.Exit(1)
	}

	brands, err := readFile(brandsFile)
	if err != nil {
		fmt.Println("Error reading brands file:", err)
		os.Exit(1)
	}

	issuers, err := readFile(issuersFile)
	if err != nil {
		fmt.Println("Error reading issuers file:", err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	cardNumber, err := GenerateRandomCard(brand, issuer, brands, issuers)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if lunacheckS(cardNumber) {
		fmt.Println(cardNumber)
	} else {
		fmt.Println("Generated card number is not valid.")
		os.Exit(1)
	}
}

func findPrefix(name string, data [][2]string) string {
	for _, entry := range data {
		if compareStrings(entry[0], name) {
			return entry[1]
		}
	}
	return ""
}

func checkPrefixEquality(brandPrefix, issuerPrefix string) bool {
	brandLen := len(brandPrefix)
	issuerLen := len(issuerPrefix)

	if brandLen == 1 || brandLen == 2 {
		if issuerLen >= 1 && issuerPrefix[0] == brandPrefix[0] {
			return true
		}
		fmt.Println("Error: Brand prefix does not match the issuer prefix.")
		return false
	} else {
		fmt.Println("Error: Brand prefix length must be 1 or 2.")
	}
	return true
}

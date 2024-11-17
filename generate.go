package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var Count int

func RandomCard(result string, count int) []string {
	var numbers []string

	if count == 4 {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 10; k++ {
					for l := 0; l < 10; l++ {
						number := fmt.Sprintf("%d%d%d%d", i, j, k, l)
						concatenated := result + number
						numbers = append(numbers, concatenated)
					}
				}
			}
		}
	}
	if count == 3 {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 10; k++ {
					number := fmt.Sprintf("%d%d%d", i, j, k)
					concatenated := result + number
					numbers = append(numbers, concatenated)
				}
			}
		}
	}
	if count == 2 {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				number := fmt.Sprintf("%d%d", i, j)
				concatenated := result + number
				numbers = append(numbers, concatenated)
			}
		}
	}
	if count == 1 {
		for i := 0; i < 10; i++ {
			number := fmt.Sprintf("%d", i)
			concatenated := result + number
			numbers = append(numbers, concatenated)
		}
	}
	return numbers
}

func generate(args []string) {
	if len(args) != 2 || len(args[1]) < 12 {
		fmt.Println("Invalid inputxdwx")
		os.Exit(1)
		return
	}
	var asterics string
	var result string
	count := 4
	for _, char := range args {
		if len(char) > 16 {
			fmt.Println("Invalid input")
			os.Exit(1)
		}
	}
	for _, char := range args[1] {

		if !(char >= '0' && char <= '9') && char != '*' {
			fmt.Println("Invalid input")
			os.Exit(1)
		}

		asterics = AppendRune(asterics, char)
	}
	reversed := reverseString(asterics)
	if containsNumberFollowedByAsterisk(reversed) {
		fmt.Println("String contains a number followed by an asterisk.")
	} else {
		for idx, char := range args[1] {
			if idx < len(args[1])-4 {
				if char == '*' {
					fmt.Println("Invalid input")
					os.Exit(1)
					return
				}
				if !(char >= '0' && char <= '9') && char != '*' {
					fmt.Println("Invalidbf input")
					os.Exit(1)
				}
				if char >= '0' && char <= '9' {
					result = AppendRune(result, char)
				} else {
					os.Exit(1)
					return
				}
			} else if char >= '0' && char <= '9' {
				count--
				result = AppendRune(result, char)
				if count < 0 {
					os.Exit(1)
					return
				}
			} else if char != '*' {
				os.Exit(1)
				return
			}
		}
		allNumbers := RandomCard(result, count)
		for _, num := range allNumbers {
			if lunacheckS(num) {
				fmt.Println(num)
			}
		}
	}
}

func pick(numbers []string) {
	if len(numbers) == 0 {
		fmt.Println("No number to picking")
		return
	}
	rand.Seed(time.Now().UnixNano())
	randomindex := rand.Intn(len(numbers))
	fmt.Println(numbers[randomindex])
}

func pickr(args []string) {
	count := 4
	var asterics string
	var result string
	for _, char := range args[2] {
		if !(char >= '0' && char <= '9') && char != '*' {
			fmt.Println("Invalid input")
			os.Exit(1)
		}

		asterics = AppendRune(asterics, char)
	}
	reversed := reverseString(asterics)
	if containsNumberFollowedByAsterisk(reversed) {
		fmt.Println("String contains a number followed by an asterisk.")
	} else {
		for idx, char := range args[2] {
			if idx < len(args[2])-4 {
				if char >= '0' && char <= '9' {
					result = AppendRune(result, char)
				} else {
					os.Exit(1)
					return
				}
			} else {
				if char >= '0' && char <= '9' {
					count--
					result = AppendRune(result, char)
					if count < 1 {
						fmt.Println("Error: number of asterisks more than 4")
						return
					}
				} else if char != '*' {
					fmt.Println("very Invalid pattern")
					return
				}
			}
		}
		allNumbers := RandomCard(result, count)
		var validNumbers []string
		for _, num := range allNumbers {
			if lunacheckS(num) {
				validNumbers = append(validNumbers, num)
			}
		}

		pick(validNumbers)
	}
}

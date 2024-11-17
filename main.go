package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || (args[0] != "issue" && args[0] != "information" && args[0] != "validate" && args[0] != "generate") {
		fmt.Println("Invalid input. Please enter a valid command (issue, information, validate, generate).")
		os.Exit(1)
	}

	switch args[0] {
	case "validate":
		if len(args) > 1 && args[1] == "--stdin" {
			validateStdin()
		} else if len(args) > 1 && args[0] == "validate" {
			validate(args)
		} else {
			fmt.Println("Invalid input for generate command.")
			os.Exit(1)
		}

	case "generate":
		if len(args) == 2 {
			generate(args)
		} else if len(args) == 3 && args[1] == "--pick" {
			pickr(args)
		} else {
			fmt.Println("Invalid input for generate command.")
			os.Exit(1)
		}

	case "information":
		if len(args) > 2 && (args[2] == "--stdin" || len(args) > 3 && args[3] == "--stdin") {
			processStdin(args)
		} else if len(args) > 2 {
			information(args)
		} else {
			fmt.Println("Invalid input for information command.")
			os.Exit(1)
		}

	case "issue":
		if len(args) > 1 {
			issue(args)
		} else {
			fmt.Println("Invalid input for issue command.")
			os.Exit(1)
		}

	default:
		fmt.Println("Invalid command")
	}
}

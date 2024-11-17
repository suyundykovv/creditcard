package main

import (
	"bufio"
	"math/rand"
	"os"
)

func AppendRune(s string, r rune) string {
	return s + string(r)
}

func byteToString(b byte) string {
	return string([]byte{b})
}

func readFile(filename string) ([][2]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][2]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := splitLine(line, ':')
		if len(parts) == 2 {
			data = append(data, [2]string{parts[0], parts[1]})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func splitLine(line string, delimiter rune) []string {
	var parts []string
	var current string
	for _, char := range line {
		if char == delimiter {
			parts = append(parts, current)
			current = ""
		} else {
			current += string(char)
		}
	}
	parts = append(parts, current)
	return parts
}

func compareStrings(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func randomByte() byte {
	return byte(rand.Intn(10) + '0')
}

func containsNumberFollowedByAsterisk(s string) bool {
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			if i+1 < len(s) && s[i+1] == '*' {
				return true
			}
		}
	}
	return false
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func reverseString(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

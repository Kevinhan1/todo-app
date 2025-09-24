package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi input untuk baca string dari user
func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

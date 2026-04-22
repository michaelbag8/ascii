package main

import (
	"fmt"
	"os"
	"strings"
)
func processBanner(input, fontFile string) error {
	data, err := os.ReadFile(fontFile)
	if err != nil {
		return err
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	var blocks [][]string

	const Height = 8
	const Size = 9
	const start = 1
	for i := start; i+Height <= len(lines); i += Size {
		blocks = append(blocks, lines[i:i+Height])
	}

	segments := strings.Split(input, "\\n")

	for s, segment := range segments {
		if segment == "" {
			if s != len(segments)-1 {
				fmt.Println()
			}
			continue
		}

		for row := 0; row < 8; row++ {
			var result strings.Builder

			for _, ch := range segment {
				index := int(ch - ' ')
				if index < 0 || index >= len(blocks) {
					continue
				}
				result.WriteString(blocks[index][row])
			}

			fmt.Println(result.String())
		}
	}

	return nil
}


func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . <text> [banner]")
		os.Exit(1)
	}

	fontFile := "standard.txt"
	if len(os.Args) == 3 {
		fontFile = os.Args[2] + ".txt"
	}

	input := os.Args[1]

	err := processBanner(input, fontFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

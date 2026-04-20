package main

import (
	"fmt"
	"os"
	"strings"
)

func ParseLoad() {

}
func main() {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error", err)
		return
	}
	content := string(data)

	input := strings.ReplaceAll(content, "\r\n", "\n")

	raw := strings.Split(input, "\n\n")

	var blocks [][]string

	for _, r := range raw {
		row := strings.Split(r, "\n")
		if len(row) < 8 {
			continue
		}

		blocks = append(blocks, row[:8])
	}

	for i, row := range blocks[33] {
		fmt.Printf("row %d: %s\n", i, row)
	}
}

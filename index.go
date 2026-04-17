package main

import (
	"fmt"
	"os"
	"strings"
)
func loadFont(filename string) ([][]string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("cannot open font file %q: %w", filename, err)
    }

    content := strings.ReplaceAll(string(data), "\r\n", "\n")
    rawBlocks := strings.Split(content, "\n\n")

    blocks := make([][]string, 0, 95)
    for _, raw := range rawBlocks {
        rows := strings.Split(raw, "\n")
        if len(rows) < 8 {
            continue
        }
        blocks = append(blocks, rows[:8])
    }

    if len(blocks) != 95 {
        return nil, fmt.Errorf("font file malformed: got %d blocks, need 95", len(blocks))
    }
    return blocks, nil
}


func renderLine(line string, blocks [][]string) {
    var sb strings.Builder
    for row := 0; row < 8; row++ {
        sb.Reset()
        for _, ch := range line {
            index := int(ch) - 32
            if index >= 0 && index < len(blocks) {
                sb.WriteString(blocks[index][row])
            } else {
                fmt.Fprintf(os.Stderr, "Error: unsupported character %q\n", ch)
                os.Exit(1)
            }
        }
        fmt.Println(sb.String())
    }
}

func renderAll(input string, blocks [][]string) {
    lines := strings.Split(input, "\\n")
    for _, line := range lines {
        if line == "" {
            fmt.Println()
            continue
        }
        renderLine(line, blocks)
    }
}


func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "Usage: go run . <string> ")
        os.Exit(1)
    }

    fontFile := "standard.txt"
    if len(os.Args) == 3 {
        fontFile = os.Args[2] + ".txt"
    }

    blocks, err := loadFont(fontFile)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    renderAll(os.Args[1], blocks)
}


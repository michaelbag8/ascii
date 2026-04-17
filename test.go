package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    data, err := os.ReadFile("min_banner.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")

    // Each character has 3 lines in this mini example
    blockSize := 3
    charMap := make(map[string][]string)

    // Step through lines in chunks of 3
    for i := 0; i < len(lines); i += blockSize {
        // Key: first line’s first character (simplified for demo)
        key := string(lines[i][0])
        charMap[key] = lines[i : i+blockSize]
    }

    // Print the map
    for k, v := range charMap {
        fmt.Println("Character:", k)
        for _, line := range v {
            fmt.Println(line)
        }
    }
}

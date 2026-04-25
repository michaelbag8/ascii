package main

import (
	"fmt"
	"os"
	"strings"
)

//reading the font file and cleaning it
func LoadFontFile(input, filename string)error{
	data, err := os.ReadFile(filename)
	if err != nil{
		fmt.Fprintf(os.Stderr, "Error reading fontfile", err)
		os.Exit(1)
	}

	content := string(data)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.TrimLeft(content, "\n")

	rawFile := strings.Split(content, "\n\n")
	
	var blocks [][]string
	for _, raw := range rawFile{
		lines := strings.Split(raw, "\n")
		if len(lines) < 8{
			continue
		}
		blocks = append(blocks, lines[:8])
	}

	//user input cleaning
	segments := strings.Split(input, "\\n")
	for i, segment := range segments{
		if segment == ""{
			if i < len(segments)-1{
				fmt.Println()
			}
			continue
		}
		
		//rendering row by row
		for row:=0; row < 8; row++{
			var result strings.Builder

			for _, ch := range segment{
				index := int(ch - ' ')
				if index >= 0 && index < len(blocks){
					result.WriteString(blocks[index][row])
				}else{
					fmt.Fprintf(os.Stderr, "Non Printable ASCII Character %q\n", ch)
					os.Exit(1)
				}
			}
			fmt.Println(result.String())
		}
		
	}
	return nil
}

func main(){
	if len(os.Args) < 2 || len(os.Args) > 3{
		fmt.Fprintf(os.Stderr, "Usage: go run . text or go run . text banner")
		os.Exit(1)
	}

	fontFile := "standard.txt"
	input := os.Args[1]

	if len(os.Args) == 3{
		fontFile = os.Args[2] + ".txt"
	}

	err := LoadFontFile(input, fontFile)
	if err != nil{
		fmt.Fprintf(os.Stderr, "Error loading banner", err)
		os.Exit(1)
	}
}
package main

import (
	"fmt"
	"os"
	"strings"
)

func hexdump(s string) ([]byte, error) {
	content, err := os.ReadFile(s)
	if err != nil {
		panic("Cannot open file")
	}

	return content, nil
}

func view(arr []byte) {
	reset := "\033[0m"
	green := "\033[32m"
	row_width := 16
	curr_row_width := 16
	seen_bytes := 0

	for i := 0; i < len(arr); i += row_width {
		fmt.Print(green + "| " + reset)
		if (len(arr) - seen_bytes) < row_width {
			curr_row_width = (len(arr) - seen_bytes)
		}
		row := arr[i:(seen_bytes + curr_row_width)]

		for i := 0; i < row_width; i++ {
			if i < curr_row_width {
				fmt.Printf("%02x ", row[i])
			} else {
				fmt.Print(strings.Repeat(" ", 3))
			}
		}

		fmt.Print(green + "| " + reset)
		fmt.Print(" ")

		for i := 0; i < row_width; i++ {
			if i < curr_row_width {
				if row[i] >= 0x20 && row[i] < 0x7f {
					fmt.Print(string(row[i]))
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(strings.Repeat(" ", 3))
			}
		}
		fmt.Print(green + "| " + reset)
		fmt.Print("\n")

		seen_bytes += row_width
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: iris <FILEPATH>")
		os.Exit(1)
	}

	var content, _ = hexdump(os.Args[1])
	view(content)
	fmt.Println("\033[0m")
}

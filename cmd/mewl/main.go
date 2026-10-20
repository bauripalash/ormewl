package main

import (
	"bufio"
	"fmt"
	"mewl/lexer"
	"os"
	"strings"
)

func repl() int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("𑄝੭ ")
		if !scanner.Scan() {
			break
		}
		if err := scanner.Err(); err != nil {
			break
		}

		tline := scanner.Text()
		if strings.TrimSpace(tline) == "" {
			continue
		}
		if tline == "q" {
			break
		}
		fmt.Println(lexer.NewDummyToken())

	}
	return 0
}

func runner() int {
	filename := os.Args[1]
	fmt.Printf("Running file : %s\n", filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file %s\n", filename)
		return 1
	}

	strContent := string(content)

	lx := lexer.NewMewlLexer(strContent)
	lx.ScanTokens()

	return 0
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ฅ[^•ﻌ•^]ฅ Mewl, The Programming Language of Cats'")
		status := repl()

		os.Exit(status)
	} else if len(os.Args) == 2 {
		status := runner()
		os.Exit(status)
	}
}

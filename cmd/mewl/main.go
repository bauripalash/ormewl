package main

import (
	"bufio"
	"fmt"
	"mewl/parser"
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

		parser.MwlTokenize(tline)

	}
	return 0
}

func runner() int {
	fmt.Println("Usage: mewl [filename]")
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

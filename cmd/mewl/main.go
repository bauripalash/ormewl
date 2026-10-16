package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ฅ[^•ﻌ•^]ฅ Mewl, The Programming Language of Cats'")
		fmt.Println("Usage: mewl <FILENAME>")
		os.Exit(1)
	}
}


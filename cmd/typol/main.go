package main

import (
	"fmt"
	"os"

	"git.sr.ht/grauwoelfchen/typol/typol/command"
)

func main() {
	out, err := command.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if out != "" {
		fmt.Println(out)
	}
}

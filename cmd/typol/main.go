package main

import (
	"fmt"
	"os"

	"git.sr.ht/grauwoelfchen/typol/typol/service"
)

func main() {
	out, err := service.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if out != "" {
		fmt.Println(out)
	}
}

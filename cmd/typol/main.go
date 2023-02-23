package main

import (
	"os"

	"git.sr.ht/grauwoelfchen/typol/typol/command"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

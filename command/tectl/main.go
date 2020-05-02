package main

import (
	"github.com/william20111/go-thousandeyes/command/tectl/cmd"
)

var (
	Version string
)

func main() {
	cmd.Execute(Version)
}

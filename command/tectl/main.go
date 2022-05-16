package main

import (
	"github.com/thousandeyes/go-thousandeyes/command/tectl/cmd"
)

var (
	Version string
)

func main() {
	cmd.Execute(Version)
}

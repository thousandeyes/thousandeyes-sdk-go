package main

import (
	"github.com/thousandeyes/go-thousandeyes/v2/command/tectl/cmd"
)

var (
	Version string
)

func main() {
	cmd.Execute(Version)
}

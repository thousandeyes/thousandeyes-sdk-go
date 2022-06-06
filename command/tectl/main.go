package main

import (
	"github.com/thousandeyes/thousandeyes-sdk-go/v2/command/tectl/cmd"
)

var (
	Version string
)

func main() {
	cmd.Execute(Version)
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Id     string
	GetCmd = &cobra.Command{
		Use:   "get",
		Short: "used to get details about resources",
	}
	GetAgentsCmd = &cobra.Command{
		Use:   "agents",
		Short: "get agents",
		Run: func(cmd *cobra.Command, args []string) {
			err := GetAgentsExecute()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	GetTestsCmd = &cobra.Command{
		Use:   "tests",
		Short: "get tests",
		Run: func(cmd *cobra.Command, args []string) {
			err := GetTestsExecute()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	GetCmd.AddCommand(GetAgentsCmd)
	GetCmd.AddCommand(GetTestsCmd)
	GetCmd.PersistentFlags().StringVar(&Id, "id", "", "tectl get [command] -i 1000")
}

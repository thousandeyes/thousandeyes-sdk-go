package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/thousandeyes/thousandeyes-sdk-go/v2"
	"os"
)

var (
	VERSION    string
	ClientOpts = thousandeyes.ClientOptions{
		AuthToken: os.Getenv("TE_TOKEN"),
		AccountID: os.Getenv("TE_AID"),
	}
	RootCmd = &cobra.Command{
		Use:   "tectl",
		Short: "tectl is a cli tool for managing thousandeyes",
		Long:  "tectl is a cli tool for managing thousandeyes and querying the existing configuration",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("tectl version: %s", VERSION)
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of tectl",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version: %s", VERSION)
		},
	}
)

func Execute(version string) error {
	VERSION = version
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(GetCmd)
}

func TableOuput() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	return table
}

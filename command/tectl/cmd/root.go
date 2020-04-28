package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

var (
	Output  string
	Version string
	RootCmd = &cobra.Command{
		Use:   "tectl",
		Short: "tectl is a cli tool for managing thousandeyes",
		Long:  "tectl is a cli tool for managing thousandeyes and querying the existing configuration",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("tectl version: %s", Version)
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of tectl",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version: %s", Version)
		},
	}
	cfgFile string
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "config file (default is $HOME/.tectl.yaml)")
	RootCmd.PersistentFlags().StringVar(&Output, "output", "o", "text")
}

type Display interface {
	Render()
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

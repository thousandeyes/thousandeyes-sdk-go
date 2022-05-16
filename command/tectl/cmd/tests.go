package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/thousandeyes/go-thousandeyes"
)

var TestsCmd = &cobra.Command{
	Use:   "tests",
	Short: "allows for viewing test details",
	Long:  `This sub-command displays test details`,
	Run: func(cmd *cobra.Command, args []string) {
		err := GetTestsExecute()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func GetTestsExecute() error {
	client := thousandeyes.NewClient(&ClientOpts)
	var table *tablewriter.Table
	if GetCmd.Flags().Changed("id") {
		id, err := GetCmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		table, err = getTest(client, id)
		if err != nil {
			return err
		}
	} else {
		var err error
		table, err = getTests(client)
		if err != nil {
			return err
		}
	}
	table.Render()
	return nil
}

func getTest(client *thousandeyes.Client, id string) (*tablewriter.Table, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	test, err := client.GetTest(intId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	table := TableOuput()
	table.SetHeader([]string{"Test Name", "TestID", "Type", "Enabled"})
	fields := []string{test.TestName, strconv.Itoa(test.TestID), test.Type, strconv.Itoa(test.Enabled)}
	table.Append(fields)
	return table, nil
}

func getTests(client *thousandeyes.Client) (*tablewriter.Table, error) {
	tests, err := client.GetTests()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	table := TableOuput()
	table.SetHeader([]string{"Test Name", "TestID", "Type", "Enabled"})
	for _, v := range *tests {
		fields := []string{v.TestName, strconv.Itoa(v.TestID), v.Type, strconv.Itoa(v.Enabled)}
		table.Append(fields)
	}
	return table, nil
}

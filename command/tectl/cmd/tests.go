package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/william20111/go-thousandeyes"
	"os"
	"strconv"
)

var TestsCmd = &cobra.Command{
	Use:   "tests",
	Short: "allows for viewing test details",
	Long:  `This sub-command displays test details`,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := testExecute()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		out.Render()
	},
}

func init() {
	RootCmd.AddCommand(TestsCmd)
}

func testExecute() (Display, error) {
	client := thousandeyes.NewClient(os.Getenv("TE_TOKEN"))
	tests, err := client.GetTests()
	if err != nil {
		return nil, err
	}
	table := TableOuput()
	table.SetHeader([]string{"Test Name", "TestID", "Type", "Enabled"})
	for _, v := range *tests {
		fields := []string{v.TestName, strconv.Itoa(v.TestID), v.Type, strconv.Itoa(v.Enabled)}
		table.Append(fields)
	}
	return table, nil
}

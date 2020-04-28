package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/william20111/go-thousandeyes"
	"os"
	"strconv"
)

var AgentCmd = &cobra.Command{
	Use:   "agents",
	Short: "allows for viewing agent details",
	Long:  `This sub-command displays agent details`,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := agentExecute()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		out.Render()
	},
}

func init() {
	RootCmd.AddCommand(AgentCmd)
}

func agentExecute() (Display, error) {
	client := thousandeyes.NewClient(os.Getenv("TE_TOKEN"))
	agents, err := client.GetAgents()
	if err != nil {
		return nil, err
	}
	table := TableOuput()
	table.SetHeader([]string{"Agent Name", "AgentID", "Enabled"})
	for _, v := range *agents {
		fields := []string{v.AgentName, strconv.Itoa(v.AgentID), strconv.Itoa(v.Enabled)}
		table.Append(fields)
	}
	return table, nil
}

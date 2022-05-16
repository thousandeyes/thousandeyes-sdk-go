package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/thousandeyes/go-thousandeyes"
)

func GetAgentsExecute() error {
	client := thousandeyes.NewClient(&ClientOpts)
	var table *tablewriter.Table
	if GetCmd.Flags().Changed("id") {
		id, err := GetCmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		table, err = getAgent(client, id)
		if err != nil {
			return err
		}
	} else {
		var err error
		table, err = getAgents(client)
		if err != nil {
			return err
		}
	}
	table.Render()
	return nil
}

func getAgents(client *thousandeyes.Client) (*tablewriter.Table, error) {
	agents, err := client.GetAgents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	table := TableOuput()
	table.SetHeader([]string{"Agent Name", "AgentID", "Enabled", "Location", "IpAddresses"})
	for _, v := range *agents {
		fields := []string{v.AgentName, strconv.Itoa(v.AgentID), strconv.Itoa(v.Enabled), v.Location, strings.Join(v.IPAddresses, ",")}
		table.Append(fields)
	}
	return table, nil
}

func getAgent(client *thousandeyes.Client, id string) (*tablewriter.Table, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	agent, err := client.GetAgent(intId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	table := TableOuput()
	table.SetHeader([]string{"Agent Name", "AgentID", "Enabled", "Location", "IpAddresses"})
	fields := []string{agent.AgentName, strconv.Itoa(agent.AgentID), strconv.Itoa(agent.Enabled), agent.Location, strings.Join(agent.IPAddresses, ",")}
	table.Append(fields)
	return table, nil
}

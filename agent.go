package thousandeyes

import "fmt"

type Agents []Agent

type Agent struct {
	AgentId               int                 `json:"agentId,omitempty"`
	AgentName             string              `json:"agentName,omitempty"`
	AgentType             string              `json:"agentType,omitempty"`
	CountryId             string              `json:"countryId,omitempty"`
	ClusterMembers        []ClusterMember     `json:"clusterMembers,omitempty"`
	IpAddresses           []string            `json:"ipAddresses,omitempty"`
	Groups                GroupLabels         `json:"groups,omitempty"`
	Location              string              `json:"location,omitempty"`
	ErrorDetails          []AgentErrorDetails `json:"errorDetails,omitempty"`
	Hostname              string              `json:"hostname,omitempty"`
	Prefix                string              `json:"prefix,omitempty"`
	Enabled               int                 `json:"enabled,omitempty"`
	Network               string              `json:"network,omitempty"`
	CreatedDate           string              `json:"createdDate,omitempty"`
	LastSeen              string              `json:"lastSeen,omitempty"`
	AgentState            string              `json:"agentType,omitempty"`
	VerifySslCertificates bool                `json:"agentType,omitempty"`
	KeepBrowserCache      bool                `json:"agentType,omitempty"`
	Utilization           int                 `json:"agentType,omitempty"`
	Ipv6Policy            string              `json:"agentType,omitempty"`
	TargetForTests        string              `json:"agentType,omitempty"`
}

type ClusterMember struct {
	MemberId          int      `json:"memberId,omitempty"`
	Name              string   `json:"name,omitempty"`
	IpAddresses       []string `json:"IpAddresses,omitempty"`
	PublicIpAddresses []string `json:"PublicIpAddresses,omitempty"`
	Prefix            string   `json:"Prefix,omitempty"`
	Network           string   `json:"network,omitempty"`
	LastSeen          string   `json:"lastSeen,omitempty"`
	AgentState        string   `json:"agentState,omitempty"`
	Utilization       int      `json:"utilization,omitempty"`
	TargetForTests    string   `json:"targetForTests,omitempty"`
}

type AgentErrorDetails struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *Client) GetAgents() (*Agents, error) {
	resp, err := c.get("/agents")
	if err != nil {
		return &Agents{}, err
	}
	var target map[string]Agents
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	agents := target["agents"]
	return &agents, nil
}

package thousandeyes

type Monitor struct {
	MonitorId   int    `json:"monitorId,omitempty"`
	IpAddress   string `json:"ipAddress,omitempty"`
	CountryId   string `json:"countryId,omitempty"`
	MonitorName string `json:"monitorName,omitempty"`
	Network     string `json:"network,omitempty"`
	MonitorType string `json:"monitorType,omitempty"`
}

package thousandeyes

// Monitors - List of monitor
type Monitors []Monitor

// Monitor - A monitor
type Monitor struct {
	MonitorID   *int64  `json:"monitorId,omitempty"`
	IPAddress   *string `json:"ipAddress,omitempty"`
	CountryID   *string `json:"countryId,omitempty"`
	MonitorName *string `json:"monitorName,omitempty"`
	Network     *string `json:"network,omitempty"`
	MonitorType *string `json:"monitorType,omitempty"`
}

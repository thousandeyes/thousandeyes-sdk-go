package thousandeyes

type WebTransaction struct {
	Agents                []Agent
	AuthType              string
	BandwidthMeasurements int
	ContentRegex          string
	Credentials           []int
	CustomHeaders         []map[string]string
	DesiredStatusCode     string
	HttpTargetTime        int
	HttpTimeLimit         int
	HttpVersion           int
	IncludeHeaders        int
	Interval              int
	MtuMeasurements       int
	NetworkMeasurements   int
	NumPathTraces         int
	Password              string
	ProbeMode             string
	Protocol              string
	SslVersionId          int
	Subinterval           int
	TargetTime            int
	TimeLimit             int
	TransactionScript     string
	Url                   string
	UseNtlm               int
	UserAgent             string
	Username              string
	VerifyCertificate     int
}

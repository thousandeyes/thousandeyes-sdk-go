package thousandeyes

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetNetMetrics(t *testing.T){
	var authToken = os.Getenv("TE_TOKEN")
	var client = NewClient(&ClientOptions{AuthToken: authToken})
	res, err := client.GetNetMetrics(1869241, "5m")
	assert.Nil(t, err)
	assert.Equal(t, 1869241, res.Net.Test.TestID)
	assert.Equal(t, "https://app.thousandeyes.com/view/tests?__a=208901&testId=1869241&roundId=1639582200&agentId=14410", res.Net.Metrics[0].PermaLink)
}
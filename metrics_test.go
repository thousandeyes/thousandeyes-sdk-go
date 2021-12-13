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
}
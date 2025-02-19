# AgentToServerEndpointInstantScheduledTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgentToServerScheduledInstantTest**](AgentToServerEndpointInstantScheduledTestsAPI.md#CreateAgentToServerScheduledInstantTest) | **Post** /endpoint/tests/scheduled-tests/agent-to-server/instant | Run agent to server instant scheduled test



## CreateAgentToServerScheduledInstantTest

> EndpointAgentToServerTest CreateAgentToServerScheduledInstantTest().EndpointAgentToServerInstantTest(endpointAgentToServerInstantTest).Aid(aid).Execute()

Run agent to server instant scheduled test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointinstanttests"
)

func main() {
	endpointAgentToServerInstantTest := *endpointinstanttests.NewEndpointAgentToServerInstantTest("Test name", "www.example.com") // EndpointAgentToServerInstantTest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointinstanttests.AgentToServerEndpointInstantScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAgentToServerScheduledInstantTest().EndpointAgentToServerInstantTest(endpointAgentToServerInstantTest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointInstantScheduledTestsAPI.CreateAgentToServerScheduledInstantTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgentToServerScheduledInstantTest`: EndpointAgentToServerTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AgentToServerEndpointInstantScheduledTestsAPI.CreateAgentToServerScheduledInstantTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAgentToServerScheduledInstantTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointAgentToServerInstantTest** | [**EndpointAgentToServerInstantTest**](EndpointAgentToServerInstantTest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointAgentToServerTest**](EndpointAgentToServerTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


# EndpointAgentsTransferAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransferEndpointAgent**](EndpointAgentsTransferAPI.md#TransferEndpointAgent) | **Post** /endpoint/agents/{agentId}/transfer | Transfer endpoint agent
[**TransferEndpointAgents**](EndpointAgentsTransferAPI.md#TransferEndpointAgents) | **Post** /endpoint/agents/transfer/bulk | Bulk transfer agents



## TransferEndpointAgent

> TransferEndpointAgent(agentId).AgentTransferRequest(agentTransferRequest).Aid(aid).Execute()

Transfer endpoint agent



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointagents"
)

func main() {
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the agent to operate on.
	agentTransferRequest := *endpointagents.NewAgentTransferRequest() // AgentTransferRequest | The request to move an agent between accounts.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsTransferAPIService)(&apiClient.Common)

	r, err := api.TransferEndpointAgent(agentId).AgentTransferRequest(agentTransferRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsTransferAPI.TransferEndpointAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | The identifier of the agent to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiTransferEndpointAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentTransferRequest** | [**AgentTransferRequest**](AgentTransferRequest.md) | The request to move an agent between accounts. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## TransferEndpointAgents

> BulkAgentTransferResponse TransferEndpointAgents().Aid(aid).BulkAgentTransferRequest(bulkAgentTransferRequest).Execute()

Bulk transfer agents



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointagents"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	bulkAgentTransferRequest := *endpointagents.NewBulkAgentTransferRequest() // BulkAgentTransferRequest | A collection of `AgentTransfers`. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsTransferAPIService)(&apiClient.Common)

	resp, r, err := api.TransferEndpointAgents().Aid(aid).BulkAgentTransferRequest(bulkAgentTransferRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsTransferAPI.TransferEndpointAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferEndpointAgents`: BulkAgentTransferResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentsTransferAPI.TransferEndpointAgents`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiTransferEndpointAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **bulkAgentTransferRequest** | [**BulkAgentTransferRequest**](BulkAgentTransferRequest.md) | A collection of &#x60;AgentTransfers&#x60;. | 

### Return type

[**BulkAgentTransferResponse**](BulkAgentTransferResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, text/csv, text/plain
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


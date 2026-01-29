# TestsAssignmentOnAgentsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignTests**](TestsAssignmentOnAgentsAPI.md#AssignTests) | **Post** /agents/{agentId}/tests/assign | Assign tests to an agent
[**OverwriteTests**](TestsAssignmentOnAgentsAPI.md#OverwriteTests) | **Post** /agents/{agentId}/tests/override | Overwrite tests assigned to an agent
[**UnassignTests**](TestsAssignmentOnAgentsAPI.md#UnassignTests) | **Post** /agents/{agentId}/tests/unassign | Unassign tests from an agent



## AssignTests

> AgentDetails AssignTests(agentId).AgentTestsAssignRequest(agentTestsAssignRequest).Aid(aid).Execute()

Assign tests to an agent



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/agents"
)

func main() {
	agentId := "281474976710706" // string | Unique ID for the Enterprise Agent cluster to add new agents to.
	agentTestsAssignRequest := *agents.NewAgentTestsAssignRequest() // AgentTestsAssignRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.TestsAssignmentOnAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.AssignTests(agentId).AgentTestsAssignRequest(agentTestsAssignRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAssignmentOnAgentsAPI.AssignTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignTests`: AgentDetails
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TestsAssignmentOnAgentsAPI.AssignTests`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | Unique ID for the Enterprise Agent cluster to add new agents to. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiAssignTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentTestsAssignRequest** | [**AgentTestsAssignRequest**](AgentTestsAssignRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AgentDetails**](AgentDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## OverwriteTests

> AgentDetails OverwriteTests(agentId).AgentTestsAssignRequest(agentTestsAssignRequest).Aid(aid).Execute()

Overwrite tests assigned to an agent



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/agents"
)

func main() {
	agentId := "281474976710706" // string | Unique ID for the Enterprise Agent cluster to add new agents to.
	agentTestsAssignRequest := *agents.NewAgentTestsAssignRequest() // AgentTestsAssignRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.TestsAssignmentOnAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.OverwriteTests(agentId).AgentTestsAssignRequest(agentTestsAssignRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAssignmentOnAgentsAPI.OverwriteTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverwriteTests`: AgentDetails
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TestsAssignmentOnAgentsAPI.OverwriteTests`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | Unique ID for the Enterprise Agent cluster to add new agents to. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiOverwriteTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentTestsAssignRequest** | [**AgentTestsAssignRequest**](AgentTestsAssignRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AgentDetails**](AgentDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UnassignTests

> AgentDetails UnassignTests(agentId).AgentTestsAssignRequest(agentTestsAssignRequest).Aid(aid).Execute()

Unassign tests from an agent



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/agents"
)

func main() {
	agentId := "281474976710706" // string | Unique ID for the Enterprise Agent cluster to add new agents to.
	agentTestsAssignRequest := *agents.NewAgentTestsAssignRequest() // AgentTestsAssignRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.TestsAssignmentOnAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.UnassignTests(agentId).AgentTestsAssignRequest(agentTestsAssignRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAssignmentOnAgentsAPI.UnassignTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnassignTests`: AgentDetails
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TestsAssignmentOnAgentsAPI.UnassignTests`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | Unique ID for the Enterprise Agent cluster to add new agents to. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUnassignTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentTestsAssignRequest** | [**AgentTestsAssignRequest**](AgentTestsAssignRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AgentDetails**](AgentDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


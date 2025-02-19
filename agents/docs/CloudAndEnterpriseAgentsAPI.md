# CloudAndEnterpriseAgentsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAgent**](CloudAndEnterpriseAgentsAPI.md#DeleteAgent) | **Delete** /agents/{agentId} | Delete Enterprise Agent
[**GetAgent**](CloudAndEnterpriseAgentsAPI.md#GetAgent) | **Get** /agents/{agentId} | Retrieve Cloud and Enterprise Agent
[**GetAgents**](CloudAndEnterpriseAgentsAPI.md#GetAgents) | **Get** /agents | List Cloud and Enterprise Agents
[**UpdateAgent**](CloudAndEnterpriseAgentsAPI.md#UpdateAgent) | **Put** /agents/{agentId} | Update Enterprise Agent



## DeleteAgent

> DeleteAgent(agentId).Aid(aid).Execute()

Delete Enterprise Agent



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
	agentId := "281474976710706" // string | Unique ID for the agent.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.CloudAndEnterpriseAgentsAPIService)(&apiClient.Common)

	r, err := api.DeleteAgent(agentId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAndEnterpriseAgentsAPI.DeleteAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | Unique ID for the agent. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAgent

> AgentDetails GetAgent(agentId).Aid(aid).Expand(expand).Execute()

Retrieve Cloud and Enterprise Agent



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
	agentId := "281474976710706" // string | Unique ID for the agent.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []agents.AgentDetailsExpand{agents.AgentDetailsExpand("cluster-member")} // []AgentDetailsExpand | Optional parameter, off by default. Indicates which agent sub-resource to expand. For example, if you wish to expand the `clusterMembers` sub-resource, pass the `?expand=cluster-member` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.CloudAndEnterpriseAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAgent(agentId).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAndEnterpriseAgentsAPI.GetAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgent`: AgentDetails
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudAndEnterpriseAgentsAPI.GetAgent`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | Unique ID for the agent. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]AgentDetailsExpand**](AgentDetailsExpand.md) | Optional parameter, off by default. Indicates which agent sub-resource to expand. For example, if you wish to expand the &#x60;clusterMembers&#x60; sub-resource, pass the &#x60;?expand&#x3D;cluster-member&#x60; query. | 

### Return type

[**AgentDetails**](AgentDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAgents

> CloudEnterpriseAgents GetAgents().Aid(aid).Expand(expand).AgentTypes(agentTypes).Execute()

List Cloud and Enterprise Agents



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []agents.AgentListExpand{agents.AgentListExpand("cluster-member")} // []AgentListExpand | Optional parameter, off by default. Indicates which agent sub-resource to expand. For example, if you wish to expand the `clusterMembers` sub-resource, pass the `?expand=cluster-member` query. (optional)
	agentTypes := []agents.CloudEnterpriseAgentType{agents.CloudEnterpriseAgentType("cloud")} // []CloudEnterpriseAgentType | Specifies the type of agent to request. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.CloudAndEnterpriseAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAgents().Aid(aid).Expand(expand).AgentTypes(agentTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAndEnterpriseAgentsAPI.GetAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgents`: CloudEnterpriseAgents
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudAndEnterpriseAgentsAPI.GetAgents`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]AgentListExpand**](AgentListExpand.md) | Optional parameter, off by default. Indicates which agent sub-resource to expand. For example, if you wish to expand the &#x60;clusterMembers&#x60; sub-resource, pass the &#x60;?expand&#x3D;cluster-member&#x60; query. | 
 **agentTypes** | [**[]CloudEnterpriseAgentType**](CloudEnterpriseAgentType.md) | Specifies the type of agent to request. | 

### Return type

[**CloudEnterpriseAgents**](CloudEnterpriseAgents.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAgent

> AgentDetails UpdateAgent(agentId).AgentRequest(agentRequest).Aid(aid).Expand(expand).Execute()

Update Enterprise Agent



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
	agentId := "281474976710706" // string | Unique ID for the agent.
	agentRequest := *agents.NewAgentRequest() // AgentRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []agents.AgentDetailsExpand{agents.AgentDetailsExpand("cluster-member")} // []AgentDetailsExpand | Optional parameter, off by default. Indicates which agent sub-resource to expand. For example, if you wish to expand the `clusterMembers` sub-resource, pass the `?expand=cluster-member` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.CloudAndEnterpriseAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAgent(agentId).AgentRequest(agentRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAndEnterpriseAgentsAPI.UpdateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgent`: AgentDetails
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudAndEnterpriseAgentsAPI.UpdateAgent`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | Unique ID for the agent. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentRequest** | [**AgentRequest**](AgentRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]AgentDetailsExpand**](AgentDetailsExpand.md) | Optional parameter, off by default. Indicates which agent sub-resource to expand. For example, if you wish to expand the &#x60;clusterMembers&#x60; sub-resource, pass the &#x60;?expand&#x3D;cluster-member&#x60; query. | 

### Return type

[**AgentDetails**](AgentDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


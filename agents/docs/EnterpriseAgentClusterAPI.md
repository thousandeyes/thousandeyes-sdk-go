# EnterpriseAgentClusterAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAgentToCluster**](EnterpriseAgentClusterAPI.md#AssignAgentToCluster) | **Post** /agents/{agentId}/cluster/assign | Add member to Enterprise Agent cluster
[**UnassignAgentFromCluster**](EnterpriseAgentClusterAPI.md#UnassignAgentFromCluster) | **Post** /agents/{agentId}/cluster/unassign | Remove member from Enterprise Agent cluster



## AssignAgentToCluster

> AgentDetails AssignAgentToCluster(agentId).AgentClusterAssignRequest(agentClusterAssignRequest).Aid(aid).Expand(expand).Execute()

Add member to Enterprise Agent cluster



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
	agentClusterAssignRequest := *agents.NewAgentClusterAssignRequest() // AgentClusterAssignRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []agents.AgentDetailsExpand{agents.AgentDetailsExpand("cluster-member")} // []AgentDetailsExpand | Optional parameter, off by default. Indicates which agent sub-resource to expand. For example, if you wish to expand the `clusterMembers` sub-resource, pass the `?expand=cluster-member` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.EnterpriseAgentClusterAPIService)(&apiClient.Common)

	resp, r, err := api.AssignAgentToCluster(agentId).AgentClusterAssignRequest(agentClusterAssignRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAgentClusterAPI.AssignAgentToCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignAgentToCluster`: AgentDetails
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EnterpriseAgentClusterAPI.AssignAgentToCluster`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | Unique ID for the Enterprise Agent cluster to add new agents to. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiAssignAgentToClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentClusterAssignRequest** | [**AgentClusterAssignRequest**](AgentClusterAssignRequest.md) |  | 
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


## UnassignAgentFromCluster

> CloudEnterpriseAgents UnassignAgentFromCluster(agentId).AgentClusterUnassignRequest(agentClusterUnassignRequest).Aid(aid).Expand(expand).Execute()

Remove member from Enterprise Agent cluster



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
	agentId := "281474976710706" // string | Unique ID for the Enterprise Agent cluster to remove agents from.
	agentClusterUnassignRequest := *agents.NewAgentClusterUnassignRequest() // AgentClusterUnassignRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []agents.AgentDetailsExpand{agents.AgentDetailsExpand("cluster-member")} // []AgentDetailsExpand | Optional parameter, off by default. Indicates which agent sub-resource to expand. For example, if you wish to expand the `clusterMembers` sub-resource, pass the `?expand=cluster-member` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.EnterpriseAgentClusterAPIService)(&apiClient.Common)

	resp, r, err := api.UnassignAgentFromCluster(agentId).AgentClusterUnassignRequest(agentClusterUnassignRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAgentClusterAPI.UnassignAgentFromCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnassignAgentFromCluster`: CloudEnterpriseAgents
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EnterpriseAgentClusterAPI.UnassignAgentFromCluster`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | Unique ID for the Enterprise Agent cluster to remove agents from. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUnassignAgentFromClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentClusterUnassignRequest** | [**AgentClusterUnassignRequest**](AgentClusterUnassignRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]AgentDetailsExpand**](AgentDetailsExpand.md) | Optional parameter, off by default. Indicates which agent sub-resource to expand. For example, if you wish to expand the &#x60;clusterMembers&#x60; sub-resource, pass the &#x60;?expand&#x3D;cluster-member&#x60; query. | 

### Return type

[**CloudEnterpriseAgents**](CloudEnterpriseAgents.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


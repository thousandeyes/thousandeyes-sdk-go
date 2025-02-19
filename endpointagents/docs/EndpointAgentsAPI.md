# EndpointAgentsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEndpointAgent**](EndpointAgentsAPI.md#DeleteEndpointAgent) | **Delete** /endpoint/agents/{agentId} | Delete endpoint agent
[**DisableEndpointAgent**](EndpointAgentsAPI.md#DisableEndpointAgent) | **Post** /endpoint/agents/{agentId}/disable | Disable endpoint agent
[**EnableEndpointAgent**](EndpointAgentsAPI.md#EnableEndpointAgent) | **Post** /endpoint/agents/{agentId}/enable | Enable endpoint agent
[**FilterEndpointAgents**](EndpointAgentsAPI.md#FilterEndpointAgents) | **Post** /endpoint/agents/filter | Filter endpoint agents
[**GetEndpointAgent**](EndpointAgentsAPI.md#GetEndpointAgent) | **Get** /endpoint/agents/{agentId} | Retrieve endpoint agent
[**GetEndpointAgents**](EndpointAgentsAPI.md#GetEndpointAgents) | **Get** /endpoint/agents | List endpoint agents
[**GetEndpointAgentsConnectionString**](EndpointAgentsAPI.md#GetEndpointAgentsConnectionString) | **Get** /endpoint/agents/connection-string | Get agent connection string
[**UpdateEndpointAgent**](EndpointAgentsAPI.md#UpdateEndpointAgent) | **Patch** /endpoint/agents/{agentId} | Update endpoint agent



## DeleteEndpointAgent

> DeleteEndpointAgent(agentId).Aid(aid).Expand(expand).Execute()

Delete endpoint agent



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []endpointagents.ExpandEndpointAgentOptions{endpointagents.ExpandEndpointAgentOptions("clients")} // []ExpandEndpointAgentOptions | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \"clients\" resource, include the query parameter `?expand=clients`.  For multiple expansions, you have two options:    * Separate the values with commas. For example, `?expandAgent=clients,tasks`. * Specify the parameter multiple times. For example, `?expandAgent=clients&expandAgent=tasks`.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsAPIService)(&apiClient.Common)

	r, err := api.DeleteEndpointAgent(agentId).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsAPI.DeleteEndpointAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | The identifier of the agent to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteEndpointAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandEndpointAgentOptions**](ExpandEndpointAgentOptions.md) | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \&quot;clients\&quot; resource, include the query parameter &#x60;?expand&#x3D;clients&#x60;.  For multiple expansions, you have two options:    * Separate the values with commas. For example, &#x60;?expandAgent&#x3D;clients,tasks&#x60;. * Specify the parameter multiple times. For example, &#x60;?expandAgent&#x3D;clients&amp;expandAgent&#x3D;tasks&#x60;.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DisableEndpointAgent

> EndpointAgent DisableEndpointAgent(agentId).Aid(aid).Execute()

Disable endpoint agent



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.DisableEndpointAgent(agentId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsAPI.DisableEndpointAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableEndpointAgent`: EndpointAgent
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentsAPI.DisableEndpointAgent`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | The identifier of the agent to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDisableEndpointAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointAgent**](EndpointAgent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## EnableEndpointAgent

> EndpointAgent EnableEndpointAgent(agentId).Aid(aid).Execute()

Enable endpoint agent



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.EnableEndpointAgent(agentId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsAPI.EnableEndpointAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableEndpointAgent`: EndpointAgent
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentsAPI.EnableEndpointAgent`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | The identifier of the agent to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiEnableEndpointAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointAgent**](EndpointAgent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## FilterEndpointAgents

> FilterEndpointAgentsResponse FilterEndpointAgents().AgentSearchRequest(agentSearchRequest).Max(max).Cursor(cursor).Aid(aid).Expand(expand).IncludeDeleted(includeDeleted).Execute()

Filter endpoint agents



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
	agentSearchRequest := *endpointagents.NewAgentSearchRequest() // AgentSearchRequest | The filter options for advanced search filtering for agents.
	max := int32(5) // int32 | (Optional) Maximum number of objects to return. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []endpointagents.ExpandEndpointAgentOptions{endpointagents.ExpandEndpointAgentOptions("clients")} // []ExpandEndpointAgentOptions | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \"clients\" resource, include the query parameter `?expand=clients`.  For multiple expansions, you have two options:    * Separate the values with commas. For example, `?expandAgent=clients,tasks`. * Specify the parameter multiple times. For example, `?expandAgent=clients&expandAgent=tasks`.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  (optional)
	includeDeleted := false // bool | When requesting entities, set to `true` if you want to see deleted entities. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.FilterEndpointAgents().AgentSearchRequest(agentSearchRequest).Max(max).Cursor(cursor).Aid(aid).Expand(expand).IncludeDeleted(includeDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsAPI.FilterEndpointAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterEndpointAgents`: FilterEndpointAgentsResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentsAPI.FilterEndpointAgents`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiFilterEndpointAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentSearchRequest** | [**AgentSearchRequest**](AgentSearchRequest.md) | The filter options for advanced search filtering for agents. | 
 **max** | **int32** | (Optional) Maximum number of objects to return. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandEndpointAgentOptions**](ExpandEndpointAgentOptions.md) | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \&quot;clients\&quot; resource, include the query parameter &#x60;?expand&#x3D;clients&#x60;.  For multiple expansions, you have two options:    * Separate the values with commas. For example, &#x60;?expandAgent&#x3D;clients,tasks&#x60;. * Specify the parameter multiple times. For example, &#x60;?expandAgent&#x3D;clients&amp;expandAgent&#x3D;tasks&#x60;.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  | 
 **includeDeleted** | **bool** | When requesting entities, set to &#x60;true&#x60; if you want to see deleted entities. | 

### Return type

[**FilterEndpointAgentsResponse**](FilterEndpointAgentsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetEndpointAgent

> EndpointAgent GetEndpointAgent(agentId).Aid(aid).Expand(expand).IncludeDeleted(includeDeleted).Execute()

Retrieve endpoint agent



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []endpointagents.ExpandEndpointAgentOptions{endpointagents.ExpandEndpointAgentOptions("clients")} // []ExpandEndpointAgentOptions | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \"clients\" resource, include the query parameter `?expand=clients`.  For multiple expansions, you have two options:    * Separate the values with commas. For example, `?expandAgent=clients,tasks`. * Specify the parameter multiple times. For example, `?expandAgent=clients&expandAgent=tasks`.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  (optional)
	includeDeleted := false // bool | When requesting entities, set to `true` if you want to see deleted entities. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.GetEndpointAgent(agentId).Aid(aid).Expand(expand).IncludeDeleted(includeDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsAPI.GetEndpointAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointAgent`: EndpointAgent
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentsAPI.GetEndpointAgent`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | The identifier of the agent to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetEndpointAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandEndpointAgentOptions**](ExpandEndpointAgentOptions.md) | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \&quot;clients\&quot; resource, include the query parameter &#x60;?expand&#x3D;clients&#x60;.  For multiple expansions, you have two options:    * Separate the values with commas. For example, &#x60;?expandAgent&#x3D;clients,tasks&#x60;. * Specify the parameter multiple times. For example, &#x60;?expandAgent&#x3D;clients&amp;expandAgent&#x3D;tasks&#x60;.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  | 
 **includeDeleted** | **bool** | When requesting entities, set to &#x60;true&#x60; if you want to see deleted entities. | 

### Return type

[**EndpointAgent**](EndpointAgent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetEndpointAgents

> ListEndpointAgentsResponse GetEndpointAgents().Max(max).Cursor(cursor).Aid(aid).Expand(expand).IncludeDeleted(includeDeleted).UseAllPermittedAids(useAllPermittedAids).AgentName(agentName).ComputerName(computerName).Execute()

List endpoint agents



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
	max := int32(5) // int32 | (Optional) Maximum number of objects to return. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []endpointagents.ExpandEndpointAgentOptions{endpointagents.ExpandEndpointAgentOptions("clients")} // []ExpandEndpointAgentOptions | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \"clients\" resource, include the query parameter `?expand=clients`.  For multiple expansions, you have two options:    * Separate the values with commas. For example, `?expandAgent=clients,tasks`. * Specify the parameter multiple times. For example, `?expandAgent=clients&expandAgent=tasks`.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  (optional)
	includeDeleted := false // bool | When requesting entities, set to `true` if you want to see deleted entities. (optional)
	useAllPermittedAids := false // bool | Set to `true` to load data from all accounts the user has access to. (optional) (default to false)
	agentName := "agentName_example" // string | Returns only agents with the specified name.  This is an exact match only.  (optional)
	computerName := "computerName_example" // string | Returns only agents with the specified computer name. This is an exact match only.  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.GetEndpointAgents().Max(max).Cursor(cursor).Aid(aid).Expand(expand).IncludeDeleted(includeDeleted).UseAllPermittedAids(useAllPermittedAids).AgentName(agentName).ComputerName(computerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsAPI.GetEndpointAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointAgents`: ListEndpointAgentsResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentsAPI.GetEndpointAgents`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetEndpointAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int32** | (Optional) Maximum number of objects to return. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandEndpointAgentOptions**](ExpandEndpointAgentOptions.md) | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \&quot;clients\&quot; resource, include the query parameter &#x60;?expand&#x3D;clients&#x60;.  For multiple expansions, you have two options:    * Separate the values with commas. For example, &#x60;?expandAgent&#x3D;clients,tasks&#x60;. * Specify the parameter multiple times. For example, &#x60;?expandAgent&#x3D;clients&amp;expandAgent&#x3D;tasks&#x60;.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  | 
 **includeDeleted** | **bool** | When requesting entities, set to &#x60;true&#x60; if you want to see deleted entities. | 
 **useAllPermittedAids** | **bool** | Set to &#x60;true&#x60; to load data from all accounts the user has access to. | [default to false]
 **agentName** | **string** | Returns only agents with the specified name.  This is an exact match only.  | 
 **computerName** | **string** | Returns only agents with the specified computer name. This is an exact match only.  | 

### Return type

[**ListEndpointAgentsResponse**](ListEndpointAgentsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetEndpointAgentsConnectionString

> ConnectionString GetEndpointAgentsConnectionString().Aid(aid).Execute()

Get agent connection string

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

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.GetEndpointAgentsConnectionString().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsAPI.GetEndpointAgentsConnectionString``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointAgentsConnectionString`: ConnectionString
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentsAPI.GetEndpointAgentsConnectionString`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetEndpointAgentsConnectionStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ConnectionString**](ConnectionString.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateEndpointAgent

> EndpointAgent UpdateEndpointAgent(agentId).Aid(aid).Expand(expand).EndpointAgentUpdate(endpointAgentUpdate).Execute()

Update endpoint agent



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []endpointagents.ExpandEndpointAgentOptions{endpointagents.ExpandEndpointAgentOptions("clients")} // []ExpandEndpointAgentOptions | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \"clients\" resource, include the query parameter `?expand=clients`.  For multiple expansions, you have two options:    * Separate the values with commas. For example, `?expandAgent=clients,tasks`. * Specify the parameter multiple times. For example, `?expandAgent=clients&expandAgent=tasks`.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  (optional)
	endpointAgentUpdate := *endpointagents.NewEndpointAgentUpdate() // EndpointAgentUpdate | Fields to modify on the agent (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateEndpointAgent(agentId).Aid(aid).Expand(expand).EndpointAgentUpdate(endpointAgentUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentsAPI.UpdateEndpointAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointAgent`: EndpointAgent
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentsAPI.UpdateEndpointAgent`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | The identifier of the agent to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateEndpointAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandEndpointAgentOptions**](ExpandEndpointAgentOptions.md) | This optional parameter allows you to control the expansion of test resources associated with the agent. By default, no expansion occurs when this query parameter is omitted. To expand the \&quot;clients\&quot; resource, include the query parameter &#x60;?expand&#x3D;clients&#x60;.  For multiple expansions, you have two options:    * Separate the values with commas. For example, &#x60;?expandAgent&#x3D;clients,tasks&#x60;. * Specify the parameter multiple times. For example, &#x60;?expandAgent&#x3D;clients&amp;expandAgent&#x3D;tasks&#x60;.  This parameter offers flexibility for users to customize the expansion of specific resources related to the agent.  | 
 **endpointAgentUpdate** | [**EndpointAgentUpdate**](EndpointAgentUpdate.md) | Fields to modify on the agent | 

### Return type

[**EndpointAgent**](EndpointAgent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


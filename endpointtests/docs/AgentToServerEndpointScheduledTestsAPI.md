# AgentToServerEndpointScheduledTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgentToServerEndpointScheduledTest**](AgentToServerEndpointScheduledTestsAPI.md#CreateAgentToServerEndpointScheduledTest) | **Post** /endpoint/tests/scheduled-tests/agent-to-server | Creates agent to server endpoint scheduled test
[**DeleteAgentToServerEndpointScheduledTest**](AgentToServerEndpointScheduledTestsAPI.md#DeleteAgentToServerEndpointScheduledTest) | **Delete** /endpoint/tests/scheduled-tests/agent-to-server/{testId} | Delete agent to server scheduled test
[**GetAgentToServerEndpointScheduledTest**](AgentToServerEndpointScheduledTestsAPI.md#GetAgentToServerEndpointScheduledTest) | **Get** /endpoint/tests/scheduled-tests/agent-to-server/{testId} | Retrieve agent to server endpoint scheduled test
[**GetAgentToServerEndpointScheduledTests**](AgentToServerEndpointScheduledTestsAPI.md#GetAgentToServerEndpointScheduledTests) | **Get** /endpoint/tests/scheduled-tests/agent-to-server | List agent to server endpoint scheduled tests
[**UpdateAgentToServerEndpointScheduledTest**](AgentToServerEndpointScheduledTestsAPI.md#UpdateAgentToServerEndpointScheduledTest) | **Patch** /endpoint/tests/scheduled-tests/agent-to-server/{testId} | Update agent to server endpoint scheduled test



## CreateAgentToServerEndpointScheduledTest

> EndpointAgentToServerTest CreateAgentToServerEndpointScheduledTest().EndpointAgentToServerTestRequest(endpointAgentToServerTestRequest).Aid(aid).Execute()

Creates agent to server endpoint scheduled test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtests"
)

func main() {
	endpointAgentToServerTestRequest := *endpointtests.NewEndpointAgentToServerTestRequest("Test name", "www.example.com") // EndpointAgentToServerTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtests.AgentToServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAgentToServerEndpointScheduledTest().EndpointAgentToServerTestRequest(endpointAgentToServerTestRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointScheduledTestsAPI.CreateAgentToServerEndpointScheduledTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgentToServerEndpointScheduledTest`: EndpointAgentToServerTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AgentToServerEndpointScheduledTestsAPI.CreateAgentToServerEndpointScheduledTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAgentToServerEndpointScheduledTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointAgentToServerTestRequest** | [**EndpointAgentToServerTestRequest**](EndpointAgentToServerTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointAgentToServerTest**](EndpointAgentToServerTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteAgentToServerEndpointScheduledTest

> DeleteAgentToServerEndpointScheduledTest(testId).Aid(aid).Execute()

Delete agent to server scheduled test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtests"
)

func main() {
	testId := "584739201" // string | Unique ID of endpoint test.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtests.AgentToServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	r, err := api.DeleteAgentToServerEndpointScheduledTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointScheduledTestsAPI.DeleteAgentToServerEndpointScheduledTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Unique ID of endpoint test. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteAgentToServerEndpointScheduledTestRequest struct via the builder pattern


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


## GetAgentToServerEndpointScheduledTest

> EndpointAgentToServerTest GetAgentToServerEndpointScheduledTest(testId).Aid(aid).Execute()

Retrieve agent to server endpoint scheduled test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtests"
)

func main() {
	testId := "584739201" // string | Unique ID of endpoint test.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtests.AgentToServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAgentToServerEndpointScheduledTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointScheduledTestsAPI.GetAgentToServerEndpointScheduledTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentToServerEndpointScheduledTest`: EndpointAgentToServerTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AgentToServerEndpointScheduledTestsAPI.GetAgentToServerEndpointScheduledTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Unique ID of endpoint test. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAgentToServerEndpointScheduledTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointAgentToServerTest**](EndpointAgentToServerTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAgentToServerEndpointScheduledTests

> EndpointAgentToServerTests GetAgentToServerEndpointScheduledTests().Aid(aid).Execute()

List agent to server endpoint scheduled tests



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtests"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtests.AgentToServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAgentToServerEndpointScheduledTests().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointScheduledTestsAPI.GetAgentToServerEndpointScheduledTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentToServerEndpointScheduledTests`: EndpointAgentToServerTests
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AgentToServerEndpointScheduledTestsAPI.GetAgentToServerEndpointScheduledTests`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAgentToServerEndpointScheduledTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointAgentToServerTests**](EndpointAgentToServerTests.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAgentToServerEndpointScheduledTest

> EndpointAgentToServerTest UpdateAgentToServerEndpointScheduledTest(testId).EndpointNetworkTestUpdate(endpointNetworkTestUpdate).Aid(aid).Execute()

Update agent to server endpoint scheduled test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtests"
)

func main() {
	testId := "584739201" // string | Unique ID of endpoint test.
	endpointNetworkTestUpdate := *endpointtests.NewEndpointNetworkTestUpdate() // EndpointNetworkTestUpdate | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtests.AgentToServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAgentToServerEndpointScheduledTest(testId).EndpointNetworkTestUpdate(endpointNetworkTestUpdate).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointScheduledTestsAPI.UpdateAgentToServerEndpointScheduledTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgentToServerEndpointScheduledTest`: EndpointAgentToServerTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AgentToServerEndpointScheduledTestsAPI.UpdateAgentToServerEndpointScheduledTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Unique ID of endpoint test. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAgentToServerEndpointScheduledTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointNetworkTestUpdate** | [**EndpointNetworkTestUpdate**](EndpointNetworkTestUpdate.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointAgentToServerTest**](EndpointAgentToServerTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


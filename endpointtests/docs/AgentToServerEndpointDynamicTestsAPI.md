# AgentToServerEndpointDynamicTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgentToServerEndpointDynamicTest**](AgentToServerEndpointDynamicTestsAPI.md#CreateAgentToServerEndpointDynamicTest) | **Post** /endpoint/tests/dynamic-tests/agent-to-server | Create endpoint dynamic test
[**DeleteAgentToServerEndpointDynamicTest**](AgentToServerEndpointDynamicTestsAPI.md#DeleteAgentToServerEndpointDynamicTest) | **Delete** /endpoint/tests/dynamic-tests/agent-to-server/{testId} | Delete agent to server dynamic test
[**GetAgentToServerEndpointDynamicTest**](AgentToServerEndpointDynamicTestsAPI.md#GetAgentToServerEndpointDynamicTest) | **Get** /endpoint/tests/dynamic-tests/agent-to-server/{testId} | Retrieve endpoint dynamic test
[**GetAgentToServerEndpointDynamicTests**](AgentToServerEndpointDynamicTestsAPI.md#GetAgentToServerEndpointDynamicTests) | **Get** /endpoint/tests/dynamic-tests/agent-to-server | List endpoint dynamic tests
[**UpdateAgentToServerEndpointDynamicTest**](AgentToServerEndpointDynamicTestsAPI.md#UpdateAgentToServerEndpointDynamicTest) | **Patch** /endpoint/tests/dynamic-tests/agent-to-server/{testId} | Update agent to server dynamic test



## CreateAgentToServerEndpointDynamicTest

> DynamicTest CreateAgentToServerEndpointDynamicTest().DynamicTestRequest(dynamicTestRequest).Aid(aid).Execute()

Create endpoint dynamic test



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
	dynamicTestRequest := *endpointtests.NewDynamicTestRequest("webex", "Test name") // DynamicTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtests.AgentToServerEndpointDynamicTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAgentToServerEndpointDynamicTest().DynamicTestRequest(dynamicTestRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointDynamicTestsAPI.CreateAgentToServerEndpointDynamicTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgentToServerEndpointDynamicTest`: DynamicTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AgentToServerEndpointDynamicTestsAPI.CreateAgentToServerEndpointDynamicTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAgentToServerEndpointDynamicTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dynamicTestRequest** | [**DynamicTestRequest**](DynamicTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**DynamicTest**](DynamicTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteAgentToServerEndpointDynamicTest

> DeleteAgentToServerEndpointDynamicTest(testId).Aid(aid).Execute()

Delete agent to server dynamic test



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

	api := (*endpointtests.AgentToServerEndpointDynamicTestsAPIService)(&apiClient.Common)

	r, err := api.DeleteAgentToServerEndpointDynamicTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointDynamicTestsAPI.DeleteAgentToServerEndpointDynamicTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Unique ID of endpoint test. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteAgentToServerEndpointDynamicTestRequest struct via the builder pattern


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


## GetAgentToServerEndpointDynamicTest

> DynamicTest GetAgentToServerEndpointDynamicTest(testId).Aid(aid).Execute()

Retrieve endpoint dynamic test



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

	api := (*endpointtests.AgentToServerEndpointDynamicTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAgentToServerEndpointDynamicTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointDynamicTestsAPI.GetAgentToServerEndpointDynamicTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentToServerEndpointDynamicTest`: DynamicTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AgentToServerEndpointDynamicTestsAPI.GetAgentToServerEndpointDynamicTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Unique ID of endpoint test. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAgentToServerEndpointDynamicTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**DynamicTest**](DynamicTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAgentToServerEndpointDynamicTests

> DynamicTests GetAgentToServerEndpointDynamicTests().Aid(aid).Execute()

List endpoint dynamic tests



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

	api := (*endpointtests.AgentToServerEndpointDynamicTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAgentToServerEndpointDynamicTests().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointDynamicTestsAPI.GetAgentToServerEndpointDynamicTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentToServerEndpointDynamicTests`: DynamicTests
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AgentToServerEndpointDynamicTestsAPI.GetAgentToServerEndpointDynamicTests`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAgentToServerEndpointDynamicTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**DynamicTests**](DynamicTests.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAgentToServerEndpointDynamicTest

> DynamicTest UpdateAgentToServerEndpointDynamicTest(testId).EndpointDynamicTestUpdate(endpointDynamicTestUpdate).Aid(aid).Execute()

Update agent to server dynamic test



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
	endpointDynamicTestUpdate := *endpointtests.NewEndpointDynamicTestUpdate() // EndpointDynamicTestUpdate | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtests.AgentToServerEndpointDynamicTestsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAgentToServerEndpointDynamicTest(testId).EndpointDynamicTestUpdate(endpointDynamicTestUpdate).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentToServerEndpointDynamicTestsAPI.UpdateAgentToServerEndpointDynamicTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgentToServerEndpointDynamicTest`: DynamicTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AgentToServerEndpointDynamicTestsAPI.UpdateAgentToServerEndpointDynamicTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Unique ID of endpoint test. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAgentToServerEndpointDynamicTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointDynamicTestUpdate** | [**EndpointDynamicTestUpdate**](EndpointDynamicTestUpdate.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**DynamicTest**](DynamicTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


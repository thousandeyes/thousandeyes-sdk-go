# APITestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiTest**](APITestsAPI.md#CreateApiTest) | **Post** /tests/api | Create API test
[**DeleteApiTest**](APITestsAPI.md#DeleteApiTest) | **Delete** /tests/api/{testId} | Delete API test
[**GetApiTest**](APITestsAPI.md#GetApiTest) | **Get** /tests/api/{testId} | Get API test
[**GetApiTests**](APITestsAPI.md#GetApiTests) | **Get** /tests/api | List API tests
[**UpdateApiTest**](APITestsAPI.md#UpdateApiTest) | **Put** /tests/api/{testId} | Update API test



## CreateApiTest

> ApiTestResponse CreateApiTest().ApiTestRequest(apiTestRequest).Aid(aid).Expand(expand).Execute()

Create API test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	apiTestRequest := *tests.NewApiTestRequest(tests.TestInterval(60), []tests.ApiRequest{*tests.NewApiRequest("Step 1", "https://api.thousandeyes.com/v7/status")}, "www.thousandeyes.com", []tests.TestAgentRequest{*tests.NewTestAgentRequest("AgentId_example")}) // ApiTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandTestOptions{tests.ExpandTestOptions("agent")} // []ExpandTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the `agents` sub-resource, they need to pass the `?expand=agent` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.APITestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateApiTest().ApiTestRequest(apiTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APITestsAPI.CreateApiTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiTest`: ApiTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `APITestsAPI.CreateApiTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateApiTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiTestRequest** | [**ApiTestRequest**](ApiTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**ApiTestResponse**](ApiTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteApiTest

> DeleteApiTest(testId).Aid(aid).Execute()

Delete API test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	testId := "202701" // string | Test ID
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.APITestsAPIService)(&apiClient.Common)

	r, err := api.DeleteApiTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APITestsAPI.DeleteApiTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteApiTestRequest struct via the builder pattern


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


## GetApiTest

> ApiTestResponse GetApiTest(testId).Aid(aid).Expand(expand).Execute()

Get API test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	testId := "202701" // string | Test ID
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandTestOptions{tests.ExpandTestOptions("agent")} // []ExpandTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the `agents` sub-resource, they need to pass the `?expand=agent` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.APITestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetApiTest(testId).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APITestsAPI.GetApiTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiTest`: ApiTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `APITestsAPI.GetApiTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetApiTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**ApiTestResponse**](ApiTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetApiTests

> ApiTests GetApiTests().Aid(aid).Execute()

List API tests



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.APITestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetApiTests().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APITestsAPI.GetApiTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiTests`: ApiTests
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `APITestsAPI.GetApiTests`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetApiTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiTests**](ApiTests.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateApiTest

> ApiTestResponse UpdateApiTest(testId).ApiTestRequest(apiTestRequest).Aid(aid).Expand(expand).Execute()

Update API test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	testId := "202701" // string | Test ID
	apiTestRequest := *tests.NewApiTestRequest(tests.TestInterval(60), []tests.ApiRequest{*tests.NewApiRequest("Step 1", "https://api.thousandeyes.com/v7/status")}, "www.thousandeyes.com", []tests.TestAgentRequest{*tests.NewTestAgentRequest("AgentId_example")}) // ApiTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandTestOptions{tests.ExpandTestOptions("agent")} // []ExpandTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the `agents` sub-resource, they need to pass the `?expand=agent` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.APITestsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateApiTest(testId).ApiTestRequest(apiTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APITestsAPI.UpdateApiTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiTest`: ApiTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `APITestsAPI.UpdateApiTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateApiTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiTestRequest** | [**ApiTestRequest**](ApiTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**ApiTestResponse**](ApiTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


# HTTPServerEndpointScheduledTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHttpServerEndpointScheduledTest**](HTTPServerEndpointScheduledTestsAPI.md#CreateHttpServerEndpointScheduledTest) | **Post** /endpoint/tests/scheduled-tests/http-server | Create HTTP server endpoint scheduled test
[**DeleteHttpServerEndpointScheduledTest**](HTTPServerEndpointScheduledTestsAPI.md#DeleteHttpServerEndpointScheduledTest) | **Delete** /endpoint/tests/scheduled-tests/http-server/{testId} | Delete HTTP server scheduled test
[**GetHttpServerEndpointScheduledTest**](HTTPServerEndpointScheduledTestsAPI.md#GetHttpServerEndpointScheduledTest) | **Get** /endpoint/tests/scheduled-tests/http-server/{testId} | Retrieves HTTP server endpoint scheduled test
[**GetHttpServerEndpointScheduledTests**](HTTPServerEndpointScheduledTestsAPI.md#GetHttpServerEndpointScheduledTests) | **Get** /endpoint/tests/scheduled-tests/http-server | List HTTP server endpoint scheduled tests
[**UpdateHttpServerEndpointScheduledTest**](HTTPServerEndpointScheduledTestsAPI.md#UpdateHttpServerEndpointScheduledTest) | **Patch** /endpoint/tests/scheduled-tests/http-server/{testId} | Update HTTP server endpoint scheduled test



## CreateHttpServerEndpointScheduledTest

> EndpointHttpServerTest CreateHttpServerEndpointScheduledTest().EndpointHttpServerTestRequest(endpointHttpServerTestRequest).Aid(aid).Execute()

Create HTTP server endpoint scheduled test



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
	endpointHttpServerTestRequest := *endpointtests.NewEndpointHttpServerTestRequest("Test name", "https://example.com:443") // EndpointHttpServerTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtests.HTTPServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateHttpServerEndpointScheduledTest().EndpointHttpServerTestRequest(endpointHttpServerTestRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPServerEndpointScheduledTestsAPI.CreateHttpServerEndpointScheduledTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHttpServerEndpointScheduledTest`: EndpointHttpServerTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `HTTPServerEndpointScheduledTestsAPI.CreateHttpServerEndpointScheduledTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateHttpServerEndpointScheduledTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointHttpServerTestRequest** | [**EndpointHttpServerTestRequest**](EndpointHttpServerTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointHttpServerTest**](EndpointHttpServerTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteHttpServerEndpointScheduledTest

> DeleteHttpServerEndpointScheduledTest(testId).Aid(aid).Execute()

Delete HTTP server scheduled test



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

	api := (*endpointtests.HTTPServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	r, err := api.DeleteHttpServerEndpointScheduledTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPServerEndpointScheduledTestsAPI.DeleteHttpServerEndpointScheduledTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Unique ID of endpoint test. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteHttpServerEndpointScheduledTestRequest struct via the builder pattern


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


## GetHttpServerEndpointScheduledTest

> EndpointHttpServerTest GetHttpServerEndpointScheduledTest(testId).Aid(aid).Execute()

Retrieves HTTP server endpoint scheduled test



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

	api := (*endpointtests.HTTPServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetHttpServerEndpointScheduledTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPServerEndpointScheduledTestsAPI.GetHttpServerEndpointScheduledTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHttpServerEndpointScheduledTest`: EndpointHttpServerTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `HTTPServerEndpointScheduledTestsAPI.GetHttpServerEndpointScheduledTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Unique ID of endpoint test. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetHttpServerEndpointScheduledTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointHttpServerTest**](EndpointHttpServerTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetHttpServerEndpointScheduledTests

> EndpointHttpServerTests GetHttpServerEndpointScheduledTests().Aid(aid).Execute()

List HTTP server endpoint scheduled tests



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

	api := (*endpointtests.HTTPServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetHttpServerEndpointScheduledTests().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPServerEndpointScheduledTestsAPI.GetHttpServerEndpointScheduledTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHttpServerEndpointScheduledTests`: EndpointHttpServerTests
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `HTTPServerEndpointScheduledTestsAPI.GetHttpServerEndpointScheduledTests`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetHttpServerEndpointScheduledTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointHttpServerTests**](EndpointHttpServerTests.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateHttpServerEndpointScheduledTest

> EndpointHttpServerTest UpdateHttpServerEndpointScheduledTest(testId).EndpointHttpTestUpdate(endpointHttpTestUpdate).Aid(aid).Execute()

Update HTTP server endpoint scheduled test



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
	endpointHttpTestUpdate := *endpointtests.NewEndpointHttpTestUpdate() // EndpointHttpTestUpdate | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtests.HTTPServerEndpointScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateHttpServerEndpointScheduledTest(testId).EndpointHttpTestUpdate(endpointHttpTestUpdate).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPServerEndpointScheduledTestsAPI.UpdateHttpServerEndpointScheduledTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHttpServerEndpointScheduledTest`: EndpointHttpServerTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `HTTPServerEndpointScheduledTestsAPI.UpdateHttpServerEndpointScheduledTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Unique ID of endpoint test. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateHttpServerEndpointScheduledTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointHttpTestUpdate** | [**EndpointHttpTestUpdate**](EndpointHttpTestUpdate.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointHttpServerTest**](EndpointHttpServerTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


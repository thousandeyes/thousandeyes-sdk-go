# BGPTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBgpTest**](BGPTestsAPI.md#CreateBgpTest) | **Post** /tests/bgp | Create BGP test
[**DeleteBgpTest**](BGPTestsAPI.md#DeleteBgpTest) | **Delete** /tests/bgp/{testId} | Delete BGP test
[**GetBgpTest**](BGPTestsAPI.md#GetBgpTest) | **Get** /tests/bgp/{testId} | Get BGP test
[**GetBgpTests**](BGPTestsAPI.md#GetBgpTests) | **Get** /tests/bgp | List BGP tests
[**UpdateBgpTest**](BGPTestsAPI.md#UpdateBgpTest) | **Put** /tests/bgp/{testId} | Update BGP test



## CreateBgpTest

> BgpTestResponse CreateBgpTest().BgpTestRequest(bgpTestRequest).Aid(aid).Expand(expand).Execute()

Create BGP test



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
	bgpTestRequest := *tests.NewBgpTestRequest("Prefix_example") // BgpTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandBgpTestOptions{tests.ExpandBgpTestOptions("alert-rule")} // []ExpandBgpTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion takes place if the query parameter is not present. To expand the `monitors` sub-resource, pass the `?expand=monitor` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.BGPTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateBgpTest().BgpTestRequest(bgpTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTestsAPI.CreateBgpTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBgpTest`: BgpTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `BGPTestsAPI.CreateBgpTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateBgpTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bgpTestRequest** | [**BgpTestRequest**](BgpTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandBgpTestOptions**](ExpandBgpTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion takes place if the query parameter is not present. To expand the &#x60;monitors&#x60; sub-resource, pass the &#x60;?expand&#x3D;monitor&#x60; query. | 

### Return type

[**BgpTestResponse**](BgpTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteBgpTest

> DeleteBgpTest(testId).Aid(aid).Execute()

Delete BGP test



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

	api := (*tests.BGPTestsAPIService)(&apiClient.Common)

	r, err := api.DeleteBgpTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTestsAPI.DeleteBgpTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteBgpTestRequest struct via the builder pattern


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


## GetBgpTest

> BgpTestResponse GetBgpTest(testId).Aid(aid).Expand(expand).Execute()

Get BGP test



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
	expand := []tests.ExpandBgpTestOptions{tests.ExpandBgpTestOptions("alert-rule")} // []ExpandBgpTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion takes place if the query parameter is not present. To expand the `monitors` sub-resource, pass the `?expand=monitor` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.BGPTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetBgpTest(testId).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTestsAPI.GetBgpTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBgpTest`: BgpTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `BGPTestsAPI.GetBgpTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetBgpTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandBgpTestOptions**](ExpandBgpTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion takes place if the query parameter is not present. To expand the &#x60;monitors&#x60; sub-resource, pass the &#x60;?expand&#x3D;monitor&#x60; query. | 

### Return type

[**BgpTestResponse**](BgpTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetBgpTests

> BgpTests GetBgpTests().Aid(aid).Execute()

List BGP tests



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

	api := (*tests.BGPTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetBgpTests().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTestsAPI.GetBgpTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBgpTests`: BgpTests
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `BGPTestsAPI.GetBgpTests`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetBgpTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**BgpTests**](BgpTests.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateBgpTest

> BgpTestResponse UpdateBgpTest(testId).UpdateBgpTestRequest(updateBgpTestRequest).Aid(aid).Expand(expand).Execute()

Update BGP test



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
	updateBgpTestRequest := *tests.NewUpdateBgpTestRequest() // UpdateBgpTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandBgpTestOptions{tests.ExpandBgpTestOptions("alert-rule")} // []ExpandBgpTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion takes place if the query parameter is not present. To expand the `monitors` sub-resource, pass the `?expand=monitor` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.BGPTestsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateBgpTest(testId).UpdateBgpTestRequest(updateBgpTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPTestsAPI.UpdateBgpTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBgpTest`: BgpTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `BGPTestsAPI.UpdateBgpTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateBgpTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBgpTestRequest** | [**UpdateBgpTestRequest**](UpdateBgpTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandBgpTestOptions**](ExpandBgpTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion takes place if the query parameter is not present. To expand the &#x60;monitors&#x60; sub-resource, pass the &#x60;?expand&#x3D;monitor&#x60; query. | 

### Return type

[**BgpTestResponse**](BgpTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


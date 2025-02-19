# DNSSECTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnsSecTest**](DNSSECTestsAPI.md#CreateDnsSecTest) | **Post** /tests/dnssec | Create DNSSEC test
[**DeleteDnsSecTest**](DNSSECTestsAPI.md#DeleteDnsSecTest) | **Delete** /tests/dnssec/{testId} | Delete DNSSEC test
[**GetDnsSecTest**](DNSSECTestsAPI.md#GetDnsSecTest) | **Get** /tests/dnssec/{testId} | Get DNSSEC test
[**GetDnsSecTests**](DNSSECTestsAPI.md#GetDnsSecTests) | **Get** /tests/dnssec | List DNSSEC tests
[**UpdateDnsSecTest**](DNSSECTestsAPI.md#UpdateDnsSecTest) | **Put** /tests/dnssec/{testId} | Update DNSSEC test



## CreateDnsSecTest

> DnsSecTestResponse CreateDnsSecTest().DnsSecTestRequest(dnsSecTestRequest).Aid(aid).Expand(expand).Execute()

Create DNSSEC test



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
	dnsSecTestRequest := *tests.NewDnsSecTestRequest(tests.TestInterval(60), "www.thousandeyes.com", []tests.TestAgentRequest{*tests.NewTestAgentRequest("AgentId_example")}) // DnsSecTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandTestOptions{tests.ExpandTestOptions("agent")} // []ExpandTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the `agents` sub-resource, they need to pass the `?expand=agent` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.DNSSECTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateDnsSecTest().DnsSecTestRequest(dnsSecTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSSECTestsAPI.CreateDnsSecTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDnsSecTest`: DnsSecTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSSECTestsAPI.CreateDnsSecTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateDnsSecTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsSecTestRequest** | [**DnsSecTestRequest**](DnsSecTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**DnsSecTestResponse**](DnsSecTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteDnsSecTest

> DeleteDnsSecTest(testId).Aid(aid).Execute()

Delete DNSSEC test



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

	api := (*tests.DNSSECTestsAPIService)(&apiClient.Common)

	r, err := api.DeleteDnsSecTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSSECTestsAPI.DeleteDnsSecTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteDnsSecTestRequest struct via the builder pattern


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


## GetDnsSecTest

> DnsSecTestResponse GetDnsSecTest(testId).Aid(aid).Expand(expand).Execute()

Get DNSSEC test



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

	api := (*tests.DNSSECTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDnsSecTest(testId).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSSECTestsAPI.GetDnsSecTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsSecTest`: DnsSecTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSSECTestsAPI.GetDnsSecTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDnsSecTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**DnsSecTestResponse**](DnsSecTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDnsSecTests

> DnsSecTests GetDnsSecTests().Aid(aid).Execute()

List DNSSEC tests



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

	api := (*tests.DNSSECTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDnsSecTests().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSSECTestsAPI.GetDnsSecTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsSecTests`: DnsSecTests
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSSECTestsAPI.GetDnsSecTests`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetDnsSecTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**DnsSecTests**](DnsSecTests.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateDnsSecTest

> DnsSecTestResponse UpdateDnsSecTest(testId).DnsSecTestRequest(dnsSecTestRequest).Aid(aid).Expand(expand).Execute()

Update DNSSEC test



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
	dnsSecTestRequest := *tests.NewDnsSecTestRequest(tests.TestInterval(60), "www.thousandeyes.com", []tests.TestAgentRequest{*tests.NewTestAgentRequest("AgentId_example")}) // DnsSecTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandTestOptions{tests.ExpandTestOptions("agent")} // []ExpandTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the `agents` sub-resource, they need to pass the `?expand=agent` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.DNSSECTestsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateDnsSecTest(testId).DnsSecTestRequest(dnsSecTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSSECTestsAPI.UpdateDnsSecTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDnsSecTest`: DnsSecTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSSECTestsAPI.UpdateDnsSecTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateDnsSecTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsSecTestRequest** | [**DnsSecTestRequest**](DnsSecTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**DnsSecTestResponse**](DnsSecTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


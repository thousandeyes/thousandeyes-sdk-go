# DNSTraceTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnsTraceTest**](DNSTraceTestsAPI.md#CreateDnsTraceTest) | **Post** /tests/dns-trace | Create DNS Trace test
[**DeleteDnsTraceTest**](DNSTraceTestsAPI.md#DeleteDnsTraceTest) | **Delete** /tests/dns-trace/{testId} | Delete DNS Trace test
[**GetDnsTraceTest**](DNSTraceTestsAPI.md#GetDnsTraceTest) | **Get** /tests/dns-trace/{testId} | Get DNS Trace test
[**GetDnsTraceTests**](DNSTraceTestsAPI.md#GetDnsTraceTests) | **Get** /tests/dns-trace | List DNS Trace tests
[**UpdateDnsTraceTest**](DNSTraceTestsAPI.md#UpdateDnsTraceTest) | **Put** /tests/dns-trace/{testId} | Update DNS Trace test



## CreateDnsTraceTest

> DnsTraceTestResponse CreateDnsTraceTest().DnsTraceTestRequest(dnsTraceTestRequest).Aid(aid).Expand(expand).Execute()

Create DNS Trace test



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
	dnsTraceTestRequest := *tests.NewDnsTraceTestRequest(tests.TestInterval(60), "www.thousandeyes.com", []tests.TestAgentRequest{*tests.NewTestAgentRequest("AgentId_example")}) // DnsTraceTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandTestOptions{tests.ExpandTestOptions("agent")} // []ExpandTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the `agents` sub-resource, they need to pass the `?expand=agent` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.DNSTraceTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateDnsTraceTest().DnsTraceTestRequest(dnsTraceTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTraceTestsAPI.CreateDnsTraceTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDnsTraceTest`: DnsTraceTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSTraceTestsAPI.CreateDnsTraceTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateDnsTraceTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsTraceTestRequest** | [**DnsTraceTestRequest**](DnsTraceTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**DnsTraceTestResponse**](DnsTraceTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteDnsTraceTest

> DeleteDnsTraceTest(testId).Aid(aid).Execute()

Delete DNS Trace test



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

	api := (*tests.DNSTraceTestsAPIService)(&apiClient.Common)

	r, err := api.DeleteDnsTraceTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTraceTestsAPI.DeleteDnsTraceTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteDnsTraceTestRequest struct via the builder pattern


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


## GetDnsTraceTest

> DnsTraceTestResponse GetDnsTraceTest(testId).Aid(aid).Expand(expand).Execute()

Get DNS Trace test



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

	api := (*tests.DNSTraceTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDnsTraceTest(testId).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTraceTestsAPI.GetDnsTraceTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsTraceTest`: DnsTraceTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSTraceTestsAPI.GetDnsTraceTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDnsTraceTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**DnsTraceTestResponse**](DnsTraceTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDnsTraceTests

> DnsTraceTests GetDnsTraceTests().Aid(aid).Execute()

List DNS Trace tests



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

	api := (*tests.DNSTraceTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDnsTraceTests().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTraceTestsAPI.GetDnsTraceTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsTraceTests`: DnsTraceTests
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSTraceTestsAPI.GetDnsTraceTests`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetDnsTraceTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**DnsTraceTests**](DnsTraceTests.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateDnsTraceTest

> DnsTraceTestResponse UpdateDnsTraceTest(testId).DnsTraceTestRequest(dnsTraceTestRequest).Aid(aid).Expand(expand).Execute()

Update DNS Trace test



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
	dnsTraceTestRequest := *tests.NewDnsTraceTestRequest(tests.TestInterval(60), "www.thousandeyes.com", []tests.TestAgentRequest{*tests.NewTestAgentRequest("AgentId_example")}) // DnsTraceTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandTestOptions{tests.ExpandTestOptions("agent")} // []ExpandTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the `agents` sub-resource, they need to pass the `?expand=agent` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.DNSTraceTestsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateDnsTraceTest(testId).DnsTraceTestRequest(dnsTraceTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTraceTestsAPI.UpdateDnsTraceTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDnsTraceTest`: DnsTraceTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSTraceTestsAPI.UpdateDnsTraceTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateDnsTraceTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsTraceTestRequest** | [**DnsTraceTestRequest**](DnsTraceTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**DnsTraceTestResponse**](DnsTraceTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


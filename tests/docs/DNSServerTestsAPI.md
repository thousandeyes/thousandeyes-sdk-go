# DNSServerTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnsServerTest**](DNSServerTestsAPI.md#CreateDnsServerTest) | **Post** /tests/dns-server | Create DNS Server test
[**DeleteDnsServerTest**](DNSServerTestsAPI.md#DeleteDnsServerTest) | **Delete** /tests/dns-server/{testId} | Delete DNS Server test
[**GetDnsServerTest**](DNSServerTestsAPI.md#GetDnsServerTest) | **Get** /tests/dns-server/{testId} | Get DNS Server test
[**GetDnsServerTests**](DNSServerTestsAPI.md#GetDnsServerTests) | **Get** /tests/dns-server | List DNS Server tests
[**UpdateDnsServerTest**](DNSServerTestsAPI.md#UpdateDnsServerTest) | **Put** /tests/dns-server/{testId} | Update DNS Server test



## CreateDnsServerTest

> DnsServerTestResponse CreateDnsServerTest().DnsServerTestRequest(dnsServerTestRequest).Aid(aid).Expand(expand).Execute()

Create DNS Server test



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
	dnsServerTestRequest := *tests.NewDnsServerTestRequest(tests.TestInterval(60), []string{"DnsServers_example"}, "www.thousandeyes.com", []tests.TestAgentRequest{*tests.NewTestAgentRequest("AgentId_example")}) // DnsServerTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandTestOptions{tests.ExpandTestOptions("agent")} // []ExpandTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the `agents` sub-resource, they need to pass the `?expand=agent` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.DNSServerTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateDnsServerTest().DnsServerTestRequest(dnsServerTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSServerTestsAPI.CreateDnsServerTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDnsServerTest`: DnsServerTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSServerTestsAPI.CreateDnsServerTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateDnsServerTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsServerTestRequest** | [**DnsServerTestRequest**](DnsServerTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**DnsServerTestResponse**](DnsServerTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteDnsServerTest

> DeleteDnsServerTest(testId).Aid(aid).Execute()

Delete DNS Server test



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

	api := (*tests.DNSServerTestsAPIService)(&apiClient.Common)

	r, err := api.DeleteDnsServerTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSServerTestsAPI.DeleteDnsServerTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteDnsServerTestRequest struct via the builder pattern


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


## GetDnsServerTest

> DnsServerTestResponse GetDnsServerTest(testId).Aid(aid).Expand(expand).Execute()

Get DNS Server test



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

	api := (*tests.DNSServerTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDnsServerTest(testId).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSServerTestsAPI.GetDnsServerTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsServerTest`: DnsServerTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSServerTestsAPI.GetDnsServerTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDnsServerTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**DnsServerTestResponse**](DnsServerTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDnsServerTests

> DnsServerTests GetDnsServerTests().Aid(aid).Execute()

List DNS Server tests



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

	api := (*tests.DNSServerTestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDnsServerTests().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSServerTestsAPI.GetDnsServerTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsServerTests`: DnsServerTests
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSServerTestsAPI.GetDnsServerTests`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetDnsServerTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**DnsServerTests**](DnsServerTests.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateDnsServerTest

> DnsServerTestResponse UpdateDnsServerTest(testId).DnsServerTestRequest(dnsServerTestRequest).Aid(aid).Expand(expand).Execute()

Update DNS Server test



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
	dnsServerTestRequest := *tests.NewDnsServerTestRequest(tests.TestInterval(60), []string{"DnsServers_example"}, "www.thousandeyes.com", []tests.TestAgentRequest{*tests.NewTestAgentRequest("AgentId_example")}) // DnsServerTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tests.ExpandTestOptions{tests.ExpandTestOptions("agent")} // []ExpandTestOptions | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the `agents` sub-resource, they need to pass the `?expand=agent` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.DNSServerTestsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateDnsServerTest(testId).DnsServerTestRequest(dnsServerTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSServerTestsAPI.UpdateDnsServerTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDnsServerTest`: DnsServerTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DNSServerTestsAPI.UpdateDnsServerTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateDnsServerTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsServerTestRequest** | [**DnsServerTestRequest**](DnsServerTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTestOptions**](ExpandTestOptions.md) | Optional parameter on whether or not to expand the test sub-resources. By default no expansion is going to take place if the query parameter is not present. If the user wishes to expand the &#x60;agents&#x60; sub-resource, they need to pass the &#x60;?expand&#x3D;agent&#x60; query. | 

### Return type

[**DnsServerTestResponse**](DnsServerTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


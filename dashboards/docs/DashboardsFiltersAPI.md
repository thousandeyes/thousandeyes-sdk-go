# DashboardsFiltersAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboardFilter**](DashboardsFiltersAPI.md#CreateDashboardFilter) | **Post** /dashboards/filters | Create dashboard filter
[**DeleteDashboardFilter**](DashboardsFiltersAPI.md#DeleteDashboardFilter) | **Delete** /dashboards/filters/{id} | Delete dashboard filter
[**GetDashboardFilter**](DashboardsFiltersAPI.md#GetDashboardFilter) | **Get** /dashboards/filters/{id} | Get dashboard filter
[**GetDashboardsFilters**](DashboardsFiltersAPI.md#GetDashboardsFilters) | **Get** /dashboards/filters | List dashboard filters
[**UpdateDashboardFilter**](DashboardsFiltersAPI.md#UpdateDashboardFilter) | **Put** /dashboards/filters/{id} | Update dashboard filter



## CreateDashboardFilter

> ApiContextFilterResponse CreateDashboardFilter().ApiContextFilterRequest(apiContextFilterRequest).Aid(aid).Execute()

Create dashboard filter



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/dashboards"
)

func main() {
	apiContextFilterRequest := *dashboards.NewApiContextFilterRequest([]dashboards.ApiDataSourceFilters{*dashboards.NewApiDataSourceFilters("VIRTUAL_AGENT", []dashboards.ApiDataSourceFilter{*dashboards.NewApiDataSourceFilter("TEST_LABEL", []string{"Values_example"}, []string{"MetricIds_example"})})}, "cea-filter") // ApiContextFilterRequest | Dashboard filter object to be created and saved
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsFiltersAPIService)(&apiClient.Common)

	resp, r, err := api.CreateDashboardFilter().ApiContextFilterRequest(apiContextFilterRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsFiltersAPI.CreateDashboardFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboardFilter`: ApiContextFilterResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardsFiltersAPI.CreateDashboardFilter`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateDashboardFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiContextFilterRequest** | [**ApiContextFilterRequest**](ApiContextFilterRequest.md) | Dashboard filter object to be created and saved | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiContextFilterResponse**](ApiContextFilterResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteDashboardFilter

> DeleteDashboardFilter(id).Aid(aid).Execute()

Delete dashboard filter



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/dashboards"
)

func main() {
	id := "65bc18e8f2073a4a469cd958" // string | Unique dashboard filter ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsFiltersAPIService)(&apiClient.Common)

	r, err := api.DeleteDashboardFilter(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsFiltersAPI.DeleteDashboardFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Unique dashboard filter ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteDashboardFilterRequest struct via the builder pattern


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


## GetDashboardFilter

> ApiContextFilterResponse GetDashboardFilter(id).Aid(aid).Execute()

Get dashboard filter



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/dashboards"
)

func main() {
	id := "65bc18e8f2073a4a469cd958" // string | Unique dashboard filter ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsFiltersAPIService)(&apiClient.Common)

	resp, r, err := api.GetDashboardFilter(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsFiltersAPI.GetDashboardFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardFilter`: ApiContextFilterResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardsFiltersAPI.GetDashboardFilter`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Unique dashboard filter ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDashboardFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiContextFilterResponse**](ApiContextFilterResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDashboardsFilters

> ApiContextFiltersResponse GetDashboardsFilters().SearchPattern(searchPattern).Aid(aid).Execute()

List dashboard filters



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/dashboards"
)

func main() {
	searchPattern := "cea-filter" // string | Optional search pattern parameter to filter list of dashboard filters by either name or description values. (optional)
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsFiltersAPIService)(&apiClient.Common)

	resp, r, err := api.GetDashboardsFilters().SearchPattern(searchPattern).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsFiltersAPI.GetDashboardsFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardsFilters`: ApiContextFiltersResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardsFiltersAPI.GetDashboardsFilters`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetDashboardsFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchPattern** | **string** | Optional search pattern parameter to filter list of dashboard filters by either name or description values. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiContextFiltersResponse**](ApiContextFiltersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateDashboardFilter

> ApiContextFilterResponse UpdateDashboardFilter(id).ApiContextFilterRequest(apiContextFilterRequest).Aid(aid).Execute()

Update dashboard filter



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/dashboards"
)

func main() {
	id := "65bc18e8f2073a4a469cd958" // string | Unique dashboard filter ID.
	apiContextFilterRequest := *dashboards.NewApiContextFilterRequest([]dashboards.ApiDataSourceFilters{*dashboards.NewApiDataSourceFilters("VIRTUAL_AGENT", []dashboards.ApiDataSourceFilter{*dashboards.NewApiDataSourceFilter("TEST_LABEL", []string{"Values_example"}, []string{"MetricIds_example"})})}, "cea-filter") // ApiContextFilterRequest | Updated dashboard filter context object
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsFiltersAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateDashboardFilter(id).ApiContextFilterRequest(apiContextFilterRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsFiltersAPI.UpdateDashboardFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboardFilter`: ApiContextFilterResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardsFiltersAPI.UpdateDashboardFilter`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Unique dashboard filter ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateDashboardFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiContextFilterRequest** | [**ApiContextFilterRequest**](ApiContextFilterRequest.md) | Updated dashboard filter context object | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiContextFilterResponse**](ApiContextFilterResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


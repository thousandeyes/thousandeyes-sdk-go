# RealUserEndpointTestResultsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterRealUserTestsNetworkResults**](RealUserEndpointTestResultsAPI.md#FilterRealUserTestsNetworkResults) | **Post** /endpoint/test-results/real-user-tests/networks/filter | List endpoint real user tests
[**FilterRealUserTestsResults**](RealUserEndpointTestResultsAPI.md#FilterRealUserTestsResults) | **Post** /endpoint/test-results/real-user-tests/filter | List endpoint real user tests
[**FilterRealUserTestsVisitedPagesResults**](RealUserEndpointTestResultsAPI.md#FilterRealUserTestsVisitedPagesResults) | **Post** /endpoint/test-results/real-user-tests/pages/filter | List endpoint real user tests visited pages
[**GetRealUserTestPageResults**](RealUserEndpointTestResultsAPI.md#GetRealUserTestPageResults) | **Get** /endpoint/test-results/real-user-tests/{id}/pages/{pageId} | Retrieve endpoint real user test page
[**GetRealUserTestResults**](RealUserEndpointTestResultsAPI.md#GetRealUserTestResults) | **Get** /endpoint/test-results/real-user-tests/{id} | Retrieve endpoint real user test



## FilterRealUserTestsNetworkResults

> RealUserEndpointTestNetworkResults FilterRealUserTestsNetworkResults().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).RealUserEndpointTestResultsRequest(realUserEndpointTestResultsRequest).Execute()

List endpoint real user tests



### Example

```go
package main

import (
	"fmt"
	"os"
    "time"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtestresults"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	window := "12h" // string | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: `s` for seconds (default if no type is specified), `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. For a precise date range, use `startDate` and `endDate`. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	realUserEndpointTestResultsRequest := *endpointtestresults.NewRealUserEndpointTestResultsRequest() // RealUserEndpointTestResultsRequest |  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.RealUserEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.FilterRealUserTestsNetworkResults().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).RealUserEndpointTestResultsRequest(realUserEndpointTestResultsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealUserEndpointTestResultsAPI.FilterRealUserTestsNetworkResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterRealUserTestsNetworkResults`: RealUserEndpointTestNetworkResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `RealUserEndpointTestResultsAPI.FilterRealUserTestsNetworkResults`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiFilterRealUserTestsNetworkResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **realUserEndpointTestResultsRequest** | [**RealUserEndpointTestResultsRequest**](RealUserEndpointTestResultsRequest.md) |  | 

### Return type

[**RealUserEndpointTestNetworkResults**](RealUserEndpointTestNetworkResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## FilterRealUserTestsResults

> RealUserEndpointTestResults FilterRealUserTestsResults().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).RealUserEndpointTestResultsRequest(realUserEndpointTestResultsRequest).Execute()

List endpoint real user tests



### Example

```go
package main

import (
	"fmt"
	"os"
    "time"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtestresults"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	window := "12h" // string | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: `s` for seconds (default if no type is specified), `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. For a precise date range, use `startDate` and `endDate`. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	realUserEndpointTestResultsRequest := *endpointtestresults.NewRealUserEndpointTestResultsRequest() // RealUserEndpointTestResultsRequest |  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.RealUserEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.FilterRealUserTestsResults().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).RealUserEndpointTestResultsRequest(realUserEndpointTestResultsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealUserEndpointTestResultsAPI.FilterRealUserTestsResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterRealUserTestsResults`: RealUserEndpointTestResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `RealUserEndpointTestResultsAPI.FilterRealUserTestsResults`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiFilterRealUserTestsResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **realUserEndpointTestResultsRequest** | [**RealUserEndpointTestResultsRequest**](RealUserEndpointTestResultsRequest.md) |  | 

### Return type

[**RealUserEndpointTestResults**](RealUserEndpointTestResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## FilterRealUserTestsVisitedPagesResults

> RealUserEndpointTestPageResults FilterRealUserTestsVisitedPagesResults().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).RealUserEndpointTestResultRequestFilter(realUserEndpointTestResultRequestFilter).Execute()

List endpoint real user tests visited pages



### Example

```go
package main

import (
	"fmt"
	"os"
    "time"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtestresults"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	window := "12h" // string | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: `s` for seconds (default if no type is specified), `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. For a precise date range, use `startDate` and `endDate`. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	realUserEndpointTestResultRequestFilter := *endpointtestresults.NewRealUserEndpointTestResultRequestFilter() // RealUserEndpointTestResultRequestFilter |  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.RealUserEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.FilterRealUserTestsVisitedPagesResults().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).RealUserEndpointTestResultRequestFilter(realUserEndpointTestResultRequestFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealUserEndpointTestResultsAPI.FilterRealUserTestsVisitedPagesResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterRealUserTestsVisitedPagesResults`: RealUserEndpointTestPageResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `RealUserEndpointTestResultsAPI.FilterRealUserTestsVisitedPagesResults`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiFilterRealUserTestsVisitedPagesResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **realUserEndpointTestResultRequestFilter** | [**RealUserEndpointTestResultRequestFilter**](RealUserEndpointTestResultRequestFilter.md) |  | 

### Return type

[**RealUserEndpointTestPageResults**](RealUserEndpointTestPageResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetRealUserTestPageResults

> RealUserEndpointTestPageDetailResult GetRealUserTestPageResults(idpageId).Aid(aid).Execute()

Retrieve endpoint real user test page



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtestresults"
)

func main() {
	id := "07625:1490529480:h3qJQTpl" // string | The real user test id.
	pageId := "281474976710706" // string | Web page ID
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.RealUserEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetRealUserTestPageResults(idpageId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealUserEndpointTestResultsAPI.GetRealUserTestPageResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealUserTestPageResults`: RealUserEndpointTestPageDetailResult
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `RealUserEndpointTestResultsAPI.GetRealUserTestPageResults`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The real user test id. | **pageId** | **string** | Web page ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetRealUserTestPageResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**RealUserEndpointTestPageDetailResult**](RealUserEndpointTestPageDetailResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetRealUserTestResults

> RealUserEndpointTestDetailResults GetRealUserTestResults(id).Aid(aid).Execute()

Retrieve endpoint real user test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtestresults"
)

func main() {
	id := "07625:1490529480:h3qJQTpl" // string | The real user test id.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.RealUserEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetRealUserTestResults(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealUserEndpointTestResultsAPI.GetRealUserTestResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealUserTestResults`: RealUserEndpointTestDetailResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `RealUserEndpointTestResultsAPI.GetRealUserTestResults`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The real user test id. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetRealUserTestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**RealUserEndpointTestDetailResults**](RealUserEndpointTestDetailResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


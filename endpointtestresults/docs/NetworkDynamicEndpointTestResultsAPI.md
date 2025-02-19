# NetworkDynamicEndpointTestResultsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterDynamicTestNetworkResults**](NetworkDynamicEndpointTestResultsAPI.md#FilterDynamicTestNetworkResults) | **Post** /endpoint/test-results/dynamic-tests/{testId}/network/filter | Retrieve network dynamic test results
[**GetDynamicTestPathVisAgentRoundResults**](NetworkDynamicEndpointTestResultsAPI.md#GetDynamicTestPathVisAgentRoundResults) | **Get** /endpoint/test-results/dynamic-tests/{testId}/path-vis/agent/{agentId}/round/{roundId} | Retrieve path visualization network dynamic test results details
[**GetDynamicTestPathVisResults**](NetworkDynamicEndpointTestResultsAPI.md#GetDynamicTestPathVisResults) | **Get** /endpoint/test-results/dynamic-tests/{testId}/path-vis | Retrieve path visualization network dynamic test results



## FilterDynamicTestNetworkResults

> NetworkDynamicEndpointTestResults FilterDynamicTestNetworkResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).DynamicEndpointTestsDataRoundSearch(dynamicEndpointTestsDataRoundSearch).Execute()

Retrieve network dynamic test results



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
	testId := "202701" // string | Test ID
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	window := "12h" // string | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: `s` for seconds (default if no type is specified), `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. For a precise date range, use `startDate` and `endDate`. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	dynamicEndpointTestsDataRoundSearch := *endpointtestresults.NewDynamicEndpointTestsDataRoundSearch() // DynamicEndpointTestsDataRoundSearch | Tests data search filters. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.NetworkDynamicEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.FilterDynamicTestNetworkResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).DynamicEndpointTestsDataRoundSearch(dynamicEndpointTestsDataRoundSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDynamicEndpointTestResultsAPI.FilterDynamicTestNetworkResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterDynamicTestNetworkResults`: NetworkDynamicEndpointTestResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `NetworkDynamicEndpointTestResultsAPI.FilterDynamicTestNetworkResults`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiFilterDynamicTestNetworkResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **dynamicEndpointTestsDataRoundSearch** | [**DynamicEndpointTestsDataRoundSearch**](DynamicEndpointTestsDataRoundSearch.md) | Tests data search filters. | 

### Return type

[**NetworkDynamicEndpointTestResults**](NetworkDynamicEndpointTestResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDynamicTestPathVisAgentRoundResults

> PathVisDetailDynamicEndpointTestResults GetDynamicTestPathVisAgentRoundResults(testIdagentIdroundId).Aid(aid).Execute()

Retrieve path visualization network dynamic test results details



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
	testId := "202701" // string | Test ID
	agentId := "11" // string | Agent ID
	roundId := "1384309800" // string | Round ID
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.NetworkDynamicEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDynamicTestPathVisAgentRoundResults(testIdagentIdroundId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDynamicEndpointTestResultsAPI.GetDynamicTestPathVisAgentRoundResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicTestPathVisAgentRoundResults`: PathVisDetailDynamicEndpointTestResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `NetworkDynamicEndpointTestResultsAPI.GetDynamicTestPathVisAgentRoundResults`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | **agentId** | **string** | Agent ID | **roundId** | **string** | Round ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDynamicTestPathVisAgentRoundResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**PathVisDetailDynamicEndpointTestResults**](PathVisDetailDynamicEndpointTestResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDynamicTestPathVisResults

> PathVisDynamicEndpointTestResults GetDynamicTestPathVisResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Execute()

Retrieve path visualization network dynamic test results



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
	testId := "202701" // string | Test ID
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	window := "12h" // string | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: `s` for seconds (default if no type is specified), `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. For a precise date range, use `startDate` and `endDate`. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.NetworkDynamicEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDynamicTestPathVisResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkDynamicEndpointTestResultsAPI.GetDynamicTestPathVisResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicTestPathVisResults`: PathVisDynamicEndpointTestResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `NetworkDynamicEndpointTestResultsAPI.GetDynamicTestPathVisResults`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDynamicTestPathVisResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 

### Return type

[**PathVisDynamicEndpointTestResults**](PathVisDynamicEndpointTestResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


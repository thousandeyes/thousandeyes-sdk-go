# NetworkTestResultsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTestNetworkResults**](NetworkTestResultsAPI.md#GetTestNetworkResults) | **Get** /test-results/{testId}/network | Get network test results
[**GetTestPathVisAgentRoundResults**](NetworkTestResultsAPI.md#GetTestPathVisAgentRoundResults) | **Get** /test-results/{testId}/path-vis/agent/{agentId}/round/{roundId} | Get path visualization test results by agent and round
[**GetTestPathVisResults**](NetworkTestResultsAPI.md#GetTestPathVisResults) | **Get** /test-results/{testId}/path-vis | Get path visualization network test results



## GetTestNetworkResults

> NetworkTestResults GetTestNetworkResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Direction(direction).Execute()

Get network test results



### Example

```go
package main

import (
	"fmt"
	"os"
    "time"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/testresults"
)

func main() {
	testId := "202701" // string | Test ID
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	window := "12h" // string | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: `s` for seconds (default if no type is specified), `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. For a precise date range, use `startDate` and `endDate`. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	direction := testresults.TestDirection("to-target") // TestDirection | Choose the direction for the metrics you want: [`from-target`, `to-target`, `bidirectional`]. This applies when you're doing bidirectional Agent-to-Agent tests. For bidirectional data, you'll get combined results; otherwise, you'll get data for one direction. If you try to get unidirectional test data with an incorrect direction parameter, it will trigger an error response. (optional) (default to "to-target")

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*testresults.NetworkTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetTestNetworkResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTestResultsAPI.GetTestNetworkResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestNetworkResults`: NetworkTestResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `NetworkTestResultsAPI.GetTestNetworkResults`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetTestNetworkResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **direction** | [**TestDirection**](TestDirection.md) | Choose the direction for the metrics you want: [&#x60;from-target&#x60;, &#x60;to-target&#x60;, &#x60;bidirectional&#x60;]. This applies when you&#39;re doing bidirectional Agent-to-Agent tests. For bidirectional data, you&#39;ll get combined results; otherwise, you&#39;ll get data for one direction. If you try to get unidirectional test data with an incorrect direction parameter, it will trigger an error response. | [default to &quot;to-target&quot;]

### Return type

[**NetworkTestResults**](NetworkTestResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetTestPathVisAgentRoundResults

> PathVisDetailTestResults GetTestPathVisAgentRoundResults(testIdagentIdroundId).Aid(aid).Direction(direction).Execute()

Get path visualization test results by agent and round



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/testresults"
)

func main() {
	testId := "202701" // string | Test ID
	agentId := "11" // string | Agent ID
	roundId := "1384309800" // string | Round ID
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	direction := testresults.PathVisDirection("to-target") // PathVisDirection | Choose the direction for the metrics you want: [`from-target`, `to-target`]. This applies when you're doing bidirectional Agent-to-Agent tests. Omitting the parameter will default the results to both `from-target` and `to-target` values (bidirectional); otherwise, you'll get data for one direction. If you try to get unidirectional test data with an incorrect direction parameter, it will trigger an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*testresults.NetworkTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetTestPathVisAgentRoundResults(testIdagentIdroundId).Aid(aid).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTestResultsAPI.GetTestPathVisAgentRoundResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestPathVisAgentRoundResults`: PathVisDetailTestResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `NetworkTestResultsAPI.GetTestPathVisAgentRoundResults`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | **agentId** | **string** | Agent ID | **roundId** | **string** | Round ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetTestPathVisAgentRoundResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **direction** | [**PathVisDirection**](PathVisDirection.md) | Choose the direction for the metrics you want: [&#x60;from-target&#x60;, &#x60;to-target&#x60;]. This applies when you&#39;re doing bidirectional Agent-to-Agent tests. Omitting the parameter will default the results to both &#x60;from-target&#x60; and &#x60;to-target&#x60; values (bidirectional); otherwise, you&#39;ll get data for one direction. If you try to get unidirectional test data with an incorrect direction parameter, it will trigger an error response. | 

### Return type

[**PathVisDetailTestResults**](PathVisDetailTestResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetTestPathVisResults

> PathVisTestResults GetTestPathVisResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Direction(direction).Execute()

Get path visualization network test results



### Example

```go
package main

import (
	"fmt"
	"os"
    "time"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/testresults"
)

func main() {
	testId := "202701" // string | Test ID
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	window := "12h" // string | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: `s` for seconds (default if no type is specified), `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. For a precise date range, use `startDate` and `endDate`. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	direction := testresults.PathVisDirection("to-target") // PathVisDirection | Choose the direction for the metrics you want: [`from-target`, `to-target`]. This applies when you're doing bidirectional Agent-to-Agent tests. Omitting the parameter will default the results to both `from-target` and `to-target` values (bidirectional); otherwise, you'll get data for one direction. If you try to get unidirectional test data with an incorrect direction parameter, it will trigger an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*testresults.NetworkTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetTestPathVisResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTestResultsAPI.GetTestPathVisResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestPathVisResults`: PathVisTestResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `NetworkTestResultsAPI.GetTestPathVisResults`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetTestPathVisResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **direction** | [**PathVisDirection**](PathVisDirection.md) | Choose the direction for the metrics you want: [&#x60;from-target&#x60;, &#x60;to-target&#x60;]. This applies when you&#39;re doing bidirectional Agent-to-Agent tests. Omitting the parameter will default the results to both &#x60;from-target&#x60; and &#x60;to-target&#x60; values (bidirectional); otherwise, you&#39;ll get data for one direction. If you try to get unidirectional test data with an incorrect direction parameter, it will trigger an error response. | 

### Return type

[**PathVisTestResults**](PathVisTestResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


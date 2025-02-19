# HTTPServerEndpointScheduledTestResultsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHttpServerScheduledTestResults**](HTTPServerEndpointScheduledTestResultsAPI.md#GetHttpServerScheduledTestResults) | **Get** /endpoint/test-results/scheduled-tests/{testId}/http-server | Retrieve HTTP server scheduled test results
[**GetMultiTestFilteredHttpServerScheduledTestResults**](HTTPServerEndpointScheduledTestResultsAPI.md#GetMultiTestFilteredHttpServerScheduledTestResults) | **Post** /endpoint/test-results/scheduled-tests/http-server/filter | Filter HTTP server scheduled test results



## GetHttpServerScheduledTestResults

> HttpEndpointTestResults GetHttpServerScheduledTestResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Expand(expand).Execute()

Retrieve HTTP server scheduled test results



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
	expand := []endpointtestresults.ExpandEndpointHttpServerOptions{endpointtestresults.ExpandEndpointHttpServerOptions("header")} // []ExpandEndpointHttpServerOptions | This parameter is optional and determines whether to expand resources related to test results. By default, no expansion occurs when this query parameter is omitted. To expand a specific resource, such as \"header,\" append `?expand=header` to the query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.HTTPServerEndpointScheduledTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetHttpServerScheduledTestResults(testId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPServerEndpointScheduledTestResultsAPI.GetHttpServerScheduledTestResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHttpServerScheduledTestResults`: HttpEndpointTestResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `HTTPServerEndpointScheduledTestResultsAPI.GetHttpServerScheduledTestResults`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetHttpServerScheduledTestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **expand** | [**[]ExpandEndpointHttpServerOptions**](ExpandEndpointHttpServerOptions.md) | This parameter is optional and determines whether to expand resources related to test results. By default, no expansion occurs when this query parameter is omitted. To expand a specific resource, such as \&quot;header,\&quot; append &#x60;?expand&#x3D;header&#x60; to the query. | 

### Return type

[**HttpEndpointTestResults**](HttpEndpointTestResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetMultiTestFilteredHttpServerScheduledTestResults

> HttpMultiEndpointTestResults GetMultiTestFilteredHttpServerScheduledTestResults().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Expand(expand).HttpEndpointTestsDataRoundsSearch(httpEndpointTestsDataRoundsSearch).Execute()

Filter HTTP server scheduled test results



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
	expand := []endpointtestresults.ExpandEndpointHttpServerOptions{endpointtestresults.ExpandEndpointHttpServerOptions("header")} // []ExpandEndpointHttpServerOptions | This parameter is optional and determines whether to expand resources related to test results. By default, no expansion occurs when this query parameter is omitted. To expand a specific resource, such as \"header,\" append `?expand=header` to the query. (optional)
	httpEndpointTestsDataRoundsSearch := *endpointtestresults.NewHttpEndpointTestsDataRoundsSearch() // HttpEndpointTestsDataRoundsSearch | Test data search filters. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.HTTPServerEndpointScheduledTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetMultiTestFilteredHttpServerScheduledTestResults().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Expand(expand).HttpEndpointTestsDataRoundsSearch(httpEndpointTestsDataRoundsSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPServerEndpointScheduledTestResultsAPI.GetMultiTestFilteredHttpServerScheduledTestResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultiTestFilteredHttpServerScheduledTestResults`: HttpMultiEndpointTestResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `HTTPServerEndpointScheduledTestResultsAPI.GetMultiTestFilteredHttpServerScheduledTestResults`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetMultiTestFilteredHttpServerScheduledTestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **expand** | [**[]ExpandEndpointHttpServerOptions**](ExpandEndpointHttpServerOptions.md) | This parameter is optional and determines whether to expand resources related to test results. By default, no expansion occurs when this query parameter is omitted. To expand a specific resource, such as \&quot;header,\&quot; append &#x60;?expand&#x3D;header&#x60; to the query. | 
 **httpEndpointTestsDataRoundsSearch** | [**HttpEndpointTestsDataRoundsSearch**](HttpEndpointTestsDataRoundsSearch.md) | Test data search filters. | 

### Return type

[**HttpMultiEndpointTestResults**](HttpMultiEndpointTestResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


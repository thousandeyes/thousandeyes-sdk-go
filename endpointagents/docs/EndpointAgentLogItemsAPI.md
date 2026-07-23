# EndpointAgentLogItemsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEndpointAgentLogItems**](EndpointAgentLogItemsAPI.md#GetEndpointAgentLogItems) | **Get** /endpoint/agents/{agentId}/logs | List endpoint agent log items



## GetEndpointAgentLogItems

> EndpointAgentLogItemsResponse GetEndpointAgentLogItems(agentId).Aid(aid).Max(max).Cursor(cursor).Window(window).StartDate(startDate).EndDate(endDate).Execute()

List endpoint agent log items



### Example

```go
package main

import (
	"fmt"
	"os"
    "time"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointagents"
)

func main() {
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the agent to operate on.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	max := int32(1000) // int32 | Maximum number of log items returned per page. (optional) (default to 1000)
	cursor := "WyIxNzA5MjQwMDAwMDAwIl0=" // string | Opaque cursor from the `_links.next.href` URL in the previous response. (optional)
	window := "12h" // string | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: `s` for seconds (default if no type is specified), `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. For a precise date range, use `startDate` and `endDate`. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointagents.EndpointAgentLogItemsAPIService)(&apiClient.Common)

	resp, r, err := api.GetEndpointAgentLogItems(agentId).Aid(aid).Max(max).Cursor(cursor).Window(window).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentLogItemsAPI.GetEndpointAgentLogItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointAgentLogItems`: EndpointAgentLogItemsResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentLogItemsAPI.GetEndpointAgentLogItems`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**agentId** | **string** | The identifier of the agent to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetEndpointAgentLogItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **max** | **int32** | Maximum number of log items returned per page. | [default to 1000]
 **cursor** | **string** | Opaque cursor from the &#x60;_links.next.href&#x60; URL in the previous response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 

### Return type

[**EndpointAgentLogItemsResponse**](EndpointAgentLogItemsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


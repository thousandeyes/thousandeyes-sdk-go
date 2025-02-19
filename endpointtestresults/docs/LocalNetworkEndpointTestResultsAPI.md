# LocalNetworkEndpointTestResultsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterLocalNetworksTestResultsTopologies**](LocalNetworkEndpointTestResultsAPI.md#FilterLocalNetworksTestResultsTopologies) | **Post** /endpoint/test-results/local-networks/topologies/filter | List endpoint network topologies probes
[**GetLocalNetworksTestResults**](LocalNetworkEndpointTestResultsAPI.md#GetLocalNetworksTestResults) | **Get** /endpoint/test-results/local-networks | List local networks
[**GetLocalNetworksTestResultsTopology**](LocalNetworkEndpointTestResultsAPI.md#GetLocalNetworksTestResultsTopology) | **Get** /endpoint/test-results/local-networks/topologies/{networkTopologyId} | Retrieve endpoint local network topology



## FilterLocalNetworksTestResultsTopologies

> LocalNetworkTopologyResults FilterLocalNetworksTestResultsTopologies().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Expand(expand).EndpointNetworkTopologyResultRequest(endpointNetworkTopologyResultRequest).Execute()

List endpoint network topologies probes



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
	expand := []endpointtestresults.ExpandLocalNetworkTopologyOptions{endpointtestresults.ExpandLocalNetworkTopologyOptions("system-metric-detail")} // []ExpandLocalNetworkTopologyOptions | This parameter is optional and determines whether to expand resources related to local network topologies. By default, no expansion occurs when this query parameter is omitted. To expand a specific resource, such as `systemMetricDetails`, append  `?expand=system-metric-detail` to the query. (optional)
	endpointNetworkTopologyResultRequest := *endpointtestresults.NewEndpointNetworkTopologyResultRequest() // EndpointNetworkTopologyResultRequest |  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.LocalNetworkEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.FilterLocalNetworksTestResultsTopologies().Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Expand(expand).EndpointNetworkTopologyResultRequest(endpointNetworkTopologyResultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalNetworkEndpointTestResultsAPI.FilterLocalNetworksTestResultsTopologies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterLocalNetworksTestResultsTopologies`: LocalNetworkTopologyResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `LocalNetworkEndpointTestResultsAPI.FilterLocalNetworksTestResultsTopologies`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiFilterLocalNetworksTestResultsTopologiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **expand** | [**[]ExpandLocalNetworkTopologyOptions**](ExpandLocalNetworkTopologyOptions.md) | This parameter is optional and determines whether to expand resources related to local network topologies. By default, no expansion occurs when this query parameter is omitted. To expand a specific resource, such as &#x60;systemMetricDetails&#x60;, append  &#x60;?expand&#x3D;system-metric-detail&#x60; to the query. | 
 **endpointNetworkTopologyResultRequest** | [**EndpointNetworkTopologyResultRequest**](EndpointNetworkTopologyResultRequest.md) |  | 

### Return type

[**LocalNetworkTopologyResults**](LocalNetworkTopologyResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetLocalNetworksTestResults

> LocalNetworkResults GetLocalNetworksTestResults().Aid(aid).Execute()

List local networks



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.LocalNetworkEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetLocalNetworksTestResults().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalNetworkEndpointTestResultsAPI.GetLocalNetworksTestResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalNetworksTestResults`: LocalNetworkResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `LocalNetworkEndpointTestResultsAPI.GetLocalNetworksTestResults`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetLocalNetworksTestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**LocalNetworkResults**](LocalNetworkResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetLocalNetworksTestResultsTopology

> LocalNetworkTopologyDetailResults GetLocalNetworksTestResultsTopology(networkTopologyId).Aid(aid).Expand(expand).Execute()

Retrieve endpoint local network topology



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
	networkTopologyId := "00160:39c518560de9:1491651900:236e6f18" // string | The network topology ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []endpointtestresults.ExpandLocalNetworkTopologyOptions{endpointtestresults.ExpandLocalNetworkTopologyOptions("system-metric-detail")} // []ExpandLocalNetworkTopologyOptions | This parameter is optional and determines whether to expand resources related to local network topologies. By default, no expansion occurs when this query parameter is omitted. To expand a specific resource, such as `systemMetricDetails`, append  `?expand=system-metric-detail` to the query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointtestresults.LocalNetworkEndpointTestResultsAPIService)(&apiClient.Common)

	resp, r, err := api.GetLocalNetworksTestResultsTopology(networkTopologyId).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalNetworkEndpointTestResultsAPI.GetLocalNetworksTestResultsTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalNetworksTestResultsTopology`: LocalNetworkTopologyDetailResults
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `LocalNetworkEndpointTestResultsAPI.GetLocalNetworksTestResultsTopology`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**networkTopologyId** | **string** | The network topology ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetLocalNetworksTestResultsTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandLocalNetworkTopologyOptions**](ExpandLocalNetworkTopologyOptions.md) | This parameter is optional and determines whether to expand resources related to local network topologies. By default, no expansion occurs when this query parameter is omitted. To expand a specific resource, such as &#x60;systemMetricDetails&#x60;, append  &#x60;?expand&#x3D;system-metric-detail&#x60; to the query. | 

### Return type

[**LocalNetworkTopologyDetailResults**](LocalNetworkTopologyDetailResults.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


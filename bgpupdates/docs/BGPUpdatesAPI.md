# BGPUpdatesAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBgpUpdates**](BGPUpdatesAPI.md#GetBgpUpdates) | **Get** /bgp/updates | List BGP updates



## GetBgpUpdates

> BgpUpdates GetBgpUpdates().Aid(aid).Max(max).Cursor(cursor).Expand(expand).StartDate(startDate).EndDate(endDate).Prefix(prefix).OriginAs(originAs).AsPath(asPath).RpkiStatus(rpkiStatus).UpdateType(updateType).Monitor(monitor).Communities(communities).Execute()

List BGP updates



### Example

```go
package main

import (
	"fmt"
	"os"
    "time"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/bgpupdates"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	max := int32(20) // int32 | Maximum number of BGP updates to return. (optional) (default to 20)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	expand := []bgpupdates.BgpDataExpandOption{bgpupdates.BgpDataExpandOption("monitor")} // []BgpDataExpandOption | Optional expansions. Pass `expand=monitor` to replace monitor IDs with full BGP monitor objects. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	prefix := []string{"Inner_example"} // []string | Prefix CIDR filters. Repeat the parameter to filter by multiple prefixes. (optional)
	originAs := []int64{int64(123)} // []int64 | Origin AS filters. Repeat the parameter to filter by multiple ASNs. (optional)
	asPath := []string{"Inner_example"} // []string | AS path filter, expressed as a space-separated list of ASNs. Repeat the parameter to filter by multiple AS paths. (optional)
	rpkiStatus := []bgpupdates.BgpRpkiStatus{bgpupdates.BgpRpkiStatus("Valid")} // []BgpRpkiStatus | RPKI status filters. (optional)
	updateType := []bgpupdates.BgpUpdateType{bgpupdates.BgpUpdateType("announcement")} // []BgpUpdateType | BGP update type filters. (optional)
	monitor := []string{"Inner_example"} // []string | BGP monitor ID filters. Repeat the parameter to filter by multiple monitors. Get `monitorId` from the `/monitors` endpoint. (optional)
	communities := []string{"Inner_example"} // []string | BGP community filters. Repeat the parameter to filter by multiple communities. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*bgpupdates.BGPUpdatesAPIService)(&apiClient.Common)

	resp, r, err := api.GetBgpUpdates().Aid(aid).Max(max).Cursor(cursor).Expand(expand).StartDate(startDate).EndDate(endDate).Prefix(prefix).OriginAs(originAs).AsPath(asPath).RpkiStatus(rpkiStatus).UpdateType(updateType).Monitor(monitor).Communities(communities).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPUpdatesAPI.GetBgpUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBgpUpdates`: BgpUpdates
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `BGPUpdatesAPI.GetBgpUpdates`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetBgpUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **max** | **int32** | Maximum number of BGP updates to return. | [default to 20]
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **expand** | [**[]BgpDataExpandOption**](BgpDataExpandOption.md) | Optional expansions. Pass &#x60;expand&#x3D;monitor&#x60; to replace monitor IDs with full BGP monitor objects. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **prefix** | **[]string** | Prefix CIDR filters. Repeat the parameter to filter by multiple prefixes. | 
 **originAs** | **[]int64** | Origin AS filters. Repeat the parameter to filter by multiple ASNs. | 
 **asPath** | **[]string** | AS path filter, expressed as a space-separated list of ASNs. Repeat the parameter to filter by multiple AS paths. | 
 **rpkiStatus** | [**[]BgpRpkiStatus**](BgpRpkiStatus.md) | RPKI status filters. | 
 **updateType** | [**[]BgpUpdateType**](BgpUpdateType.md) | BGP update type filters. | 
 **monitor** | **[]string** | BGP monitor ID filters. Repeat the parameter to filter by multiple monitors. Get &#x60;monitorId&#x60; from the &#x60;/monitors&#x60; endpoint. | 
 **communities** | **[]string** | BGP community filters. Repeat the parameter to filter by multiple communities. | 

### Return type

[**BgpUpdates**](BgpUpdates.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


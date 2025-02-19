# DashboardsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardsAPI.md#CreateDashboard) | **Post** /dashboards | Create dashboard
[**DeleteDashboard**](DashboardsAPI.md#DeleteDashboard) | **Delete** /dashboards/{dashboardId} | Delete dashboard
[**GetDashboard**](DashboardsAPI.md#GetDashboard) | **Get** /dashboards/{dashboardId} | Retrieve dashboard
[**GetDashboardWidgetData**](DashboardsAPI.md#GetDashboardWidgetData) | **Get** /dashboards/{dashboardId}/widgets/{widgetId} | Retrieve dashboard widget data
[**GetDashboards**](DashboardsAPI.md#GetDashboards) | **Get** /dashboards | List dashboards
[**UpdateDashboard**](DashboardsAPI.md#UpdateDashboard) | **Put** /dashboards/{dashboardId} | Update dashboard



## CreateDashboard

> Dashboard CreateDashboard().Dashboard(dashboard).Aid(aid).Execute()

Create dashboard



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
	dashboard := *dashboards.NewDashboard() // Dashboard | Request body schema to create a dashboard.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateDashboard().Dashboard(dashboard).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.CreateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboard`: Dashboard
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.CreateDashboard`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**Dashboard**](Dashboard.md) | Request body schema to create a dashboard. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteDashboard

> DeleteDashboard(dashboardId).Aid(aid).Execute()

Delete dashboard



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
	dashboardId := "646f4d2ce3c99b0536c3821e" // string | A Identifier for a dashboard which can be obtained from the `/dashboards` endpoint.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsAPIService)(&apiClient.Common)

	r, err := api.DeleteDashboard(dashboardId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DeleteDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dashboardId** | **string** | A Identifier for a dashboard which can be obtained from the &#x60;/dashboards&#x60; endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteDashboardRequest struct via the builder pattern


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


## GetDashboard

> ApiDashboard GetDashboard(dashboardId).Aid(aid).Execute()

Retrieve dashboard



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
	dashboardId := "646f4d2ce3c99b0536c3821e" // string | A Identifier for a dashboard which can be obtained from the `/dashboards` endpoint.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDashboard(dashboardId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboard`: ApiDashboard
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboard`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dashboardId** | **string** | A Identifier for a dashboard which can be obtained from the &#x60;/dashboards&#x60; endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiDashboard**](ApiDashboard.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDashboardWidgetData

> ApiWidgetDataResponse GetDashboardWidgetData(dashboardIdwidgetId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Max(max).Cursor(cursor).Sort(sort).Order(order).Execute()

Retrieve dashboard widget data



### Example

```go
package main

import (
	"fmt"
	"os"
    "time"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/dashboards"
)

func main() {
	dashboardId := "646f4d2ce3c99b0536c3821e" // string | A Identifier for a dashboard which can be obtained from the `/dashboards` endpoint.
	widgetId := "unpmg" // string | A Identifier for a widget.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	window := "12h" // string | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: `s` for seconds (default if no type is specified), `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. For a precise date range, use `startDate` and `endDate`. (optional)
	startDate := time.Now() // time.Time | Use with the `endDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	endDate := time.Now() // time.Time | Defaults to current time the request is made. Use with the `startDate` parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can't be used with `window`. (optional)
	max := float32(10) // float32 | Optionally specify the maximum number of objects to retrieve. This only applies to the **Alert List** and **Test Table** Widgets. * The default for the **Alert List** widget is set by its limitBy configuration. * The default value for the **Test Table** widget is 10. (optional)
	cursor := "bGFzdFJvdW5kSWQ9MTY4MTQxMDQ4MA" // string | An optional pagination cursor. This parameter should not not be used directly. Instead, use the `_links` returned by the API. This feature is only available in the **Test Table** widget. (optional)
	sort := "alertStatus" // string | Optional sorting parameter with attributes listed comma-separated. This only applies to the **Alert List** and **Test Table** Widgets. * For the **Alert List** widget, you can sort by `alertStatus` or `startTime`. The default is `alertStatus`. * For the **Test Table** widget, you can sort by `alertStatus`, `testName`, or `testType`. The sequence might vary from the web application. The default sort attribute is `alertStatus`. (optional)
	order := dashboards.DashboardOrder("asc") // DashboardOrder | Optional sorting order parameter that accepts either `asc` (ascending) or `desc` (descending) values. This only applies to the **Alert List** and **Test Table** Widgets. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDashboardWidgetData(dashboardIdwidgetId).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Max(max).Cursor(cursor).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboardWidgetData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardWidgetData`: ApiWidgetDataResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboardWidgetData`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dashboardId** | **string** | A Identifier for a dashboard which can be obtained from the &#x60;/dashboards&#x60; endpoint. | **widgetId** | **string** | A Identifier for a widget. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDashboardWidgetDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **window** | **string** | A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;. | 
 **startDate** | **time.Time** | Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **endDate** | **time.Time** | Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;. | 
 **max** | **float32** | Optionally specify the maximum number of objects to retrieve. This only applies to the **Alert List** and **Test Table** Widgets. * The default for the **Alert List** widget is set by its limitBy configuration. * The default value for the **Test Table** widget is 10. | 
 **cursor** | **string** | An optional pagination cursor. This parameter should not not be used directly. Instead, use the &#x60;_links&#x60; returned by the API. This feature is only available in the **Test Table** widget. | 
 **sort** | **string** | Optional sorting parameter with attributes listed comma-separated. This only applies to the **Alert List** and **Test Table** Widgets. * For the **Alert List** widget, you can sort by &#x60;alertStatus&#x60; or &#x60;startTime&#x60;. The default is &#x60;alertStatus&#x60;. * For the **Test Table** widget, you can sort by &#x60;alertStatus&#x60;, &#x60;testName&#x60;, or &#x60;testType&#x60;. The sequence might vary from the web application. The default sort attribute is &#x60;alertStatus&#x60;. | 
 **order** | [**DashboardOrder**](DashboardOrder.md) | Optional sorting order parameter that accepts either &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) values. This only applies to the **Alert List** and **Test Table** Widgets. | 

### Return type

[**ApiWidgetDataResponse**](ApiWidgetDataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDashboards

> []ApiDashboard GetDashboards().Aid(aid).Execute()

List dashboards



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDashboards().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboards`: []ApiDashboard
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboards`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**[]ApiDashboard**](ApiDashboard.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateDashboard

> Dashboard UpdateDashboard(dashboardId).Dashboard(dashboard).Aid(aid).Execute()

Update dashboard



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
	dashboardId := "646f4d2ce3c99b0536c3821e" // string | A Identifier for a dashboard which can be obtained from the `/dashboards` endpoint.
	dashboard := *dashboards.NewDashboard() // Dashboard | Request body schema to update a dashboard.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateDashboard(dashboardId).Dashboard(dashboard).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.UpdateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboard`: Dashboard
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.UpdateDashboard`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dashboardId** | **string** | A Identifier for a dashboard which can be obtained from the &#x60;/dashboards&#x60; endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**Dashboard**](Dashboard.md) | Request body schema to update a dashboard. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


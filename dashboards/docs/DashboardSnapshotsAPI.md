# DashboardSnapshotsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboardSnapshot**](DashboardSnapshotsAPI.md#CreateDashboardSnapshot) | **Post** /dashboard-snapshots | Create dashboard snapshot
[**DeleteDashboardSnapshot**](DashboardSnapshotsAPI.md#DeleteDashboardSnapshot) | **Delete** /dashboard-snapshots/{snapshotId} | Delete dashboard snapshot
[**GetDashboardSnapshot**](DashboardSnapshotsAPI.md#GetDashboardSnapshot) | **Get** /dashboard-snapshots/{snapshotId} | Retrieve dashboard snapshot
[**GetDashboardSnapshotWidgetData**](DashboardSnapshotsAPI.md#GetDashboardSnapshotWidgetData) | **Get** /dashboard-snapshots/{snapshotId}/widgets/{widgetId} | Retrieve dashboard snapshot data
[**GetDashboardSnapshots**](DashboardSnapshotsAPI.md#GetDashboardSnapshots) | **Get** /dashboard-snapshots | List dashboard snapshots
[**UpdateDashboardSnapshotExpirationDate**](DashboardSnapshotsAPI.md#UpdateDashboardSnapshotExpirationDate) | **Patch** /dashboard-snapshots/{snapshotId} | Update snapshot expiration



## CreateDashboardSnapshot

> DashboardSnapshotResponse CreateDashboardSnapshot().GenerateDashboardSnapshotRequest(generateDashboardSnapshotRequest).Aid(aid).Execute()

Create dashboard snapshot



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
	generateDashboardSnapshotRequest := *dashboards.NewGenerateDashboardSnapshotRequest() // GenerateDashboardSnapshotRequest | Request body schema to create a dashboard snapshot.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardSnapshotsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateDashboardSnapshot().GenerateDashboardSnapshotRequest(generateDashboardSnapshotRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSnapshotsAPI.CreateDashboardSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboardSnapshot`: DashboardSnapshotResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardSnapshotsAPI.CreateDashboardSnapshot`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateDashboardSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateDashboardSnapshotRequest** | [**GenerateDashboardSnapshotRequest**](GenerateDashboardSnapshotRequest.md) | Request body schema to create a dashboard snapshot. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**DashboardSnapshotResponse**](DashboardSnapshotResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteDashboardSnapshot

> DeleteDashboardSnapshot(snapshotId).Aid(aid).Execute()

Delete dashboard snapshot



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
	snapshotId := "d28bb71f-5a47-4783-8f12-d4b115e61b0c" // string | A Identifier for a dashboard snapshot which can be obtained from the `/dashboards-snapshots` endpoint.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardSnapshotsAPIService)(&apiClient.Common)

	r, err := api.DeleteDashboardSnapshot(snapshotId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSnapshotsAPI.DeleteDashboardSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**snapshotId** | **string** | A Identifier for a dashboard snapshot which can be obtained from the &#x60;/dashboards-snapshots&#x60; endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteDashboardSnapshotRequest struct via the builder pattern


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


## GetDashboardSnapshot

> ApiDashboardSnapshot GetDashboardSnapshot(snapshotId).Aid(aid).Execute()

Retrieve dashboard snapshot



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
	snapshotId := "d28bb71f-5a47-4783-8f12-d4b115e61b0c" // string | A Identifier for a dashboard snapshot which can be obtained from the `/dashboards-snapshots` endpoint.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardSnapshotsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDashboardSnapshot(snapshotId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSnapshotsAPI.GetDashboardSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardSnapshot`: ApiDashboardSnapshot
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardSnapshotsAPI.GetDashboardSnapshot`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**snapshotId** | **string** | A Identifier for a dashboard snapshot which can be obtained from the &#x60;/dashboards-snapshots&#x60; endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDashboardSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiDashboardSnapshot**](ApiDashboardSnapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDashboardSnapshotWidgetData

> ApiWidgetDataSnapshotResponse GetDashboardSnapshotWidgetData(snapshotIdwidgetId).Aid(aid).Execute()

Retrieve dashboard snapshot data



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
	snapshotId := "d28bb71f-5a47-4783-8f12-d4b115e61b0c" // string | A Identifier for a dashboard snapshot which can be obtained from the `/dashboards-snapshots` endpoint.
	widgetId := "unpmg" // string | A Identifier for a widget.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardSnapshotsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDashboardSnapshotWidgetData(snapshotIdwidgetId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSnapshotsAPI.GetDashboardSnapshotWidgetData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardSnapshotWidgetData`: ApiWidgetDataSnapshotResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardSnapshotsAPI.GetDashboardSnapshotWidgetData`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**snapshotId** | **string** | A Identifier for a dashboard snapshot which can be obtained from the &#x60;/dashboards-snapshots&#x60; endpoint. | **widgetId** | **string** | A Identifier for a widget. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetDashboardSnapshotWidgetDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiWidgetDataSnapshotResponse**](ApiWidgetDataSnapshotResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetDashboardSnapshots

> DashboardSnapshotsPage GetDashboardSnapshots().Aid(aid).DashboardId(dashboardId).Cursor(cursor).Execute()

List dashboard snapshots



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
	dashboardId := "646f4d2ce3c99b0536c3821e" // string |  (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardSnapshotsAPIService)(&apiClient.Common)

	resp, r, err := api.GetDashboardSnapshots().Aid(aid).DashboardId(dashboardId).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSnapshotsAPI.GetDashboardSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardSnapshots`: DashboardSnapshotsPage
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `DashboardSnapshotsAPI.GetDashboardSnapshots`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetDashboardSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **dashboardId** | **string** |  | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 

### Return type

[**DashboardSnapshotsPage**](DashboardSnapshotsPage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateDashboardSnapshotExpirationDate

> UpdateDashboardSnapshotExpirationDate(snapshotId).UpdateSnapshotExpirationDateApiRequest(updateSnapshotExpirationDateApiRequest).Aid(aid).Execute()

Update snapshot expiration



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
	snapshotId := "d28bb71f-5a47-4783-8f12-d4b115e61b0c" // string | A Identifier for a dashboard snapshot which can be obtained from the `/dashboards-snapshots` endpoint.
	updateSnapshotExpirationDateApiRequest := *dashboards.NewUpdateSnapshotExpirationDateApiRequest() // UpdateSnapshotExpirationDateApiRequest | Request body schema to update a snapshot expiration.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*dashboards.DashboardSnapshotsAPIService)(&apiClient.Common)

	r, err := api.UpdateDashboardSnapshotExpirationDate(snapshotId).UpdateSnapshotExpirationDateApiRequest(updateSnapshotExpirationDateApiRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSnapshotsAPI.UpdateDashboardSnapshotExpirationDate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**snapshotId** | **string** | A Identifier for a dashboard snapshot which can be obtained from the &#x60;/dashboards-snapshots&#x60; endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateDashboardSnapshotExpirationDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSnapshotExpirationDateApiRequest** | [**UpdateSnapshotExpirationDateApiRequest**](UpdateSnapshotExpirationDateApiRequest.md) | Request body schema to update a snapshot expiration. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


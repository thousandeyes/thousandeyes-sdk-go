# AlertSuppressionWindowsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertSuppressionWindow**](AlertSuppressionWindowsAPI.md#CreateAlertSuppressionWindow) | **Post** /alert-suppression-windows | Create alert suppression window
[**DeleteAlertSuppressionWindow**](AlertSuppressionWindowsAPI.md#DeleteAlertSuppressionWindow) | **Delete** /alert-suppression-windows/{windowId} | Delete alert suppression window
[**GetAlertSuppressionWindow**](AlertSuppressionWindowsAPI.md#GetAlertSuppressionWindow) | **Get** /alert-suppression-windows/{windowId} | Retrieve alert suppression window
[**GetAlertSuppressionWindows**](AlertSuppressionWindowsAPI.md#GetAlertSuppressionWindows) | **Get** /alert-suppression-windows | List alert suppression windows
[**UpdateAlertSuppressionWindow**](AlertSuppressionWindowsAPI.md#UpdateAlertSuppressionWindow) | **Put** /alert-suppression-windows/{windowId} | Update alert suppression window



## CreateAlertSuppressionWindow

> AlertSuppressionWindowDetail CreateAlertSuppressionWindow().AlertSuppressionWindowRequest(alertSuppressionWindowRequest).Aid(aid).Expand(expand).Execute()

Create alert suppression window



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	alertSuppressionWindowRequest := *alerts.NewAlertSuppressionWindowRequest() // AlertSuppressionWindowRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []alerts.ExpandAlertTestOptions{alerts.ExpandAlertTestOptions("test")} // []ExpandAlertTestOptions | Optional parameter on whether or not to expand alert related resources.  Without this parameter, there's no default expansion. For example, to expand the \"tests\" resource, use the `?expand=test` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertSuppressionWindowsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAlertSuppressionWindow().AlertSuppressionWindowRequest(alertSuppressionWindowRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertSuppressionWindowsAPI.CreateAlertSuppressionWindow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertSuppressionWindow`: AlertSuppressionWindowDetail
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AlertSuppressionWindowsAPI.CreateAlertSuppressionWindow`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAlertSuppressionWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertSuppressionWindowRequest** | [**AlertSuppressionWindowRequest**](AlertSuppressionWindowRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandAlertTestOptions**](ExpandAlertTestOptions.md) | Optional parameter on whether or not to expand alert related resources.  Without this parameter, there&#39;s no default expansion. For example, to expand the \&quot;tests\&quot; resource, use the &#x60;?expand&#x3D;test&#x60; query. | 

### Return type

[**AlertSuppressionWindowDetail**](AlertSuppressionWindowDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteAlertSuppressionWindow

> DeleteAlertSuppressionWindow(windowId).Aid(aid).Execute()

Delete alert suppression window



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	windowId := "2411" // string | Unique window ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertSuppressionWindowsAPIService)(&apiClient.Common)

	r, err := api.DeleteAlertSuppressionWindow(windowId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertSuppressionWindowsAPI.DeleteAlertSuppressionWindow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**windowId** | **string** | Unique window ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteAlertSuppressionWindowRequest struct via the builder pattern


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


## GetAlertSuppressionWindow

> AlertSuppressionWindowDetail GetAlertSuppressionWindow(windowId).Aid(aid).Expand(expand).Execute()

Retrieve alert suppression window



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	windowId := "2411" // string | Unique window ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []alerts.ExpandAlertTestOptions{alerts.ExpandAlertTestOptions("test")} // []ExpandAlertTestOptions | Optional parameter on whether or not to expand alert related resources.  Without this parameter, there's no default expansion. For example, to expand the \"tests\" resource, use the `?expand=test` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertSuppressionWindowsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAlertSuppressionWindow(windowId).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertSuppressionWindowsAPI.GetAlertSuppressionWindow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertSuppressionWindow`: AlertSuppressionWindowDetail
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AlertSuppressionWindowsAPI.GetAlertSuppressionWindow`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**windowId** | **string** | Unique window ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAlertSuppressionWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandAlertTestOptions**](ExpandAlertTestOptions.md) | Optional parameter on whether or not to expand alert related resources.  Without this parameter, there&#39;s no default expansion. For example, to expand the \&quot;tests\&quot; resource, use the &#x60;?expand&#x3D;test&#x60; query. | 

### Return type

[**AlertSuppressionWindowDetail**](AlertSuppressionWindowDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAlertSuppressionWindows

> AlertSuppressionWindows GetAlertSuppressionWindows().Aid(aid).Execute()

List alert suppression windows



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertSuppressionWindowsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAlertSuppressionWindows().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertSuppressionWindowsAPI.GetAlertSuppressionWindows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertSuppressionWindows`: AlertSuppressionWindows
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AlertSuppressionWindowsAPI.GetAlertSuppressionWindows`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAlertSuppressionWindowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AlertSuppressionWindows**](AlertSuppressionWindows.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAlertSuppressionWindow

> AlertSuppressionWindowDetail UpdateAlertSuppressionWindow(windowId).AlertSuppressionWindowRequest(alertSuppressionWindowRequest).Aid(aid).Expand(expand).Execute()

Update alert suppression window



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	windowId := "2411" // string | Unique window ID.
	alertSuppressionWindowRequest := *alerts.NewAlertSuppressionWindowRequest() // AlertSuppressionWindowRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []alerts.ExpandAlertTestOptions{alerts.ExpandAlertTestOptions("test")} // []ExpandAlertTestOptions | Optional parameter on whether or not to expand alert related resources.  Without this parameter, there's no default expansion. For example, to expand the \"tests\" resource, use the `?expand=test` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertSuppressionWindowsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAlertSuppressionWindow(windowId).AlertSuppressionWindowRequest(alertSuppressionWindowRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertSuppressionWindowsAPI.UpdateAlertSuppressionWindow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertSuppressionWindow`: AlertSuppressionWindowDetail
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AlertSuppressionWindowsAPI.UpdateAlertSuppressionWindow`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**windowId** | **string** | Unique window ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAlertSuppressionWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertSuppressionWindowRequest** | [**AlertSuppressionWindowRequest**](AlertSuppressionWindowRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandAlertTestOptions**](ExpandAlertTestOptions.md) | Optional parameter on whether or not to expand alert related resources.  Without this parameter, there&#39;s no default expansion. For example, to expand the \&quot;tests\&quot; resource, use the &#x60;?expand&#x3D;test&#x60; query. | 

### Return type

[**AlertSuppressionWindowDetail**](AlertSuppressionWindowDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


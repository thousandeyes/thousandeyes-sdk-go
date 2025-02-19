# PathVisualizationInterfaceGroupsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePathVisInterfaceGroups**](PathVisualizationInterfaceGroupsAPI.md#CreatePathVisInterfaceGroups) | **Post** /network/path-vis/interface-groups | Create interface group for path visualization
[**DeletePathVisInterfaceGroup**](PathVisualizationInterfaceGroupsAPI.md#DeletePathVisInterfaceGroup) | **Delete** /network/path-vis/interface-groups/{interfaceGroupId} | Delete interface group
[**GetPathVisInterfaceGroups**](PathVisualizationInterfaceGroupsAPI.md#GetPathVisInterfaceGroups) | **Get** /network/path-vis/interface-groups | List interface groups for path visualization
[**UpdatePathVisInterfaceGroup**](PathVisualizationInterfaceGroupsAPI.md#UpdatePathVisInterfaceGroup) | **Put** /network/path-vis/interface-groups/{interfaceGroupId} | Update interface group



## CreatePathVisInterfaceGroups

> InterfaceGroup CreatePathVisInterfaceGroups().InterfaceGroup(interfaceGroup).Aid(aid).Execute()

Create interface group for path visualization



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	interfaceGroup := *tests.NewInterfaceGroup() // InterfaceGroup | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.PathVisualizationInterfaceGroupsAPIService)(&apiClient.Common)

	resp, r, err := api.CreatePathVisInterfaceGroups().InterfaceGroup(interfaceGroup).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PathVisualizationInterfaceGroupsAPI.CreatePathVisInterfaceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePathVisInterfaceGroups`: InterfaceGroup
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `PathVisualizationInterfaceGroupsAPI.CreatePathVisInterfaceGroups`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreatePathVisInterfaceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interfaceGroup** | [**InterfaceGroup**](InterfaceGroup.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**InterfaceGroup**](InterfaceGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeletePathVisInterfaceGroup

> DeletePathVisInterfaceGroup(interfaceGroupId).Aid(aid).Execute()

Delete interface group



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	interfaceGroupId := "281474976710706" // string | ID of the network path vis interface group
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.PathVisualizationInterfaceGroupsAPIService)(&apiClient.Common)

	r, err := api.DeletePathVisInterfaceGroup(interfaceGroupId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PathVisualizationInterfaceGroupsAPI.DeletePathVisInterfaceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**interfaceGroupId** | **string** | ID of the network path vis interface group | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeletePathVisInterfaceGroupRequest struct via the builder pattern


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


## GetPathVisInterfaceGroups

> InterfaceGroups GetPathVisInterfaceGroups().Aid(aid).Execute()

List interface groups for path visualization



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.PathVisualizationInterfaceGroupsAPIService)(&apiClient.Common)

	resp, r, err := api.GetPathVisInterfaceGroups().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PathVisualizationInterfaceGroupsAPI.GetPathVisInterfaceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPathVisInterfaceGroups`: InterfaceGroups
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `PathVisualizationInterfaceGroupsAPI.GetPathVisInterfaceGroups`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetPathVisInterfaceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**InterfaceGroups**](InterfaceGroups.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdatePathVisInterfaceGroup

> InterfaceGroup UpdatePathVisInterfaceGroup(interfaceGroupId).InterfaceGroup(interfaceGroup).Aid(aid).Execute()

Update interface group



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	interfaceGroupId := "281474976710706" // string | ID of the network path vis interface group
	interfaceGroup := *tests.NewInterfaceGroup() // InterfaceGroup | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.PathVisualizationInterfaceGroupsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdatePathVisInterfaceGroup(interfaceGroupId).InterfaceGroup(interfaceGroup).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PathVisualizationInterfaceGroupsAPI.UpdatePathVisInterfaceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePathVisInterfaceGroup`: InterfaceGroup
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `PathVisualizationInterfaceGroupsAPI.UpdatePathVisInterfaceGroup`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**interfaceGroupId** | **string** | ID of the network path vis interface group | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdatePathVisInterfaceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interfaceGroup** | [**InterfaceGroup**](InterfaceGroup.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**InterfaceGroup**](InterfaceGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


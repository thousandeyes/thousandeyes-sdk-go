# EndpointAgentLabelsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpointLabel**](EndpointAgentLabelsAPI.md#CreateEndpointLabel) | **Post** /endpoint/labels | Create label
[**DeleteEndpointLabel**](EndpointAgentLabelsAPI.md#DeleteEndpointLabel) | **Delete** /endpoint/labels/{id} | Delete label
[**GetEndpointLabel**](EndpointAgentLabelsAPI.md#GetEndpointLabel) | **Get** /endpoint/labels/{id} | Retrieve label
[**GetEndpointLabels**](EndpointAgentLabelsAPI.md#GetEndpointLabels) | **Get** /endpoint/labels | List labels
[**UpdateEndpointLabel**](EndpointAgentLabelsAPI.md#UpdateEndpointLabel) | **Patch** /endpoint/labels/{id} | Update label



## CreateEndpointLabel

> LabelResponse CreateEndpointLabel().Aid(aid).LabelRequest(labelRequest).Execute()

Create label



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointlabels"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	labelRequest := *endpointlabels.NewLabelRequest("Head office meeting rooms", endpointlabels.MatchType("and"), []endpointlabels.Filter{*endpointlabels.NewFilter()}) // LabelRequest | Label settings (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointlabels.EndpointAgentLabelsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateEndpointLabel().Aid(aid).LabelRequest(labelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentLabelsAPI.CreateEndpointLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpointLabel`: LabelResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentLabelsAPI.CreateEndpointLabel`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateEndpointLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **labelRequest** | [**LabelRequest**](LabelRequest.md) | Label settings | 

### Return type

[**LabelResponse**](LabelResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteEndpointLabel

> DeleteEndpointLabel(id).Aid(aid).Execute()

Delete label



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointlabels"
)

func main() {
	id := "id_example" // string | The unique identifier of the label to operate on.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointlabels.EndpointAgentLabelsAPIService)(&apiClient.Common)

	r, err := api.DeleteEndpointLabel(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentLabelsAPI.DeleteEndpointLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The unique identifier of the label to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteEndpointLabelRequest struct via the builder pattern


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


## GetEndpointLabel

> LabelResponse GetEndpointLabel(id).Expand(expand).Aid(aid).Execute()

Retrieve label



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointlabels"
)

func main() {
	id := "id_example" // string | The unique identifier of the label to operate on.
	expand := []endpointlabels.ExpandLabelOptions{endpointlabels.ExpandLabelOptions("filters")} // []ExpandLabelOptions | This parameter is optional and determines whether to include additional details in the response. To specify multiple expansions, you can either separate the values with commas or specify the parameter multiple times. (optional)
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointlabels.EndpointAgentLabelsAPIService)(&apiClient.Common)

	resp, r, err := api.GetEndpointLabel(id).Expand(expand).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentLabelsAPI.GetEndpointLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointLabel`: LabelResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentLabelsAPI.GetEndpointLabel`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The unique identifier of the label to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetEndpointLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | [**[]ExpandLabelOptions**](ExpandLabelOptions.md) | This parameter is optional and determines whether to include additional details in the response. To specify multiple expansions, you can either separate the values with commas or specify the parameter multiple times. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**LabelResponse**](LabelResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetEndpointLabels

> Labels GetEndpointLabels().Max(max).Cursor(cursor).Expand(expand).Aid(aid).Execute()

List labels



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointlabels"
)

func main() {
	max := int32(5) // int32 | (Optional) Maximum number of objects to return. (optional)
	cursor := "cursor_example" // string | (Optional) Opaque cursor used for pagination. Clients should use `next` value from `_links` instead of this parameter. (optional)
	expand := []endpointlabels.ExpandLabelOptions{endpointlabels.ExpandLabelOptions("filters")} // []ExpandLabelOptions | This parameter is optional and determines whether to include additional details in the response. To specify multiple expansions, you can either separate the values with commas or specify the parameter multiple times. (optional)
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointlabels.EndpointAgentLabelsAPIService)(&apiClient.Common)

	resp, r, err := api.GetEndpointLabels().Max(max).Cursor(cursor).Expand(expand).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentLabelsAPI.GetEndpointLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointLabels`: Labels
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentLabelsAPI.GetEndpointLabels`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetEndpointLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int32** | (Optional) Maximum number of objects to return. | 
 **cursor** | **string** | (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter. | 
 **expand** | [**[]ExpandLabelOptions**](ExpandLabelOptions.md) | This parameter is optional and determines whether to include additional details in the response. To specify multiple expansions, you can either separate the values with commas or specify the parameter multiple times. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Labels**](Labels.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateEndpointLabel

> LabelResponse UpdateEndpointLabel(id).Aid(aid).Label(label).Execute()

Update label



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointlabels"
)

func main() {
	id := "id_example" // string | The unique identifier of the label to operate on.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	label := *endpointlabels.NewLabel() // Label | Fields to change on the agent (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointlabels.EndpointAgentLabelsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateEndpointLabel(id).Aid(aid).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAgentLabelsAPI.UpdateEndpointLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpointLabel`: LabelResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EndpointAgentLabelsAPI.UpdateEndpointLabel`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The unique identifier of the label to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateEndpointLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **label** | [**Label**](Label.md) | Fields to change on the agent | 

### Return type

[**LabelResponse**](LabelResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


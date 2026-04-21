# OperationConnectorsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOperationConnectors**](OperationConnectorsAPI.md#GetOperationConnectors) | **Get** /operations/{type}/{id}/connectors | Retrieve connectors assigned to an operation
[**SetOperationConnectors**](OperationConnectorsAPI.md#SetOperationConnectors) | **Put** /operations/{type}/{id}/connectors | Assign connectors to an operation



## GetOperationConnectors

> Assignments GetOperationConnectors(type_, id).Aid(aid).Execute()

Retrieve connectors assigned to an operation



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/connectors"
)

func main() {
	type_ := "webhooks" // string | The operation type.
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The operation ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.OperationConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.GetOperationConnectors(type_, id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationConnectorsAPI.GetOperationConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperationConnectors`: Assignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `OperationConnectorsAPI.GetOperationConnectors`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**type_** | **string** | The operation type. | **id** | **string** | The operation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetOperationConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Assignments**](Assignments.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## SetOperationConnectors

> Assignments SetOperationConnectors(type_, id).RequestBody(requestBody).Aid(aid).Execute()

Assign connectors to an operation



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/connectors"
)

func main() {
	type_ := "webhooks" // string | The operation type.
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The operation ID.
	requestBody := []string{"Property_example"} // []string | List of connector IDs to assign to the operation.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.OperationConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.SetOperationConnectors(type_, id).RequestBody(requestBody).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationConnectorsAPI.SetOperationConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOperationConnectors`: Assignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `OperationConnectorsAPI.SetOperationConnectors`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**type_** | **string** | The operation type. | **id** | **string** | The operation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiSetOperationConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | List of connector IDs to assign to the operation. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Assignments**](Assignments.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


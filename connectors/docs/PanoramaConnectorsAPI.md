# PanoramaConnectorsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePanoramaConnector**](PanoramaConnectorsAPI.md#CreatePanoramaConnector) | **Post** /connectors/panorama | Create Panorama connector
[**DeletePanoramaConnector**](PanoramaConnectorsAPI.md#DeletePanoramaConnector) | **Delete** /connectors/panorama/{id} | Delete Panorama connector
[**GetPanoramaConnector**](PanoramaConnectorsAPI.md#GetPanoramaConnector) | **Get** /connectors/panorama/{id} | Retrieve Panorama connector
[**GetPanoramaConnectorOperations**](PanoramaConnectorsAPI.md#GetPanoramaConnectorOperations) | **Get** /connectors/panorama/{id}/operations | List operation IDs for Panorama connector
[**GetPanoramaConnectors**](PanoramaConnectorsAPI.md#GetPanoramaConnectors) | **Get** /connectors/panorama | List Panorama connectors
[**SetPanoramaConnectorOperations**](PanoramaConnectorsAPI.md#SetPanoramaConnectorOperations) | **Put** /connectors/panorama/{id}/operations | Assign operations to Panorama connector
[**UpdatePanoramaConnector**](PanoramaConnectorsAPI.md#UpdatePanoramaConnector) | **Put** /connectors/panorama/{id} | Update Panorama connector



## CreatePanoramaConnector

> PanoramaConnector CreatePanoramaConnector().PanoramaConnector(panoramaConnector).Aid(aid).Execute()

Create Panorama connector



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
	panoramaConnector := *connectors.NewPanoramaConnector("panorama", "Branch Panorama", "https://panorama.example.com", connectors.PanoramaConnectorAuth{PanoramaKeyGenAuthentication: connectors.NewPanoramaKeyGenAuthentication("panorama-user", "Password_example", int64(40), "Type_example")}) // PanoramaConnector | Panorama connector configuration.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.PanoramaConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.CreatePanoramaConnector().PanoramaConnector(panoramaConnector).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PanoramaConnectorsAPI.CreatePanoramaConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePanoramaConnector`: PanoramaConnector
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `PanoramaConnectorsAPI.CreatePanoramaConnector`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreatePanoramaConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **panoramaConnector** | [**PanoramaConnector**](PanoramaConnector.md) | Panorama connector configuration. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**PanoramaConnector**](PanoramaConnector.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeletePanoramaConnector

> DeletePanoramaConnector(id).Aid(aid).Execute()

Delete Panorama connector



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
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The connector ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.PanoramaConnectorsAPIService)(&apiClient.Common)

	r, err := api.DeletePanoramaConnector(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PanoramaConnectorsAPI.DeletePanoramaConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeletePanoramaConnectorRequest struct via the builder pattern


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


## GetPanoramaConnector

> PanoramaConnector GetPanoramaConnector(id).Aid(aid).Execute()

Retrieve Panorama connector



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
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The connector ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.PanoramaConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.GetPanoramaConnector(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PanoramaConnectorsAPI.GetPanoramaConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPanoramaConnector`: PanoramaConnector
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `PanoramaConnectorsAPI.GetPanoramaConnector`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetPanoramaConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**PanoramaConnector**](PanoramaConnector.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetPanoramaConnectorOperations

> Assignments GetPanoramaConnectorOperations(id).Aid(aid).Execute()

List operation IDs for Panorama connector



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
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The connector ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.PanoramaConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.GetPanoramaConnectorOperations(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PanoramaConnectorsAPI.GetPanoramaConnectorOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPanoramaConnectorOperations`: Assignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `PanoramaConnectorsAPI.GetPanoramaConnectorOperations`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetPanoramaConnectorOperationsRequest struct via the builder pattern


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


## GetPanoramaConnectors

> PanoramaConnectors GetPanoramaConnectors().Aid(aid).Execute()

List Panorama connectors



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.PanoramaConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.GetPanoramaConnectors().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PanoramaConnectorsAPI.GetPanoramaConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPanoramaConnectors`: PanoramaConnectors
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `PanoramaConnectorsAPI.GetPanoramaConnectors`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetPanoramaConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**PanoramaConnectors**](PanoramaConnectors.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## SetPanoramaConnectorOperations

> Assignments SetPanoramaConnectorOperations(id).RequestBody(requestBody).Aid(aid).Execute()

Assign operations to Panorama connector



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
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The connector ID.
	requestBody := []string{"Property_example"} // []string | List of operation IDs to assign to the connector.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.PanoramaConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.SetPanoramaConnectorOperations(id).RequestBody(requestBody).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PanoramaConnectorsAPI.SetPanoramaConnectorOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPanoramaConnectorOperations`: Assignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `PanoramaConnectorsAPI.SetPanoramaConnectorOperations`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiSetPanoramaConnectorOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | List of operation IDs to assign to the connector. | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Assignments**](Assignments.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdatePanoramaConnector

> PanoramaConnector UpdatePanoramaConnector(id).PanoramaConnector(panoramaConnector).Aid(aid).Execute()

Update Panorama connector



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
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The connector ID.
	panoramaConnector := *connectors.NewPanoramaConnector("panorama", "Branch Panorama", "https://panorama.example.com", connectors.PanoramaConnectorAuth{PanoramaKeyGenAuthentication: connectors.NewPanoramaKeyGenAuthentication("panorama-user", "Password_example", int64(40), "Type_example")}) // PanoramaConnector | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.PanoramaConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdatePanoramaConnector(id).PanoramaConnector(panoramaConnector).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PanoramaConnectorsAPI.UpdatePanoramaConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePanoramaConnector`: PanoramaConnector
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `PanoramaConnectorsAPI.UpdatePanoramaConnector`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdatePanoramaConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **panoramaConnector** | [**PanoramaConnector**](PanoramaConnector.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**PanoramaConnector**](PanoramaConnector.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


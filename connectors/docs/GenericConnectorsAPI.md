# GenericConnectorsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGenericConnector**](GenericConnectorsAPI.md#CreateGenericConnector) | **Post** /connectors/generic | Create connector
[**DeleteGenericConnector**](GenericConnectorsAPI.md#DeleteGenericConnector) | **Delete** /connectors/generic/{id} | Delete connector
[**GetGenericConnector**](GenericConnectorsAPI.md#GetGenericConnector) | **Get** /connectors/generic/{id} | Retrieve connector
[**GetGenericConnectors**](GenericConnectorsAPI.md#GetGenericConnectors) | **Get** /connectors/generic | List connectors
[**ListGenericConnectorOperations**](GenericConnectorsAPI.md#ListGenericConnectorOperations) | **Get** /connectors/generic/{id}/operations | List operation IDs assigned to a connector
[**SetGenericConnectorOperations**](GenericConnectorsAPI.md#SetGenericConnectorOperations) | **Put** /connectors/generic/{id}/operations | Assign operations to a connector
[**UpdateGenericConnector**](GenericConnectorsAPI.md#UpdateGenericConnector) | **Put** /connectors/generic/{id} | Update connector



## CreateGenericConnector

> GenericConnector CreateGenericConnector().GenericConnector(genericConnector).Aid(aid).Execute()

Create connector



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
	genericConnector := *connectors.NewGenericConnector(connectors.ConnectorType("generic"), "Cisco Slack", "https://hooks.slack.com/services/abc/xyz") // GenericConnector | 
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.GenericConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateGenericConnector().GenericConnector(genericConnector).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericConnectorsAPI.CreateGenericConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGenericConnector`: GenericConnector
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `GenericConnectorsAPI.CreateGenericConnector`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateGenericConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericConnector** | [**GenericConnector**](GenericConnector.md) |  | 
 **aid** | **float32** | Account ID | 

### Return type

[**GenericConnector**](GenericConnector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteGenericConnector

> DeleteGenericConnector(id).Aid(aid).Execute()

Delete connector



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
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.GenericConnectorsAPIService)(&apiClient.Common)

	r, err := api.DeleteGenericConnector(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericConnectorsAPI.DeleteGenericConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteGenericConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **float32** | Account ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetGenericConnector

> GenericConnector GetGenericConnector(id).Aid(aid).Execute()

Retrieve connector



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
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.GenericConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.GetGenericConnector(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericConnectorsAPI.GetGenericConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericConnector`: GenericConnector
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `GenericConnectorsAPI.GetGenericConnector`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetGenericConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **float32** | Account ID | 

### Return type

[**GenericConnector**](GenericConnector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetGenericConnectors

> GenericConnectors GetGenericConnectors().Aid(aid).Execute()

List connectors



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
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.GenericConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.GetGenericConnectors().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericConnectorsAPI.GetGenericConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericConnectors`: GenericConnectors
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `GenericConnectorsAPI.GetGenericConnectors`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetGenericConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **float32** | Account ID | 

### Return type

[**GenericConnectors**](GenericConnectors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## ListGenericConnectorOperations

> Assignments ListGenericConnectorOperations(id).Aid(aid).Execute()

List operation IDs assigned to a connector



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
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.GenericConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.ListGenericConnectorOperations(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericConnectorsAPI.ListGenericConnectorOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGenericConnectorOperations`: Assignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `GenericConnectorsAPI.ListGenericConnectorOperations`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiListGenericConnectorOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **float32** | Account ID | 

### Return type

[**Assignments**](Assignments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## SetGenericConnectorOperations

> Assignments SetGenericConnectorOperations(id).RequestBody(requestBody).Aid(aid).Execute()

Assign operations to a connector



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
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.GenericConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.SetGenericConnectorOperations(id).RequestBody(requestBody).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericConnectorsAPI.SetGenericConnectorOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGenericConnectorOperations`: Assignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `GenericConnectorsAPI.SetGenericConnectorOperations`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiSetGenericConnectorOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | List of operation IDs to assign to the connector. | 
 **aid** | **float32** | Account ID | 

### Return type

[**Assignments**](Assignments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateGenericConnector

> GenericConnector UpdateGenericConnector(id).GenericConnector(genericConnector).Aid(aid).Execute()

Update connector



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
	genericConnector := *connectors.NewGenericConnector(connectors.ConnectorType("generic"), "Cisco Slack", "https://hooks.slack.com/services/abc/xyz") // GenericConnector | 
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.GenericConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateGenericConnector(id).GenericConnector(genericConnector).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericConnectorsAPI.UpdateGenericConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGenericConnector`: GenericConnector
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `GenericConnectorsAPI.UpdateGenericConnector`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateGenericConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericConnector** | [**GenericConnector**](GenericConnector.md) |  | 
 **aid** | **float32** | Account ID | 

### Return type

[**GenericConnector**](GenericConnector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


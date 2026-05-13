# CyberArkConjurConnectorsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConjurConnector**](CyberArkConjurConnectorsAPI.md#CreateConjurConnector) | **Post** /connectors/conjur | Create Conjur connector
[**DeleteConjurConnector**](CyberArkConjurConnectorsAPI.md#DeleteConjurConnector) | **Delete** /connectors/conjur/{id} | Delete a Conjur connector
[**GetConjurConnector**](CyberArkConjurConnectorsAPI.md#GetConjurConnector) | **Get** /connectors/conjur/{id} | Retrieve a Conjur connector
[**GetConjurConnectorOperations**](CyberArkConjurConnectorsAPI.md#GetConjurConnectorOperations) | **Get** /connectors/conjur/{id}/operations | List operation IDs for a Conjur connector
[**GetConjurConnectors**](CyberArkConjurConnectorsAPI.md#GetConjurConnectors) | **Get** /connectors/conjur | List Conjur connectors
[**SetConjurConnectorOperations**](CyberArkConjurConnectorsAPI.md#SetConjurConnectorOperations) | **Put** /connectors/conjur/{id}/operations | Assign operations to a Conjur connector
[**UpdateConjurConnector**](CyberArkConjurConnectorsAPI.md#UpdateConjurConnector) | **Put** /connectors/conjur/{id} | Update a Conjur connector



## CreateConjurConnector

> ConjurConnector CreateConjurConnector().ConjurConnector(conjurConnector).Aid(aid).Execute()

Create Conjur connector



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
	conjurConnector := *connectors.NewConjurConnector(connectors.ConnectorType("generic"), "Cisco Slack", "https://eval.conjur.org/secrets", "My CyberArk Account", *connectors.NewConjurHostAuthentication("host1", "abc123", connectors.AuthenticationType("basic"))) // ConjurConnector | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.CyberArkConjurConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateConjurConnector().ConjurConnector(conjurConnector).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CyberArkConjurConnectorsAPI.CreateConjurConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConjurConnector`: ConjurConnector
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CyberArkConjurConnectorsAPI.CreateConjurConnector`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateConjurConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conjurConnector** | [**ConjurConnector**](ConjurConnector.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ConjurConnector**](ConjurConnector.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteConjurConnector

> DeleteConjurConnector(id).ConfirmDisabledObjects(confirmDisabledObjects).Aid(aid).Execute()

Delete a Conjur connector



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
	confirmDisabledObjects := true // bool | Confirmation to disable affected objects (for example, tests) for Conjur connectors. (default to false)
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.CyberArkConjurConnectorsAPIService)(&apiClient.Common)

	r, err := api.DeleteConjurConnector(id).ConfirmDisabledObjects(confirmDisabledObjects).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CyberArkConjurConnectorsAPI.DeleteConjurConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteConjurConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirmDisabledObjects** | **bool** | Confirmation to disable affected objects (for example, tests) for Conjur connectors. | [default to false]
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetConjurConnector

> ConjurConnector GetConjurConnector(id).Aid(aid).Execute()

Retrieve a Conjur connector



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

	api := (*connectors.CyberArkConjurConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.GetConjurConnector(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CyberArkConjurConnectorsAPI.GetConjurConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConjurConnector`: ConjurConnector
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CyberArkConjurConnectorsAPI.GetConjurConnector`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetConjurConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ConjurConnector**](ConjurConnector.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetConjurConnectorOperations

> Assignments GetConjurConnectorOperations(id).Aid(aid).Execute()

List operation IDs for a Conjur connector



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

	api := (*connectors.CyberArkConjurConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.GetConjurConnectorOperations(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CyberArkConjurConnectorsAPI.GetConjurConnectorOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConjurConnectorOperations`: Assignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CyberArkConjurConnectorsAPI.GetConjurConnectorOperations`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetConjurConnectorOperationsRequest struct via the builder pattern


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


## GetConjurConnectors

> ConjurConnectors GetConjurConnectors().Aid(aid).Execute()

List Conjur connectors



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

	api := (*connectors.CyberArkConjurConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.GetConjurConnectors().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CyberArkConjurConnectorsAPI.GetConjurConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConjurConnectors`: ConjurConnectors
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CyberArkConjurConnectorsAPI.GetConjurConnectors`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetConjurConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ConjurConnectors**](ConjurConnectors.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## SetConjurConnectorOperations

> Assignments SetConjurConnectorOperations(id).ConfirmDisabledObjects(confirmDisabledObjects).RequestBody(requestBody).Aid(aid).Execute()

Assign operations to a Conjur connector



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
	confirmDisabledObjects := true // bool | Confirmation to disable affected objects (for example, tests) for Conjur connectors. (default to false)
	requestBody := []string{"Property_example"} // []string | List of operation IDs to assign to the connector.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.CyberArkConjurConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.SetConjurConnectorOperations(id).ConfirmDisabledObjects(confirmDisabledObjects).RequestBody(requestBody).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CyberArkConjurConnectorsAPI.SetConjurConnectorOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetConjurConnectorOperations`: Assignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CyberArkConjurConnectorsAPI.SetConjurConnectorOperations`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiSetConjurConnectorOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirmDisabledObjects** | **bool** | Confirmation to disable affected objects (for example, tests) for Conjur connectors. | [default to false]
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


## UpdateConjurConnector

> ConjurConnector UpdateConjurConnector(id).ConjurConnector(conjurConnector).Aid(aid).Execute()

Update a Conjur connector



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
	conjurConnector := *connectors.NewConjurConnector(connectors.ConnectorType("generic"), "Cisco Slack", "https://eval.conjur.org/secrets", "My CyberArk Account", *connectors.NewConjurHostAuthentication("host1", "abc123", connectors.AuthenticationType("basic"))) // ConjurConnector | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.CyberArkConjurConnectorsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateConjurConnector(id).ConjurConnector(conjurConnector).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CyberArkConjurConnectorsAPI.UpdateConjurConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConjurConnector`: ConjurConnector
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CyberArkConjurConnectorsAPI.UpdateConjurConnector`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The connector ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateConjurConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conjurConnector** | [**ConjurConnector**](ConjurConnector.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ConjurConnector**](ConjurConnector.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


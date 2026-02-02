# WebhookOperationsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookOperation**](WebhookOperationsAPI.md#CreateWebhookOperation) | **Post** /operations/webhooks | Create webhook operation
[**DeleteWebhookOperation**](WebhookOperationsAPI.md#DeleteWebhookOperation) | **Delete** /operations/webhooks/{id} | Delete webhook operation
[**GetWebhookOperation**](WebhookOperationsAPI.md#GetWebhookOperation) | **Get** /operations/webhooks/{id} | Retrieve webhook operation
[**GetWebhookOperations**](WebhookOperationsAPI.md#GetWebhookOperations) | **Get** /operations/webhooks | List webhook operations
[**UpdateWebhookOperation**](WebhookOperationsAPI.md#UpdateWebhookOperation) | **Put** /operations/webhooks/{id} | Update webhook operation



## CreateWebhookOperation

> WebhookOperation CreateWebhookOperation().WebhookOperation(webhookOperation).Aid(aid).Execute()

Create webhook operation



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
	webhookOperation := *connectors.NewWebhookOperation("My operation", connectors.OperationCategory("alerts"), connectors.OperationStatus("pending")) // WebhookOperation | 
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.WebhookOperationsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateWebhookOperation().WebhookOperation(webhookOperation).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookOperationsAPI.CreateWebhookOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookOperation`: WebhookOperation
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `WebhookOperationsAPI.CreateWebhookOperation`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateWebhookOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookOperation** | [**WebhookOperation**](WebhookOperation.md) |  | 
 **aid** | **float32** | Account ID | 

### Return type

[**WebhookOperation**](WebhookOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteWebhookOperation

> DeleteWebhookOperation(id).Aid(aid).Execute()

Delete webhook operation



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
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The operation ID.
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.WebhookOperationsAPIService)(&apiClient.Common)

	r, err := api.DeleteWebhookOperation(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookOperationsAPI.DeleteWebhookOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The operation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteWebhookOperationRequest struct via the builder pattern


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


## GetWebhookOperation

> WebhookOperation GetWebhookOperation(id).Aid(aid).Execute()

Retrieve webhook operation



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
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The operation ID.
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.WebhookOperationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetWebhookOperation(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookOperationsAPI.GetWebhookOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookOperation`: WebhookOperation
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `WebhookOperationsAPI.GetWebhookOperation`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The operation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetWebhookOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **float32** | Account ID | 

### Return type

[**WebhookOperation**](WebhookOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetWebhookOperations

> WebhookOperations GetWebhookOperations().Aid(aid).Execute()

List webhook operations



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

	api := (*connectors.WebhookOperationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetWebhookOperations().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookOperationsAPI.GetWebhookOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookOperations`: WebhookOperations
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `WebhookOperationsAPI.GetWebhookOperations`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetWebhookOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **float32** | Account ID | 

### Return type

[**WebhookOperations**](WebhookOperations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateWebhookOperation

> WebhookOperation UpdateWebhookOperation(id).WebhookOperation(webhookOperation).Aid(aid).Execute()

Update webhook operation



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
	id := "cb1b8033-ea2d-4e9b-a920-fe87850693cf" // string | The operation ID.
	webhookOperation := *connectors.NewWebhookOperation("My operation", connectors.OperationCategory("alerts"), connectors.OperationStatus("pending")) // WebhookOperation | 
	aid := float32(123456) // float32 | Account ID (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.WebhookOperationsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateWebhookOperation(id).WebhookOperation(webhookOperation).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookOperationsAPI.UpdateWebhookOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookOperation`: WebhookOperation
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `WebhookOperationsAPI.UpdateWebhookOperation`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The operation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateWebhookOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookOperation** | [**WebhookOperation**](WebhookOperation.md) |  | 
 **aid** | **float32** | Account ID | 

### Return type

[**WebhookOperation**](WebhookOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


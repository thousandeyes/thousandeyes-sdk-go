# CredentialVaultOperationsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialVaultOperation**](CredentialVaultOperationsAPI.md#CreateCredentialVaultOperation) | **Post** /operations/credential-vault | Create Credential Vault operation
[**DeleteCredentialVaultOperation**](CredentialVaultOperationsAPI.md#DeleteCredentialVaultOperation) | **Delete** /operations/credential-vault/{id} | Delete Credential Vault operation
[**GetCredentialVaultOperation**](CredentialVaultOperationsAPI.md#GetCredentialVaultOperation) | **Get** /operations/credential-vault/{id} | Get Credential Vault operation
[**GetCredentialVaultOperations**](CredentialVaultOperationsAPI.md#GetCredentialVaultOperations) | **Get** /operations/credential-vault | List Credential Vault operations
[**UpdateCredentialVaultOperation**](CredentialVaultOperationsAPI.md#UpdateCredentialVaultOperation) | **Put** /operations/credential-vault/{id} | Update Credential Vault operation



## CreateCredentialVaultOperation

> CredentialVaultOperation CreateCredentialVaultOperation().CredentialVaultOperation(credentialVaultOperation).Aid(aid).Execute()

Create Credential Vault operation



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
	credentialVaultOperation := *connectors.NewCredentialVaultOperation("My operation", []connectors.CredentialVaultSecret{*connectors.NewCredentialVaultSecret("secret_name", "secret/key")}) // CredentialVaultOperation | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.CredentialVaultOperationsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateCredentialVaultOperation().CredentialVaultOperation(credentialVaultOperation).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultOperationsAPI.CreateCredentialVaultOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCredentialVaultOperation`: CredentialVaultOperation
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CredentialVaultOperationsAPI.CreateCredentialVaultOperation`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateCredentialVaultOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialVaultOperation** | [**CredentialVaultOperation**](CredentialVaultOperation.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**CredentialVaultOperation**](CredentialVaultOperation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteCredentialVaultOperation

> DeleteCredentialVaultOperation(id).ConfirmDisabledObjects(confirmDisabledObjects).Aid(aid).Execute()

Delete Credential Vault operation



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
	confirmDisabledObjects := true // bool | Confirmation to disable affected objects (for example, tests) for credential-vault operations. (default to false)
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.CredentialVaultOperationsAPIService)(&apiClient.Common)

	r, err := api.DeleteCredentialVaultOperation(id).ConfirmDisabledObjects(confirmDisabledObjects).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultOperationsAPI.DeleteCredentialVaultOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The operation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteCredentialVaultOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirmDisabledObjects** | **bool** | Confirmation to disable affected objects (for example, tests) for credential-vault operations. | [default to false]
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetCredentialVaultOperation

> CredentialVaultOperation GetCredentialVaultOperation(id).Aid(aid).Execute()

Get Credential Vault operation



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.CredentialVaultOperationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetCredentialVaultOperation(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultOperationsAPI.GetCredentialVaultOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialVaultOperation`: CredentialVaultOperation
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CredentialVaultOperationsAPI.GetCredentialVaultOperation`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The operation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetCredentialVaultOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**CredentialVaultOperation**](CredentialVaultOperation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetCredentialVaultOperations

> CredentialVaultOperations GetCredentialVaultOperations().Aid(aid).Execute()

List Credential Vault operations



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

	api := (*connectors.CredentialVaultOperationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetCredentialVaultOperations().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultOperationsAPI.GetCredentialVaultOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialVaultOperations`: CredentialVaultOperations
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CredentialVaultOperationsAPI.GetCredentialVaultOperations`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetCredentialVaultOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**CredentialVaultOperations**](CredentialVaultOperations.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateCredentialVaultOperation

> CredentialVaultOperation UpdateCredentialVaultOperation(id).CredentialVaultOperation(credentialVaultOperation).Aid(aid).Execute()

Update Credential Vault operation



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
	credentialVaultOperation := *connectors.NewCredentialVaultOperation("My operation", []connectors.CredentialVaultSecret{*connectors.NewCredentialVaultSecret("secret_name", "secret/key")}) // CredentialVaultOperation | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*connectors.CredentialVaultOperationsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateCredentialVaultOperation(id).CredentialVaultOperation(credentialVaultOperation).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultOperationsAPI.UpdateCredentialVaultOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredentialVaultOperation`: CredentialVaultOperation
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CredentialVaultOperationsAPI.UpdateCredentialVaultOperation`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The operation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateCredentialVaultOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialVaultOperation** | [**CredentialVaultOperation**](CredentialVaultOperation.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**CredentialVaultOperation**](CredentialVaultOperation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


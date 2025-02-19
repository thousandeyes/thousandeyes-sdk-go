# CredentialsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredential**](CredentialsAPI.md#CreateCredential) | **Post** /credentials | Create credential
[**DeleteCredential**](CredentialsAPI.md#DeleteCredential) | **Delete** /credentials/{id} | Delete credential
[**GetCredential**](CredentialsAPI.md#GetCredential) | **Get** /credentials/{id} | Retrieve credential
[**GetCredentials**](CredentialsAPI.md#GetCredentials) | **Get** /credentials | List credentials
[**UpdateCredential**](CredentialsAPI.md#UpdateCredential) | **Put** /credentials/{id} | Update credential



## CreateCredential

> CredentialWithoutValue CreateCredential().CredentialRequest(credentialRequest).Aid(aid).Execute()

Create credential



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/credentials"
)

func main() {
	credentialRequest := *credentials.NewCredentialRequest() // CredentialRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*credentials.CredentialsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateCredential().CredentialRequest(credentialRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.CreateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCredential`: CredentialWithoutValue
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.CreateCredential`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialRequest** | [**CredentialRequest**](CredentialRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**CredentialWithoutValue**](CredentialWithoutValue.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteCredential

> DeleteCredential(id).Aid(aid).Execute()

Delete credential



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/credentials"
)

func main() {
	id := "3247" // string | The ID of the desired credential.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*credentials.CredentialsAPIService)(&apiClient.Common)

	r, err := api.DeleteCredential(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.DeleteCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The ID of the desired credential. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteCredentialRequest struct via the builder pattern


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


## GetCredential

> Credential GetCredential(id).Aid(aid).Execute()

Retrieve credential



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/credentials"
)

func main() {
	id := "3247" // string | The ID of the desired credential.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*credentials.CredentialsAPIService)(&apiClient.Common)

	resp, r, err := api.GetCredential(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.GetCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredential`: Credential
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.GetCredential`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The ID of the desired credential. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Credential**](Credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetCredentials

> Credentials GetCredentials().Aid(aid).Execute()

List credentials



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/credentials"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*credentials.CredentialsAPIService)(&apiClient.Common)

	resp, r, err := api.GetCredentials().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.GetCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentials`: Credentials
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.GetCredentials`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Credentials**](Credentials.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateCredential

> CredentialWithoutValue UpdateCredential(id).CredentialRequest(credentialRequest).Aid(aid).Execute()

Update credential



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/credentials"
)

func main() {
	id := "3247" // string | The ID of the desired credential.
	credentialRequest := *credentials.NewCredentialRequest() // CredentialRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*credentials.CredentialsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateCredential(id).CredentialRequest(credentialRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.UpdateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredential`: CredentialWithoutValue
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.UpdateCredential`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | The ID of the desired credential. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialRequest** | [**CredentialRequest**](CredentialRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**CredentialWithoutValue**](CredentialWithoutValue.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


# AccountGroupsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountGroup**](AccountGroupsAPI.md#CreateAccountGroup) | **Post** /account-groups | Create account group
[**DeleteAccountGroup**](AccountGroupsAPI.md#DeleteAccountGroup) | **Delete** /account-groups/{id} | Delete account group
[**GetAccountGroup**](AccountGroupsAPI.md#GetAccountGroup) | **Get** /account-groups/{id} | Retrieve account group
[**GetAccountGroups**](AccountGroupsAPI.md#GetAccountGroups) | **Get** /account-groups | List account groups
[**UpdateAccountGroup**](AccountGroupsAPI.md#UpdateAccountGroup) | **Put** /account-groups/{id} | Update account group



## CreateAccountGroup

> CreatedAccountGroup CreateAccountGroup().AccountGroupRequest(accountGroupRequest).Expand(expand).Execute()

Create account group



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
)

func main() {
	accountGroupRequest := *administrative.NewAccountGroupRequest("My testing account group") // AccountGroupRequest | 
	expand := []administrative.ExpandAccountGroupOptions{administrative.ExpandAccountGroupOptions("user")} // []ExpandAccountGroupOptions | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the `users` resource, pass the `?expand=user` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*administrative.AccountGroupsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAccountGroup().AccountGroupRequest(accountGroupRequest).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.CreateAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountGroup`: CreatedAccountGroup
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.CreateAccountGroup`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountGroupRequest** | [**AccountGroupRequest**](AccountGroupRequest.md) |  | 
 **expand** | [**[]ExpandAccountGroupOptions**](ExpandAccountGroupOptions.md) | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the &#x60;users&#x60; resource, pass the &#x60;?expand&#x3D;user&#x60; query. | 

### Return type

[**CreatedAccountGroup**](CreatedAccountGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteAccountGroup

> DeleteAccountGroup(id).Execute()

Delete account group



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
)

func main() {
	id := "1234" // string | Identifier for the account group.

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*administrative.AccountGroupsAPIService)(&apiClient.Common)

	r, err := api.DeleteAccountGroup(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.DeleteAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Identifier for the account group. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAccountGroup

> AccountGroupDetail GetAccountGroup(id).Expand(expand).Execute()

Retrieve account group



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
)

func main() {
	id := "1234" // string | Identifier for the account group.
	expand := []administrative.ExpandAccountGroupOptions{administrative.ExpandAccountGroupOptions("user")} // []ExpandAccountGroupOptions | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the `users` resource, pass the `?expand=user` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*administrative.AccountGroupsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAccountGroup(id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.GetAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountGroup`: AccountGroupDetail
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.GetAccountGroup`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Identifier for the account group. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | [**[]ExpandAccountGroupOptions**](ExpandAccountGroupOptions.md) | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the &#x60;users&#x60; resource, pass the &#x60;?expand&#x3D;user&#x60; query. | 

### Return type

[**AccountGroupDetail**](AccountGroupDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAccountGroups

> AccountGroups GetAccountGroups().Execute()

List account groups



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
)

func main() {

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*administrative.AccountGroupsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAccountGroups().Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.GetAccountGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountGroups`: AccountGroups
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.GetAccountGroups`: %v\n", string(json))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAccountGroupsRequest struct via the builder pattern


### Return type

[**AccountGroups**](AccountGroups.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAccountGroup

> AccountGroupDetail UpdateAccountGroup(id).AccountGroupRequest(accountGroupRequest).Expand(expand).Execute()

Update account group



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
)

func main() {
	id := "1234" // string | Identifier for the account group.
	accountGroupRequest := *administrative.NewAccountGroupRequest("My testing account group") // AccountGroupRequest | 
	expand := []administrative.ExpandAccountGroupOptions{administrative.ExpandAccountGroupOptions("user")} // []ExpandAccountGroupOptions | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the `users` resource, pass the `?expand=user` query. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*administrative.AccountGroupsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAccountGroup(id).AccountGroupRequest(accountGroupRequest).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.UpdateAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountGroup`: AccountGroupDetail
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.UpdateAccountGroup`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Identifier for the account group. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountGroupRequest** | [**AccountGroupRequest**](AccountGroupRequest.md) |  | 
 **expand** | [**[]ExpandAccountGroupOptions**](ExpandAccountGroupOptions.md) | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the &#x60;users&#x60; resource, pass the &#x60;?expand&#x3D;user&#x60; query. | 

### Return type

[**AccountGroupDetail**](AccountGroupDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


# InternetInsightsCatalogProvidersAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterCatalogProviders**](InternetInsightsCatalogProvidersAPI.md#FilterCatalogProviders) | **Post** /internet-insights/catalog/providers/filter | List catalog providers
[**GetCatalogProvider**](InternetInsightsCatalogProvidersAPI.md#GetCatalogProvider) | **Get** /internet-insights/catalog/providers/{providerId} | Retrieve a catalog provider



## FilterCatalogProviders

> ApiCatalogProviderResponse FilterCatalogProviders().ApiCatalogProviderFilter(apiCatalogProviderFilter).Aid(aid).Execute()

List catalog providers



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/internetinsights"
)

func main() {
	apiCatalogProviderFilter := *internetinsights.NewApiCatalogProviderFilter() // ApiCatalogProviderFilter | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*internetinsights.InternetInsightsCatalogProvidersAPIService)(&apiClient.Common)

	resp, r, err := api.FilterCatalogProviders().ApiCatalogProviderFilter(apiCatalogProviderFilter).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetInsightsCatalogProvidersAPI.FilterCatalogProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterCatalogProviders`: ApiCatalogProviderResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `InternetInsightsCatalogProvidersAPI.FilterCatalogProviders`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiFilterCatalogProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCatalogProviderFilter** | [**ApiCatalogProviderFilter**](ApiCatalogProviderFilter.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiCatalogProviderResponse**](ApiCatalogProviderResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetCatalogProvider

> ApiCatalogProviderDetails GetCatalogProvider(providerId).Aid(aid).Execute()

Retrieve a catalog provider



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/internetinsights"
)

func main() {
	providerId := "85602a0a-54a7-4e97-946e-67492ef1fa26" // string | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*internetinsights.InternetInsightsCatalogProvidersAPIService)(&apiClient.Common)

	resp, r, err := api.GetCatalogProvider(providerId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetInsightsCatalogProvidersAPI.GetCatalogProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogProvider`: ApiCatalogProviderDetails
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `InternetInsightsCatalogProvidersAPI.GetCatalogProvider`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetCatalogProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiCatalogProviderDetails**](ApiCatalogProviderDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


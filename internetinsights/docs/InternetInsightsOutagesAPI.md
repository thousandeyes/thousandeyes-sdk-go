# InternetInsightsOutagesAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterOutages**](InternetInsightsOutagesAPI.md#FilterOutages) | **Post** /internet-insights/outages/filter | List network and application outages
[**GetAppOutage**](InternetInsightsOutagesAPI.md#GetAppOutage) | **Get** /internet-insights/outages/app/{outageId} | Retrieve application outage
[**GetNetworkOutage**](InternetInsightsOutagesAPI.md#GetNetworkOutage) | **Get** /internet-insights/outages/net/{outageId} | Retrieve network outage



## FilterOutages

> ApiOutagesResponse FilterOutages().ApiOutageFilter(apiOutageFilter).Aid(aid).Execute()

List network and application outages



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
	apiOutageFilter := *internetinsights.NewApiOutageFilter() // ApiOutageFilter | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*internetinsights.InternetInsightsOutagesAPIService)(&apiClient.Common)

	resp, r, err := api.FilterOutages().ApiOutageFilter(apiOutageFilter).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetInsightsOutagesAPI.FilterOutages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilterOutages`: ApiOutagesResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `InternetInsightsOutagesAPI.FilterOutages`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiFilterOutagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiOutageFilter** | [**ApiOutageFilter**](ApiOutageFilter.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiOutagesResponse**](ApiOutagesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAppOutage

> ApiApplicationOutageDetails GetAppOutage(outageId).Aid(aid).Execute()

Retrieve application outage



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
	outageId := "F73E24F17E4996923196826A208BB572508A8EB13BEE14B0" // string | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*internetinsights.InternetInsightsOutagesAPIService)(&apiClient.Common)

	resp, r, err := api.GetAppOutage(outageId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetInsightsOutagesAPI.GetAppOutage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppOutage`: ApiApplicationOutageDetails
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `InternetInsightsOutagesAPI.GetAppOutage`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**outageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAppOutageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiApplicationOutageDetails**](ApiApplicationOutageDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetNetworkOutage

> ApiNetworkOutageDetails GetNetworkOutage(outageId).Aid(aid).Execute()

Retrieve network outage



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
	outageId := "694D8656960F34F76489BCE5E9BCD58EC53027462740D75F" // string | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*internetinsights.InternetInsightsOutagesAPIService)(&apiClient.Common)

	resp, r, err := api.GetNetworkOutage(outageId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetInsightsOutagesAPI.GetNetworkOutage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkOutage`: ApiNetworkOutageDetails
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `InternetInsightsOutagesAPI.GetNetworkOutage`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**outageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetNetworkOutageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ApiNetworkOutageDetails**](ApiNetworkOutageDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


# CloudInsightsIntegrationsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAWSFlowLogsMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#CreateAWSFlowLogsMonitoringIntegration) | **Post** /cloud-insights/integration/aws/flow-logs | Create AWS flow logs monitoring integration
[**CreateAWSInventoryMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#CreateAWSInventoryMonitoringIntegration) | **Post** /cloud-insights/integration/aws/inventory | Create AWS inventory monitoring integration
[**CreateAzureFlowLogsMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#CreateAzureFlowLogsMonitoringIntegration) | **Post** /cloud-insights/integration/azure/flow-logs | Create Azure flow logs monitoring integration
[**CreateAzureInventoryMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#CreateAzureInventoryMonitoringIntegration) | **Post** /cloud-insights/integration/azure/inventory | Create Azure inventory monitoring integration
[**DeleteAwsMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#DeleteAwsMonitoringIntegration) | **Delete** /cloud-insights/integration/aws/{integrationId} | Delete AWS integration
[**DeleteAzureMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#DeleteAzureMonitoringIntegration) | **Delete** /cloud-insights/integration/azure/{integrationId} | Delete Azure integration
[**GetAWSFlowlogsMonitoringIntegrationPolicies**](CloudInsightsIntegrationsAPI.md#GetAWSFlowlogsMonitoringIntegrationPolicies) | **Get** /cloud-insights/integration/aws/flow-logs/policies | Get AWS flow logs monitoring IAM policies
[**GetAWSInventoryMonitoringIntegrationPolicies**](CloudInsightsIntegrationsAPI.md#GetAWSInventoryMonitoringIntegrationPolicies) | **Get** /cloud-insights/integration/aws/inventory/policies | Get AWS inventory monitoring IAM policies
[**GetAWSMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#GetAWSMonitoringIntegration) | **Get** /cloud-insights/integration/aws/{integrationId} | Get AWS integration
[**GetAllAWSMonitoringIntegrations**](CloudInsightsIntegrationsAPI.md#GetAllAWSMonitoringIntegrations) | **Get** /cloud-insights/integration/aws | List AWS integrations
[**GetAllAzureMonitoringIntegrations**](CloudInsightsIntegrationsAPI.md#GetAllAzureMonitoringIntegrations) | **Get** /cloud-insights/integration/azure | List Azure integrations
[**GetAzureMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#GetAzureMonitoringIntegration) | **Get** /cloud-insights/integration/azure/{integrationId} | Get Azure integration
[**UpdateAzureFlowLogsMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#UpdateAzureFlowLogsMonitoringIntegration) | **Put** /cloud-insights/integration/azure/flow-logs/{integrationId} | Update Azure flow logs monitoring integration
[**UpdateAzureInventoryMonitoringIntegration**](CloudInsightsIntegrationsAPI.md#UpdateAzureInventoryMonitoringIntegration) | **Put** /cloud-insights/integration/azure/inventory/{integrationId} | Update Azure inventory monitoring integration



## CreateAWSFlowLogsMonitoringIntegration

> AwsMonitoringIntegration CreateAWSFlowLogsMonitoringIntegration().AwsFlowLogsIntegrationRequest(awsFlowLogsIntegrationRequest).Aid(aid).Execute()

Create AWS flow logs monitoring integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	awsFlowLogsIntegrationRequest := *cloudinsightsintegrations.NewAwsFlowLogsIntegrationRequest("integration name", "arn:aws:iam::01234567890:role/aws-flow-logs-ro", []string{"SnsTopicsArns_example"}) // AwsFlowLogsIntegrationRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAWSFlowLogsMonitoringIntegration().AwsFlowLogsIntegrationRequest(awsFlowLogsIntegrationRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.CreateAWSFlowLogsMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAWSFlowLogsMonitoringIntegration`: AwsMonitoringIntegration
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.CreateAWSFlowLogsMonitoringIntegration`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAWSFlowLogsMonitoringIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsFlowLogsIntegrationRequest** | [**AwsFlowLogsIntegrationRequest**](AwsFlowLogsIntegrationRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AwsMonitoringIntegration**](AwsMonitoringIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## CreateAWSInventoryMonitoringIntegration

> AwsMonitoringIntegration CreateAWSInventoryMonitoringIntegration().AwsInventoryIntegrationRequest(awsInventoryIntegrationRequest).Aid(aid).Execute()

Create AWS inventory monitoring integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	awsInventoryIntegrationRequest := *cloudinsightsintegrations.NewAwsInventoryIntegrationRequest("integration name", "arn:aws:iam::01234567890:role/aws-inventory-ro") // AwsInventoryIntegrationRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAWSInventoryMonitoringIntegration().AwsInventoryIntegrationRequest(awsInventoryIntegrationRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.CreateAWSInventoryMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAWSInventoryMonitoringIntegration`: AwsMonitoringIntegration
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.CreateAWSInventoryMonitoringIntegration`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAWSInventoryMonitoringIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsInventoryIntegrationRequest** | [**AwsInventoryIntegrationRequest**](AwsInventoryIntegrationRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AwsMonitoringIntegration**](AwsMonitoringIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## CreateAzureFlowLogsMonitoringIntegration

> AzureMonitoringIntegration CreateAzureFlowLogsMonitoringIntegration().AzureFlowLogsIntegrationRequest(azureFlowLogsIntegrationRequest).Aid(aid).Execute()

Create Azure flow logs monitoring integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	azureFlowLogsIntegrationRequest := *cloudinsightsintegrations.NewAzureFlowLogsIntegrationRequest("integration name", "2f9b2a1c-6d87-4c92-9d51-37efc93c0a4f", "p9z4Q~U7LkdX1yZr3TcB6fJqG8aM0sDeT5R9hVnKw", "e3a72c9b-42b1-4e0b-9b9f-7f6a3b2a1c44", "https://your-service-bus-namespace.servicebus.windows.net/your-queue-name") // AzureFlowLogsIntegrationRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAzureFlowLogsMonitoringIntegration().AzureFlowLogsIntegrationRequest(azureFlowLogsIntegrationRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.CreateAzureFlowLogsMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAzureFlowLogsMonitoringIntegration`: AzureMonitoringIntegration
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.CreateAzureFlowLogsMonitoringIntegration`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAzureFlowLogsMonitoringIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureFlowLogsIntegrationRequest** | [**AzureFlowLogsIntegrationRequest**](AzureFlowLogsIntegrationRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AzureMonitoringIntegration**](AzureMonitoringIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## CreateAzureInventoryMonitoringIntegration

> AzureMonitoringIntegration CreateAzureInventoryMonitoringIntegration().AzureInventoryIntegrationRequest(azureInventoryIntegrationRequest).Aid(aid).Execute()

Create Azure inventory monitoring integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	azureInventoryIntegrationRequest := *cloudinsightsintegrations.NewAzureInventoryIntegrationRequest("integration name", "2f9b2a1c-6d87-4c92-9d51-37efc93c0a4f", "p9z4Q~U7LkdX1yZr3TcB6fJqG8aM0sDeT5R9hVnKw", "e3a72c9b-42b1-4e0b-9b9f-7f6a3b2a1c44") // AzureInventoryIntegrationRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAzureInventoryMonitoringIntegration().AzureInventoryIntegrationRequest(azureInventoryIntegrationRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.CreateAzureInventoryMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAzureInventoryMonitoringIntegration`: AzureMonitoringIntegration
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.CreateAzureInventoryMonitoringIntegration`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAzureInventoryMonitoringIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureInventoryIntegrationRequest** | [**AzureInventoryIntegrationRequest**](AzureInventoryIntegrationRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AzureMonitoringIntegration**](AzureMonitoringIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteAwsMonitoringIntegration

> DeleteAwsMonitoringIntegration(integrationId).Aid(aid).Execute()

Delete AWS integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	integrationId := "e9c3bf02-a48c-4aa8-9e5f-898800d6f569" // string | The unique ID of the AWS or Azure inventory or flow logs monitoring integration.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	r, err := api.DeleteAwsMonitoringIntegration(integrationId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.DeleteAwsMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**integrationId** | **string** | The unique ID of the AWS or Azure inventory or flow logs monitoring integration. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteAwsMonitoringIntegrationRequest struct via the builder pattern


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


## DeleteAzureMonitoringIntegration

> DeleteAzureMonitoringIntegration(integrationId).Aid(aid).Execute()

Delete Azure integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	integrationId := "e9c3bf02-a48c-4aa8-9e5f-898800d6f569" // string | The unique ID of the AWS or Azure inventory or flow logs monitoring integration.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	r, err := api.DeleteAzureMonitoringIntegration(integrationId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.DeleteAzureMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**integrationId** | **string** | The unique ID of the AWS or Azure inventory or flow logs monitoring integration. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteAzureMonitoringIntegrationRequest struct via the builder pattern


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


## GetAWSFlowlogsMonitoringIntegrationPolicies

> string GetAWSFlowlogsMonitoringIntegrationPolicies().Aid(aid).Execute()

Get AWS flow logs monitoring IAM policies



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAWSFlowlogsMonitoringIntegrationPolicies().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.GetAWSFlowlogsMonitoringIntegrationPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAWSFlowlogsMonitoringIntegrationPolicies`: string
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.GetAWSFlowlogsMonitoringIntegrationPolicies`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAWSFlowlogsMonitoringIntegrationPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAWSInventoryMonitoringIntegrationPolicies

> string GetAWSInventoryMonitoringIntegrationPolicies().Aid(aid).Execute()

Get AWS inventory monitoring IAM policies



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAWSInventoryMonitoringIntegrationPolicies().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.GetAWSInventoryMonitoringIntegrationPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAWSInventoryMonitoringIntegrationPolicies`: string
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.GetAWSInventoryMonitoringIntegrationPolicies`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAWSInventoryMonitoringIntegrationPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAWSMonitoringIntegration

> AwsMonitoringIntegration GetAWSMonitoringIntegration(integrationId).Aid(aid).Execute()

Get AWS integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	integrationId := "e9c3bf02-a48c-4aa8-9e5f-898800d6f569" // string | The unique ID of the AWS or Azure inventory or flow logs monitoring integration.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAWSMonitoringIntegration(integrationId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.GetAWSMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAWSMonitoringIntegration`: AwsMonitoringIntegration
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.GetAWSMonitoringIntegration`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**integrationId** | **string** | The unique ID of the AWS or Azure inventory or flow logs monitoring integration. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAWSMonitoringIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AwsMonitoringIntegration**](AwsMonitoringIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAllAWSMonitoringIntegrations

> AwsMonitoringIntegrations GetAllAWSMonitoringIntegrations().Aid(aid).Execute()

List AWS integrations



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAllAWSMonitoringIntegrations().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.GetAllAWSMonitoringIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAWSMonitoringIntegrations`: AwsMonitoringIntegrations
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.GetAllAWSMonitoringIntegrations`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAllAWSMonitoringIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AwsMonitoringIntegrations**](AwsMonitoringIntegrations.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAllAzureMonitoringIntegrations

> AzureMonitoringIntegrations GetAllAzureMonitoringIntegrations().Aid(aid).Execute()

List Azure integrations



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAllAzureMonitoringIntegrations().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.GetAllAzureMonitoringIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAzureMonitoringIntegrations`: AzureMonitoringIntegrations
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.GetAllAzureMonitoringIntegrations`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAllAzureMonitoringIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AzureMonitoringIntegrations**](AzureMonitoringIntegrations.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAzureMonitoringIntegration

> AzureMonitoringIntegration GetAzureMonitoringIntegration(integrationId).Aid(aid).Execute()

Get Azure integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	integrationId := "e9c3bf02-a48c-4aa8-9e5f-898800d6f569" // string | The unique ID of the AWS or Azure inventory or flow logs monitoring integration.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAzureMonitoringIntegration(integrationId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.GetAzureMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAzureMonitoringIntegration`: AzureMonitoringIntegration
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.GetAzureMonitoringIntegration`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**integrationId** | **string** | The unique ID of the AWS or Azure inventory or flow logs monitoring integration. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAzureMonitoringIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AzureMonitoringIntegration**](AzureMonitoringIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAzureFlowLogsMonitoringIntegration

> AzureMonitoringIntegration UpdateAzureFlowLogsMonitoringIntegration(integrationId).AzureFlowLogsIntegrationRequest(azureFlowLogsIntegrationRequest).Aid(aid).Execute()

Update Azure flow logs monitoring integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	integrationId := "e9c3bf02-a48c-4aa8-9e5f-898800d6f569" // string | The unique ID of the AWS or Azure inventory or flow logs monitoring integration.
	azureFlowLogsIntegrationRequest := *cloudinsightsintegrations.NewAzureFlowLogsIntegrationRequest("integration name", "2f9b2a1c-6d87-4c92-9d51-37efc93c0a4f", "p9z4Q~U7LkdX1yZr3TcB6fJqG8aM0sDeT5R9hVnKw", "e3a72c9b-42b1-4e0b-9b9f-7f6a3b2a1c44", "https://your-service-bus-namespace.servicebus.windows.net/your-queue-name") // AzureFlowLogsIntegrationRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAzureFlowLogsMonitoringIntegration(integrationId).AzureFlowLogsIntegrationRequest(azureFlowLogsIntegrationRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.UpdateAzureFlowLogsMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAzureFlowLogsMonitoringIntegration`: AzureMonitoringIntegration
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.UpdateAzureFlowLogsMonitoringIntegration`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**integrationId** | **string** | The unique ID of the AWS or Azure inventory or flow logs monitoring integration. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAzureFlowLogsMonitoringIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureFlowLogsIntegrationRequest** | [**AzureFlowLogsIntegrationRequest**](AzureFlowLogsIntegrationRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AzureMonitoringIntegration**](AzureMonitoringIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAzureInventoryMonitoringIntegration

> AzureMonitoringIntegration UpdateAzureInventoryMonitoringIntegration(integrationId).AzureInventoryIntegrationRequest(azureInventoryIntegrationRequest).Aid(aid).Execute()

Update Azure inventory monitoring integration



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/cloudinsightsintegrations"
)

func main() {
	integrationId := "e9c3bf02-a48c-4aa8-9e5f-898800d6f569" // string | The unique ID of the AWS or Azure inventory or flow logs monitoring integration.
	azureInventoryIntegrationRequest := *cloudinsightsintegrations.NewAzureInventoryIntegrationRequest("integration name", "2f9b2a1c-6d87-4c92-9d51-37efc93c0a4f", "p9z4Q~U7LkdX1yZr3TcB6fJqG8aM0sDeT5R9hVnKw", "e3a72c9b-42b1-4e0b-9b9f-7f6a3b2a1c44") // AzureInventoryIntegrationRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAzureInventoryMonitoringIntegration(integrationId).AzureInventoryIntegrationRequest(azureInventoryIntegrationRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationsAPI.UpdateAzureInventoryMonitoringIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAzureInventoryMonitoringIntegration`: AzureMonitoringIntegration
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationsAPI.UpdateAzureInventoryMonitoringIntegration`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**integrationId** | **string** | The unique ID of the AWS or Azure inventory or flow logs monitoring integration. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAzureInventoryMonitoringIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureInventoryIntegrationRequest** | [**AzureInventoryIntegrationRequest**](AzureInventoryIntegrationRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AzureMonitoringIntegration**](AzureMonitoringIntegration.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


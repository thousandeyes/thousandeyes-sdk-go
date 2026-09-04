# CloudInsightsIntegrationPolicySettingsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAWSIntegrationPolicySettings**](CloudInsightsIntegrationPolicySettingsAPI.md#GetAWSIntegrationPolicySettings) | **Get** /cloud-insights/integration/aws/policy/settings | Get AWS integration policy settings
[**GetAzureIntegrationPolicySettings**](CloudInsightsIntegrationPolicySettingsAPI.md#GetAzureIntegrationPolicySettings) | **Get** /cloud-insights/integration/azure/policy/settings | Get Azure integration policy settings
[**UpdateAWSIntegrationPolicySettings**](CloudInsightsIntegrationPolicySettingsAPI.md#UpdateAWSIntegrationPolicySettings) | **Put** /cloud-insights/integration/aws/policy/settings | Update AWS integration policy settings
[**UpdateAzureIntegrationPolicySettings**](CloudInsightsIntegrationPolicySettingsAPI.md#UpdateAzureIntegrationPolicySettings) | **Put** /cloud-insights/integration/azure/policy/settings | Update Azure integration policy settings



## GetAWSIntegrationPolicySettings

> AwsIntegrationPolicySetting GetAWSIntegrationPolicySettings().Aid(aid).Execute()

Get AWS integration policy settings



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

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationPolicySettingsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAWSIntegrationPolicySettings().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationPolicySettingsAPI.GetAWSIntegrationPolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAWSIntegrationPolicySettings`: AwsIntegrationPolicySetting
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationPolicySettingsAPI.GetAWSIntegrationPolicySettings`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAWSIntegrationPolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AwsIntegrationPolicySetting**](AwsIntegrationPolicySetting.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAzureIntegrationPolicySettings

> AzureIntegrationPolicySetting GetAzureIntegrationPolicySettings().Aid(aid).Execute()

Get Azure integration policy settings



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

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationPolicySettingsAPIService)(&apiClient.Common)

	resp, r, err := api.GetAzureIntegrationPolicySettings().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationPolicySettingsAPI.GetAzureIntegrationPolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAzureIntegrationPolicySettings`: AzureIntegrationPolicySetting
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationPolicySettingsAPI.GetAzureIntegrationPolicySettings`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAzureIntegrationPolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AzureIntegrationPolicySetting**](AzureIntegrationPolicySetting.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAWSIntegrationPolicySettings

> AwsIntegrationPolicySetting UpdateAWSIntegrationPolicySettings().AwsIntegrationPolicySetting(awsIntegrationPolicySetting).Aid(aid).Execute()

Update AWS integration policy settings



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
	awsIntegrationPolicySetting := *cloudinsightsintegrations.NewAwsIntegrationPolicySetting([]cloudinsightsintegrations.AwsResourceGroupType{cloudinsightsintegrations.AwsResourceGroupType("cloudfront")}, []cloudinsightsintegrations.AwsRegion{cloudinsightsintegrations.AwsRegion("us-east-1")}, true) // AwsIntegrationPolicySetting | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationPolicySettingsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAWSIntegrationPolicySettings().AwsIntegrationPolicySetting(awsIntegrationPolicySetting).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationPolicySettingsAPI.UpdateAWSIntegrationPolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAWSIntegrationPolicySettings`: AwsIntegrationPolicySetting
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationPolicySettingsAPI.UpdateAWSIntegrationPolicySettings`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAWSIntegrationPolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsIntegrationPolicySetting** | [**AwsIntegrationPolicySetting**](AwsIntegrationPolicySetting.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AwsIntegrationPolicySetting**](AwsIntegrationPolicySetting.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAzureIntegrationPolicySettings

> AzureIntegrationPolicySetting UpdateAzureIntegrationPolicySettings().AzureIntegrationPolicySetting(azureIntegrationPolicySetting).Aid(aid).Execute()

Update Azure integration policy settings



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
	azureIntegrationPolicySetting := *cloudinsightsintegrations.NewAzureIntegrationPolicySetting([]cloudinsightsintegrations.AzureResourceGroupType{cloudinsightsintegrations.AzureResourceGroupType("afd")}, *cloudinsightsintegrations.NewSubscriptionsPolicy([]cloudinsightsintegrations.SubscriptionsPolicyRule{*cloudinsightsintegrations.NewSubscriptionsPolicyRule(cloudinsightsintegrations.SubscriptionsPolicyRuleField("subscription-id"), "^prod-.*$", cloudinsightsintegrations.SubscriptionsPolicyRuleAction("include"))}, cloudinsightsintegrations.SubscriptionsPolicyRuleAction("include"))) // AzureIntegrationPolicySetting | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*cloudinsightsintegrations.CloudInsightsIntegrationPolicySettingsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAzureIntegrationPolicySettings().AzureIntegrationPolicySetting(azureIntegrationPolicySetting).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsIntegrationPolicySettingsAPI.UpdateAzureIntegrationPolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAzureIntegrationPolicySettings`: AzureIntegrationPolicySetting
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudInsightsIntegrationPolicySettingsAPI.UpdateAzureIntegrationPolicySettings`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAzureIntegrationPolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureIntegrationPolicySetting** | [**AzureIntegrationPolicySetting**](AzureIntegrationPolicySetting.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**AzureIntegrationPolicySetting**](AzureIntegrationPolicySetting.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


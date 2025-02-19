# CloudAndEnterpriseAgentNotificationRulesAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentsNotificationRule**](CloudAndEnterpriseAgentNotificationRulesAPI.md#GetAgentsNotificationRule) | **Get** /agents/notification-rules/{notificationRuleId} | Retrieve agent notification rule
[**GetAgentsNotificationRules**](CloudAndEnterpriseAgentNotificationRulesAPI.md#GetAgentsNotificationRules) | **Get** /agents/notification-rules | List agent notification rules



## GetAgentsNotificationRule

> NotificationRuleDetail GetAgentsNotificationRule(notificationRuleId).Aid(aid).Execute()

Retrieve agent notification rule



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/agents"
)

func main() {
	notificationRuleId := "281474976710706" // string | Unique ID for the agent notification rule.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.CloudAndEnterpriseAgentNotificationRulesAPIService)(&apiClient.Common)

	resp, r, err := api.GetAgentsNotificationRule(notificationRuleId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAndEnterpriseAgentNotificationRulesAPI.GetAgentsNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsNotificationRule`: NotificationRuleDetail
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudAndEnterpriseAgentNotificationRulesAPI.GetAgentsNotificationRule`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**notificationRuleId** | **string** | Unique ID for the agent notification rule. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAgentsNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**NotificationRuleDetail**](NotificationRuleDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAgentsNotificationRules

> ListNotificationRulesResponse GetAgentsNotificationRules().Aid(aid).Execute()

List agent notification rules



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/agents"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*agents.CloudAndEnterpriseAgentNotificationRulesAPIService)(&apiClient.Common)

	resp, r, err := api.GetAgentsNotificationRules().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAndEnterpriseAgentNotificationRulesAPI.GetAgentsNotificationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsNotificationRules`: ListNotificationRulesResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `CloudAndEnterpriseAgentNotificationRulesAPI.GetAgentsNotificationRules`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAgentsNotificationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**ListNotificationRulesResponse**](ListNotificationRulesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


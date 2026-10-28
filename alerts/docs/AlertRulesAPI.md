# AlertRulesAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertRule**](AlertRulesAPI.md#CreateAlertRule) | **Post** /alerts/rules | Create alert rule
[**DeleteAlertRule**](AlertRulesAPI.md#DeleteAlertRule) | **Delete** /alerts/rules/{ruleId} | Delete alert rule
[**GetAlertRule**](AlertRulesAPI.md#GetAlertRule) | **Get** /alerts/rules/{ruleId} | Retrieve alert rule
[**GetAlertsRules**](AlertRulesAPI.md#GetAlertsRules) | **Get** /alerts/rules | List alert rules
[**UpdateAlertRule**](AlertRulesAPI.md#UpdateAlertRule) | **Put** /alerts/rules/{ruleId} | Update alert rule



## CreateAlertRule

> Rule CreateAlertRule().RuleDetailUpdate(ruleDetailUpdate).Aid(aid).Execute()

Create alert rule



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	ruleDetailUpdate := *alerts.NewRuleDetailUpdate("The End of the Internet", "((hops((hopDelay >= 100 ms))))", alerts.AlertType("page-load"), int32(5)) // RuleDetailUpdate | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertRulesAPIService)(&apiClient.Common)

	resp, r, err := api.CreateAlertRule().RuleDetailUpdate(ruleDetailUpdate).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.CreateAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertRule`: Rule
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.CreateAlertRule`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleDetailUpdate** | [**RuleDetailUpdate**](RuleDetailUpdate.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteAlertRule

> DeleteAlertRule(ruleId).Aid(aid).Execute()

Delete alert rule



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	ruleId := "127094" // string | Unique alert rule ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertRulesAPIService)(&apiClient.Common)

	r, err := api.DeleteAlertRule(ruleId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.DeleteAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ruleId** | **string** | Unique alert rule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteAlertRuleRequest struct via the builder pattern


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


## GetAlertRule

> RuleDetail GetAlertRule(ruleId).Aid(aid).Execute()

Retrieve alert rule



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	ruleId := "127094" // string | Unique alert rule ID.
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertRulesAPIService)(&apiClient.Common)

	resp, r, err := api.GetAlertRule(ruleId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.GetAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertRule`: RuleDetail
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.GetAlertRule`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ruleId** | **string** | Unique alert rule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**RuleDetail**](RuleDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetAlertsRules

> Rules GetAlertsRules().Aid(aid).HasOptimization(hasOptimization).Execute()

List alert rules



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	hasOptimization := true // bool | Filters alert rules by whether they currently have an active optimization. `true` returns only rules with an active optimization; every successfully returned row contains `hasOptimization: true`. `false` returns only rules without an active optimization; every successfully returned row contains `hasOptimization: false`. When omitted, the list is unfiltered. Rows whose lookup succeeds include `hasOptimization` `true` or `false`; the field is omitted when enrichment is unavailable. When this parameter is explicitly supplied and optimization data cannot be retrieved, the request returns `503` with `OPTIMIZATION_DATA_UNAVAILABLE`. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertRulesAPIService)(&apiClient.Common)

	resp, r, err := api.GetAlertsRules().Aid(aid).HasOptimization(hasOptimization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.GetAlertsRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertsRules`: Rules
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.GetAlertsRules`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetAlertsRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **hasOptimization** | **bool** | Filters alert rules by whether they currently have an active optimization. &#x60;true&#x60; returns only rules with an active optimization; every successfully returned row contains &#x60;hasOptimization: true&#x60;. &#x60;false&#x60; returns only rules without an active optimization; every successfully returned row contains &#x60;hasOptimization: false&#x60;. When omitted, the list is unfiltered. Rows whose lookup succeeds include &#x60;hasOptimization&#x60; &#x60;true&#x60; or &#x60;false&#x60;; the field is omitted when enrichment is unavailable. When this parameter is explicitly supplied and optimization data cannot be retrieved, the request returns &#x60;503&#x60; with &#x60;OPTIMIZATION_DATA_UNAVAILABLE&#x60;. | 

### Return type

[**Rules**](Rules.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateAlertRule

> Rule UpdateAlertRule(ruleId).RuleDetailUpdate(ruleDetailUpdate).Aid(aid).Execute()

Update alert rule



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/alerts"
)

func main() {
	ruleId := "127094" // string | Unique alert rule ID.
	ruleDetailUpdate := *alerts.NewRuleDetailUpdate("The End of the Internet", "((hops((hopDelay >= 100 ms))))", alerts.AlertType("page-load"), int32(5)) // RuleDetailUpdate | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*alerts.AlertRulesAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateAlertRule(ruleId).RuleDetailUpdate(ruleDetailUpdate).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.UpdateAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertRule`: Rule
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.UpdateAlertRule`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ruleId** | **string** | Unique alert rule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleDetailUpdate** | [**RuleDetailUpdate**](RuleDetailUpdate.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


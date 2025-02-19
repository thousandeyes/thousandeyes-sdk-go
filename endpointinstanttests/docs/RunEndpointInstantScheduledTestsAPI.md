# RunEndpointInstantScheduledTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunEndpointScheduledInstantTest**](RunEndpointInstantScheduledTestsAPI.md#RunEndpointScheduledInstantTest) | **Post** /endpoint/tests/scheduled-tests/{testId}/run | Run endpoint instant scheduled test



## RunEndpointScheduledInstantTest

> EndpointRunScheduledInstantTestResult RunEndpointScheduledInstantTest(testId).Aid(aid).Execute()

Run endpoint instant scheduled test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointinstanttests"
)

func main() {
	testId := "765231567" // string | ID of the endpoint instant scheduled test to rerun
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointinstanttests.RunEndpointInstantScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.RunEndpointScheduledInstantTest(testId).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunEndpointInstantScheduledTestsAPI.RunEndpointScheduledInstantTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunEndpointScheduledInstantTest`: EndpointRunScheduledInstantTestResult
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `RunEndpointInstantScheduledTestsAPI.RunEndpointScheduledInstantTest`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | ID of the endpoint instant scheduled test to rerun | 

### Other Parameters

Other parameters are passed through a pointer to a ApiRunEndpointScheduledInstantTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointRunScheduledInstantTestResult**](EndpointRunScheduledInstantTestResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


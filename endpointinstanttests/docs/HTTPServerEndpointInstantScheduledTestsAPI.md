# HTTPServerEndpointInstantScheduledTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHttpServerScheduledInstantTest**](HTTPServerEndpointInstantScheduledTestsAPI.md#CreateHttpServerScheduledInstantTest) | **Post** /endpoint/tests/scheduled-tests/http-server/instant | Run http server instant scheduled test



## CreateHttpServerScheduledInstantTest

> EndpointHttpServerTest CreateHttpServerScheduledInstantTest().EndpointHttpServerInstantTest(endpointHttpServerInstantTest).Aid(aid).Execute()

Run http server instant scheduled test



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
	endpointHttpServerInstantTest := *endpointinstanttests.NewEndpointHttpServerInstantTest("Test name", "https://example.com:443") // EndpointHttpServerInstantTest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*endpointinstanttests.HTTPServerEndpointInstantScheduledTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateHttpServerScheduledInstantTest().EndpointHttpServerInstantTest(endpointHttpServerInstantTest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPServerEndpointInstantScheduledTestsAPI.CreateHttpServerScheduledInstantTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHttpServerScheduledInstantTest`: EndpointHttpServerTest
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `HTTPServerEndpointInstantScheduledTestsAPI.CreateHttpServerScheduledInstantTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateHttpServerScheduledInstantTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointHttpServerInstantTest** | [**EndpointHttpServerInstantTest**](EndpointHttpServerInstantTest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EndpointHttpServerTest**](EndpointHttpServerTest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


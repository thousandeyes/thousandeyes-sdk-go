# SIPServerInstantTestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipServerInstantTest**](SIPServerInstantTestsAPI.md#CreateSipServerInstantTest) | **Post** /tests/sip-server/instant | Create SIP server instant test



## CreateSipServerInstantTest

> SipServerInstantTestResponse CreateSipServerInstantTest().SipServerInstantTestRequest(sipServerInstantTestRequest).Aid(aid).Expand(expand).Execute()

Create SIP server instant test



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/instanttests"
)

func main() {
	sipServerInstantTestRequest := *instanttests.NewSipServerInstantTestRequest([]instanttests.TestAgent{*instanttests.NewTestAgent()}, *instanttests.NewTestSipCredentials(int32(49153))) // SipServerInstantTestRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []instanttests.ExpandInstantTestOptions{instanttests.ExpandInstantTestOptions("agent")} // []ExpandInstantTestOptions | (Optional) Indicates if the test sub-resources should be expanded. Defaults to no expansion. To expand the `agents` sub-resource, use the query `?expand=agent`. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*instanttests.SIPServerInstantTestsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateSipServerInstantTest().SipServerInstantTestRequest(sipServerInstantTestRequest).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIPServerInstantTestsAPI.CreateSipServerInstantTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSipServerInstantTest`: SipServerInstantTestResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `SIPServerInstantTestsAPI.CreateSipServerInstantTest`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateSipServerInstantTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sipServerInstantTestRequest** | [**SipServerInstantTestRequest**](SipServerInstantTestRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandInstantTestOptions**](ExpandInstantTestOptions.md) | (Optional) Indicates if the test sub-resources should be expanded. Defaults to no expansion. To expand the &#x60;agents&#x60; sub-resource, use the query &#x60;?expand&#x3D;agent&#x60;. | 

### Return type

[**SipServerInstantTestResponse**](SipServerInstantTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


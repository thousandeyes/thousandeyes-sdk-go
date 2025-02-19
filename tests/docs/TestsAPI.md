# TestsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTests**](TestsAPI.md#GetTests) | **Get** /tests | List configured tests



## GetTests

> Tests GetTests().Aid(aid).Execute()

List configured tests



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tests.TestsAPIService)(&apiClient.Common)

	resp, r, err := api.GetTests().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.GetTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTests`: Tests
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.GetTests`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**Tests**](Tests.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


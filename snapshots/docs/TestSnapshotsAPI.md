# TestSnapshotsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTestSnapshot**](TestSnapshotsAPI.md#CreateTestSnapshot) | **Post** /tests/{testId}/snapshot | Create test snapshot



## CreateTestSnapshot

> SnapshotResponse CreateTestSnapshot(testId).SnapshotRequest(snapshotRequest).Aid(aid).Execute()

Create test snapshot



### Example

```go
package main

import (
	"fmt"
	"os"
    "time"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/snapshots"
)

func main() {
	testId := "202701" // string | Test ID
	snapshotRequest := *snapshots.NewSnapshotRequest("Snapshot created through API", time.Now(), time.Now()) // SnapshotRequest | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*snapshots.TestSnapshotsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateTestSnapshot(testId).SnapshotRequest(snapshotRequest).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSnapshotsAPI.CreateTestSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTestSnapshot`: SnapshotResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TestSnapshotsAPI.CreateTestSnapshot`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**testId** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiCreateTestSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotRequest** | [**SnapshotRequest**](SnapshotRequest.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**SnapshotResponse**](SnapshotResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


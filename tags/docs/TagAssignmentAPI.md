# TagAssignmentAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignTag**](TagAssignmentAPI.md#AssignTag) | **Post** /tags/{id}/assign | Assign tag to multiple objects
[**AssignTags**](TagAssignmentAPI.md#AssignTags) | **Post** /tags/assign | Assign multiple tags to multiple objects
[**UnassignTag**](TagAssignmentAPI.md#UnassignTag) | **Post** /tags/{id}/unassign | Remove tag from multiple objects
[**UnassignTags**](TagAssignmentAPI.md#UnassignTags) | **Post** /tags/unassign | Remove multiple tags from multiple objects



## AssignTag

> BulkTagAssignment AssignTag(id).TagAssignment(tagAssignment).Aid(aid).Execute()

Assign tag to multiple objects



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tags"
)

func main() {
	id := "c6b78e57-81a2-4c5f-a11a-d96c3c664d55" // string | Tag ID
	tagAssignment := *tags.NewTagAssignment() // TagAssignment | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagAssignmentAPIService)(&apiClient.Common)

	resp, r, err := api.AssignTag(id).TagAssignment(tagAssignment).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAssignmentAPI.AssignTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignTag`: BulkTagAssignment
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TagAssignmentAPI.AssignTag`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Tag ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiAssignTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagAssignment** | [**TagAssignment**](TagAssignment.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**BulkTagAssignment**](BulkTagAssignment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## AssignTags

> BulkTagAssignments AssignTags().BulkTagAssignments(bulkTagAssignments).Aid(aid).Execute()

Assign multiple tags to multiple objects



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tags"
)

func main() {
	bulkTagAssignments := *tags.NewBulkTagAssignments() // BulkTagAssignments | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagAssignmentAPIService)(&apiClient.Common)

	resp, r, err := api.AssignTags().BulkTagAssignments(bulkTagAssignments).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAssignmentAPI.AssignTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignTags`: BulkTagAssignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TagAssignmentAPI.AssignTags`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiAssignTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkTagAssignments** | [**BulkTagAssignments**](BulkTagAssignments.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**BulkTagAssignments**](BulkTagAssignments.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UnassignTag

> UnassignTag(id).TagAssignment(tagAssignment).Aid(aid).Execute()

Remove tag from multiple objects



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tags"
)

func main() {
	id := "c6b78e57-81a2-4c5f-a11a-d96c3c664d55" // string | Tag ID
	tagAssignment := *tags.NewTagAssignment() // TagAssignment | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagAssignmentAPIService)(&apiClient.Common)

	r, err := api.UnassignTag(id).TagAssignment(tagAssignment).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAssignmentAPI.UnassignTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Tag ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUnassignTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagAssignment** | [**TagAssignment**](TagAssignment.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UnassignTags

> BulkTagAssignments UnassignTags().BulkTagAssignments(bulkTagAssignments).Aid(aid).Execute()

Remove multiple tags from multiple objects



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tags"
)

func main() {
	bulkTagAssignments := *tags.NewBulkTagAssignments() // BulkTagAssignments | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagAssignmentAPIService)(&apiClient.Common)

	resp, r, err := api.UnassignTags().BulkTagAssignments(bulkTagAssignments).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAssignmentAPI.UnassignTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnassignTags`: BulkTagAssignments
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TagAssignmentAPI.UnassignTags`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiUnassignTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkTagAssignments** | [**BulkTagAssignments**](BulkTagAssignments.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**BulkTagAssignments**](BulkTagAssignments.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


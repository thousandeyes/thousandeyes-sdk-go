# TagsAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagsAPI.md#CreateTag) | **Post** /tags | Create tag
[**CreateTags**](TagsAPI.md#CreateTags) | **Post** /tags/bulk | Create multiple tags
[**DeleteTag**](TagsAPI.md#DeleteTag) | **Delete** /tags/{id} | Delete tag
[**GetTag**](TagsAPI.md#GetTag) | **Get** /tags/{id} | Retrieve tag
[**GetTags**](TagsAPI.md#GetTags) | **Get** /tags | List tags
[**UpdateTag**](TagsAPI.md#UpdateTag) | **Put** /tags/{id} | Update tag



## CreateTag

> TagInfo CreateTag().Aid(aid).TagInfo(tagInfo).Execute()

Create tag



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	tagInfo := *tags.NewTagInfo() // TagInfo | Tag resource (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateTag().Aid(aid).TagInfo(tagInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: TagInfo
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.CreateTag`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **tagInfo** | [**TagInfo**](TagInfo.md) | Tag resource | 

### Return type

[**TagInfo**](TagInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## CreateTags

> BulkTagResponse CreateTags().Aid(aid).BulkTagResponse(bulkTagResponse).Execute()

Create multiple tags



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	bulkTagResponse := *tags.NewBulkTagResponse() // BulkTagResponse | Tag resource (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagsAPIService)(&apiClient.Common)

	resp, r, err := api.CreateTags().Aid(aid).BulkTagResponse(bulkTagResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTags`: BulkTagResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.CreateTags`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **bulkTagResponse** | [**BulkTagResponse**](BulkTagResponse.md) | Tag resource | 

### Return type

[**BulkTagResponse**](BulkTagResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(id).Aid(aid).Execute()

Delete tag



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagsAPIService)(&apiClient.Common)

	r, err := api.DeleteTag(id).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Tag ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiDeleteTagRequest struct via the builder pattern


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


## GetTag

> Tag GetTag(id).Aid(aid).Expand(expand).Execute()

Retrieve tag



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tags.ExpandTagsOptions{tags.ExpandTagsOptions("assignments")} // []ExpandTagsOptions | Optional, to retrieve associated assignments (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagsAPIService)(&apiClient.Common)

	resp, r, err := api.GetTag(id).Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTag`: Tag
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTag`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Tag ID | 

### Other Parameters

Other parameters are passed through a pointer to a ApiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTagsOptions**](ExpandTagsOptions.md) | Optional, to retrieve associated assignments | 

### Return type

[**Tag**](Tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetTags

> Tags GetTags().Aid(aid).Expand(expand).Execute()

List tags



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
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	expand := []tags.ExpandTagsOptions{tags.ExpandTagsOptions("assignments")} // []ExpandTagsOptions | Optional, to retrieve associated assignments (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagsAPIService)(&apiClient.Common)

	resp, r, err := api.GetTags().Aid(aid).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: Tags
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTags`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **expand** | [**[]ExpandTagsOptions**](ExpandTagsOptions.md) | Optional, to retrieve associated assignments | 

### Return type

[**Tags**](Tags.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UpdateTag

> TagInfo UpdateTag(id).Aid(aid).TagInfo(tagInfo).Execute()

Update tag



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
	id := "c6b78e57-81a2-4c5f-a11a-d96c3c664d55" // string | ID of tag to update
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)
	tagInfo := *tags.NewTagInfo() // TagInfo |  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*tags.TagsAPIService)(&apiClient.Common)

	resp, r, err := api.UpdateTag(id).Aid(aid).TagInfo(tagInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.UpdateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTag`: TagInfo
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.UpdateTag`: %v\n", string(json))
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | ID of tag to update | 

### Other Parameters

Other parameters are passed through a pointer to a ApiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 
 **tagInfo** | [**TagInfo**](TagInfo.md) |  | 

### Return type

[**TagInfo**](TagInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


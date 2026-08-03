# QuotasAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignOrganizationsAccountGroupsQuotas**](QuotasAPI.md#AssignOrganizationsAccountGroupsQuotas) | **Post** /quotas/account-groups/assign | Create or update accout group quotas
[**AssignOrganizationsQuotas**](QuotasAPI.md#AssignOrganizationsQuotas) | **Post** /quotas/assign | Create or update organizations quotas
[**GetQuotas**](QuotasAPI.md#GetQuotas) | **Get** /quotas | Get organization and account group usage quota
[**UnassignOrganizationsAccountGroupsQuotas**](QuotasAPI.md#UnassignOrganizationsAccountGroupsQuotas) | **Post** /quotas/account-groups/unassign | Remove account group quotas from organizations
[**UnassignOrganizationsQuotas**](QuotasAPI.md#UnassignOrganizationsQuotas) | **Post** /quotas/unassign | Remove organization quotas



## AssignOrganizationsAccountGroupsQuotas

> OrganizationsQuotasAssign AssignOrganizationsAccountGroupsQuotas().OrganizationsQuotasAssign(organizationsQuotasAssign).Execute()

Create or update accout group quotas



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/usage"
)

func main() {
	organizationsQuotasAssign := *usage.NewOrganizationsQuotasAssign() // OrganizationsQuotasAssign |  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*usage.QuotasAPIService)(&apiClient.Common)

	resp, r, err := api.AssignOrganizationsAccountGroupsQuotas().OrganizationsQuotasAssign(organizationsQuotasAssign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.AssignOrganizationsAccountGroupsQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignOrganizationsAccountGroupsQuotas`: OrganizationsQuotasAssign
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `QuotasAPI.AssignOrganizationsAccountGroupsQuotas`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiAssignOrganizationsAccountGroupsQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationsQuotasAssign** | [**OrganizationsQuotasAssign**](OrganizationsQuotasAssign.md) |  | 

### Return type

[**OrganizationsQuotasAssign**](OrganizationsQuotasAssign.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## AssignOrganizationsQuotas

> QuotasAssignResponse AssignOrganizationsQuotas().QuotasAssignRequest(quotasAssignRequest).Execute()

Create or update organizations quotas



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/usage"
)

func main() {
	quotasAssignRequest := *usage.NewQuotasAssignRequest() // QuotasAssignRequest |  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*usage.QuotasAPIService)(&apiClient.Common)

	resp, r, err := api.AssignOrganizationsQuotas().QuotasAssignRequest(quotasAssignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.AssignOrganizationsQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignOrganizationsQuotas`: QuotasAssignResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `QuotasAPI.AssignOrganizationsQuotas`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiAssignOrganizationsQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotasAssignRequest** | [**QuotasAssignRequest**](QuotasAssignRequest.md) |  | 

### Return type

[**QuotasAssignResponse**](QuotasAssignResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetQuotas

> Quotas GetQuotas().Execute()

Get organization and account group usage quota



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/usage"
)

func main() {

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*usage.QuotasAPIService)(&apiClient.Common)

	resp, r, err := api.GetQuotas().Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.GetQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotas`: Quotas
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `QuotasAPI.GetQuotas`: %v\n", string(json))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a ApiGetQuotasRequest struct via the builder pattern


### Return type

[**Quotas**](Quotas.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UnassignOrganizationsAccountGroupsQuotas

> UnassignOrganizationsAccountGroupsQuotas().OrganizationsQuotasUnassign(organizationsQuotasUnassign).Execute()

Remove account group quotas from organizations



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/usage"
)

func main() {
	organizationsQuotasUnassign := *usage.NewOrganizationsQuotasUnassign() // OrganizationsQuotasUnassign |  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*usage.QuotasAPIService)(&apiClient.Common)

	r, err := api.UnassignOrganizationsAccountGroupsQuotas().OrganizationsQuotasUnassign(organizationsQuotasUnassign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.UnassignOrganizationsAccountGroupsQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiUnassignOrganizationsAccountGroupsQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationsQuotasUnassign** | [**OrganizationsQuotasUnassign**](OrganizationsQuotasUnassign.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## UnassignOrganizationsQuotas

> UnassignOrganizationsQuotas().QuotasUnassign(quotasUnassign).Execute()

Remove organization quotas



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/usage"
)

func main() {
	quotasUnassign := *usage.NewQuotasUnassign() // QuotasUnassign |  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*usage.QuotasAPIService)(&apiClient.Common)

	r, err := api.UnassignOrganizationsQuotas().QuotasUnassign(quotasUnassign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.UnassignOrganizationsQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiUnassignOrganizationsQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotasUnassign** | [**QuotasUnassign**](QuotasUnassign.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


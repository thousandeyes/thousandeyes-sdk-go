# EmulationAPI

All URIs are relative to *https://api.thousandeyes.com/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmulatedDevice**](EmulationAPI.md#CreateEmulatedDevice) | **Post** /emulated-devices | Create emulated device
[**GetEmulatedDevices**](EmulationAPI.md#GetEmulatedDevices) | **Get** /emulated-devices | List emulated devices
[**GetUserAgents**](EmulationAPI.md#GetUserAgents) | **Get** /user-agents | List user-agents



## CreateEmulatedDevice

> EmulatedDeviceResponse CreateEmulatedDevice().EmulatedDevice(emulatedDevice).Aid(aid).Execute()

Create emulated device



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/emulation"
)

func main() {
	emulatedDevice := *emulation.NewEmulatedDevice(emulation.EmulatedDeviceCategory("desktop"), int32(1024), int32(768)) // EmulatedDevice | 
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*emulation.EmulationAPIService)(&apiClient.Common)

	resp, r, err := api.CreateEmulatedDevice().EmulatedDevice(emulatedDevice).Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmulationAPI.CreateEmulatedDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmulatedDevice`: EmulatedDeviceResponse
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EmulationAPI.CreateEmulatedDevice`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiCreateEmulatedDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emulatedDevice** | [**EmulatedDevice**](EmulatedDevice.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**EmulatedDeviceResponse**](EmulatedDeviceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetEmulatedDevices

> EmulatedDeviceResponses GetEmulatedDevices().Expand(expand).Execute()

List emulated devices



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/emulation"
)

func main() {
	expand := []emulation.ExpandEmulatedDeviceOptions{emulation.ExpandEmulatedDeviceOptions("user-agent")} // []ExpandEmulatedDeviceOptions | Optional query parameter that controls whether user-agent templates are included in the response. By default, user-agent templates are not included. To include them, add `?expand=user-agent` to the request.  (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*emulation.EmulationAPIService)(&apiClient.Common)

	resp, r, err := api.GetEmulatedDevices().Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmulationAPI.GetEmulatedDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmulatedDevices`: EmulatedDeviceResponses
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EmulationAPI.GetEmulatedDevices`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetEmulatedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | [**[]ExpandEmulatedDeviceOptions**](ExpandEmulatedDeviceOptions.md) | Optional query parameter that controls whether user-agent templates are included in the response. By default, user-agent templates are not included. To include them, add &#x60;?expand&#x3D;user-agent&#x60; to the request.  | 

### Return type

[**EmulatedDeviceResponses**](EmulatedDeviceResponses.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


## GetUserAgents

> UserAgents GetUserAgents().Aid(aid).Execute()

List user-agents



### Example

```go
package main

import (
	"fmt"
	"os"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/emulation"
)

func main() {
	aid := "1234" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. (optional)

	configuration := client.NewConfiguration().WithAuthToken("<bearer-token>")
	apiClient := client.NewAPIClient(configuration)

	api := (*emulation.EmulationAPIService)(&apiClient.Common)

	resp, r, err := api.GetUserAgents().Aid(aid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmulationAPI.GetUserAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAgents`: UserAgents
	json, _ := resp.MarshalJSON()
	fmt.Fprintf(os.Stdout, "Response from `EmulationAPI.GetUserAgents`: %v\n", string(json))
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a ApiGetUserAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response. | 

### Return type

[**UserAgents**](UserAgents.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)[[Back to README]](../README.md)


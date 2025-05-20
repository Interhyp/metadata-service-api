# \RestApiV1ServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteService**](RestApiV1ServicesAPI.md#DeleteService) | **Delete** /rest/api/v1/services/{service} | delete the service with a given name
[**GetService**](RestApiV1ServicesAPI.md#GetService) | **Get** /rest/api/v1/services/{service} | get a single service by name
[**GetServicePromoters**](RestApiV1ServicesAPI.md#GetServicePromoters) | **Get** /rest/api/v1/services/{service}/promoters | get all users who may promote a service
[**GetServices**](RestApiV1ServicesAPI.md#GetServices) | **Get** /rest/api/v1/services | get services
[**PatchService**](RestApiV1ServicesAPI.md#PatchService) | **Patch** /rest/api/v1/services/{service} | patch an existing service with the given name
[**RegisterService**](RestApiV1ServicesAPI.md#RegisterService) | **Post** /rest/api/v1/services/{service} | register a new service with the given name
[**UpdateService**](RestApiV1ServicesAPI.md#UpdateService) | **Put** /rest/api/v1/services/{service} | update an existing service with the given name



## DeleteService

> DeleteService(ctx, service).DeletionDto(deletionDto).Execute()

delete the service with a given name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Interhyp/metadata-service-api"
)

func main() {
	service := "service_example" // string | The (globally unique) name of the service, must match `^[a-z](-?[a-z0-9]+)*$`.
	deletionDto := *openapiclient.NewDeletionDto("JiraIssue_example") // DeletionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RestApiV1ServicesAPI.DeleteService(context.Background(), service).DeletionDto(deletionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1ServicesAPI.DeleteService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | The (globally unique) name of the service, must match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deletionDto** | [**DeletionDto**](DeletionDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> ServiceDto GetService(ctx, service).Execute()

get a single service by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Interhyp/metadata-service-api"
)

func main() {
	service := "service_example" // string | The (globally unique) name of the service, must match `^[a-z](-?[a-z0-9]+)*$`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1ServicesAPI.GetService(context.Background(), service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1ServicesAPI.GetService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetService`: ServiceDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1ServicesAPI.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | The (globally unique) name of the service, must match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDto**](ServiceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePromoters

> ServicePromotersDto GetServicePromoters(ctx, service).Execute()

get all users who may promote a service

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Interhyp/metadata-service-api"
)

func main() {
	service := "service_example" // string | The (globally unique) name of the service, must match `^[a-z](-?[a-z0-9]+)*$`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1ServicesAPI.GetServicePromoters(context.Background(), service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1ServicesAPI.GetServicePromoters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicePromoters`: ServicePromotersDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1ServicesAPI.GetServicePromoters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | The (globally unique) name of the service, must match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicePromotersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicePromotersDto**](ServicePromotersDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> ServiceListDto GetServices(ctx).Owner(owner).Execute()

get services



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Interhyp/metadata-service-api"
)

func main() {
	owner := "some-owner" // string | Allows filtering the output by owner alias. Valid aliases match `^[a-z](-?[a-z0-9]+)*$`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1ServicesAPI.GetServices(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1ServicesAPI.GetServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServices`: ServiceListDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1ServicesAPI.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Allows filtering the output by owner alias. Valid aliases match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 

### Return type

[**ServiceListDto**](ServiceListDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchService

> ServiceDto PatchService(ctx, service).ServicePatchDto(servicePatchDto).Execute()

patch an existing service with the given name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Interhyp/metadata-service-api"
)

func main() {
	service := "service_example" // string | The (globally unique) name of the service, must match `^[a-z](-?[a-z0-9]+)*$`.
	servicePatchDto := *openapiclient.NewServicePatchDto("TimeStamp_example", "CommitHash_example", "JiraIssue_example") // ServicePatchDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1ServicesAPI.PatchService(context.Background(), service).ServicePatchDto(servicePatchDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1ServicesAPI.PatchService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchService`: ServiceDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1ServicesAPI.PatchService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | The (globally unique) name of the service, must match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servicePatchDto** | [**ServicePatchDto**](ServicePatchDto.md) |  | 

### Return type

[**ServiceDto**](ServiceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterService

> ServiceDto RegisterService(ctx, service).ServiceCreateDto(serviceCreateDto).Execute()

register a new service with the given name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Interhyp/metadata-service-api"
)

func main() {
	service := "service_example" // string | The (globally unique) name of the service, must match `^[a-z](-?[a-z0-9]+)*$`.
	serviceCreateDto := *openapiclient.NewServiceCreateDto("Owner_example", []openapiclient.Quicklink{*openapiclient.NewQuicklink()}, []string{"Repositories_example"}, "AlertTarget_example", "JiraIssue_example") // ServiceCreateDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1ServicesAPI.RegisterService(context.Background(), service).ServiceCreateDto(serviceCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1ServicesAPI.RegisterService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterService`: ServiceDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1ServicesAPI.RegisterService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | The (globally unique) name of the service, must match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceCreateDto** | [**ServiceCreateDto**](ServiceCreateDto.md) |  | 

### Return type

[**ServiceDto**](ServiceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ServiceDto UpdateService(ctx, service).ServiceDto(serviceDto).Execute()

update an existing service with the given name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Interhyp/metadata-service-api"
)

func main() {
	service := "service_example" // string | The (globally unique) name of the service, must match `^[a-z](-?[a-z0-9]+)*$`.
	serviceDto := *openapiclient.NewServiceDto("Owner_example", []openapiclient.Quicklink{*openapiclient.NewQuicklink()}, []string{"Repositories_example"}, "AlertTarget_example", "TimeStamp_example", "CommitHash_example", "JiraIssue_example") // ServiceDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1ServicesAPI.UpdateService(context.Background(), service).ServiceDto(serviceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1ServicesAPI.UpdateService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateService`: ServiceDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1ServicesAPI.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | The (globally unique) name of the service, must match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceDto** | [**ServiceDto**](ServiceDto.md) |  | 

### Return type

[**ServiceDto**](ServiceDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


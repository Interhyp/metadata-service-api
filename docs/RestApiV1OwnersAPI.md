# \RestApiV1OwnersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOwner**](RestApiV1OwnersAPI.md#CreateOwner) | **Post** /rest/api/v1/owners/{owner} | create a new owner with a given alias
[**DeleteOwner**](RestApiV1OwnersAPI.md#DeleteOwner) | **Delete** /rest/api/v1/owners/{owner} | delete the owner with a given alias
[**GetOwner**](RestApiV1OwnersAPI.md#GetOwner) | **Get** /rest/api/v1/owners/{owner} | get a single owner by alias
[**GetOwners**](RestApiV1OwnersAPI.md#GetOwners) | **Get** /rest/api/v1/owners | get owners
[**PatchOwner**](RestApiV1OwnersAPI.md#PatchOwner) | **Patch** /rest/api/v1/owners/{owner} | patch an existing owner with a given alias
[**UpdateOwner**](RestApiV1OwnersAPI.md#UpdateOwner) | **Put** /rest/api/v1/owners/{owner} | update an existing owner with a given alias



## CreateOwner

> OwnerDto CreateOwner(ctx, owner).OwnerCreateDto(ownerCreateDto).Execute()

create a new owner with a given alias



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
	owner := "owner_example" // string | 
	ownerCreateDto := *openapiclient.NewOwnerCreateDto("Contact_example", "JiraIssue_example") // OwnerCreateDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1OwnersAPI.CreateOwner(context.Background(), owner).OwnerCreateDto(ownerCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1OwnersAPI.CreateOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOwner`: OwnerDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1OwnersAPI.CreateOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ownerCreateDto** | [**OwnerCreateDto**](OwnerCreateDto.md) |  | 

### Return type

[**OwnerDto**](OwnerDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOwner

> DeleteOwner(ctx, owner).DeletionDto(deletionDto).Execute()

delete the owner with a given alias



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
	owner := "owner_example" // string | 
	deletionDto := *openapiclient.NewDeletionDto("JiraIssue_example") // DeletionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RestApiV1OwnersAPI.DeleteOwner(context.Background(), owner).DeletionDto(deletionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1OwnersAPI.DeleteOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOwnerRequest struct via the builder pattern


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


## GetOwner

> OwnerDto GetOwner(ctx, owner).Execute()

get a single owner by alias



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
	owner := "owner_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1OwnersAPI.GetOwner(context.Background(), owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1OwnersAPI.GetOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwner`: OwnerDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1OwnersAPI.GetOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OwnerDto**](OwnerDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwners

> OwnerListDto GetOwners(ctx).Execute()

get owners



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1OwnersAPI.GetOwners(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1OwnersAPI.GetOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwners`: OwnerListDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1OwnersAPI.GetOwners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnersRequest struct via the builder pattern


### Return type

[**OwnerListDto**](OwnerListDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOwner

> OwnerDto PatchOwner(ctx, owner).OwnerPatchDto(ownerPatchDto).Execute()

patch an existing owner with a given alias



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
	owner := "owner_example" // string | 
	ownerPatchDto := *openapiclient.NewOwnerPatchDto("TimeStamp_example", "CommitHash_example", "JiraIssue_example") // OwnerPatchDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1OwnersAPI.PatchOwner(context.Background(), owner).OwnerPatchDto(ownerPatchDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1OwnersAPI.PatchOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOwner`: OwnerDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1OwnersAPI.PatchOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ownerPatchDto** | [**OwnerPatchDto**](OwnerPatchDto.md) |  | 

### Return type

[**OwnerDto**](OwnerDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOwner

> OwnerDto UpdateOwner(ctx, owner).OwnerDto(ownerDto).Execute()

update an existing owner with a given alias



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
	owner := "owner_example" // string | 
	ownerDto := *openapiclient.NewOwnerDto("Contact_example", "TimeStamp_example", "CommitHash_example", "JiraIssue_example") // OwnerDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1OwnersAPI.UpdateOwner(context.Background(), owner).OwnerDto(ownerDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1OwnersAPI.UpdateOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOwner`: OwnerDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1OwnersAPI.UpdateOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ownerDto** | [**OwnerDto**](OwnerDto.md) |  | 

### Return type

[**OwnerDto**](OwnerDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


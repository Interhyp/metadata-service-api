# \RestApiV1RepositoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRepositoriesOfOwner**](RestApiV1RepositoriesAPI.md#GetRepositoriesOfOwner) | **Get** /rest/api/v1/repositories | get repositories
[**GetRepository**](RestApiV1RepositoriesAPI.md#GetRepository) | **Get** /rest/api/v1/repositories/{repository} | get a single repository by key
[**PatchRepository**](RestApiV1RepositoriesAPI.md#PatchRepository) | **Patch** /rest/api/v1/repositories/{repository} | patch an existing repository with the given key
[**RegisterRepository**](RestApiV1RepositoriesAPI.md#RegisterRepository) | **Post** /rest/api/v1/repositories/{repository} | register a new repository with the given key
[**RemoveRepository**](RestApiV1RepositoriesAPI.md#RemoveRepository) | **Delete** /rest/api/v1/repositories/{repository} | remove the repository with the given key
[**UpdateRepository**](RestApiV1RepositoriesAPI.md#UpdateRepository) | **Put** /rest/api/v1/repositories/{repository} | update an existing repository with the given key



## GetRepositoriesOfOwner

> RepositoryListDto GetRepositoriesOfOwner(ctx).Url(url).Owner(owner).Service(service).Name(name).Type_(type_).Execute()

get repositories



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
	url := "git@github.com:some-org/some-repo.git" // string | Optional - allows filtering the output by repository url. Must match `^[a-z](-?:?@?.?/?[a-z0-9]+)*$`. (optional)
	owner := "some-owner" // string | Optional - the alias of an owner. If present, only repositories with this owner are returned. Must match `^[a-z](-?[a-z0-9]+)*$`. (optional)
	service := "unicorn-finder-service" // string | Optional - the name of a service. If present, only repositories referenced by the given service are returned. Must match `^[a-z](-?[a-z0-9]+)*$`. (optional)
	name := "some-service" // string | Optional - allows filtering the output by repository name (the first part of the key before the .). Must match `^[a-z](-?[a-z0-9]+)*$`. (optional)
	type_ := "helm-chart" // string | Optional - allows filtering the output by repository type (the second part of the key after the .). Must currently be one of api, helm-chart, helm-deployment, implementation, terraform-module, javascript-module. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1RepositoriesAPI.GetRepositoriesOfOwner(context.Background()).Url(url).Owner(owner).Service(service).Name(name).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1RepositoriesAPI.GetRepositoriesOfOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoriesOfOwner`: RepositoryListDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1RepositoriesAPI.GetRepositoriesOfOwner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesOfOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | Optional - allows filtering the output by repository url. Must match &#x60;^[a-z](-?:?@?.?/?[a-z0-9]+)*$&#x60;. | 
 **owner** | **string** | Optional - the alias of an owner. If present, only repositories with this owner are returned. Must match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 
 **service** | **string** | Optional - the name of a service. If present, only repositories referenced by the given service are returned. Must match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 
 **name** | **string** | Optional - allows filtering the output by repository name (the first part of the key before the .). Must match &#x60;^[a-z](-?[a-z0-9]+)*$&#x60;. | 
 **type_** | **string** | Optional - allows filtering the output by repository type (the second part of the key after the .). Must currently be one of api, helm-chart, helm-deployment, implementation, terraform-module, javascript-module. | 

### Return type

[**RepositoryListDto**](RepositoryListDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> RepositoryDto GetRepository(ctx, repository).Execute()

get a single repository by key

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
	repository := "unicorn-finder-service.implementation" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1RepositoriesAPI.GetRepository(context.Background(), repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1RepositoriesAPI.GetRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository`: RepositoryDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1RepositoriesAPI.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepositoryDto**](RepositoryDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRepository

> RepositoryDto PatchRepository(ctx, repository).RepositoryPatchDto(repositoryPatchDto).Execute()

patch an existing repository with the given key



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
	repository := "unicorn-finder-service.implementation" // string | 
	repositoryPatchDto := *openapiclient.NewRepositoryPatchDto("TimeStamp_example", "CommitHash_example", "JiraIssue_example") // RepositoryPatchDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1RepositoriesAPI.PatchRepository(context.Background(), repository).RepositoryPatchDto(repositoryPatchDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1RepositoriesAPI.PatchRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRepository`: RepositoryDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1RepositoriesAPI.PatchRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryPatchDto** | [**RepositoryPatchDto**](RepositoryPatchDto.md) |  | 

### Return type

[**RepositoryDto**](RepositoryDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterRepository

> RepositoryDto RegisterRepository(ctx, repository).RepositoryCreateDto(repositoryCreateDto).Execute()

register a new repository with the given key



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
	repository := "unicorn-finder-service.implementation" // string | 
	repositoryCreateDto := *openapiclient.NewRepositoryCreateDto("Owner_example", "Url_example", "Mainline_example", "JiraIssue_example") // RepositoryCreateDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1RepositoriesAPI.RegisterRepository(context.Background(), repository).RepositoryCreateDto(repositoryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1RepositoriesAPI.RegisterRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterRepository`: RepositoryDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1RepositoriesAPI.RegisterRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryCreateDto** | [**RepositoryCreateDto**](RepositoryCreateDto.md) |  | 

### Return type

[**RepositoryDto**](RepositoryDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRepository

> RemoveRepository(ctx, repository).DeletionDto(deletionDto).Execute()

remove the repository with the given key



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
	repository := "unicorn-finder-service.implementation" // string | 
	deletionDto := *openapiclient.NewDeletionDto("JiraIssue_example") // DeletionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RestApiV1RepositoriesAPI.RemoveRepository(context.Background(), repository).DeletionDto(deletionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1RepositoriesAPI.RemoveRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRepositoryRequest struct via the builder pattern


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


## UpdateRepository

> RepositoryDto UpdateRepository(ctx, repository).RepositoryDto(repositoryDto).Execute()

update an existing repository with the given key



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
	repository := "unicorn-finder-service.implementation" // string | 
	repositoryDto := *openapiclient.NewRepositoryDto("Owner_example", "Url_example", "Mainline_example", "TimeStamp_example", "CommitHash_example", "JiraIssue_example") // RepositoryDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestApiV1RepositoriesAPI.UpdateRepository(context.Background(), repository).RepositoryDto(repositoryDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestApiV1RepositoriesAPI.UpdateRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRepository`: RepositoryDto
	fmt.Fprintf(os.Stdout, "Response from `RestApiV1RepositoriesAPI.UpdateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryDto** | [**RepositoryDto**](RepositoryDto.md) |  | 

### Return type

[**RepositoryDto**](RepositoryDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


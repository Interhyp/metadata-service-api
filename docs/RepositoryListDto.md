# RepositoryListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | [**map[string]RepositoryDto**](RepositoryDto.md) |  | 
**TimeStamp** | **string** | ISO-8601 UTC date time at which the list of repositories was obtained from service-metadata | 

## Methods

### NewRepositoryListDto

`func NewRepositoryListDto(repositories map[string]RepositoryDto, timeStamp string, ) *RepositoryListDto`

NewRepositoryListDto instantiates a new RepositoryListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryListDtoWithDefaults

`func NewRepositoryListDtoWithDefaults() *RepositoryListDto`

NewRepositoryListDtoWithDefaults instantiates a new RepositoryListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *RepositoryListDto) GetRepositories() map[string]RepositoryDto`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *RepositoryListDto) GetRepositoriesOk() (*map[string]RepositoryDto, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *RepositoryListDto) SetRepositories(v map[string]RepositoryDto)`

SetRepositories sets Repositories field to given value.


### GetTimeStamp

`func (o *RepositoryListDto) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *RepositoryListDto) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *RepositoryListDto) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



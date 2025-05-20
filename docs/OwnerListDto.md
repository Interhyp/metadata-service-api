# OwnerListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owners** | [**map[string]OwnerDto**](OwnerDto.md) |  | 
**TimeStamp** | **string** | ISO-8601 UTC date time at which the list of owners was obtained from service-metadata | 

## Methods

### NewOwnerListDto

`func NewOwnerListDto(owners map[string]OwnerDto, timeStamp string, ) *OwnerListDto`

NewOwnerListDto instantiates a new OwnerListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerListDtoWithDefaults

`func NewOwnerListDtoWithDefaults() *OwnerListDto`

NewOwnerListDtoWithDefaults instantiates a new OwnerListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwners

`func (o *OwnerListDto) GetOwners() map[string]OwnerDto`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *OwnerListDto) GetOwnersOk() (*map[string]OwnerDto, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *OwnerListDto) SetOwners(v map[string]OwnerDto)`

SetOwners sets Owners field to given value.


### GetTimeStamp

`func (o *OwnerListDto) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *OwnerListDto) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *OwnerListDto) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



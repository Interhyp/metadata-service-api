# ServiceListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | [**map[string]ServiceDto**](ServiceDto.md) |  | 
**TimeStamp** | **string** | ISO-8601 UTC date time at which the list of services was obtained from service-metadata | 

## Methods

### NewServiceListDto

`func NewServiceListDto(services map[string]ServiceDto, timeStamp string, ) *ServiceListDto`

NewServiceListDto instantiates a new ServiceListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceListDtoWithDefaults

`func NewServiceListDtoWithDefaults() *ServiceListDto`

NewServiceListDtoWithDefaults instantiates a new ServiceListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *ServiceListDto) GetServices() map[string]ServiceDto`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServiceListDto) GetServicesOk() (*map[string]ServiceDto, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServiceListDto) SetServices(v map[string]ServiceDto)`

SetServices sets Services field to given value.


### GetTimeStamp

`func (o *ServiceListDto) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ServiceListDto) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ServiceListDto) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



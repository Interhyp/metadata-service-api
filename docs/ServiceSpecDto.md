# ServiceSpecDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**System** | Pointer to **string** | A reference to the system that the component belongs to | [optional] 
**DependsOn** | Pointer to **[]string** | A relation denoting a dependency on another entity | [optional] 
**ProvidesApis** | Pointer to **[]string** | A relation with an API, provided by this entity | [optional] 
**ConsumesApis** | Pointer to **[]string** | A relation with an API, consumed by this entity | [optional] 

## Methods

### NewServiceSpecDto

`func NewServiceSpecDto() *ServiceSpecDto`

NewServiceSpecDto instantiates a new ServiceSpecDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSpecDtoWithDefaults

`func NewServiceSpecDtoWithDefaults() *ServiceSpecDto`

NewServiceSpecDtoWithDefaults instantiates a new ServiceSpecDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystem

`func (o *ServiceSpecDto) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ServiceSpecDto) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ServiceSpecDto) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ServiceSpecDto) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetDependsOn

`func (o *ServiceSpecDto) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *ServiceSpecDto) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *ServiceSpecDto) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *ServiceSpecDto) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetProvidesApis

`func (o *ServiceSpecDto) GetProvidesApis() []string`

GetProvidesApis returns the ProvidesApis field if non-nil, zero value otherwise.

### GetProvidesApisOk

`func (o *ServiceSpecDto) GetProvidesApisOk() (*[]string, bool)`

GetProvidesApisOk returns a tuple with the ProvidesApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidesApis

`func (o *ServiceSpecDto) SetProvidesApis(v []string)`

SetProvidesApis sets ProvidesApis field to given value.

### HasProvidesApis

`func (o *ServiceSpecDto) HasProvidesApis() bool`

HasProvidesApis returns a boolean if a field has been set.

### GetConsumesApis

`func (o *ServiceSpecDto) GetConsumesApis() []string`

GetConsumesApis returns the ConsumesApis field if non-nil, zero value otherwise.

### GetConsumesApisOk

`func (o *ServiceSpecDto) GetConsumesApisOk() (*[]string, bool)`

GetConsumesApisOk returns a tuple with the ConsumesApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumesApis

`func (o *ServiceSpecDto) SetConsumesApis(v []string)`

SetConsumesApis sets ConsumesApis field to given value.

### HasConsumesApis

`func (o *ServiceSpecDto) HasConsumesApis() bool`

HasConsumesApis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



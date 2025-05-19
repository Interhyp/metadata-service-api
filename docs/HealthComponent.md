# HealthComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthComponent

`func NewHealthComponent() *HealthComponent`

NewHealthComponent instantiates a new HealthComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthComponentWithDefaults

`func NewHealthComponentWithDefaults() *HealthComponent`

NewHealthComponentWithDefaults instantiates a new HealthComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HealthComponent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthComponent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthComponent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthComponent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *HealthComponent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthComponent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthComponent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthComponent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



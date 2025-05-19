# ConditionReferenceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefMatcher** | **string** | Reference of a branch. | 
**Exemptions** | Pointer to **[]string** | list of users or groups for which this protection does not apply. | [optional] 
**Source** | Pointer to **string** | The expected source for the required conditional build. | [optional] 

## Methods

### NewConditionReferenceDto

`func NewConditionReferenceDto(refMatcher string, ) *ConditionReferenceDto`

NewConditionReferenceDto instantiates a new ConditionReferenceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionReferenceDtoWithDefaults

`func NewConditionReferenceDtoWithDefaults() *ConditionReferenceDto`

NewConditionReferenceDtoWithDefaults instantiates a new ConditionReferenceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefMatcher

`func (o *ConditionReferenceDto) GetRefMatcher() string`

GetRefMatcher returns the RefMatcher field if non-nil, zero value otherwise.

### GetRefMatcherOk

`func (o *ConditionReferenceDto) GetRefMatcherOk() (*string, bool)`

GetRefMatcherOk returns a tuple with the RefMatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefMatcher

`func (o *ConditionReferenceDto) SetRefMatcher(v string)`

SetRefMatcher sets RefMatcher field to given value.


### GetExemptions

`func (o *ConditionReferenceDto) GetExemptions() []string`

GetExemptions returns the Exemptions field if non-nil, zero value otherwise.

### GetExemptionsOk

`func (o *ConditionReferenceDto) GetExemptionsOk() (*[]string, bool)`

GetExemptionsOk returns a tuple with the Exemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptions

`func (o *ConditionReferenceDto) SetExemptions(v []string)`

SetExemptions sets Exemptions field to given value.

### HasExemptions

`func (o *ConditionReferenceDto) HasExemptions() bool`

HasExemptions returns a boolean if a field has been set.

### GetSource

`func (o *ConditionReferenceDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConditionReferenceDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConditionReferenceDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConditionReferenceDto) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



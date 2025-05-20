# ProtectedRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | **string** | fnmatch pattern to define protected refs. Must not start with refs/heads/ or refs/tags/. Special value :MAINLINE: matches the currently configured mainline for branch protections. | 
**Exemptions** | Pointer to **[]string** | list of users or groups for which this protection does not apply. | [optional] 
**ExemptionsRoles** | Pointer to **[]string** | list of group identifiers for which this protection does not apply. This field is read-only and will be filled automatically from the exemptions fields. | [optional] 

## Methods

### NewProtectedRef

`func NewProtectedRef(pattern string, ) *ProtectedRef`

NewProtectedRef instantiates a new ProtectedRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedRefWithDefaults

`func NewProtectedRefWithDefaults() *ProtectedRef`

NewProtectedRefWithDefaults instantiates a new ProtectedRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *ProtectedRef) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ProtectedRef) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ProtectedRef) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetExemptions

`func (o *ProtectedRef) GetExemptions() []string`

GetExemptions returns the Exemptions field if non-nil, zero value otherwise.

### GetExemptionsOk

`func (o *ProtectedRef) GetExemptionsOk() (*[]string, bool)`

GetExemptionsOk returns a tuple with the Exemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptions

`func (o *ProtectedRef) SetExemptions(v []string)`

SetExemptions sets Exemptions field to given value.

### HasExemptions

`func (o *ProtectedRef) HasExemptions() bool`

HasExemptions returns a boolean if a field has been set.

### GetExemptionsRoles

`func (o *ProtectedRef) GetExemptionsRoles() []string`

GetExemptionsRoles returns the ExemptionsRoles field if non-nil, zero value otherwise.

### GetExemptionsRolesOk

`func (o *ProtectedRef) GetExemptionsRolesOk() (*[]string, bool)`

GetExemptionsRolesOk returns a tuple with the ExemptionsRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptionsRoles

`func (o *ProtectedRef) SetExemptionsRoles(v []string)`

SetExemptionsRoles sets ExemptionsRoles field to given value.

### HasExemptionsRoles

`func (o *ProtectedRef) HasExemptionsRoles() bool`

HasExemptionsRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



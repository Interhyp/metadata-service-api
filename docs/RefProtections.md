# RefProtections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**RefProtectionsBranches**](RefProtectionsBranches.md) |  | [optional] 
**Tags** | Pointer to [**RefProtectionsTags**](RefProtectionsTags.md) |  | [optional] 

## Methods

### NewRefProtections

`func NewRefProtections() *RefProtections`

NewRefProtections instantiates a new RefProtections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefProtectionsWithDefaults

`func NewRefProtectionsWithDefaults() *RefProtections`

NewRefProtectionsWithDefaults instantiates a new RefProtections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *RefProtections) GetBranches() RefProtectionsBranches`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *RefProtections) GetBranchesOk() (*RefProtectionsBranches, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *RefProtections) SetBranches(v RefProtectionsBranches)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *RefProtections) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetTags

`func (o *RefProtections) GetTags() RefProtectionsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RefProtections) GetTagsOk() (*RefProtectionsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RefProtections) SetTags(v RefProtectionsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RefProtections) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



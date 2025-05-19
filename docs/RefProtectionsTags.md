# RefProtectionsTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreventAllChanges** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Prevents all changes of the protected refs. | [optional] 
**PreventCreation** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Prevents creation of the protected refs. | [optional] 
**PreventDeletion** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Prevents deletion of the protected refs. | [optional] 
**PreventForcePush** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Prevents force pushes to the protected refs for users with push permission. | [optional] 

## Methods

### NewRefProtectionsTags

`func NewRefProtectionsTags() *RefProtectionsTags`

NewRefProtectionsTags instantiates a new RefProtectionsTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefProtectionsTagsWithDefaults

`func NewRefProtectionsTagsWithDefaults() *RefProtectionsTags`

NewRefProtectionsTagsWithDefaults instantiates a new RefProtectionsTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreventAllChanges

`func (o *RefProtectionsTags) GetPreventAllChanges() []ProtectedRef`

GetPreventAllChanges returns the PreventAllChanges field if non-nil, zero value otherwise.

### GetPreventAllChangesOk

`func (o *RefProtectionsTags) GetPreventAllChangesOk() (*[]ProtectedRef, bool)`

GetPreventAllChangesOk returns a tuple with the PreventAllChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAllChanges

`func (o *RefProtectionsTags) SetPreventAllChanges(v []ProtectedRef)`

SetPreventAllChanges sets PreventAllChanges field to given value.

### HasPreventAllChanges

`func (o *RefProtectionsTags) HasPreventAllChanges() bool`

HasPreventAllChanges returns a boolean if a field has been set.

### GetPreventCreation

`func (o *RefProtectionsTags) GetPreventCreation() []ProtectedRef`

GetPreventCreation returns the PreventCreation field if non-nil, zero value otherwise.

### GetPreventCreationOk

`func (o *RefProtectionsTags) GetPreventCreationOk() (*[]ProtectedRef, bool)`

GetPreventCreationOk returns a tuple with the PreventCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventCreation

`func (o *RefProtectionsTags) SetPreventCreation(v []ProtectedRef)`

SetPreventCreation sets PreventCreation field to given value.

### HasPreventCreation

`func (o *RefProtectionsTags) HasPreventCreation() bool`

HasPreventCreation returns a boolean if a field has been set.

### GetPreventDeletion

`func (o *RefProtectionsTags) GetPreventDeletion() []ProtectedRef`

GetPreventDeletion returns the PreventDeletion field if non-nil, zero value otherwise.

### GetPreventDeletionOk

`func (o *RefProtectionsTags) GetPreventDeletionOk() (*[]ProtectedRef, bool)`

GetPreventDeletionOk returns a tuple with the PreventDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventDeletion

`func (o *RefProtectionsTags) SetPreventDeletion(v []ProtectedRef)`

SetPreventDeletion sets PreventDeletion field to given value.

### HasPreventDeletion

`func (o *RefProtectionsTags) HasPreventDeletion() bool`

HasPreventDeletion returns a boolean if a field has been set.

### GetPreventForcePush

`func (o *RefProtectionsTags) GetPreventForcePush() []ProtectedRef`

GetPreventForcePush returns the PreventForcePush field if non-nil, zero value otherwise.

### GetPreventForcePushOk

`func (o *RefProtectionsTags) GetPreventForcePushOk() (*[]ProtectedRef, bool)`

GetPreventForcePushOk returns a tuple with the PreventForcePush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventForcePush

`func (o *RefProtectionsTags) SetPreventForcePush(v []ProtectedRef)`

SetPreventForcePush sets PreventForcePush field to given value.

### HasPreventForcePush

`func (o *RefProtectionsTags) HasPreventForcePush() bool`

HasPreventForcePush returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



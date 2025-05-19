# RefProtectionsBranches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirePR** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Forces creating a PR to update the protected refs. | [optional] 
**PreventAllChanges** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Prevents all changes of the protected refs. | [optional] 
**PreventCreation** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Prevents creation of the protected refs. | [optional] 
**PreventDeletion** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Prevents deletion of the protected refs. | [optional] 
**PreventPush** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Prevents pushes to the protected refs. | [optional] 
**PreventForcePush** | Pointer to [**[]ProtectedRef**](ProtectedRef.md) | Prevents force pushes to the protected refs for users with push permission. | [optional] 

## Methods

### NewRefProtectionsBranches

`func NewRefProtectionsBranches() *RefProtectionsBranches`

NewRefProtectionsBranches instantiates a new RefProtectionsBranches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefProtectionsBranchesWithDefaults

`func NewRefProtectionsBranchesWithDefaults() *RefProtectionsBranches`

NewRefProtectionsBranchesWithDefaults instantiates a new RefProtectionsBranches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirePR

`func (o *RefProtectionsBranches) GetRequirePR() []ProtectedRef`

GetRequirePR returns the RequirePR field if non-nil, zero value otherwise.

### GetRequirePROk

`func (o *RefProtectionsBranches) GetRequirePROk() (*[]ProtectedRef, bool)`

GetRequirePROk returns a tuple with the RequirePR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePR

`func (o *RefProtectionsBranches) SetRequirePR(v []ProtectedRef)`

SetRequirePR sets RequirePR field to given value.

### HasRequirePR

`func (o *RefProtectionsBranches) HasRequirePR() bool`

HasRequirePR returns a boolean if a field has been set.

### GetPreventAllChanges

`func (o *RefProtectionsBranches) GetPreventAllChanges() []ProtectedRef`

GetPreventAllChanges returns the PreventAllChanges field if non-nil, zero value otherwise.

### GetPreventAllChangesOk

`func (o *RefProtectionsBranches) GetPreventAllChangesOk() (*[]ProtectedRef, bool)`

GetPreventAllChangesOk returns a tuple with the PreventAllChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAllChanges

`func (o *RefProtectionsBranches) SetPreventAllChanges(v []ProtectedRef)`

SetPreventAllChanges sets PreventAllChanges field to given value.

### HasPreventAllChanges

`func (o *RefProtectionsBranches) HasPreventAllChanges() bool`

HasPreventAllChanges returns a boolean if a field has been set.

### GetPreventCreation

`func (o *RefProtectionsBranches) GetPreventCreation() []ProtectedRef`

GetPreventCreation returns the PreventCreation field if non-nil, zero value otherwise.

### GetPreventCreationOk

`func (o *RefProtectionsBranches) GetPreventCreationOk() (*[]ProtectedRef, bool)`

GetPreventCreationOk returns a tuple with the PreventCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventCreation

`func (o *RefProtectionsBranches) SetPreventCreation(v []ProtectedRef)`

SetPreventCreation sets PreventCreation field to given value.

### HasPreventCreation

`func (o *RefProtectionsBranches) HasPreventCreation() bool`

HasPreventCreation returns a boolean if a field has been set.

### GetPreventDeletion

`func (o *RefProtectionsBranches) GetPreventDeletion() []ProtectedRef`

GetPreventDeletion returns the PreventDeletion field if non-nil, zero value otherwise.

### GetPreventDeletionOk

`func (o *RefProtectionsBranches) GetPreventDeletionOk() (*[]ProtectedRef, bool)`

GetPreventDeletionOk returns a tuple with the PreventDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventDeletion

`func (o *RefProtectionsBranches) SetPreventDeletion(v []ProtectedRef)`

SetPreventDeletion sets PreventDeletion field to given value.

### HasPreventDeletion

`func (o *RefProtectionsBranches) HasPreventDeletion() bool`

HasPreventDeletion returns a boolean if a field has been set.

### GetPreventPush

`func (o *RefProtectionsBranches) GetPreventPush() []ProtectedRef`

GetPreventPush returns the PreventPush field if non-nil, zero value otherwise.

### GetPreventPushOk

`func (o *RefProtectionsBranches) GetPreventPushOk() (*[]ProtectedRef, bool)`

GetPreventPushOk returns a tuple with the PreventPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventPush

`func (o *RefProtectionsBranches) SetPreventPush(v []ProtectedRef)`

SetPreventPush sets PreventPush field to given value.

### HasPreventPush

`func (o *RefProtectionsBranches) HasPreventPush() bool`

HasPreventPush returns a boolean if a field has been set.

### GetPreventForcePush

`func (o *RefProtectionsBranches) GetPreventForcePush() []ProtectedRef`

GetPreventForcePush returns the PreventForcePush field if non-nil, zero value otherwise.

### GetPreventForcePushOk

`func (o *RefProtectionsBranches) GetPreventForcePushOk() (*[]ProtectedRef, bool)`

GetPreventForcePushOk returns a tuple with the PreventForcePush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventForcePush

`func (o *RefProtectionsBranches) SetPreventForcePush(v []ProtectedRef)`

SetPreventForcePush sets PreventForcePush field to given value.

### HasPreventForcePush

`func (o *RefProtectionsBranches) HasPreventForcePush() bool`

HasPreventForcePush returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



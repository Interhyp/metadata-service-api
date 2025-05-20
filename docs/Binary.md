# Binary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The group id of binary | 
**ArtifactId** | **string** | The artifact id of binary | 
**VersionPrefix** | **string** | The version prefix of binary | 
**Classifier** | Pointer to **string** | The classifier of binary | [optional] 
**FileType** | Pointer to **string** | The file type of binary e.g. tar.gz | [optional] 

## Methods

### NewBinary

`func NewBinary(groupId string, artifactId string, versionPrefix string, ) *Binary`

NewBinary instantiates a new Binary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryWithDefaults

`func NewBinaryWithDefaults() *Binary`

NewBinaryWithDefaults instantiates a new Binary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *Binary) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Binary) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Binary) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetArtifactId

`func (o *Binary) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *Binary) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *Binary) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.


### GetVersionPrefix

`func (o *Binary) GetVersionPrefix() string`

GetVersionPrefix returns the VersionPrefix field if non-nil, zero value otherwise.

### GetVersionPrefixOk

`func (o *Binary) GetVersionPrefixOk() (*string, bool)`

GetVersionPrefixOk returns a tuple with the VersionPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPrefix

`func (o *Binary) SetVersionPrefix(v string)`

SetVersionPrefix sets VersionPrefix field to given value.


### GetClassifier

`func (o *Binary) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *Binary) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *Binary) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *Binary) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetFileType

`func (o *Binary) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *Binary) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *Binary) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *Binary) HasFileType() bool`

HasFileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



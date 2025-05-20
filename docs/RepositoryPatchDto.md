# RepositoryPatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **string** | The alias of the repository owner | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Mainline** | Pointer to **string** |  | [optional] 
**Generator** | Pointer to **string** | the generator used for the initial contents of this repository | [optional] 
**Configuration** | Pointer to [**RepositoryConfigurationPatchDto**](RepositoryConfigurationPatchDto.md) |  | [optional] 
**TimeStamp** | **string** | ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates. | 
**CommitHash** | **string** | The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates. | 
**JiraIssue** | **string** | The jira issue to use for committing a change, or the last jira issue used. | 
**Labels** | Pointer to **map[string]string** | A map of arbitrary string labels attached to this repository. | [optional] 

## Methods

### NewRepositoryPatchDto

`func NewRepositoryPatchDto(timeStamp string, commitHash string, jiraIssue string, ) *RepositoryPatchDto`

NewRepositoryPatchDto instantiates a new RepositoryPatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryPatchDtoWithDefaults

`func NewRepositoryPatchDtoWithDefaults() *RepositoryPatchDto`

NewRepositoryPatchDtoWithDefaults instantiates a new RepositoryPatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *RepositoryPatchDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RepositoryPatchDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RepositoryPatchDto) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *RepositoryPatchDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetUrl

`func (o *RepositoryPatchDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RepositoryPatchDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RepositoryPatchDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RepositoryPatchDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMainline

`func (o *RepositoryPatchDto) GetMainline() string`

GetMainline returns the Mainline field if non-nil, zero value otherwise.

### GetMainlineOk

`func (o *RepositoryPatchDto) GetMainlineOk() (*string, bool)`

GetMainlineOk returns a tuple with the Mainline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainline

`func (o *RepositoryPatchDto) SetMainline(v string)`

SetMainline sets Mainline field to given value.

### HasMainline

`func (o *RepositoryPatchDto) HasMainline() bool`

HasMainline returns a boolean if a field has been set.

### GetGenerator

`func (o *RepositoryPatchDto) GetGenerator() string`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *RepositoryPatchDto) GetGeneratorOk() (*string, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *RepositoryPatchDto) SetGenerator(v string)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *RepositoryPatchDto) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.

### GetConfiguration

`func (o *RepositoryPatchDto) GetConfiguration() RepositoryConfigurationPatchDto`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RepositoryPatchDto) GetConfigurationOk() (*RepositoryConfigurationPatchDto, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RepositoryPatchDto) SetConfiguration(v RepositoryConfigurationPatchDto)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RepositoryPatchDto) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTimeStamp

`func (o *RepositoryPatchDto) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *RepositoryPatchDto) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *RepositoryPatchDto) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetCommitHash

`func (o *RepositoryPatchDto) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *RepositoryPatchDto) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *RepositoryPatchDto) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### GetJiraIssue

`func (o *RepositoryPatchDto) GetJiraIssue() string`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *RepositoryPatchDto) GetJiraIssueOk() (*string, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *RepositoryPatchDto) SetJiraIssue(v string)`

SetJiraIssue sets JiraIssue field to given value.


### GetLabels

`func (o *RepositoryPatchDto) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RepositoryPatchDto) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RepositoryPatchDto) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RepositoryPatchDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



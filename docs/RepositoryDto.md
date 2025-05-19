# RepositoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the repository as determined by its key. | [optional] 
**Owner** | **string** | The alias of the repository owner | 
**Description** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Mainline** | **string** |  | 
**Generator** | Pointer to **string** | the generator used for the initial contents of this repository | [optional] 
**Configuration** | Pointer to [**RepositoryConfigurationDto**](RepositoryConfigurationDto.md) |  | [optional] 
**TimeStamp** | **string** | ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates. | 
**CommitHash** | **string** | The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates. | 
**JiraIssue** | **string** | The jira issue to use for committing a change, or the last jira issue used. | 
**Labels** | Pointer to **map[string]string** | A map of arbitrary string labels attached to this repository. | [optional] 

## Methods

### NewRepositoryDto

`func NewRepositoryDto(owner string, url string, mainline string, timeStamp string, commitHash string, jiraIssue string, ) *RepositoryDto`

NewRepositoryDto instantiates a new RepositoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryDtoWithDefaults

`func NewRepositoryDtoWithDefaults() *RepositoryDto`

NewRepositoryDtoWithDefaults instantiates a new RepositoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RepositoryDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositoryDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositoryDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RepositoryDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *RepositoryDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RepositoryDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RepositoryDto) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetDescription

`func (o *RepositoryDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositoryDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositoryDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RepositoryDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *RepositoryDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RepositoryDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RepositoryDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMainline

`func (o *RepositoryDto) GetMainline() string`

GetMainline returns the Mainline field if non-nil, zero value otherwise.

### GetMainlineOk

`func (o *RepositoryDto) GetMainlineOk() (*string, bool)`

GetMainlineOk returns a tuple with the Mainline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainline

`func (o *RepositoryDto) SetMainline(v string)`

SetMainline sets Mainline field to given value.


### GetGenerator

`func (o *RepositoryDto) GetGenerator() string`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *RepositoryDto) GetGeneratorOk() (*string, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *RepositoryDto) SetGenerator(v string)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *RepositoryDto) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.

### GetConfiguration

`func (o *RepositoryDto) GetConfiguration() RepositoryConfigurationDto`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RepositoryDto) GetConfigurationOk() (*RepositoryConfigurationDto, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RepositoryDto) SetConfiguration(v RepositoryConfigurationDto)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RepositoryDto) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTimeStamp

`func (o *RepositoryDto) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *RepositoryDto) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *RepositoryDto) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetCommitHash

`func (o *RepositoryDto) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *RepositoryDto) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *RepositoryDto) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### GetJiraIssue

`func (o *RepositoryDto) GetJiraIssue() string`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *RepositoryDto) GetJiraIssueOk() (*string, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *RepositoryDto) SetJiraIssue(v string)`

SetJiraIssue sets JiraIssue field to given value.


### GetLabels

`func (o *RepositoryDto) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RepositoryDto) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RepositoryDto) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RepositoryDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



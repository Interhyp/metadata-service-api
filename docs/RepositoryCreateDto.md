# RepositoryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | The alias of the repository owner | 
**Url** | **string** |  | 
**Mainline** | **string** |  | 
**Generator** | Pointer to **string** | the generator used for the initial contents of this repository | [optional] 
**Configuration** | Pointer to [**RepositoryConfigurationDto**](RepositoryConfigurationDto.md) |  | [optional] 
**JiraIssue** | **string** | The jira issue to use for committing a change, or the last jira issue used. | 
**Labels** | Pointer to **map[string]string** | A map of arbitrary string labels attached to this repository. | [optional] 

## Methods

### NewRepositoryCreateDto

`func NewRepositoryCreateDto(owner string, url string, mainline string, jiraIssue string, ) *RepositoryCreateDto`

NewRepositoryCreateDto instantiates a new RepositoryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryCreateDtoWithDefaults

`func NewRepositoryCreateDtoWithDefaults() *RepositoryCreateDto`

NewRepositoryCreateDtoWithDefaults instantiates a new RepositoryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *RepositoryCreateDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RepositoryCreateDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RepositoryCreateDto) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetUrl

`func (o *RepositoryCreateDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RepositoryCreateDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RepositoryCreateDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMainline

`func (o *RepositoryCreateDto) GetMainline() string`

GetMainline returns the Mainline field if non-nil, zero value otherwise.

### GetMainlineOk

`func (o *RepositoryCreateDto) GetMainlineOk() (*string, bool)`

GetMainlineOk returns a tuple with the Mainline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainline

`func (o *RepositoryCreateDto) SetMainline(v string)`

SetMainline sets Mainline field to given value.


### GetGenerator

`func (o *RepositoryCreateDto) GetGenerator() string`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *RepositoryCreateDto) GetGeneratorOk() (*string, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *RepositoryCreateDto) SetGenerator(v string)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *RepositoryCreateDto) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.

### GetConfiguration

`func (o *RepositoryCreateDto) GetConfiguration() RepositoryConfigurationDto`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RepositoryCreateDto) GetConfigurationOk() (*RepositoryConfigurationDto, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RepositoryCreateDto) SetConfiguration(v RepositoryConfigurationDto)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RepositoryCreateDto) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetJiraIssue

`func (o *RepositoryCreateDto) GetJiraIssue() string`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *RepositoryCreateDto) GetJiraIssueOk() (*string, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *RepositoryCreateDto) SetJiraIssue(v string)`

SetJiraIssue sets JiraIssue field to given value.


### GetLabels

`func (o *RepositoryCreateDto) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RepositoryCreateDto) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RepositoryCreateDto) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RepositoryCreateDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



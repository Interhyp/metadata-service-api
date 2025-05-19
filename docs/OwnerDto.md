# OwnerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | **string** | The contact information of the owner | 
**TeamsChannelURL** | Pointer to **string** | The teams channel url information of the owner | [optional] 
**ProductOwner** | Pointer to **string** | The product owner of this owner space | [optional] 
**Members** | Pointer to **[]string** | A list of users which constitute this owner | [optional] 
**Groups** | Pointer to **map[string][]string** | Collection of arbitrary user groups which can be referenced in service configurations. Map of string (group name e.g. some-owner) of strings (list of usernames), one username for each group is required. | [optional] 
**Promoters** | Pointer to **[]string** | A list of users that are allowed to promote services in this owner space | [optional] 
**DefaultJiraProject** | Pointer to **string** | The default jira project that is used by this owner space | [optional] 
**TimeStamp** | **string** | ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates. | 
**CommitHash** | **string** | The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates. | 
**JiraIssue** | **string** | The jira issue to use for committing a change, or the last jira issue used. | 
**DisplayName** | Pointer to **string** | A display name of the owner, to be presented in user interfaces instead of the owner&#39;s name, when available | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewOwnerDto

`func NewOwnerDto(contact string, timeStamp string, commitHash string, jiraIssue string, ) *OwnerDto`

NewOwnerDto instantiates a new OwnerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerDtoWithDefaults

`func NewOwnerDtoWithDefaults() *OwnerDto`

NewOwnerDtoWithDefaults instantiates a new OwnerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *OwnerDto) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OwnerDto) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OwnerDto) SetContact(v string)`

SetContact sets Contact field to given value.


### GetTeamsChannelURL

`func (o *OwnerDto) GetTeamsChannelURL() string`

GetTeamsChannelURL returns the TeamsChannelURL field if non-nil, zero value otherwise.

### GetTeamsChannelURLOk

`func (o *OwnerDto) GetTeamsChannelURLOk() (*string, bool)`

GetTeamsChannelURLOk returns a tuple with the TeamsChannelURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsChannelURL

`func (o *OwnerDto) SetTeamsChannelURL(v string)`

SetTeamsChannelURL sets TeamsChannelURL field to given value.

### HasTeamsChannelURL

`func (o *OwnerDto) HasTeamsChannelURL() bool`

HasTeamsChannelURL returns a boolean if a field has been set.

### GetProductOwner

`func (o *OwnerDto) GetProductOwner() string`

GetProductOwner returns the ProductOwner field if non-nil, zero value otherwise.

### GetProductOwnerOk

`func (o *OwnerDto) GetProductOwnerOk() (*string, bool)`

GetProductOwnerOk returns a tuple with the ProductOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOwner

`func (o *OwnerDto) SetProductOwner(v string)`

SetProductOwner sets ProductOwner field to given value.

### HasProductOwner

`func (o *OwnerDto) HasProductOwner() bool`

HasProductOwner returns a boolean if a field has been set.

### GetMembers

`func (o *OwnerDto) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OwnerDto) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OwnerDto) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OwnerDto) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetGroups

`func (o *OwnerDto) GetGroups() map[string][]string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *OwnerDto) GetGroupsOk() (*map[string][]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *OwnerDto) SetGroups(v map[string][]string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *OwnerDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetPromoters

`func (o *OwnerDto) GetPromoters() []string`

GetPromoters returns the Promoters field if non-nil, zero value otherwise.

### GetPromotersOk

`func (o *OwnerDto) GetPromotersOk() (*[]string, bool)`

GetPromotersOk returns a tuple with the Promoters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoters

`func (o *OwnerDto) SetPromoters(v []string)`

SetPromoters sets Promoters field to given value.

### HasPromoters

`func (o *OwnerDto) HasPromoters() bool`

HasPromoters returns a boolean if a field has been set.

### GetDefaultJiraProject

`func (o *OwnerDto) GetDefaultJiraProject() string`

GetDefaultJiraProject returns the DefaultJiraProject field if non-nil, zero value otherwise.

### GetDefaultJiraProjectOk

`func (o *OwnerDto) GetDefaultJiraProjectOk() (*string, bool)`

GetDefaultJiraProjectOk returns a tuple with the DefaultJiraProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultJiraProject

`func (o *OwnerDto) SetDefaultJiraProject(v string)`

SetDefaultJiraProject sets DefaultJiraProject field to given value.

### HasDefaultJiraProject

`func (o *OwnerDto) HasDefaultJiraProject() bool`

HasDefaultJiraProject returns a boolean if a field has been set.

### GetTimeStamp

`func (o *OwnerDto) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *OwnerDto) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *OwnerDto) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetCommitHash

`func (o *OwnerDto) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *OwnerDto) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *OwnerDto) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### GetJiraIssue

`func (o *OwnerDto) GetJiraIssue() string`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *OwnerDto) GetJiraIssueOk() (*string, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *OwnerDto) SetJiraIssue(v string)`

SetJiraIssue sets JiraIssue field to given value.


### GetDisplayName

`func (o *OwnerDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OwnerDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OwnerDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OwnerDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLinks

`func (o *OwnerDto) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OwnerDto) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OwnerDto) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OwnerDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



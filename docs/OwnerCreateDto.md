# OwnerCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | **string** | The contact information of the owner | 
**TeamsChannelURL** | Pointer to **string** | The teams channel url information of the owner | [optional] 
**ProductOwner** | Pointer to **string** | The product owner of this owner space | [optional] 
**Promoters** | Pointer to **[]string** | A list of users that are allowed to promote services in this owner space | [optional] 
**Members** | Pointer to **[]string** | A list of users which constitute this owner | [optional] 
**Groups** | Pointer to **map[string][]string** | Map of string (group name e.g. some-owner) of strings (list of usernames), one username for each group is required. | [optional] 
**DefaultJiraProject** | Pointer to **string** | The default jira project that is used by this owner space | [optional] 
**JiraIssue** | **string** | The jira issue to use for committing a change, or the last jira issue used. | 
**DisplayName** | Pointer to **string** | A display name of the owner, to be presented in user interfaces instead of the owner&#39;s name, when available | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewOwnerCreateDto

`func NewOwnerCreateDto(contact string, jiraIssue string, ) *OwnerCreateDto`

NewOwnerCreateDto instantiates a new OwnerCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerCreateDtoWithDefaults

`func NewOwnerCreateDtoWithDefaults() *OwnerCreateDto`

NewOwnerCreateDtoWithDefaults instantiates a new OwnerCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *OwnerCreateDto) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OwnerCreateDto) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OwnerCreateDto) SetContact(v string)`

SetContact sets Contact field to given value.


### GetTeamsChannelURL

`func (o *OwnerCreateDto) GetTeamsChannelURL() string`

GetTeamsChannelURL returns the TeamsChannelURL field if non-nil, zero value otherwise.

### GetTeamsChannelURLOk

`func (o *OwnerCreateDto) GetTeamsChannelURLOk() (*string, bool)`

GetTeamsChannelURLOk returns a tuple with the TeamsChannelURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsChannelURL

`func (o *OwnerCreateDto) SetTeamsChannelURL(v string)`

SetTeamsChannelURL sets TeamsChannelURL field to given value.

### HasTeamsChannelURL

`func (o *OwnerCreateDto) HasTeamsChannelURL() bool`

HasTeamsChannelURL returns a boolean if a field has been set.

### GetProductOwner

`func (o *OwnerCreateDto) GetProductOwner() string`

GetProductOwner returns the ProductOwner field if non-nil, zero value otherwise.

### GetProductOwnerOk

`func (o *OwnerCreateDto) GetProductOwnerOk() (*string, bool)`

GetProductOwnerOk returns a tuple with the ProductOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOwner

`func (o *OwnerCreateDto) SetProductOwner(v string)`

SetProductOwner sets ProductOwner field to given value.

### HasProductOwner

`func (o *OwnerCreateDto) HasProductOwner() bool`

HasProductOwner returns a boolean if a field has been set.

### GetPromoters

`func (o *OwnerCreateDto) GetPromoters() []string`

GetPromoters returns the Promoters field if non-nil, zero value otherwise.

### GetPromotersOk

`func (o *OwnerCreateDto) GetPromotersOk() (*[]string, bool)`

GetPromotersOk returns a tuple with the Promoters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoters

`func (o *OwnerCreateDto) SetPromoters(v []string)`

SetPromoters sets Promoters field to given value.

### HasPromoters

`func (o *OwnerCreateDto) HasPromoters() bool`

HasPromoters returns a boolean if a field has been set.

### GetMembers

`func (o *OwnerCreateDto) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OwnerCreateDto) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OwnerCreateDto) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OwnerCreateDto) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetGroups

`func (o *OwnerCreateDto) GetGroups() map[string][]string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *OwnerCreateDto) GetGroupsOk() (*map[string][]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *OwnerCreateDto) SetGroups(v map[string][]string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *OwnerCreateDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetDefaultJiraProject

`func (o *OwnerCreateDto) GetDefaultJiraProject() string`

GetDefaultJiraProject returns the DefaultJiraProject field if non-nil, zero value otherwise.

### GetDefaultJiraProjectOk

`func (o *OwnerCreateDto) GetDefaultJiraProjectOk() (*string, bool)`

GetDefaultJiraProjectOk returns a tuple with the DefaultJiraProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultJiraProject

`func (o *OwnerCreateDto) SetDefaultJiraProject(v string)`

SetDefaultJiraProject sets DefaultJiraProject field to given value.

### HasDefaultJiraProject

`func (o *OwnerCreateDto) HasDefaultJiraProject() bool`

HasDefaultJiraProject returns a boolean if a field has been set.

### GetJiraIssue

`func (o *OwnerCreateDto) GetJiraIssue() string`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *OwnerCreateDto) GetJiraIssueOk() (*string, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *OwnerCreateDto) SetJiraIssue(v string)`

SetJiraIssue sets JiraIssue field to given value.


### GetDisplayName

`func (o *OwnerCreateDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OwnerCreateDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OwnerCreateDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OwnerCreateDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLinks

`func (o *OwnerCreateDto) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OwnerCreateDto) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OwnerCreateDto) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OwnerCreateDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



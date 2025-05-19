# ServiceCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | The alias of the service owner. Note, an update with changed owner will move the service and any associated repositories to the new owner, but of course this will not move e.g. Jenkins jobs. That&#39;s your job. | 
**Description** | Pointer to **string** | A short description of the functionality of the service. | [optional] 
**Quicklinks** | [**[]Quicklink**](Quicklink.md) | A list of quicklinks related to the service | 
**Repositories** | **[]string** | The keys of repositories associated with the service. When sending an update, they must refer to repositories that belong to this service, or the update will fail | 
**AlertTarget** | **string** | The default channel used to send any alerts of the service to. Can be an email address or a Teams webhook URL | 
**OperationType** | Pointer to **string** | The operation type of the service. &#39;WORKLOAD&#39; follows the default deployment strategy of one instance per environment, &#39;PLATFORM&#39; one instance per cluster or node and &#39;APPLICATION&#39; is a standalone application that is not deployed via the common strategies. | [optional] [default to "WORKLOAD"]
**InternetExposed** | Pointer to **bool** | The value defines if the service is available from the internet and the time period in which security holes must be processed. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Spec** | Pointer to [**ServiceSpecDto**](ServiceSpecDto.md) |  | [optional] 
**PostPromotes** | Pointer to [**PostPromote**](PostPromote.md) | Post promote dependencies. | [optional] 
**JiraIssue** | **string** | The jira issue to use for committing a change, or the last jira issue used. | 

## Methods

### NewServiceCreateDto

`func NewServiceCreateDto(owner string, quicklinks []Quicklink, repositories []string, alertTarget string, jiraIssue string, ) *ServiceCreateDto`

NewServiceCreateDto instantiates a new ServiceCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCreateDtoWithDefaults

`func NewServiceCreateDtoWithDefaults() *ServiceCreateDto`

NewServiceCreateDtoWithDefaults instantiates a new ServiceCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *ServiceCreateDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ServiceCreateDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ServiceCreateDto) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetDescription

`func (o *ServiceCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuicklinks

`func (o *ServiceCreateDto) GetQuicklinks() []Quicklink`

GetQuicklinks returns the Quicklinks field if non-nil, zero value otherwise.

### GetQuicklinksOk

`func (o *ServiceCreateDto) GetQuicklinksOk() (*[]Quicklink, bool)`

GetQuicklinksOk returns a tuple with the Quicklinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuicklinks

`func (o *ServiceCreateDto) SetQuicklinks(v []Quicklink)`

SetQuicklinks sets Quicklinks field to given value.


### GetRepositories

`func (o *ServiceCreateDto) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ServiceCreateDto) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ServiceCreateDto) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.


### GetAlertTarget

`func (o *ServiceCreateDto) GetAlertTarget() string`

GetAlertTarget returns the AlertTarget field if non-nil, zero value otherwise.

### GetAlertTargetOk

`func (o *ServiceCreateDto) GetAlertTargetOk() (*string, bool)`

GetAlertTargetOk returns a tuple with the AlertTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTarget

`func (o *ServiceCreateDto) SetAlertTarget(v string)`

SetAlertTarget sets AlertTarget field to given value.


### GetOperationType

`func (o *ServiceCreateDto) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ServiceCreateDto) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ServiceCreateDto) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *ServiceCreateDto) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetInternetExposed

`func (o *ServiceCreateDto) GetInternetExposed() bool`

GetInternetExposed returns the InternetExposed field if non-nil, zero value otherwise.

### GetInternetExposedOk

`func (o *ServiceCreateDto) GetInternetExposedOk() (*bool, bool)`

GetInternetExposedOk returns a tuple with the InternetExposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetExposed

`func (o *ServiceCreateDto) SetInternetExposed(v bool)`

SetInternetExposed sets InternetExposed field to given value.

### HasInternetExposed

`func (o *ServiceCreateDto) HasInternetExposed() bool`

HasInternetExposed returns a boolean if a field has been set.

### GetTags

`func (o *ServiceCreateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceCreateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceCreateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceCreateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLabels

`func (o *ServiceCreateDto) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ServiceCreateDto) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ServiceCreateDto) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ServiceCreateDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSpec

`func (o *ServiceCreateDto) GetSpec() ServiceSpecDto`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ServiceCreateDto) GetSpecOk() (*ServiceSpecDto, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ServiceCreateDto) SetSpec(v ServiceSpecDto)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ServiceCreateDto) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetPostPromotes

`func (o *ServiceCreateDto) GetPostPromotes() PostPromote`

GetPostPromotes returns the PostPromotes field if non-nil, zero value otherwise.

### GetPostPromotesOk

`func (o *ServiceCreateDto) GetPostPromotesOk() (*PostPromote, bool)`

GetPostPromotesOk returns a tuple with the PostPromotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPromotes

`func (o *ServiceCreateDto) SetPostPromotes(v PostPromote)`

SetPostPromotes sets PostPromotes field to given value.

### HasPostPromotes

`func (o *ServiceCreateDto) HasPostPromotes() bool`

HasPostPromotes returns a boolean if a field has been set.

### GetJiraIssue

`func (o *ServiceCreateDto) GetJiraIssue() string`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *ServiceCreateDto) GetJiraIssueOk() (*string, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *ServiceCreateDto) SetJiraIssue(v string)`

SetJiraIssue sets JiraIssue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



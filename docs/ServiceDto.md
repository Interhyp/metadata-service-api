# ServiceDto

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
**TimeStamp** | **string** | ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates. | 
**CommitHash** | **string** | The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates. | 
**JiraIssue** | **string** | The jira issue to use for committing a change, or the last jira issue used. | 
**Lifecycle** | Pointer to **string** | The current phase of the service&#39;s development. A service usually starts off as &#39;experimental&#39;, then becomes &#39;operational&#39; (i. e. can be reliably used and/or consumed). Once &#39;deprecated&#39;, the service doesn’t guarantee reliable use/consumption any longer and if &#39;decommissionable&#39;, the service will soon cease to exist. | [optional] 

## Methods

### NewServiceDto

`func NewServiceDto(owner string, quicklinks []Quicklink, repositories []string, alertTarget string, timeStamp string, commitHash string, jiraIssue string, ) *ServiceDto`

NewServiceDto instantiates a new ServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDtoWithDefaults

`func NewServiceDtoWithDefaults() *ServiceDto`

NewServiceDtoWithDefaults instantiates a new ServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *ServiceDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ServiceDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ServiceDto) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetDescription

`func (o *ServiceDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuicklinks

`func (o *ServiceDto) GetQuicklinks() []Quicklink`

GetQuicklinks returns the Quicklinks field if non-nil, zero value otherwise.

### GetQuicklinksOk

`func (o *ServiceDto) GetQuicklinksOk() (*[]Quicklink, bool)`

GetQuicklinksOk returns a tuple with the Quicklinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuicklinks

`func (o *ServiceDto) SetQuicklinks(v []Quicklink)`

SetQuicklinks sets Quicklinks field to given value.


### GetRepositories

`func (o *ServiceDto) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ServiceDto) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ServiceDto) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.


### GetAlertTarget

`func (o *ServiceDto) GetAlertTarget() string`

GetAlertTarget returns the AlertTarget field if non-nil, zero value otherwise.

### GetAlertTargetOk

`func (o *ServiceDto) GetAlertTargetOk() (*string, bool)`

GetAlertTargetOk returns a tuple with the AlertTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTarget

`func (o *ServiceDto) SetAlertTarget(v string)`

SetAlertTarget sets AlertTarget field to given value.


### GetOperationType

`func (o *ServiceDto) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ServiceDto) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ServiceDto) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *ServiceDto) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetInternetExposed

`func (o *ServiceDto) GetInternetExposed() bool`

GetInternetExposed returns the InternetExposed field if non-nil, zero value otherwise.

### GetInternetExposedOk

`func (o *ServiceDto) GetInternetExposedOk() (*bool, bool)`

GetInternetExposedOk returns a tuple with the InternetExposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetExposed

`func (o *ServiceDto) SetInternetExposed(v bool)`

SetInternetExposed sets InternetExposed field to given value.

### HasInternetExposed

`func (o *ServiceDto) HasInternetExposed() bool`

HasInternetExposed returns a boolean if a field has been set.

### GetTags

`func (o *ServiceDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLabels

`func (o *ServiceDto) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ServiceDto) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ServiceDto) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ServiceDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSpec

`func (o *ServiceDto) GetSpec() ServiceSpecDto`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ServiceDto) GetSpecOk() (*ServiceSpecDto, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ServiceDto) SetSpec(v ServiceSpecDto)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ServiceDto) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetPostPromotes

`func (o *ServiceDto) GetPostPromotes() PostPromote`

GetPostPromotes returns the PostPromotes field if non-nil, zero value otherwise.

### GetPostPromotesOk

`func (o *ServiceDto) GetPostPromotesOk() (*PostPromote, bool)`

GetPostPromotesOk returns a tuple with the PostPromotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPromotes

`func (o *ServiceDto) SetPostPromotes(v PostPromote)`

SetPostPromotes sets PostPromotes field to given value.

### HasPostPromotes

`func (o *ServiceDto) HasPostPromotes() bool`

HasPostPromotes returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ServiceDto) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ServiceDto) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ServiceDto) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetCommitHash

`func (o *ServiceDto) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ServiceDto) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ServiceDto) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### GetJiraIssue

`func (o *ServiceDto) GetJiraIssue() string`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *ServiceDto) GetJiraIssueOk() (*string, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *ServiceDto) SetJiraIssue(v string)`

SetJiraIssue sets JiraIssue field to given value.


### GetLifecycle

`func (o *ServiceDto) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ServiceDto) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ServiceDto) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ServiceDto) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



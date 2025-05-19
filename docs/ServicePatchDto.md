# ServicePatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **string** | The alias of the service owner. Note, a patch with changed owner will move the service and any associated repositories to the new owner, but of course this will not move e.g. Jenkins jobs. That&#39;s your job. | [optional] 
**Description** | Pointer to **string** | A short description of the functionality of the service. | [optional] 
**Quicklinks** | Pointer to [**[]Quicklink**](Quicklink.md) | A list of quicklinks related to the service | [optional] 
**Repositories** | Pointer to **[]string** | The keys of repositories associated with the service. When sending an update, they must refer to repositories that belong to this service, or the update will fail | [optional] 
**AlertTarget** | Pointer to **string** | The default channel used to send any alerts of the service to. Can be an email address or a Teams webhook URL | [optional] 
**OperationType** | Pointer to **string** | The operation type of the service. &#39;WORKLOAD&#39; follows the default deployment strategy of one instance per environment, &#39;PLATFORM&#39; one instance per cluster or node and &#39;APPLICATION&#39; is a standalone application that is not deployed via the common strategies. | [optional] [default to "WORKLOAD"]
**InternetExposed** | Pointer to **bool** | The value defines if the service is available from the internet and the time period in which security holes must be processed. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Spec** | Pointer to [**ServiceSpecDto**](ServiceSpecDto.md) |  | [optional] 
**PostPromotes** | Pointer to [**PostPromote**](PostPromote.md) | Post promote dependencies. | [optional] 
**TimeStamp** | **string** | ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates. | 
**CommitHash** | **string** | The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates. | 
**JiraIssue** | **string** | The jira issue to use for committing a change, or the last jira issue used. | 
**Lifecycle** | Pointer to **string** | The current phase of the service&#39;s development. A service usually starts off as &#39;experimental&#39;, then becomes &#39;operational&#39; (i. e. can be reliably used and/or consumed). Once &#39;deprecated&#39;, the service doesn’t guarantee reliable use/consumption any longer. | [optional] 

## Methods

### NewServicePatchDto

`func NewServicePatchDto(timeStamp string, commitHash string, jiraIssue string, ) *ServicePatchDto`

NewServicePatchDto instantiates a new ServicePatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePatchDtoWithDefaults

`func NewServicePatchDtoWithDefaults() *ServicePatchDto`

NewServicePatchDtoWithDefaults instantiates a new ServicePatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *ServicePatchDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ServicePatchDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ServicePatchDto) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ServicePatchDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDescription

`func (o *ServicePatchDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePatchDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePatchDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePatchDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuicklinks

`func (o *ServicePatchDto) GetQuicklinks() []Quicklink`

GetQuicklinks returns the Quicklinks field if non-nil, zero value otherwise.

### GetQuicklinksOk

`func (o *ServicePatchDto) GetQuicklinksOk() (*[]Quicklink, bool)`

GetQuicklinksOk returns a tuple with the Quicklinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuicklinks

`func (o *ServicePatchDto) SetQuicklinks(v []Quicklink)`

SetQuicklinks sets Quicklinks field to given value.

### HasQuicklinks

`func (o *ServicePatchDto) HasQuicklinks() bool`

HasQuicklinks returns a boolean if a field has been set.

### GetRepositories

`func (o *ServicePatchDto) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ServicePatchDto) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ServicePatchDto) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *ServicePatchDto) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetAlertTarget

`func (o *ServicePatchDto) GetAlertTarget() string`

GetAlertTarget returns the AlertTarget field if non-nil, zero value otherwise.

### GetAlertTargetOk

`func (o *ServicePatchDto) GetAlertTargetOk() (*string, bool)`

GetAlertTargetOk returns a tuple with the AlertTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTarget

`func (o *ServicePatchDto) SetAlertTarget(v string)`

SetAlertTarget sets AlertTarget field to given value.

### HasAlertTarget

`func (o *ServicePatchDto) HasAlertTarget() bool`

HasAlertTarget returns a boolean if a field has been set.

### GetOperationType

`func (o *ServicePatchDto) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ServicePatchDto) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ServicePatchDto) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *ServicePatchDto) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetInternetExposed

`func (o *ServicePatchDto) GetInternetExposed() bool`

GetInternetExposed returns the InternetExposed field if non-nil, zero value otherwise.

### GetInternetExposedOk

`func (o *ServicePatchDto) GetInternetExposedOk() (*bool, bool)`

GetInternetExposedOk returns a tuple with the InternetExposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetExposed

`func (o *ServicePatchDto) SetInternetExposed(v bool)`

SetInternetExposed sets InternetExposed field to given value.

### HasInternetExposed

`func (o *ServicePatchDto) HasInternetExposed() bool`

HasInternetExposed returns a boolean if a field has been set.

### GetTags

`func (o *ServicePatchDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServicePatchDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServicePatchDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServicePatchDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLabels

`func (o *ServicePatchDto) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ServicePatchDto) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ServicePatchDto) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ServicePatchDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSpec

`func (o *ServicePatchDto) GetSpec() ServiceSpecDto`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ServicePatchDto) GetSpecOk() (*ServiceSpecDto, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ServicePatchDto) SetSpec(v ServiceSpecDto)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ServicePatchDto) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetPostPromotes

`func (o *ServicePatchDto) GetPostPromotes() PostPromote`

GetPostPromotes returns the PostPromotes field if non-nil, zero value otherwise.

### GetPostPromotesOk

`func (o *ServicePatchDto) GetPostPromotesOk() (*PostPromote, bool)`

GetPostPromotesOk returns a tuple with the PostPromotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPromotes

`func (o *ServicePatchDto) SetPostPromotes(v PostPromote)`

SetPostPromotes sets PostPromotes field to given value.

### HasPostPromotes

`func (o *ServicePatchDto) HasPostPromotes() bool`

HasPostPromotes returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ServicePatchDto) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ServicePatchDto) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ServicePatchDto) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetCommitHash

`func (o *ServicePatchDto) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ServicePatchDto) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ServicePatchDto) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### GetJiraIssue

`func (o *ServicePatchDto) GetJiraIssue() string`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *ServicePatchDto) GetJiraIssueOk() (*string, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *ServicePatchDto) SetJiraIssue(v string)`

SetJiraIssue sets JiraIssue field to given value.


### GetLifecycle

`func (o *ServicePatchDto) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ServicePatchDto) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ServicePatchDto) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ServicePatchDto) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



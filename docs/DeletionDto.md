# DeletionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JiraIssue** | **string** | The jira issue to use for committing the deletion. | 

## Methods

### NewDeletionDto

`func NewDeletionDto(jiraIssue string, ) *DeletionDto`

NewDeletionDto instantiates a new DeletionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletionDtoWithDefaults

`func NewDeletionDtoWithDefaults() *DeletionDto`

NewDeletionDtoWithDefaults instantiates a new DeletionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJiraIssue

`func (o *DeletionDto) GetJiraIssue() string`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *DeletionDto) GetJiraIssueOk() (*string, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *DeletionDto) SetJiraIssue(v string)`

SetJiraIssue sets JiraIssue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



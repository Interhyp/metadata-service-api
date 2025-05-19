# PullRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMergeCommits** | Pointer to **bool** | Allows merge commits on pull requests | [optional] [default to true]
**AllowRebaseMerging** | Pointer to **bool** | Allows rebase merging on pull requests | [optional] [default to true]

## Methods

### NewPullRequests

`func NewPullRequests() *PullRequests`

NewPullRequests instantiates a new PullRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestsWithDefaults

`func NewPullRequestsWithDefaults() *PullRequests`

NewPullRequestsWithDefaults instantiates a new PullRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMergeCommits

`func (o *PullRequests) GetAllowMergeCommits() bool`

GetAllowMergeCommits returns the AllowMergeCommits field if non-nil, zero value otherwise.

### GetAllowMergeCommitsOk

`func (o *PullRequests) GetAllowMergeCommitsOk() (*bool, bool)`

GetAllowMergeCommitsOk returns a tuple with the AllowMergeCommits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMergeCommits

`func (o *PullRequests) SetAllowMergeCommits(v bool)`

SetAllowMergeCommits sets AllowMergeCommits field to given value.

### HasAllowMergeCommits

`func (o *PullRequests) HasAllowMergeCommits() bool`

HasAllowMergeCommits returns a boolean if a field has been set.

### GetAllowRebaseMerging

`func (o *PullRequests) GetAllowRebaseMerging() bool`

GetAllowRebaseMerging returns the AllowRebaseMerging field if non-nil, zero value otherwise.

### GetAllowRebaseMergingOk

`func (o *PullRequests) GetAllowRebaseMergingOk() (*bool, bool)`

GetAllowRebaseMergingOk returns a tuple with the AllowRebaseMerging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRebaseMerging

`func (o *PullRequests) SetAllowRebaseMerging(v bool)`

SetAllowRebaseMerging sets AllowRebaseMerging field to given value.

### HasAllowRebaseMerging

`func (o *PullRequests) HasAllowRebaseMerging() bool`

HasAllowRebaseMerging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



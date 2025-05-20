# RepositoryConfigurationPatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeys** | Pointer to [**[]RepositoryConfigurationAccessKeyDto**](RepositoryConfigurationAccessKeyDto.md) | Ssh-Keys configured on the repository. | [optional] 
**MergeConfig** | Pointer to [**RepositoryConfigurationDtoMergeConfig**](RepositoryConfigurationDtoMergeConfig.md) |  | [optional] 
**DefaultTasks** | Pointer to [**[]RepositoryConfigurationDefaultTaskDto**](RepositoryConfigurationDefaultTaskDto.md) |  | [optional] 
**BranchNameRegex** | Pointer to **string** | Use an explicit branch name regex. | [optional] 
**CommitMessageRegex** | Pointer to **string** | Use an explicit commit message regex. | [optional] 
**CommitMessageType** | Pointer to **string** | Adds a corresponding commit message regex. | [optional] 
**RequireSuccessfulBuilds** | Pointer to **int32** | Set the required successful builds counter. | [optional] 
**RequireApprovals** | Pointer to **int32** | Set the required approvals counter. | [optional] 
**ExcludeMergeCommits** | Pointer to **bool** | Exclude merge commits from commit checks. | [optional] 
**ExcludeMergeCheckUsers** | Pointer to [**[]ExcludeMergeCheckUserDto**](ExcludeMergeCheckUserDto.md) | Exclude users from commit checks. | [optional] 
**Webhooks** | Pointer to [**RepositoryConfigurationWebhooksDto**](RepositoryConfigurationWebhooksDto.md) |  | [optional] 
**Approvers** | Pointer to **map[string][]string** | Map of string (group name e.g. some-owner) of strings (list of approvers), one approval for each group is required. | [optional] 
**Watchers** | Pointer to **[]string** | List of strings (list of watchers, either usernames or group identifier), which are added as reviewers but require no approval. | [optional] 
**Archived** | Pointer to **bool** | Moves the repository into the archive. | [optional] 
**Unmanaged** | Pointer to **bool** | Repository will not be configured, also not archived. | [optional] 
**RefProtections** | Pointer to [**RefProtections**](RefProtections.md) |  | [optional] 
**RequireIssue** | Pointer to **bool** | Configures JQL matcher with query: issuetype in (Story, Bug) AND &#39;Risk Level&#39; is not EMPTY | [optional] 
**RequireConditions** | Pointer to [**map[string]ConditionReferenceDto**](ConditionReferenceDto.md) | Configuration of conditional builds as map of structs (key name e.g. some-key) of target references. | [optional] 
**ActionsAccess** | Pointer to **string** | Control how the repository is used by GitHub Actions workflows in other repositories | [optional] 
**PullRequests** | Pointer to [**PullRequests**](PullRequests.md) |  | [optional] 

## Methods

### NewRepositoryConfigurationPatchDto

`func NewRepositoryConfigurationPatchDto() *RepositoryConfigurationPatchDto`

NewRepositoryConfigurationPatchDto instantiates a new RepositoryConfigurationPatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryConfigurationPatchDtoWithDefaults

`func NewRepositoryConfigurationPatchDtoWithDefaults() *RepositoryConfigurationPatchDto`

NewRepositoryConfigurationPatchDtoWithDefaults instantiates a new RepositoryConfigurationPatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeys

`func (o *RepositoryConfigurationPatchDto) GetAccessKeys() []RepositoryConfigurationAccessKeyDto`

GetAccessKeys returns the AccessKeys field if non-nil, zero value otherwise.

### GetAccessKeysOk

`func (o *RepositoryConfigurationPatchDto) GetAccessKeysOk() (*[]RepositoryConfigurationAccessKeyDto, bool)`

GetAccessKeysOk returns a tuple with the AccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeys

`func (o *RepositoryConfigurationPatchDto) SetAccessKeys(v []RepositoryConfigurationAccessKeyDto)`

SetAccessKeys sets AccessKeys field to given value.

### HasAccessKeys

`func (o *RepositoryConfigurationPatchDto) HasAccessKeys() bool`

HasAccessKeys returns a boolean if a field has been set.

### GetMergeConfig

`func (o *RepositoryConfigurationPatchDto) GetMergeConfig() RepositoryConfigurationDtoMergeConfig`

GetMergeConfig returns the MergeConfig field if non-nil, zero value otherwise.

### GetMergeConfigOk

`func (o *RepositoryConfigurationPatchDto) GetMergeConfigOk() (*RepositoryConfigurationDtoMergeConfig, bool)`

GetMergeConfigOk returns a tuple with the MergeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeConfig

`func (o *RepositoryConfigurationPatchDto) SetMergeConfig(v RepositoryConfigurationDtoMergeConfig)`

SetMergeConfig sets MergeConfig field to given value.

### HasMergeConfig

`func (o *RepositoryConfigurationPatchDto) HasMergeConfig() bool`

HasMergeConfig returns a boolean if a field has been set.

### GetDefaultTasks

`func (o *RepositoryConfigurationPatchDto) GetDefaultTasks() []RepositoryConfigurationDefaultTaskDto`

GetDefaultTasks returns the DefaultTasks field if non-nil, zero value otherwise.

### GetDefaultTasksOk

`func (o *RepositoryConfigurationPatchDto) GetDefaultTasksOk() (*[]RepositoryConfigurationDefaultTaskDto, bool)`

GetDefaultTasksOk returns a tuple with the DefaultTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTasks

`func (o *RepositoryConfigurationPatchDto) SetDefaultTasks(v []RepositoryConfigurationDefaultTaskDto)`

SetDefaultTasks sets DefaultTasks field to given value.

### HasDefaultTasks

`func (o *RepositoryConfigurationPatchDto) HasDefaultTasks() bool`

HasDefaultTasks returns a boolean if a field has been set.

### GetBranchNameRegex

`func (o *RepositoryConfigurationPatchDto) GetBranchNameRegex() string`

GetBranchNameRegex returns the BranchNameRegex field if non-nil, zero value otherwise.

### GetBranchNameRegexOk

`func (o *RepositoryConfigurationPatchDto) GetBranchNameRegexOk() (*string, bool)`

GetBranchNameRegexOk returns a tuple with the BranchNameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchNameRegex

`func (o *RepositoryConfigurationPatchDto) SetBranchNameRegex(v string)`

SetBranchNameRegex sets BranchNameRegex field to given value.

### HasBranchNameRegex

`func (o *RepositoryConfigurationPatchDto) HasBranchNameRegex() bool`

HasBranchNameRegex returns a boolean if a field has been set.

### GetCommitMessageRegex

`func (o *RepositoryConfigurationPatchDto) GetCommitMessageRegex() string`

GetCommitMessageRegex returns the CommitMessageRegex field if non-nil, zero value otherwise.

### GetCommitMessageRegexOk

`func (o *RepositoryConfigurationPatchDto) GetCommitMessageRegexOk() (*string, bool)`

GetCommitMessageRegexOk returns a tuple with the CommitMessageRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessageRegex

`func (o *RepositoryConfigurationPatchDto) SetCommitMessageRegex(v string)`

SetCommitMessageRegex sets CommitMessageRegex field to given value.

### HasCommitMessageRegex

`func (o *RepositoryConfigurationPatchDto) HasCommitMessageRegex() bool`

HasCommitMessageRegex returns a boolean if a field has been set.

### GetCommitMessageType

`func (o *RepositoryConfigurationPatchDto) GetCommitMessageType() string`

GetCommitMessageType returns the CommitMessageType field if non-nil, zero value otherwise.

### GetCommitMessageTypeOk

`func (o *RepositoryConfigurationPatchDto) GetCommitMessageTypeOk() (*string, bool)`

GetCommitMessageTypeOk returns a tuple with the CommitMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessageType

`func (o *RepositoryConfigurationPatchDto) SetCommitMessageType(v string)`

SetCommitMessageType sets CommitMessageType field to given value.

### HasCommitMessageType

`func (o *RepositoryConfigurationPatchDto) HasCommitMessageType() bool`

HasCommitMessageType returns a boolean if a field has been set.

### GetRequireSuccessfulBuilds

`func (o *RepositoryConfigurationPatchDto) GetRequireSuccessfulBuilds() int32`

GetRequireSuccessfulBuilds returns the RequireSuccessfulBuilds field if non-nil, zero value otherwise.

### GetRequireSuccessfulBuildsOk

`func (o *RepositoryConfigurationPatchDto) GetRequireSuccessfulBuildsOk() (*int32, bool)`

GetRequireSuccessfulBuildsOk returns a tuple with the RequireSuccessfulBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSuccessfulBuilds

`func (o *RepositoryConfigurationPatchDto) SetRequireSuccessfulBuilds(v int32)`

SetRequireSuccessfulBuilds sets RequireSuccessfulBuilds field to given value.

### HasRequireSuccessfulBuilds

`func (o *RepositoryConfigurationPatchDto) HasRequireSuccessfulBuilds() bool`

HasRequireSuccessfulBuilds returns a boolean if a field has been set.

### GetRequireApprovals

`func (o *RepositoryConfigurationPatchDto) GetRequireApprovals() int32`

GetRequireApprovals returns the RequireApprovals field if non-nil, zero value otherwise.

### GetRequireApprovalsOk

`func (o *RepositoryConfigurationPatchDto) GetRequireApprovalsOk() (*int32, bool)`

GetRequireApprovalsOk returns a tuple with the RequireApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireApprovals

`func (o *RepositoryConfigurationPatchDto) SetRequireApprovals(v int32)`

SetRequireApprovals sets RequireApprovals field to given value.

### HasRequireApprovals

`func (o *RepositoryConfigurationPatchDto) HasRequireApprovals() bool`

HasRequireApprovals returns a boolean if a field has been set.

### GetExcludeMergeCommits

`func (o *RepositoryConfigurationPatchDto) GetExcludeMergeCommits() bool`

GetExcludeMergeCommits returns the ExcludeMergeCommits field if non-nil, zero value otherwise.

### GetExcludeMergeCommitsOk

`func (o *RepositoryConfigurationPatchDto) GetExcludeMergeCommitsOk() (*bool, bool)`

GetExcludeMergeCommitsOk returns a tuple with the ExcludeMergeCommits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMergeCommits

`func (o *RepositoryConfigurationPatchDto) SetExcludeMergeCommits(v bool)`

SetExcludeMergeCommits sets ExcludeMergeCommits field to given value.

### HasExcludeMergeCommits

`func (o *RepositoryConfigurationPatchDto) HasExcludeMergeCommits() bool`

HasExcludeMergeCommits returns a boolean if a field has been set.

### GetExcludeMergeCheckUsers

`func (o *RepositoryConfigurationPatchDto) GetExcludeMergeCheckUsers() []ExcludeMergeCheckUserDto`

GetExcludeMergeCheckUsers returns the ExcludeMergeCheckUsers field if non-nil, zero value otherwise.

### GetExcludeMergeCheckUsersOk

`func (o *RepositoryConfigurationPatchDto) GetExcludeMergeCheckUsersOk() (*[]ExcludeMergeCheckUserDto, bool)`

GetExcludeMergeCheckUsersOk returns a tuple with the ExcludeMergeCheckUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMergeCheckUsers

`func (o *RepositoryConfigurationPatchDto) SetExcludeMergeCheckUsers(v []ExcludeMergeCheckUserDto)`

SetExcludeMergeCheckUsers sets ExcludeMergeCheckUsers field to given value.

### HasExcludeMergeCheckUsers

`func (o *RepositoryConfigurationPatchDto) HasExcludeMergeCheckUsers() bool`

HasExcludeMergeCheckUsers returns a boolean if a field has been set.

### GetWebhooks

`func (o *RepositoryConfigurationPatchDto) GetWebhooks() RepositoryConfigurationWebhooksDto`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *RepositoryConfigurationPatchDto) GetWebhooksOk() (*RepositoryConfigurationWebhooksDto, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *RepositoryConfigurationPatchDto) SetWebhooks(v RepositoryConfigurationWebhooksDto)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *RepositoryConfigurationPatchDto) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetApprovers

`func (o *RepositoryConfigurationPatchDto) GetApprovers() map[string][]string`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *RepositoryConfigurationPatchDto) GetApproversOk() (*map[string][]string, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *RepositoryConfigurationPatchDto) SetApprovers(v map[string][]string)`

SetApprovers sets Approvers field to given value.

### HasApprovers

`func (o *RepositoryConfigurationPatchDto) HasApprovers() bool`

HasApprovers returns a boolean if a field has been set.

### GetWatchers

`func (o *RepositoryConfigurationPatchDto) GetWatchers() []string`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *RepositoryConfigurationPatchDto) GetWatchersOk() (*[]string, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *RepositoryConfigurationPatchDto) SetWatchers(v []string)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *RepositoryConfigurationPatchDto) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.

### GetArchived

`func (o *RepositoryConfigurationPatchDto) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *RepositoryConfigurationPatchDto) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *RepositoryConfigurationPatchDto) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *RepositoryConfigurationPatchDto) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetUnmanaged

`func (o *RepositoryConfigurationPatchDto) GetUnmanaged() bool`

GetUnmanaged returns the Unmanaged field if non-nil, zero value otherwise.

### GetUnmanagedOk

`func (o *RepositoryConfigurationPatchDto) GetUnmanagedOk() (*bool, bool)`

GetUnmanagedOk returns a tuple with the Unmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanaged

`func (o *RepositoryConfigurationPatchDto) SetUnmanaged(v bool)`

SetUnmanaged sets Unmanaged field to given value.

### HasUnmanaged

`func (o *RepositoryConfigurationPatchDto) HasUnmanaged() bool`

HasUnmanaged returns a boolean if a field has been set.

### GetRefProtections

`func (o *RepositoryConfigurationPatchDto) GetRefProtections() RefProtections`

GetRefProtections returns the RefProtections field if non-nil, zero value otherwise.

### GetRefProtectionsOk

`func (o *RepositoryConfigurationPatchDto) GetRefProtectionsOk() (*RefProtections, bool)`

GetRefProtectionsOk returns a tuple with the RefProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProtections

`func (o *RepositoryConfigurationPatchDto) SetRefProtections(v RefProtections)`

SetRefProtections sets RefProtections field to given value.

### HasRefProtections

`func (o *RepositoryConfigurationPatchDto) HasRefProtections() bool`

HasRefProtections returns a boolean if a field has been set.

### GetRequireIssue

`func (o *RepositoryConfigurationPatchDto) GetRequireIssue() bool`

GetRequireIssue returns the RequireIssue field if non-nil, zero value otherwise.

### GetRequireIssueOk

`func (o *RepositoryConfigurationPatchDto) GetRequireIssueOk() (*bool, bool)`

GetRequireIssueOk returns a tuple with the RequireIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireIssue

`func (o *RepositoryConfigurationPatchDto) SetRequireIssue(v bool)`

SetRequireIssue sets RequireIssue field to given value.

### HasRequireIssue

`func (o *RepositoryConfigurationPatchDto) HasRequireIssue() bool`

HasRequireIssue returns a boolean if a field has been set.

### GetRequireConditions

`func (o *RepositoryConfigurationPatchDto) GetRequireConditions() map[string]ConditionReferenceDto`

GetRequireConditions returns the RequireConditions field if non-nil, zero value otherwise.

### GetRequireConditionsOk

`func (o *RepositoryConfigurationPatchDto) GetRequireConditionsOk() (*map[string]ConditionReferenceDto, bool)`

GetRequireConditionsOk returns a tuple with the RequireConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireConditions

`func (o *RepositoryConfigurationPatchDto) SetRequireConditions(v map[string]ConditionReferenceDto)`

SetRequireConditions sets RequireConditions field to given value.

### HasRequireConditions

`func (o *RepositoryConfigurationPatchDto) HasRequireConditions() bool`

HasRequireConditions returns a boolean if a field has been set.

### GetActionsAccess

`func (o *RepositoryConfigurationPatchDto) GetActionsAccess() string`

GetActionsAccess returns the ActionsAccess field if non-nil, zero value otherwise.

### GetActionsAccessOk

`func (o *RepositoryConfigurationPatchDto) GetActionsAccessOk() (*string, bool)`

GetActionsAccessOk returns a tuple with the ActionsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsAccess

`func (o *RepositoryConfigurationPatchDto) SetActionsAccess(v string)`

SetActionsAccess sets ActionsAccess field to given value.

### HasActionsAccess

`func (o *RepositoryConfigurationPatchDto) HasActionsAccess() bool`

HasActionsAccess returns a boolean if a field has been set.

### GetPullRequests

`func (o *RepositoryConfigurationPatchDto) GetPullRequests() PullRequests`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *RepositoryConfigurationPatchDto) GetPullRequestsOk() (*PullRequests, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *RepositoryConfigurationPatchDto) SetPullRequests(v PullRequests)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *RepositoryConfigurationPatchDto) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



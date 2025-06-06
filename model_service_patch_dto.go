/*
Metadata

Obtain and manage metadata for owners, services, repositories. Please see [README](https://github.com/Interhyp/metadata-service/blob/main/README.md) for details. **CLIENTS MUST READ!**

API version: v1
Contact: somebody@some-organisation.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ServicePatchDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePatchDto{}

// ServicePatchDto struct for ServicePatchDto
type ServicePatchDto struct {
	// The alias of the service owner. Note, a patch with changed owner will move the service and any associated repositories to the new owner, but of course this will not move e.g. Jenkins jobs. That's your job.
	Owner *string `json:"owner,omitempty"`
	// A short description of the functionality of the service.
	Description *string `json:"description,omitempty"`
	// A list of quicklinks related to the service
	Quicklinks []Quicklink `json:"quicklinks,omitempty"`
	// The keys of repositories associated with the service. When sending an update, they must refer to repositories that belong to this service, or the update will fail
	Repositories []string `json:"repositories,omitempty"`
	// The default channel used to send any alerts of the service to. Can be an email address or a Teams webhook URL
	AlertTarget *string `json:"alertTarget,omitempty"`
	// The operation type of the service. 'WORKLOAD' follows the default deployment strategy of one instance per environment, 'PLATFORM' one instance per cluster or node and 'APPLICATION' is a standalone application that is not deployed via the common strategies.
	OperationType *string `json:"operationType,omitempty"`
	// The value defines if the service is available from the internet and the time period in which security holes must be processed.
	InternetExposed *bool `json:"internetExposed,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	Spec *ServiceSpecDto `json:"spec,omitempty"`
	// Post promote dependencies.
	PostPromotes *PostPromote `json:"postPromotes,omitempty"`
	// ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates.
	TimeStamp string `json:"timeStamp" yaml:"-"`
	// The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates.
	CommitHash string `json:"commitHash" yaml:"-"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue" yaml:"-"`
	// The current phase of the service's development. A service usually starts off as 'experimental', then becomes 'operational' (i. e. can be reliably used and/or consumed). Once 'deprecated', the service doesn’t guarantee reliable use/consumption any longer.
	Lifecycle *string `json:"lifecycle,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicePatchDto ServicePatchDto

// NewServicePatchDto instantiates a new ServicePatchDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePatchDto(timeStamp string, commitHash string, jiraIssue string) *ServicePatchDto {
	this := ServicePatchDto{}
	var operationType string = "WORKLOAD"
	this.OperationType = &operationType
	this.TimeStamp = timeStamp
	this.CommitHash = commitHash
	this.JiraIssue = jiraIssue
	return &this
}

// NewServicePatchDtoWithDefaults instantiates a new ServicePatchDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePatchDtoWithDefaults() *ServicePatchDto {
	this := ServicePatchDto{}
	var operationType string = "WORKLOAD"
	this.OperationType = &operationType
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ServicePatchDto) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ServicePatchDto) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ServicePatchDto) SetOwner(v string) {
	o.Owner = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServicePatchDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServicePatchDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServicePatchDto) SetDescription(v string) {
	o.Description = &v
}

// GetQuicklinks returns the Quicklinks field value if set, zero value otherwise.
func (o *ServicePatchDto) GetQuicklinks() []Quicklink {
	if o == nil || IsNil(o.Quicklinks) {
		var ret []Quicklink
		return ret
	}
	return o.Quicklinks
}

// GetQuicklinksOk returns a tuple with the Quicklinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetQuicklinksOk() ([]Quicklink, bool) {
	if o == nil || IsNil(o.Quicklinks) {
		return nil, false
	}
	return o.Quicklinks, true
}

// HasQuicklinks returns a boolean if a field has been set.
func (o *ServicePatchDto) HasQuicklinks() bool {
	if o != nil && !IsNil(o.Quicklinks) {
		return true
	}

	return false
}

// SetQuicklinks gets a reference to the given []Quicklink and assigns it to the Quicklinks field.
func (o *ServicePatchDto) SetQuicklinks(v []Quicklink) {
	o.Quicklinks = v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *ServicePatchDto) GetRepositories() []string {
	if o == nil || IsNil(o.Repositories) {
		var ret []string
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetRepositoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *ServicePatchDto) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []string and assigns it to the Repositories field.
func (o *ServicePatchDto) SetRepositories(v []string) {
	o.Repositories = v
}

// GetAlertTarget returns the AlertTarget field value if set, zero value otherwise.
func (o *ServicePatchDto) GetAlertTarget() string {
	if o == nil || IsNil(o.AlertTarget) {
		var ret string
		return ret
	}
	return *o.AlertTarget
}

// GetAlertTargetOk returns a tuple with the AlertTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetAlertTargetOk() (*string, bool) {
	if o == nil || IsNil(o.AlertTarget) {
		return nil, false
	}
	return o.AlertTarget, true
}

// HasAlertTarget returns a boolean if a field has been set.
func (o *ServicePatchDto) HasAlertTarget() bool {
	if o != nil && !IsNil(o.AlertTarget) {
		return true
	}

	return false
}

// SetAlertTarget gets a reference to the given string and assigns it to the AlertTarget field.
func (o *ServicePatchDto) SetAlertTarget(v string) {
	o.AlertTarget = &v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *ServicePatchDto) GetOperationType() string {
	if o == nil || IsNil(o.OperationType) {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetOperationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OperationType) {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *ServicePatchDto) HasOperationType() bool {
	if o != nil && !IsNil(o.OperationType) {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *ServicePatchDto) SetOperationType(v string) {
	o.OperationType = &v
}

// GetInternetExposed returns the InternetExposed field value if set, zero value otherwise.
func (o *ServicePatchDto) GetInternetExposed() bool {
	if o == nil || IsNil(o.InternetExposed) {
		var ret bool
		return ret
	}
	return *o.InternetExposed
}

// GetInternetExposedOk returns a tuple with the InternetExposed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetInternetExposedOk() (*bool, bool) {
	if o == nil || IsNil(o.InternetExposed) {
		return nil, false
	}
	return o.InternetExposed, true
}

// HasInternetExposed returns a boolean if a field has been set.
func (o *ServicePatchDto) HasInternetExposed() bool {
	if o != nil && !IsNil(o.InternetExposed) {
		return true
	}

	return false
}

// SetInternetExposed gets a reference to the given bool and assigns it to the InternetExposed field.
func (o *ServicePatchDto) SetInternetExposed(v bool) {
	o.InternetExposed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServicePatchDto) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServicePatchDto) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServicePatchDto) SetTags(v []string) {
	o.Tags = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ServicePatchDto) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetLabelsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return map[string]string{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ServicePatchDto) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ServicePatchDto) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ServicePatchDto) GetSpec() ServiceSpecDto {
	if o == nil || IsNil(o.Spec) {
		var ret ServiceSpecDto
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetSpecOk() (*ServiceSpecDto, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ServicePatchDto) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given ServiceSpecDto and assigns it to the Spec field.
func (o *ServicePatchDto) SetSpec(v ServiceSpecDto) {
	o.Spec = &v
}

// GetPostPromotes returns the PostPromotes field value if set, zero value otherwise.
func (o *ServicePatchDto) GetPostPromotes() PostPromote {
	if o == nil || IsNil(o.PostPromotes) {
		var ret PostPromote
		return ret
	}
	return *o.PostPromotes
}

// GetPostPromotesOk returns a tuple with the PostPromotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetPostPromotesOk() (*PostPromote, bool) {
	if o == nil || IsNil(o.PostPromotes) {
		return nil, false
	}
	return o.PostPromotes, true
}

// HasPostPromotes returns a boolean if a field has been set.
func (o *ServicePatchDto) HasPostPromotes() bool {
	if o != nil && !IsNil(o.PostPromotes) {
		return true
	}

	return false
}

// SetPostPromotes gets a reference to the given PostPromote and assigns it to the PostPromotes field.
func (o *ServicePatchDto) SetPostPromotes(v PostPromote) {
	o.PostPromotes = &v
}

// GetTimeStamp returns the TimeStamp field value
func (o *ServicePatchDto) GetTimeStamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetTimeStampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *ServicePatchDto) SetTimeStamp(v string) {
	o.TimeStamp = v
}

// GetCommitHash returns the CommitHash field value
func (o *ServicePatchDto) GetCommitHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitHash
}

// GetCommitHashOk returns a tuple with the CommitHash field value
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetCommitHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitHash, true
}

// SetCommitHash sets field value
func (o *ServicePatchDto) SetCommitHash(v string) {
	o.CommitHash = v
}

// GetJiraIssue returns the JiraIssue field value
func (o *ServicePatchDto) GetJiraIssue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JiraIssue
}

// GetJiraIssueOk returns a tuple with the JiraIssue field value
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetJiraIssueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JiraIssue, true
}

// SetJiraIssue sets field value
func (o *ServicePatchDto) SetJiraIssue(v string) {
	o.JiraIssue = v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *ServicePatchDto) GetLifecycle() string {
	if o == nil || IsNil(o.Lifecycle) {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePatchDto) GetLifecycleOk() (*string, bool) {
	if o == nil || IsNil(o.Lifecycle) {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *ServicePatchDto) HasLifecycle() bool {
	if o != nil && !IsNil(o.Lifecycle) {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *ServicePatchDto) SetLifecycle(v string) {
	o.Lifecycle = &v
}

func (o ServicePatchDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePatchDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Quicklinks) {
		toSerialize["quicklinks"] = o.Quicklinks
	}
	if !IsNil(o.Repositories) {
		toSerialize["repositories"] = o.Repositories
	}
	if !IsNil(o.AlertTarget) {
		toSerialize["alertTarget"] = o.AlertTarget
	}
	if !IsNil(o.OperationType) {
		toSerialize["operationType"] = o.OperationType
	}
	if !IsNil(o.InternetExposed) {
		toSerialize["internetExposed"] = o.InternetExposed
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if !IsNil(o.PostPromotes) {
		toSerialize["postPromotes"] = o.PostPromotes
	}
	toSerialize["timeStamp"] = o.TimeStamp
	toSerialize["commitHash"] = o.CommitHash
	toSerialize["jiraIssue"] = o.JiraIssue
	if !IsNil(o.Lifecycle) {
		toSerialize["lifecycle"] = o.Lifecycle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicePatchDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timeStamp",
		"commitHash",
		"jiraIssue",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServicePatchDto := _ServicePatchDto{}

	err = json.Unmarshal(data, &varServicePatchDto)

	if err != nil {
		return err
	}

	*o = ServicePatchDto(varServicePatchDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "owner")
		delete(additionalProperties, "description")
		delete(additionalProperties, "quicklinks")
		delete(additionalProperties, "repositories")
		delete(additionalProperties, "alertTarget")
		delete(additionalProperties, "operationType")
		delete(additionalProperties, "internetExposed")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "spec")
		delete(additionalProperties, "postPromotes")
		delete(additionalProperties, "timeStamp")
		delete(additionalProperties, "commitHash")
		delete(additionalProperties, "jiraIssue")
		delete(additionalProperties, "lifecycle")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicePatchDto struct {
	value *ServicePatchDto
	isSet bool
}

func (v NullableServicePatchDto) Get() *ServicePatchDto {
	return v.value
}

func (v *NullableServicePatchDto) Set(val *ServicePatchDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePatchDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePatchDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePatchDto(val *ServicePatchDto) *NullableServicePatchDto {
	return &NullableServicePatchDto{value: val, isSet: true}
}

func (v NullableServicePatchDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePatchDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



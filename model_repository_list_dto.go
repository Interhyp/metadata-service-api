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

// checks if the RepositoryListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryListDto{}

// RepositoryListDto struct for RepositoryListDto
type RepositoryListDto struct {
	Repositories map[string]RepositoryDto `json:"repositories"`
	// ISO-8601 UTC date time at which the list of repositories was obtained from service-metadata
	TimeStamp string `json:"timeStamp" yaml:"-"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryListDto RepositoryListDto

// NewRepositoryListDto instantiates a new RepositoryListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryListDto(repositories map[string]RepositoryDto, timeStamp string) *RepositoryListDto {
	this := RepositoryListDto{}
	this.Repositories = repositories
	this.TimeStamp = timeStamp
	return &this
}

// NewRepositoryListDtoWithDefaults instantiates a new RepositoryListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryListDtoWithDefaults() *RepositoryListDto {
	this := RepositoryListDto{}
	return &this
}

// GetRepositories returns the Repositories field value
func (o *RepositoryListDto) GetRepositories() map[string]RepositoryDto {
	if o == nil {
		var ret map[string]RepositoryDto
		return ret
	}

	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value
// and a boolean to check if the value has been set.
func (o *RepositoryListDto) GetRepositoriesOk() (map[string]RepositoryDto, bool) {
	if o == nil {
		return map[string]RepositoryDto{}, false
	}
	return o.Repositories, true
}

// SetRepositories sets field value
func (o *RepositoryListDto) SetRepositories(v map[string]RepositoryDto) {
	o.Repositories = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *RepositoryListDto) GetTimeStamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *RepositoryListDto) GetTimeStampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *RepositoryListDto) SetTimeStamp(v string) {
	o.TimeStamp = v
}

func (o RepositoryListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["repositories"] = o.Repositories
	toSerialize["timeStamp"] = o.TimeStamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"repositories",
		"timeStamp",
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

	varRepositoryListDto := _RepositoryListDto{}

	err = json.Unmarshal(data, &varRepositoryListDto)

	if err != nil {
		return err
	}

	*o = RepositoryListDto(varRepositoryListDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repositories")
		delete(additionalProperties, "timeStamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryListDto struct {
	value *RepositoryListDto
	isSet bool
}

func (v NullableRepositoryListDto) Get() *RepositoryListDto {
	return v.value
}

func (v *NullableRepositoryListDto) Set(val *RepositoryListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryListDto(val *RepositoryListDto) *NullableRepositoryListDto {
	return &NullableRepositoryListDto{value: val, isSet: true}
}

func (v NullableRepositoryListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



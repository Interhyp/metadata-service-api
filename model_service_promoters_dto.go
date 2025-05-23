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

// checks if the ServicePromotersDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePromotersDto{}

// ServicePromotersDto struct for ServicePromotersDto
type ServicePromotersDto struct {
	Promoters []string `json:"promoters"`
	AdditionalProperties map[string]interface{}
}

type _ServicePromotersDto ServicePromotersDto

// NewServicePromotersDto instantiates a new ServicePromotersDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePromotersDto(promoters []string) *ServicePromotersDto {
	this := ServicePromotersDto{}
	this.Promoters = promoters
	return &this
}

// NewServicePromotersDtoWithDefaults instantiates a new ServicePromotersDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePromotersDtoWithDefaults() *ServicePromotersDto {
	this := ServicePromotersDto{}
	return &this
}

// GetPromoters returns the Promoters field value
func (o *ServicePromotersDto) GetPromoters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Promoters
}

// GetPromotersOk returns a tuple with the Promoters field value
// and a boolean to check if the value has been set.
func (o *ServicePromotersDto) GetPromotersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Promoters, true
}

// SetPromoters sets field value
func (o *ServicePromotersDto) SetPromoters(v []string) {
	o.Promoters = v
}

func (o ServicePromotersDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePromotersDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["promoters"] = o.Promoters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicePromotersDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"promoters",
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

	varServicePromotersDto := _ServicePromotersDto{}

	err = json.Unmarshal(data, &varServicePromotersDto)

	if err != nil {
		return err
	}

	*o = ServicePromotersDto(varServicePromotersDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "promoters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicePromotersDto struct {
	value *ServicePromotersDto
	isSet bool
}

func (v NullableServicePromotersDto) Get() *ServicePromotersDto {
	return v.value
}

func (v *NullableServicePromotersDto) Set(val *ServicePromotersDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePromotersDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePromotersDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePromotersDto(val *ServicePromotersDto) *NullableServicePromotersDto {
	return &NullableServicePromotersDto{value: val, isSet: true}
}

func (v NullableServicePromotersDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePromotersDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



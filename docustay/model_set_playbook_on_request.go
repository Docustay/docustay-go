/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SetPlaybookOnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetPlaybookOnRequest{}

// SetPlaybookOnRequest struct for SetPlaybookOnRequest
type SetPlaybookOnRequest struct {
	Enabled bool `json:"enabled"`
}

type _SetPlaybookOnRequest SetPlaybookOnRequest

// NewSetPlaybookOnRequest instantiates a new SetPlaybookOnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPlaybookOnRequest(enabled bool) *SetPlaybookOnRequest {
	this := SetPlaybookOnRequest{}
	this.Enabled = enabled
	return &this
}

// NewSetPlaybookOnRequestWithDefaults instantiates a new SetPlaybookOnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPlaybookOnRequestWithDefaults() *SetPlaybookOnRequest {
	this := SetPlaybookOnRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *SetPlaybookOnRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SetPlaybookOnRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SetPlaybookOnRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SetPlaybookOnRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetPlaybookOnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *SetPlaybookOnRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varSetPlaybookOnRequest := _SetPlaybookOnRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetPlaybookOnRequest)

	if err != nil {
		return err
	}

	*o = SetPlaybookOnRequest(varSetPlaybookOnRequest)

	return err
}

type NullableSetPlaybookOnRequest struct {
	value *SetPlaybookOnRequest
	isSet bool
}

func (v NullableSetPlaybookOnRequest) Get() *SetPlaybookOnRequest {
	return v.value
}

func (v *NullableSetPlaybookOnRequest) Set(val *SetPlaybookOnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPlaybookOnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPlaybookOnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPlaybookOnRequest(val *SetPlaybookOnRequest) *NullableSetPlaybookOnRequest {
	return &NullableSetPlaybookOnRequest{value: val, isSet: true}
}

func (v NullableSetPlaybookOnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetPlaybookOnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



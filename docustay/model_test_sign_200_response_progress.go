/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
	"fmt"
)

// checks if the TestSign200ResponseProgress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSign200ResponseProgress{}

// TestSign200ResponseProgress struct for TestSign200ResponseProgress
type TestSign200ResponseProgress struct {
	Signed int32 `json:"signed"`
	Of int32 `json:"of"`
	AdditionalProperties map[string]interface{}
}

type _TestSign200ResponseProgress TestSign200ResponseProgress

// NewTestSign200ResponseProgress instantiates a new TestSign200ResponseProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSign200ResponseProgress(signed int32, of int32) *TestSign200ResponseProgress {
	this := TestSign200ResponseProgress{}
	this.Signed = signed
	this.Of = of
	return &this
}

// NewTestSign200ResponseProgressWithDefaults instantiates a new TestSign200ResponseProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSign200ResponseProgressWithDefaults() *TestSign200ResponseProgress {
	this := TestSign200ResponseProgress{}
	return &this
}

// GetSigned returns the Signed field value
func (o *TestSign200ResponseProgress) GetSigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Signed
}

// GetSignedOk returns a tuple with the Signed field value
// and a boolean to check if the value has been set.
func (o *TestSign200ResponseProgress) GetSignedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signed, true
}

// SetSigned sets field value
func (o *TestSign200ResponseProgress) SetSigned(v int32) {
	o.Signed = v
}

// GetOf returns the Of field value
func (o *TestSign200ResponseProgress) GetOf() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Of
}

// GetOfOk returns a tuple with the Of field value
// and a boolean to check if the value has been set.
func (o *TestSign200ResponseProgress) GetOfOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Of, true
}

// SetOf sets field value
func (o *TestSign200ResponseProgress) SetOf(v int32) {
	o.Of = v
}

func (o TestSign200ResponseProgress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSign200ResponseProgress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signed"] = o.Signed
	toSerialize["of"] = o.Of

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestSign200ResponseProgress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signed",
		"of",
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

	varTestSign200ResponseProgress := _TestSign200ResponseProgress{}

	err = json.Unmarshal(data, &varTestSign200ResponseProgress)

	if err != nil {
		return err
	}

	*o = TestSign200ResponseProgress(varTestSign200ResponseProgress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signed")
		delete(additionalProperties, "of")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestSign200ResponseProgress struct {
	value *TestSign200ResponseProgress
	isSet bool
}

func (v NullableTestSign200ResponseProgress) Get() *TestSign200ResponseProgress {
	return v.value
}

func (v *NullableTestSign200ResponseProgress) Set(val *TestSign200ResponseProgress) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSign200ResponseProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSign200ResponseProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSign200ResponseProgress(val *TestSign200ResponseProgress) *NullableTestSign200ResponseProgress {
	return &NullableTestSign200ResponseProgress{value: val, isSet: true}
}

func (v NullableTestSign200ResponseProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSign200ResponseProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



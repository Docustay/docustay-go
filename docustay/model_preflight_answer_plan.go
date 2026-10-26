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

// checks if the PreflightAnswerPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreflightAnswerPlan{}

// PreflightAnswerPlan struct for PreflightAnswerPlan
type PreflightAnswerPlan struct {
	Key string `json:"key"`
	Name string `json:"name"`
}

type _PreflightAnswerPlan PreflightAnswerPlan

// NewPreflightAnswerPlan instantiates a new PreflightAnswerPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreflightAnswerPlan(key string, name string) *PreflightAnswerPlan {
	this := PreflightAnswerPlan{}
	this.Key = key
	this.Name = name
	return &this
}

// NewPreflightAnswerPlanWithDefaults instantiates a new PreflightAnswerPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreflightAnswerPlanWithDefaults() *PreflightAnswerPlan {
	this := PreflightAnswerPlan{}
	return &this
}

// GetKey returns the Key field value
func (o *PreflightAnswerPlan) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerPlan) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PreflightAnswerPlan) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *PreflightAnswerPlan) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerPlan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PreflightAnswerPlan) SetName(v string) {
	o.Name = v
}

func (o PreflightAnswerPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreflightAnswerPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *PreflightAnswerPlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"name",
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

	varPreflightAnswerPlan := _PreflightAnswerPlan{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPreflightAnswerPlan)

	if err != nil {
		return err
	}

	*o = PreflightAnswerPlan(varPreflightAnswerPlan)

	return err
}

type NullablePreflightAnswerPlan struct {
	value *PreflightAnswerPlan
	isSet bool
}

func (v NullablePreflightAnswerPlan) Get() *PreflightAnswerPlan {
	return v.value
}

func (v *NullablePreflightAnswerPlan) Set(val *PreflightAnswerPlan) {
	v.value = val
	v.isSet = true
}

func (v NullablePreflightAnswerPlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePreflightAnswerPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreflightAnswerPlan(val *PreflightAnswerPlan) *NullablePreflightAnswerPlan {
	return &NullablePreflightAnswerPlan{value: val, isSet: true}
}

func (v NullablePreflightAnswerPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreflightAnswerPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



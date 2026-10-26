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

// checks if the TemplateMadeWarningsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateMadeWarningsInner{}

// TemplateMadeWarningsInner struct for TemplateMadeWarningsInner
type TemplateMadeWarningsInner struct {
	Tag string `json:"tag"`
	Warning string `json:"warning"`
	AdditionalProperties map[string]interface{}
}

type _TemplateMadeWarningsInner TemplateMadeWarningsInner

// NewTemplateMadeWarningsInner instantiates a new TemplateMadeWarningsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateMadeWarningsInner(tag string, warning string) *TemplateMadeWarningsInner {
	this := TemplateMadeWarningsInner{}
	this.Tag = tag
	this.Warning = warning
	return &this
}

// NewTemplateMadeWarningsInnerWithDefaults instantiates a new TemplateMadeWarningsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateMadeWarningsInnerWithDefaults() *TemplateMadeWarningsInner {
	this := TemplateMadeWarningsInner{}
	return &this
}

// GetTag returns the Tag field value
func (o *TemplateMadeWarningsInner) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *TemplateMadeWarningsInner) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *TemplateMadeWarningsInner) SetTag(v string) {
	o.Tag = v
}

// GetWarning returns the Warning field value
func (o *TemplateMadeWarningsInner) GetWarning() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Warning
}

// GetWarningOk returns a tuple with the Warning field value
// and a boolean to check if the value has been set.
func (o *TemplateMadeWarningsInner) GetWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warning, true
}

// SetWarning sets field value
func (o *TemplateMadeWarningsInner) SetWarning(v string) {
	o.Warning = v
}

func (o TemplateMadeWarningsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateMadeWarningsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tag"] = o.Tag
	toSerialize["warning"] = o.Warning

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateMadeWarningsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tag",
		"warning",
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

	varTemplateMadeWarningsInner := _TemplateMadeWarningsInner{}

	err = json.Unmarshal(data, &varTemplateMadeWarningsInner)

	if err != nil {
		return err
	}

	*o = TemplateMadeWarningsInner(varTemplateMadeWarningsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tag")
		delete(additionalProperties, "warning")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateMadeWarningsInner struct {
	value *TemplateMadeWarningsInner
	isSet bool
}

func (v NullableTemplateMadeWarningsInner) Get() *TemplateMadeWarningsInner {
	return v.value
}

func (v *NullableTemplateMadeWarningsInner) Set(val *TemplateMadeWarningsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateMadeWarningsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateMadeWarningsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateMadeWarningsInner(val *TemplateMadeWarningsInner) *NullableTemplateMadeWarningsInner {
	return &NullableTemplateMadeWarningsInner{value: val, isSet: true}
}

func (v NullableTemplateMadeWarningsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateMadeWarningsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



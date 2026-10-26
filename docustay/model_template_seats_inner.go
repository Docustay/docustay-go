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

// checks if the TemplateSeatsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateSeatsInner{}

// TemplateSeatsInner struct for TemplateSeatsInner
type TemplateSeatsInner struct {
	Key string `json:"key"`
	Label string `json:"label"`
	AdditionalProperties map[string]interface{}
}

type _TemplateSeatsInner TemplateSeatsInner

// NewTemplateSeatsInner instantiates a new TemplateSeatsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSeatsInner(key string, label string) *TemplateSeatsInner {
	this := TemplateSeatsInner{}
	this.Key = key
	this.Label = label
	return &this
}

// NewTemplateSeatsInnerWithDefaults instantiates a new TemplateSeatsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSeatsInnerWithDefaults() *TemplateSeatsInner {
	this := TemplateSeatsInner{}
	return &this
}

// GetKey returns the Key field value
func (o *TemplateSeatsInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TemplateSeatsInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TemplateSeatsInner) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value
func (o *TemplateSeatsInner) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TemplateSeatsInner) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *TemplateSeatsInner) SetLabel(v string) {
	o.Label = v
}

func (o TemplateSeatsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateSeatsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["label"] = o.Label

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateSeatsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"label",
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

	varTemplateSeatsInner := _TemplateSeatsInner{}

	err = json.Unmarshal(data, &varTemplateSeatsInner)

	if err != nil {
		return err
	}

	*o = TemplateSeatsInner(varTemplateSeatsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateSeatsInner struct {
	value *TemplateSeatsInner
	isSet bool
}

func (v NullableTemplateSeatsInner) Get() *TemplateSeatsInner {
	return v.value
}

func (v *NullableTemplateSeatsInner) Set(val *TemplateSeatsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSeatsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSeatsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSeatsInner(val *TemplateSeatsInner) *NullableTemplateSeatsInner {
	return &NullableTemplateSeatsInner{value: val, isSet: true}
}

func (v NullableTemplateSeatsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSeatsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



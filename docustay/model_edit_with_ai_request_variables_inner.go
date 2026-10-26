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

// checks if the EditWithAiRequestVariablesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditWithAiRequestVariablesInner{}

// EditWithAiRequestVariablesInner struct for EditWithAiRequestVariablesInner
type EditWithAiRequestVariablesInner struct {
	Key string `json:"key"`
	Label *string `json:"label,omitempty"`
	Kind *string `json:"kind,omitempty"`
}

type _EditWithAiRequestVariablesInner EditWithAiRequestVariablesInner

// NewEditWithAiRequestVariablesInner instantiates a new EditWithAiRequestVariablesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditWithAiRequestVariablesInner(key string) *EditWithAiRequestVariablesInner {
	this := EditWithAiRequestVariablesInner{}
	this.Key = key
	return &this
}

// NewEditWithAiRequestVariablesInnerWithDefaults instantiates a new EditWithAiRequestVariablesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditWithAiRequestVariablesInnerWithDefaults() *EditWithAiRequestVariablesInner {
	this := EditWithAiRequestVariablesInner{}
	return &this
}

// GetKey returns the Key field value
func (o *EditWithAiRequestVariablesInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EditWithAiRequestVariablesInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EditWithAiRequestVariablesInner) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EditWithAiRequestVariablesInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditWithAiRequestVariablesInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EditWithAiRequestVariablesInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *EditWithAiRequestVariablesInner) SetLabel(v string) {
	o.Label = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *EditWithAiRequestVariablesInner) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditWithAiRequestVariablesInner) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *EditWithAiRequestVariablesInner) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *EditWithAiRequestVariablesInner) SetKind(v string) {
	o.Kind = &v
}

func (o EditWithAiRequestVariablesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditWithAiRequestVariablesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return toSerialize, nil
}

func (o *EditWithAiRequestVariablesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varEditWithAiRequestVariablesInner := _EditWithAiRequestVariablesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditWithAiRequestVariablesInner)

	if err != nil {
		return err
	}

	*o = EditWithAiRequestVariablesInner(varEditWithAiRequestVariablesInner)

	return err
}

type NullableEditWithAiRequestVariablesInner struct {
	value *EditWithAiRequestVariablesInner
	isSet bool
}

func (v NullableEditWithAiRequestVariablesInner) Get() *EditWithAiRequestVariablesInner {
	return v.value
}

func (v *NullableEditWithAiRequestVariablesInner) Set(val *EditWithAiRequestVariablesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEditWithAiRequestVariablesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEditWithAiRequestVariablesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditWithAiRequestVariablesInner(val *EditWithAiRequestVariablesInner) *NullableEditWithAiRequestVariablesInner {
	return &NullableEditWithAiRequestVariablesInner{value: val, isSet: true}
}

func (v NullableEditWithAiRequestVariablesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditWithAiRequestVariablesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



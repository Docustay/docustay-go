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

// checks if the DetectFieldsResultPagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectFieldsResultPagesInner{}

// DetectFieldsResultPagesInner struct for DetectFieldsResultPagesInner
type DetectFieldsResultPagesInner struct {
	Page int32 `json:"page"`
	Kind string `json:"kind"`
	Fields int32 `json:"fields"`
	AdditionalProperties map[string]interface{}
}

type _DetectFieldsResultPagesInner DetectFieldsResultPagesInner

// NewDetectFieldsResultPagesInner instantiates a new DetectFieldsResultPagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectFieldsResultPagesInner(page int32, kind string, fields int32) *DetectFieldsResultPagesInner {
	this := DetectFieldsResultPagesInner{}
	this.Page = page
	this.Kind = kind
	this.Fields = fields
	return &this
}

// NewDetectFieldsResultPagesInnerWithDefaults instantiates a new DetectFieldsResultPagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectFieldsResultPagesInnerWithDefaults() *DetectFieldsResultPagesInner {
	this := DetectFieldsResultPagesInner{}
	return &this
}

// GetPage returns the Page field value
func (o *DetectFieldsResultPagesInner) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultPagesInner) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *DetectFieldsResultPagesInner) SetPage(v int32) {
	o.Page = v
}

// GetKind returns the Kind field value
func (o *DetectFieldsResultPagesInner) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultPagesInner) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *DetectFieldsResultPagesInner) SetKind(v string) {
	o.Kind = v
}

// GetFields returns the Fields field value
func (o *DetectFieldsResultPagesInner) GetFields() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultPagesInner) GetFieldsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *DetectFieldsResultPagesInner) SetFields(v int32) {
	o.Fields = v
}

func (o DetectFieldsResultPagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectFieldsResultPagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page
	toSerialize["kind"] = o.Kind
	toSerialize["fields"] = o.Fields

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DetectFieldsResultPagesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"page",
		"kind",
		"fields",
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

	varDetectFieldsResultPagesInner := _DetectFieldsResultPagesInner{}

	err = json.Unmarshal(data, &varDetectFieldsResultPagesInner)

	if err != nil {
		return err
	}

	*o = DetectFieldsResultPagesInner(varDetectFieldsResultPagesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetectFieldsResultPagesInner struct {
	value *DetectFieldsResultPagesInner
	isSet bool
}

func (v NullableDetectFieldsResultPagesInner) Get() *DetectFieldsResultPagesInner {
	return v.value
}

func (v *NullableDetectFieldsResultPagesInner) Set(val *DetectFieldsResultPagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectFieldsResultPagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectFieldsResultPagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectFieldsResultPagesInner(val *DetectFieldsResultPagesInner) *NullableDetectFieldsResultPagesInner {
	return &NullableDetectFieldsResultPagesInner{value: val, isSet: true}
}

func (v NullableDetectFieldsResultPagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectFieldsResultPagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



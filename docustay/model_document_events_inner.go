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

// checks if the DocumentEventsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentEventsInner{}

// DocumentEventsInner struct for DocumentEventsInner
type DocumentEventsInner struct {
	At NullableString `json:"at"`
	Kind string `json:"kind"`
	Text string `json:"text"`
	AdditionalProperties map[string]interface{}
}

type _DocumentEventsInner DocumentEventsInner

// NewDocumentEventsInner instantiates a new DocumentEventsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentEventsInner(at NullableString, kind string, text string) *DocumentEventsInner {
	this := DocumentEventsInner{}
	this.At = at
	this.Kind = kind
	this.Text = text
	return &this
}

// NewDocumentEventsInnerWithDefaults instantiates a new DocumentEventsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentEventsInnerWithDefaults() *DocumentEventsInner {
	this := DocumentEventsInner{}
	return &this
}

// GetAt returns the At field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentEventsInner) GetAt() string {
	if o == nil || o.At.Get() == nil {
		var ret string
		return ret
	}

	return *o.At.Get()
}

// GetAtOk returns a tuple with the At field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentEventsInner) GetAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.At.Get(), o.At.IsSet()
}

// SetAt sets field value
func (o *DocumentEventsInner) SetAt(v string) {
	o.At.Set(&v)
}

// GetKind returns the Kind field value
func (o *DocumentEventsInner) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DocumentEventsInner) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *DocumentEventsInner) SetKind(v string) {
	o.Kind = v
}

// GetText returns the Text field value
func (o *DocumentEventsInner) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *DocumentEventsInner) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *DocumentEventsInner) SetText(v string) {
	o.Text = v
}

func (o DocumentEventsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentEventsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["at"] = o.At.Get()
	toSerialize["kind"] = o.Kind
	toSerialize["text"] = o.Text

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DocumentEventsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"at",
		"kind",
		"text",
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

	varDocumentEventsInner := _DocumentEventsInner{}

	err = json.Unmarshal(data, &varDocumentEventsInner)

	if err != nil {
		return err
	}

	*o = DocumentEventsInner(varDocumentEventsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "at")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentEventsInner struct {
	value *DocumentEventsInner
	isSet bool
}

func (v NullableDocumentEventsInner) Get() *DocumentEventsInner {
	return v.value
}

func (v *NullableDocumentEventsInner) Set(val *DocumentEventsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentEventsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentEventsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentEventsInner(val *DocumentEventsInner) *NullableDocumentEventsInner {
	return &NullableDocumentEventsInner{value: val, isSet: true}
}

func (v NullableDocumentEventsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentEventsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



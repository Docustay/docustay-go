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

// checks if the DetectFieldsResultDocumentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectFieldsResultDocumentsInner{}

// DetectFieldsResultDocumentsInner struct for DetectFieldsResultDocumentsInner
type DetectFieldsResultDocumentsInner struct {
	From int32 `json:"from"`
	To int32 `json:"to"`
	Kind string `json:"kind"`
	Title NullableString `json:"title"`
	AdditionalProperties map[string]interface{}
}

type _DetectFieldsResultDocumentsInner DetectFieldsResultDocumentsInner

// NewDetectFieldsResultDocumentsInner instantiates a new DetectFieldsResultDocumentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectFieldsResultDocumentsInner(from int32, to int32, kind string, title NullableString) *DetectFieldsResultDocumentsInner {
	this := DetectFieldsResultDocumentsInner{}
	this.From = from
	this.To = to
	this.Kind = kind
	this.Title = title
	return &this
}

// NewDetectFieldsResultDocumentsInnerWithDefaults instantiates a new DetectFieldsResultDocumentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectFieldsResultDocumentsInnerWithDefaults() *DetectFieldsResultDocumentsInner {
	this := DetectFieldsResultDocumentsInner{}
	return &this
}

// GetFrom returns the From field value
func (o *DetectFieldsResultDocumentsInner) GetFrom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultDocumentsInner) GetFromOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *DetectFieldsResultDocumentsInner) SetFrom(v int32) {
	o.From = v
}

// GetTo returns the To field value
func (o *DetectFieldsResultDocumentsInner) GetTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultDocumentsInner) GetToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *DetectFieldsResultDocumentsInner) SetTo(v int32) {
	o.To = v
}

// GetKind returns the Kind field value
func (o *DetectFieldsResultDocumentsInner) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultDocumentsInner) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *DetectFieldsResultDocumentsInner) SetKind(v string) {
	o.Kind = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DetectFieldsResultDocumentsInner) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetectFieldsResultDocumentsInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *DetectFieldsResultDocumentsInner) SetTitle(v string) {
	o.Title.Set(&v)
}

func (o DetectFieldsResultDocumentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectFieldsResultDocumentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	toSerialize["kind"] = o.Kind
	toSerialize["title"] = o.Title.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DetectFieldsResultDocumentsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from",
		"to",
		"kind",
		"title",
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

	varDetectFieldsResultDocumentsInner := _DetectFieldsResultDocumentsInner{}

	err = json.Unmarshal(data, &varDetectFieldsResultDocumentsInner)

	if err != nil {
		return err
	}

	*o = DetectFieldsResultDocumentsInner(varDetectFieldsResultDocumentsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetectFieldsResultDocumentsInner struct {
	value *DetectFieldsResultDocumentsInner
	isSet bool
}

func (v NullableDetectFieldsResultDocumentsInner) Get() *DetectFieldsResultDocumentsInner {
	return v.value
}

func (v *NullableDetectFieldsResultDocumentsInner) Set(val *DetectFieldsResultDocumentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectFieldsResultDocumentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectFieldsResultDocumentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectFieldsResultDocumentsInner(val *DetectFieldsResultDocumentsInner) *NullableDetectFieldsResultDocumentsInner {
	return &NullableDetectFieldsResultDocumentsInner{value: val, isSet: true}
}

func (v NullableDetectFieldsResultDocumentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectFieldsResultDocumentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



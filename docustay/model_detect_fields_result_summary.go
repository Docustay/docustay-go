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

// checks if the DetectFieldsResultSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectFieldsResultSummary{}

// DetectFieldsResultSummary struct for DetectFieldsResultSummary
type DetectFieldsResultSummary struct {
	Fields int32 `json:"fields"`
	Signers int32 `json:"signers"`
	Uncertain int32 `json:"uncertain"`
	Truncated bool `json:"truncated"`
	Cached bool `json:"cached"`
	AdditionalProperties map[string]interface{}
}

type _DetectFieldsResultSummary DetectFieldsResultSummary

// NewDetectFieldsResultSummary instantiates a new DetectFieldsResultSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectFieldsResultSummary(fields int32, signers int32, uncertain int32, truncated bool, cached bool) *DetectFieldsResultSummary {
	this := DetectFieldsResultSummary{}
	this.Fields = fields
	this.Signers = signers
	this.Uncertain = uncertain
	this.Truncated = truncated
	this.Cached = cached
	return &this
}

// NewDetectFieldsResultSummaryWithDefaults instantiates a new DetectFieldsResultSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectFieldsResultSummaryWithDefaults() *DetectFieldsResultSummary {
	this := DetectFieldsResultSummary{}
	return &this
}

// GetFields returns the Fields field value
func (o *DetectFieldsResultSummary) GetFields() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultSummary) GetFieldsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *DetectFieldsResultSummary) SetFields(v int32) {
	o.Fields = v
}

// GetSigners returns the Signers field value
func (o *DetectFieldsResultSummary) GetSigners() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Signers
}

// GetSignersOk returns a tuple with the Signers field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultSummary) GetSignersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signers, true
}

// SetSigners sets field value
func (o *DetectFieldsResultSummary) SetSigners(v int32) {
	o.Signers = v
}

// GetUncertain returns the Uncertain field value
func (o *DetectFieldsResultSummary) GetUncertain() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Uncertain
}

// GetUncertainOk returns a tuple with the Uncertain field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultSummary) GetUncertainOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uncertain, true
}

// SetUncertain sets field value
func (o *DetectFieldsResultSummary) SetUncertain(v int32) {
	o.Uncertain = v
}

// GetTruncated returns the Truncated field value
func (o *DetectFieldsResultSummary) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultSummary) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value
func (o *DetectFieldsResultSummary) SetTruncated(v bool) {
	o.Truncated = v
}

// GetCached returns the Cached field value
func (o *DetectFieldsResultSummary) GetCached() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cached
}

// GetCachedOk returns a tuple with the Cached field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsResultSummary) GetCachedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cached, true
}

// SetCached sets field value
func (o *DetectFieldsResultSummary) SetCached(v bool) {
	o.Cached = v
}

func (o DetectFieldsResultSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectFieldsResultSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	toSerialize["signers"] = o.Signers
	toSerialize["uncertain"] = o.Uncertain
	toSerialize["truncated"] = o.Truncated
	toSerialize["cached"] = o.Cached

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DetectFieldsResultSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields",
		"signers",
		"uncertain",
		"truncated",
		"cached",
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

	varDetectFieldsResultSummary := _DetectFieldsResultSummary{}

	err = json.Unmarshal(data, &varDetectFieldsResultSummary)

	if err != nil {
		return err
	}

	*o = DetectFieldsResultSummary(varDetectFieldsResultSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fields")
		delete(additionalProperties, "signers")
		delete(additionalProperties, "uncertain")
		delete(additionalProperties, "truncated")
		delete(additionalProperties, "cached")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetectFieldsResultSummary struct {
	value *DetectFieldsResultSummary
	isSet bool
}

func (v NullableDetectFieldsResultSummary) Get() *DetectFieldsResultSummary {
	return v.value
}

func (v *NullableDetectFieldsResultSummary) Set(val *DetectFieldsResultSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectFieldsResultSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectFieldsResultSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectFieldsResultSummary(val *DetectFieldsResultSummary) *NullableDetectFieldsResultSummary {
	return &NullableDetectFieldsResultSummary{value: val, isSet: true}
}

func (v NullableDetectFieldsResultSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectFieldsResultSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



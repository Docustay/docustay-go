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

// checks if the TemplateMadeDetected type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateMadeDetected{}

// TemplateMadeDetected Present when the page itself was looked at (a file without text tags): how many fields were found, how many seats, how many it was unsure of, and how many were placed (`fields: auto`).
type TemplateMadeDetected struct {
	Fields int32 `json:"fields"`
	Signers int32 `json:"signers"`
	LowConfidence int32 `json:"lowConfidence"`
	Placed int32 `json:"placed"`
	AdditionalProperties map[string]interface{}
}

type _TemplateMadeDetected TemplateMadeDetected

// NewTemplateMadeDetected instantiates a new TemplateMadeDetected object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateMadeDetected(fields int32, signers int32, lowConfidence int32, placed int32) *TemplateMadeDetected {
	this := TemplateMadeDetected{}
	this.Fields = fields
	this.Signers = signers
	this.LowConfidence = lowConfidence
	this.Placed = placed
	return &this
}

// NewTemplateMadeDetectedWithDefaults instantiates a new TemplateMadeDetected object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateMadeDetectedWithDefaults() *TemplateMadeDetected {
	this := TemplateMadeDetected{}
	return &this
}

// GetFields returns the Fields field value
func (o *TemplateMadeDetected) GetFields() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *TemplateMadeDetected) GetFieldsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *TemplateMadeDetected) SetFields(v int32) {
	o.Fields = v
}

// GetSigners returns the Signers field value
func (o *TemplateMadeDetected) GetSigners() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Signers
}

// GetSignersOk returns a tuple with the Signers field value
// and a boolean to check if the value has been set.
func (o *TemplateMadeDetected) GetSignersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signers, true
}

// SetSigners sets field value
func (o *TemplateMadeDetected) SetSigners(v int32) {
	o.Signers = v
}

// GetLowConfidence returns the LowConfidence field value
func (o *TemplateMadeDetected) GetLowConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LowConfidence
}

// GetLowConfidenceOk returns a tuple with the LowConfidence field value
// and a boolean to check if the value has been set.
func (o *TemplateMadeDetected) GetLowConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowConfidence, true
}

// SetLowConfidence sets field value
func (o *TemplateMadeDetected) SetLowConfidence(v int32) {
	o.LowConfidence = v
}

// GetPlaced returns the Placed field value
func (o *TemplateMadeDetected) GetPlaced() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Placed
}

// GetPlacedOk returns a tuple with the Placed field value
// and a boolean to check if the value has been set.
func (o *TemplateMadeDetected) GetPlacedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Placed, true
}

// SetPlaced sets field value
func (o *TemplateMadeDetected) SetPlaced(v int32) {
	o.Placed = v
}

func (o TemplateMadeDetected) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateMadeDetected) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	toSerialize["signers"] = o.Signers
	toSerialize["lowConfidence"] = o.LowConfidence
	toSerialize["placed"] = o.Placed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateMadeDetected) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields",
		"signers",
		"lowConfidence",
		"placed",
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

	varTemplateMadeDetected := _TemplateMadeDetected{}

	err = json.Unmarshal(data, &varTemplateMadeDetected)

	if err != nil {
		return err
	}

	*o = TemplateMadeDetected(varTemplateMadeDetected)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fields")
		delete(additionalProperties, "signers")
		delete(additionalProperties, "lowConfidence")
		delete(additionalProperties, "placed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateMadeDetected struct {
	value *TemplateMadeDetected
	isSet bool
}

func (v NullableTemplateMadeDetected) Get() *TemplateMadeDetected {
	return v.value
}

func (v *NullableTemplateMadeDetected) Set(val *TemplateMadeDetected) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateMadeDetected) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateMadeDetected) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateMadeDetected(val *TemplateMadeDetected) *NullableTemplateMadeDetected {
	return &NullableTemplateMadeDetected{value: val, isSet: true}
}

func (v NullableTemplateMadeDetected) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateMadeDetected) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



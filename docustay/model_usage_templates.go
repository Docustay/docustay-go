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

// checks if the UsageTemplates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageTemplates{}

// UsageTemplates New templates made this period, documents and forms together. Free includes 5 and Pay as you go 25 (from 1 November 2026); `limit` is null on Pro and Lifetime.
type UsageTemplates struct {
	Used int32 `json:"used"`
	Limit NullableInt32 `json:"limit"`
	ResetsAt NullableString `json:"resetsAt"`
}

type _UsageTemplates UsageTemplates

// NewUsageTemplates instantiates a new UsageTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageTemplates(used int32, limit NullableInt32, resetsAt NullableString) *UsageTemplates {
	this := UsageTemplates{}
	this.Used = used
	this.Limit = limit
	this.ResetsAt = resetsAt
	return &this
}

// NewUsageTemplatesWithDefaults instantiates a new UsageTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageTemplatesWithDefaults() *UsageTemplates {
	this := UsageTemplates{}
	return &this
}

// GetUsed returns the Used field value
func (o *UsageTemplates) GetUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *UsageTemplates) GetUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *UsageTemplates) SetUsed(v int32) {
	o.Used = v
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageTemplates) GetLimit() int32 {
	if o == nil || o.Limit.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageTemplates) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// SetLimit sets field value
func (o *UsageTemplates) SetLimit(v int32) {
	o.Limit.Set(&v)
}

// GetResetsAt returns the ResetsAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UsageTemplates) GetResetsAt() string {
	if o == nil || o.ResetsAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResetsAt.Get()
}

// GetResetsAtOk returns a tuple with the ResetsAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageTemplates) GetResetsAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResetsAt.Get(), o.ResetsAt.IsSet()
}

// SetResetsAt sets field value
func (o *UsageTemplates) SetResetsAt(v string) {
	o.ResetsAt.Set(&v)
}

func (o UsageTemplates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageTemplates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["used"] = o.Used
	toSerialize["limit"] = o.Limit.Get()
	toSerialize["resetsAt"] = o.ResetsAt.Get()
	return toSerialize, nil
}

func (o *UsageTemplates) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"used",
		"limit",
		"resetsAt",
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

	varUsageTemplates := _UsageTemplates{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageTemplates)

	if err != nil {
		return err
	}

	*o = UsageTemplates(varUsageTemplates)

	return err
}

type NullableUsageTemplates struct {
	value *UsageTemplates
	isSet bool
}

func (v NullableUsageTemplates) Get() *UsageTemplates {
	return v.value
}

func (v *NullableUsageTemplates) Set(val *UsageTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageTemplates(val *UsageTemplates) *NullableUsageTemplates {
	return &NullableUsageTemplates{value: val, isSet: true}
}

func (v NullableUsageTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the UsageSigners type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageSigners{}

// UsageSigners Different signers sent to this period, counted once by email address (your own team does not count). Free includes 15 from 1 November 2026; `limit` and `used` are null on the paid plans, where signers are free and unlimited.
type UsageSigners struct {
	Used NullableInt32 `json:"used"`
	Limit NullableInt32 `json:"limit"`
	ResetsAt NullableString `json:"resetsAt"`
}

type _UsageSigners UsageSigners

// NewUsageSigners instantiates a new UsageSigners object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSigners(used NullableInt32, limit NullableInt32, resetsAt NullableString) *UsageSigners {
	this := UsageSigners{}
	this.Used = used
	this.Limit = limit
	this.ResetsAt = resetsAt
	return &this
}

// NewUsageSignersWithDefaults instantiates a new UsageSigners object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSignersWithDefaults() *UsageSigners {
	this := UsageSigners{}
	return &this
}

// GetUsed returns the Used field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageSigners) GetUsed() int32 {
	if o == nil || o.Used.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Used.Get()
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageSigners) GetUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Used.Get(), o.Used.IsSet()
}

// SetUsed sets field value
func (o *UsageSigners) SetUsed(v int32) {
	o.Used.Set(&v)
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageSigners) GetLimit() int32 {
	if o == nil || o.Limit.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageSigners) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// SetLimit sets field value
func (o *UsageSigners) SetLimit(v int32) {
	o.Limit.Set(&v)
}

// GetResetsAt returns the ResetsAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UsageSigners) GetResetsAt() string {
	if o == nil || o.ResetsAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResetsAt.Get()
}

// GetResetsAtOk returns a tuple with the ResetsAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageSigners) GetResetsAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResetsAt.Get(), o.ResetsAt.IsSet()
}

// SetResetsAt sets field value
func (o *UsageSigners) SetResetsAt(v string) {
	o.ResetsAt.Set(&v)
}

func (o UsageSigners) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageSigners) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["used"] = o.Used.Get()
	toSerialize["limit"] = o.Limit.Get()
	toSerialize["resetsAt"] = o.ResetsAt.Get()
	return toSerialize, nil
}

func (o *UsageSigners) UnmarshalJSON(data []byte) (err error) {
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

	varUsageSigners := _UsageSigners{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageSigners)

	if err != nil {
		return err
	}

	*o = UsageSigners(varUsageSigners)

	return err
}

type NullableUsageSigners struct {
	value *UsageSigners
	isSet bool
}

func (v NullableUsageSigners) Get() *UsageSigners {
	return v.value
}

func (v *NullableUsageSigners) Set(val *UsageSigners) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageSigners) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageSigners) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageSigners(val *UsageSigners) *NullableUsageSigners {
	return &NullableUsageSigners{value: val, isSet: true}
}

func (v NullableUsageSigners) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageSigners) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



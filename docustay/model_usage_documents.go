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

// checks if the UsageDocuments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageDocuments{}

// UsageDocuments struct for UsageDocuments
type UsageDocuments struct {
	Used int32 `json:"used"`
	Limit NullableInt32 `json:"limit"`
}

type _UsageDocuments UsageDocuments

// NewUsageDocuments instantiates a new UsageDocuments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageDocuments(used int32, limit NullableInt32) *UsageDocuments {
	this := UsageDocuments{}
	this.Used = used
	this.Limit = limit
	return &this
}

// NewUsageDocumentsWithDefaults instantiates a new UsageDocuments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageDocumentsWithDefaults() *UsageDocuments {
	this := UsageDocuments{}
	return &this
}

// GetUsed returns the Used field value
func (o *UsageDocuments) GetUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *UsageDocuments) GetUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *UsageDocuments) SetUsed(v int32) {
	o.Used = v
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageDocuments) GetLimit() int32 {
	if o == nil || o.Limit.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageDocuments) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// SetLimit sets field value
func (o *UsageDocuments) SetLimit(v int32) {
	o.Limit.Set(&v)
}

func (o UsageDocuments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageDocuments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["used"] = o.Used
	toSerialize["limit"] = o.Limit.Get()
	return toSerialize, nil
}

func (o *UsageDocuments) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"used",
		"limit",
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

	varUsageDocuments := _UsageDocuments{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageDocuments)

	if err != nil {
		return err
	}

	*o = UsageDocuments(varUsageDocuments)

	return err
}

type NullableUsageDocuments struct {
	value *UsageDocuments
	isSet bool
}

func (v NullableUsageDocuments) Get() *UsageDocuments {
	return v.value
}

func (v *NullableUsageDocuments) Set(val *UsageDocuments) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageDocuments) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageDocuments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageDocuments(val *UsageDocuments) *NullableUsageDocuments {
	return &NullableUsageDocuments{value: val, isSet: true}
}

func (v NullableUsageDocuments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageDocuments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



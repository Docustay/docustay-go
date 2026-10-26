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

// checks if the UsagePeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsagePeriod{}

// UsagePeriod struct for UsagePeriod
type UsagePeriod struct {
	Start string `json:"start"`
	End string `json:"end"`
}

type _UsagePeriod UsagePeriod

// NewUsagePeriod instantiates a new UsagePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsagePeriod(start string, end string) *UsagePeriod {
	this := UsagePeriod{}
	this.Start = start
	this.End = end
	return &this
}

// NewUsagePeriodWithDefaults instantiates a new UsagePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsagePeriodWithDefaults() *UsagePeriod {
	this := UsagePeriod{}
	return &this
}

// GetStart returns the Start field value
func (o *UsagePeriod) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *UsagePeriod) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *UsagePeriod) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *UsagePeriod) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *UsagePeriod) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *UsagePeriod) SetEnd(v string) {
	o.End = v
}

func (o UsagePeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsagePeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	return toSerialize, nil
}

func (o *UsagePeriod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
		"end",
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

	varUsagePeriod := _UsagePeriod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsagePeriod)

	if err != nil {
		return err
	}

	*o = UsagePeriod(varUsagePeriod)

	return err
}

type NullableUsagePeriod struct {
	value *UsagePeriod
	isSet bool
}

func (v NullableUsagePeriod) Get() *UsagePeriod {
	return v.value
}

func (v *NullableUsagePeriod) Set(val *UsagePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableUsagePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableUsagePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsagePeriod(val *UsagePeriod) *NullableUsagePeriod {
	return &NullableUsagePeriod{value: val, isSet: true}
}

func (v NullableUsagePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsagePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



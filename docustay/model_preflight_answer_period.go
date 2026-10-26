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

// checks if the PreflightAnswerPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreflightAnswerPeriod{}

// PreflightAnswerPeriod struct for PreflightAnswerPeriod
type PreflightAnswerPeriod struct {
	Start string `json:"start"`
	End string `json:"end"`
	Resets string `json:"resets"`
	Kind string `json:"kind"`
}

type _PreflightAnswerPeriod PreflightAnswerPeriod

// NewPreflightAnswerPeriod instantiates a new PreflightAnswerPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreflightAnswerPeriod(start string, end string, resets string, kind string) *PreflightAnswerPeriod {
	this := PreflightAnswerPeriod{}
	this.Start = start
	this.End = end
	this.Resets = resets
	this.Kind = kind
	return &this
}

// NewPreflightAnswerPeriodWithDefaults instantiates a new PreflightAnswerPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreflightAnswerPeriodWithDefaults() *PreflightAnswerPeriod {
	this := PreflightAnswerPeriod{}
	return &this
}

// GetStart returns the Start field value
func (o *PreflightAnswerPeriod) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerPeriod) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *PreflightAnswerPeriod) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *PreflightAnswerPeriod) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerPeriod) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *PreflightAnswerPeriod) SetEnd(v string) {
	o.End = v
}

// GetResets returns the Resets field value
func (o *PreflightAnswerPeriod) GetResets() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resets
}

// GetResetsOk returns a tuple with the Resets field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerPeriod) GetResetsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resets, true
}

// SetResets sets field value
func (o *PreflightAnswerPeriod) SetResets(v string) {
	o.Resets = v
}

// GetKind returns the Kind field value
func (o *PreflightAnswerPeriod) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerPeriod) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *PreflightAnswerPeriod) SetKind(v string) {
	o.Kind = v
}

func (o PreflightAnswerPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreflightAnswerPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	toSerialize["resets"] = o.Resets
	toSerialize["kind"] = o.Kind
	return toSerialize, nil
}

func (o *PreflightAnswerPeriod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
		"end",
		"resets",
		"kind",
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

	varPreflightAnswerPeriod := _PreflightAnswerPeriod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPreflightAnswerPeriod)

	if err != nil {
		return err
	}

	*o = PreflightAnswerPeriod(varPreflightAnswerPeriod)

	return err
}

type NullablePreflightAnswerPeriod struct {
	value *PreflightAnswerPeriod
	isSet bool
}

func (v NullablePreflightAnswerPeriod) Get() *PreflightAnswerPeriod {
	return v.value
}

func (v *NullablePreflightAnswerPeriod) Set(val *PreflightAnswerPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullablePreflightAnswerPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullablePreflightAnswerPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreflightAnswerPeriod(val *PreflightAnswerPeriod) *NullablePreflightAnswerPeriod {
	return &NullablePreflightAnswerPeriod{value: val, isSet: true}
}

func (v NullablePreflightAnswerPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreflightAnswerPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



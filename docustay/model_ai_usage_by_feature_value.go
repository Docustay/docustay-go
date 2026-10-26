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

// checks if the AiUsageByFeatureValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiUsageByFeatureValue{}

// AiUsageByFeatureValue struct for AiUsageByFeatureValue
type AiUsageByFeatureValue struct {
	Calls float32 `json:"calls"`
	Credits float32 `json:"credits"`
	AdditionalProperties map[string]interface{}
}

type _AiUsageByFeatureValue AiUsageByFeatureValue

// NewAiUsageByFeatureValue instantiates a new AiUsageByFeatureValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiUsageByFeatureValue(calls float32, credits float32) *AiUsageByFeatureValue {
	this := AiUsageByFeatureValue{}
	this.Calls = calls
	this.Credits = credits
	return &this
}

// NewAiUsageByFeatureValueWithDefaults instantiates a new AiUsageByFeatureValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiUsageByFeatureValueWithDefaults() *AiUsageByFeatureValue {
	this := AiUsageByFeatureValue{}
	return &this
}

// GetCalls returns the Calls field value
func (o *AiUsageByFeatureValue) GetCalls() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value
// and a boolean to check if the value has been set.
func (o *AiUsageByFeatureValue) GetCallsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Calls, true
}

// SetCalls sets field value
func (o *AiUsageByFeatureValue) SetCalls(v float32) {
	o.Calls = v
}

// GetCredits returns the Credits field value
func (o *AiUsageByFeatureValue) GetCredits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value
// and a boolean to check if the value has been set.
func (o *AiUsageByFeatureValue) GetCreditsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credits, true
}

// SetCredits sets field value
func (o *AiUsageByFeatureValue) SetCredits(v float32) {
	o.Credits = v
}

func (o AiUsageByFeatureValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiUsageByFeatureValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["calls"] = o.Calls
	toSerialize["credits"] = o.Credits

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiUsageByFeatureValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"calls",
		"credits",
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

	varAiUsageByFeatureValue := _AiUsageByFeatureValue{}

	err = json.Unmarshal(data, &varAiUsageByFeatureValue)

	if err != nil {
		return err
	}

	*o = AiUsageByFeatureValue(varAiUsageByFeatureValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "calls")
		delete(additionalProperties, "credits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiUsageByFeatureValue struct {
	value *AiUsageByFeatureValue
	isSet bool
}

func (v NullableAiUsageByFeatureValue) Get() *AiUsageByFeatureValue {
	return v.value
}

func (v *NullableAiUsageByFeatureValue) Set(val *AiUsageByFeatureValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAiUsageByFeatureValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAiUsageByFeatureValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiUsageByFeatureValue(val *AiUsageByFeatureValue) *NullableAiUsageByFeatureValue {
	return &NullableAiUsageByFeatureValue{value: val, isSet: true}
}

func (v NullableAiUsageByFeatureValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiUsageByFeatureValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



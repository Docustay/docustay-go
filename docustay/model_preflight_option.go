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

// checks if the PreflightOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreflightOption{}

// PreflightOption struct for PreflightOption
type PreflightOption struct {
	Key string `json:"key"`
	Label string `json:"label"`
	// Whole cents; null when the option has no price of its own (raising the budget).
	PriceMinor NullableInt32 `json:"priceMinor"`
	// `team member a month`, `month`, `once` or null.
	Per NullableString `json:"per"`
	// Deprecated. Always false: there is one Pro price. Kept so existing clients keep working.
	// Deprecated
	Founders *bool `json:"founders,omitempty"`
	Unlocks []string `json:"unlocks"`
	AdditionalProperties map[string]interface{}
}

type _PreflightOption PreflightOption

// NewPreflightOption instantiates a new PreflightOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreflightOption(key string, label string, priceMinor NullableInt32, per NullableString, unlocks []string) *PreflightOption {
	this := PreflightOption{}
	this.Key = key
	this.Label = label
	this.PriceMinor = priceMinor
	this.Per = per
	this.Unlocks = unlocks
	return &this
}

// NewPreflightOptionWithDefaults instantiates a new PreflightOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreflightOptionWithDefaults() *PreflightOption {
	this := PreflightOption{}
	return &this
}

// GetKey returns the Key field value
func (o *PreflightOption) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PreflightOption) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PreflightOption) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value
func (o *PreflightOption) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PreflightOption) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PreflightOption) SetLabel(v string) {
	o.Label = v
}

// GetPriceMinor returns the PriceMinor field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PreflightOption) GetPriceMinor() int32 {
	if o == nil || o.PriceMinor.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PriceMinor.Get()
}

// GetPriceMinorOk returns a tuple with the PriceMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreflightOption) GetPriceMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceMinor.Get(), o.PriceMinor.IsSet()
}

// SetPriceMinor sets field value
func (o *PreflightOption) SetPriceMinor(v int32) {
	o.PriceMinor.Set(&v)
}

// GetPer returns the Per field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PreflightOption) GetPer() string {
	if o == nil || o.Per.Get() == nil {
		var ret string
		return ret
	}

	return *o.Per.Get()
}

// GetPerOk returns a tuple with the Per field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreflightOption) GetPerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Per.Get(), o.Per.IsSet()
}

// SetPer sets field value
func (o *PreflightOption) SetPer(v string) {
	o.Per.Set(&v)
}

// GetFounders returns the Founders field value if set, zero value otherwise.
// Deprecated
func (o *PreflightOption) GetFounders() bool {
	if o == nil || IsNil(o.Founders) {
		var ret bool
		return ret
	}
	return *o.Founders
}

// GetFoundersOk returns a tuple with the Founders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PreflightOption) GetFoundersOk() (*bool, bool) {
	if o == nil || IsNil(o.Founders) {
		return nil, false
	}
	return o.Founders, true
}

// HasFounders returns a boolean if a field has been set.
func (o *PreflightOption) HasFounders() bool {
	if o != nil && !IsNil(o.Founders) {
		return true
	}

	return false
}

// SetFounders gets a reference to the given bool and assigns it to the Founders field.
// Deprecated
func (o *PreflightOption) SetFounders(v bool) {
	o.Founders = &v
}

// GetUnlocks returns the Unlocks field value
func (o *PreflightOption) GetUnlocks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Unlocks
}

// GetUnlocksOk returns a tuple with the Unlocks field value
// and a boolean to check if the value has been set.
func (o *PreflightOption) GetUnlocksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unlocks, true
}

// SetUnlocks sets field value
func (o *PreflightOption) SetUnlocks(v []string) {
	o.Unlocks = v
}

func (o PreflightOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreflightOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["label"] = o.Label
	toSerialize["priceMinor"] = o.PriceMinor.Get()
	toSerialize["per"] = o.Per.Get()
	if !IsNil(o.Founders) {
		toSerialize["founders"] = o.Founders
	}
	toSerialize["unlocks"] = o.Unlocks

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreflightOption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"label",
		"priceMinor",
		"per",
		"unlocks",
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

	varPreflightOption := _PreflightOption{}

	err = json.Unmarshal(data, &varPreflightOption)

	if err != nil {
		return err
	}

	*o = PreflightOption(varPreflightOption)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "label")
		delete(additionalProperties, "priceMinor")
		delete(additionalProperties, "per")
		delete(additionalProperties, "founders")
		delete(additionalProperties, "unlocks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreflightOption struct {
	value *PreflightOption
	isSet bool
}

func (v NullablePreflightOption) Get() *PreflightOption {
	return v.value
}

func (v *NullablePreflightOption) Set(val *PreflightOption) {
	v.value = val
	v.isSet = true
}

func (v NullablePreflightOption) IsSet() bool {
	return v.isSet
}

func (v *NullablePreflightOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreflightOption(val *PreflightOption) *NullablePreflightOption {
	return &NullablePreflightOption{value: val, isSet: true}
}

func (v NullablePreflightOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreflightOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



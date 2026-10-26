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

// checks if the PreflightAnswerCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreflightAnswerCost{}

// PreflightAnswerCost struct for PreflightAnswerCost
type PreflightAnswerCost struct {
	Minor int32 `json:"minor"`
	Currency string `json:"currency"`
	Units float32 `json:"units"`
	PerUnitMinor NullableFloat32 `json:"perUnitMinor"`
	Per NullableString `json:"per"`
	Prorated bool `json:"prorated"`
	FeePercent *float32 `json:"feePercent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreflightAnswerCost PreflightAnswerCost

// NewPreflightAnswerCost instantiates a new PreflightAnswerCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreflightAnswerCost(minor int32, currency string, units float32, perUnitMinor NullableFloat32, per NullableString, prorated bool) *PreflightAnswerCost {
	this := PreflightAnswerCost{}
	this.Minor = minor
	this.Currency = currency
	this.Units = units
	this.PerUnitMinor = perUnitMinor
	this.Per = per
	this.Prorated = prorated
	return &this
}

// NewPreflightAnswerCostWithDefaults instantiates a new PreflightAnswerCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreflightAnswerCostWithDefaults() *PreflightAnswerCost {
	this := PreflightAnswerCost{}
	return &this
}

// GetMinor returns the Minor field value
func (o *PreflightAnswerCost) GetMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minor
}

// GetMinorOk returns a tuple with the Minor field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerCost) GetMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minor, true
}

// SetMinor sets field value
func (o *PreflightAnswerCost) SetMinor(v int32) {
	o.Minor = v
}

// GetCurrency returns the Currency field value
func (o *PreflightAnswerCost) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerCost) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PreflightAnswerCost) SetCurrency(v string) {
	o.Currency = v
}

// GetUnits returns the Units field value
func (o *PreflightAnswerCost) GetUnits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerCost) GetUnitsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *PreflightAnswerCost) SetUnits(v float32) {
	o.Units = v
}

// GetPerUnitMinor returns the PerUnitMinor field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PreflightAnswerCost) GetPerUnitMinor() float32 {
	if o == nil || o.PerUnitMinor.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PerUnitMinor.Get()
}

// GetPerUnitMinorOk returns a tuple with the PerUnitMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreflightAnswerCost) GetPerUnitMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerUnitMinor.Get(), o.PerUnitMinor.IsSet()
}

// SetPerUnitMinor sets field value
func (o *PreflightAnswerCost) SetPerUnitMinor(v float32) {
	o.PerUnitMinor.Set(&v)
}

// GetPer returns the Per field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PreflightAnswerCost) GetPer() string {
	if o == nil || o.Per.Get() == nil {
		var ret string
		return ret
	}

	return *o.Per.Get()
}

// GetPerOk returns a tuple with the Per field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreflightAnswerCost) GetPerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Per.Get(), o.Per.IsSet()
}

// SetPer sets field value
func (o *PreflightAnswerCost) SetPer(v string) {
	o.Per.Set(&v)
}

// GetProrated returns the Prorated field value
func (o *PreflightAnswerCost) GetProrated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Prorated
}

// GetProratedOk returns a tuple with the Prorated field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswerCost) GetProratedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prorated, true
}

// SetProrated sets field value
func (o *PreflightAnswerCost) SetProrated(v bool) {
	o.Prorated = v
}

// GetFeePercent returns the FeePercent field value if set, zero value otherwise.
func (o *PreflightAnswerCost) GetFeePercent() float32 {
	if o == nil || IsNil(o.FeePercent) {
		var ret float32
		return ret
	}
	return *o.FeePercent
}

// GetFeePercentOk returns a tuple with the FeePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightAnswerCost) GetFeePercentOk() (*float32, bool) {
	if o == nil || IsNil(o.FeePercent) {
		return nil, false
	}
	return o.FeePercent, true
}

// HasFeePercent returns a boolean if a field has been set.
func (o *PreflightAnswerCost) HasFeePercent() bool {
	if o != nil && !IsNil(o.FeePercent) {
		return true
	}

	return false
}

// SetFeePercent gets a reference to the given float32 and assigns it to the FeePercent field.
func (o *PreflightAnswerCost) SetFeePercent(v float32) {
	o.FeePercent = &v
}

func (o PreflightAnswerCost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreflightAnswerCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minor"] = o.Minor
	toSerialize["currency"] = o.Currency
	toSerialize["units"] = o.Units
	toSerialize["perUnitMinor"] = o.PerUnitMinor.Get()
	toSerialize["per"] = o.Per.Get()
	toSerialize["prorated"] = o.Prorated
	if !IsNil(o.FeePercent) {
		toSerialize["feePercent"] = o.FeePercent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreflightAnswerCost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"minor",
		"currency",
		"units",
		"perUnitMinor",
		"per",
		"prorated",
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

	varPreflightAnswerCost := _PreflightAnswerCost{}

	err = json.Unmarshal(data, &varPreflightAnswerCost)

	if err != nil {
		return err
	}

	*o = PreflightAnswerCost(varPreflightAnswerCost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "units")
		delete(additionalProperties, "perUnitMinor")
		delete(additionalProperties, "per")
		delete(additionalProperties, "prorated")
		delete(additionalProperties, "feePercent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreflightAnswerCost struct {
	value *PreflightAnswerCost
	isSet bool
}

func (v NullablePreflightAnswerCost) Get() *PreflightAnswerCost {
	return v.value
}

func (v *NullablePreflightAnswerCost) Set(val *PreflightAnswerCost) {
	v.value = val
	v.isSet = true
}

func (v NullablePreflightAnswerCost) IsSet() bool {
	return v.isSet
}

func (v *NullablePreflightAnswerCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreflightAnswerCost(val *PreflightAnswerCost) *NullablePreflightAnswerCost {
	return &NullablePreflightAnswerCost{value: val, isSet: true}
}

func (v NullablePreflightAnswerCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreflightAnswerCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



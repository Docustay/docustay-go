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

// checks if the AiUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiUsage{}

// AiUsage struct for AiUsage
type AiUsage struct {
	Mode string `json:"mode"`
	Plan NullableString `json:"plan"`
	Allowance NullableFloat32 `json:"allowance"`
	Used NullableFloat32 `json:"used"`
	Remaining NullableFloat32 `json:"remaining"`
	Percent NullableFloat32 `json:"percent"`
	State string `json:"state"`
	ResetsAt NullableString `json:"resetsAt"`
	ByFeature map[string]AiUsageByFeatureValue `json:"byFeature"`
	Paused bool `json:"paused"`
	AdditionalProperties map[string]interface{}
}

type _AiUsage AiUsage

// NewAiUsage instantiates a new AiUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiUsage(mode string, plan NullableString, allowance NullableFloat32, used NullableFloat32, remaining NullableFloat32, percent NullableFloat32, state string, resetsAt NullableString, byFeature map[string]AiUsageByFeatureValue, paused bool) *AiUsage {
	this := AiUsage{}
	this.Mode = mode
	this.Plan = plan
	this.Allowance = allowance
	this.Used = used
	this.Remaining = remaining
	this.Percent = percent
	this.State = state
	this.ResetsAt = resetsAt
	this.ByFeature = byFeature
	this.Paused = paused
	return &this
}

// NewAiUsageWithDefaults instantiates a new AiUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiUsageWithDefaults() *AiUsage {
	this := AiUsage{}
	return &this
}

// GetMode returns the Mode field value
func (o *AiUsage) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AiUsage) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *AiUsage) SetMode(v string) {
	o.Mode = v
}

// GetPlan returns the Plan field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiUsage) GetPlan() string {
	if o == nil || o.Plan.Get() == nil {
		var ret string
		return ret
	}

	return *o.Plan.Get()
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiUsage) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plan.Get(), o.Plan.IsSet()
}

// SetPlan sets field value
func (o *AiUsage) SetPlan(v string) {
	o.Plan.Set(&v)
}

// GetAllowance returns the Allowance field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *AiUsage) GetAllowance() float32 {
	if o == nil || o.Allowance.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Allowance.Get()
}

// GetAllowanceOk returns a tuple with the Allowance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiUsage) GetAllowanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Allowance.Get(), o.Allowance.IsSet()
}

// SetAllowance sets field value
func (o *AiUsage) SetAllowance(v float32) {
	o.Allowance.Set(&v)
}

// GetUsed returns the Used field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *AiUsage) GetUsed() float32 {
	if o == nil || o.Used.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Used.Get()
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiUsage) GetUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Used.Get(), o.Used.IsSet()
}

// SetUsed sets field value
func (o *AiUsage) SetUsed(v float32) {
	o.Used.Set(&v)
}

// GetRemaining returns the Remaining field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *AiUsage) GetRemaining() float32 {
	if o == nil || o.Remaining.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Remaining.Get()
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiUsage) GetRemainingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remaining.Get(), o.Remaining.IsSet()
}

// SetRemaining sets field value
func (o *AiUsage) SetRemaining(v float32) {
	o.Remaining.Set(&v)
}

// GetPercent returns the Percent field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *AiUsage) GetPercent() float32 {
	if o == nil || o.Percent.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Percent.Get()
}

// GetPercentOk returns a tuple with the Percent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiUsage) GetPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Percent.Get(), o.Percent.IsSet()
}

// SetPercent sets field value
func (o *AiUsage) SetPercent(v float32) {
	o.Percent.Set(&v)
}

// GetState returns the State field value
func (o *AiUsage) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AiUsage) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *AiUsage) SetState(v string) {
	o.State = v
}

// GetResetsAt returns the ResetsAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiUsage) GetResetsAt() string {
	if o == nil || o.ResetsAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResetsAt.Get()
}

// GetResetsAtOk returns a tuple with the ResetsAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiUsage) GetResetsAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResetsAt.Get(), o.ResetsAt.IsSet()
}

// SetResetsAt sets field value
func (o *AiUsage) SetResetsAt(v string) {
	o.ResetsAt.Set(&v)
}

// GetByFeature returns the ByFeature field value
func (o *AiUsage) GetByFeature() map[string]AiUsageByFeatureValue {
	if o == nil {
		var ret map[string]AiUsageByFeatureValue
		return ret
	}

	return o.ByFeature
}

// GetByFeatureOk returns a tuple with the ByFeature field value
// and a boolean to check if the value has been set.
func (o *AiUsage) GetByFeatureOk() (map[string]AiUsageByFeatureValue, bool) {
	if o == nil {
		return map[string]AiUsageByFeatureValue{}, false
	}
	return o.ByFeature, true
}

// SetByFeature sets field value
func (o *AiUsage) SetByFeature(v map[string]AiUsageByFeatureValue) {
	o.ByFeature = v
}

// GetPaused returns the Paused field value
func (o *AiUsage) GetPaused() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Paused
}

// GetPausedOk returns a tuple with the Paused field value
// and a boolean to check if the value has been set.
func (o *AiUsage) GetPausedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Paused, true
}

// SetPaused sets field value
func (o *AiUsage) SetPaused(v bool) {
	o.Paused = v
}

func (o AiUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mode"] = o.Mode
	toSerialize["plan"] = o.Plan.Get()
	toSerialize["allowance"] = o.Allowance.Get()
	toSerialize["used"] = o.Used.Get()
	toSerialize["remaining"] = o.Remaining.Get()
	toSerialize["percent"] = o.Percent.Get()
	toSerialize["state"] = o.State
	toSerialize["resetsAt"] = o.ResetsAt.Get()
	toSerialize["byFeature"] = o.ByFeature
	toSerialize["paused"] = o.Paused

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiUsage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mode",
		"plan",
		"allowance",
		"used",
		"remaining",
		"percent",
		"state",
		"resetsAt",
		"byFeature",
		"paused",
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

	varAiUsage := _AiUsage{}

	err = json.Unmarshal(data, &varAiUsage)

	if err != nil {
		return err
	}

	*o = AiUsage(varAiUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mode")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "allowance")
		delete(additionalProperties, "used")
		delete(additionalProperties, "remaining")
		delete(additionalProperties, "percent")
		delete(additionalProperties, "state")
		delete(additionalProperties, "resetsAt")
		delete(additionalProperties, "byFeature")
		delete(additionalProperties, "paused")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiUsage struct {
	value *AiUsage
	isSet bool
}

func (v NullableAiUsage) Get() *AiUsage {
	return v.value
}

func (v *NullableAiUsage) Set(val *AiUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableAiUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableAiUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiUsage(val *AiUsage) *NullableAiUsage {
	return &NullableAiUsage{value: val, isSet: true}
}

func (v NullableAiUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



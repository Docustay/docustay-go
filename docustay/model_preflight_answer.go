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

// checks if the PreflightAnswer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreflightAnswer{}

// PreflightAnswer struct for PreflightAnswer
type PreflightAnswer struct {
	Action string `json:"action"`
	Allowed bool `json:"allowed"`
	Plan NullablePreflightAnswerPlan `json:"plan"`
	Period NullablePreflightAnswerPeriod `json:"period"`
	Meter map[string]*interface{} `json:"meter"`
	Cost PreflightAnswerCost `json:"cost"`
	AfterThis map[string]*interface{} `json:"afterThis"`
	Severity string `json:"severity"`
	Reason NullableString `json:"reason"`
	Options []PreflightOption `json:"options"`
	Message string `json:"message"`
	// For `bulk_send` on a limit: how many of the rows can go now.
	AllowedCount *int32 `json:"allowedCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreflightAnswer PreflightAnswer

// NewPreflightAnswer instantiates a new PreflightAnswer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreflightAnswer(action string, allowed bool, plan NullablePreflightAnswerPlan, period NullablePreflightAnswerPeriod, meter map[string]*interface{}, cost PreflightAnswerCost, afterThis map[string]*interface{}, severity string, reason NullableString, options []PreflightOption, message string) *PreflightAnswer {
	this := PreflightAnswer{}
	this.Action = action
	this.Allowed = allowed
	this.Plan = plan
	this.Period = period
	this.Meter = meter
	this.Cost = cost
	this.AfterThis = afterThis
	this.Severity = severity
	this.Reason = reason
	this.Options = options
	this.Message = message
	return &this
}

// NewPreflightAnswerWithDefaults instantiates a new PreflightAnswer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreflightAnswerWithDefaults() *PreflightAnswer {
	this := PreflightAnswer{}
	return &this
}

// GetAction returns the Action field value
func (o *PreflightAnswer) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswer) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PreflightAnswer) SetAction(v string) {
	o.Action = v
}

// GetAllowed returns the Allowed field value
func (o *PreflightAnswer) GetAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswer) GetAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowed, true
}

// SetAllowed sets field value
func (o *PreflightAnswer) SetAllowed(v bool) {
	o.Allowed = v
}

// GetPlan returns the Plan field value
// If the value is explicit nil, the zero value for PreflightAnswerPlan will be returned
func (o *PreflightAnswer) GetPlan() PreflightAnswerPlan {
	if o == nil || o.Plan.Get() == nil {
		var ret PreflightAnswerPlan
		return ret
	}

	return *o.Plan.Get()
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreflightAnswer) GetPlanOk() (*PreflightAnswerPlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plan.Get(), o.Plan.IsSet()
}

// SetPlan sets field value
func (o *PreflightAnswer) SetPlan(v PreflightAnswerPlan) {
	o.Plan.Set(&v)
}

// GetPeriod returns the Period field value
// If the value is explicit nil, the zero value for PreflightAnswerPeriod will be returned
func (o *PreflightAnswer) GetPeriod() PreflightAnswerPeriod {
	if o == nil || o.Period.Get() == nil {
		var ret PreflightAnswerPeriod
		return ret
	}

	return *o.Period.Get()
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreflightAnswer) GetPeriodOk() (*PreflightAnswerPeriod, bool) {
	if o == nil {
		return nil, false
	}
	return o.Period.Get(), o.Period.IsSet()
}

// SetPeriod sets field value
func (o *PreflightAnswer) SetPeriod(v PreflightAnswerPeriod) {
	o.Period.Set(&v)
}

// GetMeter returns the Meter field value
// If the value is explicit nil, the zero value for map[string]*interface{} will be returned
func (o *PreflightAnswer) GetMeter() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.Meter
}

// GetMeterOk returns a tuple with the Meter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreflightAnswer) GetMeterOk() (map[string]*interface{}, bool) {
	if o == nil || IsNil(o.Meter) {
		return map[string]*interface{}{}, false
	}
	return o.Meter, true
}

// SetMeter sets field value
func (o *PreflightAnswer) SetMeter(v map[string]*interface{}) {
	o.Meter = v
}

// GetCost returns the Cost field value
func (o *PreflightAnswer) GetCost() PreflightAnswerCost {
	if o == nil {
		var ret PreflightAnswerCost
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswer) GetCostOk() (*PreflightAnswerCost, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *PreflightAnswer) SetCost(v PreflightAnswerCost) {
	o.Cost = v
}

// GetAfterThis returns the AfterThis field value
// If the value is explicit nil, the zero value for map[string]*interface{} will be returned
func (o *PreflightAnswer) GetAfterThis() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.AfterThis
}

// GetAfterThisOk returns a tuple with the AfterThis field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreflightAnswer) GetAfterThisOk() (map[string]*interface{}, bool) {
	if o == nil || IsNil(o.AfterThis) {
		return map[string]*interface{}{}, false
	}
	return o.AfterThis, true
}

// SetAfterThis sets field value
func (o *PreflightAnswer) SetAfterThis(v map[string]*interface{}) {
	o.AfterThis = v
}

// GetSeverity returns the Severity field value
func (o *PreflightAnswer) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswer) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *PreflightAnswer) SetSeverity(v string) {
	o.Severity = v
}

// GetReason returns the Reason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PreflightAnswer) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreflightAnswer) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// SetReason sets field value
func (o *PreflightAnswer) SetReason(v string) {
	o.Reason.Set(&v)
}

// GetOptions returns the Options field value
func (o *PreflightAnswer) GetOptions() []PreflightOption {
	if o == nil {
		var ret []PreflightOption
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswer) GetOptionsOk() ([]PreflightOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *PreflightAnswer) SetOptions(v []PreflightOption) {
	o.Options = v
}

// GetMessage returns the Message field value
func (o *PreflightAnswer) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PreflightAnswer) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PreflightAnswer) SetMessage(v string) {
	o.Message = v
}

// GetAllowedCount returns the AllowedCount field value if set, zero value otherwise.
func (o *PreflightAnswer) GetAllowedCount() int32 {
	if o == nil || IsNil(o.AllowedCount) {
		var ret int32
		return ret
	}
	return *o.AllowedCount
}

// GetAllowedCountOk returns a tuple with the AllowedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightAnswer) GetAllowedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowedCount) {
		return nil, false
	}
	return o.AllowedCount, true
}

// HasAllowedCount returns a boolean if a field has been set.
func (o *PreflightAnswer) HasAllowedCount() bool {
	if o != nil && !IsNil(o.AllowedCount) {
		return true
	}

	return false
}

// SetAllowedCount gets a reference to the given int32 and assigns it to the AllowedCount field.
func (o *PreflightAnswer) SetAllowedCount(v int32) {
	o.AllowedCount = &v
}

func (o PreflightAnswer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreflightAnswer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["allowed"] = o.Allowed
	toSerialize["plan"] = o.Plan.Get()
	toSerialize["period"] = o.Period.Get()
	if o.Meter != nil {
		toSerialize["meter"] = o.Meter
	}
	toSerialize["cost"] = o.Cost
	if o.AfterThis != nil {
		toSerialize["afterThis"] = o.AfterThis
	}
	toSerialize["severity"] = o.Severity
	toSerialize["reason"] = o.Reason.Get()
	toSerialize["options"] = o.Options
	toSerialize["message"] = o.Message
	if !IsNil(o.AllowedCount) {
		toSerialize["allowedCount"] = o.AllowedCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreflightAnswer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"allowed",
		"plan",
		"period",
		"meter",
		"cost",
		"afterThis",
		"severity",
		"reason",
		"options",
		"message",
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

	varPreflightAnswer := _PreflightAnswer{}

	err = json.Unmarshal(data, &varPreflightAnswer)

	if err != nil {
		return err
	}

	*o = PreflightAnswer(varPreflightAnswer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "allowed")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "period")
		delete(additionalProperties, "meter")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "afterThis")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "options")
		delete(additionalProperties, "message")
		delete(additionalProperties, "allowedCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreflightAnswer struct {
	value *PreflightAnswer
	isSet bool
}

func (v NullablePreflightAnswer) Get() *PreflightAnswer {
	return v.value
}

func (v *NullablePreflightAnswer) Set(val *PreflightAnswer) {
	v.value = val
	v.isSet = true
}

func (v NullablePreflightAnswer) IsSet() bool {
	return v.isSet
}

func (v *NullablePreflightAnswer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreflightAnswer(val *PreflightAnswer) *NullablePreflightAnswer {
	return &NullablePreflightAnswer{value: val, isSet: true}
}

func (v NullablePreflightAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreflightAnswer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the UsageEmails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageEmails{}

// UsageEmails struct for UsageEmails
type UsageEmails struct {
	UsedToday int32 `json:"usedToday"`
	// The most emails this workspace may send today. In a workspace's first 7 days it is lower than the plan's limits.emailsPerDay (10 on Free, 100 on a paid plan).
	LimitPerDay int32 `json:"limitPerDay"`
	ResetsAt string `json:"resetsAt"`
}

type _UsageEmails UsageEmails

// NewUsageEmails instantiates a new UsageEmails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageEmails(usedToday int32, limitPerDay int32, resetsAt string) *UsageEmails {
	this := UsageEmails{}
	this.UsedToday = usedToday
	this.LimitPerDay = limitPerDay
	this.ResetsAt = resetsAt
	return &this
}

// NewUsageEmailsWithDefaults instantiates a new UsageEmails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageEmailsWithDefaults() *UsageEmails {
	this := UsageEmails{}
	return &this
}

// GetUsedToday returns the UsedToday field value
func (o *UsageEmails) GetUsedToday() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsedToday
}

// GetUsedTodayOk returns a tuple with the UsedToday field value
// and a boolean to check if the value has been set.
func (o *UsageEmails) GetUsedTodayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedToday, true
}

// SetUsedToday sets field value
func (o *UsageEmails) SetUsedToday(v int32) {
	o.UsedToday = v
}

// GetLimitPerDay returns the LimitPerDay field value
func (o *UsageEmails) GetLimitPerDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LimitPerDay
}

// GetLimitPerDayOk returns a tuple with the LimitPerDay field value
// and a boolean to check if the value has been set.
func (o *UsageEmails) GetLimitPerDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitPerDay, true
}

// SetLimitPerDay sets field value
func (o *UsageEmails) SetLimitPerDay(v int32) {
	o.LimitPerDay = v
}

// GetResetsAt returns the ResetsAt field value
func (o *UsageEmails) GetResetsAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetsAt
}

// GetResetsAtOk returns a tuple with the ResetsAt field value
// and a boolean to check if the value has been set.
func (o *UsageEmails) GetResetsAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetsAt, true
}

// SetResetsAt sets field value
func (o *UsageEmails) SetResetsAt(v string) {
	o.ResetsAt = v
}

func (o UsageEmails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageEmails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["usedToday"] = o.UsedToday
	toSerialize["limitPerDay"] = o.LimitPerDay
	toSerialize["resetsAt"] = o.ResetsAt
	return toSerialize, nil
}

func (o *UsageEmails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"usedToday",
		"limitPerDay",
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

	varUsageEmails := _UsageEmails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageEmails)

	if err != nil {
		return err
	}

	*o = UsageEmails(varUsageEmails)

	return err
}

type NullableUsageEmails struct {
	value *UsageEmails
	isSet bool
}

func (v NullableUsageEmails) Get() *UsageEmails {
	return v.value
}

func (v *NullableUsageEmails) Set(val *UsageEmails) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageEmails) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageEmails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageEmails(val *UsageEmails) *NullableUsageEmails {
	return &NullableUsageEmails{value: val, isSet: true}
}

func (v NullableUsageEmails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageEmails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



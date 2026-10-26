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

// checks if the SendDocument201ResponseNotEmailedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendDocument201ResponseNotEmailedInner{}

// SendDocument201ResponseNotEmailedInner struct for SendDocument201ResponseNotEmailedInner
type SendDocument201ResponseNotEmailedInner struct {
	Seat string `json:"seat"`
	Email string `json:"email"`
	Reason string `json:"reason"`
	Detail NullableString `json:"detail"`
	AdditionalProperties map[string]interface{}
}

type _SendDocument201ResponseNotEmailedInner SendDocument201ResponseNotEmailedInner

// NewSendDocument201ResponseNotEmailedInner instantiates a new SendDocument201ResponseNotEmailedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendDocument201ResponseNotEmailedInner(seat string, email string, reason string, detail NullableString) *SendDocument201ResponseNotEmailedInner {
	this := SendDocument201ResponseNotEmailedInner{}
	this.Seat = seat
	this.Email = email
	this.Reason = reason
	this.Detail = detail
	return &this
}

// NewSendDocument201ResponseNotEmailedInnerWithDefaults instantiates a new SendDocument201ResponseNotEmailedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendDocument201ResponseNotEmailedInnerWithDefaults() *SendDocument201ResponseNotEmailedInner {
	this := SendDocument201ResponseNotEmailedInner{}
	return &this
}

// GetSeat returns the Seat field value
func (o *SendDocument201ResponseNotEmailedInner) GetSeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seat
}

// GetSeatOk returns a tuple with the Seat field value
// and a boolean to check if the value has been set.
func (o *SendDocument201ResponseNotEmailedInner) GetSeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seat, true
}

// SetSeat sets field value
func (o *SendDocument201ResponseNotEmailedInner) SetSeat(v string) {
	o.Seat = v
}

// GetEmail returns the Email field value
func (o *SendDocument201ResponseNotEmailedInner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SendDocument201ResponseNotEmailedInner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SendDocument201ResponseNotEmailedInner) SetEmail(v string) {
	o.Email = v
}

// GetReason returns the Reason field value
func (o *SendDocument201ResponseNotEmailedInner) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SendDocument201ResponseNotEmailedInner) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SendDocument201ResponseNotEmailedInner) SetReason(v string) {
	o.Reason = v
}

// GetDetail returns the Detail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SendDocument201ResponseNotEmailedInner) GetDetail() string {
	if o == nil || o.Detail.Get() == nil {
		var ret string
		return ret
	}

	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendDocument201ResponseNotEmailedInner) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// SetDetail sets field value
func (o *SendDocument201ResponseNotEmailedInner) SetDetail(v string) {
	o.Detail.Set(&v)
}

func (o SendDocument201ResponseNotEmailedInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendDocument201ResponseNotEmailedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seat"] = o.Seat
	toSerialize["email"] = o.Email
	toSerialize["reason"] = o.Reason
	toSerialize["detail"] = o.Detail.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SendDocument201ResponseNotEmailedInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seat",
		"email",
		"reason",
		"detail",
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

	varSendDocument201ResponseNotEmailedInner := _SendDocument201ResponseNotEmailedInner{}

	err = json.Unmarshal(data, &varSendDocument201ResponseNotEmailedInner)

	if err != nil {
		return err
	}

	*o = SendDocument201ResponseNotEmailedInner(varSendDocument201ResponseNotEmailedInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "seat")
		delete(additionalProperties, "email")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSendDocument201ResponseNotEmailedInner struct {
	value *SendDocument201ResponseNotEmailedInner
	isSet bool
}

func (v NullableSendDocument201ResponseNotEmailedInner) Get() *SendDocument201ResponseNotEmailedInner {
	return v.value
}

func (v *NullableSendDocument201ResponseNotEmailedInner) Set(val *SendDocument201ResponseNotEmailedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSendDocument201ResponseNotEmailedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSendDocument201ResponseNotEmailedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendDocument201ResponseNotEmailedInner(val *SendDocument201ResponseNotEmailedInner) *NullableSendDocument201ResponseNotEmailedInner {
	return &NullableSendDocument201ResponseNotEmailedInner{value: val, isSet: true}
}

func (v NullableSendDocument201ResponseNotEmailedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendDocument201ResponseNotEmailedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



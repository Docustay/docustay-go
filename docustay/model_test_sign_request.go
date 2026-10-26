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

// checks if the TestSignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSignRequest{}

// TestSignRequest struct for TestSignRequest
type TestSignRequest struct {
	// The seat to sign, for example `signer_1`.
	Seat string `json:"seat"`
	// Decline the seat instead of signing it: the document becomes `declined` and `documents.document.declined` fires, as when a person declines.
	Decline *bool `json:"decline,omitempty"`
	// With `decline`, the reason the sender sees.
	Reason *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestSignRequest TestSignRequest

// NewTestSignRequest instantiates a new TestSignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSignRequest(seat string) *TestSignRequest {
	this := TestSignRequest{}
	this.Seat = seat
	return &this
}

// NewTestSignRequestWithDefaults instantiates a new TestSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSignRequestWithDefaults() *TestSignRequest {
	this := TestSignRequest{}
	return &this
}

// GetSeat returns the Seat field value
func (o *TestSignRequest) GetSeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seat
}

// GetSeatOk returns a tuple with the Seat field value
// and a boolean to check if the value has been set.
func (o *TestSignRequest) GetSeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seat, true
}

// SetSeat sets field value
func (o *TestSignRequest) SetSeat(v string) {
	o.Seat = v
}

// GetDecline returns the Decline field value if set, zero value otherwise.
func (o *TestSignRequest) GetDecline() bool {
	if o == nil || IsNil(o.Decline) {
		var ret bool
		return ret
	}
	return *o.Decline
}

// GetDeclineOk returns a tuple with the Decline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSignRequest) GetDeclineOk() (*bool, bool) {
	if o == nil || IsNil(o.Decline) {
		return nil, false
	}
	return o.Decline, true
}

// HasDecline returns a boolean if a field has been set.
func (o *TestSignRequest) HasDecline() bool {
	if o != nil && !IsNil(o.Decline) {
		return true
	}

	return false
}

// SetDecline gets a reference to the given bool and assigns it to the Decline field.
func (o *TestSignRequest) SetDecline(v bool) {
	o.Decline = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TestSignRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSignRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TestSignRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TestSignRequest) SetReason(v string) {
	o.Reason = &v
}

func (o TestSignRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seat"] = o.Seat
	if !IsNil(o.Decline) {
		toSerialize["decline"] = o.Decline
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestSignRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seat",
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

	varTestSignRequest := _TestSignRequest{}

	err = json.Unmarshal(data, &varTestSignRequest)

	if err != nil {
		return err
	}

	*o = TestSignRequest(varTestSignRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "seat")
		delete(additionalProperties, "decline")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestSignRequest struct {
	value *TestSignRequest
	isSet bool
}

func (v NullableTestSignRequest) Get() *TestSignRequest {
	return v.value
}

func (v *NullableTestSignRequest) Set(val *TestSignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSignRequest(val *TestSignRequest) *NullableTestSignRequest {
	return &NullableTestSignRequest{value: val, isSet: true}
}

func (v NullableTestSignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



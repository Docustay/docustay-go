/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the RemindRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemindRequest{}

// RemindRequest struct for RemindRequest
type RemindRequest struct {
	// Remind only this seat (for example signer_2). Leave out to remind everyone who is up.
	Seat *string `json:"seat,omitempty"`
}

// NewRemindRequest instantiates a new RemindRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemindRequest() *RemindRequest {
	this := RemindRequest{}
	return &this
}

// NewRemindRequestWithDefaults instantiates a new RemindRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemindRequestWithDefaults() *RemindRequest {
	this := RemindRequest{}
	return &this
}

// GetSeat returns the Seat field value if set, zero value otherwise.
func (o *RemindRequest) GetSeat() string {
	if o == nil || IsNil(o.Seat) {
		var ret string
		return ret
	}
	return *o.Seat
}

// GetSeatOk returns a tuple with the Seat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemindRequest) GetSeatOk() (*string, bool) {
	if o == nil || IsNil(o.Seat) {
		return nil, false
	}
	return o.Seat, true
}

// HasSeat returns a boolean if a field has been set.
func (o *RemindRequest) HasSeat() bool {
	if o != nil && !IsNil(o.Seat) {
		return true
	}

	return false
}

// SetSeat gets a reference to the given string and assigns it to the Seat field.
func (o *RemindRequest) SetSeat(v string) {
	o.Seat = &v
}

func (o RemindRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemindRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seat) {
		toSerialize["seat"] = o.Seat
	}
	return toSerialize, nil
}

type NullableRemindRequest struct {
	value *RemindRequest
	isSet bool
}

func (v NullableRemindRequest) Get() *RemindRequest {
	return v.value
}

func (v *NullableRemindRequest) Set(val *RemindRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemindRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemindRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemindRequest(val *RemindRequest) *NullableRemindRequest {
	return &NullableRemindRequest{value: val, isSet: true}
}

func (v NullableRemindRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemindRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



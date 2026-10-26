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

// checks if the EmbedSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbedSessionRequest{}

// EmbedSessionRequest struct for EmbedSessionRequest
type EmbedSessionRequest struct {
	Seat string `json:"seat"`
	// The website that will show the form; must be on your allowed list.
	Origin string `json:"origin"`
}

type _EmbedSessionRequest EmbedSessionRequest

// NewEmbedSessionRequest instantiates a new EmbedSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbedSessionRequest(seat string, origin string) *EmbedSessionRequest {
	this := EmbedSessionRequest{}
	this.Seat = seat
	this.Origin = origin
	return &this
}

// NewEmbedSessionRequestWithDefaults instantiates a new EmbedSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbedSessionRequestWithDefaults() *EmbedSessionRequest {
	this := EmbedSessionRequest{}
	return &this
}

// GetSeat returns the Seat field value
func (o *EmbedSessionRequest) GetSeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seat
}

// GetSeatOk returns a tuple with the Seat field value
// and a boolean to check if the value has been set.
func (o *EmbedSessionRequest) GetSeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seat, true
}

// SetSeat sets field value
func (o *EmbedSessionRequest) SetSeat(v string) {
	o.Seat = v
}

// GetOrigin returns the Origin field value
func (o *EmbedSessionRequest) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *EmbedSessionRequest) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *EmbedSessionRequest) SetOrigin(v string) {
	o.Origin = v
}

func (o EmbedSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbedSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seat"] = o.Seat
	toSerialize["origin"] = o.Origin
	return toSerialize, nil
}

func (o *EmbedSessionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seat",
		"origin",
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

	varEmbedSessionRequest := _EmbedSessionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmbedSessionRequest)

	if err != nil {
		return err
	}

	*o = EmbedSessionRequest(varEmbedSessionRequest)

	return err
}

type NullableEmbedSessionRequest struct {
	value *EmbedSessionRequest
	isSet bool
}

func (v NullableEmbedSessionRequest) Get() *EmbedSessionRequest {
	return v.value
}

func (v *NullableEmbedSessionRequest) Set(val *EmbedSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbedSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbedSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbedSessionRequest(val *EmbedSessionRequest) *NullableEmbedSessionRequest {
	return &NullableEmbedSessionRequest{value: val, isSet: true}
}

func (v NullableEmbedSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbedSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



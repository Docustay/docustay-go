/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the RotateWebhookSecretRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RotateWebhookSecretRequest{}

// RotateWebhookSecretRequest struct for RotateWebhookSecretRequest
type RotateWebhookSecretRequest struct {
	WindowHours *int32 `json:"windowHours,omitempty"`
}

// NewRotateWebhookSecretRequest instantiates a new RotateWebhookSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotateWebhookSecretRequest() *RotateWebhookSecretRequest {
	this := RotateWebhookSecretRequest{}
	return &this
}

// NewRotateWebhookSecretRequestWithDefaults instantiates a new RotateWebhookSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotateWebhookSecretRequestWithDefaults() *RotateWebhookSecretRequest {
	this := RotateWebhookSecretRequest{}
	return &this
}

// GetWindowHours returns the WindowHours field value if set, zero value otherwise.
func (o *RotateWebhookSecretRequest) GetWindowHours() int32 {
	if o == nil || IsNil(o.WindowHours) {
		var ret int32
		return ret
	}
	return *o.WindowHours
}

// GetWindowHoursOk returns a tuple with the WindowHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotateWebhookSecretRequest) GetWindowHoursOk() (*int32, bool) {
	if o == nil || IsNil(o.WindowHours) {
		return nil, false
	}
	return o.WindowHours, true
}

// HasWindowHours returns a boolean if a field has been set.
func (o *RotateWebhookSecretRequest) HasWindowHours() bool {
	if o != nil && !IsNil(o.WindowHours) {
		return true
	}

	return false
}

// SetWindowHours gets a reference to the given int32 and assigns it to the WindowHours field.
func (o *RotateWebhookSecretRequest) SetWindowHours(v int32) {
	o.WindowHours = &v
}

func (o RotateWebhookSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RotateWebhookSecretRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WindowHours) {
		toSerialize["windowHours"] = o.WindowHours
	}
	return toSerialize, nil
}

type NullableRotateWebhookSecretRequest struct {
	value *RotateWebhookSecretRequest
	isSet bool
}

func (v NullableRotateWebhookSecretRequest) Get() *RotateWebhookSecretRequest {
	return v.value
}

func (v *NullableRotateWebhookSecretRequest) Set(val *RotateWebhookSecretRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRotateWebhookSecretRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRotateWebhookSecretRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotateWebhookSecretRequest(val *RotateWebhookSecretRequest) *NullableRotateWebhookSecretRequest {
	return &NullableRotateWebhookSecretRequest{value: val, isSet: true}
}

func (v NullableRotateWebhookSecretRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotateWebhookSecretRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



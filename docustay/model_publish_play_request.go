/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the PublishPlayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublishPlayRequest{}

// PublishPlayRequest struct for PublishPlayRequest
type PublishPlayRequest struct {
	Version *int32 `json:"version,omitempty"`
}

// NewPublishPlayRequest instantiates a new PublishPlayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishPlayRequest() *PublishPlayRequest {
	this := PublishPlayRequest{}
	return &this
}

// NewPublishPlayRequestWithDefaults instantiates a new PublishPlayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishPlayRequestWithDefaults() *PublishPlayRequest {
	this := PublishPlayRequest{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PublishPlayRequest) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishPlayRequest) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PublishPlayRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *PublishPlayRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o PublishPlayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublishPlayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullablePublishPlayRequest struct {
	value *PublishPlayRequest
	isSet bool
}

func (v NullablePublishPlayRequest) Get() *PublishPlayRequest {
	return v.value
}

func (v *NullablePublishPlayRequest) Set(val *PublishPlayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishPlayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishPlayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishPlayRequest(val *PublishPlayRequest) *NullablePublishPlayRequest {
	return &NullablePublishPlayRequest{value: val, isSet: true}
}

func (v NullablePublishPlayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishPlayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



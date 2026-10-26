/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the VoidDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoidDocumentRequest{}

// VoidDocumentRequest struct for VoidDocumentRequest
type VoidDocumentRequest struct {
	Reason *string `json:"reason,omitempty"`
}

// NewVoidDocumentRequest instantiates a new VoidDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoidDocumentRequest() *VoidDocumentRequest {
	this := VoidDocumentRequest{}
	return &this
}

// NewVoidDocumentRequestWithDefaults instantiates a new VoidDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoidDocumentRequestWithDefaults() *VoidDocumentRequest {
	this := VoidDocumentRequest{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *VoidDocumentRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidDocumentRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *VoidDocumentRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *VoidDocumentRequest) SetReason(v string) {
	o.Reason = &v
}

func (o VoidDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoidDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableVoidDocumentRequest struct {
	value *VoidDocumentRequest
	isSet bool
}

func (v NullableVoidDocumentRequest) Get() *VoidDocumentRequest {
	return v.value
}

func (v *NullableVoidDocumentRequest) Set(val *VoidDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVoidDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVoidDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoidDocumentRequest(val *VoidDocumentRequest) *NullableVoidDocumentRequest {
	return &NullableVoidDocumentRequest{value: val, isSet: true}
}

func (v NullableVoidDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoidDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



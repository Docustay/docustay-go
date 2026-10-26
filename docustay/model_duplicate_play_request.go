/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the DuplicatePlayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DuplicatePlayRequest{}

// DuplicatePlayRequest struct for DuplicatePlayRequest
type DuplicatePlayRequest struct {
	As *string `json:"as,omitempty" validate:"regexp=^[a-z0-9][a-z0-9-]{2\\,60}$"`
}

// NewDuplicatePlayRequest instantiates a new DuplicatePlayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuplicatePlayRequest() *DuplicatePlayRequest {
	this := DuplicatePlayRequest{}
	return &this
}

// NewDuplicatePlayRequestWithDefaults instantiates a new DuplicatePlayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuplicatePlayRequestWithDefaults() *DuplicatePlayRequest {
	this := DuplicatePlayRequest{}
	return &this
}

// GetAs returns the As field value if set, zero value otherwise.
func (o *DuplicatePlayRequest) GetAs() string {
	if o == nil || IsNil(o.As) {
		var ret string
		return ret
	}
	return *o.As
}

// GetAsOk returns a tuple with the As field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicatePlayRequest) GetAsOk() (*string, bool) {
	if o == nil || IsNil(o.As) {
		return nil, false
	}
	return o.As, true
}

// HasAs returns a boolean if a field has been set.
func (o *DuplicatePlayRequest) HasAs() bool {
	if o != nil && !IsNil(o.As) {
		return true
	}

	return false
}

// SetAs gets a reference to the given string and assigns it to the As field.
func (o *DuplicatePlayRequest) SetAs(v string) {
	o.As = &v
}

func (o DuplicatePlayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DuplicatePlayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.As) {
		toSerialize["as"] = o.As
	}
	return toSerialize, nil
}

type NullableDuplicatePlayRequest struct {
	value *DuplicatePlayRequest
	isSet bool
}

func (v NullableDuplicatePlayRequest) Get() *DuplicatePlayRequest {
	return v.value
}

func (v *NullableDuplicatePlayRequest) Set(val *DuplicatePlayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDuplicatePlayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDuplicatePlayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuplicatePlayRequest(val *DuplicatePlayRequest) *NullableDuplicatePlayRequest {
	return &NullableDuplicatePlayRequest{value: val, isSet: true}
}

func (v NullableDuplicatePlayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuplicatePlayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



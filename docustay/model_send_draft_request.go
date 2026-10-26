/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the SendDraftRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendDraftRequest{}

// SendDraftRequest struct for SendDraftRequest
type SendDraftRequest struct {
	// The boxes the sender fills in (\"Filled in by: Sender\") and the template's send-time amounts, by key: a box's label in lower case with underscores (\"Effective date\" → `effective_date`). A send that leaves one empty answers 400 naming it.
	Values map[string]string `json:"values,omitempty"`
}

// NewSendDraftRequest instantiates a new SendDraftRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendDraftRequest() *SendDraftRequest {
	this := SendDraftRequest{}
	return &this
}

// NewSendDraftRequestWithDefaults instantiates a new SendDraftRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendDraftRequestWithDefaults() *SendDraftRequest {
	this := SendDraftRequest{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SendDraftRequest) GetValues() map[string]string {
	if o == nil || IsNil(o.Values) {
		var ret map[string]string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendDraftRequest) GetValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Values) {
		return map[string]string{}, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SendDraftRequest) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]string and assigns it to the Values field.
func (o *SendDraftRequest) SetValues(v map[string]string) {
	o.Values = v
}

func (o SendDraftRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendDraftRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableSendDraftRequest struct {
	value *SendDraftRequest
	isSet bool
}

func (v NullableSendDraftRequest) Get() *SendDraftRequest {
	return v.value
}

func (v *NullableSendDraftRequest) Set(val *SendDraftRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendDraftRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendDraftRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendDraftRequest(val *SendDraftRequest) *NullableSendDraftRequest {
	return &NullableSendDraftRequest{value: val, isSet: true}
}

func (v NullableSendDraftRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendDraftRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



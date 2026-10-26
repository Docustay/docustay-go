/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the SetTemplateDefaultsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetTemplateDefaultsRequest{}

// SetTemplateDefaultsRequest struct for SetTemplateDefaultsRequest
type SetTemplateDefaultsRequest struct {
	Description NullableString `json:"description,omitempty"`
	ClientDescription NullableString `json:"clientDescription,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Send *SendChoices `json:"send,omitempty"`
}

// NewSetTemplateDefaultsRequest instantiates a new SetTemplateDefaultsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetTemplateDefaultsRequest() *SetTemplateDefaultsRequest {
	this := SetTemplateDefaultsRequest{}
	return &this
}

// NewSetTemplateDefaultsRequestWithDefaults instantiates a new SetTemplateDefaultsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetTemplateDefaultsRequestWithDefaults() *SetTemplateDefaultsRequest {
	this := SetTemplateDefaultsRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetTemplateDefaultsRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetTemplateDefaultsRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SetTemplateDefaultsRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SetTemplateDefaultsRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SetTemplateDefaultsRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SetTemplateDefaultsRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetClientDescription returns the ClientDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetTemplateDefaultsRequest) GetClientDescription() string {
	if o == nil || IsNil(o.ClientDescription.Get()) {
		var ret string
		return ret
	}
	return *o.ClientDescription.Get()
}

// GetClientDescriptionOk returns a tuple with the ClientDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetTemplateDefaultsRequest) GetClientDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientDescription.Get(), o.ClientDescription.IsSet()
}

// HasClientDescription returns a boolean if a field has been set.
func (o *SetTemplateDefaultsRequest) HasClientDescription() bool {
	if o != nil && o.ClientDescription.IsSet() {
		return true
	}

	return false
}

// SetClientDescription gets a reference to the given NullableString and assigns it to the ClientDescription field.
func (o *SetTemplateDefaultsRequest) SetClientDescription(v string) {
	o.ClientDescription.Set(&v)
}
// SetClientDescriptionNil sets the value for ClientDescription to be an explicit nil
func (o *SetTemplateDefaultsRequest) SetClientDescriptionNil() {
	o.ClientDescription.Set(nil)
}

// UnsetClientDescription ensures that no value is present for ClientDescription, not even an explicit nil
func (o *SetTemplateDefaultsRequest) UnsetClientDescription() {
	o.ClientDescription.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetTemplateDefaultsRequest) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetTemplateDefaultsRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *SetTemplateDefaultsRequest) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *SetTemplateDefaultsRequest) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *SetTemplateDefaultsRequest) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *SetTemplateDefaultsRequest) UnsetMessage() {
	o.Message.Unset()
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *SetTemplateDefaultsRequest) GetSend() SendChoices {
	if o == nil || IsNil(o.Send) {
		var ret SendChoices
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetTemplateDefaultsRequest) GetSendOk() (*SendChoices, bool) {
	if o == nil || IsNil(o.Send) {
		return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *SetTemplateDefaultsRequest) HasSend() bool {
	if o != nil && !IsNil(o.Send) {
		return true
	}

	return false
}

// SetSend gets a reference to the given SendChoices and assigns it to the Send field.
func (o *SetTemplateDefaultsRequest) SetSend(v SendChoices) {
	o.Send = &v
}

func (o SetTemplateDefaultsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetTemplateDefaultsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ClientDescription.IsSet() {
		toSerialize["clientDescription"] = o.ClientDescription.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if !IsNil(o.Send) {
		toSerialize["send"] = o.Send
	}
	return toSerialize, nil
}

type NullableSetTemplateDefaultsRequest struct {
	value *SetTemplateDefaultsRequest
	isSet bool
}

func (v NullableSetTemplateDefaultsRequest) Get() *SetTemplateDefaultsRequest {
	return v.value
}

func (v *NullableSetTemplateDefaultsRequest) Set(val *SetTemplateDefaultsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetTemplateDefaultsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetTemplateDefaultsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetTemplateDefaultsRequest(val *SetTemplateDefaultsRequest) *NullableSetTemplateDefaultsRequest {
	return &NullableSetTemplateDefaultsRequest{value: val, isSet: true}
}

func (v NullableSetTemplateDefaultsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetTemplateDefaultsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



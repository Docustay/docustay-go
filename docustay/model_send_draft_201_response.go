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

// checks if the SendDraft201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendDraft201Response{}

// SendDraft201Response struct for SendDraft201Response
type SendDraft201Response struct {
	Id string `json:"id"`
	State string `json:"state"`
	Emailed int32 `json:"emailed"`
	NotEmailed []interface{} `json:"notEmailed"`
	AdditionalProperties map[string]interface{}
}

type _SendDraft201Response SendDraft201Response

// NewSendDraft201Response instantiates a new SendDraft201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendDraft201Response(id string, state string, emailed int32, notEmailed []interface{}) *SendDraft201Response {
	this := SendDraft201Response{}
	this.Id = id
	this.State = state
	this.Emailed = emailed
	this.NotEmailed = notEmailed
	return &this
}

// NewSendDraft201ResponseWithDefaults instantiates a new SendDraft201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendDraft201ResponseWithDefaults() *SendDraft201Response {
	this := SendDraft201Response{}
	return &this
}

// GetId returns the Id field value
func (o *SendDraft201Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SendDraft201Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SendDraft201Response) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *SendDraft201Response) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SendDraft201Response) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SendDraft201Response) SetState(v string) {
	o.State = v
}

// GetEmailed returns the Emailed field value
func (o *SendDraft201Response) GetEmailed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Emailed
}

// GetEmailedOk returns a tuple with the Emailed field value
// and a boolean to check if the value has been set.
func (o *SendDraft201Response) GetEmailedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Emailed, true
}

// SetEmailed sets field value
func (o *SendDraft201Response) SetEmailed(v int32) {
	o.Emailed = v
}

// GetNotEmailed returns the NotEmailed field value
func (o *SendDraft201Response) GetNotEmailed() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.NotEmailed
}

// GetNotEmailedOk returns a tuple with the NotEmailed field value
// and a boolean to check if the value has been set.
func (o *SendDraft201Response) GetNotEmailedOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotEmailed, true
}

// SetNotEmailed sets field value
func (o *SendDraft201Response) SetNotEmailed(v []interface{}) {
	o.NotEmailed = v
}

func (o SendDraft201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendDraft201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	toSerialize["emailed"] = o.Emailed
	toSerialize["notEmailed"] = o.NotEmailed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SendDraft201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
		"emailed",
		"notEmailed",
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

	varSendDraft201Response := _SendDraft201Response{}

	err = json.Unmarshal(data, &varSendDraft201Response)

	if err != nil {
		return err
	}

	*o = SendDraft201Response(varSendDraft201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "emailed")
		delete(additionalProperties, "notEmailed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSendDraft201Response struct {
	value *SendDraft201Response
	isSet bool
}

func (v NullableSendDraft201Response) Get() *SendDraft201Response {
	return v.value
}

func (v *NullableSendDraft201Response) Set(val *SendDraft201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSendDraft201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSendDraft201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendDraft201Response(val *SendDraft201Response) *NullableSendDraft201Response {
	return &NullableSendDraft201Response{value: val, isSet: true}
}

func (v NullableSendDraft201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendDraft201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



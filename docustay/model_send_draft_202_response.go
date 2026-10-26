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

// checks if the SendDraft202Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendDraft202Response{}

// SendDraft202Response struct for SendDraft202Response
type SendDraft202Response struct {
	Id string `json:"id"`
	State string `json:"state"`
	Approval SendDraft202ResponseApproval `json:"approval"`
	AdditionalProperties map[string]interface{}
}

type _SendDraft202Response SendDraft202Response

// NewSendDraft202Response instantiates a new SendDraft202Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendDraft202Response(id string, state string, approval SendDraft202ResponseApproval) *SendDraft202Response {
	this := SendDraft202Response{}
	this.Id = id
	this.State = state
	this.Approval = approval
	return &this
}

// NewSendDraft202ResponseWithDefaults instantiates a new SendDraft202Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendDraft202ResponseWithDefaults() *SendDraft202Response {
	this := SendDraft202Response{}
	return &this
}

// GetId returns the Id field value
func (o *SendDraft202Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SendDraft202Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SendDraft202Response) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *SendDraft202Response) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SendDraft202Response) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SendDraft202Response) SetState(v string) {
	o.State = v
}

// GetApproval returns the Approval field value
func (o *SendDraft202Response) GetApproval() SendDraft202ResponseApproval {
	if o == nil {
		var ret SendDraft202ResponseApproval
		return ret
	}

	return o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value
// and a boolean to check if the value has been set.
func (o *SendDraft202Response) GetApprovalOk() (*SendDraft202ResponseApproval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Approval, true
}

// SetApproval sets field value
func (o *SendDraft202Response) SetApproval(v SendDraft202ResponseApproval) {
	o.Approval = v
}

func (o SendDraft202Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendDraft202Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	toSerialize["approval"] = o.Approval

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SendDraft202Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
		"approval",
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

	varSendDraft202Response := _SendDraft202Response{}

	err = json.Unmarshal(data, &varSendDraft202Response)

	if err != nil {
		return err
	}

	*o = SendDraft202Response(varSendDraft202Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "approval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSendDraft202Response struct {
	value *SendDraft202Response
	isSet bool
}

func (v NullableSendDraft202Response) Get() *SendDraft202Response {
	return v.value
}

func (v *NullableSendDraft202Response) Set(val *SendDraft202Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSendDraft202Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSendDraft202Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendDraft202Response(val *SendDraft202Response) *NullableSendDraft202Response {
	return &NullableSendDraft202Response{value: val, isSet: true}
}

func (v NullableSendDraft202Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendDraft202Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



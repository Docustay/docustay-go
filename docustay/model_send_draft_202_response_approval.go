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

// checks if the SendDraft202ResponseApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendDraft202ResponseApproval{}

// SendDraft202ResponseApproval Set when a draft-only key asked to send: a person approves or declines in the app, and the document stays a draft until then.
type SendDraft202ResponseApproval struct {
	Id string `json:"id"`
	Status string `json:"status"`
	InstanceId *string `json:"instanceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SendDraft202ResponseApproval SendDraft202ResponseApproval

// NewSendDraft202ResponseApproval instantiates a new SendDraft202ResponseApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendDraft202ResponseApproval(id string, status string) *SendDraft202ResponseApproval {
	this := SendDraft202ResponseApproval{}
	this.Id = id
	this.Status = status
	return &this
}

// NewSendDraft202ResponseApprovalWithDefaults instantiates a new SendDraft202ResponseApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendDraft202ResponseApprovalWithDefaults() *SendDraft202ResponseApproval {
	this := SendDraft202ResponseApproval{}
	return &this
}

// GetId returns the Id field value
func (o *SendDraft202ResponseApproval) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SendDraft202ResponseApproval) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SendDraft202ResponseApproval) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *SendDraft202ResponseApproval) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SendDraft202ResponseApproval) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SendDraft202ResponseApproval) SetStatus(v string) {
	o.Status = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *SendDraft202ResponseApproval) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendDraft202ResponseApproval) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *SendDraft202ResponseApproval) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *SendDraft202ResponseApproval) SetInstanceId(v string) {
	o.InstanceId = &v
}

func (o SendDraft202ResponseApproval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendDraft202ResponseApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SendDraft202ResponseApproval) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
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

	varSendDraft202ResponseApproval := _SendDraft202ResponseApproval{}

	err = json.Unmarshal(data, &varSendDraft202ResponseApproval)

	if err != nil {
		return err
	}

	*o = SendDraft202ResponseApproval(varSendDraft202ResponseApproval)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "instanceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSendDraft202ResponseApproval struct {
	value *SendDraft202ResponseApproval
	isSet bool
}

func (v NullableSendDraft202ResponseApproval) Get() *SendDraft202ResponseApproval {
	return v.value
}

func (v *NullableSendDraft202ResponseApproval) Set(val *SendDraft202ResponseApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableSendDraft202ResponseApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableSendDraft202ResponseApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendDraft202ResponseApproval(val *SendDraft202ResponseApproval) *NullableSendDraft202ResponseApproval {
	return &NullableSendDraft202ResponseApproval{value: val, isSet: true}
}

func (v NullableSendDraft202ResponseApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendDraft202ResponseApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



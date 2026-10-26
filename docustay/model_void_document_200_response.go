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

// checks if the VoidDocument200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoidDocument200Response{}

// VoidDocument200Response struct for VoidDocument200Response
type VoidDocument200Response struct {
	Id string `json:"id"`
	State string `json:"state"`
	// How many people were emailed that it is cancelled. A test document sends one [TEST] notice to its sender.
	Notified int32 `json:"notified"`
	AdditionalProperties map[string]interface{}
}

type _VoidDocument200Response VoidDocument200Response

// NewVoidDocument200Response instantiates a new VoidDocument200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoidDocument200Response(id string, state string, notified int32) *VoidDocument200Response {
	this := VoidDocument200Response{}
	this.Id = id
	this.State = state
	this.Notified = notified
	return &this
}

// NewVoidDocument200ResponseWithDefaults instantiates a new VoidDocument200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoidDocument200ResponseWithDefaults() *VoidDocument200Response {
	this := VoidDocument200Response{}
	return &this
}

// GetId returns the Id field value
func (o *VoidDocument200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VoidDocument200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VoidDocument200Response) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *VoidDocument200Response) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VoidDocument200Response) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VoidDocument200Response) SetState(v string) {
	o.State = v
}

// GetNotified returns the Notified field value
func (o *VoidDocument200Response) GetNotified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Notified
}

// GetNotifiedOk returns a tuple with the Notified field value
// and a boolean to check if the value has been set.
func (o *VoidDocument200Response) GetNotifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notified, true
}

// SetNotified sets field value
func (o *VoidDocument200Response) SetNotified(v int32) {
	o.Notified = v
}

func (o VoidDocument200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoidDocument200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	toSerialize["notified"] = o.Notified

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VoidDocument200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
		"notified",
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

	varVoidDocument200Response := _VoidDocument200Response{}

	err = json.Unmarshal(data, &varVoidDocument200Response)

	if err != nil {
		return err
	}

	*o = VoidDocument200Response(varVoidDocument200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "notified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVoidDocument200Response struct {
	value *VoidDocument200Response
	isSet bool
}

func (v NullableVoidDocument200Response) Get() *VoidDocument200Response {
	return v.value
}

func (v *NullableVoidDocument200Response) Set(val *VoidDocument200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVoidDocument200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVoidDocument200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoidDocument200Response(val *VoidDocument200Response) *NullableVoidDocument200Response {
	return &NullableVoidDocument200Response{value: val, isSet: true}
}

func (v NullableVoidDocument200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoidDocument200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the DeleteWebhook200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteWebhook200Response{}

// DeleteWebhook200Response struct for DeleteWebhook200Response
type DeleteWebhook200Response struct {
	Ok bool `json:"ok"`
	AdditionalProperties map[string]interface{}
}

type _DeleteWebhook200Response DeleteWebhook200Response

// NewDeleteWebhook200Response instantiates a new DeleteWebhook200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteWebhook200Response(ok bool) *DeleteWebhook200Response {
	this := DeleteWebhook200Response{}
	this.Ok = ok
	return &this
}

// NewDeleteWebhook200ResponseWithDefaults instantiates a new DeleteWebhook200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteWebhook200ResponseWithDefaults() *DeleteWebhook200Response {
	this := DeleteWebhook200Response{}
	return &this
}

// GetOk returns the Ok field value
func (o *DeleteWebhook200Response) GetOk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value
// and a boolean to check if the value has been set.
func (o *DeleteWebhook200Response) GetOkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ok, true
}

// SetOk sets field value
func (o *DeleteWebhook200Response) SetOk(v bool) {
	o.Ok = v
}

func (o DeleteWebhook200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteWebhook200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ok"] = o.Ok

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteWebhook200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ok",
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

	varDeleteWebhook200Response := _DeleteWebhook200Response{}

	err = json.Unmarshal(data, &varDeleteWebhook200Response)

	if err != nil {
		return err
	}

	*o = DeleteWebhook200Response(varDeleteWebhook200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ok")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteWebhook200Response struct {
	value *DeleteWebhook200Response
	isSet bool
}

func (v NullableDeleteWebhook200Response) Get() *DeleteWebhook200Response {
	return v.value
}

func (v *NullableDeleteWebhook200Response) Set(val *DeleteWebhook200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteWebhook200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteWebhook200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteWebhook200Response(val *DeleteWebhook200Response) *NullableDeleteWebhook200Response {
	return &NullableDeleteWebhook200Response{value: val, isSet: true}
}

func (v NullableDeleteWebhook200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteWebhook200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



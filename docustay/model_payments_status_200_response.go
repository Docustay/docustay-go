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

// checks if the PaymentsStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentsStatus200Response{}

// PaymentsStatus200Response struct for PaymentsStatus200Response
type PaymentsStatus200Response struct {
	Status string `json:"status"`
	Mode NullableString `json:"mode"`
	AdditionalProperties map[string]interface{}
}

type _PaymentsStatus200Response PaymentsStatus200Response

// NewPaymentsStatus200Response instantiates a new PaymentsStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentsStatus200Response(status string, mode NullableString) *PaymentsStatus200Response {
	this := PaymentsStatus200Response{}
	this.Status = status
	this.Mode = mode
	return &this
}

// NewPaymentsStatus200ResponseWithDefaults instantiates a new PaymentsStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentsStatus200ResponseWithDefaults() *PaymentsStatus200Response {
	this := PaymentsStatus200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *PaymentsStatus200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentsStatus200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentsStatus200Response) SetStatus(v string) {
	o.Status = v
}

// GetMode returns the Mode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentsStatus200Response) GetMode() string {
	if o == nil || o.Mode.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentsStatus200Response) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// SetMode sets field value
func (o *PaymentsStatus200Response) SetMode(v string) {
	o.Mode.Set(&v)
}

func (o PaymentsStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentsStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["mode"] = o.Mode.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentsStatus200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"mode",
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

	varPaymentsStatus200Response := _PaymentsStatus200Response{}

	err = json.Unmarshal(data, &varPaymentsStatus200Response)

	if err != nil {
		return err
	}

	*o = PaymentsStatus200Response(varPaymentsStatus200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentsStatus200Response struct {
	value *PaymentsStatus200Response
	isSet bool
}

func (v NullablePaymentsStatus200Response) Get() *PaymentsStatus200Response {
	return v.value
}

func (v *NullablePaymentsStatus200Response) Set(val *PaymentsStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentsStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentsStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentsStatus200Response(val *PaymentsStatus200Response) *NullablePaymentsStatus200Response {
	return &NullablePaymentsStatus200Response{value: val, isSet: true}
}

func (v NullablePaymentsStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentsStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



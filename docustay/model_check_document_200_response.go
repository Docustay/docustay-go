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

// checks if the CheckDocument200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckDocument200Response{}

// CheckDocument200Response struct for CheckDocument200Response
type CheckDocument200Response struct {
	Ok bool `json:"ok"`
	Errors []string `json:"errors"`
	Warnings []interface{} `json:"warnings"`
	Roles []interface{} `json:"roles"`
	AdditionalProperties map[string]interface{}
}

type _CheckDocument200Response CheckDocument200Response

// NewCheckDocument200Response instantiates a new CheckDocument200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckDocument200Response(ok bool, errors []string, warnings []interface{}, roles []interface{}) *CheckDocument200Response {
	this := CheckDocument200Response{}
	this.Ok = ok
	this.Errors = errors
	this.Warnings = warnings
	this.Roles = roles
	return &this
}

// NewCheckDocument200ResponseWithDefaults instantiates a new CheckDocument200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckDocument200ResponseWithDefaults() *CheckDocument200Response {
	this := CheckDocument200Response{}
	return &this
}

// GetOk returns the Ok field value
func (o *CheckDocument200Response) GetOk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value
// and a boolean to check if the value has been set.
func (o *CheckDocument200Response) GetOkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ok, true
}

// SetOk sets field value
func (o *CheckDocument200Response) SetOk(v bool) {
	o.Ok = v
}

// GetErrors returns the Errors field value
func (o *CheckDocument200Response) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *CheckDocument200Response) GetErrorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *CheckDocument200Response) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value
func (o *CheckDocument200Response) GetWarnings() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *CheckDocument200Response) GetWarningsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *CheckDocument200Response) SetWarnings(v []interface{}) {
	o.Warnings = v
}

// GetRoles returns the Roles field value
func (o *CheckDocument200Response) GetRoles() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *CheckDocument200Response) GetRolesOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *CheckDocument200Response) SetRoles(v []interface{}) {
	o.Roles = v
}

func (o CheckDocument200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckDocument200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ok"] = o.Ok
	toSerialize["errors"] = o.Errors
	toSerialize["warnings"] = o.Warnings
	toSerialize["roles"] = o.Roles

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckDocument200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ok",
		"errors",
		"warnings",
		"roles",
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

	varCheckDocument200Response := _CheckDocument200Response{}

	err = json.Unmarshal(data, &varCheckDocument200Response)

	if err != nil {
		return err
	}

	*o = CheckDocument200Response(varCheckDocument200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ok")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "roles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckDocument200Response struct {
	value *CheckDocument200Response
	isSet bool
}

func (v NullableCheckDocument200Response) Get() *CheckDocument200Response {
	return v.value
}

func (v *NullableCheckDocument200Response) Set(val *CheckDocument200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckDocument200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckDocument200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckDocument200Response(val *CheckDocument200Response) *NullableCheckDocument200Response {
	return &NullableCheckDocument200Response{value: val, isSet: true}
}

func (v NullableCheckDocument200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckDocument200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



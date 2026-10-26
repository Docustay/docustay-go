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

// checks if the ReviseDocument201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviseDocument201Response{}

// ReviseDocument201Response struct for ReviseDocument201Response
type ReviseDocument201Response struct {
	Id string `json:"id"`
	PreviousId string `json:"previousId"`
	Warnings []interface{} `json:"warnings"`
	AdditionalProperties map[string]interface{}
}

type _ReviseDocument201Response ReviseDocument201Response

// NewReviseDocument201Response instantiates a new ReviseDocument201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviseDocument201Response(id string, previousId string, warnings []interface{}) *ReviseDocument201Response {
	this := ReviseDocument201Response{}
	this.Id = id
	this.PreviousId = previousId
	this.Warnings = warnings
	return &this
}

// NewReviseDocument201ResponseWithDefaults instantiates a new ReviseDocument201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviseDocument201ResponseWithDefaults() *ReviseDocument201Response {
	this := ReviseDocument201Response{}
	return &this
}

// GetId returns the Id field value
func (o *ReviseDocument201Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviseDocument201Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviseDocument201Response) SetId(v string) {
	o.Id = v
}

// GetPreviousId returns the PreviousId field value
func (o *ReviseDocument201Response) GetPreviousId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousId
}

// GetPreviousIdOk returns a tuple with the PreviousId field value
// and a boolean to check if the value has been set.
func (o *ReviseDocument201Response) GetPreviousIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousId, true
}

// SetPreviousId sets field value
func (o *ReviseDocument201Response) SetPreviousId(v string) {
	o.PreviousId = v
}

// GetWarnings returns the Warnings field value
func (o *ReviseDocument201Response) GetWarnings() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *ReviseDocument201Response) GetWarningsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *ReviseDocument201Response) SetWarnings(v []interface{}) {
	o.Warnings = v
}

func (o ReviseDocument201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviseDocument201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["previousId"] = o.PreviousId
	toSerialize["warnings"] = o.Warnings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviseDocument201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"previousId",
		"warnings",
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

	varReviseDocument201Response := _ReviseDocument201Response{}

	err = json.Unmarshal(data, &varReviseDocument201Response)

	if err != nil {
		return err
	}

	*o = ReviseDocument201Response(varReviseDocument201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "previousId")
		delete(additionalProperties, "warnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviseDocument201Response struct {
	value *ReviseDocument201Response
	isSet bool
}

func (v NullableReviseDocument201Response) Get() *ReviseDocument201Response {
	return v.value
}

func (v *NullableReviseDocument201Response) Set(val *ReviseDocument201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReviseDocument201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReviseDocument201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviseDocument201Response(val *ReviseDocument201Response) *NullableReviseDocument201Response {
	return &NullableReviseDocument201Response{value: val, isSet: true}
}

func (v NullableReviseDocument201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviseDocument201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



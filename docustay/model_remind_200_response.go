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

// checks if the Remind200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Remind200Response{}

// Remind200Response struct for Remind200Response
type Remind200Response struct {
	Id string `json:"id"`
	Reminded int32 `json:"reminded"`
	AdditionalProperties map[string]interface{}
}

type _Remind200Response Remind200Response

// NewRemind200Response instantiates a new Remind200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemind200Response(id string, reminded int32) *Remind200Response {
	this := Remind200Response{}
	this.Id = id
	this.Reminded = reminded
	return &this
}

// NewRemind200ResponseWithDefaults instantiates a new Remind200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemind200ResponseWithDefaults() *Remind200Response {
	this := Remind200Response{}
	return &this
}

// GetId returns the Id field value
func (o *Remind200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Remind200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Remind200Response) SetId(v string) {
	o.Id = v
}

// GetReminded returns the Reminded field value
func (o *Remind200Response) GetReminded() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Reminded
}

// GetRemindedOk returns a tuple with the Reminded field value
// and a boolean to check if the value has been set.
func (o *Remind200Response) GetRemindedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reminded, true
}

// SetReminded sets field value
func (o *Remind200Response) SetReminded(v int32) {
	o.Reminded = v
}

func (o Remind200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Remind200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["reminded"] = o.Reminded

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Remind200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"reminded",
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

	varRemind200Response := _Remind200Response{}

	err = json.Unmarshal(data, &varRemind200Response)

	if err != nil {
		return err
	}

	*o = Remind200Response(varRemind200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "reminded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemind200Response struct {
	value *Remind200Response
	isSet bool
}

func (v NullableRemind200Response) Get() *Remind200Response {
	return v.value
}

func (v *NullableRemind200Response) Set(val *Remind200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRemind200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRemind200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemind200Response(val *Remind200Response) *NullableRemind200Response {
	return &NullableRemind200Response{value: val, isSet: true}
}

func (v NullableRemind200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemind200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



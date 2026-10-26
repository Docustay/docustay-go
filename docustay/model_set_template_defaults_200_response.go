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

// checks if the SetTemplateDefaults200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetTemplateDefaults200Response{}

// SetTemplateDefaults200Response struct for SetTemplateDefaults200Response
type SetTemplateDefaults200Response struct {
	TemplateId string `json:"templateId"`
	Saved bool `json:"saved"`
	AdditionalProperties map[string]interface{}
}

type _SetTemplateDefaults200Response SetTemplateDefaults200Response

// NewSetTemplateDefaults200Response instantiates a new SetTemplateDefaults200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetTemplateDefaults200Response(templateId string, saved bool) *SetTemplateDefaults200Response {
	this := SetTemplateDefaults200Response{}
	this.TemplateId = templateId
	this.Saved = saved
	return &this
}

// NewSetTemplateDefaults200ResponseWithDefaults instantiates a new SetTemplateDefaults200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetTemplateDefaults200ResponseWithDefaults() *SetTemplateDefaults200Response {
	this := SetTemplateDefaults200Response{}
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *SetTemplateDefaults200Response) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *SetTemplateDefaults200Response) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *SetTemplateDefaults200Response) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetSaved returns the Saved field value
func (o *SetTemplateDefaults200Response) GetSaved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Saved
}

// GetSavedOk returns a tuple with the Saved field value
// and a boolean to check if the value has been set.
func (o *SetTemplateDefaults200Response) GetSavedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Saved, true
}

// SetSaved sets field value
func (o *SetTemplateDefaults200Response) SetSaved(v bool) {
	o.Saved = v
}

func (o SetTemplateDefaults200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetTemplateDefaults200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateId"] = o.TemplateId
	toSerialize["saved"] = o.Saved

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetTemplateDefaults200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateId",
		"saved",
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

	varSetTemplateDefaults200Response := _SetTemplateDefaults200Response{}

	err = json.Unmarshal(data, &varSetTemplateDefaults200Response)

	if err != nil {
		return err
	}

	*o = SetTemplateDefaults200Response(varSetTemplateDefaults200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "saved")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetTemplateDefaults200Response struct {
	value *SetTemplateDefaults200Response
	isSet bool
}

func (v NullableSetTemplateDefaults200Response) Get() *SetTemplateDefaults200Response {
	return v.value
}

func (v *NullableSetTemplateDefaults200Response) Set(val *SetTemplateDefaults200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSetTemplateDefaults200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSetTemplateDefaults200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetTemplateDefaults200Response(val *SetTemplateDefaults200Response) *NullableSetTemplateDefaults200Response {
	return &NullableSetTemplateDefaults200Response{value: val, isSet: true}
}

func (v NullableSetTemplateDefaults200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetTemplateDefaults200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



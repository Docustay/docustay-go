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

// checks if the ListTemplates200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTemplates200Response{}

// ListTemplates200Response struct for ListTemplates200Response
type ListTemplates200Response struct {
	Templates []Template `json:"templates"`
	AdditionalProperties map[string]interface{}
}

type _ListTemplates200Response ListTemplates200Response

// NewListTemplates200Response instantiates a new ListTemplates200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTemplates200Response(templates []Template) *ListTemplates200Response {
	this := ListTemplates200Response{}
	this.Templates = templates
	return &this
}

// NewListTemplates200ResponseWithDefaults instantiates a new ListTemplates200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTemplates200ResponseWithDefaults() *ListTemplates200Response {
	this := ListTemplates200Response{}
	return &this
}

// GetTemplates returns the Templates field value
func (o *ListTemplates200Response) GetTemplates() []Template {
	if o == nil {
		var ret []Template
		return ret
	}

	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
func (o *ListTemplates200Response) GetTemplatesOk() ([]Template, bool) {
	if o == nil {
		return nil, false
	}
	return o.Templates, true
}

// SetTemplates sets field value
func (o *ListTemplates200Response) SetTemplates(v []Template) {
	o.Templates = v
}

func (o ListTemplates200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTemplates200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templates"] = o.Templates

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListTemplates200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templates",
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

	varListTemplates200Response := _ListTemplates200Response{}

	err = json.Unmarshal(data, &varListTemplates200Response)

	if err != nil {
		return err
	}

	*o = ListTemplates200Response(varListTemplates200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "templates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListTemplates200Response struct {
	value *ListTemplates200Response
	isSet bool
}

func (v NullableListTemplates200Response) Get() *ListTemplates200Response {
	return v.value
}

func (v *NullableListTemplates200Response) Set(val *ListTemplates200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListTemplates200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListTemplates200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTemplates200Response(val *ListTemplates200Response) *NullableListTemplates200Response {
	return &NullableListTemplates200Response{value: val, isSet: true}
}

func (v NullableListTemplates200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTemplates200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



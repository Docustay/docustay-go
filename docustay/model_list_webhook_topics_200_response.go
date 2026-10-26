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

// checks if the ListWebhookTopics200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhookTopics200Response{}

// ListWebhookTopics200Response struct for ListWebhookTopics200Response
type ListWebhookTopics200Response struct {
	Topics []string `json:"topics"`
	AdditionalProperties map[string]interface{}
}

type _ListWebhookTopics200Response ListWebhookTopics200Response

// NewListWebhookTopics200Response instantiates a new ListWebhookTopics200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhookTopics200Response(topics []string) *ListWebhookTopics200Response {
	this := ListWebhookTopics200Response{}
	this.Topics = topics
	return &this
}

// NewListWebhookTopics200ResponseWithDefaults instantiates a new ListWebhookTopics200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhookTopics200ResponseWithDefaults() *ListWebhookTopics200Response {
	this := ListWebhookTopics200Response{}
	return &this
}

// GetTopics returns the Topics field value
func (o *ListWebhookTopics200Response) GetTopics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value
// and a boolean to check if the value has been set.
func (o *ListWebhookTopics200Response) GetTopicsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Topics, true
}

// SetTopics sets field value
func (o *ListWebhookTopics200Response) SetTopics(v []string) {
	o.Topics = v
}

func (o ListWebhookTopics200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhookTopics200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["topics"] = o.Topics

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListWebhookTopics200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"topics",
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

	varListWebhookTopics200Response := _ListWebhookTopics200Response{}

	err = json.Unmarshal(data, &varListWebhookTopics200Response)

	if err != nil {
		return err
	}

	*o = ListWebhookTopics200Response(varListWebhookTopics200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "topics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListWebhookTopics200Response struct {
	value *ListWebhookTopics200Response
	isSet bool
}

func (v NullableListWebhookTopics200Response) Get() *ListWebhookTopics200Response {
	return v.value
}

func (v *NullableListWebhookTopics200Response) Set(val *ListWebhookTopics200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhookTopics200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhookTopics200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhookTopics200Response(val *ListWebhookTopics200Response) *NullableListWebhookTopics200Response {
	return &NullableListWebhookTopics200Response{value: val, isSet: true}
}

func (v NullableListWebhookTopics200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhookTopics200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ListWebhooks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhooks200Response{}

// ListWebhooks200Response struct for ListWebhooks200Response
type ListWebhooks200Response struct {
	Webhooks []WebhookEndpoint `json:"webhooks"`
	AdditionalProperties map[string]interface{}
}

type _ListWebhooks200Response ListWebhooks200Response

// NewListWebhooks200Response instantiates a new ListWebhooks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhooks200Response(webhooks []WebhookEndpoint) *ListWebhooks200Response {
	this := ListWebhooks200Response{}
	this.Webhooks = webhooks
	return &this
}

// NewListWebhooks200ResponseWithDefaults instantiates a new ListWebhooks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhooks200ResponseWithDefaults() *ListWebhooks200Response {
	this := ListWebhooks200Response{}
	return &this
}

// GetWebhooks returns the Webhooks field value
func (o *ListWebhooks200Response) GetWebhooks() []WebhookEndpoint {
	if o == nil {
		var ret []WebhookEndpoint
		return ret
	}

	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value
// and a boolean to check if the value has been set.
func (o *ListWebhooks200Response) GetWebhooksOk() ([]WebhookEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// SetWebhooks sets field value
func (o *ListWebhooks200Response) SetWebhooks(v []WebhookEndpoint) {
	o.Webhooks = v
}

func (o ListWebhooks200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhooks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["webhooks"] = o.Webhooks

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListWebhooks200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"webhooks",
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

	varListWebhooks200Response := _ListWebhooks200Response{}

	err = json.Unmarshal(data, &varListWebhooks200Response)

	if err != nil {
		return err
	}

	*o = ListWebhooks200Response(varListWebhooks200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "webhooks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListWebhooks200Response struct {
	value *ListWebhooks200Response
	isSet bool
}

func (v NullableListWebhooks200Response) Get() *ListWebhooks200Response {
	return v.value
}

func (v *NullableListWebhooks200Response) Set(val *ListWebhooks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhooks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhooks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhooks200Response(val *ListWebhooks200Response) *NullableListWebhooks200Response {
	return &NullableListWebhooks200Response{value: val, isSet: true}
}

func (v NullableListWebhooks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhooks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



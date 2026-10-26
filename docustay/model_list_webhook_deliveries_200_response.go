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

// checks if the ListWebhookDeliveries200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhookDeliveries200Response{}

// ListWebhookDeliveries200Response struct for ListWebhookDeliveries200Response
type ListWebhookDeliveries200Response struct {
	Deliveries []ListWebhookDeliveries200ResponseDeliveriesInner `json:"deliveries"`
	AdditionalProperties map[string]interface{}
}

type _ListWebhookDeliveries200Response ListWebhookDeliveries200Response

// NewListWebhookDeliveries200Response instantiates a new ListWebhookDeliveries200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhookDeliveries200Response(deliveries []ListWebhookDeliveries200ResponseDeliveriesInner) *ListWebhookDeliveries200Response {
	this := ListWebhookDeliveries200Response{}
	this.Deliveries = deliveries
	return &this
}

// NewListWebhookDeliveries200ResponseWithDefaults instantiates a new ListWebhookDeliveries200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhookDeliveries200ResponseWithDefaults() *ListWebhookDeliveries200Response {
	this := ListWebhookDeliveries200Response{}
	return &this
}

// GetDeliveries returns the Deliveries field value
func (o *ListWebhookDeliveries200Response) GetDeliveries() []ListWebhookDeliveries200ResponseDeliveriesInner {
	if o == nil {
		var ret []ListWebhookDeliveries200ResponseDeliveriesInner
		return ret
	}

	return o.Deliveries
}

// GetDeliveriesOk returns a tuple with the Deliveries field value
// and a boolean to check if the value has been set.
func (o *ListWebhookDeliveries200Response) GetDeliveriesOk() ([]ListWebhookDeliveries200ResponseDeliveriesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deliveries, true
}

// SetDeliveries sets field value
func (o *ListWebhookDeliveries200Response) SetDeliveries(v []ListWebhookDeliveries200ResponseDeliveriesInner) {
	o.Deliveries = v
}

func (o ListWebhookDeliveries200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhookDeliveries200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliveries"] = o.Deliveries

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListWebhookDeliveries200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deliveries",
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

	varListWebhookDeliveries200Response := _ListWebhookDeliveries200Response{}

	err = json.Unmarshal(data, &varListWebhookDeliveries200Response)

	if err != nil {
		return err
	}

	*o = ListWebhookDeliveries200Response(varListWebhookDeliveries200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deliveries")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListWebhookDeliveries200Response struct {
	value *ListWebhookDeliveries200Response
	isSet bool
}

func (v NullableListWebhookDeliveries200Response) Get() *ListWebhookDeliveries200Response {
	return v.value
}

func (v *NullableListWebhookDeliveries200Response) Set(val *ListWebhookDeliveries200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhookDeliveries200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhookDeliveries200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhookDeliveries200Response(val *ListWebhookDeliveries200Response) *NullableListWebhookDeliveries200Response {
	return &NullableListWebhookDeliveries200Response{value: val, isSet: true}
}

func (v NullableListWebhookDeliveries200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhookDeliveries200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



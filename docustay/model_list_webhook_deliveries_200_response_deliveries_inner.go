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

// checks if the ListWebhookDeliveries200ResponseDeliveriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhookDeliveries200ResponseDeliveriesInner{}

// ListWebhookDeliveries200ResponseDeliveriesInner struct for ListWebhookDeliveries200ResponseDeliveriesInner
type ListWebhookDeliveries200ResponseDeliveriesInner struct {
	Id string `json:"id"`
	Event string `json:"event"`
	Attempt int32 `json:"attempt"`
	Status string `json:"status"`
	ResponseStatus NullableFloat32 `json:"responseStatus"`
	Error NullableString `json:"error"`
	AttemptedAt string `json:"attemptedAt"`
	Test bool `json:"test"`
	AdditionalProperties map[string]interface{}
}

type _ListWebhookDeliveries200ResponseDeliveriesInner ListWebhookDeliveries200ResponseDeliveriesInner

// NewListWebhookDeliveries200ResponseDeliveriesInner instantiates a new ListWebhookDeliveries200ResponseDeliveriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhookDeliveries200ResponseDeliveriesInner(id string, event string, attempt int32, status string, responseStatus NullableFloat32, error_ NullableString, attemptedAt string, test bool) *ListWebhookDeliveries200ResponseDeliveriesInner {
	this := ListWebhookDeliveries200ResponseDeliveriesInner{}
	this.Id = id
	this.Event = event
	this.Attempt = attempt
	this.Status = status
	this.ResponseStatus = responseStatus
	this.Error = error_
	this.AttemptedAt = attemptedAt
	this.Test = test
	return &this
}

// NewListWebhookDeliveries200ResponseDeliveriesInnerWithDefaults instantiates a new ListWebhookDeliveries200ResponseDeliveriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhookDeliveries200ResponseDeliveriesInnerWithDefaults() *ListWebhookDeliveries200ResponseDeliveriesInner {
	this := ListWebhookDeliveries200ResponseDeliveriesInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetId(v string) {
	o.Id = v
}

// GetEvent returns the Event field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetEvent(v string) {
	o.Event = v
}

// GetAttempt returns the Attempt field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetAttempt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Attempt
}

// GetAttemptOk returns a tuple with the Attempt field value
// and a boolean to check if the value has been set.
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetAttemptOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attempt, true
}

// SetAttempt sets field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetAttempt(v int32) {
	o.Attempt = v
}

// GetStatus returns the Status field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetStatus(v string) {
	o.Status = v
}

// GetResponseStatus returns the ResponseStatus field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetResponseStatus() float32 {
	if o == nil || o.ResponseStatus.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ResponseStatus.Get()
}

// GetResponseStatusOk returns a tuple with the ResponseStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetResponseStatusOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseStatus.Get(), o.ResponseStatus.IsSet()
}

// SetResponseStatus sets field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetResponseStatus(v float32) {
	o.ResponseStatus.Set(&v)
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetError(v string) {
	o.Error.Set(&v)
}

// GetAttemptedAt returns the AttemptedAt field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetAttemptedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttemptedAt
}

// GetAttemptedAtOk returns a tuple with the AttemptedAt field value
// and a boolean to check if the value has been set.
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetAttemptedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttemptedAt, true
}

// SetAttemptedAt sets field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetAttemptedAt(v string) {
	o.AttemptedAt = v
}

// GetTest returns the Test field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Test
}

// GetTestOk returns a tuple with the Test field value
// and a boolean to check if the value has been set.
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Test, true
}

// SetTest sets field value
func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetTest(v bool) {
	o.Test = v
}

func (o ListWebhookDeliveries200ResponseDeliveriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhookDeliveries200ResponseDeliveriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["event"] = o.Event
	toSerialize["attempt"] = o.Attempt
	toSerialize["status"] = o.Status
	toSerialize["responseStatus"] = o.ResponseStatus.Get()
	toSerialize["error"] = o.Error.Get()
	toSerialize["attemptedAt"] = o.AttemptedAt
	toSerialize["test"] = o.Test

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListWebhookDeliveries200ResponseDeliveriesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"event",
		"attempt",
		"status",
		"responseStatus",
		"error",
		"attemptedAt",
		"test",
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

	varListWebhookDeliveries200ResponseDeliveriesInner := _ListWebhookDeliveries200ResponseDeliveriesInner{}

	err = json.Unmarshal(data, &varListWebhookDeliveries200ResponseDeliveriesInner)

	if err != nil {
		return err
	}

	*o = ListWebhookDeliveries200ResponseDeliveriesInner(varListWebhookDeliveries200ResponseDeliveriesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "event")
		delete(additionalProperties, "attempt")
		delete(additionalProperties, "status")
		delete(additionalProperties, "responseStatus")
		delete(additionalProperties, "error")
		delete(additionalProperties, "attemptedAt")
		delete(additionalProperties, "test")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListWebhookDeliveries200ResponseDeliveriesInner struct {
	value *ListWebhookDeliveries200ResponseDeliveriesInner
	isSet bool
}

func (v NullableListWebhookDeliveries200ResponseDeliveriesInner) Get() *ListWebhookDeliveries200ResponseDeliveriesInner {
	return v.value
}

func (v *NullableListWebhookDeliveries200ResponseDeliveriesInner) Set(val *ListWebhookDeliveries200ResponseDeliveriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhookDeliveries200ResponseDeliveriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhookDeliveries200ResponseDeliveriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhookDeliveries200ResponseDeliveriesInner(val *ListWebhookDeliveries200ResponseDeliveriesInner) *NullableListWebhookDeliveries200ResponseDeliveriesInner {
	return &NullableListWebhookDeliveries200ResponseDeliveriesInner{value: val, isSet: true}
}

func (v NullableListWebhookDeliveries200ResponseDeliveriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhookDeliveries200ResponseDeliveriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the FetchRelayEvents200ResponseEventsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchRelayEvents200ResponseEventsInner{}

// FetchRelayEvents200ResponseEventsInner struct for FetchRelayEvents200ResponseEventsInner
type FetchRelayEvents200ResponseEventsInner struct {
	Id string `json:"id"`
	Topic NullableString `json:"topic"`
	Headers map[string]string `json:"headers"`
	Body string `json:"body"`
	At string `json:"at"`
	AdditionalProperties map[string]interface{}
}

type _FetchRelayEvents200ResponseEventsInner FetchRelayEvents200ResponseEventsInner

// NewFetchRelayEvents200ResponseEventsInner instantiates a new FetchRelayEvents200ResponseEventsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchRelayEvents200ResponseEventsInner(id string, topic NullableString, headers map[string]string, body string, at string) *FetchRelayEvents200ResponseEventsInner {
	this := FetchRelayEvents200ResponseEventsInner{}
	this.Id = id
	this.Topic = topic
	this.Headers = headers
	this.Body = body
	this.At = at
	return &this
}

// NewFetchRelayEvents200ResponseEventsInnerWithDefaults instantiates a new FetchRelayEvents200ResponseEventsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchRelayEvents200ResponseEventsInnerWithDefaults() *FetchRelayEvents200ResponseEventsInner {
	this := FetchRelayEvents200ResponseEventsInner{}
	return &this
}

// GetId returns the Id field value
func (o *FetchRelayEvents200ResponseEventsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FetchRelayEvents200ResponseEventsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FetchRelayEvents200ResponseEventsInner) SetId(v string) {
	o.Id = v
}

// GetTopic returns the Topic field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FetchRelayEvents200ResponseEventsInner) GetTopic() string {
	if o == nil || o.Topic.Get() == nil {
		var ret string
		return ret
	}

	return *o.Topic.Get()
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchRelayEvents200ResponseEventsInner) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Topic.Get(), o.Topic.IsSet()
}

// SetTopic sets field value
func (o *FetchRelayEvents200ResponseEventsInner) SetTopic(v string) {
	o.Topic.Set(&v)
}

// GetHeaders returns the Headers field value
func (o *FetchRelayEvents200ResponseEventsInner) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *FetchRelayEvents200ResponseEventsInner) GetHeadersOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.Headers, true
}

// SetHeaders sets field value
func (o *FetchRelayEvents200ResponseEventsInner) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetBody returns the Body field value
func (o *FetchRelayEvents200ResponseEventsInner) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *FetchRelayEvents200ResponseEventsInner) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *FetchRelayEvents200ResponseEventsInner) SetBody(v string) {
	o.Body = v
}

// GetAt returns the At field value
func (o *FetchRelayEvents200ResponseEventsInner) GetAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.At
}

// GetAtOk returns a tuple with the At field value
// and a boolean to check if the value has been set.
func (o *FetchRelayEvents200ResponseEventsInner) GetAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.At, true
}

// SetAt sets field value
func (o *FetchRelayEvents200ResponseEventsInner) SetAt(v string) {
	o.At = v
}

func (o FetchRelayEvents200ResponseEventsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchRelayEvents200ResponseEventsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["topic"] = o.Topic.Get()
	toSerialize["headers"] = o.Headers
	toSerialize["body"] = o.Body
	toSerialize["at"] = o.At

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FetchRelayEvents200ResponseEventsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"topic",
		"headers",
		"body",
		"at",
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

	varFetchRelayEvents200ResponseEventsInner := _FetchRelayEvents200ResponseEventsInner{}

	err = json.Unmarshal(data, &varFetchRelayEvents200ResponseEventsInner)

	if err != nil {
		return err
	}

	*o = FetchRelayEvents200ResponseEventsInner(varFetchRelayEvents200ResponseEventsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "topic")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "body")
		delete(additionalProperties, "at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFetchRelayEvents200ResponseEventsInner struct {
	value *FetchRelayEvents200ResponseEventsInner
	isSet bool
}

func (v NullableFetchRelayEvents200ResponseEventsInner) Get() *FetchRelayEvents200ResponseEventsInner {
	return v.value
}

func (v *NullableFetchRelayEvents200ResponseEventsInner) Set(val *FetchRelayEvents200ResponseEventsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchRelayEvents200ResponseEventsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchRelayEvents200ResponseEventsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchRelayEvents200ResponseEventsInner(val *FetchRelayEvents200ResponseEventsInner) *NullableFetchRelayEvents200ResponseEventsInner {
	return &NullableFetchRelayEvents200ResponseEventsInner{value: val, isSet: true}
}

func (v NullableFetchRelayEvents200ResponseEventsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchRelayEvents200ResponseEventsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



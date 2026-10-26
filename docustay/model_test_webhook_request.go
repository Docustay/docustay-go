/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the TestWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestWebhookRequest{}

// TestWebhookRequest struct for TestWebhookRequest
type TestWebhookRequest struct {
	Topic *string `json:"topic,omitempty"`
}

// NewTestWebhookRequest instantiates a new TestWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestWebhookRequest() *TestWebhookRequest {
	this := TestWebhookRequest{}
	return &this
}

// NewTestWebhookRequestWithDefaults instantiates a new TestWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestWebhookRequestWithDefaults() *TestWebhookRequest {
	this := TestWebhookRequest{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *TestWebhookRequest) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestWebhookRequest) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *TestWebhookRequest) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *TestWebhookRequest) SetTopic(v string) {
	o.Topic = &v
}

func (o TestWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	return toSerialize, nil
}

type NullableTestWebhookRequest struct {
	value *TestWebhookRequest
	isSet bool
}

func (v NullableTestWebhookRequest) Get() *TestWebhookRequest {
	return v.value
}

func (v *NullableTestWebhookRequest) Set(val *TestWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestWebhookRequest(val *TestWebhookRequest) *NullableTestWebhookRequest {
	return &NullableTestWebhookRequest{value: val, isSet: true}
}

func (v NullableTestWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



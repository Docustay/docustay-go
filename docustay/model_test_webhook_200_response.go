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

// checks if the TestWebhook200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestWebhook200Response{}

// TestWebhook200Response struct for TestWebhook200Response
type TestWebhook200Response struct {
	Outcome string `json:"outcome"`
	Status NullableFloat32 `json:"status"`
	DurationMs float32 `json:"durationMs"`
	ResponseExcerpt NullableString `json:"responseExcerpt"`
	Error NullableString `json:"error"`
	Topic string `json:"topic"`
	AdditionalProperties map[string]interface{}
}

type _TestWebhook200Response TestWebhook200Response

// NewTestWebhook200Response instantiates a new TestWebhook200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestWebhook200Response(outcome string, status NullableFloat32, durationMs float32, responseExcerpt NullableString, error_ NullableString, topic string) *TestWebhook200Response {
	this := TestWebhook200Response{}
	this.Outcome = outcome
	this.Status = status
	this.DurationMs = durationMs
	this.ResponseExcerpt = responseExcerpt
	this.Error = error_
	this.Topic = topic
	return &this
}

// NewTestWebhook200ResponseWithDefaults instantiates a new TestWebhook200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestWebhook200ResponseWithDefaults() *TestWebhook200Response {
	this := TestWebhook200Response{}
	return &this
}

// GetOutcome returns the Outcome field value
func (o *TestWebhook200Response) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *TestWebhook200Response) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *TestWebhook200Response) SetOutcome(v string) {
	o.Outcome = v
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *TestWebhook200Response) GetStatus() float32 {
	if o == nil || o.Status.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestWebhook200Response) GetStatusOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *TestWebhook200Response) SetStatus(v float32) {
	o.Status.Set(&v)
}

// GetDurationMs returns the DurationMs field value
func (o *TestWebhook200Response) GetDurationMs() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value
// and a boolean to check if the value has been set.
func (o *TestWebhook200Response) GetDurationMsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMs, true
}

// SetDurationMs sets field value
func (o *TestWebhook200Response) SetDurationMs(v float32) {
	o.DurationMs = v
}

// GetResponseExcerpt returns the ResponseExcerpt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TestWebhook200Response) GetResponseExcerpt() string {
	if o == nil || o.ResponseExcerpt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResponseExcerpt.Get()
}

// GetResponseExcerptOk returns a tuple with the ResponseExcerpt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestWebhook200Response) GetResponseExcerptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseExcerpt.Get(), o.ResponseExcerpt.IsSet()
}

// SetResponseExcerpt sets field value
func (o *TestWebhook200Response) SetResponseExcerpt(v string) {
	o.ResponseExcerpt.Set(&v)
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TestWebhook200Response) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestWebhook200Response) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *TestWebhook200Response) SetError(v string) {
	o.Error.Set(&v)
}

// GetTopic returns the Topic field value
func (o *TestWebhook200Response) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *TestWebhook200Response) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *TestWebhook200Response) SetTopic(v string) {
	o.Topic = v
}

func (o TestWebhook200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestWebhook200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["outcome"] = o.Outcome
	toSerialize["status"] = o.Status.Get()
	toSerialize["durationMs"] = o.DurationMs
	toSerialize["responseExcerpt"] = o.ResponseExcerpt.Get()
	toSerialize["error"] = o.Error.Get()
	toSerialize["topic"] = o.Topic

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestWebhook200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"outcome",
		"status",
		"durationMs",
		"responseExcerpt",
		"error",
		"topic",
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

	varTestWebhook200Response := _TestWebhook200Response{}

	err = json.Unmarshal(data, &varTestWebhook200Response)

	if err != nil {
		return err
	}

	*o = TestWebhook200Response(varTestWebhook200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "status")
		delete(additionalProperties, "durationMs")
		delete(additionalProperties, "responseExcerpt")
		delete(additionalProperties, "error")
		delete(additionalProperties, "topic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestWebhook200Response struct {
	value *TestWebhook200Response
	isSet bool
}

func (v NullableTestWebhook200Response) Get() *TestWebhook200Response {
	return v.value
}

func (v *NullableTestWebhook200Response) Set(val *TestWebhook200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTestWebhook200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTestWebhook200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestWebhook200Response(val *TestWebhook200Response) *NullableTestWebhook200Response {
	return &NullableTestWebhook200Response{value: val, isSet: true}
}

func (v NullableTestWebhook200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestWebhook200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



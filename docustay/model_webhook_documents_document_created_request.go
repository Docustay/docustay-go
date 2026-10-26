/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the WebhookDocumentsDocumentCreatedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookDocumentsDocumentCreatedRequest{}

// WebhookDocumentsDocumentCreatedRequest struct for WebhookDocumentsDocumentCreatedRequest
type WebhookDocumentsDocumentCreatedRequest struct {
	// The event's id; the same value as the webhook-id header. Time-ordered, and unchanged when a delivery is retried.
	Id string `json:"id"`
	// What happened.
	Topic string `json:"topic"`
	SubjectType string `json:"subject_type"`
	SubjectId string `json:"subject_id"`
	// Facts about the subject. Never a signing link or a code. `payload.test` is true for every event of a test-mode document.
	Payload map[string]interface{} `json:"payload"`
	EmittedAt time.Time `json:"emitted_at"`
}

type _WebhookDocumentsDocumentCreatedRequest WebhookDocumentsDocumentCreatedRequest

// NewWebhookDocumentsDocumentCreatedRequest instantiates a new WebhookDocumentsDocumentCreatedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDocumentsDocumentCreatedRequest(id string, topic string, subjectType string, subjectId string, payload map[string]interface{}, emittedAt time.Time) *WebhookDocumentsDocumentCreatedRequest {
	this := WebhookDocumentsDocumentCreatedRequest{}
	this.Id = id
	this.Topic = topic
	this.SubjectType = subjectType
	this.SubjectId = subjectId
	this.Payload = payload
	this.EmittedAt = emittedAt
	return &this
}

// NewWebhookDocumentsDocumentCreatedRequestWithDefaults instantiates a new WebhookDocumentsDocumentCreatedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDocumentsDocumentCreatedRequestWithDefaults() *WebhookDocumentsDocumentCreatedRequest {
	this := WebhookDocumentsDocumentCreatedRequest{}
	return &this
}

// GetId returns the Id field value
func (o *WebhookDocumentsDocumentCreatedRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookDocumentsDocumentCreatedRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookDocumentsDocumentCreatedRequest) SetId(v string) {
	o.Id = v
}

// GetTopic returns the Topic field value
func (o *WebhookDocumentsDocumentCreatedRequest) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *WebhookDocumentsDocumentCreatedRequest) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *WebhookDocumentsDocumentCreatedRequest) SetTopic(v string) {
	o.Topic = v
}

// GetSubjectType returns the SubjectType field value
func (o *WebhookDocumentsDocumentCreatedRequest) GetSubjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *WebhookDocumentsDocumentCreatedRequest) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value
func (o *WebhookDocumentsDocumentCreatedRequest) SetSubjectType(v string) {
	o.SubjectType = v
}

// GetSubjectId returns the SubjectId field value
func (o *WebhookDocumentsDocumentCreatedRequest) GetSubjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value
// and a boolean to check if the value has been set.
func (o *WebhookDocumentsDocumentCreatedRequest) GetSubjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectId, true
}

// SetSubjectId sets field value
func (o *WebhookDocumentsDocumentCreatedRequest) SetSubjectId(v string) {
	o.SubjectId = v
}

// GetPayload returns the Payload field value
func (o *WebhookDocumentsDocumentCreatedRequest) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *WebhookDocumentsDocumentCreatedRequest) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *WebhookDocumentsDocumentCreatedRequest) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetEmittedAt returns the EmittedAt field value
func (o *WebhookDocumentsDocumentCreatedRequest) GetEmittedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EmittedAt
}

// GetEmittedAtOk returns a tuple with the EmittedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookDocumentsDocumentCreatedRequest) GetEmittedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmittedAt, true
}

// SetEmittedAt sets field value
func (o *WebhookDocumentsDocumentCreatedRequest) SetEmittedAt(v time.Time) {
	o.EmittedAt = v
}

func (o WebhookDocumentsDocumentCreatedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookDocumentsDocumentCreatedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["topic"] = o.Topic
	toSerialize["subject_type"] = o.SubjectType
	toSerialize["subject_id"] = o.SubjectId
	toSerialize["payload"] = o.Payload
	toSerialize["emitted_at"] = o.EmittedAt
	return toSerialize, nil
}

func (o *WebhookDocumentsDocumentCreatedRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"topic",
		"subject_type",
		"subject_id",
		"payload",
		"emitted_at",
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

	varWebhookDocumentsDocumentCreatedRequest := _WebhookDocumentsDocumentCreatedRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookDocumentsDocumentCreatedRequest)

	if err != nil {
		return err
	}

	*o = WebhookDocumentsDocumentCreatedRequest(varWebhookDocumentsDocumentCreatedRequest)

	return err
}

type NullableWebhookDocumentsDocumentCreatedRequest struct {
	value *WebhookDocumentsDocumentCreatedRequest
	isSet bool
}

func (v NullableWebhookDocumentsDocumentCreatedRequest) Get() *WebhookDocumentsDocumentCreatedRequest {
	return v.value
}

func (v *NullableWebhookDocumentsDocumentCreatedRequest) Set(val *WebhookDocumentsDocumentCreatedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDocumentsDocumentCreatedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDocumentsDocumentCreatedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDocumentsDocumentCreatedRequest(val *WebhookDocumentsDocumentCreatedRequest) *NullableWebhookDocumentsDocumentCreatedRequest {
	return &NullableWebhookDocumentsDocumentCreatedRequest{value: val, isSet: true}
}

func (v NullableWebhookDocumentsDocumentCreatedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDocumentsDocumentCreatedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



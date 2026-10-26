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

// checks if the DocumentSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentSummary{}

// DocumentSummary struct for DocumentSummary
type DocumentSummary struct {
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId NullableString `json:"externalId"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata"`
	// True for a document made in test mode (a test API key or the workspace's test switch).
	Test bool `json:"test"`
	Id string `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
	Parties int32 `json:"parties"`
	CreatedAt NullableString `json:"createdAt"`
	SentAt NullableString `json:"sentAt"`
	ExecutedAt NullableString `json:"executedAt"`
	ExpiresAt NullableString `json:"expiresAt"`
	AdditionalProperties map[string]interface{}
}

type _DocumentSummary DocumentSummary

// NewDocumentSummary instantiates a new DocumentSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentSummary(externalId NullableString, metadata map[string]string, test bool, id string, title string, state string, parties int32, createdAt NullableString, sentAt NullableString, executedAt NullableString, expiresAt NullableString) *DocumentSummary {
	this := DocumentSummary{}
	this.ExternalId = externalId
	this.Metadata = metadata
	this.Test = test
	this.Id = id
	this.Title = title
	this.State = state
	this.Parties = parties
	this.CreatedAt = createdAt
	this.SentAt = sentAt
	this.ExecutedAt = executedAt
	this.ExpiresAt = expiresAt
	return &this
}

// NewDocumentSummaryWithDefaults instantiates a new DocumentSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentSummaryWithDefaults() *DocumentSummary {
	this := DocumentSummary{}
	return &this
}

// GetExternalId returns the ExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentSummary) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSummary) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// SetExternalId sets field value
func (o *DocumentSummary) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// GetMetadata returns the Metadata field value
func (o *DocumentSummary) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *DocumentSummary) GetMetadataOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *DocumentSummary) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetTest returns the Test field value
func (o *DocumentSummary) GetTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Test
}

// GetTestOk returns a tuple with the Test field value
// and a boolean to check if the value has been set.
func (o *DocumentSummary) GetTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Test, true
}

// SetTest sets field value
func (o *DocumentSummary) SetTest(v bool) {
	o.Test = v
}

// GetId returns the Id field value
func (o *DocumentSummary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DocumentSummary) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DocumentSummary) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *DocumentSummary) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DocumentSummary) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DocumentSummary) SetTitle(v string) {
	o.Title = v
}

// GetState returns the State field value
func (o *DocumentSummary) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DocumentSummary) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DocumentSummary) SetState(v string) {
	o.State = v
}

// GetParties returns the Parties field value
func (o *DocumentSummary) GetParties() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Parties
}

// GetPartiesOk returns a tuple with the Parties field value
// and a boolean to check if the value has been set.
func (o *DocumentSummary) GetPartiesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parties, true
}

// SetParties sets field value
func (o *DocumentSummary) SetParties(v int32) {
	o.Parties = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentSummary) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSummary) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *DocumentSummary) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetSentAt returns the SentAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentSummary) GetSentAt() string {
	if o == nil || o.SentAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.SentAt.Get()
}

// GetSentAtOk returns a tuple with the SentAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSummary) GetSentAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SentAt.Get(), o.SentAt.IsSet()
}

// SetSentAt sets field value
func (o *DocumentSummary) SetSentAt(v string) {
	o.SentAt.Set(&v)
}

// GetExecutedAt returns the ExecutedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentSummary) GetExecutedAt() string {
	if o == nil || o.ExecutedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExecutedAt.Get()
}

// GetExecutedAtOk returns a tuple with the ExecutedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSummary) GetExecutedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutedAt.Get(), o.ExecutedAt.IsSet()
}

// SetExecutedAt sets field value
func (o *DocumentSummary) SetExecutedAt(v string) {
	o.ExecutedAt.Set(&v)
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentSummary) GetExpiresAt() string {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSummary) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *DocumentSummary) SetExpiresAt(v string) {
	o.ExpiresAt.Set(&v)
}

func (o DocumentSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId.Get()
	toSerialize["metadata"] = o.Metadata
	toSerialize["test"] = o.Test
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["state"] = o.State
	toSerialize["parties"] = o.Parties
	toSerialize["createdAt"] = o.CreatedAt.Get()
	toSerialize["sentAt"] = o.SentAt.Get()
	toSerialize["executedAt"] = o.ExecutedAt.Get()
	toSerialize["expiresAt"] = o.ExpiresAt.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DocumentSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalId",
		"metadata",
		"test",
		"id",
		"title",
		"state",
		"parties",
		"createdAt",
		"sentAt",
		"executedAt",
		"expiresAt",
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

	varDocumentSummary := _DocumentSummary{}

	err = json.Unmarshal(data, &varDocumentSummary)

	if err != nil {
		return err
	}

	*o = DocumentSummary(varDocumentSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "test")
		delete(additionalProperties, "id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "state")
		delete(additionalProperties, "parties")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "sentAt")
		delete(additionalProperties, "executedAt")
		delete(additionalProperties, "expiresAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentSummary struct {
	value *DocumentSummary
	isSet bool
}

func (v NullableDocumentSummary) Get() *DocumentSummary {
	return v.value
}

func (v *NullableDocumentSummary) Set(val *DocumentSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentSummary(val *DocumentSummary) *NullableDocumentSummary {
	return &NullableDocumentSummary{value: val, isSet: true}
}

func (v NullableDocumentSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



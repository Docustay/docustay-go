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

// checks if the Document type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Document{}

// Document struct for Document
type Document struct {
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId NullableString `json:"externalId"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata"`
	Test bool `json:"test"`
	Id string `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
	SentAt NullableString `json:"sentAt"`
	ExpiresAt NullableString `json:"expiresAt"`
	ExecutedAt NullableString `json:"executedAt"`
	VoidedAt NullableString `json:"voidedAt"`
	VoidReason NullableString `json:"voidReason"`
	Progress TestSign200ResponseProgress `json:"progress"`
	SignedCopyAvailable bool `json:"signedCopyAvailable"`
	Parties []Party `json:"parties"`
	// Payments made through payment fields on this document (empty when it has none). A payment field's amount is one set amount (optionally paid as a deposit now), or an amount you give when you send the document — see the text tag setting `amount`.
	Payments []Payment `json:"payments"`
	Events []DocumentEventsInner `json:"events"`
	AdditionalProperties map[string]interface{}
}

type _Document Document

// NewDocument instantiates a new Document object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocument(externalId NullableString, metadata map[string]string, test bool, id string, title string, state string, sentAt NullableString, expiresAt NullableString, executedAt NullableString, voidedAt NullableString, voidReason NullableString, progress TestSign200ResponseProgress, signedCopyAvailable bool, parties []Party, payments []Payment, events []DocumentEventsInner) *Document {
	this := Document{}
	this.ExternalId = externalId
	this.Metadata = metadata
	this.Test = test
	this.Id = id
	this.Title = title
	this.State = state
	this.SentAt = sentAt
	this.ExpiresAt = expiresAt
	this.ExecutedAt = executedAt
	this.VoidedAt = voidedAt
	this.VoidReason = voidReason
	this.Progress = progress
	this.SignedCopyAvailable = signedCopyAvailable
	this.Parties = parties
	this.Payments = payments
	this.Events = events
	return &this
}

// NewDocumentWithDefaults instantiates a new Document object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentWithDefaults() *Document {
	this := Document{}
	return &this
}

// GetExternalId returns the ExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Document) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Document) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// SetExternalId sets field value
func (o *Document) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// GetMetadata returns the Metadata field value
func (o *Document) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Document) GetMetadataOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Document) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetTest returns the Test field value
func (o *Document) GetTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Test
}

// GetTestOk returns a tuple with the Test field value
// and a boolean to check if the value has been set.
func (o *Document) GetTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Test, true
}

// SetTest sets field value
func (o *Document) SetTest(v bool) {
	o.Test = v
}

// GetId returns the Id field value
func (o *Document) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Document) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Document) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *Document) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Document) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Document) SetTitle(v string) {
	o.Title = v
}

// GetState returns the State field value
func (o *Document) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Document) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Document) SetState(v string) {
	o.State = v
}

// GetSentAt returns the SentAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Document) GetSentAt() string {
	if o == nil || o.SentAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.SentAt.Get()
}

// GetSentAtOk returns a tuple with the SentAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Document) GetSentAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SentAt.Get(), o.SentAt.IsSet()
}

// SetSentAt sets field value
func (o *Document) SetSentAt(v string) {
	o.SentAt.Set(&v)
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Document) GetExpiresAt() string {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Document) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *Document) SetExpiresAt(v string) {
	o.ExpiresAt.Set(&v)
}

// GetExecutedAt returns the ExecutedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Document) GetExecutedAt() string {
	if o == nil || o.ExecutedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExecutedAt.Get()
}

// GetExecutedAtOk returns a tuple with the ExecutedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Document) GetExecutedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutedAt.Get(), o.ExecutedAt.IsSet()
}

// SetExecutedAt sets field value
func (o *Document) SetExecutedAt(v string) {
	o.ExecutedAt.Set(&v)
}

// GetVoidedAt returns the VoidedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Document) GetVoidedAt() string {
	if o == nil || o.VoidedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.VoidedAt.Get()
}

// GetVoidedAtOk returns a tuple with the VoidedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Document) GetVoidedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoidedAt.Get(), o.VoidedAt.IsSet()
}

// SetVoidedAt sets field value
func (o *Document) SetVoidedAt(v string) {
	o.VoidedAt.Set(&v)
}

// GetVoidReason returns the VoidReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Document) GetVoidReason() string {
	if o == nil || o.VoidReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.VoidReason.Get()
}

// GetVoidReasonOk returns a tuple with the VoidReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Document) GetVoidReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoidReason.Get(), o.VoidReason.IsSet()
}

// SetVoidReason sets field value
func (o *Document) SetVoidReason(v string) {
	o.VoidReason.Set(&v)
}

// GetProgress returns the Progress field value
func (o *Document) GetProgress() TestSign200ResponseProgress {
	if o == nil {
		var ret TestSign200ResponseProgress
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *Document) GetProgressOk() (*TestSign200ResponseProgress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *Document) SetProgress(v TestSign200ResponseProgress) {
	o.Progress = v
}

// GetSignedCopyAvailable returns the SignedCopyAvailable field value
func (o *Document) GetSignedCopyAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SignedCopyAvailable
}

// GetSignedCopyAvailableOk returns a tuple with the SignedCopyAvailable field value
// and a boolean to check if the value has been set.
func (o *Document) GetSignedCopyAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedCopyAvailable, true
}

// SetSignedCopyAvailable sets field value
func (o *Document) SetSignedCopyAvailable(v bool) {
	o.SignedCopyAvailable = v
}

// GetParties returns the Parties field value
func (o *Document) GetParties() []Party {
	if o == nil {
		var ret []Party
		return ret
	}

	return o.Parties
}

// GetPartiesOk returns a tuple with the Parties field value
// and a boolean to check if the value has been set.
func (o *Document) GetPartiesOk() ([]Party, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parties, true
}

// SetParties sets field value
func (o *Document) SetParties(v []Party) {
	o.Parties = v
}

// GetPayments returns the Payments field value
func (o *Document) GetPayments() []Payment {
	if o == nil {
		var ret []Payment
		return ret
	}

	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value
// and a boolean to check if the value has been set.
func (o *Document) GetPaymentsOk() ([]Payment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payments, true
}

// SetPayments sets field value
func (o *Document) SetPayments(v []Payment) {
	o.Payments = v
}

// GetEvents returns the Events field value
func (o *Document) GetEvents() []DocumentEventsInner {
	if o == nil {
		var ret []DocumentEventsInner
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *Document) GetEventsOk() ([]DocumentEventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *Document) SetEvents(v []DocumentEventsInner) {
	o.Events = v
}

func (o Document) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Document) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId.Get()
	toSerialize["metadata"] = o.Metadata
	toSerialize["test"] = o.Test
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["state"] = o.State
	toSerialize["sentAt"] = o.SentAt.Get()
	toSerialize["expiresAt"] = o.ExpiresAt.Get()
	toSerialize["executedAt"] = o.ExecutedAt.Get()
	toSerialize["voidedAt"] = o.VoidedAt.Get()
	toSerialize["voidReason"] = o.VoidReason.Get()
	toSerialize["progress"] = o.Progress
	toSerialize["signedCopyAvailable"] = o.SignedCopyAvailable
	toSerialize["parties"] = o.Parties
	toSerialize["payments"] = o.Payments
	toSerialize["events"] = o.Events

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Document) UnmarshalJSON(data []byte) (err error) {
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
		"sentAt",
		"expiresAt",
		"executedAt",
		"voidedAt",
		"voidReason",
		"progress",
		"signedCopyAvailable",
		"parties",
		"payments",
		"events",
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

	varDocument := _Document{}

	err = json.Unmarshal(data, &varDocument)

	if err != nil {
		return err
	}

	*o = Document(varDocument)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "test")
		delete(additionalProperties, "id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "state")
		delete(additionalProperties, "sentAt")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "executedAt")
		delete(additionalProperties, "voidedAt")
		delete(additionalProperties, "voidReason")
		delete(additionalProperties, "progress")
		delete(additionalProperties, "signedCopyAvailable")
		delete(additionalProperties, "parties")
		delete(additionalProperties, "payments")
		delete(additionalProperties, "events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocument struct {
	value *Document
	isSet bool
}

func (v NullableDocument) Get() *Document {
	return v.value
}

func (v *NullableDocument) Set(val *Document) {
	v.value = val
	v.isSet = true
}

func (v NullableDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocument(val *Document) *NullableDocument {
	return &NullableDocument{value: val, isSet: true}
}

func (v NullableDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



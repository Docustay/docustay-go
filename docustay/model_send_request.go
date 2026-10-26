/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SendRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendRequest{}

// SendRequest struct for SendRequest
type SendRequest struct {
	TemplateId string `json:"templateId"`
	Email *EmailChoice `json:"email,omitempty"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId *string `json:"externalId,omitempty"`
	Title *string `json:"title,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Subject *string `json:"subject,omitempty"`
	ExpiryDays *int32 `json:"expiryDays,omitempty"`
	// `sequential`: each seat signs in the template's seat order (signer_1 first), and an embedded seat gets its session only when it is that person's turn. Left out, the template's own setting.
	SigningOrder *string `json:"signingOrder,omitempty"`
	// The boxes the sender fills in (\"Filled in by: Sender\") and the template's send-time amounts, by key: a box's label in lower case with underscores (\"Effective date\" → `effective_date`). A send that leaves one empty answers 400 naming it.
	Values map[string]string `json:"values,omitempty"`
	// Seats signed inside your own app instead of by emailed link (for example [\"signer_1\"]). Nobody is emailed for these seats: you ask for a session with POST /api/v1/documents/{id}/embed-session and show it on a website you listed under Settings → Developers → Embedding.
	Embedded []string `json:"embedded,omitempty"`
	// One entry for each seat of the template (see GET /api/v1/documents/templates).
	Parties []SendRequestPartiesInner `json:"parties"`
}

type _SendRequest SendRequest

// NewSendRequest instantiates a new SendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendRequest(templateId string, parties []SendRequestPartiesInner) *SendRequest {
	this := SendRequest{}
	this.TemplateId = templateId
	this.Parties = parties
	return &this
}

// NewSendRequestWithDefaults instantiates a new SendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendRequestWithDefaults() *SendRequest {
	this := SendRequest{}
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *SendRequest) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *SendRequest) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *SendRequest) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SendRequest) GetEmail() EmailChoice {
	if o == nil || IsNil(o.Email) {
		var ret EmailChoice
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequest) GetEmailOk() (*EmailChoice, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SendRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailChoice and assigns it to the Email field.
func (o *SendRequest) SetEmail(v EmailChoice) {
	o.Email = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SendRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequest) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SendRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *SendRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *SendRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *SendRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *SendRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SendRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SendRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SendRequest) SetTitle(v string) {
	o.Title = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendRequest) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *SendRequest) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *SendRequest) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *SendRequest) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *SendRequest) UnsetMessage() {
	o.Message.Unset()
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SendRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SendRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SendRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetExpiryDays returns the ExpiryDays field value if set, zero value otherwise.
func (o *SendRequest) GetExpiryDays() int32 {
	if o == nil || IsNil(o.ExpiryDays) {
		var ret int32
		return ret
	}
	return *o.ExpiryDays
}

// GetExpiryDaysOk returns a tuple with the ExpiryDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequest) GetExpiryDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiryDays) {
		return nil, false
	}
	return o.ExpiryDays, true
}

// HasExpiryDays returns a boolean if a field has been set.
func (o *SendRequest) HasExpiryDays() bool {
	if o != nil && !IsNil(o.ExpiryDays) {
		return true
	}

	return false
}

// SetExpiryDays gets a reference to the given int32 and assigns it to the ExpiryDays field.
func (o *SendRequest) SetExpiryDays(v int32) {
	o.ExpiryDays = &v
}

// GetSigningOrder returns the SigningOrder field value if set, zero value otherwise.
func (o *SendRequest) GetSigningOrder() string {
	if o == nil || IsNil(o.SigningOrder) {
		var ret string
		return ret
	}
	return *o.SigningOrder
}

// GetSigningOrderOk returns a tuple with the SigningOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequest) GetSigningOrderOk() (*string, bool) {
	if o == nil || IsNil(o.SigningOrder) {
		return nil, false
	}
	return o.SigningOrder, true
}

// HasSigningOrder returns a boolean if a field has been set.
func (o *SendRequest) HasSigningOrder() bool {
	if o != nil && !IsNil(o.SigningOrder) {
		return true
	}

	return false
}

// SetSigningOrder gets a reference to the given string and assigns it to the SigningOrder field.
func (o *SendRequest) SetSigningOrder(v string) {
	o.SigningOrder = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SendRequest) GetValues() map[string]string {
	if o == nil || IsNil(o.Values) {
		var ret map[string]string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequest) GetValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Values) {
		return map[string]string{}, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SendRequest) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]string and assigns it to the Values field.
func (o *SendRequest) SetValues(v map[string]string) {
	o.Values = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *SendRequest) GetEmbedded() []string {
	if o == nil || IsNil(o.Embedded) {
		var ret []string
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequest) GetEmbeddedOk() ([]string, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *SendRequest) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given []string and assigns it to the Embedded field.
func (o *SendRequest) SetEmbedded(v []string) {
	o.Embedded = v
}

// GetParties returns the Parties field value
func (o *SendRequest) GetParties() []SendRequestPartiesInner {
	if o == nil {
		var ret []SendRequestPartiesInner
		return ret
	}

	return o.Parties
}

// GetPartiesOk returns a tuple with the Parties field value
// and a boolean to check if the value has been set.
func (o *SendRequest) GetPartiesOk() ([]SendRequestPartiesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parties, true
}

// SetParties sets field value
func (o *SendRequest) SetParties(v []SendRequestPartiesInner) {
	o.Parties = v
}

func (o SendRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateId"] = o.TemplateId
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.ExpiryDays) {
		toSerialize["expiryDays"] = o.ExpiryDays
	}
	if !IsNil(o.SigningOrder) {
		toSerialize["signingOrder"] = o.SigningOrder
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.Embedded) {
		toSerialize["embedded"] = o.Embedded
	}
	toSerialize["parties"] = o.Parties
	return toSerialize, nil
}

func (o *SendRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateId",
		"parties",
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

	varSendRequest := _SendRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendRequest)

	if err != nil {
		return err
	}

	*o = SendRequest(varSendRequest)

	return err
}

type NullableSendRequest struct {
	value *SendRequest
	isSet bool
}

func (v NullableSendRequest) Get() *SendRequest {
	return v.value
}

func (v *NullableSendRequest) Set(val *SendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendRequest(val *SendRequest) *NullableSendRequest {
	return &NullableSendRequest{value: val, isSet: true}
}

func (v NullableSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



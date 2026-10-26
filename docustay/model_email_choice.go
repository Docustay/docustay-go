/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the EmailChoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailChoice{}

// EmailChoice Your own wording for the emails about THIS document. What you give here is used for this document.
type EmailChoice struct {
	Subject *string `json:"subject,omitempty"`
	Message *string `json:"message,omitempty"`
	// Where replies to this document's emails go.
	ReplyTo *string `json:"replyTo,omitempty"`
	// The name the emails say they are from, for this document.
	SenderName *string `json:"senderName,omitempty"`
}

// NewEmailChoice instantiates a new EmailChoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailChoice() *EmailChoice {
	this := EmailChoice{}
	return &this
}

// NewEmailChoiceWithDefaults instantiates a new EmailChoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailChoiceWithDefaults() *EmailChoice {
	this := EmailChoice{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailChoice) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChoice) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailChoice) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *EmailChoice) SetSubject(v string) {
	o.Subject = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EmailChoice) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChoice) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EmailChoice) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EmailChoice) SetMessage(v string) {
	o.Message = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *EmailChoice) GetReplyTo() string {
	if o == nil || IsNil(o.ReplyTo) {
		var ret string
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChoice) GetReplyToOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *EmailChoice) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given string and assigns it to the ReplyTo field.
func (o *EmailChoice) SetReplyTo(v string) {
	o.ReplyTo = &v
}

// GetSenderName returns the SenderName field value if set, zero value otherwise.
func (o *EmailChoice) GetSenderName() string {
	if o == nil || IsNil(o.SenderName) {
		var ret string
		return ret
	}
	return *o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChoice) GetSenderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SenderName) {
		return nil, false
	}
	return o.SenderName, true
}

// HasSenderName returns a boolean if a field has been set.
func (o *EmailChoice) HasSenderName() bool {
	if o != nil && !IsNil(o.SenderName) {
		return true
	}

	return false
}

// SetSenderName gets a reference to the given string and assigns it to the SenderName field.
func (o *EmailChoice) SetSenderName(v string) {
	o.SenderName = &v
}

func (o EmailChoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailChoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.SenderName) {
		toSerialize["senderName"] = o.SenderName
	}
	return toSerialize, nil
}

type NullableEmailChoice struct {
	value *EmailChoice
	isSet bool
}

func (v NullableEmailChoice) Get() *EmailChoice {
	return v.value
}

func (v *NullableEmailChoice) Set(val *EmailChoice) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailChoice) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailChoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailChoice(val *EmailChoice) *NullableEmailChoice {
	return &NullableEmailChoice{value: val, isSet: true}
}

func (v NullableEmailChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailChoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



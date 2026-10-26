/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the SendChoices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendChoices{}

// SendChoices struct for SendChoices
type SendChoices struct {
	ReplyTo *string `json:"replyTo,omitempty"`
	SenderName *string `json:"senderName,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Message *string `json:"message,omitempty"`
	ExpiryDays NullableInt32 `json:"expiryDays,omitempty"`
	ReminderDays []int32 `json:"reminderDays,omitempty"`
	SigningOrder *string `json:"signingOrder,omitempty"`
	// A note only the sender sees; never in anything a signer gets.
	InternalNote *string `json:"internalNote,omitempty"`
	RequireDeclineReason *bool `json:"requireDeclineReason,omitempty"`
}

// NewSendChoices instantiates a new SendChoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendChoices() *SendChoices {
	this := SendChoices{}
	return &this
}

// NewSendChoicesWithDefaults instantiates a new SendChoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendChoicesWithDefaults() *SendChoices {
	this := SendChoices{}
	return &this
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *SendChoices) GetReplyTo() string {
	if o == nil || IsNil(o.ReplyTo) {
		var ret string
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChoices) GetReplyToOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *SendChoices) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given string and assigns it to the ReplyTo field.
func (o *SendChoices) SetReplyTo(v string) {
	o.ReplyTo = &v
}

// GetSenderName returns the SenderName field value if set, zero value otherwise.
func (o *SendChoices) GetSenderName() string {
	if o == nil || IsNil(o.SenderName) {
		var ret string
		return ret
	}
	return *o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChoices) GetSenderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SenderName) {
		return nil, false
	}
	return o.SenderName, true
}

// HasSenderName returns a boolean if a field has been set.
func (o *SendChoices) HasSenderName() bool {
	if o != nil && !IsNil(o.SenderName) {
		return true
	}

	return false
}

// SetSenderName gets a reference to the given string and assigns it to the SenderName field.
func (o *SendChoices) SetSenderName(v string) {
	o.SenderName = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SendChoices) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChoices) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SendChoices) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SendChoices) SetSubject(v string) {
	o.Subject = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SendChoices) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChoices) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SendChoices) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SendChoices) SetMessage(v string) {
	o.Message = &v
}

// GetExpiryDays returns the ExpiryDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendChoices) GetExpiryDays() int32 {
	if o == nil || IsNil(o.ExpiryDays.Get()) {
		var ret int32
		return ret
	}
	return *o.ExpiryDays.Get()
}

// GetExpiryDaysOk returns a tuple with the ExpiryDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendChoices) GetExpiryDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryDays.Get(), o.ExpiryDays.IsSet()
}

// HasExpiryDays returns a boolean if a field has been set.
func (o *SendChoices) HasExpiryDays() bool {
	if o != nil && o.ExpiryDays.IsSet() {
		return true
	}

	return false
}

// SetExpiryDays gets a reference to the given NullableInt32 and assigns it to the ExpiryDays field.
func (o *SendChoices) SetExpiryDays(v int32) {
	o.ExpiryDays.Set(&v)
}
// SetExpiryDaysNil sets the value for ExpiryDays to be an explicit nil
func (o *SendChoices) SetExpiryDaysNil() {
	o.ExpiryDays.Set(nil)
}

// UnsetExpiryDays ensures that no value is present for ExpiryDays, not even an explicit nil
func (o *SendChoices) UnsetExpiryDays() {
	o.ExpiryDays.Unset()
}

// GetReminderDays returns the ReminderDays field value if set, zero value otherwise.
func (o *SendChoices) GetReminderDays() []int32 {
	if o == nil || IsNil(o.ReminderDays) {
		var ret []int32
		return ret
	}
	return o.ReminderDays
}

// GetReminderDaysOk returns a tuple with the ReminderDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChoices) GetReminderDaysOk() ([]int32, bool) {
	if o == nil || IsNil(o.ReminderDays) {
		return nil, false
	}
	return o.ReminderDays, true
}

// HasReminderDays returns a boolean if a field has been set.
func (o *SendChoices) HasReminderDays() bool {
	if o != nil && !IsNil(o.ReminderDays) {
		return true
	}

	return false
}

// SetReminderDays gets a reference to the given []int32 and assigns it to the ReminderDays field.
func (o *SendChoices) SetReminderDays(v []int32) {
	o.ReminderDays = v
}

// GetSigningOrder returns the SigningOrder field value if set, zero value otherwise.
func (o *SendChoices) GetSigningOrder() string {
	if o == nil || IsNil(o.SigningOrder) {
		var ret string
		return ret
	}
	return *o.SigningOrder
}

// GetSigningOrderOk returns a tuple with the SigningOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChoices) GetSigningOrderOk() (*string, bool) {
	if o == nil || IsNil(o.SigningOrder) {
		return nil, false
	}
	return o.SigningOrder, true
}

// HasSigningOrder returns a boolean if a field has been set.
func (o *SendChoices) HasSigningOrder() bool {
	if o != nil && !IsNil(o.SigningOrder) {
		return true
	}

	return false
}

// SetSigningOrder gets a reference to the given string and assigns it to the SigningOrder field.
func (o *SendChoices) SetSigningOrder(v string) {
	o.SigningOrder = &v
}

// GetInternalNote returns the InternalNote field value if set, zero value otherwise.
func (o *SendChoices) GetInternalNote() string {
	if o == nil || IsNil(o.InternalNote) {
		var ret string
		return ret
	}
	return *o.InternalNote
}

// GetInternalNoteOk returns a tuple with the InternalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChoices) GetInternalNoteOk() (*string, bool) {
	if o == nil || IsNil(o.InternalNote) {
		return nil, false
	}
	return o.InternalNote, true
}

// HasInternalNote returns a boolean if a field has been set.
func (o *SendChoices) HasInternalNote() bool {
	if o != nil && !IsNil(o.InternalNote) {
		return true
	}

	return false
}

// SetInternalNote gets a reference to the given string and assigns it to the InternalNote field.
func (o *SendChoices) SetInternalNote(v string) {
	o.InternalNote = &v
}

// GetRequireDeclineReason returns the RequireDeclineReason field value if set, zero value otherwise.
func (o *SendChoices) GetRequireDeclineReason() bool {
	if o == nil || IsNil(o.RequireDeclineReason) {
		var ret bool
		return ret
	}
	return *o.RequireDeclineReason
}

// GetRequireDeclineReasonOk returns a tuple with the RequireDeclineReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChoices) GetRequireDeclineReasonOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireDeclineReason) {
		return nil, false
	}
	return o.RequireDeclineReason, true
}

// HasRequireDeclineReason returns a boolean if a field has been set.
func (o *SendChoices) HasRequireDeclineReason() bool {
	if o != nil && !IsNil(o.RequireDeclineReason) {
		return true
	}

	return false
}

// SetRequireDeclineReason gets a reference to the given bool and assigns it to the RequireDeclineReason field.
func (o *SendChoices) SetRequireDeclineReason(v bool) {
	o.RequireDeclineReason = &v
}

func (o SendChoices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendChoices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.SenderName) {
		toSerialize["senderName"] = o.SenderName
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.ExpiryDays.IsSet() {
		toSerialize["expiryDays"] = o.ExpiryDays.Get()
	}
	if !IsNil(o.ReminderDays) {
		toSerialize["reminderDays"] = o.ReminderDays
	}
	if !IsNil(o.SigningOrder) {
		toSerialize["signingOrder"] = o.SigningOrder
	}
	if !IsNil(o.InternalNote) {
		toSerialize["internalNote"] = o.InternalNote
	}
	if !IsNil(o.RequireDeclineReason) {
		toSerialize["requireDeclineReason"] = o.RequireDeclineReason
	}
	return toSerialize, nil
}

type NullableSendChoices struct {
	value *SendChoices
	isSet bool
}

func (v NullableSendChoices) Get() *SendChoices {
	return v.value
}

func (v *NullableSendChoices) Set(val *SendChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableSendChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableSendChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendChoices(val *SendChoices) *NullableSendChoices {
	return &NullableSendChoices{value: val, isSet: true}
}

func (v NullableSendChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the SendRequestPartiesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendRequestPartiesInner{}

// SendRequestPartiesInner struct for SendRequestPartiesInner
type SendRequestPartiesInner struct {
	Seat string `json:"seat"`
	Name string `json:"name"`
	Email string `json:"email"`
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId *string `json:"externalId,omitempty"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
}

type _SendRequestPartiesInner SendRequestPartiesInner

// NewSendRequestPartiesInner instantiates a new SendRequestPartiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendRequestPartiesInner(seat string, name string, email string) *SendRequestPartiesInner {
	this := SendRequestPartiesInner{}
	this.Seat = seat
	this.Name = name
	this.Email = email
	return &this
}

// NewSendRequestPartiesInnerWithDefaults instantiates a new SendRequestPartiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendRequestPartiesInnerWithDefaults() *SendRequestPartiesInner {
	this := SendRequestPartiesInner{}
	return &this
}

// GetSeat returns the Seat field value
func (o *SendRequestPartiesInner) GetSeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seat
}

// GetSeatOk returns a tuple with the Seat field value
// and a boolean to check if the value has been set.
func (o *SendRequestPartiesInner) GetSeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seat, true
}

// SetSeat sets field value
func (o *SendRequestPartiesInner) SetSeat(v string) {
	o.Seat = v
}

// GetName returns the Name field value
func (o *SendRequestPartiesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SendRequestPartiesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SendRequestPartiesInner) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *SendRequestPartiesInner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SendRequestPartiesInner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SendRequestPartiesInner) SetEmail(v string) {
	o.Email = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *SendRequestPartiesInner) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequestPartiesInner) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *SendRequestPartiesInner) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *SendRequestPartiesInner) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SendRequestPartiesInner) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendRequestPartiesInner) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SendRequestPartiesInner) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *SendRequestPartiesInner) SetMetadata(v map[string]string) {
	o.Metadata = v
}

func (o SendRequestPartiesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendRequestPartiesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seat"] = o.Seat
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *SendRequestPartiesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seat",
		"name",
		"email",
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

	varSendRequestPartiesInner := _SendRequestPartiesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendRequestPartiesInner)

	if err != nil {
		return err
	}

	*o = SendRequestPartiesInner(varSendRequestPartiesInner)

	return err
}

type NullableSendRequestPartiesInner struct {
	value *SendRequestPartiesInner
	isSet bool
}

func (v NullableSendRequestPartiesInner) Get() *SendRequestPartiesInner {
	return v.value
}

func (v *NullableSendRequestPartiesInner) Set(val *SendRequestPartiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSendRequestPartiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSendRequestPartiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendRequestPartiesInner(val *SendRequestPartiesInner) *NullableSendRequestPartiesInner {
	return &NullableSendRequestPartiesInner{value: val, isSet: true}
}

func (v NullableSendRequestPartiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendRequestPartiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



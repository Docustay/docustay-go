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

// checks if the SendDocument201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendDocument201Response{}

// SendDocument201Response struct for SendDocument201Response
type SendDocument201Response struct {
	Id string `json:"id"`
	State string `json:"state"`
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId NullableString `json:"externalId,omitempty"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	// How many signing emails went out now.
	Emailed int32 `json:"emailed"`
	// Anyone who was NOT emailed, and why (for example `limit_reached`: the daily email limit). The document is still sent; fetch their link or send a reminder later.
	NotEmailed []SendDocument201ResponseNotEmailedInner `json:"notEmailed"`
	AdditionalProperties map[string]interface{}
}

type _SendDocument201Response SendDocument201Response

// NewSendDocument201Response instantiates a new SendDocument201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendDocument201Response(id string, state string, emailed int32, notEmailed []SendDocument201ResponseNotEmailedInner) *SendDocument201Response {
	this := SendDocument201Response{}
	this.Id = id
	this.State = state
	this.Emailed = emailed
	this.NotEmailed = notEmailed
	return &this
}

// NewSendDocument201ResponseWithDefaults instantiates a new SendDocument201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendDocument201ResponseWithDefaults() *SendDocument201Response {
	this := SendDocument201Response{}
	return &this
}

// GetId returns the Id field value
func (o *SendDocument201Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SendDocument201Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SendDocument201Response) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *SendDocument201Response) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SendDocument201Response) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SendDocument201Response) SetState(v string) {
	o.State = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendDocument201Response) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendDocument201Response) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *SendDocument201Response) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *SendDocument201Response) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *SendDocument201Response) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *SendDocument201Response) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SendDocument201Response) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendDocument201Response) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SendDocument201Response) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *SendDocument201Response) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetEmailed returns the Emailed field value
func (o *SendDocument201Response) GetEmailed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Emailed
}

// GetEmailedOk returns a tuple with the Emailed field value
// and a boolean to check if the value has been set.
func (o *SendDocument201Response) GetEmailedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Emailed, true
}

// SetEmailed sets field value
func (o *SendDocument201Response) SetEmailed(v int32) {
	o.Emailed = v
}

// GetNotEmailed returns the NotEmailed field value
func (o *SendDocument201Response) GetNotEmailed() []SendDocument201ResponseNotEmailedInner {
	if o == nil {
		var ret []SendDocument201ResponseNotEmailedInner
		return ret
	}

	return o.NotEmailed
}

// GetNotEmailedOk returns a tuple with the NotEmailed field value
// and a boolean to check if the value has been set.
func (o *SendDocument201Response) GetNotEmailedOk() ([]SendDocument201ResponseNotEmailedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotEmailed, true
}

// SetNotEmailed sets field value
func (o *SendDocument201Response) SetNotEmailed(v []SendDocument201ResponseNotEmailedInner) {
	o.NotEmailed = v
}

func (o SendDocument201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendDocument201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["emailed"] = o.Emailed
	toSerialize["notEmailed"] = o.NotEmailed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SendDocument201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
		"emailed",
		"notEmailed",
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

	varSendDocument201Response := _SendDocument201Response{}

	err = json.Unmarshal(data, &varSendDocument201Response)

	if err != nil {
		return err
	}

	*o = SendDocument201Response(varSendDocument201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "emailed")
		delete(additionalProperties, "notEmailed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSendDocument201Response struct {
	value *SendDocument201Response
	isSet bool
}

func (v NullableSendDocument201Response) Get() *SendDocument201Response {
	return v.value
}

func (v *NullableSendDocument201Response) Set(val *SendDocument201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSendDocument201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSendDocument201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendDocument201Response(val *SendDocument201Response) *NullableSendDocument201Response {
	return &NullableSendDocument201Response{value: val, isSet: true}
}

func (v NullableSendDocument201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendDocument201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



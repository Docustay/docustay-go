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

// checks if the UseTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UseTemplateRequest{}

// UseTemplateRequest struct for UseTemplateRequest
type UseTemplateRequest struct {
	Email *EmailChoice `json:"email,omitempty"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId *string `json:"externalId,omitempty"`
	Title *string `json:"title,omitempty"`
	People []Person `json:"people"`
	VariableValues map[string]string `json:"variableValues,omitempty"`
	Facts map[string]string `json:"facts,omitempty"`
	Send *SendChoices `json:"send,omitempty"`
	SendNow *bool `json:"sendNow,omitempty"`
}

type _UseTemplateRequest UseTemplateRequest

// NewUseTemplateRequest instantiates a new UseTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUseTemplateRequest(people []Person) *UseTemplateRequest {
	this := UseTemplateRequest{}
	this.People = people
	return &this
}

// NewUseTemplateRequestWithDefaults instantiates a new UseTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUseTemplateRequestWithDefaults() *UseTemplateRequest {
	this := UseTemplateRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UseTemplateRequest) GetEmail() EmailChoice {
	if o == nil || IsNil(o.Email) {
		var ret EmailChoice
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UseTemplateRequest) GetEmailOk() (*EmailChoice, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UseTemplateRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailChoice and assigns it to the Email field.
func (o *UseTemplateRequest) SetEmail(v EmailChoice) {
	o.Email = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UseTemplateRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UseTemplateRequest) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UseTemplateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UseTemplateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *UseTemplateRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UseTemplateRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *UseTemplateRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *UseTemplateRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UseTemplateRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UseTemplateRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UseTemplateRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UseTemplateRequest) SetTitle(v string) {
	o.Title = &v
}

// GetPeople returns the People field value
func (o *UseTemplateRequest) GetPeople() []Person {
	if o == nil {
		var ret []Person
		return ret
	}

	return o.People
}

// GetPeopleOk returns a tuple with the People field value
// and a boolean to check if the value has been set.
func (o *UseTemplateRequest) GetPeopleOk() ([]Person, bool) {
	if o == nil {
		return nil, false
	}
	return o.People, true
}

// SetPeople sets field value
func (o *UseTemplateRequest) SetPeople(v []Person) {
	o.People = v
}

// GetVariableValues returns the VariableValues field value if set, zero value otherwise.
func (o *UseTemplateRequest) GetVariableValues() map[string]string {
	if o == nil || IsNil(o.VariableValues) {
		var ret map[string]string
		return ret
	}
	return o.VariableValues
}

// GetVariableValuesOk returns a tuple with the VariableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UseTemplateRequest) GetVariableValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.VariableValues) {
		return map[string]string{}, false
	}
	return o.VariableValues, true
}

// HasVariableValues returns a boolean if a field has been set.
func (o *UseTemplateRequest) HasVariableValues() bool {
	if o != nil && !IsNil(o.VariableValues) {
		return true
	}

	return false
}

// SetVariableValues gets a reference to the given map[string]string and assigns it to the VariableValues field.
func (o *UseTemplateRequest) SetVariableValues(v map[string]string) {
	o.VariableValues = v
}

// GetFacts returns the Facts field value if set, zero value otherwise.
func (o *UseTemplateRequest) GetFacts() map[string]string {
	if o == nil || IsNil(o.Facts) {
		var ret map[string]string
		return ret
	}
	return o.Facts
}

// GetFactsOk returns a tuple with the Facts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UseTemplateRequest) GetFactsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Facts) {
		return map[string]string{}, false
	}
	return o.Facts, true
}

// HasFacts returns a boolean if a field has been set.
func (o *UseTemplateRequest) HasFacts() bool {
	if o != nil && !IsNil(o.Facts) {
		return true
	}

	return false
}

// SetFacts gets a reference to the given map[string]string and assigns it to the Facts field.
func (o *UseTemplateRequest) SetFacts(v map[string]string) {
	o.Facts = v
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *UseTemplateRequest) GetSend() SendChoices {
	if o == nil || IsNil(o.Send) {
		var ret SendChoices
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UseTemplateRequest) GetSendOk() (*SendChoices, bool) {
	if o == nil || IsNil(o.Send) {
		return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *UseTemplateRequest) HasSend() bool {
	if o != nil && !IsNil(o.Send) {
		return true
	}

	return false
}

// SetSend gets a reference to the given SendChoices and assigns it to the Send field.
func (o *UseTemplateRequest) SetSend(v SendChoices) {
	o.Send = &v
}

// GetSendNow returns the SendNow field value if set, zero value otherwise.
func (o *UseTemplateRequest) GetSendNow() bool {
	if o == nil || IsNil(o.SendNow) {
		var ret bool
		return ret
	}
	return *o.SendNow
}

// GetSendNowOk returns a tuple with the SendNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UseTemplateRequest) GetSendNowOk() (*bool, bool) {
	if o == nil || IsNil(o.SendNow) {
		return nil, false
	}
	return o.SendNow, true
}

// HasSendNow returns a boolean if a field has been set.
func (o *UseTemplateRequest) HasSendNow() bool {
	if o != nil && !IsNil(o.SendNow) {
		return true
	}

	return false
}

// SetSendNow gets a reference to the given bool and assigns it to the SendNow field.
func (o *UseTemplateRequest) SetSendNow(v bool) {
	o.SendNow = &v
}

func (o UseTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UseTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	toSerialize["people"] = o.People
	if !IsNil(o.VariableValues) {
		toSerialize["variableValues"] = o.VariableValues
	}
	if !IsNil(o.Facts) {
		toSerialize["facts"] = o.Facts
	}
	if !IsNil(o.Send) {
		toSerialize["send"] = o.Send
	}
	if !IsNil(o.SendNow) {
		toSerialize["sendNow"] = o.SendNow
	}
	return toSerialize, nil
}

func (o *UseTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"people",
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

	varUseTemplateRequest := _UseTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUseTemplateRequest)

	if err != nil {
		return err
	}

	*o = UseTemplateRequest(varUseTemplateRequest)

	return err
}

type NullableUseTemplateRequest struct {
	value *UseTemplateRequest
	isSet bool
}

func (v NullableUseTemplateRequest) Get() *UseTemplateRequest {
	return v.value
}

func (v *NullableUseTemplateRequest) Set(val *UseTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUseTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUseTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUseTemplateRequest(val *UseTemplateRequest) *NullableUseTemplateRequest {
	return &NullableUseTemplateRequest{value: val, isSet: true}
}

func (v NullableUseTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUseTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



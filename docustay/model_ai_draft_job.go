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

// checks if the AiDraftJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiDraftJob{}

// AiDraftJob struct for AiDraftJob
type AiDraftJob struct {
	Id string `json:"id"`
	State string `json:"state"`
	Message NullableString `json:"message,omitempty"`
	Source map[string]*interface{} `json:"source,omitempty"`
	Patches map[string]*interface{} `json:"patches,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AiDraftJob AiDraftJob

// NewAiDraftJob instantiates a new AiDraftJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiDraftJob(id string, state string) *AiDraftJob {
	this := AiDraftJob{}
	this.Id = id
	this.State = state
	return &this
}

// NewAiDraftJobWithDefaults instantiates a new AiDraftJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiDraftJobWithDefaults() *AiDraftJob {
	this := AiDraftJob{}
	return &this
}

// GetId returns the Id field value
func (o *AiDraftJob) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiDraftJob) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiDraftJob) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *AiDraftJob) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AiDraftJob) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *AiDraftJob) SetState(v string) {
	o.State = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiDraftJob) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiDraftJob) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *AiDraftJob) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *AiDraftJob) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *AiDraftJob) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *AiDraftJob) UnsetMessage() {
	o.Message.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiDraftJob) GetSource() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiDraftJob) GetSourceOk() (map[string]*interface{}, bool) {
	if o == nil || IsNil(o.Source) {
		return map[string]*interface{}{}, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AiDraftJob) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given map[string]*interface{} and assigns it to the Source field.
func (o *AiDraftJob) SetSource(v map[string]*interface{}) {
	o.Source = v
}

// GetPatches returns the Patches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiDraftJob) GetPatches() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}
	return o.Patches
}

// GetPatchesOk returns a tuple with the Patches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiDraftJob) GetPatchesOk() (map[string]*interface{}, bool) {
	if o == nil || IsNil(o.Patches) {
		return map[string]*interface{}{}, false
	}
	return o.Patches, true
}

// HasPatches returns a boolean if a field has been set.
func (o *AiDraftJob) HasPatches() bool {
	if o != nil && !IsNil(o.Patches) {
		return true
	}

	return false
}

// SetPatches gets a reference to the given map[string]*interface{} and assigns it to the Patches field.
func (o *AiDraftJob) SetPatches(v map[string]*interface{}) {
	o.Patches = v
}

func (o AiDraftJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiDraftJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Patches != nil {
		toSerialize["patches"] = o.Patches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiDraftJob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
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

	varAiDraftJob := _AiDraftJob{}

	err = json.Unmarshal(data, &varAiDraftJob)

	if err != nil {
		return err
	}

	*o = AiDraftJob(varAiDraftJob)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "message")
		delete(additionalProperties, "source")
		delete(additionalProperties, "patches")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiDraftJob struct {
	value *AiDraftJob
	isSet bool
}

func (v NullableAiDraftJob) Get() *AiDraftJob {
	return v.value
}

func (v *NullableAiDraftJob) Set(val *AiDraftJob) {
	v.value = val
	v.isSet = true
}

func (v NullableAiDraftJob) IsSet() bool {
	return v.isSet
}

func (v *NullableAiDraftJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiDraftJob(val *AiDraftJob) *NullableAiDraftJob {
	return &NullableAiDraftJob{value: val, isSet: true}
}

func (v NullableAiDraftJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiDraftJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



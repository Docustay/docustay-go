/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the DraftWithAiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DraftWithAiRequest{}

// DraftWithAiRequest struct for DraftWithAiRequest
type DraftWithAiRequest struct {
	Prompt *string `json:"prompt,omitempty"`
	Brief *WizardBrief `json:"brief,omitempty"`
	Mode *string `json:"mode,omitempty"`
}

// NewDraftWithAiRequest instantiates a new DraftWithAiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDraftWithAiRequest() *DraftWithAiRequest {
	this := DraftWithAiRequest{}
	return &this
}

// NewDraftWithAiRequestWithDefaults instantiates a new DraftWithAiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDraftWithAiRequestWithDefaults() *DraftWithAiRequest {
	this := DraftWithAiRequest{}
	return &this
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *DraftWithAiRequest) GetPrompt() string {
	if o == nil || IsNil(o.Prompt) {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftWithAiRequest) GetPromptOk() (*string, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *DraftWithAiRequest) HasPrompt() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *DraftWithAiRequest) SetPrompt(v string) {
	o.Prompt = &v
}

// GetBrief returns the Brief field value if set, zero value otherwise.
func (o *DraftWithAiRequest) GetBrief() WizardBrief {
	if o == nil || IsNil(o.Brief) {
		var ret WizardBrief
		return ret
	}
	return *o.Brief
}

// GetBriefOk returns a tuple with the Brief field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftWithAiRequest) GetBriefOk() (*WizardBrief, bool) {
	if o == nil || IsNil(o.Brief) {
		return nil, false
	}
	return o.Brief, true
}

// HasBrief returns a boolean if a field has been set.
func (o *DraftWithAiRequest) HasBrief() bool {
	if o != nil && !IsNil(o.Brief) {
		return true
	}

	return false
}

// SetBrief gets a reference to the given WizardBrief and assigns it to the Brief field.
func (o *DraftWithAiRequest) SetBrief(v WizardBrief) {
	o.Brief = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DraftWithAiRequest) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftWithAiRequest) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DraftWithAiRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DraftWithAiRequest) SetMode(v string) {
	o.Mode = &v
}

func (o DraftWithAiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DraftWithAiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.Brief) {
		toSerialize["brief"] = o.Brief
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableDraftWithAiRequest struct {
	value *DraftWithAiRequest
	isSet bool
}

func (v NullableDraftWithAiRequest) Get() *DraftWithAiRequest {
	return v.value
}

func (v *NullableDraftWithAiRequest) Set(val *DraftWithAiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDraftWithAiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDraftWithAiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDraftWithAiRequest(val *DraftWithAiRequest) *NullableDraftWithAiRequest {
	return &NullableDraftWithAiRequest{value: val, isSet: true}
}

func (v NullableDraftWithAiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDraftWithAiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



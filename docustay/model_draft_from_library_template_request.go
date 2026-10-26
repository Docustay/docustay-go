/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the DraftFromLibraryTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DraftFromLibraryTemplateRequest{}

// DraftFromLibraryTemplateRequest struct for DraftFromLibraryTemplateRequest
type DraftFromLibraryTemplateRequest struct {
	Title *string `json:"title,omitempty"`
}

// NewDraftFromLibraryTemplateRequest instantiates a new DraftFromLibraryTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDraftFromLibraryTemplateRequest() *DraftFromLibraryTemplateRequest {
	this := DraftFromLibraryTemplateRequest{}
	return &this
}

// NewDraftFromLibraryTemplateRequestWithDefaults instantiates a new DraftFromLibraryTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDraftFromLibraryTemplateRequestWithDefaults() *DraftFromLibraryTemplateRequest {
	this := DraftFromLibraryTemplateRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DraftFromLibraryTemplateRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftFromLibraryTemplateRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DraftFromLibraryTemplateRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DraftFromLibraryTemplateRequest) SetTitle(v string) {
	o.Title = &v
}

func (o DraftFromLibraryTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DraftFromLibraryTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableDraftFromLibraryTemplateRequest struct {
	value *DraftFromLibraryTemplateRequest
	isSet bool
}

func (v NullableDraftFromLibraryTemplateRequest) Get() *DraftFromLibraryTemplateRequest {
	return v.value
}

func (v *NullableDraftFromLibraryTemplateRequest) Set(val *DraftFromLibraryTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDraftFromLibraryTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDraftFromLibraryTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDraftFromLibraryTemplateRequest(val *DraftFromLibraryTemplateRequest) *NullableDraftFromLibraryTemplateRequest {
	return &NullableDraftFromLibraryTemplateRequest{value: val, isSet: true}
}

func (v NullableDraftFromLibraryTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDraftFromLibraryTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



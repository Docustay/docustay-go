/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
)

// checks if the CustomizeLibraryTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomizeLibraryTemplateRequest{}

// CustomizeLibraryTemplateRequest struct for CustomizeLibraryTemplateRequest
type CustomizeLibraryTemplateRequest struct {
	Name *string `json:"name,omitempty"`
}

// NewCustomizeLibraryTemplateRequest instantiates a new CustomizeLibraryTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomizeLibraryTemplateRequest() *CustomizeLibraryTemplateRequest {
	this := CustomizeLibraryTemplateRequest{}
	return &this
}

// NewCustomizeLibraryTemplateRequestWithDefaults instantiates a new CustomizeLibraryTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomizeLibraryTemplateRequestWithDefaults() *CustomizeLibraryTemplateRequest {
	this := CustomizeLibraryTemplateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomizeLibraryTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizeLibraryTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomizeLibraryTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomizeLibraryTemplateRequest) SetName(v string) {
	o.Name = &v
}

func (o CustomizeLibraryTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomizeLibraryTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCustomizeLibraryTemplateRequest struct {
	value *CustomizeLibraryTemplateRequest
	isSet bool
}

func (v NullableCustomizeLibraryTemplateRequest) Get() *CustomizeLibraryTemplateRequest {
	return v.value
}

func (v *NullableCustomizeLibraryTemplateRequest) Set(val *CustomizeLibraryTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomizeLibraryTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomizeLibraryTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomizeLibraryTemplateRequest(val *CustomizeLibraryTemplateRequest) *NullableCustomizeLibraryTemplateRequest {
	return &NullableCustomizeLibraryTemplateRequest{value: val, isSet: true}
}

func (v NullableCustomizeLibraryTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomizeLibraryTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



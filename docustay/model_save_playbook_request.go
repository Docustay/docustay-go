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

// checks if the SavePlaybookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavePlaybookRequest{}

// SavePlaybookRequest struct for SavePlaybookRequest
type SavePlaybookRequest struct {
	Name string `json:"name" validate:"regexp=^[a-z0-9][a-z0-9-]{2\\,60}$"`
	Title string `json:"title"`
	Description *string `json:"description,omitempty"`
	Plays []string `json:"plays,omitempty"`
	// `builtin:<slug>` for a library template, or one of your templates' family ids.
	Templates []string `json:"templates,omitempty"`
}

type _SavePlaybookRequest SavePlaybookRequest

// NewSavePlaybookRequest instantiates a new SavePlaybookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavePlaybookRequest(name string, title string) *SavePlaybookRequest {
	this := SavePlaybookRequest{}
	this.Name = name
	this.Title = title
	return &this
}

// NewSavePlaybookRequestWithDefaults instantiates a new SavePlaybookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavePlaybookRequestWithDefaults() *SavePlaybookRequest {
	this := SavePlaybookRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SavePlaybookRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SavePlaybookRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SavePlaybookRequest) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *SavePlaybookRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SavePlaybookRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SavePlaybookRequest) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SavePlaybookRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavePlaybookRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SavePlaybookRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SavePlaybookRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPlays returns the Plays field value if set, zero value otherwise.
func (o *SavePlaybookRequest) GetPlays() []string {
	if o == nil || IsNil(o.Plays) {
		var ret []string
		return ret
	}
	return o.Plays
}

// GetPlaysOk returns a tuple with the Plays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavePlaybookRequest) GetPlaysOk() ([]string, bool) {
	if o == nil || IsNil(o.Plays) {
		return nil, false
	}
	return o.Plays, true
}

// HasPlays returns a boolean if a field has been set.
func (o *SavePlaybookRequest) HasPlays() bool {
	if o != nil && !IsNil(o.Plays) {
		return true
	}

	return false
}

// SetPlays gets a reference to the given []string and assigns it to the Plays field.
func (o *SavePlaybookRequest) SetPlays(v []string) {
	o.Plays = v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *SavePlaybookRequest) GetTemplates() []string {
	if o == nil || IsNil(o.Templates) {
		var ret []string
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavePlaybookRequest) GetTemplatesOk() ([]string, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *SavePlaybookRequest) HasTemplates() bool {
	if o != nil && !IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []string and assigns it to the Templates field.
func (o *SavePlaybookRequest) SetTemplates(v []string) {
	o.Templates = v
}

func (o SavePlaybookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavePlaybookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Plays) {
		toSerialize["plays"] = o.Plays
	}
	if !IsNil(o.Templates) {
		toSerialize["templates"] = o.Templates
	}
	return toSerialize, nil
}

func (o *SavePlaybookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"title",
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

	varSavePlaybookRequest := _SavePlaybookRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSavePlaybookRequest)

	if err != nil {
		return err
	}

	*o = SavePlaybookRequest(varSavePlaybookRequest)

	return err
}

type NullableSavePlaybookRequest struct {
	value *SavePlaybookRequest
	isSet bool
}

func (v NullableSavePlaybookRequest) Get() *SavePlaybookRequest {
	return v.value
}

func (v *NullableSavePlaybookRequest) Set(val *SavePlaybookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSavePlaybookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSavePlaybookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavePlaybookRequest(val *SavePlaybookRequest) *NullableSavePlaybookRequest {
	return &NullableSavePlaybookRequest{value: val, isSet: true}
}

func (v NullableSavePlaybookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavePlaybookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



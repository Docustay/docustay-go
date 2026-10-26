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

// checks if the CreateTemplateFromSource201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTemplateFromSource201Response{}

// CreateTemplateFromSource201Response struct for CreateTemplateFromSource201Response
type CreateTemplateFromSource201Response struct {
	TemplateId string `json:"templateId"`
	Profile string `json:"profile"`
	Version int32 `json:"version"`
	Hash string `json:"hash"`
	Warnings []interface{} `json:"warnings"`
	AdditionalProperties map[string]interface{}
}

type _CreateTemplateFromSource201Response CreateTemplateFromSource201Response

// NewCreateTemplateFromSource201Response instantiates a new CreateTemplateFromSource201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTemplateFromSource201Response(templateId string, profile string, version int32, hash string, warnings []interface{}) *CreateTemplateFromSource201Response {
	this := CreateTemplateFromSource201Response{}
	this.TemplateId = templateId
	this.Profile = profile
	this.Version = version
	this.Hash = hash
	this.Warnings = warnings
	return &this
}

// NewCreateTemplateFromSource201ResponseWithDefaults instantiates a new CreateTemplateFromSource201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTemplateFromSource201ResponseWithDefaults() *CreateTemplateFromSource201Response {
	this := CreateTemplateFromSource201Response{}
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *CreateTemplateFromSource201Response) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSource201Response) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *CreateTemplateFromSource201Response) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetProfile returns the Profile field value
func (o *CreateTemplateFromSource201Response) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSource201Response) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *CreateTemplateFromSource201Response) SetProfile(v string) {
	o.Profile = v
}

// GetVersion returns the Version field value
func (o *CreateTemplateFromSource201Response) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSource201Response) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CreateTemplateFromSource201Response) SetVersion(v int32) {
	o.Version = v
}

// GetHash returns the Hash field value
func (o *CreateTemplateFromSource201Response) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSource201Response) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *CreateTemplateFromSource201Response) SetHash(v string) {
	o.Hash = v
}

// GetWarnings returns the Warnings field value
func (o *CreateTemplateFromSource201Response) GetWarnings() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSource201Response) GetWarningsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *CreateTemplateFromSource201Response) SetWarnings(v []interface{}) {
	o.Warnings = v
}

func (o CreateTemplateFromSource201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTemplateFromSource201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateId"] = o.TemplateId
	toSerialize["profile"] = o.Profile
	toSerialize["version"] = o.Version
	toSerialize["hash"] = o.Hash
	toSerialize["warnings"] = o.Warnings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateTemplateFromSource201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateId",
		"profile",
		"version",
		"hash",
		"warnings",
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

	varCreateTemplateFromSource201Response := _CreateTemplateFromSource201Response{}

	err = json.Unmarshal(data, &varCreateTemplateFromSource201Response)

	if err != nil {
		return err
	}

	*o = CreateTemplateFromSource201Response(varCreateTemplateFromSource201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "version")
		delete(additionalProperties, "hash")
		delete(additionalProperties, "warnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateTemplateFromSource201Response struct {
	value *CreateTemplateFromSource201Response
	isSet bool
}

func (v NullableCreateTemplateFromSource201Response) Get() *CreateTemplateFromSource201Response {
	return v.value
}

func (v *NullableCreateTemplateFromSource201Response) Set(val *CreateTemplateFromSource201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTemplateFromSource201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTemplateFromSource201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTemplateFromSource201Response(val *CreateTemplateFromSource201Response) *NullableCreateTemplateFromSource201Response {
	return &NullableCreateTemplateFromSource201Response{value: val, isSet: true}
}

func (v NullableCreateTemplateFromSource201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTemplateFromSource201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



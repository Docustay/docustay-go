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

// checks if the BuilderSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuilderSession{}

// BuilderSession struct for BuilderSession
type BuilderSession struct {
	Url string `json:"url"`
	ExpiresAt string `json:"expiresAt"`
	AdditionalProperties map[string]interface{}
}

type _BuilderSession BuilderSession

// NewBuilderSession instantiates a new BuilderSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuilderSession(url string, expiresAt string) *BuilderSession {
	this := BuilderSession{}
	this.Url = url
	this.ExpiresAt = expiresAt
	return &this
}

// NewBuilderSessionWithDefaults instantiates a new BuilderSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuilderSessionWithDefaults() *BuilderSession {
	this := BuilderSession{}
	return &this
}

// GetUrl returns the Url field value
func (o *BuilderSession) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BuilderSession) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BuilderSession) SetUrl(v string) {
	o.Url = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *BuilderSession) GetExpiresAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *BuilderSession) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *BuilderSession) SetExpiresAt(v string) {
	o.ExpiresAt = v
}

func (o BuilderSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuilderSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["expiresAt"] = o.ExpiresAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BuilderSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"expiresAt",
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

	varBuilderSession := _BuilderSession{}

	err = json.Unmarshal(data, &varBuilderSession)

	if err != nil {
		return err
	}

	*o = BuilderSession(varBuilderSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "expiresAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuilderSession struct {
	value *BuilderSession
	isSet bool
}

func (v NullableBuilderSession) Get() *BuilderSession {
	return v.value
}

func (v *NullableBuilderSession) Set(val *BuilderSession) {
	v.value = val
	v.isSet = true
}

func (v NullableBuilderSession) IsSet() bool {
	return v.isSet
}

func (v *NullableBuilderSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuilderSession(val *BuilderSession) *NullableBuilderSession {
	return &NullableBuilderSession{value: val, isSet: true}
}

func (v NullableBuilderSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuilderSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



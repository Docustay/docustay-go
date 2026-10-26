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

// checks if the TemplateMade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateMade{}

// TemplateMade struct for TemplateMade
type TemplateMade struct {
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId NullableString `json:"externalId,omitempty"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	TemplateId string `json:"templateId"`
	Fields int32 `json:"fields"`
	Warnings []TemplateMadeWarningsInner `json:"warnings"`
	Detected *TemplateMadeDetected `json:"detected,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateMade TemplateMade

// NewTemplateMade instantiates a new TemplateMade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateMade(templateId string, fields int32, warnings []TemplateMadeWarningsInner) *TemplateMade {
	this := TemplateMade{}
	this.TemplateId = templateId
	this.Fields = fields
	this.Warnings = warnings
	return &this
}

// NewTemplateMadeWithDefaults instantiates a new TemplateMade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateMadeWithDefaults() *TemplateMade {
	this := TemplateMade{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateMade) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateMade) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *TemplateMade) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *TemplateMade) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *TemplateMade) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *TemplateMade) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TemplateMade) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMade) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TemplateMade) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *TemplateMade) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetTemplateId returns the TemplateId field value
func (o *TemplateMade) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *TemplateMade) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *TemplateMade) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetFields returns the Fields field value
func (o *TemplateMade) GetFields() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *TemplateMade) GetFieldsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *TemplateMade) SetFields(v int32) {
	o.Fields = v
}

// GetWarnings returns the Warnings field value
func (o *TemplateMade) GetWarnings() []TemplateMadeWarningsInner {
	if o == nil {
		var ret []TemplateMadeWarningsInner
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *TemplateMade) GetWarningsOk() ([]TemplateMadeWarningsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *TemplateMade) SetWarnings(v []TemplateMadeWarningsInner) {
	o.Warnings = v
}

// GetDetected returns the Detected field value if set, zero value otherwise.
func (o *TemplateMade) GetDetected() TemplateMadeDetected {
	if o == nil || IsNil(o.Detected) {
		var ret TemplateMadeDetected
		return ret
	}
	return *o.Detected
}

// GetDetectedOk returns a tuple with the Detected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMade) GetDetectedOk() (*TemplateMadeDetected, bool) {
	if o == nil || IsNil(o.Detected) {
		return nil, false
	}
	return o.Detected, true
}

// HasDetected returns a boolean if a field has been set.
func (o *TemplateMade) HasDetected() bool {
	if o != nil && !IsNil(o.Detected) {
		return true
	}

	return false
}

// SetDetected gets a reference to the given TemplateMadeDetected and assigns it to the Detected field.
func (o *TemplateMade) SetDetected(v TemplateMadeDetected) {
	o.Detected = &v
}

func (o TemplateMade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateMade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["templateId"] = o.TemplateId
	toSerialize["fields"] = o.Fields
	toSerialize["warnings"] = o.Warnings
	if !IsNil(o.Detected) {
		toSerialize["detected"] = o.Detected
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateMade) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateId",
		"fields",
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

	varTemplateMade := _TemplateMade{}

	err = json.Unmarshal(data, &varTemplateMade)

	if err != nil {
		return err
	}

	*o = TemplateMade(varTemplateMade)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "detected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateMade struct {
	value *TemplateMade
	isSet bool
}

func (v NullableTemplateMade) Get() *TemplateMade {
	return v.value
}

func (v *NullableTemplateMade) Set(val *TemplateMade) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateMade) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateMade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateMade(val *TemplateMade) *NullableTemplateMade {
	return &NullableTemplateMade{value: val, isSet: true}
}

func (v NullableTemplateMade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateMade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



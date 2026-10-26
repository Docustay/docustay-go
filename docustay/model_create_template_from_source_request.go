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

// checks if the CreateTemplateFromSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTemplateFromSourceRequest{}

// CreateTemplateFromSourceRequest struct for CreateTemplateFromSourceRequest
type CreateTemplateFromSourceRequest struct {
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId *string `json:"externalId,omitempty"`
	Name string `json:"name"`
	// The document source: a ProseMirror/TipTap JSON document (`{type:\"doc\",content:[…]}`) using headings, paragraphs, lists, tables, `fieldTag`, `signatureBlock`, `variableChip`. See the guide on document sources.
	Source map[string]*interface{} `json:"source"`
	Profile *string `json:"profile,omitempty"`
	Paper *string `json:"paper,omitempty"`
}

type _CreateTemplateFromSourceRequest CreateTemplateFromSourceRequest

// NewCreateTemplateFromSourceRequest instantiates a new CreateTemplateFromSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTemplateFromSourceRequest(name string, source map[string]*interface{}) *CreateTemplateFromSourceRequest {
	this := CreateTemplateFromSourceRequest{}
	this.Name = name
	this.Source = source
	return &this
}

// NewCreateTemplateFromSourceRequestWithDefaults instantiates a new CreateTemplateFromSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTemplateFromSourceRequestWithDefaults() *CreateTemplateFromSourceRequest {
	this := CreateTemplateFromSourceRequest{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateTemplateFromSourceRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSourceRequest) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateTemplateFromSourceRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateTemplateFromSourceRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *CreateTemplateFromSourceRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSourceRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *CreateTemplateFromSourceRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *CreateTemplateFromSourceRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetName returns the Name field value
func (o *CreateTemplateFromSourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTemplateFromSourceRequest) SetName(v string) {
	o.Name = v
}

// GetSource returns the Source field value
func (o *CreateTemplateFromSourceRequest) GetSource() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSourceRequest) GetSourceOk() (map[string]*interface{}, bool) {
	if o == nil {
		return map[string]*interface{}{}, false
	}
	return o.Source, true
}

// SetSource sets field value
func (o *CreateTemplateFromSourceRequest) SetSource(v map[string]*interface{}) {
	o.Source = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *CreateTemplateFromSourceRequest) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSourceRequest) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *CreateTemplateFromSourceRequest) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *CreateTemplateFromSourceRequest) SetProfile(v string) {
	o.Profile = &v
}

// GetPaper returns the Paper field value if set, zero value otherwise.
func (o *CreateTemplateFromSourceRequest) GetPaper() string {
	if o == nil || IsNil(o.Paper) {
		var ret string
		return ret
	}
	return *o.Paper
}

// GetPaperOk returns a tuple with the Paper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateFromSourceRequest) GetPaperOk() (*string, bool) {
	if o == nil || IsNil(o.Paper) {
		return nil, false
	}
	return o.Paper, true
}

// HasPaper returns a boolean if a field has been set.
func (o *CreateTemplateFromSourceRequest) HasPaper() bool {
	if o != nil && !IsNil(o.Paper) {
		return true
	}

	return false
}

// SetPaper gets a reference to the given string and assigns it to the Paper field.
func (o *CreateTemplateFromSourceRequest) SetPaper(v string) {
	o.Paper = &v
}

func (o CreateTemplateFromSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTemplateFromSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	toSerialize["name"] = o.Name
	toSerialize["source"] = o.Source
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Paper) {
		toSerialize["paper"] = o.Paper
	}
	return toSerialize, nil
}

func (o *CreateTemplateFromSourceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"source",
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

	varCreateTemplateFromSourceRequest := _CreateTemplateFromSourceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTemplateFromSourceRequest)

	if err != nil {
		return err
	}

	*o = CreateTemplateFromSourceRequest(varCreateTemplateFromSourceRequest)

	return err
}

type NullableCreateTemplateFromSourceRequest struct {
	value *CreateTemplateFromSourceRequest
	isSet bool
}

func (v NullableCreateTemplateFromSourceRequest) Get() *CreateTemplateFromSourceRequest {
	return v.value
}

func (v *NullableCreateTemplateFromSourceRequest) Set(val *CreateTemplateFromSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTemplateFromSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTemplateFromSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTemplateFromSourceRequest(val *CreateTemplateFromSourceRequest) *NullableCreateTemplateFromSourceRequest {
	return &NullableCreateTemplateFromSourceRequest{value: val, isSet: true}
}

func (v NullableCreateTemplateFromSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTemplateFromSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



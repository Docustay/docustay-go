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

// checks if the FileTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileTemplateRequest{}

// FileTemplateRequest struct for FileTemplateRequest
type FileTemplateRequest struct {
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId *string `json:"externalId,omitempty"`
	Name string `json:"name"`
	FileName string `json:"fileName"`
	// The file, base64-encoded: a PDF file with text tags written in it.
	File string `json:"file"`
	// `tags` (default) places fields only where a text tag is written. `auto` also looks at the page itself — lines, boxes, tick boxes, highlighted bars and the words beside them — and places the fields it is confident about, with the people grouped into seats.
	Fields *string `json:"fields,omitempty"`
	// With `fields: auto`: the least sure a field may be and still be placed (default 0.6).
	MinConfidence *float32 `json:"minConfidence,omitempty"`
}

type _FileTemplateRequest FileTemplateRequest

// NewFileTemplateRequest instantiates a new FileTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileTemplateRequest(name string, fileName string, file string) *FileTemplateRequest {
	this := FileTemplateRequest{}
	this.Name = name
	this.FileName = fileName
	this.File = file
	return &this
}

// NewFileTemplateRequestWithDefaults instantiates a new FileTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileTemplateRequestWithDefaults() *FileTemplateRequest {
	this := FileTemplateRequest{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FileTemplateRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileTemplateRequest) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FileTemplateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *FileTemplateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *FileTemplateRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileTemplateRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *FileTemplateRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *FileTemplateRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetName returns the Name field value
func (o *FileTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetFileName returns the FileName field value
func (o *FileTemplateRequest) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *FileTemplateRequest) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *FileTemplateRequest) SetFileName(v string) {
	o.FileName = v
}

// GetFile returns the File field value
func (o *FileTemplateRequest) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *FileTemplateRequest) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *FileTemplateRequest) SetFile(v string) {
	o.File = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *FileTemplateRequest) GetFields() string {
	if o == nil || IsNil(o.Fields) {
		var ret string
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileTemplateRequest) GetFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *FileTemplateRequest) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given string and assigns it to the Fields field.
func (o *FileTemplateRequest) SetFields(v string) {
	o.Fields = &v
}

// GetMinConfidence returns the MinConfidence field value if set, zero value otherwise.
func (o *FileTemplateRequest) GetMinConfidence() float32 {
	if o == nil || IsNil(o.MinConfidence) {
		var ret float32
		return ret
	}
	return *o.MinConfidence
}

// GetMinConfidenceOk returns a tuple with the MinConfidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileTemplateRequest) GetMinConfidenceOk() (*float32, bool) {
	if o == nil || IsNil(o.MinConfidence) {
		return nil, false
	}
	return o.MinConfidence, true
}

// HasMinConfidence returns a boolean if a field has been set.
func (o *FileTemplateRequest) HasMinConfidence() bool {
	if o != nil && !IsNil(o.MinConfidence) {
		return true
	}

	return false
}

// SetMinConfidence gets a reference to the given float32 and assigns it to the MinConfidence field.
func (o *FileTemplateRequest) SetMinConfidence(v float32) {
	o.MinConfidence = &v
}

func (o FileTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	toSerialize["name"] = o.Name
	toSerialize["fileName"] = o.FileName
	toSerialize["file"] = o.File
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.MinConfidence) {
		toSerialize["minConfidence"] = o.MinConfidence
	}
	return toSerialize, nil
}

func (o *FileTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"fileName",
		"file",
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

	varFileTemplateRequest := _FileTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileTemplateRequest)

	if err != nil {
		return err
	}

	*o = FileTemplateRequest(varFileTemplateRequest)

	return err
}

type NullableFileTemplateRequest struct {
	value *FileTemplateRequest
	isSet bool
}

func (v NullableFileTemplateRequest) Get() *FileTemplateRequest {
	return v.value
}

func (v *NullableFileTemplateRequest) Set(val *FileTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileTemplateRequest(val *FileTemplateRequest) *NullableFileTemplateRequest {
	return &NullableFileTemplateRequest{value: val, isSet: true}
}

func (v NullableFileTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the DetectFieldsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectFieldsRequest{}

// DetectFieldsRequest struct for DetectFieldsRequest
type DetectFieldsRequest struct {
	FileName *string `json:"fileName,omitempty"`
	// The file, base64-encoded: a PDF file.
	File string `json:"file"`
	MinConfidence *float32 `json:"minConfidence,omitempty"`
	// Read pages that are only a picture by character recognition (default true; English).
	Ocr *bool `json:"ocr,omitempty"`
}

type _DetectFieldsRequest DetectFieldsRequest

// NewDetectFieldsRequest instantiates a new DetectFieldsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectFieldsRequest(file string) *DetectFieldsRequest {
	this := DetectFieldsRequest{}
	this.File = file
	return &this
}

// NewDetectFieldsRequestWithDefaults instantiates a new DetectFieldsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectFieldsRequestWithDefaults() *DetectFieldsRequest {
	this := DetectFieldsRequest{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *DetectFieldsRequest) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectFieldsRequest) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *DetectFieldsRequest) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *DetectFieldsRequest) SetFileName(v string) {
	o.FileName = &v
}

// GetFile returns the File field value
func (o *DetectFieldsRequest) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *DetectFieldsRequest) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *DetectFieldsRequest) SetFile(v string) {
	o.File = v
}

// GetMinConfidence returns the MinConfidence field value if set, zero value otherwise.
func (o *DetectFieldsRequest) GetMinConfidence() float32 {
	if o == nil || IsNil(o.MinConfidence) {
		var ret float32
		return ret
	}
	return *o.MinConfidence
}

// GetMinConfidenceOk returns a tuple with the MinConfidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectFieldsRequest) GetMinConfidenceOk() (*float32, bool) {
	if o == nil || IsNil(o.MinConfidence) {
		return nil, false
	}
	return o.MinConfidence, true
}

// HasMinConfidence returns a boolean if a field has been set.
func (o *DetectFieldsRequest) HasMinConfidence() bool {
	if o != nil && !IsNil(o.MinConfidence) {
		return true
	}

	return false
}

// SetMinConfidence gets a reference to the given float32 and assigns it to the MinConfidence field.
func (o *DetectFieldsRequest) SetMinConfidence(v float32) {
	o.MinConfidence = &v
}

// GetOcr returns the Ocr field value if set, zero value otherwise.
func (o *DetectFieldsRequest) GetOcr() bool {
	if o == nil || IsNil(o.Ocr) {
		var ret bool
		return ret
	}
	return *o.Ocr
}

// GetOcrOk returns a tuple with the Ocr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectFieldsRequest) GetOcrOk() (*bool, bool) {
	if o == nil || IsNil(o.Ocr) {
		return nil, false
	}
	return o.Ocr, true
}

// HasOcr returns a boolean if a field has been set.
func (o *DetectFieldsRequest) HasOcr() bool {
	if o != nil && !IsNil(o.Ocr) {
		return true
	}

	return false
}

// SetOcr gets a reference to the given bool and assigns it to the Ocr field.
func (o *DetectFieldsRequest) SetOcr(v bool) {
	o.Ocr = &v
}

func (o DetectFieldsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectFieldsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	toSerialize["file"] = o.File
	if !IsNil(o.MinConfidence) {
		toSerialize["minConfidence"] = o.MinConfidence
	}
	if !IsNil(o.Ocr) {
		toSerialize["ocr"] = o.Ocr
	}
	return toSerialize, nil
}

func (o *DetectFieldsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDetectFieldsRequest := _DetectFieldsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDetectFieldsRequest)

	if err != nil {
		return err
	}

	*o = DetectFieldsRequest(varDetectFieldsRequest)

	return err
}

type NullableDetectFieldsRequest struct {
	value *DetectFieldsRequest
	isSet bool
}

func (v NullableDetectFieldsRequest) Get() *DetectFieldsRequest {
	return v.value
}

func (v *NullableDetectFieldsRequest) Set(val *DetectFieldsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectFieldsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectFieldsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectFieldsRequest(val *DetectFieldsRequest) *NullableDetectFieldsRequest {
	return &NullableDetectFieldsRequest{value: val, isSet: true}
}

func (v NullableDetectFieldsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectFieldsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



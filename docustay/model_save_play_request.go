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

// checks if the SavePlayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavePlayRequest{}

// SavePlayRequest struct for SavePlayRequest
type SavePlayRequest struct {
	Source string `json:"source"`
}

type _SavePlayRequest SavePlayRequest

// NewSavePlayRequest instantiates a new SavePlayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavePlayRequest(source string) *SavePlayRequest {
	this := SavePlayRequest{}
	this.Source = source
	return &this
}

// NewSavePlayRequestWithDefaults instantiates a new SavePlayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavePlayRequestWithDefaults() *SavePlayRequest {
	this := SavePlayRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *SavePlayRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SavePlayRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SavePlayRequest) SetSource(v string) {
	o.Source = v
}

func (o SavePlayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavePlayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

func (o *SavePlayRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varSavePlayRequest := _SavePlayRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSavePlayRequest)

	if err != nil {
		return err
	}

	*o = SavePlayRequest(varSavePlayRequest)

	return err
}

type NullableSavePlayRequest struct {
	value *SavePlayRequest
	isSet bool
}

func (v NullableSavePlayRequest) Get() *SavePlayRequest {
	return v.value
}

func (v *NullableSavePlayRequest) Set(val *SavePlayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSavePlayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSavePlayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavePlayRequest(val *SavePlayRequest) *NullableSavePlayRequest {
	return &NullableSavePlayRequest{value: val, isSet: true}
}

func (v NullableSavePlayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavePlayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



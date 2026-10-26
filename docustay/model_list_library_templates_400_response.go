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

// checks if the ListLibraryTemplates400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLibraryTemplates400Response{}

// ListLibraryTemplates400Response struct for ListLibraryTemplates400Response
type ListLibraryTemplates400Response struct {
	Error ListLibraryTemplates400ResponseError `json:"error"`
}

type _ListLibraryTemplates400Response ListLibraryTemplates400Response

// NewListLibraryTemplates400Response instantiates a new ListLibraryTemplates400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLibraryTemplates400Response(error_ ListLibraryTemplates400ResponseError) *ListLibraryTemplates400Response {
	this := ListLibraryTemplates400Response{}
	this.Error = error_
	return &this
}

// NewListLibraryTemplates400ResponseWithDefaults instantiates a new ListLibraryTemplates400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLibraryTemplates400ResponseWithDefaults() *ListLibraryTemplates400Response {
	this := ListLibraryTemplates400Response{}
	return &this
}

// GetError returns the Error field value
func (o *ListLibraryTemplates400Response) GetError() ListLibraryTemplates400ResponseError {
	if o == nil {
		var ret ListLibraryTemplates400ResponseError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ListLibraryTemplates400Response) GetErrorOk() (*ListLibraryTemplates400ResponseError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ListLibraryTemplates400Response) SetError(v ListLibraryTemplates400ResponseError) {
	o.Error = v
}

func (o ListLibraryTemplates400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListLibraryTemplates400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *ListLibraryTemplates400Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varListLibraryTemplates400Response := _ListLibraryTemplates400Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListLibraryTemplates400Response)

	if err != nil {
		return err
	}

	*o = ListLibraryTemplates400Response(varListLibraryTemplates400Response)

	return err
}

type NullableListLibraryTemplates400Response struct {
	value *ListLibraryTemplates400Response
	isSet bool
}

func (v NullableListLibraryTemplates400Response) Get() *ListLibraryTemplates400Response {
	return v.value
}

func (v *NullableListLibraryTemplates400Response) Set(val *ListLibraryTemplates400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListLibraryTemplates400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListLibraryTemplates400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLibraryTemplates400Response(val *ListLibraryTemplates400Response) *NullableListLibraryTemplates400Response {
	return &NullableListLibraryTemplates400Response{value: val, isSet: true}
}

func (v NullableListLibraryTemplates400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLibraryTemplates400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



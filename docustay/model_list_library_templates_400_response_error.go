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

// checks if the ListLibraryTemplates400ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLibraryTemplates400ResponseError{}

// ListLibraryTemplates400ResponseError struct for ListLibraryTemplates400ResponseError
type ListLibraryTemplates400ResponseError struct {
	Kind string `json:"kind"`
	Message string `json:"message"`
	RequestId *string `json:"requestId,omitempty"`
	Retryable *bool `json:"retryable,omitempty"`
}

type _ListLibraryTemplates400ResponseError ListLibraryTemplates400ResponseError

// NewListLibraryTemplates400ResponseError instantiates a new ListLibraryTemplates400ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLibraryTemplates400ResponseError(kind string, message string) *ListLibraryTemplates400ResponseError {
	this := ListLibraryTemplates400ResponseError{}
	this.Kind = kind
	this.Message = message
	return &this
}

// NewListLibraryTemplates400ResponseErrorWithDefaults instantiates a new ListLibraryTemplates400ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLibraryTemplates400ResponseErrorWithDefaults() *ListLibraryTemplates400ResponseError {
	this := ListLibraryTemplates400ResponseError{}
	return &this
}

// GetKind returns the Kind field value
func (o *ListLibraryTemplates400ResponseError) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ListLibraryTemplates400ResponseError) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ListLibraryTemplates400ResponseError) SetKind(v string) {
	o.Kind = v
}

// GetMessage returns the Message field value
func (o *ListLibraryTemplates400ResponseError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ListLibraryTemplates400ResponseError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ListLibraryTemplates400ResponseError) SetMessage(v string) {
	o.Message = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ListLibraryTemplates400ResponseError) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLibraryTemplates400ResponseError) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ListLibraryTemplates400ResponseError) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ListLibraryTemplates400ResponseError) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *ListLibraryTemplates400ResponseError) GetRetryable() bool {
	if o == nil || IsNil(o.Retryable) {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLibraryTemplates400ResponseError) GetRetryableOk() (*bool, bool) {
	if o == nil || IsNil(o.Retryable) {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *ListLibraryTemplates400ResponseError) HasRetryable() bool {
	if o != nil && !IsNil(o.Retryable) {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *ListLibraryTemplates400ResponseError) SetRetryable(v bool) {
	o.Retryable = &v
}

func (o ListLibraryTemplates400ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListLibraryTemplates400ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["message"] = o.Message
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.Retryable) {
		toSerialize["retryable"] = o.Retryable
	}
	return toSerialize, nil
}

func (o *ListLibraryTemplates400ResponseError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"message",
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

	varListLibraryTemplates400ResponseError := _ListLibraryTemplates400ResponseError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListLibraryTemplates400ResponseError)

	if err != nil {
		return err
	}

	*o = ListLibraryTemplates400ResponseError(varListLibraryTemplates400ResponseError)

	return err
}

type NullableListLibraryTemplates400ResponseError struct {
	value *ListLibraryTemplates400ResponseError
	isSet bool
}

func (v NullableListLibraryTemplates400ResponseError) Get() *ListLibraryTemplates400ResponseError {
	return v.value
}

func (v *NullableListLibraryTemplates400ResponseError) Set(val *ListLibraryTemplates400ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableListLibraryTemplates400ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableListLibraryTemplates400ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLibraryTemplates400ResponseError(val *ListLibraryTemplates400ResponseError) *NullableListLibraryTemplates400ResponseError {
	return &NullableListLibraryTemplates400ResponseError{value: val, isSet: true}
}

func (v NullableListLibraryTemplates400ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLibraryTemplates400ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



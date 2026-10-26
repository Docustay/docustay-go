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

// checks if the ErrorError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorError{}

// ErrorError struct for ErrorError
type ErrorError struct {
	Kind string `json:"kind"`
	Message string `json:"message"`
	RequestId *string `json:"requestId,omitempty"`
	Retryable *bool `json:"retryable,omitempty"`
	Detail map[string]*interface{} `json:"detail,omitempty"`
}

type _ErrorError ErrorError

// NewErrorError instantiates a new ErrorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorError(kind string, message string) *ErrorError {
	this := ErrorError{}
	this.Kind = kind
	this.Message = message
	return &this
}

// NewErrorErrorWithDefaults instantiates a new ErrorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorErrorWithDefaults() *ErrorError {
	this := ErrorError{}
	return &this
}

// GetKind returns the Kind field value
func (o *ErrorError) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ErrorError) SetKind(v string) {
	o.Kind = v
}

// GetMessage returns the Message field value
func (o *ErrorError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorError) SetMessage(v string) {
	o.Message = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ErrorError) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorError) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ErrorError) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ErrorError) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *ErrorError) GetRetryable() bool {
	if o == nil || IsNil(o.Retryable) {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorError) GetRetryableOk() (*bool, bool) {
	if o == nil || IsNil(o.Retryable) {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *ErrorError) HasRetryable() bool {
	if o != nil && !IsNil(o.Retryable) {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *ErrorError) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorError) GetDetail() map[string]*interface{} {
	if o == nil || IsNil(o.Detail) {
		var ret map[string]*interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorError) GetDetailOk() (map[string]*interface{}, bool) {
	if o == nil || IsNil(o.Detail) {
		return map[string]*interface{}{}, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorError) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given map[string]*interface{} and assigns it to the Detail field.
func (o *ErrorError) SetDetail(v map[string]*interface{}) {
	o.Detail = v
}

func (o ErrorError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["message"] = o.Message
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.Retryable) {
		toSerialize["retryable"] = o.Retryable
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

func (o *ErrorError) UnmarshalJSON(data []byte) (err error) {
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

	varErrorError := _ErrorError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorError)

	if err != nil {
		return err
	}

	*o = ErrorError(varErrorError)

	return err
}

type NullableErrorError struct {
	value *ErrorError
	isSet bool
}

func (v NullableErrorError) Get() *ErrorError {
	return v.value
}

func (v *NullableErrorError) Set(val *ErrorError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorError(val *ErrorError) *NullableErrorError {
	return &NullableErrorError{value: val, isSet: true}
}

func (v NullableErrorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



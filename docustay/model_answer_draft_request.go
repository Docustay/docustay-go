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

// checks if the AnswerDraftRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnswerDraftRequest{}

// AnswerDraftRequest struct for AnswerDraftRequest
type AnswerDraftRequest struct {
	Answers map[string]StartDraftRequestAnswersValue `json:"answers"`
	// Send-time amounts the template's words need, by key (a `needs_input` answer lists them under `values`), e.g. `{ \"deposit\": \"$500\" }`.
	Values map[string]string `json:"values,omitempty"`
	UseDefaults *bool `json:"useDefaults,omitempty"`
}

type _AnswerDraftRequest AnswerDraftRequest

// NewAnswerDraftRequest instantiates a new AnswerDraftRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnswerDraftRequest(answers map[string]StartDraftRequestAnswersValue) *AnswerDraftRequest {
	this := AnswerDraftRequest{}
	this.Answers = answers
	return &this
}

// NewAnswerDraftRequestWithDefaults instantiates a new AnswerDraftRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnswerDraftRequestWithDefaults() *AnswerDraftRequest {
	this := AnswerDraftRequest{}
	return &this
}

// GetAnswers returns the Answers field value
func (o *AnswerDraftRequest) GetAnswers() map[string]StartDraftRequestAnswersValue {
	if o == nil {
		var ret map[string]StartDraftRequestAnswersValue
		return ret
	}

	return o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value
// and a boolean to check if the value has been set.
func (o *AnswerDraftRequest) GetAnswersOk() (map[string]StartDraftRequestAnswersValue, bool) {
	if o == nil {
		return map[string]StartDraftRequestAnswersValue{}, false
	}
	return o.Answers, true
}

// SetAnswers sets field value
func (o *AnswerDraftRequest) SetAnswers(v map[string]StartDraftRequestAnswersValue) {
	o.Answers = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *AnswerDraftRequest) GetValues() map[string]string {
	if o == nil || IsNil(o.Values) {
		var ret map[string]string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerDraftRequest) GetValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Values) {
		return map[string]string{}, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *AnswerDraftRequest) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]string and assigns it to the Values field.
func (o *AnswerDraftRequest) SetValues(v map[string]string) {
	o.Values = v
}

// GetUseDefaults returns the UseDefaults field value if set, zero value otherwise.
func (o *AnswerDraftRequest) GetUseDefaults() bool {
	if o == nil || IsNil(o.UseDefaults) {
		var ret bool
		return ret
	}
	return *o.UseDefaults
}

// GetUseDefaultsOk returns a tuple with the UseDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerDraftRequest) GetUseDefaultsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDefaults) {
		return nil, false
	}
	return o.UseDefaults, true
}

// HasUseDefaults returns a boolean if a field has been set.
func (o *AnswerDraftRequest) HasUseDefaults() bool {
	if o != nil && !IsNil(o.UseDefaults) {
		return true
	}

	return false
}

// SetUseDefaults gets a reference to the given bool and assigns it to the UseDefaults field.
func (o *AnswerDraftRequest) SetUseDefaults(v bool) {
	o.UseDefaults = &v
}

func (o AnswerDraftRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnswerDraftRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["answers"] = o.Answers
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.UseDefaults) {
		toSerialize["useDefaults"] = o.UseDefaults
	}
	return toSerialize, nil
}

func (o *AnswerDraftRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"answers",
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

	varAnswerDraftRequest := _AnswerDraftRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnswerDraftRequest)

	if err != nil {
		return err
	}

	*o = AnswerDraftRequest(varAnswerDraftRequest)

	return err
}

type NullableAnswerDraftRequest struct {
	value *AnswerDraftRequest
	isSet bool
}

func (v NullableAnswerDraftRequest) Get() *AnswerDraftRequest {
	return v.value
}

func (v *NullableAnswerDraftRequest) Set(val *AnswerDraftRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnswerDraftRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnswerDraftRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnswerDraftRequest(val *AnswerDraftRequest) *NullableAnswerDraftRequest {
	return &NullableAnswerDraftRequest{value: val, isSet: true}
}

func (v NullableAnswerDraftRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnswerDraftRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



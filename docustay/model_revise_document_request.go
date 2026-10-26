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

// checks if the ReviseDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviseDocumentRequest{}

// ReviseDocumentRequest struct for ReviseDocumentRequest
type ReviseDocumentRequest struct {
	// A written document: TipTap JSON `{type:\"doc\",content:[…]}`. Field tags (`fieldTag` with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (`variableChip`; `custom.<key>` chips are the template's questions), and headings with `audience: \"sender\"` (Sender only: in nothing a signer sees, signs or receives).
	Source map[string]*interface{} `json:"source"`
	Facts map[string]string `json:"facts,omitempty"`
	VariableValues map[string]string `json:"variableValues,omitempty"`
	// Old role word → new role word when you renamed a role.
	RoleMap map[string]string `json:"roleMap,omitempty"`
}

type _ReviseDocumentRequest ReviseDocumentRequest

// NewReviseDocumentRequest instantiates a new ReviseDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviseDocumentRequest(source map[string]*interface{}) *ReviseDocumentRequest {
	this := ReviseDocumentRequest{}
	this.Source = source
	return &this
}

// NewReviseDocumentRequestWithDefaults instantiates a new ReviseDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviseDocumentRequestWithDefaults() *ReviseDocumentRequest {
	this := ReviseDocumentRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *ReviseDocumentRequest) GetSource() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ReviseDocumentRequest) GetSourceOk() (map[string]*interface{}, bool) {
	if o == nil {
		return map[string]*interface{}{}, false
	}
	return o.Source, true
}

// SetSource sets field value
func (o *ReviseDocumentRequest) SetSource(v map[string]*interface{}) {
	o.Source = v
}

// GetFacts returns the Facts field value if set, zero value otherwise.
func (o *ReviseDocumentRequest) GetFacts() map[string]string {
	if o == nil || IsNil(o.Facts) {
		var ret map[string]string
		return ret
	}
	return o.Facts
}

// GetFactsOk returns a tuple with the Facts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviseDocumentRequest) GetFactsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Facts) {
		return map[string]string{}, false
	}
	return o.Facts, true
}

// HasFacts returns a boolean if a field has been set.
func (o *ReviseDocumentRequest) HasFacts() bool {
	if o != nil && !IsNil(o.Facts) {
		return true
	}

	return false
}

// SetFacts gets a reference to the given map[string]string and assigns it to the Facts field.
func (o *ReviseDocumentRequest) SetFacts(v map[string]string) {
	o.Facts = v
}

// GetVariableValues returns the VariableValues field value if set, zero value otherwise.
func (o *ReviseDocumentRequest) GetVariableValues() map[string]string {
	if o == nil || IsNil(o.VariableValues) {
		var ret map[string]string
		return ret
	}
	return o.VariableValues
}

// GetVariableValuesOk returns a tuple with the VariableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviseDocumentRequest) GetVariableValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.VariableValues) {
		return map[string]string{}, false
	}
	return o.VariableValues, true
}

// HasVariableValues returns a boolean if a field has been set.
func (o *ReviseDocumentRequest) HasVariableValues() bool {
	if o != nil && !IsNil(o.VariableValues) {
		return true
	}

	return false
}

// SetVariableValues gets a reference to the given map[string]string and assigns it to the VariableValues field.
func (o *ReviseDocumentRequest) SetVariableValues(v map[string]string) {
	o.VariableValues = v
}

// GetRoleMap returns the RoleMap field value if set, zero value otherwise.
func (o *ReviseDocumentRequest) GetRoleMap() map[string]string {
	if o == nil || IsNil(o.RoleMap) {
		var ret map[string]string
		return ret
	}
	return o.RoleMap
}

// GetRoleMapOk returns a tuple with the RoleMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviseDocumentRequest) GetRoleMapOk() (map[string]string, bool) {
	if o == nil || IsNil(o.RoleMap) {
		return map[string]string{}, false
	}
	return o.RoleMap, true
}

// HasRoleMap returns a boolean if a field has been set.
func (o *ReviseDocumentRequest) HasRoleMap() bool {
	if o != nil && !IsNil(o.RoleMap) {
		return true
	}

	return false
}

// SetRoleMap gets a reference to the given map[string]string and assigns it to the RoleMap field.
func (o *ReviseDocumentRequest) SetRoleMap(v map[string]string) {
	o.RoleMap = v
}

func (o ReviseDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviseDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	if !IsNil(o.Facts) {
		toSerialize["facts"] = o.Facts
	}
	if !IsNil(o.VariableValues) {
		toSerialize["variableValues"] = o.VariableValues
	}
	if !IsNil(o.RoleMap) {
		toSerialize["roleMap"] = o.RoleMap
	}
	return toSerialize, nil
}

func (o *ReviseDocumentRequest) UnmarshalJSON(data []byte) (err error) {
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

	varReviseDocumentRequest := _ReviseDocumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReviseDocumentRequest)

	if err != nil {
		return err
	}

	*o = ReviseDocumentRequest(varReviseDocumentRequest)

	return err
}

type NullableReviseDocumentRequest struct {
	value *ReviseDocumentRequest
	isSet bool
}

func (v NullableReviseDocumentRequest) Get() *ReviseDocumentRequest {
	return v.value
}

func (v *NullableReviseDocumentRequest) Set(val *ReviseDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReviseDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReviseDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviseDocumentRequest(val *ReviseDocumentRequest) *NullableReviseDocumentRequest {
	return &NullableReviseDocumentRequest{value: val, isSet: true}
}

func (v NullableReviseDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviseDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the EditWithAiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditWithAiRequest{}

// EditWithAiRequest struct for EditWithAiRequest
type EditWithAiRequest struct {
	Request string `json:"request"`
	// A written document: TipTap JSON `{type:\"doc\",content:[…]}`. Field tags (`fieldTag` with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (`variableChip`; `custom.<key>` chips are the template's questions), and headings with `audience: \"sender\"` (Sender only: in nothing a signer sees, signs or receives).
	Source map[string]*interface{} `json:"source"`
	Mode *string `json:"mode,omitempty"`
	Variables []EditWithAiRequestVariablesInner `json:"variables,omitempty"`
}

type _EditWithAiRequest EditWithAiRequest

// NewEditWithAiRequest instantiates a new EditWithAiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditWithAiRequest(request string, source map[string]*interface{}) *EditWithAiRequest {
	this := EditWithAiRequest{}
	this.Request = request
	this.Source = source
	return &this
}

// NewEditWithAiRequestWithDefaults instantiates a new EditWithAiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditWithAiRequestWithDefaults() *EditWithAiRequest {
	this := EditWithAiRequest{}
	return &this
}

// GetRequest returns the Request field value
func (o *EditWithAiRequest) GetRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *EditWithAiRequest) GetRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *EditWithAiRequest) SetRequest(v string) {
	o.Request = v
}

// GetSource returns the Source field value
func (o *EditWithAiRequest) GetSource() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EditWithAiRequest) GetSourceOk() (map[string]*interface{}, bool) {
	if o == nil {
		return map[string]*interface{}{}, false
	}
	return o.Source, true
}

// SetSource sets field value
func (o *EditWithAiRequest) SetSource(v map[string]*interface{}) {
	o.Source = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *EditWithAiRequest) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditWithAiRequest) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *EditWithAiRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *EditWithAiRequest) SetMode(v string) {
	o.Mode = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *EditWithAiRequest) GetVariables() []EditWithAiRequestVariablesInner {
	if o == nil || IsNil(o.Variables) {
		var ret []EditWithAiRequestVariablesInner
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditWithAiRequest) GetVariablesOk() ([]EditWithAiRequestVariablesInner, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *EditWithAiRequest) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []EditWithAiRequestVariablesInner and assigns it to the Variables field.
func (o *EditWithAiRequest) SetVariables(v []EditWithAiRequestVariablesInner) {
	o.Variables = v
}

func (o EditWithAiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditWithAiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request"] = o.Request
	toSerialize["source"] = o.Source
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

func (o *EditWithAiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request",
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

	varEditWithAiRequest := _EditWithAiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditWithAiRequest)

	if err != nil {
		return err
	}

	*o = EditWithAiRequest(varEditWithAiRequest)

	return err
}

type NullableEditWithAiRequest struct {
	value *EditWithAiRequest
	isSet bool
}

func (v NullableEditWithAiRequest) Get() *EditWithAiRequest {
	return v.value
}

func (v *NullableEditWithAiRequest) Set(val *EditWithAiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditWithAiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditWithAiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditWithAiRequest(val *EditWithAiRequest) *NullableEditWithAiRequest {
	return &NullableEditWithAiRequest{value: val, isSet: true}
}

func (v NullableEditWithAiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditWithAiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



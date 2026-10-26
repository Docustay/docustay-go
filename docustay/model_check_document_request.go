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

// checks if the CheckDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckDocumentRequest{}

// CheckDocumentRequest struct for CheckDocumentRequest
type CheckDocumentRequest struct {
	// A written document: TipTap JSON `{type:\"doc\",content:[…]}`. Field tags (`fieldTag` with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (`variableChip`; `custom.<key>` chips are the template's questions), and headings with `audience: \"sender\"` (Sender only: in nothing a signer sees, signs or receives).
	Source map[string]*interface{} `json:"source"`
	Facts map[string]string `json:"facts,omitempty"`
	// With people, each person's own merge text is known, as when the document is made.
	People []Person `json:"people,omitempty"`
	VariableValues map[string]string `json:"variableValues,omitempty"`
	RequireSignature *bool `json:"requireSignature,omitempty"`
	ForTemplate *bool `json:"forTemplate,omitempty"`
}

type _CheckDocumentRequest CheckDocumentRequest

// NewCheckDocumentRequest instantiates a new CheckDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckDocumentRequest(source map[string]*interface{}) *CheckDocumentRequest {
	this := CheckDocumentRequest{}
	this.Source = source
	return &this
}

// NewCheckDocumentRequestWithDefaults instantiates a new CheckDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckDocumentRequestWithDefaults() *CheckDocumentRequest {
	this := CheckDocumentRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *CheckDocumentRequest) GetSource() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CheckDocumentRequest) GetSourceOk() (map[string]*interface{}, bool) {
	if o == nil {
		return map[string]*interface{}{}, false
	}
	return o.Source, true
}

// SetSource sets field value
func (o *CheckDocumentRequest) SetSource(v map[string]*interface{}) {
	o.Source = v
}

// GetFacts returns the Facts field value if set, zero value otherwise.
func (o *CheckDocumentRequest) GetFacts() map[string]string {
	if o == nil || IsNil(o.Facts) {
		var ret map[string]string
		return ret
	}
	return o.Facts
}

// GetFactsOk returns a tuple with the Facts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDocumentRequest) GetFactsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Facts) {
		return map[string]string{}, false
	}
	return o.Facts, true
}

// HasFacts returns a boolean if a field has been set.
func (o *CheckDocumentRequest) HasFacts() bool {
	if o != nil && !IsNil(o.Facts) {
		return true
	}

	return false
}

// SetFacts gets a reference to the given map[string]string and assigns it to the Facts field.
func (o *CheckDocumentRequest) SetFacts(v map[string]string) {
	o.Facts = v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *CheckDocumentRequest) GetPeople() []Person {
	if o == nil || IsNil(o.People) {
		var ret []Person
		return ret
	}
	return o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDocumentRequest) GetPeopleOk() ([]Person, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *CheckDocumentRequest) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given []Person and assigns it to the People field.
func (o *CheckDocumentRequest) SetPeople(v []Person) {
	o.People = v
}

// GetVariableValues returns the VariableValues field value if set, zero value otherwise.
func (o *CheckDocumentRequest) GetVariableValues() map[string]string {
	if o == nil || IsNil(o.VariableValues) {
		var ret map[string]string
		return ret
	}
	return o.VariableValues
}

// GetVariableValuesOk returns a tuple with the VariableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDocumentRequest) GetVariableValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.VariableValues) {
		return map[string]string{}, false
	}
	return o.VariableValues, true
}

// HasVariableValues returns a boolean if a field has been set.
func (o *CheckDocumentRequest) HasVariableValues() bool {
	if o != nil && !IsNil(o.VariableValues) {
		return true
	}

	return false
}

// SetVariableValues gets a reference to the given map[string]string and assigns it to the VariableValues field.
func (o *CheckDocumentRequest) SetVariableValues(v map[string]string) {
	o.VariableValues = v
}

// GetRequireSignature returns the RequireSignature field value if set, zero value otherwise.
func (o *CheckDocumentRequest) GetRequireSignature() bool {
	if o == nil || IsNil(o.RequireSignature) {
		var ret bool
		return ret
	}
	return *o.RequireSignature
}

// GetRequireSignatureOk returns a tuple with the RequireSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDocumentRequest) GetRequireSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSignature) {
		return nil, false
	}
	return o.RequireSignature, true
}

// HasRequireSignature returns a boolean if a field has been set.
func (o *CheckDocumentRequest) HasRequireSignature() bool {
	if o != nil && !IsNil(o.RequireSignature) {
		return true
	}

	return false
}

// SetRequireSignature gets a reference to the given bool and assigns it to the RequireSignature field.
func (o *CheckDocumentRequest) SetRequireSignature(v bool) {
	o.RequireSignature = &v
}

// GetForTemplate returns the ForTemplate field value if set, zero value otherwise.
func (o *CheckDocumentRequest) GetForTemplate() bool {
	if o == nil || IsNil(o.ForTemplate) {
		var ret bool
		return ret
	}
	return *o.ForTemplate
}

// GetForTemplateOk returns a tuple with the ForTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDocumentRequest) GetForTemplateOk() (*bool, bool) {
	if o == nil || IsNil(o.ForTemplate) {
		return nil, false
	}
	return o.ForTemplate, true
}

// HasForTemplate returns a boolean if a field has been set.
func (o *CheckDocumentRequest) HasForTemplate() bool {
	if o != nil && !IsNil(o.ForTemplate) {
		return true
	}

	return false
}

// SetForTemplate gets a reference to the given bool and assigns it to the ForTemplate field.
func (o *CheckDocumentRequest) SetForTemplate(v bool) {
	o.ForTemplate = &v
}

func (o CheckDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	if !IsNil(o.Facts) {
		toSerialize["facts"] = o.Facts
	}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}
	if !IsNil(o.VariableValues) {
		toSerialize["variableValues"] = o.VariableValues
	}
	if !IsNil(o.RequireSignature) {
		toSerialize["requireSignature"] = o.RequireSignature
	}
	if !IsNil(o.ForTemplate) {
		toSerialize["forTemplate"] = o.ForTemplate
	}
	return toSerialize, nil
}

func (o *CheckDocumentRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCheckDocumentRequest := _CheckDocumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCheckDocumentRequest)

	if err != nil {
		return err
	}

	*o = CheckDocumentRequest(varCheckDocumentRequest)

	return err
}

type NullableCheckDocumentRequest struct {
	value *CheckDocumentRequest
	isSet bool
}

func (v NullableCheckDocumentRequest) Get() *CheckDocumentRequest {
	return v.value
}

func (v *NullableCheckDocumentRequest) Set(val *CheckDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckDocumentRequest(val *CheckDocumentRequest) *NullableCheckDocumentRequest {
	return &NullableCheckDocumentRequest{value: val, isSet: true}
}

func (v NullableCheckDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



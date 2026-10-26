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

// checks if the WizardBrief type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WizardBrief{}

// WizardBrief The wizard's answers (see GET /api/v1/ai/wizard for the questions).
type WizardBrief struct {
	Type string `json:"type"`
	Kind *string `json:"kind,omitempty"`
	Answers map[string]string `json:"answers,omitempty"`
	Final string `json:"final"`
	Payment *WizardBriefPayment `json:"payment,omitempty"`
}

type _WizardBrief WizardBrief

// NewWizardBrief instantiates a new WizardBrief object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWizardBrief(type_ string, final string) *WizardBrief {
	this := WizardBrief{}
	this.Type = type_
	this.Final = final
	return &this
}

// NewWizardBriefWithDefaults instantiates a new WizardBrief object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWizardBriefWithDefaults() *WizardBrief {
	this := WizardBrief{}
	return &this
}

// GetType returns the Type field value
func (o *WizardBrief) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WizardBrief) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WizardBrief) SetType(v string) {
	o.Type = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *WizardBrief) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardBrief) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *WizardBrief) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *WizardBrief) SetKind(v string) {
	o.Kind = &v
}

// GetAnswers returns the Answers field value if set, zero value otherwise.
func (o *WizardBrief) GetAnswers() map[string]string {
	if o == nil || IsNil(o.Answers) {
		var ret map[string]string
		return ret
	}
	return o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardBrief) GetAnswersOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Answers) {
		return map[string]string{}, false
	}
	return o.Answers, true
}

// HasAnswers returns a boolean if a field has been set.
func (o *WizardBrief) HasAnswers() bool {
	if o != nil && !IsNil(o.Answers) {
		return true
	}

	return false
}

// SetAnswers gets a reference to the given map[string]string and assigns it to the Answers field.
func (o *WizardBrief) SetAnswers(v map[string]string) {
	o.Answers = v
}

// GetFinal returns the Final field value
func (o *WizardBrief) GetFinal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Final
}

// GetFinalOk returns a tuple with the Final field value
// and a boolean to check if the value has been set.
func (o *WizardBrief) GetFinalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Final, true
}

// SetFinal sets field value
func (o *WizardBrief) SetFinal(v string) {
	o.Final = v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *WizardBrief) GetPayment() WizardBriefPayment {
	if o == nil || IsNil(o.Payment) {
		var ret WizardBriefPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardBrief) GetPaymentOk() (*WizardBriefPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *WizardBrief) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given WizardBriefPayment and assigns it to the Payment field.
func (o *WizardBrief) SetPayment(v WizardBriefPayment) {
	o.Payment = &v
}

func (o WizardBrief) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WizardBrief) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Answers) {
		toSerialize["answers"] = o.Answers
	}
	toSerialize["final"] = o.Final
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	return toSerialize, nil
}

func (o *WizardBrief) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"final",
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

	varWizardBrief := _WizardBrief{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWizardBrief)

	if err != nil {
		return err
	}

	*o = WizardBrief(varWizardBrief)

	return err
}

type NullableWizardBrief struct {
	value *WizardBrief
	isSet bool
}

func (v NullableWizardBrief) Get() *WizardBrief {
	return v.value
}

func (v *NullableWizardBrief) Set(val *WizardBrief) {
	v.value = val
	v.isSet = true
}

func (v NullableWizardBrief) IsSet() bool {
	return v.isSet
}

func (v *NullableWizardBrief) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWizardBrief(val *WizardBrief) *NullableWizardBrief {
	return &NullableWizardBrief{value: val, isSet: true}
}

func (v NullableWizardBrief) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWizardBrief) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



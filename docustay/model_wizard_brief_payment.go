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

// checks if the WizardBriefPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WizardBriefPayment{}

// WizardBriefPayment A payment collected at signing. The worker (plain code) places the Payment field with exactly this amount; later payments are written as text.
type WizardBriefPayment struct {
	Collect bool `json:"collect"`
	// In major units, typed by the person (for example \"1500\"). The assistant never decides an amount.
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	PayerRole *string `json:"payer_role,omitempty"`
	Description string `json:"description"`
	ScheduleNote *string `json:"schedule_note,omitempty"`
}

type _WizardBriefPayment WizardBriefPayment

// NewWizardBriefPayment instantiates a new WizardBriefPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWizardBriefPayment(collect bool, amount string, currency string, description string) *WizardBriefPayment {
	this := WizardBriefPayment{}
	this.Collect = collect
	this.Amount = amount
	this.Currency = currency
	this.Description = description
	return &this
}

// NewWizardBriefPaymentWithDefaults instantiates a new WizardBriefPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWizardBriefPaymentWithDefaults() *WizardBriefPayment {
	this := WizardBriefPayment{}
	return &this
}

// GetCollect returns the Collect field value
func (o *WizardBriefPayment) GetCollect() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Collect
}

// GetCollectOk returns a tuple with the Collect field value
// and a boolean to check if the value has been set.
func (o *WizardBriefPayment) GetCollectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collect, true
}

// SetCollect sets field value
func (o *WizardBriefPayment) SetCollect(v bool) {
	o.Collect = v
}

// GetAmount returns the Amount field value
func (o *WizardBriefPayment) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *WizardBriefPayment) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *WizardBriefPayment) SetAmount(v string) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *WizardBriefPayment) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *WizardBriefPayment) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *WizardBriefPayment) SetCurrency(v string) {
	o.Currency = v
}

// GetPayerRole returns the PayerRole field value if set, zero value otherwise.
func (o *WizardBriefPayment) GetPayerRole() string {
	if o == nil || IsNil(o.PayerRole) {
		var ret string
		return ret
	}
	return *o.PayerRole
}

// GetPayerRoleOk returns a tuple with the PayerRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardBriefPayment) GetPayerRoleOk() (*string, bool) {
	if o == nil || IsNil(o.PayerRole) {
		return nil, false
	}
	return o.PayerRole, true
}

// HasPayerRole returns a boolean if a field has been set.
func (o *WizardBriefPayment) HasPayerRole() bool {
	if o != nil && !IsNil(o.PayerRole) {
		return true
	}

	return false
}

// SetPayerRole gets a reference to the given string and assigns it to the PayerRole field.
func (o *WizardBriefPayment) SetPayerRole(v string) {
	o.PayerRole = &v
}

// GetDescription returns the Description field value
func (o *WizardBriefPayment) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WizardBriefPayment) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WizardBriefPayment) SetDescription(v string) {
	o.Description = v
}

// GetScheduleNote returns the ScheduleNote field value if set, zero value otherwise.
func (o *WizardBriefPayment) GetScheduleNote() string {
	if o == nil || IsNil(o.ScheduleNote) {
		var ret string
		return ret
	}
	return *o.ScheduleNote
}

// GetScheduleNoteOk returns a tuple with the ScheduleNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardBriefPayment) GetScheduleNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleNote) {
		return nil, false
	}
	return o.ScheduleNote, true
}

// HasScheduleNote returns a boolean if a field has been set.
func (o *WizardBriefPayment) HasScheduleNote() bool {
	if o != nil && !IsNil(o.ScheduleNote) {
		return true
	}

	return false
}

// SetScheduleNote gets a reference to the given string and assigns it to the ScheduleNote field.
func (o *WizardBriefPayment) SetScheduleNote(v string) {
	o.ScheduleNote = &v
}

func (o WizardBriefPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WizardBriefPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collect"] = o.Collect
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	if !IsNil(o.PayerRole) {
		toSerialize["payer_role"] = o.PayerRole
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.ScheduleNote) {
		toSerialize["schedule_note"] = o.ScheduleNote
	}
	return toSerialize, nil
}

func (o *WizardBriefPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"collect",
		"amount",
		"currency",
		"description",
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

	varWizardBriefPayment := _WizardBriefPayment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWizardBriefPayment)

	if err != nil {
		return err
	}

	*o = WizardBriefPayment(varWizardBriefPayment)

	return err
}

type NullableWizardBriefPayment struct {
	value *WizardBriefPayment
	isSet bool
}

func (v NullableWizardBriefPayment) Get() *WizardBriefPayment {
	return v.value
}

func (v *NullableWizardBriefPayment) Set(val *WizardBriefPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableWizardBriefPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableWizardBriefPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWizardBriefPayment(val *WizardBriefPayment) *NullableWizardBriefPayment {
	return &NullableWizardBriefPayment{value: val, isSet: true}
}

func (v NullableWizardBriefPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWizardBriefPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



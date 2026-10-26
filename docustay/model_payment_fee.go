/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
	"fmt"
)

// checks if the PaymentFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentFee{}

// PaymentFee Docustay's platform fee on this payment, on top of Stripe's own fees (which Stripe charges to your own Stripe account).
type PaymentFee struct {
	// Docustay's fee, decided when the payment started: 200 (2%) on the Free plan, 0 on Pro.
	BasisPoints int32 `json:"basisPoints"`
	AmountMinor int32 `json:"amountMinor"`
	// How much of the fee went back with refunds.
	RefundedMinor int32 `json:"refundedMinor"`
	Plan NullableString `json:"plan"`
	AdditionalProperties map[string]interface{}
}

type _PaymentFee PaymentFee

// NewPaymentFee instantiates a new PaymentFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentFee(basisPoints int32, amountMinor int32, refundedMinor int32, plan NullableString) *PaymentFee {
	this := PaymentFee{}
	this.BasisPoints = basisPoints
	this.AmountMinor = amountMinor
	this.RefundedMinor = refundedMinor
	this.Plan = plan
	return &this
}

// NewPaymentFeeWithDefaults instantiates a new PaymentFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentFeeWithDefaults() *PaymentFee {
	this := PaymentFee{}
	return &this
}

// GetBasisPoints returns the BasisPoints field value
func (o *PaymentFee) GetBasisPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BasisPoints
}

// GetBasisPointsOk returns a tuple with the BasisPoints field value
// and a boolean to check if the value has been set.
func (o *PaymentFee) GetBasisPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasisPoints, true
}

// SetBasisPoints sets field value
func (o *PaymentFee) SetBasisPoints(v int32) {
	o.BasisPoints = v
}

// GetAmountMinor returns the AmountMinor field value
func (o *PaymentFee) GetAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *PaymentFee) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *PaymentFee) SetAmountMinor(v int32) {
	o.AmountMinor = v
}

// GetRefundedMinor returns the RefundedMinor field value
func (o *PaymentFee) GetRefundedMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RefundedMinor
}

// GetRefundedMinorOk returns a tuple with the RefundedMinor field value
// and a boolean to check if the value has been set.
func (o *PaymentFee) GetRefundedMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundedMinor, true
}

// SetRefundedMinor sets field value
func (o *PaymentFee) SetRefundedMinor(v int32) {
	o.RefundedMinor = v
}

// GetPlan returns the Plan field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentFee) GetPlan() string {
	if o == nil || o.Plan.Get() == nil {
		var ret string
		return ret
	}

	return *o.Plan.Get()
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentFee) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plan.Get(), o.Plan.IsSet()
}

// SetPlan sets field value
func (o *PaymentFee) SetPlan(v string) {
	o.Plan.Set(&v)
}

func (o PaymentFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["basisPoints"] = o.BasisPoints
	toSerialize["amountMinor"] = o.AmountMinor
	toSerialize["refundedMinor"] = o.RefundedMinor
	toSerialize["plan"] = o.Plan.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"basisPoints",
		"amountMinor",
		"refundedMinor",
		"plan",
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

	varPaymentFee := _PaymentFee{}

	err = json.Unmarshal(data, &varPaymentFee)

	if err != nil {
		return err
	}

	*o = PaymentFee(varPaymentFee)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "basisPoints")
		delete(additionalProperties, "amountMinor")
		delete(additionalProperties, "refundedMinor")
		delete(additionalProperties, "plan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentFee struct {
	value *PaymentFee
	isSet bool
}

func (v NullablePaymentFee) Get() *PaymentFee {
	return v.value
}

func (v *NullablePaymentFee) Set(val *PaymentFee) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentFee) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentFee(val *PaymentFee) *NullablePaymentFee {
	return &NullablePaymentFee{value: val, isSet: true}
}

func (v NullablePaymentFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



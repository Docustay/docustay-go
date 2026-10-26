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

// checks if the Payment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Payment{}

// Payment struct for Payment
type Payment struct {
	Id string `json:"id"`
	FieldId NullableString `json:"fieldId"`
	Person string `json:"person"`
	// `paid` once Stripe says so. A refund moves it to `partially_refunded` or `refunded`.
	Status string `json:"status"`
	// What was charged, in the smallest unit of `currency` (cents). Worked out on the server; never taken from the browser.
	AmountMinor int32 `json:"amountMinor"`
	Currency string `json:"currency"`
	PaidAt NullableString `json:"paidAt"`
	// Stripe's own receipt for the payment.
	ReceiptUrl NullableString `json:"receiptUrl"`
	RefundedMinor int32 `json:"refundedMinor"`
	Fee PaymentFee `json:"fee"`
	Test bool `json:"test"`
	AdditionalProperties map[string]interface{}
}

type _Payment Payment

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment(id string, fieldId NullableString, person string, status string, amountMinor int32, currency string, paidAt NullableString, receiptUrl NullableString, refundedMinor int32, fee PaymentFee, test bool) *Payment {
	this := Payment{}
	this.Id = id
	this.FieldId = fieldId
	this.Person = person
	this.Status = status
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.PaidAt = paidAt
	this.ReceiptUrl = receiptUrl
	this.RefundedMinor = refundedMinor
	this.Fee = fee
	this.Test = test
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	return &this
}

// GetId returns the Id field value
func (o *Payment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Payment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Payment) SetId(v string) {
	o.Id = v
}

// GetFieldId returns the FieldId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payment) GetFieldId() string {
	if o == nil || o.FieldId.Get() == nil {
		var ret string
		return ret
	}

	return *o.FieldId.Get()
}

// GetFieldIdOk returns a tuple with the FieldId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetFieldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldId.Get(), o.FieldId.IsSet()
}

// SetFieldId sets field value
func (o *Payment) SetFieldId(v string) {
	o.FieldId.Set(&v)
}

// GetPerson returns the Person field value
func (o *Payment) GetPerson() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Person
}

// GetPersonOk returns a tuple with the Person field value
// and a boolean to check if the value has been set.
func (o *Payment) GetPersonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Person, true
}

// SetPerson sets field value
func (o *Payment) SetPerson(v string) {
	o.Person = v
}

// GetStatus returns the Status field value
func (o *Payment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Payment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Payment) SetStatus(v string) {
	o.Status = v
}

// GetAmountMinor returns the AmountMinor field value
func (o *Payment) GetAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *Payment) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *Payment) SetAmountMinor(v int32) {
	o.AmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *Payment) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Payment) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Payment) SetCurrency(v string) {
	o.Currency = v
}

// GetPaidAt returns the PaidAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payment) GetPaidAt() string {
	if o == nil || o.PaidAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaidAt.Get()
}

// GetPaidAtOk returns a tuple with the PaidAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetPaidAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidAt.Get(), o.PaidAt.IsSet()
}

// SetPaidAt sets field value
func (o *Payment) SetPaidAt(v string) {
	o.PaidAt.Set(&v)
}

// GetReceiptUrl returns the ReceiptUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payment) GetReceiptUrl() string {
	if o == nil || o.ReceiptUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReceiptUrl.Get()
}

// GetReceiptUrlOk returns a tuple with the ReceiptUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetReceiptUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiptUrl.Get(), o.ReceiptUrl.IsSet()
}

// SetReceiptUrl sets field value
func (o *Payment) SetReceiptUrl(v string) {
	o.ReceiptUrl.Set(&v)
}

// GetRefundedMinor returns the RefundedMinor field value
func (o *Payment) GetRefundedMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RefundedMinor
}

// GetRefundedMinorOk returns a tuple with the RefundedMinor field value
// and a boolean to check if the value has been set.
func (o *Payment) GetRefundedMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundedMinor, true
}

// SetRefundedMinor sets field value
func (o *Payment) SetRefundedMinor(v int32) {
	o.RefundedMinor = v
}

// GetFee returns the Fee field value
func (o *Payment) GetFee() PaymentFee {
	if o == nil {
		var ret PaymentFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *Payment) GetFeeOk() (*PaymentFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *Payment) SetFee(v PaymentFee) {
	o.Fee = v
}

// GetTest returns the Test field value
func (o *Payment) GetTest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Test
}

// GetTestOk returns a tuple with the Test field value
// and a boolean to check if the value has been set.
func (o *Payment) GetTestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Test, true
}

// SetTest sets field value
func (o *Payment) SetTest(v bool) {
	o.Test = v
}

func (o Payment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fieldId"] = o.FieldId.Get()
	toSerialize["person"] = o.Person
	toSerialize["status"] = o.Status
	toSerialize["amountMinor"] = o.AmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["paidAt"] = o.PaidAt.Get()
	toSerialize["receiptUrl"] = o.ReceiptUrl.Get()
	toSerialize["refundedMinor"] = o.RefundedMinor
	toSerialize["fee"] = o.Fee
	toSerialize["test"] = o.Test

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Payment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"fieldId",
		"person",
		"status",
		"amountMinor",
		"currency",
		"paidAt",
		"receiptUrl",
		"refundedMinor",
		"fee",
		"test",
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

	varPayment := _Payment{}

	err = json.Unmarshal(data, &varPayment)

	if err != nil {
		return err
	}

	*o = Payment(varPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "fieldId")
		delete(additionalProperties, "person")
		delete(additionalProperties, "status")
		delete(additionalProperties, "amountMinor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "paidAt")
		delete(additionalProperties, "receiptUrl")
		delete(additionalProperties, "refundedMinor")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "test")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the RefundPaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefundPaymentRequest{}

// RefundPaymentRequest struct for RefundPaymentRequest
type RefundPaymentRequest struct {
	AmountMinor *int32 `json:"amountMinor,omitempty"`
	Reason string `json:"reason"`
}

type _RefundPaymentRequest RefundPaymentRequest

// NewRefundPaymentRequest instantiates a new RefundPaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundPaymentRequest(reason string) *RefundPaymentRequest {
	this := RefundPaymentRequest{}
	this.Reason = reason
	return &this
}

// NewRefundPaymentRequestWithDefaults instantiates a new RefundPaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundPaymentRequestWithDefaults() *RefundPaymentRequest {
	this := RefundPaymentRequest{}
	return &this
}

// GetAmountMinor returns the AmountMinor field value if set, zero value otherwise.
func (o *RefundPaymentRequest) GetAmountMinor() int32 {
	if o == nil || IsNil(o.AmountMinor) {
		var ret int32
		return ret
	}
	return *o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundPaymentRequest) GetAmountMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountMinor) {
		return nil, false
	}
	return o.AmountMinor, true
}

// HasAmountMinor returns a boolean if a field has been set.
func (o *RefundPaymentRequest) HasAmountMinor() bool {
	if o != nil && !IsNil(o.AmountMinor) {
		return true
	}

	return false
}

// SetAmountMinor gets a reference to the given int32 and assigns it to the AmountMinor field.
func (o *RefundPaymentRequest) SetAmountMinor(v int32) {
	o.AmountMinor = &v
}

// GetReason returns the Reason field value
func (o *RefundPaymentRequest) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *RefundPaymentRequest) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *RefundPaymentRequest) SetReason(v string) {
	o.Reason = v
}

func (o RefundPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundPaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmountMinor) {
		toSerialize["amountMinor"] = o.AmountMinor
	}
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *RefundPaymentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reason",
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

	varRefundPaymentRequest := _RefundPaymentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefundPaymentRequest)

	if err != nil {
		return err
	}

	*o = RefundPaymentRequest(varRefundPaymentRequest)

	return err
}

type NullableRefundPaymentRequest struct {
	value *RefundPaymentRequest
	isSet bool
}

func (v NullableRefundPaymentRequest) Get() *RefundPaymentRequest {
	return v.value
}

func (v *NullableRefundPaymentRequest) Set(val *RefundPaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundPaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundPaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundPaymentRequest(val *RefundPaymentRequest) *NullableRefundPaymentRequest {
	return &NullableRefundPaymentRequest{value: val, isSet: true}
}

func (v NullableRefundPaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundPaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



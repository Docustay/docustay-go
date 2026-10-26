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

// checks if the UsageBilling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageBilling{}

// UsageBilling Usage that costs money on Pay as you go, Pro and Lifetime: what has been billed per unit this period, the budget (default $50), and whether sending is paused at 100% of it. Amounts are whole cents.
type UsageBilling struct {
	PerDocumentMinor NullableInt32 `json:"perDocumentMinor"`
	PerDraftMinor NullableInt32 `json:"perDraftMinor"`
	BillableDocuments int32 `json:"billableDocuments"`
	BillableDrafts int32 `json:"billableDrafts"`
	UnbilledMinor int32 `json:"unbilledMinor"`
	BudgetMinor int32 `json:"budgetMinor"`
	PercentOfBudget int32 `json:"percentOfBudget"`
	Paused bool `json:"paused"`
	Currency string `json:"currency"`
}

type _UsageBilling UsageBilling

// NewUsageBilling instantiates a new UsageBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageBilling(perDocumentMinor NullableInt32, perDraftMinor NullableInt32, billableDocuments int32, billableDrafts int32, unbilledMinor int32, budgetMinor int32, percentOfBudget int32, paused bool, currency string) *UsageBilling {
	this := UsageBilling{}
	this.PerDocumentMinor = perDocumentMinor
	this.PerDraftMinor = perDraftMinor
	this.BillableDocuments = billableDocuments
	this.BillableDrafts = billableDrafts
	this.UnbilledMinor = unbilledMinor
	this.BudgetMinor = budgetMinor
	this.PercentOfBudget = percentOfBudget
	this.Paused = paused
	this.Currency = currency
	return &this
}

// NewUsageBillingWithDefaults instantiates a new UsageBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageBillingWithDefaults() *UsageBilling {
	this := UsageBilling{}
	return &this
}

// GetPerDocumentMinor returns the PerDocumentMinor field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageBilling) GetPerDocumentMinor() int32 {
	if o == nil || o.PerDocumentMinor.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PerDocumentMinor.Get()
}

// GetPerDocumentMinorOk returns a tuple with the PerDocumentMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageBilling) GetPerDocumentMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerDocumentMinor.Get(), o.PerDocumentMinor.IsSet()
}

// SetPerDocumentMinor sets field value
func (o *UsageBilling) SetPerDocumentMinor(v int32) {
	o.PerDocumentMinor.Set(&v)
}

// GetPerDraftMinor returns the PerDraftMinor field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageBilling) GetPerDraftMinor() int32 {
	if o == nil || o.PerDraftMinor.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PerDraftMinor.Get()
}

// GetPerDraftMinorOk returns a tuple with the PerDraftMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageBilling) GetPerDraftMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerDraftMinor.Get(), o.PerDraftMinor.IsSet()
}

// SetPerDraftMinor sets field value
func (o *UsageBilling) SetPerDraftMinor(v int32) {
	o.PerDraftMinor.Set(&v)
}

// GetBillableDocuments returns the BillableDocuments field value
func (o *UsageBilling) GetBillableDocuments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BillableDocuments
}

// GetBillableDocumentsOk returns a tuple with the BillableDocuments field value
// and a boolean to check if the value has been set.
func (o *UsageBilling) GetBillableDocumentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillableDocuments, true
}

// SetBillableDocuments sets field value
func (o *UsageBilling) SetBillableDocuments(v int32) {
	o.BillableDocuments = v
}

// GetBillableDrafts returns the BillableDrafts field value
func (o *UsageBilling) GetBillableDrafts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BillableDrafts
}

// GetBillableDraftsOk returns a tuple with the BillableDrafts field value
// and a boolean to check if the value has been set.
func (o *UsageBilling) GetBillableDraftsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillableDrafts, true
}

// SetBillableDrafts sets field value
func (o *UsageBilling) SetBillableDrafts(v int32) {
	o.BillableDrafts = v
}

// GetUnbilledMinor returns the UnbilledMinor field value
func (o *UsageBilling) GetUnbilledMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnbilledMinor
}

// GetUnbilledMinorOk returns a tuple with the UnbilledMinor field value
// and a boolean to check if the value has been set.
func (o *UsageBilling) GetUnbilledMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnbilledMinor, true
}

// SetUnbilledMinor sets field value
func (o *UsageBilling) SetUnbilledMinor(v int32) {
	o.UnbilledMinor = v
}

// GetBudgetMinor returns the BudgetMinor field value
func (o *UsageBilling) GetBudgetMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BudgetMinor
}

// GetBudgetMinorOk returns a tuple with the BudgetMinor field value
// and a boolean to check if the value has been set.
func (o *UsageBilling) GetBudgetMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetMinor, true
}

// SetBudgetMinor sets field value
func (o *UsageBilling) SetBudgetMinor(v int32) {
	o.BudgetMinor = v
}

// GetPercentOfBudget returns the PercentOfBudget field value
func (o *UsageBilling) GetPercentOfBudget() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PercentOfBudget
}

// GetPercentOfBudgetOk returns a tuple with the PercentOfBudget field value
// and a boolean to check if the value has been set.
func (o *UsageBilling) GetPercentOfBudgetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentOfBudget, true
}

// SetPercentOfBudget sets field value
func (o *UsageBilling) SetPercentOfBudget(v int32) {
	o.PercentOfBudget = v
}

// GetPaused returns the Paused field value
func (o *UsageBilling) GetPaused() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Paused
}

// GetPausedOk returns a tuple with the Paused field value
// and a boolean to check if the value has been set.
func (o *UsageBilling) GetPausedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Paused, true
}

// SetPaused sets field value
func (o *UsageBilling) SetPaused(v bool) {
	o.Paused = v
}

// GetCurrency returns the Currency field value
func (o *UsageBilling) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UsageBilling) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UsageBilling) SetCurrency(v string) {
	o.Currency = v
}

func (o UsageBilling) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageBilling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["perDocumentMinor"] = o.PerDocumentMinor.Get()
	toSerialize["perDraftMinor"] = o.PerDraftMinor.Get()
	toSerialize["billableDocuments"] = o.BillableDocuments
	toSerialize["billableDrafts"] = o.BillableDrafts
	toSerialize["unbilledMinor"] = o.UnbilledMinor
	toSerialize["budgetMinor"] = o.BudgetMinor
	toSerialize["percentOfBudget"] = o.PercentOfBudget
	toSerialize["paused"] = o.Paused
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *UsageBilling) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"perDocumentMinor",
		"perDraftMinor",
		"billableDocuments",
		"billableDrafts",
		"unbilledMinor",
		"budgetMinor",
		"percentOfBudget",
		"paused",
		"currency",
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

	varUsageBilling := _UsageBilling{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageBilling)

	if err != nil {
		return err
	}

	*o = UsageBilling(varUsageBilling)

	return err
}

type NullableUsageBilling struct {
	value *UsageBilling
	isSet bool
}

func (v NullableUsageBilling) Get() *UsageBilling {
	return v.value
}

func (v *NullableUsageBilling) Set(val *UsageBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageBilling(val *UsageBilling) *NullableUsageBilling {
	return &NullableUsageBilling{value: val, isSet: true}
}

func (v NullableUsageBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



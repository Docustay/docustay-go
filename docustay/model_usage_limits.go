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

// checks if the UsageLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageLimits{}

// UsageLimits The plan's limits and prices in one place (whole cents; null = no limit or no price): documents, AI drafts, new templates and different signers per period, what each extra document and draft costs, team members, the Pay as you go price per team member a month, the largest upload in MB, emails a day, the most rows in one list on Free, and the fee on payments in percent.
type UsageLimits struct {
	DocumentsPerPeriod NullableInt32 `json:"documentsPerPeriod"`
	PerDocumentMinor NullableInt32 `json:"perDocumentMinor"`
	DraftsPerPeriod NullableInt32 `json:"draftsPerPeriod"`
	PerDraftMinor NullableInt32 `json:"perDraftMinor"`
	TemplatesPerPeriod NullableInt32 `json:"templatesPerPeriod"`
	SignersPerPeriod NullableInt32 `json:"signersPerPeriod"`
	People NullableInt32 `json:"people"`
	PersonMonthlyMinor NullableInt32 `json:"personMonthlyMinor"`
	UploadMb NullableInt32 `json:"uploadMb"`
	// The plan's daily email limit. Today's actual limit is emails.limitPerDay, which is lower in a workspace's first 7 days.
	EmailsPerDay NullableInt32 `json:"emailsPerDay"`
	BulkSendMax NullableInt32 `json:"bulkSendMax"`
	PaymentsFeePercent float32 `json:"paymentsFeePercent"`
}

type _UsageLimits UsageLimits

// NewUsageLimits instantiates a new UsageLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLimits(documentsPerPeriod NullableInt32, perDocumentMinor NullableInt32, draftsPerPeriod NullableInt32, perDraftMinor NullableInt32, templatesPerPeriod NullableInt32, signersPerPeriod NullableInt32, people NullableInt32, personMonthlyMinor NullableInt32, uploadMb NullableInt32, emailsPerDay NullableInt32, bulkSendMax NullableInt32, paymentsFeePercent float32) *UsageLimits {
	this := UsageLimits{}
	this.DocumentsPerPeriod = documentsPerPeriod
	this.PerDocumentMinor = perDocumentMinor
	this.DraftsPerPeriod = draftsPerPeriod
	this.PerDraftMinor = perDraftMinor
	this.TemplatesPerPeriod = templatesPerPeriod
	this.SignersPerPeriod = signersPerPeriod
	this.People = people
	this.PersonMonthlyMinor = personMonthlyMinor
	this.UploadMb = uploadMb
	this.EmailsPerDay = emailsPerDay
	this.BulkSendMax = bulkSendMax
	this.PaymentsFeePercent = paymentsFeePercent
	return &this
}

// NewUsageLimitsWithDefaults instantiates a new UsageLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageLimitsWithDefaults() *UsageLimits {
	this := UsageLimits{}
	return &this
}

// GetDocumentsPerPeriod returns the DocumentsPerPeriod field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetDocumentsPerPeriod() int32 {
	if o == nil || o.DocumentsPerPeriod.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DocumentsPerPeriod.Get()
}

// GetDocumentsPerPeriodOk returns a tuple with the DocumentsPerPeriod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetDocumentsPerPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentsPerPeriod.Get(), o.DocumentsPerPeriod.IsSet()
}

// SetDocumentsPerPeriod sets field value
func (o *UsageLimits) SetDocumentsPerPeriod(v int32) {
	o.DocumentsPerPeriod.Set(&v)
}

// GetPerDocumentMinor returns the PerDocumentMinor field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetPerDocumentMinor() int32 {
	if o == nil || o.PerDocumentMinor.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PerDocumentMinor.Get()
}

// GetPerDocumentMinorOk returns a tuple with the PerDocumentMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetPerDocumentMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerDocumentMinor.Get(), o.PerDocumentMinor.IsSet()
}

// SetPerDocumentMinor sets field value
func (o *UsageLimits) SetPerDocumentMinor(v int32) {
	o.PerDocumentMinor.Set(&v)
}

// GetDraftsPerPeriod returns the DraftsPerPeriod field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetDraftsPerPeriod() int32 {
	if o == nil || o.DraftsPerPeriod.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DraftsPerPeriod.Get()
}

// GetDraftsPerPeriodOk returns a tuple with the DraftsPerPeriod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetDraftsPerPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DraftsPerPeriod.Get(), o.DraftsPerPeriod.IsSet()
}

// SetDraftsPerPeriod sets field value
func (o *UsageLimits) SetDraftsPerPeriod(v int32) {
	o.DraftsPerPeriod.Set(&v)
}

// GetPerDraftMinor returns the PerDraftMinor field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetPerDraftMinor() int32 {
	if o == nil || o.PerDraftMinor.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PerDraftMinor.Get()
}

// GetPerDraftMinorOk returns a tuple with the PerDraftMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetPerDraftMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerDraftMinor.Get(), o.PerDraftMinor.IsSet()
}

// SetPerDraftMinor sets field value
func (o *UsageLimits) SetPerDraftMinor(v int32) {
	o.PerDraftMinor.Set(&v)
}

// GetTemplatesPerPeriod returns the TemplatesPerPeriod field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetTemplatesPerPeriod() int32 {
	if o == nil || o.TemplatesPerPeriod.Get() == nil {
		var ret int32
		return ret
	}

	return *o.TemplatesPerPeriod.Get()
}

// GetTemplatesPerPeriodOk returns a tuple with the TemplatesPerPeriod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetTemplatesPerPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplatesPerPeriod.Get(), o.TemplatesPerPeriod.IsSet()
}

// SetTemplatesPerPeriod sets field value
func (o *UsageLimits) SetTemplatesPerPeriod(v int32) {
	o.TemplatesPerPeriod.Set(&v)
}

// GetSignersPerPeriod returns the SignersPerPeriod field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetSignersPerPeriod() int32 {
	if o == nil || o.SignersPerPeriod.Get() == nil {
		var ret int32
		return ret
	}

	return *o.SignersPerPeriod.Get()
}

// GetSignersPerPeriodOk returns a tuple with the SignersPerPeriod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetSignersPerPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignersPerPeriod.Get(), o.SignersPerPeriod.IsSet()
}

// SetSignersPerPeriod sets field value
func (o *UsageLimits) SetSignersPerPeriod(v int32) {
	o.SignersPerPeriod.Set(&v)
}

// GetPeople returns the People field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetPeople() int32 {
	if o == nil || o.People.Get() == nil {
		var ret int32
		return ret
	}

	return *o.People.Get()
}

// GetPeopleOk returns a tuple with the People field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetPeopleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.People.Get(), o.People.IsSet()
}

// SetPeople sets field value
func (o *UsageLimits) SetPeople(v int32) {
	o.People.Set(&v)
}

// GetPersonMonthlyMinor returns the PersonMonthlyMinor field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetPersonMonthlyMinor() int32 {
	if o == nil || o.PersonMonthlyMinor.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PersonMonthlyMinor.Get()
}

// GetPersonMonthlyMinorOk returns a tuple with the PersonMonthlyMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetPersonMonthlyMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PersonMonthlyMinor.Get(), o.PersonMonthlyMinor.IsSet()
}

// SetPersonMonthlyMinor sets field value
func (o *UsageLimits) SetPersonMonthlyMinor(v int32) {
	o.PersonMonthlyMinor.Set(&v)
}

// GetUploadMb returns the UploadMb field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetUploadMb() int32 {
	if o == nil || o.UploadMb.Get() == nil {
		var ret int32
		return ret
	}

	return *o.UploadMb.Get()
}

// GetUploadMbOk returns a tuple with the UploadMb field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetUploadMbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadMb.Get(), o.UploadMb.IsSet()
}

// SetUploadMb sets field value
func (o *UsageLimits) SetUploadMb(v int32) {
	o.UploadMb.Set(&v)
}

// GetEmailsPerDay returns the EmailsPerDay field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetEmailsPerDay() int32 {
	if o == nil || o.EmailsPerDay.Get() == nil {
		var ret int32
		return ret
	}

	return *o.EmailsPerDay.Get()
}

// GetEmailsPerDayOk returns a tuple with the EmailsPerDay field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetEmailsPerDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailsPerDay.Get(), o.EmailsPerDay.IsSet()
}

// SetEmailsPerDay sets field value
func (o *UsageLimits) SetEmailsPerDay(v int32) {
	o.EmailsPerDay.Set(&v)
}

// GetBulkSendMax returns the BulkSendMax field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UsageLimits) GetBulkSendMax() int32 {
	if o == nil || o.BulkSendMax.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BulkSendMax.Get()
}

// GetBulkSendMaxOk returns a tuple with the BulkSendMax field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageLimits) GetBulkSendMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BulkSendMax.Get(), o.BulkSendMax.IsSet()
}

// SetBulkSendMax sets field value
func (o *UsageLimits) SetBulkSendMax(v int32) {
	o.BulkSendMax.Set(&v)
}

// GetPaymentsFeePercent returns the PaymentsFeePercent field value
func (o *UsageLimits) GetPaymentsFeePercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PaymentsFeePercent
}

// GetPaymentsFeePercentOk returns a tuple with the PaymentsFeePercent field value
// and a boolean to check if the value has been set.
func (o *UsageLimits) GetPaymentsFeePercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentsFeePercent, true
}

// SetPaymentsFeePercent sets field value
func (o *UsageLimits) SetPaymentsFeePercent(v float32) {
	o.PaymentsFeePercent = v
}

func (o UsageLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["documentsPerPeriod"] = o.DocumentsPerPeriod.Get()
	toSerialize["perDocumentMinor"] = o.PerDocumentMinor.Get()
	toSerialize["draftsPerPeriod"] = o.DraftsPerPeriod.Get()
	toSerialize["perDraftMinor"] = o.PerDraftMinor.Get()
	toSerialize["templatesPerPeriod"] = o.TemplatesPerPeriod.Get()
	toSerialize["signersPerPeriod"] = o.SignersPerPeriod.Get()
	toSerialize["people"] = o.People.Get()
	toSerialize["personMonthlyMinor"] = o.PersonMonthlyMinor.Get()
	toSerialize["uploadMb"] = o.UploadMb.Get()
	toSerialize["emailsPerDay"] = o.EmailsPerDay.Get()
	toSerialize["bulkSendMax"] = o.BulkSendMax.Get()
	toSerialize["paymentsFeePercent"] = o.PaymentsFeePercent
	return toSerialize, nil
}

func (o *UsageLimits) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"documentsPerPeriod",
		"perDocumentMinor",
		"draftsPerPeriod",
		"perDraftMinor",
		"templatesPerPeriod",
		"signersPerPeriod",
		"people",
		"personMonthlyMinor",
		"uploadMb",
		"emailsPerDay",
		"bulkSendMax",
		"paymentsFeePercent",
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

	varUsageLimits := _UsageLimits{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageLimits)

	if err != nil {
		return err
	}

	*o = UsageLimits(varUsageLimits)

	return err
}

type NullableUsageLimits struct {
	value *UsageLimits
	isSet bool
}

func (v NullableUsageLimits) Get() *UsageLimits {
	return v.value
}

func (v *NullableUsageLimits) Set(val *UsageLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageLimits(val *UsageLimits) *NullableUsageLimits {
	return &NullableUsageLimits{value: val, isSet: true}
}

func (v NullableUsageLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



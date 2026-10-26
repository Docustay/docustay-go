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

// checks if the Usage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Usage{}

// Usage struct for Usage
type Usage struct {
	Plan NullableString `json:"plan"`
	Period NullableUsagePeriod `json:"period"`
	// When this allowance period began (ISO 8601, UTC): the start of the billing period on Pay as you go, Pro and Lifetime, the 1st of the month on Free.
	PeriodStart NullableString `json:"periodStart,omitempty"`
	// When it ends and every count resets (ISO 8601, UTC); shown in the app as 'Resets Nov 19'.
	PeriodEnd NullableString `json:"periodEnd,omitempty"`
	Documents NullableUsageDocuments `json:"documents"`
	Limits NullableUsageLimits `json:"limits,omitempty"`
	Drafts NullableUsageDrafts `json:"drafts,omitempty"`
	Billing NullableUsageBilling `json:"billing,omitempty"`
	Emails NullableUsageEmails `json:"emails"`
	Templates NullableUsageTemplates `json:"templates,omitempty"`
	Signers NullableUsageSigners `json:"signers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Usage Usage

// NewUsage instantiates a new Usage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsage(plan NullableString, period NullableUsagePeriod, documents NullableUsageDocuments, emails NullableUsageEmails) *Usage {
	this := Usage{}
	this.Plan = plan
	this.Period = period
	this.Documents = documents
	this.Emails = emails
	return &this
}

// NewUsageWithDefaults instantiates a new Usage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageWithDefaults() *Usage {
	this := Usage{}
	return &this
}

// GetPlan returns the Plan field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Usage) GetPlan() string {
	if o == nil || o.Plan.Get() == nil {
		var ret string
		return ret
	}

	return *o.Plan.Get()
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plan.Get(), o.Plan.IsSet()
}

// SetPlan sets field value
func (o *Usage) SetPlan(v string) {
	o.Plan.Set(&v)
}

// GetPeriod returns the Period field value
// If the value is explicit nil, the zero value for UsagePeriod will be returned
func (o *Usage) GetPeriod() UsagePeriod {
	if o == nil || o.Period.Get() == nil {
		var ret UsagePeriod
		return ret
	}

	return *o.Period.Get()
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetPeriodOk() (*UsagePeriod, bool) {
	if o == nil {
		return nil, false
	}
	return o.Period.Get(), o.Period.IsSet()
}

// SetPeriod sets field value
func (o *Usage) SetPeriod(v UsagePeriod) {
	o.Period.Set(&v)
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Usage) GetPeriodStart() string {
	if o == nil || IsNil(o.PeriodStart.Get()) {
		var ret string
		return ret
	}
	return *o.PeriodStart.Get()
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetPeriodStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeriodStart.Get(), o.PeriodStart.IsSet()
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *Usage) HasPeriodStart() bool {
	if o != nil && o.PeriodStart.IsSet() {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given NullableString and assigns it to the PeriodStart field.
func (o *Usage) SetPeriodStart(v string) {
	o.PeriodStart.Set(&v)
}
// SetPeriodStartNil sets the value for PeriodStart to be an explicit nil
func (o *Usage) SetPeriodStartNil() {
	o.PeriodStart.Set(nil)
}

// UnsetPeriodStart ensures that no value is present for PeriodStart, not even an explicit nil
func (o *Usage) UnsetPeriodStart() {
	o.PeriodStart.Unset()
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Usage) GetPeriodEnd() string {
	if o == nil || IsNil(o.PeriodEnd.Get()) {
		var ret string
		return ret
	}
	return *o.PeriodEnd.Get()
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetPeriodEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeriodEnd.Get(), o.PeriodEnd.IsSet()
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *Usage) HasPeriodEnd() bool {
	if o != nil && o.PeriodEnd.IsSet() {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given NullableString and assigns it to the PeriodEnd field.
func (o *Usage) SetPeriodEnd(v string) {
	o.PeriodEnd.Set(&v)
}
// SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil
func (o *Usage) SetPeriodEndNil() {
	o.PeriodEnd.Set(nil)
}

// UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
func (o *Usage) UnsetPeriodEnd() {
	o.PeriodEnd.Unset()
}

// GetDocuments returns the Documents field value
// If the value is explicit nil, the zero value for UsageDocuments will be returned
func (o *Usage) GetDocuments() UsageDocuments {
	if o == nil || o.Documents.Get() == nil {
		var ret UsageDocuments
		return ret
	}

	return *o.Documents.Get()
}

// GetDocumentsOk returns a tuple with the Documents field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetDocumentsOk() (*UsageDocuments, bool) {
	if o == nil {
		return nil, false
	}
	return o.Documents.Get(), o.Documents.IsSet()
}

// SetDocuments sets field value
func (o *Usage) SetDocuments(v UsageDocuments) {
	o.Documents.Set(&v)
}

// GetLimits returns the Limits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Usage) GetLimits() UsageLimits {
	if o == nil || IsNil(o.Limits.Get()) {
		var ret UsageLimits
		return ret
	}
	return *o.Limits.Get()
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetLimitsOk() (*UsageLimits, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limits.Get(), o.Limits.IsSet()
}

// HasLimits returns a boolean if a field has been set.
func (o *Usage) HasLimits() bool {
	if o != nil && o.Limits.IsSet() {
		return true
	}

	return false
}

// SetLimits gets a reference to the given NullableUsageLimits and assigns it to the Limits field.
func (o *Usage) SetLimits(v UsageLimits) {
	o.Limits.Set(&v)
}
// SetLimitsNil sets the value for Limits to be an explicit nil
func (o *Usage) SetLimitsNil() {
	o.Limits.Set(nil)
}

// UnsetLimits ensures that no value is present for Limits, not even an explicit nil
func (o *Usage) UnsetLimits() {
	o.Limits.Unset()
}

// GetDrafts returns the Drafts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Usage) GetDrafts() UsageDrafts {
	if o == nil || IsNil(o.Drafts.Get()) {
		var ret UsageDrafts
		return ret
	}
	return *o.Drafts.Get()
}

// GetDraftsOk returns a tuple with the Drafts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetDraftsOk() (*UsageDrafts, bool) {
	if o == nil {
		return nil, false
	}
	return o.Drafts.Get(), o.Drafts.IsSet()
}

// HasDrafts returns a boolean if a field has been set.
func (o *Usage) HasDrafts() bool {
	if o != nil && o.Drafts.IsSet() {
		return true
	}

	return false
}

// SetDrafts gets a reference to the given NullableUsageDrafts and assigns it to the Drafts field.
func (o *Usage) SetDrafts(v UsageDrafts) {
	o.Drafts.Set(&v)
}
// SetDraftsNil sets the value for Drafts to be an explicit nil
func (o *Usage) SetDraftsNil() {
	o.Drafts.Set(nil)
}

// UnsetDrafts ensures that no value is present for Drafts, not even an explicit nil
func (o *Usage) UnsetDrafts() {
	o.Drafts.Unset()
}

// GetBilling returns the Billing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Usage) GetBilling() UsageBilling {
	if o == nil || IsNil(o.Billing.Get()) {
		var ret UsageBilling
		return ret
	}
	return *o.Billing.Get()
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetBillingOk() (*UsageBilling, bool) {
	if o == nil {
		return nil, false
	}
	return o.Billing.Get(), o.Billing.IsSet()
}

// HasBilling returns a boolean if a field has been set.
func (o *Usage) HasBilling() bool {
	if o != nil && o.Billing.IsSet() {
		return true
	}

	return false
}

// SetBilling gets a reference to the given NullableUsageBilling and assigns it to the Billing field.
func (o *Usage) SetBilling(v UsageBilling) {
	o.Billing.Set(&v)
}
// SetBillingNil sets the value for Billing to be an explicit nil
func (o *Usage) SetBillingNil() {
	o.Billing.Set(nil)
}

// UnsetBilling ensures that no value is present for Billing, not even an explicit nil
func (o *Usage) UnsetBilling() {
	o.Billing.Unset()
}

// GetEmails returns the Emails field value
// If the value is explicit nil, the zero value for UsageEmails will be returned
func (o *Usage) GetEmails() UsageEmails {
	if o == nil || o.Emails.Get() == nil {
		var ret UsageEmails
		return ret
	}

	return *o.Emails.Get()
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetEmailsOk() (*UsageEmails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails.Get(), o.Emails.IsSet()
}

// SetEmails sets field value
func (o *Usage) SetEmails(v UsageEmails) {
	o.Emails.Set(&v)
}

// GetTemplates returns the Templates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Usage) GetTemplates() UsageTemplates {
	if o == nil || IsNil(o.Templates.Get()) {
		var ret UsageTemplates
		return ret
	}
	return *o.Templates.Get()
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetTemplatesOk() (*UsageTemplates, bool) {
	if o == nil {
		return nil, false
	}
	return o.Templates.Get(), o.Templates.IsSet()
}

// HasTemplates returns a boolean if a field has been set.
func (o *Usage) HasTemplates() bool {
	if o != nil && o.Templates.IsSet() {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given NullableUsageTemplates and assigns it to the Templates field.
func (o *Usage) SetTemplates(v UsageTemplates) {
	o.Templates.Set(&v)
}
// SetTemplatesNil sets the value for Templates to be an explicit nil
func (o *Usage) SetTemplatesNil() {
	o.Templates.Set(nil)
}

// UnsetTemplates ensures that no value is present for Templates, not even an explicit nil
func (o *Usage) UnsetTemplates() {
	o.Templates.Unset()
}

// GetSigners returns the Signers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Usage) GetSigners() UsageSigners {
	if o == nil || IsNil(o.Signers.Get()) {
		var ret UsageSigners
		return ret
	}
	return *o.Signers.Get()
}

// GetSignersOk returns a tuple with the Signers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Usage) GetSignersOk() (*UsageSigners, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signers.Get(), o.Signers.IsSet()
}

// HasSigners returns a boolean if a field has been set.
func (o *Usage) HasSigners() bool {
	if o != nil && o.Signers.IsSet() {
		return true
	}

	return false
}

// SetSigners gets a reference to the given NullableUsageSigners and assigns it to the Signers field.
func (o *Usage) SetSigners(v UsageSigners) {
	o.Signers.Set(&v)
}
// SetSignersNil sets the value for Signers to be an explicit nil
func (o *Usage) SetSignersNil() {
	o.Signers.Set(nil)
}

// UnsetSigners ensures that no value is present for Signers, not even an explicit nil
func (o *Usage) UnsetSigners() {
	o.Signers.Unset()
}

func (o Usage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Usage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plan"] = o.Plan.Get()
	toSerialize["period"] = o.Period.Get()
	if o.PeriodStart.IsSet() {
		toSerialize["periodStart"] = o.PeriodStart.Get()
	}
	if o.PeriodEnd.IsSet() {
		toSerialize["periodEnd"] = o.PeriodEnd.Get()
	}
	toSerialize["documents"] = o.Documents.Get()
	if o.Limits.IsSet() {
		toSerialize["limits"] = o.Limits.Get()
	}
	if o.Drafts.IsSet() {
		toSerialize["drafts"] = o.Drafts.Get()
	}
	if o.Billing.IsSet() {
		toSerialize["billing"] = o.Billing.Get()
	}
	toSerialize["emails"] = o.Emails.Get()
	if o.Templates.IsSet() {
		toSerialize["templates"] = o.Templates.Get()
	}
	if o.Signers.IsSet() {
		toSerialize["signers"] = o.Signers.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Usage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plan",
		"period",
		"documents",
		"emails",
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

	varUsage := _Usage{}

	err = json.Unmarshal(data, &varUsage)

	if err != nil {
		return err
	}

	*o = Usage(varUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "plan")
		delete(additionalProperties, "period")
		delete(additionalProperties, "periodStart")
		delete(additionalProperties, "periodEnd")
		delete(additionalProperties, "documents")
		delete(additionalProperties, "limits")
		delete(additionalProperties, "drafts")
		delete(additionalProperties, "billing")
		delete(additionalProperties, "emails")
		delete(additionalProperties, "templates")
		delete(additionalProperties, "signers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsage struct {
	value *Usage
	isSet bool
}

func (v NullableUsage) Get() *Usage {
	return v.value
}

func (v *NullableUsage) Set(val *Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsage(val *Usage) *NullableUsage {
	return &NullableUsage{value: val, isSet: true}
}

func (v NullableUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



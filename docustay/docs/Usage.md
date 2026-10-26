# Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | **NullableString** |  | 
**Period** | [**NullableUsagePeriod**](UsagePeriod.md) |  | 
**PeriodStart** | Pointer to **NullableString** | When this allowance period began (ISO 8601, UTC): the start of the billing period on Pay as you go, Pro and Lifetime, the 1st of the month on Free. | [optional] 
**PeriodEnd** | Pointer to **NullableString** | When it ends and every count resets (ISO 8601, UTC); shown in the app as &#39;Resets Nov 19&#39;. | [optional] 
**Documents** | [**NullableUsageDocuments**](UsageDocuments.md) |  | 
**Limits** | Pointer to [**NullableUsageLimits**](UsageLimits.md) |  | [optional] 
**Drafts** | Pointer to [**NullableUsageDrafts**](UsageDrafts.md) |  | [optional] 
**Billing** | Pointer to [**NullableUsageBilling**](UsageBilling.md) |  | [optional] 
**Emails** | [**NullableUsageEmails**](UsageEmails.md) |  | 
**Templates** | Pointer to [**NullableUsageTemplates**](UsageTemplates.md) |  | [optional] 
**Signers** | Pointer to [**NullableUsageSigners**](UsageSigners.md) |  | [optional] 

## Methods

### NewUsage

`func NewUsage(plan NullableString, period NullableUsagePeriod, documents NullableUsageDocuments, emails NullableUsageEmails, ) *Usage`

NewUsage instantiates a new Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageWithDefaults

`func NewUsageWithDefaults() *Usage`

NewUsageWithDefaults instantiates a new Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *Usage) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Usage) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Usage) SetPlan(v string)`

SetPlan sets Plan field to given value.


### SetPlanNil

`func (o *Usage) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *Usage) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetPeriod

`func (o *Usage) GetPeriod() UsagePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Usage) GetPeriodOk() (*UsagePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Usage) SetPeriod(v UsagePeriod)`

SetPeriod sets Period field to given value.


### SetPeriodNil

`func (o *Usage) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *Usage) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetPeriodStart

`func (o *Usage) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *Usage) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *Usage) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *Usage) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### SetPeriodStartNil

`func (o *Usage) SetPeriodStartNil(b bool)`

 SetPeriodStartNil sets the value for PeriodStart to be an explicit nil

### UnsetPeriodStart
`func (o *Usage) UnsetPeriodStart()`

UnsetPeriodStart ensures that no value is present for PeriodStart, not even an explicit nil
### GetPeriodEnd

`func (o *Usage) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *Usage) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *Usage) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *Usage) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### SetPeriodEndNil

`func (o *Usage) SetPeriodEndNil(b bool)`

 SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil

### UnsetPeriodEnd
`func (o *Usage) UnsetPeriodEnd()`

UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
### GetDocuments

`func (o *Usage) GetDocuments() UsageDocuments`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Usage) GetDocumentsOk() (*UsageDocuments, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Usage) SetDocuments(v UsageDocuments)`

SetDocuments sets Documents field to given value.


### SetDocumentsNil

`func (o *Usage) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *Usage) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetLimits

`func (o *Usage) GetLimits() UsageLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Usage) GetLimitsOk() (*UsageLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Usage) SetLimits(v UsageLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *Usage) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimitsNil

`func (o *Usage) SetLimitsNil(b bool)`

 SetLimitsNil sets the value for Limits to be an explicit nil

### UnsetLimits
`func (o *Usage) UnsetLimits()`

UnsetLimits ensures that no value is present for Limits, not even an explicit nil
### GetDrafts

`func (o *Usage) GetDrafts() UsageDrafts`

GetDrafts returns the Drafts field if non-nil, zero value otherwise.

### GetDraftsOk

`func (o *Usage) GetDraftsOk() (*UsageDrafts, bool)`

GetDraftsOk returns a tuple with the Drafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrafts

`func (o *Usage) SetDrafts(v UsageDrafts)`

SetDrafts sets Drafts field to given value.

### HasDrafts

`func (o *Usage) HasDrafts() bool`

HasDrafts returns a boolean if a field has been set.

### SetDraftsNil

`func (o *Usage) SetDraftsNil(b bool)`

 SetDraftsNil sets the value for Drafts to be an explicit nil

### UnsetDrafts
`func (o *Usage) UnsetDrafts()`

UnsetDrafts ensures that no value is present for Drafts, not even an explicit nil
### GetBilling

`func (o *Usage) GetBilling() UsageBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Usage) GetBillingOk() (*UsageBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Usage) SetBilling(v UsageBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *Usage) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### SetBillingNil

`func (o *Usage) SetBillingNil(b bool)`

 SetBillingNil sets the value for Billing to be an explicit nil

### UnsetBilling
`func (o *Usage) UnsetBilling()`

UnsetBilling ensures that no value is present for Billing, not even an explicit nil
### GetEmails

`func (o *Usage) GetEmails() UsageEmails`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *Usage) GetEmailsOk() (*UsageEmails, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *Usage) SetEmails(v UsageEmails)`

SetEmails sets Emails field to given value.


### SetEmailsNil

`func (o *Usage) SetEmailsNil(b bool)`

 SetEmailsNil sets the value for Emails to be an explicit nil

### UnsetEmails
`func (o *Usage) UnsetEmails()`

UnsetEmails ensures that no value is present for Emails, not even an explicit nil
### GetTemplates

`func (o *Usage) GetTemplates() UsageTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *Usage) GetTemplatesOk() (*UsageTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *Usage) SetTemplates(v UsageTemplates)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *Usage) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### SetTemplatesNil

`func (o *Usage) SetTemplatesNil(b bool)`

 SetTemplatesNil sets the value for Templates to be an explicit nil

### UnsetTemplates
`func (o *Usage) UnsetTemplates()`

UnsetTemplates ensures that no value is present for Templates, not even an explicit nil
### GetSigners

`func (o *Usage) GetSigners() UsageSigners`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *Usage) GetSignersOk() (*UsageSigners, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *Usage) SetSigners(v UsageSigners)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *Usage) HasSigners() bool`

HasSigners returns a boolean if a field has been set.

### SetSignersNil

`func (o *Usage) SetSignersNil(b bool)`

 SetSignersNil sets the value for Signers to be an explicit nil

### UnsetSigners
`func (o *Usage) UnsetSigners()`

UnsetSigners ensures that no value is present for Signers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



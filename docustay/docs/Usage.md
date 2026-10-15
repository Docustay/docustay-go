# Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | **NullableString** |  | 
**Period** | [**UsagePeriod**](UsagePeriod.md) |  | 
**Documents** | [**UsageDocuments**](UsageDocuments.md) |  | 
**Emails** | [**UsageEmails**](UsageEmails.md) |  | 
**Templates** | Pointer to [**UsageTemplates**](UsageTemplates.md) |  | [optional] 

## Methods

### NewUsage

`func NewUsage(plan NullableString, period UsagePeriod, documents UsageDocuments, emails UsageEmails, ) *Usage`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



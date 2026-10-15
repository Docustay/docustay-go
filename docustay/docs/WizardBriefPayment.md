# WizardBriefPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collect** | **bool** |  | 
**Amount** | **string** | In major units, typed by the person (for example \&quot;1500\&quot;). The assistant never decides an amount. | 
**Currency** | **string** |  | 
**PayerRole** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**ScheduleNote** | Pointer to **string** |  | [optional] 

## Methods

### NewWizardBriefPayment

`func NewWizardBriefPayment(collect bool, amount string, currency string, description string, ) *WizardBriefPayment`

NewWizardBriefPayment instantiates a new WizardBriefPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardBriefPaymentWithDefaults

`func NewWizardBriefPaymentWithDefaults() *WizardBriefPayment`

NewWizardBriefPaymentWithDefaults instantiates a new WizardBriefPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollect

`func (o *WizardBriefPayment) GetCollect() bool`

GetCollect returns the Collect field if non-nil, zero value otherwise.

### GetCollectOk

`func (o *WizardBriefPayment) GetCollectOk() (*bool, bool)`

GetCollectOk returns a tuple with the Collect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollect

`func (o *WizardBriefPayment) SetCollect(v bool)`

SetCollect sets Collect field to given value.


### GetAmount

`func (o *WizardBriefPayment) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WizardBriefPayment) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WizardBriefPayment) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *WizardBriefPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WizardBriefPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WizardBriefPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPayerRole

`func (o *WizardBriefPayment) GetPayerRole() string`

GetPayerRole returns the PayerRole field if non-nil, zero value otherwise.

### GetPayerRoleOk

`func (o *WizardBriefPayment) GetPayerRoleOk() (*string, bool)`

GetPayerRoleOk returns a tuple with the PayerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerRole

`func (o *WizardBriefPayment) SetPayerRole(v string)`

SetPayerRole sets PayerRole field to given value.

### HasPayerRole

`func (o *WizardBriefPayment) HasPayerRole() bool`

HasPayerRole returns a boolean if a field has been set.

### GetDescription

`func (o *WizardBriefPayment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WizardBriefPayment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WizardBriefPayment) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetScheduleNote

`func (o *WizardBriefPayment) GetScheduleNote() string`

GetScheduleNote returns the ScheduleNote field if non-nil, zero value otherwise.

### GetScheduleNoteOk

`func (o *WizardBriefPayment) GetScheduleNoteOk() (*string, bool)`

GetScheduleNoteOk returns a tuple with the ScheduleNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleNote

`func (o *WizardBriefPayment) SetScheduleNote(v string)`

SetScheduleNote sets ScheduleNote field to given value.

### HasScheduleNote

`func (o *WizardBriefPayment) HasScheduleNote() bool`

HasScheduleNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



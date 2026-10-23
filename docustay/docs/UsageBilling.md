# UsageBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerDocumentMinor** | **NullableInt32** |  | 
**PerDraftMinor** | **NullableInt32** |  | 
**BillableDocuments** | **int32** |  | 
**BillableDrafts** | **int32** |  | 
**UnbilledMinor** | **int32** |  | 
**BudgetMinor** | **int32** |  | 
**PercentOfBudget** | **int32** |  | 
**Paused** | **bool** |  | 
**Currency** | **string** |  | 

## Methods

### NewUsageBilling

`func NewUsageBilling(perDocumentMinor NullableInt32, perDraftMinor NullableInt32, billableDocuments int32, billableDrafts int32, unbilledMinor int32, budgetMinor int32, percentOfBudget int32, paused bool, currency string, ) *UsageBilling`

NewUsageBilling instantiates a new UsageBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageBillingWithDefaults

`func NewUsageBillingWithDefaults() *UsageBilling`

NewUsageBillingWithDefaults instantiates a new UsageBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerDocumentMinor

`func (o *UsageBilling) GetPerDocumentMinor() int32`

GetPerDocumentMinor returns the PerDocumentMinor field if non-nil, zero value otherwise.

### GetPerDocumentMinorOk

`func (o *UsageBilling) GetPerDocumentMinorOk() (*int32, bool)`

GetPerDocumentMinorOk returns a tuple with the PerDocumentMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerDocumentMinor

`func (o *UsageBilling) SetPerDocumentMinor(v int32)`

SetPerDocumentMinor sets PerDocumentMinor field to given value.


### SetPerDocumentMinorNil

`func (o *UsageBilling) SetPerDocumentMinorNil(b bool)`

 SetPerDocumentMinorNil sets the value for PerDocumentMinor to be an explicit nil

### UnsetPerDocumentMinor
`func (o *UsageBilling) UnsetPerDocumentMinor()`

UnsetPerDocumentMinor ensures that no value is present for PerDocumentMinor, not even an explicit nil
### GetPerDraftMinor

`func (o *UsageBilling) GetPerDraftMinor() int32`

GetPerDraftMinor returns the PerDraftMinor field if non-nil, zero value otherwise.

### GetPerDraftMinorOk

`func (o *UsageBilling) GetPerDraftMinorOk() (*int32, bool)`

GetPerDraftMinorOk returns a tuple with the PerDraftMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerDraftMinor

`func (o *UsageBilling) SetPerDraftMinor(v int32)`

SetPerDraftMinor sets PerDraftMinor field to given value.


### SetPerDraftMinorNil

`func (o *UsageBilling) SetPerDraftMinorNil(b bool)`

 SetPerDraftMinorNil sets the value for PerDraftMinor to be an explicit nil

### UnsetPerDraftMinor
`func (o *UsageBilling) UnsetPerDraftMinor()`

UnsetPerDraftMinor ensures that no value is present for PerDraftMinor, not even an explicit nil
### GetBillableDocuments

`func (o *UsageBilling) GetBillableDocuments() int32`

GetBillableDocuments returns the BillableDocuments field if non-nil, zero value otherwise.

### GetBillableDocumentsOk

`func (o *UsageBilling) GetBillableDocumentsOk() (*int32, bool)`

GetBillableDocumentsOk returns a tuple with the BillableDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableDocuments

`func (o *UsageBilling) SetBillableDocuments(v int32)`

SetBillableDocuments sets BillableDocuments field to given value.


### GetBillableDrafts

`func (o *UsageBilling) GetBillableDrafts() int32`

GetBillableDrafts returns the BillableDrafts field if non-nil, zero value otherwise.

### GetBillableDraftsOk

`func (o *UsageBilling) GetBillableDraftsOk() (*int32, bool)`

GetBillableDraftsOk returns a tuple with the BillableDrafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableDrafts

`func (o *UsageBilling) SetBillableDrafts(v int32)`

SetBillableDrafts sets BillableDrafts field to given value.


### GetUnbilledMinor

`func (o *UsageBilling) GetUnbilledMinor() int32`

GetUnbilledMinor returns the UnbilledMinor field if non-nil, zero value otherwise.

### GetUnbilledMinorOk

`func (o *UsageBilling) GetUnbilledMinorOk() (*int32, bool)`

GetUnbilledMinorOk returns a tuple with the UnbilledMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbilledMinor

`func (o *UsageBilling) SetUnbilledMinor(v int32)`

SetUnbilledMinor sets UnbilledMinor field to given value.


### GetBudgetMinor

`func (o *UsageBilling) GetBudgetMinor() int32`

GetBudgetMinor returns the BudgetMinor field if non-nil, zero value otherwise.

### GetBudgetMinorOk

`func (o *UsageBilling) GetBudgetMinorOk() (*int32, bool)`

GetBudgetMinorOk returns a tuple with the BudgetMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetMinor

`func (o *UsageBilling) SetBudgetMinor(v int32)`

SetBudgetMinor sets BudgetMinor field to given value.


### GetPercentOfBudget

`func (o *UsageBilling) GetPercentOfBudget() int32`

GetPercentOfBudget returns the PercentOfBudget field if non-nil, zero value otherwise.

### GetPercentOfBudgetOk

`func (o *UsageBilling) GetPercentOfBudgetOk() (*int32, bool)`

GetPercentOfBudgetOk returns a tuple with the PercentOfBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOfBudget

`func (o *UsageBilling) SetPercentOfBudget(v int32)`

SetPercentOfBudget sets PercentOfBudget field to given value.


### GetPaused

`func (o *UsageBilling) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *UsageBilling) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *UsageBilling) SetPaused(v bool)`

SetPaused sets Paused field to given value.


### GetCurrency

`func (o *UsageBilling) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UsageBilling) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UsageBilling) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



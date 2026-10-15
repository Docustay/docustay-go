# SetContractDateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Date** | Pointer to **NullableString** |  | [optional] 
**AmountMinor** | Pointer to **NullableInt32** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSetContractDateRequest

`func NewSetContractDateRequest(kind string, ) *SetContractDateRequest`

NewSetContractDateRequest instantiates a new SetContractDateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetContractDateRequestWithDefaults

`func NewSetContractDateRequestWithDefaults() *SetContractDateRequest`

NewSetContractDateRequestWithDefaults instantiates a new SetContractDateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SetContractDateRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SetContractDateRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SetContractDateRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDate

`func (o *SetContractDateRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SetContractDateRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SetContractDateRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *SetContractDateRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *SetContractDateRequest) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *SetContractDateRequest) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetAmountMinor

`func (o *SetContractDateRequest) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *SetContractDateRequest) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *SetContractDateRequest) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *SetContractDateRequest) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### SetAmountMinorNil

`func (o *SetContractDateRequest) SetAmountMinorNil(b bool)`

 SetAmountMinorNil sets the value for AmountMinor to be an explicit nil

### UnsetAmountMinor
`func (o *SetContractDateRequest) UnsetAmountMinor()`

UnsetAmountMinor ensures that no value is present for AmountMinor, not even an explicit nil
### GetCurrency

`func (o *SetContractDateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SetContractDateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SetContractDateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SetContractDateRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *SetContractDateRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *SetContractDateRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



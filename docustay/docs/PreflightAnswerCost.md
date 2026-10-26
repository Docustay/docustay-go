# PreflightAnswerCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minor** | **int32** |  | 
**Currency** | **string** |  | 
**Units** | **float32** |  | 
**PerUnitMinor** | **NullableFloat32** |  | 
**Per** | **NullableString** |  | 
**Prorated** | **bool** |  | 
**FeePercent** | Pointer to **float32** |  | [optional] 

## Methods

### NewPreflightAnswerCost

`func NewPreflightAnswerCost(minor int32, currency string, units float32, perUnitMinor NullableFloat32, per NullableString, prorated bool, ) *PreflightAnswerCost`

NewPreflightAnswerCost instantiates a new PreflightAnswerCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreflightAnswerCostWithDefaults

`func NewPreflightAnswerCostWithDefaults() *PreflightAnswerCost`

NewPreflightAnswerCostWithDefaults instantiates a new PreflightAnswerCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinor

`func (o *PreflightAnswerCost) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *PreflightAnswerCost) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *PreflightAnswerCost) SetMinor(v int32)`

SetMinor sets Minor field to given value.


### GetCurrency

`func (o *PreflightAnswerCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PreflightAnswerCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PreflightAnswerCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUnits

`func (o *PreflightAnswerCost) GetUnits() float32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *PreflightAnswerCost) GetUnitsOk() (*float32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *PreflightAnswerCost) SetUnits(v float32)`

SetUnits sets Units field to given value.


### GetPerUnitMinor

`func (o *PreflightAnswerCost) GetPerUnitMinor() float32`

GetPerUnitMinor returns the PerUnitMinor field if non-nil, zero value otherwise.

### GetPerUnitMinorOk

`func (o *PreflightAnswerCost) GetPerUnitMinorOk() (*float32, bool)`

GetPerUnitMinorOk returns a tuple with the PerUnitMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUnitMinor

`func (o *PreflightAnswerCost) SetPerUnitMinor(v float32)`

SetPerUnitMinor sets PerUnitMinor field to given value.


### SetPerUnitMinorNil

`func (o *PreflightAnswerCost) SetPerUnitMinorNil(b bool)`

 SetPerUnitMinorNil sets the value for PerUnitMinor to be an explicit nil

### UnsetPerUnitMinor
`func (o *PreflightAnswerCost) UnsetPerUnitMinor()`

UnsetPerUnitMinor ensures that no value is present for PerUnitMinor, not even an explicit nil
### GetPer

`func (o *PreflightAnswerCost) GetPer() string`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *PreflightAnswerCost) GetPerOk() (*string, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *PreflightAnswerCost) SetPer(v string)`

SetPer sets Per field to given value.


### SetPerNil

`func (o *PreflightAnswerCost) SetPerNil(b bool)`

 SetPerNil sets the value for Per to be an explicit nil

### UnsetPer
`func (o *PreflightAnswerCost) UnsetPer()`

UnsetPer ensures that no value is present for Per, not even an explicit nil
### GetProrated

`func (o *PreflightAnswerCost) GetProrated() bool`

GetProrated returns the Prorated field if non-nil, zero value otherwise.

### GetProratedOk

`func (o *PreflightAnswerCost) GetProratedOk() (*bool, bool)`

GetProratedOk returns a tuple with the Prorated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrated

`func (o *PreflightAnswerCost) SetProrated(v bool)`

SetProrated sets Prorated field to given value.


### GetFeePercent

`func (o *PreflightAnswerCost) GetFeePercent() float32`

GetFeePercent returns the FeePercent field if non-nil, zero value otherwise.

### GetFeePercentOk

`func (o *PreflightAnswerCost) GetFeePercentOk() (*float32, bool)`

GetFeePercentOk returns a tuple with the FeePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePercent

`func (o *PreflightAnswerCost) SetFeePercent(v float32)`

SetFeePercent sets FeePercent field to given value.

### HasFeePercent

`func (o *PreflightAnswerCost) HasFeePercent() bool`

HasFeePercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



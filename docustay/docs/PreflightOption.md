# PreflightOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Label** | **string** |  | 
**PriceMinor** | **NullableInt32** | Whole cents; null when the option has no price of its own (raising the budget). | 
**Per** | **NullableString** | &#x60;team member a month&#x60;, &#x60;month&#x60;, &#x60;once&#x60; or null. | 
**Founders** | Pointer to **bool** | Deprecated. Always false: there is one Pro price. Kept so existing clients keep working. | [optional] 
**Unlocks** | **[]string** |  | 

## Methods

### NewPreflightOption

`func NewPreflightOption(key string, label string, priceMinor NullableInt32, per NullableString, unlocks []string, ) *PreflightOption`

NewPreflightOption instantiates a new PreflightOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreflightOptionWithDefaults

`func NewPreflightOptionWithDefaults() *PreflightOption`

NewPreflightOptionWithDefaults instantiates a new PreflightOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PreflightOption) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PreflightOption) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PreflightOption) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *PreflightOption) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PreflightOption) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PreflightOption) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPriceMinor

`func (o *PreflightOption) GetPriceMinor() int32`

GetPriceMinor returns the PriceMinor field if non-nil, zero value otherwise.

### GetPriceMinorOk

`func (o *PreflightOption) GetPriceMinorOk() (*int32, bool)`

GetPriceMinorOk returns a tuple with the PriceMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMinor

`func (o *PreflightOption) SetPriceMinor(v int32)`

SetPriceMinor sets PriceMinor field to given value.


### SetPriceMinorNil

`func (o *PreflightOption) SetPriceMinorNil(b bool)`

 SetPriceMinorNil sets the value for PriceMinor to be an explicit nil

### UnsetPriceMinor
`func (o *PreflightOption) UnsetPriceMinor()`

UnsetPriceMinor ensures that no value is present for PriceMinor, not even an explicit nil
### GetPer

`func (o *PreflightOption) GetPer() string`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *PreflightOption) GetPerOk() (*string, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *PreflightOption) SetPer(v string)`

SetPer sets Per field to given value.


### SetPerNil

`func (o *PreflightOption) SetPerNil(b bool)`

 SetPerNil sets the value for Per to be an explicit nil

### UnsetPer
`func (o *PreflightOption) UnsetPer()`

UnsetPer ensures that no value is present for Per, not even an explicit nil
### GetFounders

`func (o *PreflightOption) GetFounders() bool`

GetFounders returns the Founders field if non-nil, zero value otherwise.

### GetFoundersOk

`func (o *PreflightOption) GetFoundersOk() (*bool, bool)`

GetFoundersOk returns a tuple with the Founders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounders

`func (o *PreflightOption) SetFounders(v bool)`

SetFounders sets Founders field to given value.

### HasFounders

`func (o *PreflightOption) HasFounders() bool`

HasFounders returns a boolean if a field has been set.

### GetUnlocks

`func (o *PreflightOption) GetUnlocks() []string`

GetUnlocks returns the Unlocks field if non-nil, zero value otherwise.

### GetUnlocksOk

`func (o *PreflightOption) GetUnlocksOk() (*[]string, bool)`

GetUnlocksOk returns a tuple with the Unlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlocks

`func (o *PreflightOption) SetUnlocks(v []string)`

SetUnlocks sets Unlocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



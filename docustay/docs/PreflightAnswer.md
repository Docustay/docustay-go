# PreflightAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Allowed** | **bool** |  | 
**Plan** | [**NullablePreflightAnswerPlan**](PreflightAnswerPlan.md) |  | 
**Period** | [**NullablePreflightAnswerPeriod**](PreflightAnswerPeriod.md) |  | 
**Meter** | **map[string]interface{}** |  | 
**Cost** | [**PreflightAnswerCost**](PreflightAnswerCost.md) |  | 
**AfterThis** | **map[string]interface{}** |  | 
**Severity** | **string** |  | 
**Reason** | **NullableString** |  | 
**Options** | [**[]PreflightOption**](PreflightOption.md) |  | 
**Message** | **string** |  | 
**AllowedCount** | Pointer to **int32** | For &#x60;bulk_send&#x60; on a limit: how many of the rows can go now. | [optional] 

## Methods

### NewPreflightAnswer

`func NewPreflightAnswer(action string, allowed bool, plan NullablePreflightAnswerPlan, period NullablePreflightAnswerPeriod, meter map[string]*interface{}, cost PreflightAnswerCost, afterThis map[string]*interface{}, severity string, reason NullableString, options []PreflightOption, message string, ) *PreflightAnswer`

NewPreflightAnswer instantiates a new PreflightAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreflightAnswerWithDefaults

`func NewPreflightAnswerWithDefaults() *PreflightAnswer`

NewPreflightAnswerWithDefaults instantiates a new PreflightAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PreflightAnswer) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PreflightAnswer) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PreflightAnswer) SetAction(v string)`

SetAction sets Action field to given value.


### GetAllowed

`func (o *PreflightAnswer) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *PreflightAnswer) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *PreflightAnswer) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetPlan

`func (o *PreflightAnswer) GetPlan() PreflightAnswerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PreflightAnswer) GetPlanOk() (*PreflightAnswerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PreflightAnswer) SetPlan(v PreflightAnswerPlan)`

SetPlan sets Plan field to given value.


### SetPlanNil

`func (o *PreflightAnswer) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *PreflightAnswer) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetPeriod

`func (o *PreflightAnswer) GetPeriod() PreflightAnswerPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *PreflightAnswer) GetPeriodOk() (*PreflightAnswerPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *PreflightAnswer) SetPeriod(v PreflightAnswerPeriod)`

SetPeriod sets Period field to given value.


### SetPeriodNil

`func (o *PreflightAnswer) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *PreflightAnswer) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetMeter

`func (o *PreflightAnswer) GetMeter() map[string]*interface{}`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *PreflightAnswer) GetMeterOk() (*map[string]*interface{}, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *PreflightAnswer) SetMeter(v map[string]*interface{})`

SetMeter sets Meter field to given value.


### SetMeterNil

`func (o *PreflightAnswer) SetMeterNil(b bool)`

 SetMeterNil sets the value for Meter to be an explicit nil

### UnsetMeter
`func (o *PreflightAnswer) UnsetMeter()`

UnsetMeter ensures that no value is present for Meter, not even an explicit nil
### GetCost

`func (o *PreflightAnswer) GetCost() PreflightAnswerCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *PreflightAnswer) GetCostOk() (*PreflightAnswerCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *PreflightAnswer) SetCost(v PreflightAnswerCost)`

SetCost sets Cost field to given value.


### GetAfterThis

`func (o *PreflightAnswer) GetAfterThis() map[string]*interface{}`

GetAfterThis returns the AfterThis field if non-nil, zero value otherwise.

### GetAfterThisOk

`func (o *PreflightAnswer) GetAfterThisOk() (*map[string]*interface{}, bool)`

GetAfterThisOk returns a tuple with the AfterThis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterThis

`func (o *PreflightAnswer) SetAfterThis(v map[string]*interface{})`

SetAfterThis sets AfterThis field to given value.


### SetAfterThisNil

`func (o *PreflightAnswer) SetAfterThisNil(b bool)`

 SetAfterThisNil sets the value for AfterThis to be an explicit nil

### UnsetAfterThis
`func (o *PreflightAnswer) UnsetAfterThis()`

UnsetAfterThis ensures that no value is present for AfterThis, not even an explicit nil
### GetSeverity

`func (o *PreflightAnswer) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *PreflightAnswer) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *PreflightAnswer) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetReason

`func (o *PreflightAnswer) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PreflightAnswer) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PreflightAnswer) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *PreflightAnswer) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *PreflightAnswer) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetOptions

`func (o *PreflightAnswer) GetOptions() []PreflightOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PreflightAnswer) GetOptionsOk() (*[]PreflightOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PreflightAnswer) SetOptions(v []PreflightOption)`

SetOptions sets Options field to given value.


### GetMessage

`func (o *PreflightAnswer) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PreflightAnswer) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PreflightAnswer) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAllowedCount

`func (o *PreflightAnswer) GetAllowedCount() int32`

GetAllowedCount returns the AllowedCount field if non-nil, zero value otherwise.

### GetAllowedCountOk

`func (o *PreflightAnswer) GetAllowedCountOk() (*int32, bool)`

GetAllowedCountOk returns a tuple with the AllowedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCount

`func (o *PreflightAnswer) SetAllowedCount(v int32)`

SetAllowedCount sets AllowedCount field to given value.

### HasAllowedCount

`func (o *PreflightAnswer) HasAllowedCount() bool`

HasAllowedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



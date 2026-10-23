# UsageEmails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedToday** | **int32** |  | 
**LimitPerDay** | **int32** |  | 
**ResetsAt** | **string** |  | 

## Methods

### NewUsageEmails

`func NewUsageEmails(usedToday int32, limitPerDay int32, resetsAt string, ) *UsageEmails`

NewUsageEmails instantiates a new UsageEmails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageEmailsWithDefaults

`func NewUsageEmailsWithDefaults() *UsageEmails`

NewUsageEmailsWithDefaults instantiates a new UsageEmails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedToday

`func (o *UsageEmails) GetUsedToday() int32`

GetUsedToday returns the UsedToday field if non-nil, zero value otherwise.

### GetUsedTodayOk

`func (o *UsageEmails) GetUsedTodayOk() (*int32, bool)`

GetUsedTodayOk returns a tuple with the UsedToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedToday

`func (o *UsageEmails) SetUsedToday(v int32)`

SetUsedToday sets UsedToday field to given value.


### GetLimitPerDay

`func (o *UsageEmails) GetLimitPerDay() int32`

GetLimitPerDay returns the LimitPerDay field if non-nil, zero value otherwise.

### GetLimitPerDayOk

`func (o *UsageEmails) GetLimitPerDayOk() (*int32, bool)`

GetLimitPerDayOk returns a tuple with the LimitPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPerDay

`func (o *UsageEmails) SetLimitPerDay(v int32)`

SetLimitPerDay sets LimitPerDay field to given value.


### GetResetsAt

`func (o *UsageEmails) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *UsageEmails) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *UsageEmails) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



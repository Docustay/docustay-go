# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Window** | [**ReportWindow**](ReportWindow.md) |  | 
**Timezone** | **string** |  | 
**Totals** | [**ReportTotals**](ReportTotals.md) |  | 
**ByDay** | [**[]ReportByDayInner**](ReportByDayInner.md) |  | 
**ByTemplate** | [**[]ReportByTemplateInner**](ReportByTemplateInner.md) |  | 
**ByTeam** | [**[]ReportByTeamInner**](ReportByTeamInner.md) |  | 
**BySender** | [**[]ReportByTeamInner**](ReportByTeamInner.md) |  | 

## Methods

### NewReport

`func NewReport(window ReportWindow, timezone string, totals ReportTotals, byDay []ReportByDayInner, byTemplate []ReportByTemplateInner, byTeam []ReportByTeamInner, bySender []ReportByTeamInner, ) *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindow

`func (o *Report) GetWindow() ReportWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *Report) GetWindowOk() (*ReportWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *Report) SetWindow(v ReportWindow)`

SetWindow sets Window field to given value.


### GetTimezone

`func (o *Report) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Report) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Report) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetTotals

`func (o *Report) GetTotals() ReportTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Report) GetTotalsOk() (*ReportTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Report) SetTotals(v ReportTotals)`

SetTotals sets Totals field to given value.


### GetByDay

`func (o *Report) GetByDay() []ReportByDayInner`

GetByDay returns the ByDay field if non-nil, zero value otherwise.

### GetByDayOk

`func (o *Report) GetByDayOk() (*[]ReportByDayInner, bool)`

GetByDayOk returns a tuple with the ByDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByDay

`func (o *Report) SetByDay(v []ReportByDayInner)`

SetByDay sets ByDay field to given value.


### GetByTemplate

`func (o *Report) GetByTemplate() []ReportByTemplateInner`

GetByTemplate returns the ByTemplate field if non-nil, zero value otherwise.

### GetByTemplateOk

`func (o *Report) GetByTemplateOk() (*[]ReportByTemplateInner, bool)`

GetByTemplateOk returns a tuple with the ByTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByTemplate

`func (o *Report) SetByTemplate(v []ReportByTemplateInner)`

SetByTemplate sets ByTemplate field to given value.


### GetByTeam

`func (o *Report) GetByTeam() []ReportByTeamInner`

GetByTeam returns the ByTeam field if non-nil, zero value otherwise.

### GetByTeamOk

`func (o *Report) GetByTeamOk() (*[]ReportByTeamInner, bool)`

GetByTeamOk returns a tuple with the ByTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByTeam

`func (o *Report) SetByTeam(v []ReportByTeamInner)`

SetByTeam sets ByTeam field to given value.


### GetBySender

`func (o *Report) GetBySender() []ReportByTeamInner`

GetBySender returns the BySender field if non-nil, zero value otherwise.

### GetBySenderOk

`func (o *Report) GetBySenderOk() (*[]ReportByTeamInner, bool)`

GetBySenderOk returns a tuple with the BySender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySender

`func (o *Report) SetBySender(v []ReportByTeamInner)`

SetBySender sets BySender field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



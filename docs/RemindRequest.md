# RemindRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seat** | Pointer to **string** | Remind only this seat (for example signer_2). Leave out to remind everyone who is up. | [optional] 

## Methods

### NewRemindRequest

`func NewRemindRequest() *RemindRequest`

NewRemindRequest instantiates a new RemindRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemindRequestWithDefaults

`func NewRemindRequestWithDefaults() *RemindRequest`

NewRemindRequestWithDefaults instantiates a new RemindRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeat

`func (o *RemindRequest) GetSeat() string`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *RemindRequest) GetSeatOk() (*string, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *RemindRequest) SetSeat(v string)`

SetSeat sets Seat field to given value.

### HasSeat

`func (o *RemindRequest) HasSeat() bool`

HasSeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



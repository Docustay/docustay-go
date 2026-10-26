# TestSignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seat** | **string** | The seat to sign, for example &#x60;signer_1&#x60;. | 
**Decline** | Pointer to **bool** | Decline the seat instead of signing it: the document becomes &#x60;declined&#x60; and &#x60;documents.document.declined&#x60; fires, as when a person declines. | [optional] 
**Reason** | Pointer to **string** | With &#x60;decline&#x60;, the reason the sender sees. | [optional] 

## Methods

### NewTestSignRequest

`func NewTestSignRequest(seat string, ) *TestSignRequest`

NewTestSignRequest instantiates a new TestSignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSignRequestWithDefaults

`func NewTestSignRequestWithDefaults() *TestSignRequest`

NewTestSignRequestWithDefaults instantiates a new TestSignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeat

`func (o *TestSignRequest) GetSeat() string`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *TestSignRequest) GetSeatOk() (*string, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *TestSignRequest) SetSeat(v string)`

SetSeat sets Seat field to given value.


### GetDecline

`func (o *TestSignRequest) GetDecline() bool`

GetDecline returns the Decline field if non-nil, zero value otherwise.

### GetDeclineOk

`func (o *TestSignRequest) GetDeclineOk() (*bool, bool)`

GetDeclineOk returns a tuple with the Decline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecline

`func (o *TestSignRequest) SetDecline(v bool)`

SetDecline sets Decline field to given value.

### HasDecline

`func (o *TestSignRequest) HasDecline() bool`

HasDecline returns a boolean if a field has been set.

### GetReason

`func (o *TestSignRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TestSignRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TestSignRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TestSignRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



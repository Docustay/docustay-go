# EmbedSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seat** | **string** |  | 
**Origin** | **string** | The website that will show the form; must be on your allowed list. | 

## Methods

### NewEmbedSessionRequest

`func NewEmbedSessionRequest(seat string, origin string, ) *EmbedSessionRequest`

NewEmbedSessionRequest instantiates a new EmbedSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbedSessionRequestWithDefaults

`func NewEmbedSessionRequestWithDefaults() *EmbedSessionRequest`

NewEmbedSessionRequestWithDefaults instantiates a new EmbedSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeat

`func (o *EmbedSessionRequest) GetSeat() string`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *EmbedSessionRequest) GetSeatOk() (*string, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *EmbedSessionRequest) SetSeat(v string)`

SetSeat sets Seat field to given value.


### GetOrigin

`func (o *EmbedSessionRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *EmbedSessionRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *EmbedSessionRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



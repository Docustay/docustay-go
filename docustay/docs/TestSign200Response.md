# TestSign200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Seat** | **string** |  | 
**Signed** | **bool** |  | 
**Declined** | Pointer to **bool** |  | [optional] 
**State** | **string** |  | 
**Progress** | [**TestSign200ResponseProgress**](TestSign200ResponseProgress.md) |  | 
**Finished** | **bool** |  | 

## Methods

### NewTestSign200Response

`func NewTestSign200Response(id string, seat string, signed bool, state string, progress TestSign200ResponseProgress, finished bool, ) *TestSign200Response`

NewTestSign200Response instantiates a new TestSign200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSign200ResponseWithDefaults

`func NewTestSign200ResponseWithDefaults() *TestSign200Response`

NewTestSign200ResponseWithDefaults instantiates a new TestSign200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestSign200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestSign200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestSign200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSeat

`func (o *TestSign200Response) GetSeat() string`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *TestSign200Response) GetSeatOk() (*string, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *TestSign200Response) SetSeat(v string)`

SetSeat sets Seat field to given value.


### GetSigned

`func (o *TestSign200Response) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *TestSign200Response) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *TestSign200Response) SetSigned(v bool)`

SetSigned sets Signed field to given value.


### GetDeclined

`func (o *TestSign200Response) GetDeclined() bool`

GetDeclined returns the Declined field if non-nil, zero value otherwise.

### GetDeclinedOk

`func (o *TestSign200Response) GetDeclinedOk() (*bool, bool)`

GetDeclinedOk returns a tuple with the Declined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclined

`func (o *TestSign200Response) SetDeclined(v bool)`

SetDeclined sets Declined field to given value.

### HasDeclined

`func (o *TestSign200Response) HasDeclined() bool`

HasDeclined returns a boolean if a field has been set.

### GetState

`func (o *TestSign200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TestSign200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TestSign200Response) SetState(v string)`

SetState sets State field to given value.


### GetProgress

`func (o *TestSign200Response) GetProgress() TestSign200ResponseProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TestSign200Response) GetProgressOk() (*TestSign200ResponseProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TestSign200Response) SetProgress(v TestSign200ResponseProgress)`

SetProgress sets Progress field to given value.


### GetFinished

`func (o *TestSign200Response) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *TestSign200Response) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *TestSign200Response) SetFinished(v bool)`

SetFinished sets Finished field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SendDraft202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | **string** |  | 
**Approval** | [**SendDraft202ResponseApproval**](SendDraft202ResponseApproval.md) |  | 

## Methods

### NewSendDraft202Response

`func NewSendDraft202Response(id string, state string, approval SendDraft202ResponseApproval, ) *SendDraft202Response`

NewSendDraft202Response instantiates a new SendDraft202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendDraft202ResponseWithDefaults

`func NewSendDraft202ResponseWithDefaults() *SendDraft202Response`

NewSendDraft202ResponseWithDefaults instantiates a new SendDraft202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SendDraft202Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SendDraft202Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SendDraft202Response) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *SendDraft202Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SendDraft202Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SendDraft202Response) SetState(v string)`

SetState sets State field to given value.


### GetApproval

`func (o *SendDraft202Response) GetApproval() SendDraft202ResponseApproval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *SendDraft202Response) GetApprovalOk() (*SendDraft202ResponseApproval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *SendDraft202Response) SetApproval(v SendDraft202ResponseApproval)`

SetApproval sets Approval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



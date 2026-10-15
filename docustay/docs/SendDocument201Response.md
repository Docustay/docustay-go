# SendDocument201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | **string** |  | 
**Emailed** | **int32** | How many signing emails went out now. | 
**NotEmailed** | [**[]SendDocument201ResponseNotEmailedInner**](SendDocument201ResponseNotEmailedInner.md) | Anyone who was NOT emailed, and why (for example &#x60;limit_reached&#x60;: the daily email limit). The document is still sent; fetch their link or send a reminder later. | 

## Methods

### NewSendDocument201Response

`func NewSendDocument201Response(id string, state string, emailed int32, notEmailed []SendDocument201ResponseNotEmailedInner, ) *SendDocument201Response`

NewSendDocument201Response instantiates a new SendDocument201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendDocument201ResponseWithDefaults

`func NewSendDocument201ResponseWithDefaults() *SendDocument201Response`

NewSendDocument201ResponseWithDefaults instantiates a new SendDocument201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SendDocument201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SendDocument201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SendDocument201Response) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *SendDocument201Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SendDocument201Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SendDocument201Response) SetState(v string)`

SetState sets State field to given value.


### GetEmailed

`func (o *SendDocument201Response) GetEmailed() int32`

GetEmailed returns the Emailed field if non-nil, zero value otherwise.

### GetEmailedOk

`func (o *SendDocument201Response) GetEmailedOk() (*int32, bool)`

GetEmailedOk returns a tuple with the Emailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailed

`func (o *SendDocument201Response) SetEmailed(v int32)`

SetEmailed sets Emailed field to given value.


### GetNotEmailed

`func (o *SendDocument201Response) GetNotEmailed() []SendDocument201ResponseNotEmailedInner`

GetNotEmailed returns the NotEmailed field if non-nil, zero value otherwise.

### GetNotEmailedOk

`func (o *SendDocument201Response) GetNotEmailedOk() (*[]SendDocument201ResponseNotEmailedInner, bool)`

GetNotEmailedOk returns a tuple with the NotEmailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEmailed

`func (o *SendDocument201Response) SetNotEmailed(v []SendDocument201ResponseNotEmailedInner)`

SetNotEmailed sets NotEmailed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



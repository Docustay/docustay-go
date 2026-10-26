# ListWebhookDeliveries200ResponseDeliveriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Event** | **string** |  | 
**Attempt** | **int32** |  | 
**Status** | **string** |  | 
**ResponseStatus** | **NullableFloat32** |  | 
**Error** | **NullableString** |  | 
**AttemptedAt** | **string** |  | 
**Test** | **bool** |  | 

## Methods

### NewListWebhookDeliveries200ResponseDeliveriesInner

`func NewListWebhookDeliveries200ResponseDeliveriesInner(id string, event string, attempt int32, status string, responseStatus NullableFloat32, error_ NullableString, attemptedAt string, test bool, ) *ListWebhookDeliveries200ResponseDeliveriesInner`

NewListWebhookDeliveries200ResponseDeliveriesInner instantiates a new ListWebhookDeliveries200ResponseDeliveriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhookDeliveries200ResponseDeliveriesInnerWithDefaults

`func NewListWebhookDeliveries200ResponseDeliveriesInnerWithDefaults() *ListWebhookDeliveries200ResponseDeliveriesInner`

NewListWebhookDeliveries200ResponseDeliveriesInnerWithDefaults instantiates a new ListWebhookDeliveries200ResponseDeliveriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetAttempt

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.


### GetStatus

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResponseStatus

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetResponseStatus() float32`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetResponseStatusOk() (*float32, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetResponseStatus(v float32)`

SetResponseStatus sets ResponseStatus field to given value.


### SetResponseStatusNil

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetResponseStatusNil(b bool)`

 SetResponseStatusNil sets the value for ResponseStatus to be an explicit nil

### UnsetResponseStatus
`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) UnsetResponseStatus()`

UnsetResponseStatus ensures that no value is present for ResponseStatus, not even an explicit nil
### GetError

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetAttemptedAt

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetAttemptedAt() string`

GetAttemptedAt returns the AttemptedAt field if non-nil, zero value otherwise.

### GetAttemptedAtOk

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetAttemptedAtOk() (*string, bool)`

GetAttemptedAtOk returns a tuple with the AttemptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedAt

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetAttemptedAt(v string)`

SetAttemptedAt sets AttemptedAt field to given value.


### GetTest

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *ListWebhookDeliveries200ResponseDeliveriesInner) SetTest(v bool)`

SetTest sets Test field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FetchRelayEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]FetchRelayEvents200ResponseEventsInner**](FetchRelayEvents200ResponseEventsInner.md) |  | 
**Cursor** | **NullableString** |  | 

## Methods

### NewFetchRelayEvents200Response

`func NewFetchRelayEvents200Response(events []FetchRelayEvents200ResponseEventsInner, cursor NullableString, ) *FetchRelayEvents200Response`

NewFetchRelayEvents200Response instantiates a new FetchRelayEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchRelayEvents200ResponseWithDefaults

`func NewFetchRelayEvents200ResponseWithDefaults() *FetchRelayEvents200Response`

NewFetchRelayEvents200ResponseWithDefaults instantiates a new FetchRelayEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *FetchRelayEvents200Response) GetEvents() []FetchRelayEvents200ResponseEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FetchRelayEvents200Response) GetEventsOk() (*[]FetchRelayEvents200ResponseEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FetchRelayEvents200Response) SetEvents(v []FetchRelayEvents200ResponseEventsInner)`

SetEvents sets Events field to given value.


### GetCursor

`func (o *FetchRelayEvents200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FetchRelayEvents200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FetchRelayEvents200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### SetCursorNil

`func (o *FetchRelayEvents200Response) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *FetchRelayEvents200Response) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



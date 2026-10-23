# CreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Relay** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EventTypes** | **[]string** | From &#x60;GET /api/v1/webhooks/topics&#x60;. | 

## Methods

### NewCreateWebhookRequest

`func NewCreateWebhookRequest(eventTypes []string, ) *CreateWebhookRequest`

NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookRequestWithDefaults

`func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest`

NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateWebhookRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRelay

`func (o *CreateWebhookRequest) GetRelay() bool`

GetRelay returns the Relay field if non-nil, zero value otherwise.

### GetRelayOk

`func (o *CreateWebhookRequest) GetRelayOk() (*bool, bool)`

GetRelayOk returns a tuple with the Relay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelay

`func (o *CreateWebhookRequest) SetRelay(v bool)`

SetRelay sets Relay field to given value.

### HasRelay

`func (o *CreateWebhookRequest) HasRelay() bool`

HasRelay returns a boolean if a field has been set.

### GetName

`func (o *CreateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEventTypes

`func (o *CreateWebhookRequest) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *CreateWebhookRequest) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *CreateWebhookRequest) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



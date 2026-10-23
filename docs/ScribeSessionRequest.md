# ScribeSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | **string** | The website that will show Scribe; must be on your allowed list. | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewScribeSessionRequest

`func NewScribeSessionRequest(origin string, ) *ScribeSessionRequest`

NewScribeSessionRequest instantiates a new ScribeSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScribeSessionRequestWithDefaults

`func NewScribeSessionRequestWithDefaults() *ScribeSessionRequest`

NewScribeSessionRequestWithDefaults instantiates a new ScribeSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *ScribeSessionRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ScribeSessionRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ScribeSessionRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetName

`func (o *ScribeSessionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScribeSessionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScribeSessionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScribeSessionRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



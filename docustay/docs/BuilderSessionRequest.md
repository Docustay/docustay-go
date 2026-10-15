# BuilderSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | **string** | The website that will show the builder; must be on your allowed list. | 
**Name** | Pointer to **string** | A name to start the template with. | [optional] 

## Methods

### NewBuilderSessionRequest

`func NewBuilderSessionRequest(origin string, ) *BuilderSessionRequest`

NewBuilderSessionRequest instantiates a new BuilderSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuilderSessionRequestWithDefaults

`func NewBuilderSessionRequestWithDefaults() *BuilderSessionRequest`

NewBuilderSessionRequestWithDefaults instantiates a new BuilderSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *BuilderSessionRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *BuilderSessionRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *BuilderSessionRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetName

`func (o *BuilderSessionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuilderSessionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuilderSessionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BuilderSessionRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



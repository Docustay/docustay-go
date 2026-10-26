# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **NullableString** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Pages** | **NullableInt32** |  | 
**Seats** | [**[]TemplateSeatsInner**](TemplateSeatsInner.md) |  | 

## Methods

### NewTemplate

`func NewTemplate(id string, name string, pages NullableInt32, seats []TemplateSeatsInner, ) *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *Template) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Template) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Template) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Template) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Template) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Template) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetMetadata

`func (o *Template) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Template) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Template) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Template) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *Template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Template) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.


### GetPages

`func (o *Template) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *Template) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *Template) SetPages(v int32)`

SetPages sets Pages field to given value.


### SetPagesNil

`func (o *Template) SetPagesNil(b bool)`

 SetPagesNil sets the value for Pages to be an explicit nil

### UnsetPages
`func (o *Template) UnsetPages()`

UnsetPages ensures that no value is present for Pages, not even an explicit nil
### GetSeats

`func (o *Template) GetSeats() []TemplateSeatsInner`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *Template) GetSeatsOk() (*[]TemplateSeatsInner, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *Template) SetSeats(v []TemplateSeatsInner)`

SetSeats sets Seats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



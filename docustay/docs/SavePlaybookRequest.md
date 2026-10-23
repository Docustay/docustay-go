# SavePlaybookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Plays** | Pointer to **[]string** |  | [optional] 
**Templates** | Pointer to **[]string** | &#x60;builtin:&lt;slug&gt;&#x60; for a library template, or one of your templates&#39; family ids. | [optional] 

## Methods

### NewSavePlaybookRequest

`func NewSavePlaybookRequest(name string, title string, ) *SavePlaybookRequest`

NewSavePlaybookRequest instantiates a new SavePlaybookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavePlaybookRequestWithDefaults

`func NewSavePlaybookRequestWithDefaults() *SavePlaybookRequest`

NewSavePlaybookRequestWithDefaults instantiates a new SavePlaybookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SavePlaybookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavePlaybookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavePlaybookRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *SavePlaybookRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SavePlaybookRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SavePlaybookRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SavePlaybookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SavePlaybookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SavePlaybookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SavePlaybookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlays

`func (o *SavePlaybookRequest) GetPlays() []string`

GetPlays returns the Plays field if non-nil, zero value otherwise.

### GetPlaysOk

`func (o *SavePlaybookRequest) GetPlaysOk() (*[]string, bool)`

GetPlaysOk returns a tuple with the Plays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlays

`func (o *SavePlaybookRequest) SetPlays(v []string)`

SetPlays sets Plays field to given value.

### HasPlays

`func (o *SavePlaybookRequest) HasPlays() bool`

HasPlays returns a boolean if a field has been set.

### GetTemplates

`func (o *SavePlaybookRequest) GetTemplates() []string`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *SavePlaybookRequest) GetTemplatesOk() (*[]string, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *SavePlaybookRequest) SetTemplates(v []string)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *SavePlaybookRequest) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



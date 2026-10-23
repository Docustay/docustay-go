# ListTeams200ResponseTeamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Members** | [**[]ListTeams200ResponseTeamsInnerMembersInner**](ListTeams200ResponseTeamsInnerMembersInner.md) |  | 
**Documents** | **int32** |  | 
**Templates** | **int32** |  | 

## Methods

### NewListTeams200ResponseTeamsInner

`func NewListTeams200ResponseTeamsInner(id string, name string, members []ListTeams200ResponseTeamsInnerMembersInner, documents int32, templates int32, ) *ListTeams200ResponseTeamsInner`

NewListTeams200ResponseTeamsInner instantiates a new ListTeams200ResponseTeamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTeams200ResponseTeamsInnerWithDefaults

`func NewListTeams200ResponseTeamsInnerWithDefaults() *ListTeams200ResponseTeamsInner`

NewListTeams200ResponseTeamsInnerWithDefaults instantiates a new ListTeams200ResponseTeamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTeams200ResponseTeamsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTeams200ResponseTeamsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTeams200ResponseTeamsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListTeams200ResponseTeamsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTeams200ResponseTeamsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTeams200ResponseTeamsInner) SetName(v string)`

SetName sets Name field to given value.


### GetMembers

`func (o *ListTeams200ResponseTeamsInner) GetMembers() []ListTeams200ResponseTeamsInnerMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ListTeams200ResponseTeamsInner) GetMembersOk() (*[]ListTeams200ResponseTeamsInnerMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ListTeams200ResponseTeamsInner) SetMembers(v []ListTeams200ResponseTeamsInnerMembersInner)`

SetMembers sets Members field to given value.


### GetDocuments

`func (o *ListTeams200ResponseTeamsInner) GetDocuments() int32`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ListTeams200ResponseTeamsInner) GetDocumentsOk() (*int32, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ListTeams200ResponseTeamsInner) SetDocuments(v int32)`

SetDocuments sets Documents field to given value.


### GetTemplates

`func (o *ListTeams200ResponseTeamsInner) GetTemplates() int32`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *ListTeams200ResponseTeamsInner) GetTemplatesOk() (*int32, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *ListTeams200ResponseTeamsInner) SetTemplates(v int32)`

SetTemplates sets Templates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



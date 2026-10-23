# CreateTemplateFromSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 
**ExternalId** | Pointer to **string** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Name** | **string** |  | 
**Source** | **map[string]interface{}** | The document source: a ProseMirror/TipTap JSON document (&#x60;{type:\&quot;doc\&quot;,content:[…]}&#x60;) using headings, paragraphs, lists, tables, &#x60;fieldTag&#x60;, &#x60;signatureBlock&#x60;, &#x60;variableChip&#x60;. See the guide on document sources. | 
**Profile** | Pointer to **string** |  | [optional] 
**Paper** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateTemplateFromSourceRequest

`func NewCreateTemplateFromSourceRequest(name string, source map[string]interface{}, ) *CreateTemplateFromSourceRequest`

NewCreateTemplateFromSourceRequest instantiates a new CreateTemplateFromSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTemplateFromSourceRequestWithDefaults

`func NewCreateTemplateFromSourceRequestWithDefaults() *CreateTemplateFromSourceRequest`

NewCreateTemplateFromSourceRequestWithDefaults instantiates a new CreateTemplateFromSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CreateTemplateFromSourceRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateTemplateFromSourceRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateTemplateFromSourceRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateTemplateFromSourceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateTemplateFromSourceRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateTemplateFromSourceRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateTemplateFromSourceRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateTemplateFromSourceRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetName

`func (o *CreateTemplateFromSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTemplateFromSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTemplateFromSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *CreateTemplateFromSourceRequest) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateTemplateFromSourceRequest) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateTemplateFromSourceRequest) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.


### GetProfile

`func (o *CreateTemplateFromSourceRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreateTemplateFromSourceRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreateTemplateFromSourceRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CreateTemplateFromSourceRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetPaper

`func (o *CreateTemplateFromSourceRequest) GetPaper() string`

GetPaper returns the Paper field if non-nil, zero value otherwise.

### GetPaperOk

`func (o *CreateTemplateFromSourceRequest) GetPaperOk() (*string, bool)`

GetPaperOk returns a tuple with the Paper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaper

`func (o *CreateTemplateFromSourceRequest) SetPaper(v string)`

SetPaper sets Paper field to given value.

### HasPaper

`func (o *CreateTemplateFromSourceRequest) HasPaper() bool`

HasPaper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



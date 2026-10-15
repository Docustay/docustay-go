# FileTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 
**ExternalId** | Pointer to **string** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Name** | **string** |  | 
**FileName** | **string** |  | 
**File** | **string** | The file, base64-encoded: a PDF file with text tags written in it. | 
**Fields** | Pointer to **string** | &#x60;tags&#x60; (default) places fields only where a text tag is written. &#x60;auto&#x60; also looks at the page itself — lines, boxes, tick boxes, highlighted bars and the words beside them — and places the fields it is confident about, with the people grouped into seats. | [optional] 
**MinConfidence** | Pointer to **float32** | With &#x60;fields: auto&#x60;: the least sure a field may be and still be placed (default 0.6). | [optional] 

## Methods

### NewFileTemplateRequest

`func NewFileTemplateRequest(name string, fileName string, file string, ) *FileTemplateRequest`

NewFileTemplateRequest instantiates a new FileTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileTemplateRequestWithDefaults

`func NewFileTemplateRequestWithDefaults() *FileTemplateRequest`

NewFileTemplateRequestWithDefaults instantiates a new FileTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *FileTemplateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FileTemplateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FileTemplateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FileTemplateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExternalId

`func (o *FileTemplateRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *FileTemplateRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *FileTemplateRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *FileTemplateRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetName

`func (o *FileTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFileName

`func (o *FileTemplateRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FileTemplateRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FileTemplateRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFile

`func (o *FileTemplateRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileTemplateRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileTemplateRequest) SetFile(v string)`

SetFile sets File field to given value.


### GetFields

`func (o *FileTemplateRequest) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FileTemplateRequest) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FileTemplateRequest) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FileTemplateRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMinConfidence

`func (o *FileTemplateRequest) GetMinConfidence() float32`

GetMinConfidence returns the MinConfidence field if non-nil, zero value otherwise.

### GetMinConfidenceOk

`func (o *FileTemplateRequest) GetMinConfidenceOk() (*float32, bool)`

GetMinConfidenceOk returns a tuple with the MinConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinConfidence

`func (o *FileTemplateRequest) SetMinConfidence(v float32)`

SetMinConfidence sets MinConfidence field to given value.

### HasMinConfidence

`func (o *FileTemplateRequest) HasMinConfidence() bool`

HasMinConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



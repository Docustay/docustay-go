# HtmlTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer. | [optional] 
**ExternalId** | Pointer to **string** | Your own id for this thing, so you can match it to your records. Filter lists with &#x60;externalId&#x3D;&#x60;. | [optional] 
**Name** | **string** |  | 
**Html** | **string** | HTML with text tags such as {{Tenant signature;type&#x3D;signature;role&#x3D;Tenant}}. Scripts, styles, frames, forms and remote addresses are removed; pictures must be data: addresses. Up to 2 MB. | 
**Paper** | Pointer to **string** |  | [optional] 

## Methods

### NewHtmlTemplateRequest

`func NewHtmlTemplateRequest(name string, html string, ) *HtmlTemplateRequest`

NewHtmlTemplateRequest instantiates a new HtmlTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHtmlTemplateRequestWithDefaults

`func NewHtmlTemplateRequestWithDefaults() *HtmlTemplateRequest`

NewHtmlTemplateRequestWithDefaults instantiates a new HtmlTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *HtmlTemplateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HtmlTemplateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HtmlTemplateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HtmlTemplateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExternalId

`func (o *HtmlTemplateRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *HtmlTemplateRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *HtmlTemplateRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *HtmlTemplateRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetName

`func (o *HtmlTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HtmlTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HtmlTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHtml

`func (o *HtmlTemplateRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *HtmlTemplateRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *HtmlTemplateRequest) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetPaper

`func (o *HtmlTemplateRequest) GetPaper() string`

GetPaper returns the Paper field if non-nil, zero value otherwise.

### GetPaperOk

`func (o *HtmlTemplateRequest) GetPaperOk() (*string, bool)`

GetPaperOk returns a tuple with the Paper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaper

`func (o *HtmlTemplateRequest) SetPaper(v string)`

SetPaper sets Paper field to given value.

### HasPaper

`func (o *HtmlTemplateRequest) HasPaper() bool`

HasPaper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



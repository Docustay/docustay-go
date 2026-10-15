# DetectFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** |  | [optional] 
**File** | **string** | The file, base64-encoded: a PDF file. | 
**MinConfidence** | Pointer to **float32** |  | [optional] 
**Ocr** | Pointer to **bool** | Read pages that are only a picture by character recognition (default true; English). | [optional] 

## Methods

### NewDetectFieldsRequest

`func NewDetectFieldsRequest(file string, ) *DetectFieldsRequest`

NewDetectFieldsRequest instantiates a new DetectFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectFieldsRequestWithDefaults

`func NewDetectFieldsRequestWithDefaults() *DetectFieldsRequest`

NewDetectFieldsRequestWithDefaults instantiates a new DetectFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *DetectFieldsRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DetectFieldsRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DetectFieldsRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DetectFieldsRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFile

`func (o *DetectFieldsRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *DetectFieldsRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *DetectFieldsRequest) SetFile(v string)`

SetFile sets File field to given value.


### GetMinConfidence

`func (o *DetectFieldsRequest) GetMinConfidence() float32`

GetMinConfidence returns the MinConfidence field if non-nil, zero value otherwise.

### GetMinConfidenceOk

`func (o *DetectFieldsRequest) GetMinConfidenceOk() (*float32, bool)`

GetMinConfidenceOk returns a tuple with the MinConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinConfidence

`func (o *DetectFieldsRequest) SetMinConfidence(v float32)`

SetMinConfidence sets MinConfidence field to given value.

### HasMinConfidence

`func (o *DetectFieldsRequest) HasMinConfidence() bool`

HasMinConfidence returns a boolean if a field has been set.

### GetOcr

`func (o *DetectFieldsRequest) GetOcr() bool`

GetOcr returns the Ocr field if non-nil, zero value otherwise.

### GetOcrOk

`func (o *DetectFieldsRequest) GetOcrOk() (*bool, bool)`

GetOcrOk returns a tuple with the Ocr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcr

`func (o *DetectFieldsRequest) SetOcr(v bool)`

SetOcr sets Ocr field to given value.

### HasOcr

`func (o *DetectFieldsRequest) HasOcr() bool`

HasOcr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



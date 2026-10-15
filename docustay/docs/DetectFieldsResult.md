# DetectFieldsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]DetectedField**](DetectedField.md) |  | 
**Signers** | [**[]TemplateSeatsInner**](TemplateSeatsInner.md) |  | 
**Pages** | [**[]DetectFieldsResultPagesInner**](DetectFieldsResultPagesInner.md) |  | 
**Documents** | [**[]DetectFieldsResultDocumentsInner**](DetectFieldsResultDocumentsInner.md) |  | 
**Summary** | [**DetectFieldsResultSummary**](DetectFieldsResultSummary.md) |  | 
**Notes** | **[]string** |  | 

## Methods

### NewDetectFieldsResult

`func NewDetectFieldsResult(fields []DetectedField, signers []TemplateSeatsInner, pages []DetectFieldsResultPagesInner, documents []DetectFieldsResultDocumentsInner, summary DetectFieldsResultSummary, notes []string, ) *DetectFieldsResult`

NewDetectFieldsResult instantiates a new DetectFieldsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectFieldsResultWithDefaults

`func NewDetectFieldsResultWithDefaults() *DetectFieldsResult`

NewDetectFieldsResultWithDefaults instantiates a new DetectFieldsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *DetectFieldsResult) GetFields() []DetectedField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DetectFieldsResult) GetFieldsOk() (*[]DetectedField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DetectFieldsResult) SetFields(v []DetectedField)`

SetFields sets Fields field to given value.


### GetSigners

`func (o *DetectFieldsResult) GetSigners() []TemplateSeatsInner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *DetectFieldsResult) GetSignersOk() (*[]TemplateSeatsInner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *DetectFieldsResult) SetSigners(v []TemplateSeatsInner)`

SetSigners sets Signers field to given value.


### GetPages

`func (o *DetectFieldsResult) GetPages() []DetectFieldsResultPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *DetectFieldsResult) GetPagesOk() (*[]DetectFieldsResultPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *DetectFieldsResult) SetPages(v []DetectFieldsResultPagesInner)`

SetPages sets Pages field to given value.


### GetDocuments

`func (o *DetectFieldsResult) GetDocuments() []DetectFieldsResultDocumentsInner`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DetectFieldsResult) GetDocumentsOk() (*[]DetectFieldsResultDocumentsInner, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DetectFieldsResult) SetDocuments(v []DetectFieldsResultDocumentsInner)`

SetDocuments sets Documents field to given value.


### GetSummary

`func (o *DetectFieldsResult) GetSummary() DetectFieldsResultSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DetectFieldsResult) GetSummaryOk() (*DetectFieldsResultSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DetectFieldsResult) SetSummary(v DetectFieldsResultSummary)`

SetSummary sets Summary field to given value.


### GetNotes

`func (o *DetectFieldsResult) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DetectFieldsResult) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DetectFieldsResult) SetNotes(v []string)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



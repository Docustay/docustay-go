# WizardBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] 
**Answers** | Pointer to **map[string]string** |  | [optional] 
**Final** | **string** |  | 
**Payment** | Pointer to [**WizardBriefPayment**](WizardBriefPayment.md) |  | [optional] 

## Methods

### NewWizardBrief

`func NewWizardBrief(type_ string, final string, ) *WizardBrief`

NewWizardBrief instantiates a new WizardBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardBriefWithDefaults

`func NewWizardBriefWithDefaults() *WizardBrief`

NewWizardBriefWithDefaults instantiates a new WizardBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WizardBrief) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WizardBrief) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WizardBrief) SetType(v string)`

SetType sets Type field to given value.


### GetKind

`func (o *WizardBrief) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WizardBrief) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WizardBrief) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WizardBrief) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAnswers

`func (o *WizardBrief) GetAnswers() map[string]string`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *WizardBrief) GetAnswersOk() (*map[string]string, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *WizardBrief) SetAnswers(v map[string]string)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *WizardBrief) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetFinal

`func (o *WizardBrief) GetFinal() string`

GetFinal returns the Final field if non-nil, zero value otherwise.

### GetFinalOk

`func (o *WizardBrief) GetFinalOk() (*string, bool)`

GetFinalOk returns a tuple with the Final field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinal

`func (o *WizardBrief) SetFinal(v string)`

SetFinal sets Final field to given value.


### GetPayment

`func (o *WizardBrief) GetPayment() WizardBriefPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *WizardBrief) GetPaymentOk() (*WizardBriefPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *WizardBrief) SetPayment(v WizardBriefPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *WizardBrief) HasPayment() bool`

HasPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



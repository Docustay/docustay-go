# UsageLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentsPerPeriod** | **NullableInt32** |  | 
**PerDocumentMinor** | **NullableInt32** |  | 
**DraftsPerPeriod** | **NullableInt32** |  | 
**PerDraftMinor** | **NullableInt32** |  | 
**TemplatesPerPeriod** | **NullableInt32** |  | 
**SignersPerPeriod** | **NullableInt32** |  | 
**People** | **NullableInt32** |  | 
**PersonMonthlyMinor** | **NullableInt32** |  | 
**UploadMb** | **NullableInt32** |  | 
**EmailsPerDay** | **NullableInt32** | The plan&#39;s daily email limit. Today&#39;s actual limit is emails.limitPerDay, which is lower in a workspace&#39;s first 7 days. | 
**BulkSendMax** | **NullableInt32** |  | 
**PaymentsFeePercent** | **float32** |  | 

## Methods

### NewUsageLimits

`func NewUsageLimits(documentsPerPeriod NullableInt32, perDocumentMinor NullableInt32, draftsPerPeriod NullableInt32, perDraftMinor NullableInt32, templatesPerPeriod NullableInt32, signersPerPeriod NullableInt32, people NullableInt32, personMonthlyMinor NullableInt32, uploadMb NullableInt32, emailsPerDay NullableInt32, bulkSendMax NullableInt32, paymentsFeePercent float32, ) *UsageLimits`

NewUsageLimits instantiates a new UsageLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLimitsWithDefaults

`func NewUsageLimitsWithDefaults() *UsageLimits`

NewUsageLimitsWithDefaults instantiates a new UsageLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentsPerPeriod

`func (o *UsageLimits) GetDocumentsPerPeriod() int32`

GetDocumentsPerPeriod returns the DocumentsPerPeriod field if non-nil, zero value otherwise.

### GetDocumentsPerPeriodOk

`func (o *UsageLimits) GetDocumentsPerPeriodOk() (*int32, bool)`

GetDocumentsPerPeriodOk returns a tuple with the DocumentsPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentsPerPeriod

`func (o *UsageLimits) SetDocumentsPerPeriod(v int32)`

SetDocumentsPerPeriod sets DocumentsPerPeriod field to given value.


### SetDocumentsPerPeriodNil

`func (o *UsageLimits) SetDocumentsPerPeriodNil(b bool)`

 SetDocumentsPerPeriodNil sets the value for DocumentsPerPeriod to be an explicit nil

### UnsetDocumentsPerPeriod
`func (o *UsageLimits) UnsetDocumentsPerPeriod()`

UnsetDocumentsPerPeriod ensures that no value is present for DocumentsPerPeriod, not even an explicit nil
### GetPerDocumentMinor

`func (o *UsageLimits) GetPerDocumentMinor() int32`

GetPerDocumentMinor returns the PerDocumentMinor field if non-nil, zero value otherwise.

### GetPerDocumentMinorOk

`func (o *UsageLimits) GetPerDocumentMinorOk() (*int32, bool)`

GetPerDocumentMinorOk returns a tuple with the PerDocumentMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerDocumentMinor

`func (o *UsageLimits) SetPerDocumentMinor(v int32)`

SetPerDocumentMinor sets PerDocumentMinor field to given value.


### SetPerDocumentMinorNil

`func (o *UsageLimits) SetPerDocumentMinorNil(b bool)`

 SetPerDocumentMinorNil sets the value for PerDocumentMinor to be an explicit nil

### UnsetPerDocumentMinor
`func (o *UsageLimits) UnsetPerDocumentMinor()`

UnsetPerDocumentMinor ensures that no value is present for PerDocumentMinor, not even an explicit nil
### GetDraftsPerPeriod

`func (o *UsageLimits) GetDraftsPerPeriod() int32`

GetDraftsPerPeriod returns the DraftsPerPeriod field if non-nil, zero value otherwise.

### GetDraftsPerPeriodOk

`func (o *UsageLimits) GetDraftsPerPeriodOk() (*int32, bool)`

GetDraftsPerPeriodOk returns a tuple with the DraftsPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftsPerPeriod

`func (o *UsageLimits) SetDraftsPerPeriod(v int32)`

SetDraftsPerPeriod sets DraftsPerPeriod field to given value.


### SetDraftsPerPeriodNil

`func (o *UsageLimits) SetDraftsPerPeriodNil(b bool)`

 SetDraftsPerPeriodNil sets the value for DraftsPerPeriod to be an explicit nil

### UnsetDraftsPerPeriod
`func (o *UsageLimits) UnsetDraftsPerPeriod()`

UnsetDraftsPerPeriod ensures that no value is present for DraftsPerPeriod, not even an explicit nil
### GetPerDraftMinor

`func (o *UsageLimits) GetPerDraftMinor() int32`

GetPerDraftMinor returns the PerDraftMinor field if non-nil, zero value otherwise.

### GetPerDraftMinorOk

`func (o *UsageLimits) GetPerDraftMinorOk() (*int32, bool)`

GetPerDraftMinorOk returns a tuple with the PerDraftMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerDraftMinor

`func (o *UsageLimits) SetPerDraftMinor(v int32)`

SetPerDraftMinor sets PerDraftMinor field to given value.


### SetPerDraftMinorNil

`func (o *UsageLimits) SetPerDraftMinorNil(b bool)`

 SetPerDraftMinorNil sets the value for PerDraftMinor to be an explicit nil

### UnsetPerDraftMinor
`func (o *UsageLimits) UnsetPerDraftMinor()`

UnsetPerDraftMinor ensures that no value is present for PerDraftMinor, not even an explicit nil
### GetTemplatesPerPeriod

`func (o *UsageLimits) GetTemplatesPerPeriod() int32`

GetTemplatesPerPeriod returns the TemplatesPerPeriod field if non-nil, zero value otherwise.

### GetTemplatesPerPeriodOk

`func (o *UsageLimits) GetTemplatesPerPeriodOk() (*int32, bool)`

GetTemplatesPerPeriodOk returns a tuple with the TemplatesPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatesPerPeriod

`func (o *UsageLimits) SetTemplatesPerPeriod(v int32)`

SetTemplatesPerPeriod sets TemplatesPerPeriod field to given value.


### SetTemplatesPerPeriodNil

`func (o *UsageLimits) SetTemplatesPerPeriodNil(b bool)`

 SetTemplatesPerPeriodNil sets the value for TemplatesPerPeriod to be an explicit nil

### UnsetTemplatesPerPeriod
`func (o *UsageLimits) UnsetTemplatesPerPeriod()`

UnsetTemplatesPerPeriod ensures that no value is present for TemplatesPerPeriod, not even an explicit nil
### GetSignersPerPeriod

`func (o *UsageLimits) GetSignersPerPeriod() int32`

GetSignersPerPeriod returns the SignersPerPeriod field if non-nil, zero value otherwise.

### GetSignersPerPeriodOk

`func (o *UsageLimits) GetSignersPerPeriodOk() (*int32, bool)`

GetSignersPerPeriodOk returns a tuple with the SignersPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignersPerPeriod

`func (o *UsageLimits) SetSignersPerPeriod(v int32)`

SetSignersPerPeriod sets SignersPerPeriod field to given value.


### SetSignersPerPeriodNil

`func (o *UsageLimits) SetSignersPerPeriodNil(b bool)`

 SetSignersPerPeriodNil sets the value for SignersPerPeriod to be an explicit nil

### UnsetSignersPerPeriod
`func (o *UsageLimits) UnsetSignersPerPeriod()`

UnsetSignersPerPeriod ensures that no value is present for SignersPerPeriod, not even an explicit nil
### GetPeople

`func (o *UsageLimits) GetPeople() int32`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *UsageLimits) GetPeopleOk() (*int32, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *UsageLimits) SetPeople(v int32)`

SetPeople sets People field to given value.


### SetPeopleNil

`func (o *UsageLimits) SetPeopleNil(b bool)`

 SetPeopleNil sets the value for People to be an explicit nil

### UnsetPeople
`func (o *UsageLimits) UnsetPeople()`

UnsetPeople ensures that no value is present for People, not even an explicit nil
### GetPersonMonthlyMinor

`func (o *UsageLimits) GetPersonMonthlyMinor() int32`

GetPersonMonthlyMinor returns the PersonMonthlyMinor field if non-nil, zero value otherwise.

### GetPersonMonthlyMinorOk

`func (o *UsageLimits) GetPersonMonthlyMinorOk() (*int32, bool)`

GetPersonMonthlyMinorOk returns a tuple with the PersonMonthlyMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonMonthlyMinor

`func (o *UsageLimits) SetPersonMonthlyMinor(v int32)`

SetPersonMonthlyMinor sets PersonMonthlyMinor field to given value.


### SetPersonMonthlyMinorNil

`func (o *UsageLimits) SetPersonMonthlyMinorNil(b bool)`

 SetPersonMonthlyMinorNil sets the value for PersonMonthlyMinor to be an explicit nil

### UnsetPersonMonthlyMinor
`func (o *UsageLimits) UnsetPersonMonthlyMinor()`

UnsetPersonMonthlyMinor ensures that no value is present for PersonMonthlyMinor, not even an explicit nil
### GetUploadMb

`func (o *UsageLimits) GetUploadMb() int32`

GetUploadMb returns the UploadMb field if non-nil, zero value otherwise.

### GetUploadMbOk

`func (o *UsageLimits) GetUploadMbOk() (*int32, bool)`

GetUploadMbOk returns a tuple with the UploadMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadMb

`func (o *UsageLimits) SetUploadMb(v int32)`

SetUploadMb sets UploadMb field to given value.


### SetUploadMbNil

`func (o *UsageLimits) SetUploadMbNil(b bool)`

 SetUploadMbNil sets the value for UploadMb to be an explicit nil

### UnsetUploadMb
`func (o *UsageLimits) UnsetUploadMb()`

UnsetUploadMb ensures that no value is present for UploadMb, not even an explicit nil
### GetEmailsPerDay

`func (o *UsageLimits) GetEmailsPerDay() int32`

GetEmailsPerDay returns the EmailsPerDay field if non-nil, zero value otherwise.

### GetEmailsPerDayOk

`func (o *UsageLimits) GetEmailsPerDayOk() (*int32, bool)`

GetEmailsPerDayOk returns a tuple with the EmailsPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsPerDay

`func (o *UsageLimits) SetEmailsPerDay(v int32)`

SetEmailsPerDay sets EmailsPerDay field to given value.


### SetEmailsPerDayNil

`func (o *UsageLimits) SetEmailsPerDayNil(b bool)`

 SetEmailsPerDayNil sets the value for EmailsPerDay to be an explicit nil

### UnsetEmailsPerDay
`func (o *UsageLimits) UnsetEmailsPerDay()`

UnsetEmailsPerDay ensures that no value is present for EmailsPerDay, not even an explicit nil
### GetBulkSendMax

`func (o *UsageLimits) GetBulkSendMax() int32`

GetBulkSendMax returns the BulkSendMax field if non-nil, zero value otherwise.

### GetBulkSendMaxOk

`func (o *UsageLimits) GetBulkSendMaxOk() (*int32, bool)`

GetBulkSendMaxOk returns a tuple with the BulkSendMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendMax

`func (o *UsageLimits) SetBulkSendMax(v int32)`

SetBulkSendMax sets BulkSendMax field to given value.


### SetBulkSendMaxNil

`func (o *UsageLimits) SetBulkSendMaxNil(b bool)`

 SetBulkSendMaxNil sets the value for BulkSendMax to be an explicit nil

### UnsetBulkSendMax
`func (o *UsageLimits) UnsetBulkSendMax()`

UnsetBulkSendMax ensures that no value is present for BulkSendMax, not even an explicit nil
### GetPaymentsFeePercent

`func (o *UsageLimits) GetPaymentsFeePercent() float32`

GetPaymentsFeePercent returns the PaymentsFeePercent field if non-nil, zero value otherwise.

### GetPaymentsFeePercentOk

`func (o *UsageLimits) GetPaymentsFeePercentOk() (*float32, bool)`

GetPaymentsFeePercentOk returns a tuple with the PaymentsFeePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsFeePercent

`func (o *UsageLimits) SetPaymentsFeePercent(v float32)`

SetPaymentsFeePercent sets PaymentsFeePercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



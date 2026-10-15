# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FieldId** | **string** |  | 
**Person** | **string** |  | 
**Status** | **string** | &#x60;paid&#x60; once Stripe says so. A refund moves it to &#x60;partially_refunded&#x60; or &#x60;refunded&#x60;. | 
**AmountMinor** | **int32** | What was charged, in the smallest unit of &#x60;currency&#x60; (cents). Worked out on the server; never taken from the browser. | 
**Currency** | **string** |  | 
**PaidAt** | **string** |  | 
**ReceiptUrl** | **string** | Stripe&#39;s own receipt for the payment. | 
**RefundedMinor** | **int32** |  | 
**Fee** | [**PaymentFee**](PaymentFee.md) |  | 
**Test** | **bool** |  | 

## Methods

### NewPayment

`func NewPayment(id string, fieldId string, person string, status string, amountMinor int32, currency string, paidAt string, receiptUrl string, refundedMinor int32, fee PaymentFee, test bool, ) *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v string)`

SetId sets Id field to given value.


### GetFieldId

`func (o *Payment) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *Payment) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *Payment) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetPerson

`func (o *Payment) GetPerson() string`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *Payment) GetPersonOk() (*string, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *Payment) SetPerson(v string)`

SetPerson sets Person field to given value.


### GetStatus

`func (o *Payment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAmountMinor

`func (o *Payment) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *Payment) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *Payment) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *Payment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaidAt

`func (o *Payment) GetPaidAt() string`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *Payment) GetPaidAtOk() (*string, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *Payment) SetPaidAt(v string)`

SetPaidAt sets PaidAt field to given value.


### GetReceiptUrl

`func (o *Payment) GetReceiptUrl() string`

GetReceiptUrl returns the ReceiptUrl field if non-nil, zero value otherwise.

### GetReceiptUrlOk

`func (o *Payment) GetReceiptUrlOk() (*string, bool)`

GetReceiptUrlOk returns a tuple with the ReceiptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptUrl

`func (o *Payment) SetReceiptUrl(v string)`

SetReceiptUrl sets ReceiptUrl field to given value.


### GetRefundedMinor

`func (o *Payment) GetRefundedMinor() int32`

GetRefundedMinor returns the RefundedMinor field if non-nil, zero value otherwise.

### GetRefundedMinorOk

`func (o *Payment) GetRefundedMinorOk() (*int32, bool)`

GetRefundedMinorOk returns a tuple with the RefundedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedMinor

`func (o *Payment) SetRefundedMinor(v int32)`

SetRefundedMinor sets RefundedMinor field to given value.


### GetFee

`func (o *Payment) GetFee() PaymentFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *Payment) GetFeeOk() (*PaymentFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *Payment) SetFee(v PaymentFee)`

SetFee sets Fee field to given value.


### GetTest

`func (o *Payment) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *Payment) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *Payment) SetTest(v bool)`

SetTest sets Test field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



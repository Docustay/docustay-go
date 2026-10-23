# PaymentFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasisPoints** | **int32** | Docustay&#39;s fee, decided when the payment started: 200 (2%) on the Free plan, 0 on Pro. | 
**AmountMinor** | **int32** |  | 
**RefundedMinor** | **int32** | How much of the fee went back with refunds. | 
**Plan** | **string** |  | 

## Methods

### NewPaymentFee

`func NewPaymentFee(basisPoints int32, amountMinor int32, refundedMinor int32, plan string, ) *PaymentFee`

NewPaymentFee instantiates a new PaymentFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentFeeWithDefaults

`func NewPaymentFeeWithDefaults() *PaymentFee`

NewPaymentFeeWithDefaults instantiates a new PaymentFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasisPoints

`func (o *PaymentFee) GetBasisPoints() int32`

GetBasisPoints returns the BasisPoints field if non-nil, zero value otherwise.

### GetBasisPointsOk

`func (o *PaymentFee) GetBasisPointsOk() (*int32, bool)`

GetBasisPointsOk returns a tuple with the BasisPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisPoints

`func (o *PaymentFee) SetBasisPoints(v int32)`

SetBasisPoints sets BasisPoints field to given value.


### GetAmountMinor

`func (o *PaymentFee) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *PaymentFee) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *PaymentFee) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetRefundedMinor

`func (o *PaymentFee) GetRefundedMinor() int32`

GetRefundedMinor returns the RefundedMinor field if non-nil, zero value otherwise.

### GetRefundedMinorOk

`func (o *PaymentFee) GetRefundedMinorOk() (*int32, bool)`

GetRefundedMinorOk returns a tuple with the RefundedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedMinor

`func (o *PaymentFee) SetRefundedMinor(v int32)`

SetRefundedMinor sets RefundedMinor field to given value.


### GetPlan

`func (o *PaymentFee) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PaymentFee) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PaymentFee) SetPlan(v string)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



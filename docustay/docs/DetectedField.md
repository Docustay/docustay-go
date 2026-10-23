# DetectedField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Page** | **int32** |  | 
**X** | **float32** |  | 
**Y** | **float32** |  | 
**W** | **float32** |  | 
**H** | **float32** |  | 
**Confidence** | **float32** |  | 
**Source** | **string** | &#x60;native&#x60; (already a form field), &#x60;line&#x60;, &#x60;text&#x60;, &#x60;box&#x60;, &#x60;table&#x60;, &#x60;checkbox&#x60; or &#x60;highlight&#x60;. | 
**Label** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** | The seat this field belongs to: one of &#x60;signers[].key&#x60;. | [optional] 
**Sensitive** | Pointer to **bool** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** | Tick boxes that answer one question share a group: pick one. | [optional] 
**Uncertain** | **bool** |  | 

## Methods

### NewDetectedField

`func NewDetectedField(id string, type_ string, page int32, x float32, y float32, w float32, h float32, confidence float32, source string, uncertain bool, ) *DetectedField`

NewDetectedField instantiates a new DetectedField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectedFieldWithDefaults

`func NewDetectedFieldWithDefaults() *DetectedField`

NewDetectedFieldWithDefaults instantiates a new DetectedField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetectedField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetectedField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetectedField) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *DetectedField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DetectedField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DetectedField) SetType(v string)`

SetType sets Type field to given value.


### GetPage

`func (o *DetectedField) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DetectedField) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DetectedField) SetPage(v int32)`

SetPage sets Page field to given value.


### GetX

`func (o *DetectedField) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *DetectedField) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *DetectedField) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *DetectedField) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *DetectedField) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *DetectedField) SetY(v float32)`

SetY sets Y field to given value.


### GetW

`func (o *DetectedField) GetW() float32`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *DetectedField) GetWOk() (*float32, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *DetectedField) SetW(v float32)`

SetW sets W field to given value.


### GetH

`func (o *DetectedField) GetH() float32`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *DetectedField) GetHOk() (*float32, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *DetectedField) SetH(v float32)`

SetH sets H field to given value.


### GetConfidence

`func (o *DetectedField) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *DetectedField) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *DetectedField) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetSource

`func (o *DetectedField) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DetectedField) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DetectedField) SetSource(v string)`

SetSource sets Source field to given value.


### GetLabel

`func (o *DetectedField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DetectedField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DetectedField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DetectedField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRole

`func (o *DetectedField) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DetectedField) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DetectedField) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DetectedField) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSensitive

`func (o *DetectedField) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *DetectedField) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *DetectedField) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *DetectedField) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetFormat

`func (o *DetectedField) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DetectedField) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DetectedField) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DetectedField) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGroup

`func (o *DetectedField) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DetectedField) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DetectedField) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DetectedField) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetUncertain

`func (o *DetectedField) GetUncertain() bool`

GetUncertain returns the Uncertain field if non-nil, zero value otherwise.

### GetUncertainOk

`func (o *DetectedField) GetUncertainOk() (*bool, bool)`

GetUncertainOk returns a tuple with the Uncertain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertain

`func (o *DetectedField) SetUncertain(v bool)`

SetUncertain sets Uncertain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListTeams200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Teams** | [**[]ListTeams200ResponseTeamsInner**](ListTeams200ResponseTeamsInner.md) |  | 

## Methods

### NewListTeams200Response

`func NewListTeams200Response(teams []ListTeams200ResponseTeamsInner, ) *ListTeams200Response`

NewListTeams200Response instantiates a new ListTeams200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTeams200ResponseWithDefaults

`func NewListTeams200ResponseWithDefaults() *ListTeams200Response`

NewListTeams200ResponseWithDefaults instantiates a new ListTeams200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeams

`func (o *ListTeams200Response) GetTeams() []ListTeams200ResponseTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ListTeams200Response) GetTeamsOk() (*[]ListTeams200ResponseTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ListTeams200Response) SetTeams(v []ListTeams200ResponseTeamsInner)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



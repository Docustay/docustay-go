# WebhookDocumentsDocumentCreatedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The event&#39;s id; the same value as the webhook-id header. Time-ordered, and unchanged when a delivery is retried. | 
**Topic** | **string** | What happened. | 
**SubjectType** | **string** |  | 
**SubjectId** | **string** |  | 
**Payload** | **map[string]interface{}** | Facts about the subject. Never a signing link or a code. &#x60;payload.test&#x60; is true for every event of a test-mode document. | 
**EmittedAt** | **time.Time** |  | 

## Methods

### NewWebhookDocumentsDocumentCreatedRequest

`func NewWebhookDocumentsDocumentCreatedRequest(id string, topic string, subjectType string, subjectId string, payload map[string]interface{}, emittedAt time.Time, ) *WebhookDocumentsDocumentCreatedRequest`

NewWebhookDocumentsDocumentCreatedRequest instantiates a new WebhookDocumentsDocumentCreatedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDocumentsDocumentCreatedRequestWithDefaults

`func NewWebhookDocumentsDocumentCreatedRequestWithDefaults() *WebhookDocumentsDocumentCreatedRequest`

NewWebhookDocumentsDocumentCreatedRequestWithDefaults instantiates a new WebhookDocumentsDocumentCreatedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookDocumentsDocumentCreatedRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDocumentsDocumentCreatedRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDocumentsDocumentCreatedRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTopic

`func (o *WebhookDocumentsDocumentCreatedRequest) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookDocumentsDocumentCreatedRequest) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookDocumentsDocumentCreatedRequest) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetSubjectType

`func (o *WebhookDocumentsDocumentCreatedRequest) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *WebhookDocumentsDocumentCreatedRequest) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *WebhookDocumentsDocumentCreatedRequest) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetSubjectId

`func (o *WebhookDocumentsDocumentCreatedRequest) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *WebhookDocumentsDocumentCreatedRequest) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *WebhookDocumentsDocumentCreatedRequest) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetPayload

`func (o *WebhookDocumentsDocumentCreatedRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookDocumentsDocumentCreatedRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookDocumentsDocumentCreatedRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetEmittedAt

`func (o *WebhookDocumentsDocumentCreatedRequest) GetEmittedAt() time.Time`

GetEmittedAt returns the EmittedAt field if non-nil, zero value otherwise.

### GetEmittedAtOk

`func (o *WebhookDocumentsDocumentCreatedRequest) GetEmittedAtOk() (*time.Time, bool)`

GetEmittedAtOk returns a tuple with the EmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmittedAt

`func (o *WebhookDocumentsDocumentCreatedRequest) SetEmittedAt(v time.Time)`

SetEmittedAt sets EmittedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



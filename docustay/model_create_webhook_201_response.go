/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateWebhook201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebhook201Response{}

// CreateWebhook201Response struct for CreateWebhook201Response
type CreateWebhook201Response struct {
	Id string `json:"id"`
	Name string `json:"name"`
	// null for a relay endpoint
	Url NullableString `json:"url"`
	Relay bool `json:"relay"`
	EventTypes []string `json:"eventTypes"`
	Status string `json:"status"`
	DisabledReason NullableString `json:"disabledReason"`
	ConsecutiveFailures int32 `json:"consecutiveFailures"`
	Fingerprint string `json:"fingerprint"`
	CreatedAt string `json:"createdAt"`
	Secret string `json:"secret"`
}

type _CreateWebhook201Response CreateWebhook201Response

// NewCreateWebhook201Response instantiates a new CreateWebhook201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhook201Response(id string, name string, url NullableString, relay bool, eventTypes []string, status string, disabledReason NullableString, consecutiveFailures int32, fingerprint string, createdAt string, secret string) *CreateWebhook201Response {
	this := CreateWebhook201Response{}
	this.Id = id
	this.Name = name
	this.Url = url
	this.Relay = relay
	this.EventTypes = eventTypes
	this.Status = status
	this.DisabledReason = disabledReason
	this.ConsecutiveFailures = consecutiveFailures
	this.Fingerprint = fingerprint
	this.CreatedAt = createdAt
	this.Secret = secret
	return &this
}

// NewCreateWebhook201ResponseWithDefaults instantiates a new CreateWebhook201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhook201ResponseWithDefaults() *CreateWebhook201Response {
	this := CreateWebhook201Response{}
	return &this
}

// GetId returns the Id field value
func (o *CreateWebhook201Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateWebhook201Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateWebhook201Response) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CreateWebhook201Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWebhook201Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWebhook201Response) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateWebhook201Response) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWebhook201Response) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *CreateWebhook201Response) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetRelay returns the Relay field value
func (o *CreateWebhook201Response) GetRelay() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Relay
}

// GetRelayOk returns a tuple with the Relay field value
// and a boolean to check if the value has been set.
func (o *CreateWebhook201Response) GetRelayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relay, true
}

// SetRelay sets field value
func (o *CreateWebhook201Response) SetRelay(v bool) {
	o.Relay = v
}

// GetEventTypes returns the EventTypes field value
func (o *CreateWebhook201Response) GetEventTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value
// and a boolean to check if the value has been set.
func (o *CreateWebhook201Response) GetEventTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// SetEventTypes sets field value
func (o *CreateWebhook201Response) SetEventTypes(v []string) {
	o.EventTypes = v
}

// GetStatus returns the Status field value
func (o *CreateWebhook201Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateWebhook201Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateWebhook201Response) SetStatus(v string) {
	o.Status = v
}

// GetDisabledReason returns the DisabledReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateWebhook201Response) GetDisabledReason() string {
	if o == nil || o.DisabledReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisabledReason.Get()
}

// GetDisabledReasonOk returns a tuple with the DisabledReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWebhook201Response) GetDisabledReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisabledReason.Get(), o.DisabledReason.IsSet()
}

// SetDisabledReason sets field value
func (o *CreateWebhook201Response) SetDisabledReason(v string) {
	o.DisabledReason.Set(&v)
}

// GetConsecutiveFailures returns the ConsecutiveFailures field value
func (o *CreateWebhook201Response) GetConsecutiveFailures() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConsecutiveFailures
}

// GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field value
// and a boolean to check if the value has been set.
func (o *CreateWebhook201Response) GetConsecutiveFailuresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsecutiveFailures, true
}

// SetConsecutiveFailures sets field value
func (o *CreateWebhook201Response) SetConsecutiveFailures(v int32) {
	o.ConsecutiveFailures = v
}

// GetFingerprint returns the Fingerprint field value
func (o *CreateWebhook201Response) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *CreateWebhook201Response) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *CreateWebhook201Response) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreateWebhook201Response) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateWebhook201Response) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreateWebhook201Response) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetSecret returns the Secret field value
func (o *CreateWebhook201Response) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateWebhook201Response) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateWebhook201Response) SetSecret(v string) {
	o.Secret = v
}

func (o CreateWebhook201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebhook201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url.Get()
	toSerialize["relay"] = o.Relay
	toSerialize["eventTypes"] = o.EventTypes
	toSerialize["status"] = o.Status
	toSerialize["disabledReason"] = o.DisabledReason.Get()
	toSerialize["consecutiveFailures"] = o.ConsecutiveFailures
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["secret"] = o.Secret
	return toSerialize, nil
}

func (o *CreateWebhook201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"url",
		"relay",
		"eventTypes",
		"status",
		"disabledReason",
		"consecutiveFailures",
		"fingerprint",
		"createdAt",
		"secret",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateWebhook201Response := _CreateWebhook201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWebhook201Response)

	if err != nil {
		return err
	}

	*o = CreateWebhook201Response(varCreateWebhook201Response)

	return err
}

type NullableCreateWebhook201Response struct {
	value *CreateWebhook201Response
	isSet bool
}

func (v NullableCreateWebhook201Response) Get() *CreateWebhook201Response {
	return v.value
}

func (v *NullableCreateWebhook201Response) Set(val *CreateWebhook201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhook201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhook201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhook201Response(val *CreateWebhook201Response) *NullableCreateWebhook201Response {
	return &NullableCreateWebhook201Response{value: val, isSet: true}
}

func (v NullableCreateWebhook201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhook201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



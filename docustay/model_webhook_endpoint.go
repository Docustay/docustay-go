/*
Docustay API 2026-10-07
https://docustay.app. MIT licence.
*/

// Code generated from the Docustay OpenAPI file by the Docustay SDK build; DO NOT EDIT.

package docustay

import (
	"encoding/json"
	"fmt"
)

// checks if the WebhookEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEndpoint{}

// WebhookEndpoint struct for WebhookEndpoint
type WebhookEndpoint struct {
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
	AdditionalProperties map[string]interface{}
}

type _WebhookEndpoint WebhookEndpoint

// NewWebhookEndpoint instantiates a new WebhookEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEndpoint(id string, name string, url NullableString, relay bool, eventTypes []string, status string, disabledReason NullableString, consecutiveFailures int32, fingerprint string, createdAt string) *WebhookEndpoint {
	this := WebhookEndpoint{}
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
	return &this
}

// NewWebhookEndpointWithDefaults instantiates a new WebhookEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEndpointWithDefaults() *WebhookEndpoint {
	this := WebhookEndpoint{}
	return &this
}

// GetId returns the Id field value
func (o *WebhookEndpoint) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpoint) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookEndpoint) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *WebhookEndpoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebhookEndpoint) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookEndpoint) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookEndpoint) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *WebhookEndpoint) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetRelay returns the Relay field value
func (o *WebhookEndpoint) GetRelay() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Relay
}

// GetRelayOk returns a tuple with the Relay field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpoint) GetRelayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relay, true
}

// SetRelay sets field value
func (o *WebhookEndpoint) SetRelay(v bool) {
	o.Relay = v
}

// GetEventTypes returns the EventTypes field value
func (o *WebhookEndpoint) GetEventTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpoint) GetEventTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// SetEventTypes sets field value
func (o *WebhookEndpoint) SetEventTypes(v []string) {
	o.EventTypes = v
}

// GetStatus returns the Status field value
func (o *WebhookEndpoint) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpoint) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WebhookEndpoint) SetStatus(v string) {
	o.Status = v
}

// GetDisabledReason returns the DisabledReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookEndpoint) GetDisabledReason() string {
	if o == nil || o.DisabledReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisabledReason.Get()
}

// GetDisabledReasonOk returns a tuple with the DisabledReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookEndpoint) GetDisabledReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisabledReason.Get(), o.DisabledReason.IsSet()
}

// SetDisabledReason sets field value
func (o *WebhookEndpoint) SetDisabledReason(v string) {
	o.DisabledReason.Set(&v)
}

// GetConsecutiveFailures returns the ConsecutiveFailures field value
func (o *WebhookEndpoint) GetConsecutiveFailures() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConsecutiveFailures
}

// GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpoint) GetConsecutiveFailuresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsecutiveFailures, true
}

// SetConsecutiveFailures sets field value
func (o *WebhookEndpoint) SetConsecutiveFailures(v int32) {
	o.ConsecutiveFailures = v
}

// GetFingerprint returns the Fingerprint field value
func (o *WebhookEndpoint) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpoint) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *WebhookEndpoint) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WebhookEndpoint) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpoint) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WebhookEndpoint) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o WebhookEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEndpoint) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookEndpoint) UnmarshalJSON(data []byte) (err error) {
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

	varWebhookEndpoint := _WebhookEndpoint{}

	err = json.Unmarshal(data, &varWebhookEndpoint)

	if err != nil {
		return err
	}

	*o = WebhookEndpoint(varWebhookEndpoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "url")
		delete(additionalProperties, "relay")
		delete(additionalProperties, "eventTypes")
		delete(additionalProperties, "status")
		delete(additionalProperties, "disabledReason")
		delete(additionalProperties, "consecutiveFailures")
		delete(additionalProperties, "fingerprint")
		delete(additionalProperties, "createdAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookEndpoint struct {
	value *WebhookEndpoint
	isSet bool
}

func (v NullableWebhookEndpoint) Get() *WebhookEndpoint {
	return v.value
}

func (v *NullableWebhookEndpoint) Set(val *WebhookEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEndpoint(val *WebhookEndpoint) *NullableWebhookEndpoint {
	return &NullableWebhookEndpoint{value: val, isSet: true}
}

func (v NullableWebhookEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the CreateWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebhookRequest{}

// CreateWebhookRequest struct for CreateWebhookRequest
type CreateWebhookRequest struct {
	Url *string `json:"url,omitempty"`
	Relay *bool `json:"relay,omitempty"`
	Name *string `json:"name,omitempty"`
	// From `GET /api/v1/webhooks/topics`.
	EventTypes []string `json:"eventTypes"`
}

type _CreateWebhookRequest CreateWebhookRequest

// NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhookRequest(eventTypes []string) *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	this.EventTypes = eventTypes
	return &this
}

// NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateWebhookRequest) SetUrl(v string) {
	o.Url = &v
}

// GetRelay returns the Relay field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetRelay() bool {
	if o == nil || IsNil(o.Relay) {
		var ret bool
		return ret
	}
	return *o.Relay
}

// GetRelayOk returns a tuple with the Relay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetRelayOk() (*bool, bool) {
	if o == nil || IsNil(o.Relay) {
		return nil, false
	}
	return o.Relay, true
}

// HasRelay returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasRelay() bool {
	if o != nil && !IsNil(o.Relay) {
		return true
	}

	return false
}

// SetRelay gets a reference to the given bool and assigns it to the Relay field.
func (o *CreateWebhookRequest) SetRelay(v bool) {
	o.Relay = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateWebhookRequest) SetName(v string) {
	o.Name = &v
}

// GetEventTypes returns the EventTypes field value
func (o *CreateWebhookRequest) GetEventTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetEventTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// SetEventTypes sets field value
func (o *CreateWebhookRequest) SetEventTypes(v []string) {
	o.EventTypes = v
}

func (o CreateWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Relay) {
		toSerialize["relay"] = o.Relay
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["eventTypes"] = o.EventTypes
	return toSerialize, nil
}

func (o *CreateWebhookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventTypes",
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

	varCreateWebhookRequest := _CreateWebhookRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWebhookRequest)

	if err != nil {
		return err
	}

	*o = CreateWebhookRequest(varCreateWebhookRequest)

	return err
}

type NullableCreateWebhookRequest struct {
	value *CreateWebhookRequest
	isSet bool
}

func (v NullableCreateWebhookRequest) Get() *CreateWebhookRequest {
	return v.value
}

func (v *NullableCreateWebhookRequest) Set(val *CreateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhookRequest(val *CreateWebhookRequest) *NullableCreateWebhookRequest {
	return &NullableCreateWebhookRequest{value: val, isSet: true}
}

func (v NullableCreateWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



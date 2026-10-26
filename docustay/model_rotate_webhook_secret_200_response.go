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

// checks if the RotateWebhookSecret200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RotateWebhookSecret200Response{}

// RotateWebhookSecret200Response struct for RotateWebhookSecret200Response
type RotateWebhookSecret200Response struct {
	Secret string `json:"secret"`
	Fingerprint string `json:"fingerprint"`
	PreviousValidUntil string `json:"previousValidUntil"`
	AdditionalProperties map[string]interface{}
}

type _RotateWebhookSecret200Response RotateWebhookSecret200Response

// NewRotateWebhookSecret200Response instantiates a new RotateWebhookSecret200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotateWebhookSecret200Response(secret string, fingerprint string, previousValidUntil string) *RotateWebhookSecret200Response {
	this := RotateWebhookSecret200Response{}
	this.Secret = secret
	this.Fingerprint = fingerprint
	this.PreviousValidUntil = previousValidUntil
	return &this
}

// NewRotateWebhookSecret200ResponseWithDefaults instantiates a new RotateWebhookSecret200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotateWebhookSecret200ResponseWithDefaults() *RotateWebhookSecret200Response {
	this := RotateWebhookSecret200Response{}
	return &this
}

// GetSecret returns the Secret field value
func (o *RotateWebhookSecret200Response) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *RotateWebhookSecret200Response) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *RotateWebhookSecret200Response) SetSecret(v string) {
	o.Secret = v
}

// GetFingerprint returns the Fingerprint field value
func (o *RotateWebhookSecret200Response) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *RotateWebhookSecret200Response) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *RotateWebhookSecret200Response) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetPreviousValidUntil returns the PreviousValidUntil field value
func (o *RotateWebhookSecret200Response) GetPreviousValidUntil() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousValidUntil
}

// GetPreviousValidUntilOk returns a tuple with the PreviousValidUntil field value
// and a boolean to check if the value has been set.
func (o *RotateWebhookSecret200Response) GetPreviousValidUntilOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousValidUntil, true
}

// SetPreviousValidUntil sets field value
func (o *RotateWebhookSecret200Response) SetPreviousValidUntil(v string) {
	o.PreviousValidUntil = v
}

func (o RotateWebhookSecret200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RotateWebhookSecret200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secret"] = o.Secret
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["previousValidUntil"] = o.PreviousValidUntil

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RotateWebhookSecret200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"secret",
		"fingerprint",
		"previousValidUntil",
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

	varRotateWebhookSecret200Response := _RotateWebhookSecret200Response{}

	err = json.Unmarshal(data, &varRotateWebhookSecret200Response)

	if err != nil {
		return err
	}

	*o = RotateWebhookSecret200Response(varRotateWebhookSecret200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "secret")
		delete(additionalProperties, "fingerprint")
		delete(additionalProperties, "previousValidUntil")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRotateWebhookSecret200Response struct {
	value *RotateWebhookSecret200Response
	isSet bool
}

func (v NullableRotateWebhookSecret200Response) Get() *RotateWebhookSecret200Response {
	return v.value
}

func (v *NullableRotateWebhookSecret200Response) Set(val *RotateWebhookSecret200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRotateWebhookSecret200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRotateWebhookSecret200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotateWebhookSecret200Response(val *RotateWebhookSecret200Response) *NullableRotateWebhookSecret200Response {
	return &NullableRotateWebhookSecret200Response{value: val, isSet: true}
}

func (v NullableRotateWebhookSecret200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotateWebhookSecret200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



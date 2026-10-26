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

// checks if the DocumentMade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentMade{}

// DocumentMade struct for DocumentMade
type DocumentMade struct {
	Approval *SendDraft202ResponseApproval `json:"approval,omitempty"`
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId NullableString `json:"externalId,omitempty"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	Id string `json:"id"`
	State string `json:"state"`
	Roles []TemplateSeatsInner `json:"roles,omitempty"`
	Emailed *int32 `json:"emailed,omitempty"`
	NotEmailed []interface{} `json:"notEmailed,omitempty"`
	Warnings []interface{} `json:"warnings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DocumentMade DocumentMade

// NewDocumentMade instantiates a new DocumentMade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentMade(id string, state string) *DocumentMade {
	this := DocumentMade{}
	this.Id = id
	this.State = state
	return &this
}

// NewDocumentMadeWithDefaults instantiates a new DocumentMade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentMadeWithDefaults() *DocumentMade {
	this := DocumentMade{}
	return &this
}

// GetApproval returns the Approval field value if set, zero value otherwise.
func (o *DocumentMade) GetApproval() SendDraft202ResponseApproval {
	if o == nil || IsNil(o.Approval) {
		var ret SendDraft202ResponseApproval
		return ret
	}
	return *o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMade) GetApprovalOk() (*SendDraft202ResponseApproval, bool) {
	if o == nil || IsNil(o.Approval) {
		return nil, false
	}
	return o.Approval, true
}

// HasApproval returns a boolean if a field has been set.
func (o *DocumentMade) HasApproval() bool {
	if o != nil && !IsNil(o.Approval) {
		return true
	}

	return false
}

// SetApproval gets a reference to the given SendDraft202ResponseApproval and assigns it to the Approval field.
func (o *DocumentMade) SetApproval(v SendDraft202ResponseApproval) {
	o.Approval = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentMade) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentMade) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *DocumentMade) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *DocumentMade) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *DocumentMade) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *DocumentMade) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DocumentMade) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMade) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DocumentMade) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DocumentMade) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetId returns the Id field value
func (o *DocumentMade) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DocumentMade) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DocumentMade) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *DocumentMade) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DocumentMade) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DocumentMade) SetState(v string) {
	o.State = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *DocumentMade) GetRoles() []TemplateSeatsInner {
	if o == nil || IsNil(o.Roles) {
		var ret []TemplateSeatsInner
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMade) GetRolesOk() ([]TemplateSeatsInner, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *DocumentMade) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []TemplateSeatsInner and assigns it to the Roles field.
func (o *DocumentMade) SetRoles(v []TemplateSeatsInner) {
	o.Roles = v
}

// GetEmailed returns the Emailed field value if set, zero value otherwise.
func (o *DocumentMade) GetEmailed() int32 {
	if o == nil || IsNil(o.Emailed) {
		var ret int32
		return ret
	}
	return *o.Emailed
}

// GetEmailedOk returns a tuple with the Emailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMade) GetEmailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Emailed) {
		return nil, false
	}
	return o.Emailed, true
}

// HasEmailed returns a boolean if a field has been set.
func (o *DocumentMade) HasEmailed() bool {
	if o != nil && !IsNil(o.Emailed) {
		return true
	}

	return false
}

// SetEmailed gets a reference to the given int32 and assigns it to the Emailed field.
func (o *DocumentMade) SetEmailed(v int32) {
	o.Emailed = &v
}

// GetNotEmailed returns the NotEmailed field value if set, zero value otherwise.
func (o *DocumentMade) GetNotEmailed() []interface{} {
	if o == nil || IsNil(o.NotEmailed) {
		var ret []interface{}
		return ret
	}
	return o.NotEmailed
}

// GetNotEmailedOk returns a tuple with the NotEmailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMade) GetNotEmailedOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.NotEmailed) {
		return nil, false
	}
	return o.NotEmailed, true
}

// HasNotEmailed returns a boolean if a field has been set.
func (o *DocumentMade) HasNotEmailed() bool {
	if o != nil && !IsNil(o.NotEmailed) {
		return true
	}

	return false
}

// SetNotEmailed gets a reference to the given []interface{} and assigns it to the NotEmailed field.
func (o *DocumentMade) SetNotEmailed(v []interface{}) {
	o.NotEmailed = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DocumentMade) GetWarnings() []interface{} {
	if o == nil || IsNil(o.Warnings) {
		var ret []interface{}
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMade) GetWarningsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DocumentMade) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []interface{} and assigns it to the Warnings field.
func (o *DocumentMade) SetWarnings(v []interface{}) {
	o.Warnings = v
}

func (o DocumentMade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentMade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Approval) {
		toSerialize["approval"] = o.Approval
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Emailed) {
		toSerialize["emailed"] = o.Emailed
	}
	if !IsNil(o.NotEmailed) {
		toSerialize["notEmailed"] = o.NotEmailed
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DocumentMade) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
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

	varDocumentMade := _DocumentMade{}

	err = json.Unmarshal(data, &varDocumentMade)

	if err != nil {
		return err
	}

	*o = DocumentMade(varDocumentMade)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "approval")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "emailed")
		delete(additionalProperties, "notEmailed")
		delete(additionalProperties, "warnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentMade struct {
	value *DocumentMade
	isSet bool
}

func (v NullableDocumentMade) Get() *DocumentMade {
	return v.value
}

func (v *NullableDocumentMade) Set(val *DocumentMade) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentMade) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentMade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentMade(val *DocumentMade) *NullableDocumentMade {
	return &NullableDocumentMade{value: val, isSet: true}
}

func (v NullableDocumentMade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentMade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



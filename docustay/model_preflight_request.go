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

// checks if the PreflightRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreflightRequest{}

// PreflightRequest struct for PreflightRequest
type PreflightRequest struct {
	// What you are about to do: `send_document`, `bulk_send` (with `documents`), `invite_person` or `remove_person` (with `people`), `ai_draft`, `ai_edit`, `ai_read_file`, `create_template`, `upload` (with `bytes`), `use_feature` (with `item`) or `connect_payments`.
	Action string `json:"action"`
	Documents *int32 `json:"documents,omitempty"`
	People *int32 `json:"people,omitempty"`
	// For `send_document`: the email addresses you are about to send to, so a Free workspace is told how many of them are new signers this month (anyone already sent to this month, and your own team, do not count).
	Recipients []string `json:"recipients,omitempty"`
	Bytes *int32 `json:"bytes,omitempty"`
	// For `use_feature`: the item to ask about, one of `unlimited_documents`, `hide_footer`, `auto_reminders`, `bulk_send`, `team_roles`, `signing_groups`, `scribe_actions`, `ai_bulk`, `custom_plays` (your own plays and playbooks), `plays_api`. Any other name is a 400 `unknown_item`.
	Item *string `json:"item,omitempty"`
}

type _PreflightRequest PreflightRequest

// NewPreflightRequest instantiates a new PreflightRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreflightRequest(action string) *PreflightRequest {
	this := PreflightRequest{}
	this.Action = action
	return &this
}

// NewPreflightRequestWithDefaults instantiates a new PreflightRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreflightRequestWithDefaults() *PreflightRequest {
	this := PreflightRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *PreflightRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PreflightRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PreflightRequest) SetAction(v string) {
	o.Action = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *PreflightRequest) GetDocuments() int32 {
	if o == nil || IsNil(o.Documents) {
		var ret int32
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightRequest) GetDocumentsOk() (*int32, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *PreflightRequest) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given int32 and assigns it to the Documents field.
func (o *PreflightRequest) SetDocuments(v int32) {
	o.Documents = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *PreflightRequest) GetPeople() int32 {
	if o == nil || IsNil(o.People) {
		var ret int32
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightRequest) GetPeopleOk() (*int32, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *PreflightRequest) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given int32 and assigns it to the People field.
func (o *PreflightRequest) SetPeople(v int32) {
	o.People = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *PreflightRequest) GetRecipients() []string {
	if o == nil || IsNil(o.Recipients) {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightRequest) GetRecipientsOk() ([]string, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *PreflightRequest) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *PreflightRequest) SetRecipients(v []string) {
	o.Recipients = v
}

// GetBytes returns the Bytes field value if set, zero value otherwise.
func (o *PreflightRequest) GetBytes() int32 {
	if o == nil || IsNil(o.Bytes) {
		var ret int32
		return ret
	}
	return *o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightRequest) GetBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.Bytes) {
		return nil, false
	}
	return o.Bytes, true
}

// HasBytes returns a boolean if a field has been set.
func (o *PreflightRequest) HasBytes() bool {
	if o != nil && !IsNil(o.Bytes) {
		return true
	}

	return false
}

// SetBytes gets a reference to the given int32 and assigns it to the Bytes field.
func (o *PreflightRequest) SetBytes(v int32) {
	o.Bytes = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *PreflightRequest) GetItem() string {
	if o == nil || IsNil(o.Item) {
		var ret string
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreflightRequest) GetItemOk() (*string, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *PreflightRequest) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given string and assigns it to the Item field.
func (o *PreflightRequest) SetItem(v string) {
	o.Item = &v
}

func (o PreflightRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreflightRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.Bytes) {
		toSerialize["bytes"] = o.Bytes
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

func (o *PreflightRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varPreflightRequest := _PreflightRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPreflightRequest)

	if err != nil {
		return err
	}

	*o = PreflightRequest(varPreflightRequest)

	return err
}

type NullablePreflightRequest struct {
	value *PreflightRequest
	isSet bool
}

func (v NullablePreflightRequest) Get() *PreflightRequest {
	return v.value
}

func (v *NullablePreflightRequest) Set(val *PreflightRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePreflightRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePreflightRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreflightRequest(val *PreflightRequest) *NullablePreflightRequest {
	return &NullablePreflightRequest{value: val, isSet: true}
}

func (v NullablePreflightRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreflightRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



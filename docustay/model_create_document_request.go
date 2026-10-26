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

// checks if the CreateDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDocumentRequest{}

// CreateDocumentRequest struct for CreateDocumentRequest
type CreateDocumentRequest struct {
	Email *EmailChoice `json:"email,omitempty"`
	// Your own key-value text (up to 20 keys; keys 1–40 letters, digits, dot, dash or underscore; values up to 500 characters). Returned in every response and every webhook for this document; never shown to a signer.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Your own id for this thing, so you can match it to your records. Filter lists with `externalId=`.
	ExternalId *string `json:"externalId,omitempty"`
	Title string `json:"title"`
	// A written document: TipTap JSON `{type:\"doc\",content:[…]}`. Field tags (`fieldTag` with name, type, role, required, autofill, editable, placeholder, tooltip, locked, options, default…), merge chips (`variableChip`; `custom.<key>` chips are the template's questions), and headings with `audience: \"sender\"` (Sender only: in nothing a signer sees, signs or receives).
	Source map[string]*interface{} `json:"source"`
	People []Person `json:"people"`
	// Merge text: `org.name`, `sender.name`, `custom.<key>` …
	Facts map[string]string `json:"facts,omitempty"`
	VariableValues map[string]string `json:"variableValues,omitempty"`
	Paper *string `json:"paper,omitempty"`
	Send *SendChoices `json:"send,omitempty"`
	Description *string `json:"description,omitempty"`
	// Also send it. The same checks as the screen's Send apply (a payment field needs Stripe connected).
	SendNow *bool `json:"sendNow,omitempty"`
}

type _CreateDocumentRequest CreateDocumentRequest

// NewCreateDocumentRequest instantiates a new CreateDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDocumentRequest(title string, source map[string]*interface{}, people []Person) *CreateDocumentRequest {
	this := CreateDocumentRequest{}
	this.Title = title
	this.Source = source
	this.People = people
	return &this
}

// NewCreateDocumentRequestWithDefaults instantiates a new CreateDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDocumentRequestWithDefaults() *CreateDocumentRequest {
	this := CreateDocumentRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateDocumentRequest) GetEmail() EmailChoice {
	if o == nil || IsNil(o.Email) {
		var ret EmailChoice
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetEmailOk() (*EmailChoice, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateDocumentRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailChoice and assigns it to the Email field.
func (o *CreateDocumentRequest) SetEmail(v EmailChoice) {
	o.Email = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateDocumentRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateDocumentRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateDocumentRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *CreateDocumentRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *CreateDocumentRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *CreateDocumentRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetTitle returns the Title field value
func (o *CreateDocumentRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateDocumentRequest) SetTitle(v string) {
	o.Title = v
}

// GetSource returns the Source field value
func (o *CreateDocumentRequest) GetSource() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetSourceOk() (map[string]*interface{}, bool) {
	if o == nil {
		return map[string]*interface{}{}, false
	}
	return o.Source, true
}

// SetSource sets field value
func (o *CreateDocumentRequest) SetSource(v map[string]*interface{}) {
	o.Source = v
}

// GetPeople returns the People field value
func (o *CreateDocumentRequest) GetPeople() []Person {
	if o == nil {
		var ret []Person
		return ret
	}

	return o.People
}

// GetPeopleOk returns a tuple with the People field value
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetPeopleOk() ([]Person, bool) {
	if o == nil {
		return nil, false
	}
	return o.People, true
}

// SetPeople sets field value
func (o *CreateDocumentRequest) SetPeople(v []Person) {
	o.People = v
}

// GetFacts returns the Facts field value if set, zero value otherwise.
func (o *CreateDocumentRequest) GetFacts() map[string]string {
	if o == nil || IsNil(o.Facts) {
		var ret map[string]string
		return ret
	}
	return o.Facts
}

// GetFactsOk returns a tuple with the Facts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetFactsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Facts) {
		return map[string]string{}, false
	}
	return o.Facts, true
}

// HasFacts returns a boolean if a field has been set.
func (o *CreateDocumentRequest) HasFacts() bool {
	if o != nil && !IsNil(o.Facts) {
		return true
	}

	return false
}

// SetFacts gets a reference to the given map[string]string and assigns it to the Facts field.
func (o *CreateDocumentRequest) SetFacts(v map[string]string) {
	o.Facts = v
}

// GetVariableValues returns the VariableValues field value if set, zero value otherwise.
func (o *CreateDocumentRequest) GetVariableValues() map[string]string {
	if o == nil || IsNil(o.VariableValues) {
		var ret map[string]string
		return ret
	}
	return o.VariableValues
}

// GetVariableValuesOk returns a tuple with the VariableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetVariableValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.VariableValues) {
		return map[string]string{}, false
	}
	return o.VariableValues, true
}

// HasVariableValues returns a boolean if a field has been set.
func (o *CreateDocumentRequest) HasVariableValues() bool {
	if o != nil && !IsNil(o.VariableValues) {
		return true
	}

	return false
}

// SetVariableValues gets a reference to the given map[string]string and assigns it to the VariableValues field.
func (o *CreateDocumentRequest) SetVariableValues(v map[string]string) {
	o.VariableValues = v
}

// GetPaper returns the Paper field value if set, zero value otherwise.
func (o *CreateDocumentRequest) GetPaper() string {
	if o == nil || IsNil(o.Paper) {
		var ret string
		return ret
	}
	return *o.Paper
}

// GetPaperOk returns a tuple with the Paper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetPaperOk() (*string, bool) {
	if o == nil || IsNil(o.Paper) {
		return nil, false
	}
	return o.Paper, true
}

// HasPaper returns a boolean if a field has been set.
func (o *CreateDocumentRequest) HasPaper() bool {
	if o != nil && !IsNil(o.Paper) {
		return true
	}

	return false
}

// SetPaper gets a reference to the given string and assigns it to the Paper field.
func (o *CreateDocumentRequest) SetPaper(v string) {
	o.Paper = &v
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *CreateDocumentRequest) GetSend() SendChoices {
	if o == nil || IsNil(o.Send) {
		var ret SendChoices
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetSendOk() (*SendChoices, bool) {
	if o == nil || IsNil(o.Send) {
		return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *CreateDocumentRequest) HasSend() bool {
	if o != nil && !IsNil(o.Send) {
		return true
	}

	return false
}

// SetSend gets a reference to the given SendChoices and assigns it to the Send field.
func (o *CreateDocumentRequest) SetSend(v SendChoices) {
	o.Send = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDocumentRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDocumentRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDocumentRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSendNow returns the SendNow field value if set, zero value otherwise.
func (o *CreateDocumentRequest) GetSendNow() bool {
	if o == nil || IsNil(o.SendNow) {
		var ret bool
		return ret
	}
	return *o.SendNow
}

// GetSendNowOk returns a tuple with the SendNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequest) GetSendNowOk() (*bool, bool) {
	if o == nil || IsNil(o.SendNow) {
		return nil, false
	}
	return o.SendNow, true
}

// HasSendNow returns a boolean if a field has been set.
func (o *CreateDocumentRequest) HasSendNow() bool {
	if o != nil && !IsNil(o.SendNow) {
		return true
	}

	return false
}

// SetSendNow gets a reference to the given bool and assigns it to the SendNow field.
func (o *CreateDocumentRequest) SetSendNow(v bool) {
	o.SendNow = &v
}

func (o CreateDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	toSerialize["title"] = o.Title
	toSerialize["source"] = o.Source
	toSerialize["people"] = o.People
	if !IsNil(o.Facts) {
		toSerialize["facts"] = o.Facts
	}
	if !IsNil(o.VariableValues) {
		toSerialize["variableValues"] = o.VariableValues
	}
	if !IsNil(o.Paper) {
		toSerialize["paper"] = o.Paper
	}
	if !IsNil(o.Send) {
		toSerialize["send"] = o.Send
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SendNow) {
		toSerialize["sendNow"] = o.SendNow
	}
	return toSerialize, nil
}

func (o *CreateDocumentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"source",
		"people",
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

	varCreateDocumentRequest := _CreateDocumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDocumentRequest)

	if err != nil {
		return err
	}

	*o = CreateDocumentRequest(varCreateDocumentRequest)

	return err
}

type NullableCreateDocumentRequest struct {
	value *CreateDocumentRequest
	isSet bool
}

func (v NullableCreateDocumentRequest) Get() *CreateDocumentRequest {
	return v.value
}

func (v *NullableCreateDocumentRequest) Set(val *CreateDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDocumentRequest(val *CreateDocumentRequest) *NullableCreateDocumentRequest {
	return &NullableCreateDocumentRequest{value: val, isSet: true}
}

func (v NullableCreateDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



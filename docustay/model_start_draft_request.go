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

// checks if the StartDraftRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartDraftRequest{}

// StartDraftRequest struct for StartDraftRequest
type StartDraftRequest struct {
	// A library template's slug (see `GET /api/v1/templates/library`).
	Template string `json:"template" validate:"regexp=^[a-z][a-z0-9-]{1\\,60}$"`
	Title *string `json:"title,omitempty"`
	People []StartDraftRequestPeopleInner `json:"people"`
	Answers map[string]StartDraftRequestAnswersValue `json:"answers,omitempty"`
	// Send-time amounts the template's words need, by key (a `needs_input` answer lists them under `values`), e.g. `{ \"deposit\": \"$500\" }`.
	Values map[string]string `json:"values,omitempty"`
	UseDefaults *bool `json:"useDefaults,omitempty"`
}

type _StartDraftRequest StartDraftRequest

// NewStartDraftRequest instantiates a new StartDraftRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartDraftRequest(template string, people []StartDraftRequestPeopleInner) *StartDraftRequest {
	this := StartDraftRequest{}
	this.Template = template
	this.People = people
	return &this
}

// NewStartDraftRequestWithDefaults instantiates a new StartDraftRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartDraftRequestWithDefaults() *StartDraftRequest {
	this := StartDraftRequest{}
	return &this
}

// GetTemplate returns the Template field value
func (o *StartDraftRequest) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *StartDraftRequest) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *StartDraftRequest) SetTemplate(v string) {
	o.Template = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *StartDraftRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartDraftRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *StartDraftRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *StartDraftRequest) SetTitle(v string) {
	o.Title = &v
}

// GetPeople returns the People field value
func (o *StartDraftRequest) GetPeople() []StartDraftRequestPeopleInner {
	if o == nil {
		var ret []StartDraftRequestPeopleInner
		return ret
	}

	return o.People
}

// GetPeopleOk returns a tuple with the People field value
// and a boolean to check if the value has been set.
func (o *StartDraftRequest) GetPeopleOk() ([]StartDraftRequestPeopleInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.People, true
}

// SetPeople sets field value
func (o *StartDraftRequest) SetPeople(v []StartDraftRequestPeopleInner) {
	o.People = v
}

// GetAnswers returns the Answers field value if set, zero value otherwise.
func (o *StartDraftRequest) GetAnswers() map[string]StartDraftRequestAnswersValue {
	if o == nil || IsNil(o.Answers) {
		var ret map[string]StartDraftRequestAnswersValue
		return ret
	}
	return o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartDraftRequest) GetAnswersOk() (map[string]StartDraftRequestAnswersValue, bool) {
	if o == nil || IsNil(o.Answers) {
		return map[string]StartDraftRequestAnswersValue{}, false
	}
	return o.Answers, true
}

// HasAnswers returns a boolean if a field has been set.
func (o *StartDraftRequest) HasAnswers() bool {
	if o != nil && !IsNil(o.Answers) {
		return true
	}

	return false
}

// SetAnswers gets a reference to the given map[string]StartDraftRequestAnswersValue and assigns it to the Answers field.
func (o *StartDraftRequest) SetAnswers(v map[string]StartDraftRequestAnswersValue) {
	o.Answers = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *StartDraftRequest) GetValues() map[string]string {
	if o == nil || IsNil(o.Values) {
		var ret map[string]string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartDraftRequest) GetValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Values) {
		return map[string]string{}, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *StartDraftRequest) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]string and assigns it to the Values field.
func (o *StartDraftRequest) SetValues(v map[string]string) {
	o.Values = v
}

// GetUseDefaults returns the UseDefaults field value if set, zero value otherwise.
func (o *StartDraftRequest) GetUseDefaults() bool {
	if o == nil || IsNil(o.UseDefaults) {
		var ret bool
		return ret
	}
	return *o.UseDefaults
}

// GetUseDefaultsOk returns a tuple with the UseDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartDraftRequest) GetUseDefaultsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDefaults) {
		return nil, false
	}
	return o.UseDefaults, true
}

// HasUseDefaults returns a boolean if a field has been set.
func (o *StartDraftRequest) HasUseDefaults() bool {
	if o != nil && !IsNil(o.UseDefaults) {
		return true
	}

	return false
}

// SetUseDefaults gets a reference to the given bool and assigns it to the UseDefaults field.
func (o *StartDraftRequest) SetUseDefaults(v bool) {
	o.UseDefaults = &v
}

func (o StartDraftRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartDraftRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["template"] = o.Template
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	toSerialize["people"] = o.People
	if !IsNil(o.Answers) {
		toSerialize["answers"] = o.Answers
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.UseDefaults) {
		toSerialize["useDefaults"] = o.UseDefaults
	}
	return toSerialize, nil
}

func (o *StartDraftRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"template",
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

	varStartDraftRequest := _StartDraftRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStartDraftRequest)

	if err != nil {
		return err
	}

	*o = StartDraftRequest(varStartDraftRequest)

	return err
}

type NullableStartDraftRequest struct {
	value *StartDraftRequest
	isSet bool
}

func (v NullableStartDraftRequest) Get() *StartDraftRequest {
	return v.value
}

func (v *NullableStartDraftRequest) Set(val *StartDraftRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStartDraftRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStartDraftRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartDraftRequest(val *StartDraftRequest) *NullableStartDraftRequest {
	return &NullableStartDraftRequest{value: val, isSet: true}
}

func (v NullableStartDraftRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartDraftRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



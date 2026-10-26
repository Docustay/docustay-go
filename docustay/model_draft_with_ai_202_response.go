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

// checks if the DraftWithAi202Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DraftWithAi202Response{}

// DraftWithAi202Response struct for DraftWithAi202Response
type DraftWithAi202Response struct {
	JobId string `json:"jobId"`
	State string `json:"state"`
	AdditionalProperties map[string]interface{}
}

type _DraftWithAi202Response DraftWithAi202Response

// NewDraftWithAi202Response instantiates a new DraftWithAi202Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDraftWithAi202Response(jobId string, state string) *DraftWithAi202Response {
	this := DraftWithAi202Response{}
	this.JobId = jobId
	this.State = state
	return &this
}

// NewDraftWithAi202ResponseWithDefaults instantiates a new DraftWithAi202Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDraftWithAi202ResponseWithDefaults() *DraftWithAi202Response {
	this := DraftWithAi202Response{}
	return &this
}

// GetJobId returns the JobId field value
func (o *DraftWithAi202Response) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *DraftWithAi202Response) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *DraftWithAi202Response) SetJobId(v string) {
	o.JobId = v
}

// GetState returns the State field value
func (o *DraftWithAi202Response) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DraftWithAi202Response) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DraftWithAi202Response) SetState(v string) {
	o.State = v
}

func (o DraftWithAi202Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DraftWithAi202Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DraftWithAi202Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobId",
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

	varDraftWithAi202Response := _DraftWithAi202Response{}

	err = json.Unmarshal(data, &varDraftWithAi202Response)

	if err != nil {
		return err
	}

	*o = DraftWithAi202Response(varDraftWithAi202Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDraftWithAi202Response struct {
	value *DraftWithAi202Response
	isSet bool
}

func (v NullableDraftWithAi202Response) Get() *DraftWithAi202Response {
	return v.value
}

func (v *NullableDraftWithAi202Response) Set(val *DraftWithAi202Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDraftWithAi202Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDraftWithAi202Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDraftWithAi202Response(val *DraftWithAi202Response) *NullableDraftWithAi202Response {
	return &NullableDraftWithAi202Response{value: val, isSet: true}
}

func (v NullableDraftWithAi202Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDraftWithAi202Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



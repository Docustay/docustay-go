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

// checks if the TestSign200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSign200Response{}

// TestSign200Response struct for TestSign200Response
type TestSign200Response struct {
	Id string `json:"id"`
	Seat string `json:"seat"`
	Signed bool `json:"signed"`
	Declined *bool `json:"declined,omitempty"`
	State string `json:"state"`
	Progress TestSign200ResponseProgress `json:"progress"`
	Finished bool `json:"finished"`
	AdditionalProperties map[string]interface{}
}

type _TestSign200Response TestSign200Response

// NewTestSign200Response instantiates a new TestSign200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSign200Response(id string, seat string, signed bool, state string, progress TestSign200ResponseProgress, finished bool) *TestSign200Response {
	this := TestSign200Response{}
	this.Id = id
	this.Seat = seat
	this.Signed = signed
	this.State = state
	this.Progress = progress
	this.Finished = finished
	return &this
}

// NewTestSign200ResponseWithDefaults instantiates a new TestSign200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSign200ResponseWithDefaults() *TestSign200Response {
	this := TestSign200Response{}
	return &this
}

// GetId returns the Id field value
func (o *TestSign200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestSign200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestSign200Response) SetId(v string) {
	o.Id = v
}

// GetSeat returns the Seat field value
func (o *TestSign200Response) GetSeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seat
}

// GetSeatOk returns a tuple with the Seat field value
// and a boolean to check if the value has been set.
func (o *TestSign200Response) GetSeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seat, true
}

// SetSeat sets field value
func (o *TestSign200Response) SetSeat(v string) {
	o.Seat = v
}

// GetSigned returns the Signed field value
func (o *TestSign200Response) GetSigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Signed
}

// GetSignedOk returns a tuple with the Signed field value
// and a boolean to check if the value has been set.
func (o *TestSign200Response) GetSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signed, true
}

// SetSigned sets field value
func (o *TestSign200Response) SetSigned(v bool) {
	o.Signed = v
}

// GetDeclined returns the Declined field value if set, zero value otherwise.
func (o *TestSign200Response) GetDeclined() bool {
	if o == nil || IsNil(o.Declined) {
		var ret bool
		return ret
	}
	return *o.Declined
}

// GetDeclinedOk returns a tuple with the Declined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSign200Response) GetDeclinedOk() (*bool, bool) {
	if o == nil || IsNil(o.Declined) {
		return nil, false
	}
	return o.Declined, true
}

// HasDeclined returns a boolean if a field has been set.
func (o *TestSign200Response) HasDeclined() bool {
	if o != nil && !IsNil(o.Declined) {
		return true
	}

	return false
}

// SetDeclined gets a reference to the given bool and assigns it to the Declined field.
func (o *TestSign200Response) SetDeclined(v bool) {
	o.Declined = &v
}

// GetState returns the State field value
func (o *TestSign200Response) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TestSign200Response) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TestSign200Response) SetState(v string) {
	o.State = v
}

// GetProgress returns the Progress field value
func (o *TestSign200Response) GetProgress() TestSign200ResponseProgress {
	if o == nil {
		var ret TestSign200ResponseProgress
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *TestSign200Response) GetProgressOk() (*TestSign200ResponseProgress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *TestSign200Response) SetProgress(v TestSign200ResponseProgress) {
	o.Progress = v
}

// GetFinished returns the Finished field value
func (o *TestSign200Response) GetFinished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value
// and a boolean to check if the value has been set.
func (o *TestSign200Response) GetFinishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Finished, true
}

// SetFinished sets field value
func (o *TestSign200Response) SetFinished(v bool) {
	o.Finished = v
}

func (o TestSign200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSign200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["seat"] = o.Seat
	toSerialize["signed"] = o.Signed
	if !IsNil(o.Declined) {
		toSerialize["declined"] = o.Declined
	}
	toSerialize["state"] = o.State
	toSerialize["progress"] = o.Progress
	toSerialize["finished"] = o.Finished

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestSign200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"seat",
		"signed",
		"state",
		"progress",
		"finished",
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

	varTestSign200Response := _TestSign200Response{}

	err = json.Unmarshal(data, &varTestSign200Response)

	if err != nil {
		return err
	}

	*o = TestSign200Response(varTestSign200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "seat")
		delete(additionalProperties, "signed")
		delete(additionalProperties, "declined")
		delete(additionalProperties, "state")
		delete(additionalProperties, "progress")
		delete(additionalProperties, "finished")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestSign200Response struct {
	value *TestSign200Response
	isSet bool
}

func (v NullableTestSign200Response) Get() *TestSign200Response {
	return v.value
}

func (v *NullableTestSign200Response) Set(val *TestSign200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSign200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSign200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSign200Response(val *TestSign200Response) *NullableTestSign200Response {
	return &NullableTestSign200Response{value: val, isSet: true}
}

func (v NullableTestSign200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSign200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



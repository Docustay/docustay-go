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


// StartDraftRequestAnswersValue struct for StartDraftRequestAnswersValue
type StartDraftRequestAnswersValue struct {
	ArrayOfString *[]string
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *StartDraftRequestAnswersValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ArrayOfString
	err = json.Unmarshal(data, &dst.ArrayOfString);
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			return nil // data stored in dst.ArrayOfString, return on the first match
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(StartDraftRequestAnswersValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StartDraftRequestAnswersValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableStartDraftRequestAnswersValue struct {
	value *StartDraftRequestAnswersValue
	isSet bool
}

func (v NullableStartDraftRequestAnswersValue) Get() *StartDraftRequestAnswersValue {
	return v.value
}

func (v *NullableStartDraftRequestAnswersValue) Set(val *StartDraftRequestAnswersValue) {
	v.value = val
	v.isSet = true
}

func (v NullableStartDraftRequestAnswersValue) IsSet() bool {
	return v.isSet
}

func (v *NullableStartDraftRequestAnswersValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartDraftRequestAnswersValue(val *StartDraftRequestAnswersValue) *NullableStartDraftRequestAnswersValue {
	return &NullableStartDraftRequestAnswersValue{value: val, isSet: true}
}

func (v NullableStartDraftRequestAnswersValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartDraftRequestAnswersValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



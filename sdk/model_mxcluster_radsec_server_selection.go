/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// MxclusterRadsecServerSelection ordered (default) / unordered. When ordered, Mist Edge will prefer and go back to the first radius server if possible
type MxclusterRadsecServerSelection string

// List of mxcluster_radsec_server_selection
const (
	MXCLUSTERRADSECSERVERSELECTION_EMPTY MxclusterRadsecServerSelection = ""
	MXCLUSTERRADSECSERVERSELECTION_ORDERED MxclusterRadsecServerSelection = "ordered"
	MXCLUSTERRADSECSERVERSELECTION_UNORDERED MxclusterRadsecServerSelection = "unordered"
)

// All allowed values of MxclusterRadsecServerSelection enum
var AllowedMxclusterRadsecServerSelectionEnumValues = []MxclusterRadsecServerSelection{
	"",
	"ordered",
	"unordered",
}

func (v *MxclusterRadsecServerSelection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MxclusterRadsecServerSelection(value)
	for _, existing := range AllowedMxclusterRadsecServerSelectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MxclusterRadsecServerSelection", value)
}

// NewMxclusterRadsecServerSelectionFromValue returns a pointer to a valid MxclusterRadsecServerSelection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMxclusterRadsecServerSelectionFromValue(v string) (*MxclusterRadsecServerSelection, error) {
	ev := MxclusterRadsecServerSelection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MxclusterRadsecServerSelection: valid values are %v", v, AllowedMxclusterRadsecServerSelectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MxclusterRadsecServerSelection) IsValid() bool {
	for _, existing := range AllowedMxclusterRadsecServerSelectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to mxcluster_radsec_server_selection value
func (v MxclusterRadsecServerSelection) Ptr() *MxclusterRadsecServerSelection {
	return &v
}

type NullableMxclusterRadsecServerSelection struct {
	value *MxclusterRadsecServerSelection
	isSet bool
}

func (v NullableMxclusterRadsecServerSelection) Get() *MxclusterRadsecServerSelection {
	return v.value
}

func (v *NullableMxclusterRadsecServerSelection) Set(val *MxclusterRadsecServerSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableMxclusterRadsecServerSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableMxclusterRadsecServerSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxclusterRadsecServerSelection(val *MxclusterRadsecServerSelection) *NullableMxclusterRadsecServerSelection {
	return &NullableMxclusterRadsecServerSelection{value: val, isSet: true}
}

func (v NullableMxclusterRadsecServerSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxclusterRadsecServerSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


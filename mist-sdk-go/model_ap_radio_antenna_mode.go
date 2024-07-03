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

// ApRadioAntennaMode the model 'ApRadioAntennaMode'
type ApRadioAntennaMode string

// List of ap_radio_antenna_mode
const (
	APRADIOANTENNAMODE_EMPTY ApRadioAntennaMode = ""
	APRADIOANTENNAMODE_DEFAULT ApRadioAntennaMode = "default"
	APRADIOANTENNAMODE__1X1 ApRadioAntennaMode = "1x1"
	APRADIOANTENNAMODE__2X2 ApRadioAntennaMode = "2x2"
	APRADIOANTENNAMODE__3X3 ApRadioAntennaMode = "3x3"
	APRADIOANTENNAMODE__4X4 ApRadioAntennaMode = "4x4"
)

// All allowed values of ApRadioAntennaMode enum
var AllowedApRadioAntennaModeEnumValues = []ApRadioAntennaMode{
	"",
	"default",
	"1x1",
	"2x2",
	"3x3",
	"4x4",
}

func (v *ApRadioAntennaMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApRadioAntennaMode(value)
	for _, existing := range AllowedApRadioAntennaModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApRadioAntennaMode", value)
}

// NewApRadioAntennaModeFromValue returns a pointer to a valid ApRadioAntennaMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApRadioAntennaModeFromValue(v string) (*ApRadioAntennaMode, error) {
	ev := ApRadioAntennaMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApRadioAntennaMode: valid values are %v", v, AllowedApRadioAntennaModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApRadioAntennaMode) IsValid() bool {
	for _, existing := range AllowedApRadioAntennaModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ap_radio_antenna_mode value
func (v ApRadioAntennaMode) Ptr() *ApRadioAntennaMode {
	return &v
}

type NullableApRadioAntennaMode struct {
	value *ApRadioAntennaMode
	isSet bool
}

func (v NullableApRadioAntennaMode) Get() *ApRadioAntennaMode {
	return v.value
}

func (v *NullableApRadioAntennaMode) Set(val *ApRadioAntennaMode) {
	v.value = val
	v.isSet = true
}

func (v NullableApRadioAntennaMode) IsSet() bool {
	return v.isSet
}

func (v *NullableApRadioAntennaMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApRadioAntennaMode(val *ApRadioAntennaMode) *NullableApRadioAntennaMode {
	return &NullableApRadioAntennaMode{value: val, isSet: true}
}

func (v NullableApRadioAntennaMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApRadioAntennaMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// ApUsbType usb config type
type ApUsbType string

// List of ap_usb_type
const (
	APUSBTYPE_EMPTY ApUsbType = ""
	APUSBTYPE_IMAGOTAG ApUsbType = "imagotag"
	APUSBTYPE_SOLUM ApUsbType = "solum"
	APUSBTYPE_HANSHOW ApUsbType = "hanshow"
)

// All allowed values of ApUsbType enum
var AllowedApUsbTypeEnumValues = []ApUsbType{
	"",
	"imagotag",
	"solum",
	"hanshow",
}

func (v *ApUsbType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApUsbType(value)
	for _, existing := range AllowedApUsbTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApUsbType", value)
}

// NewApUsbTypeFromValue returns a pointer to a valid ApUsbType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApUsbTypeFromValue(v string) (*ApUsbType, error) {
	ev := ApUsbType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApUsbType: valid values are %v", v, AllowedApUsbTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApUsbType) IsValid() bool {
	for _, existing := range AllowedApUsbTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ap_usb_type value
func (v ApUsbType) Ptr() *ApUsbType {
	return &v
}

type NullableApUsbType struct {
	value *ApUsbType
	isSet bool
}

func (v NullableApUsbType) Get() *ApUsbType {
	return v.value
}

func (v *NullableApUsbType) Set(val *ApUsbType) {
	v.value = val
	v.isSet = true
}

func (v NullableApUsbType) IsSet() bool {
	return v.isSet
}

func (v *NullableApUsbType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApUsbType(val *ApUsbType) *NullableApUsbType {
	return &NullableApUsbType{value: val, isSet: true}
}

func (v NullableApUsbType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApUsbType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


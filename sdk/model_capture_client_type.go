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

// CaptureClientType client
type CaptureClientType string

// List of capture_client_type
const (
	CAPTURECLIENTTYPE_EMPTY CaptureClientType = ""
	CAPTURECLIENTTYPE_CLIENT CaptureClientType = "client"
)

// All allowed values of CaptureClientType enum
var AllowedCaptureClientTypeEnumValues = []CaptureClientType{
	"",
	"client",
}

func (v *CaptureClientType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaptureClientType(value)
	for _, existing := range AllowedCaptureClientTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaptureClientType", value)
}

// NewCaptureClientTypeFromValue returns a pointer to a valid CaptureClientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaptureClientTypeFromValue(v string) (*CaptureClientType, error) {
	ev := CaptureClientType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaptureClientType: valid values are %v", v, AllowedCaptureClientTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaptureClientType) IsValid() bool {
	for _, existing := range AllowedCaptureClientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to capture_client_type value
func (v CaptureClientType) Ptr() *CaptureClientType {
	return &v
}

type NullableCaptureClientType struct {
	value *CaptureClientType
	isSet bool
}

func (v NullableCaptureClientType) Get() *CaptureClientType {
	return v.value
}

func (v *NullableCaptureClientType) Set(val *CaptureClientType) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureClientType) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureClientType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureClientType(val *CaptureClientType) *NullableCaptureClientType {
	return &NullableCaptureClientType{value: val, isSet: true}
}

func (v NullableCaptureClientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureClientType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


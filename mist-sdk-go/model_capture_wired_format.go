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

// CaptureWiredFormat pcap format
type CaptureWiredFormat string

// List of capture_wired_format
const (
	CAPTUREWIREDFORMAT_EMPTY CaptureWiredFormat = ""
	CAPTUREWIREDFORMAT_PCAP CaptureWiredFormat = "pcap"
	CAPTUREWIREDFORMAT_STREAM CaptureWiredFormat = "stream"
)

// All allowed values of CaptureWiredFormat enum
var AllowedCaptureWiredFormatEnumValues = []CaptureWiredFormat{
	"",
	"pcap",
	"stream",
}

func (v *CaptureWiredFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaptureWiredFormat(value)
	for _, existing := range AllowedCaptureWiredFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaptureWiredFormat", value)
}

// NewCaptureWiredFormatFromValue returns a pointer to a valid CaptureWiredFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaptureWiredFormatFromValue(v string) (*CaptureWiredFormat, error) {
	ev := CaptureWiredFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaptureWiredFormat: valid values are %v", v, AllowedCaptureWiredFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaptureWiredFormat) IsValid() bool {
	for _, existing := range AllowedCaptureWiredFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to capture_wired_format value
func (v CaptureWiredFormat) Ptr() *CaptureWiredFormat {
	return &v
}

type NullableCaptureWiredFormat struct {
	value *CaptureWiredFormat
	isSet bool
}

func (v NullableCaptureWiredFormat) Get() *CaptureWiredFormat {
	return v.value
}

func (v *NullableCaptureWiredFormat) Set(val *CaptureWiredFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureWiredFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureWiredFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureWiredFormat(val *CaptureWiredFormat) *NullableCaptureWiredFormat {
	return &NullableCaptureWiredFormat{value: val, isSet: true}
}

func (v NullableCaptureWiredFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureWiredFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

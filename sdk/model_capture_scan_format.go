/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// CaptureScanFormat the model 'CaptureScanFormat'
type CaptureScanFormat string

// List of capture_scan_format
const (
	CAPTURESCANFORMAT_EMPTY CaptureScanFormat = ""
	CAPTURESCANFORMAT_PCAP CaptureScanFormat = "pcap"
	CAPTURESCANFORMAT_STREAM CaptureScanFormat = "stream"
)

// All allowed values of CaptureScanFormat enum
var AllowedCaptureScanFormatEnumValues = []CaptureScanFormat{
	"",
	"pcap",
	"stream",
}

func (v *CaptureScanFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaptureScanFormat(value)
	for _, existing := range AllowedCaptureScanFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaptureScanFormat", value)
}

// NewCaptureScanFormatFromValue returns a pointer to a valid CaptureScanFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaptureScanFormatFromValue(v string) (*CaptureScanFormat, error) {
	ev := CaptureScanFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaptureScanFormat: valid values are %v", v, AllowedCaptureScanFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaptureScanFormat) IsValid() bool {
	for _, existing := range AllowedCaptureScanFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to capture_scan_format value
func (v CaptureScanFormat) Ptr() *CaptureScanFormat {
	return &v
}

type NullableCaptureScanFormat struct {
	value *CaptureScanFormat
	isSet bool
}

func (v NullableCaptureScanFormat) Get() *CaptureScanFormat {
	return v.value
}

func (v *NullableCaptureScanFormat) Set(val *CaptureScanFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureScanFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureScanFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureScanFormat(val *CaptureScanFormat) *NullableCaptureScanFormat {
	return &NullableCaptureScanFormat{value: val, isSet: true}
}

func (v NullableCaptureScanFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureScanFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


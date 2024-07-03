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

// MxclusterRadAuthServerKeywrapFormat if used for Mist APs
type MxclusterRadAuthServerKeywrapFormat string

// List of mxcluster_rad_auth_server_keywrap_format
const (
	MXCLUSTERRADAUTHSERVERKEYWRAPFORMAT_EMPTY MxclusterRadAuthServerKeywrapFormat = ""
	MXCLUSTERRADAUTHSERVERKEYWRAPFORMAT_HEX MxclusterRadAuthServerKeywrapFormat = "hex"
	MXCLUSTERRADAUTHSERVERKEYWRAPFORMAT_ASCII MxclusterRadAuthServerKeywrapFormat = "ascii"
)

// All allowed values of MxclusterRadAuthServerKeywrapFormat enum
var AllowedMxclusterRadAuthServerKeywrapFormatEnumValues = []MxclusterRadAuthServerKeywrapFormat{
	"",
	"hex",
	"ascii",
}

func (v *MxclusterRadAuthServerKeywrapFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MxclusterRadAuthServerKeywrapFormat(value)
	for _, existing := range AllowedMxclusterRadAuthServerKeywrapFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MxclusterRadAuthServerKeywrapFormat", value)
}

// NewMxclusterRadAuthServerKeywrapFormatFromValue returns a pointer to a valid MxclusterRadAuthServerKeywrapFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMxclusterRadAuthServerKeywrapFormatFromValue(v string) (*MxclusterRadAuthServerKeywrapFormat, error) {
	ev := MxclusterRadAuthServerKeywrapFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MxclusterRadAuthServerKeywrapFormat: valid values are %v", v, AllowedMxclusterRadAuthServerKeywrapFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MxclusterRadAuthServerKeywrapFormat) IsValid() bool {
	for _, existing := range AllowedMxclusterRadAuthServerKeywrapFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to mxcluster_rad_auth_server_keywrap_format value
func (v MxclusterRadAuthServerKeywrapFormat) Ptr() *MxclusterRadAuthServerKeywrapFormat {
	return &v
}

type NullableMxclusterRadAuthServerKeywrapFormat struct {
	value *MxclusterRadAuthServerKeywrapFormat
	isSet bool
}

func (v NullableMxclusterRadAuthServerKeywrapFormat) Get() *MxclusterRadAuthServerKeywrapFormat {
	return v.value
}

func (v *NullableMxclusterRadAuthServerKeywrapFormat) Set(val *MxclusterRadAuthServerKeywrapFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableMxclusterRadAuthServerKeywrapFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableMxclusterRadAuthServerKeywrapFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxclusterRadAuthServerKeywrapFormat(val *MxclusterRadAuthServerKeywrapFormat) *NullableMxclusterRadAuthServerKeywrapFormat {
	return &NullableMxclusterRadAuthServerKeywrapFormat{value: val, isSet: true}
}

func (v NullableMxclusterRadAuthServerKeywrapFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxclusterRadAuthServerKeywrapFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


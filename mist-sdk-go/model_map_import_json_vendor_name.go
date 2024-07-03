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

// MapImportJsonVendorName the model 'MapImportJsonVendorName'
type MapImportJsonVendorName string

// List of map_import_json_vendor_name
const (
	MAPIMPORTJSONVENDORNAME_EMPTY MapImportJsonVendorName = ""
	MAPIMPORTJSONVENDORNAME_EKAHAU MapImportJsonVendorName = "ekahau"
	MAPIMPORTJSONVENDORNAME_IBWAVE MapImportJsonVendorName = "ibwave"
)

// All allowed values of MapImportJsonVendorName enum
var AllowedMapImportJsonVendorNameEnumValues = []MapImportJsonVendorName{
	"",
	"ekahau",
	"ibwave",
}

func (v *MapImportJsonVendorName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MapImportJsonVendorName(value)
	for _, existing := range AllowedMapImportJsonVendorNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MapImportJsonVendorName", value)
}

// NewMapImportJsonVendorNameFromValue returns a pointer to a valid MapImportJsonVendorName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMapImportJsonVendorNameFromValue(v string) (*MapImportJsonVendorName, error) {
	ev := MapImportJsonVendorName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MapImportJsonVendorName: valid values are %v", v, AllowedMapImportJsonVendorNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MapImportJsonVendorName) IsValid() bool {
	for _, existing := range AllowedMapImportJsonVendorNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to map_import_json_vendor_name value
func (v MapImportJsonVendorName) Ptr() *MapImportJsonVendorName {
	return &v
}

type NullableMapImportJsonVendorName struct {
	value *MapImportJsonVendorName
	isSet bool
}

func (v NullableMapImportJsonVendorName) Get() *MapImportJsonVendorName {
	return v.value
}

func (v *NullableMapImportJsonVendorName) Set(val *MapImportJsonVendorName) {
	v.value = val
	v.isSet = true
}

func (v NullableMapImportJsonVendorName) IsSet() bool {
	return v.isSet
}

func (v *NullableMapImportJsonVendorName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapImportJsonVendorName(val *MapImportJsonVendorName) *NullableMapImportJsonVendorName {
	return &NullableMapImportJsonVendorName{value: val, isSet: true}
}

func (v NullableMapImportJsonVendorName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapImportJsonVendorName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


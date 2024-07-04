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

// WxlanTagOperation the model 'WxlanTagOperation'
type WxlanTagOperation string

// List of wxlan_tag_operation
const (
	WXLANTAGOPERATION_EMPTY WxlanTagOperation = ""
	WXLANTAGOPERATION_IN WxlanTagOperation = "in"
	WXLANTAGOPERATION_NOT_IN WxlanTagOperation = "not_in"
)

// All allowed values of WxlanTagOperation enum
var AllowedWxlanTagOperationEnumValues = []WxlanTagOperation{
	"",
	"in",
	"not_in",
}

func (v *WxlanTagOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WxlanTagOperation(value)
	for _, existing := range AllowedWxlanTagOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WxlanTagOperation", value)
}

// NewWxlanTagOperationFromValue returns a pointer to a valid WxlanTagOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWxlanTagOperationFromValue(v string) (*WxlanTagOperation, error) {
	ev := WxlanTagOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WxlanTagOperation: valid values are %v", v, AllowedWxlanTagOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WxlanTagOperation) IsValid() bool {
	for _, existing := range AllowedWxlanTagOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wxlan_tag_operation value
func (v WxlanTagOperation) Ptr() *WxlanTagOperation {
	return &v
}

type NullableWxlanTagOperation struct {
	value *WxlanTagOperation
	isSet bool
}

func (v NullableWxlanTagOperation) Get() *WxlanTagOperation {
	return v.value
}

func (v *NullableWxlanTagOperation) Set(val *WxlanTagOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableWxlanTagOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableWxlanTagOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWxlanTagOperation(val *WxlanTagOperation) *NullableWxlanTagOperation {
	return &NullableWxlanTagOperation{value: val, isSet: true}
}

func (v NullableWxlanTagOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWxlanTagOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


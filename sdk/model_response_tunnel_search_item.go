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

// ResponseTunnelSearchItem struct for ResponseTunnelSearchItem
type ResponseTunnelSearchItem struct {
	MxtunnelStats *MxtunnelStats
	WanTunnelStats *WanTunnelStats
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResponseTunnelSearchItem) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MxtunnelStats
	err = json.Unmarshal(data, &dst.MxtunnelStats);
	if err == nil {
		jsonMxtunnelStats, _ := json.Marshal(dst.MxtunnelStats)
		if string(jsonMxtunnelStats) == "{}" { // empty struct
			dst.MxtunnelStats = nil
		} else {
			return nil // data stored in dst.MxtunnelStats, return on the first match
		}
	} else {
		dst.MxtunnelStats = nil
	}

	// try to unmarshal JSON data into WanTunnelStats
	err = json.Unmarshal(data, &dst.WanTunnelStats);
	if err == nil {
		jsonWanTunnelStats, _ := json.Marshal(dst.WanTunnelStats)
		if string(jsonWanTunnelStats) == "{}" { // empty struct
			dst.WanTunnelStats = nil
		} else {
			return nil // data stored in dst.WanTunnelStats, return on the first match
		}
	} else {
		dst.WanTunnelStats = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ResponseTunnelSearchItem)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResponseTunnelSearchItem) MarshalJSON() ([]byte, error) {
	if src.MxtunnelStats != nil {
		return json.Marshal(&src.MxtunnelStats)
	}

	if src.WanTunnelStats != nil {
		return json.Marshal(&src.WanTunnelStats)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResponseTunnelSearchItem struct {
	value *ResponseTunnelSearchItem
	isSet bool
}

func (v NullableResponseTunnelSearchItem) Get() *ResponseTunnelSearchItem {
	return v.value
}

func (v *NullableResponseTunnelSearchItem) Set(val *ResponseTunnelSearchItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTunnelSearchItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTunnelSearchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTunnelSearchItem(val *ResponseTunnelSearchItem) *NullableResponseTunnelSearchItem {
	return &NullableResponseTunnelSearchItem{value: val, isSet: true}
}

func (v NullableResponseTunnelSearchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTunnelSearchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



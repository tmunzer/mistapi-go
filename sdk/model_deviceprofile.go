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

// Deviceprofile struct for Deviceprofile
type Deviceprofile struct {
	DeviceprofileAp *DeviceprofileAp
	GatewayTemplate *GatewayTemplate
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Deviceprofile) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DeviceprofileAp
	err = json.Unmarshal(data, &dst.DeviceprofileAp);
	if err == nil {
		jsonDeviceprofileAp, _ := json.Marshal(dst.DeviceprofileAp)
		if string(jsonDeviceprofileAp) == "{}" { // empty struct
			dst.DeviceprofileAp = nil
		} else {
			return nil // data stored in dst.DeviceprofileAp, return on the first match
		}
	} else {
		dst.DeviceprofileAp = nil
	}

	// try to unmarshal JSON data into GatewayTemplate
	err = json.Unmarshal(data, &dst.GatewayTemplate);
	if err == nil {
		jsonGatewayTemplate, _ := json.Marshal(dst.GatewayTemplate)
		if string(jsonGatewayTemplate) == "{}" { // empty struct
			dst.GatewayTemplate = nil
		} else {
			return nil // data stored in dst.GatewayTemplate, return on the first match
		}
	} else {
		dst.GatewayTemplate = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Deviceprofile)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Deviceprofile) MarshalJSON() ([]byte, error) {
	if src.DeviceprofileAp != nil {
		return json.Marshal(&src.DeviceprofileAp)
	}

	if src.GatewayTemplate != nil {
		return json.Marshal(&src.GatewayTemplate)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDeviceprofile struct {
	value *Deviceprofile
	isSet bool
}

func (v NullableDeviceprofile) Get() *Deviceprofile {
	return v.value
}

func (v *NullableDeviceprofile) Set(val *Deviceprofile) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceprofile) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceprofile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceprofile(val *Deviceprofile) *NullableDeviceprofile {
	return &NullableDeviceprofile{value: val, isSet: true}
}

func (v NullableDeviceprofile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceprofile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



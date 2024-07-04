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

// GatewayClusterSwapOp the model 'GatewayClusterSwapOp'
type GatewayClusterSwapOp string

// List of gateway_cluster_swap_op
const (
	GATEWAYCLUSTERSWAPOP_EMPTY GatewayClusterSwapOp = ""
	GATEWAYCLUSTERSWAPOP_SWAP GatewayClusterSwapOp = "swap"
	GATEWAYCLUSTERSWAPOP_REPLACE_NODE1 GatewayClusterSwapOp = "replace_node1"
)

// All allowed values of GatewayClusterSwapOp enum
var AllowedGatewayClusterSwapOpEnumValues = []GatewayClusterSwapOp{
	"",
	"swap",
	"replace_node1",
}

func (v *GatewayClusterSwapOp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayClusterSwapOp(value)
	for _, existing := range AllowedGatewayClusterSwapOpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayClusterSwapOp", value)
}

// NewGatewayClusterSwapOpFromValue returns a pointer to a valid GatewayClusterSwapOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayClusterSwapOpFromValue(v string) (*GatewayClusterSwapOp, error) {
	ev := GatewayClusterSwapOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayClusterSwapOp: valid values are %v", v, AllowedGatewayClusterSwapOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayClusterSwapOp) IsValid() bool {
	for _, existing := range AllowedGatewayClusterSwapOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_cluster_swap_op value
func (v GatewayClusterSwapOp) Ptr() *GatewayClusterSwapOp {
	return &v
}

type NullableGatewayClusterSwapOp struct {
	value *GatewayClusterSwapOp
	isSet bool
}

func (v NullableGatewayClusterSwapOp) Get() *GatewayClusterSwapOp {
	return v.value
}

func (v *NullableGatewayClusterSwapOp) Set(val *GatewayClusterSwapOp) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayClusterSwapOp) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayClusterSwapOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayClusterSwapOp(val *GatewayClusterSwapOp) *NullableGatewayClusterSwapOp {
	return &NullableGatewayClusterSwapOp{value: val, isSet: true}
}

func (v NullableGatewayClusterSwapOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayClusterSwapOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


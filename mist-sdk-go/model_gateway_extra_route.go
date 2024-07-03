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
)

// checks if the GatewayExtraRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayExtraRoute{}

// GatewayExtraRoute Property key is the destination CIDR (e.g. \"10.0.0.0/8\")
type GatewayExtraRoute struct {
	Via *string `json:"via,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayExtraRoute GatewayExtraRoute

// NewGatewayExtraRoute instantiates a new GatewayExtraRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayExtraRoute() *GatewayExtraRoute {
	this := GatewayExtraRoute{}
	return &this
}

// NewGatewayExtraRouteWithDefaults instantiates a new GatewayExtraRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayExtraRouteWithDefaults() *GatewayExtraRoute {
	this := GatewayExtraRoute{}
	return &this
}

// GetVia returns the Via field value if set, zero value otherwise.
func (o *GatewayExtraRoute) GetVia() string {
	if o == nil || IsNil(o.Via) {
		var ret string
		return ret
	}
	return *o.Via
}

// GetViaOk returns a tuple with the Via field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayExtraRoute) GetViaOk() (*string, bool) {
	if o == nil || IsNil(o.Via) {
		return nil, false
	}
	return o.Via, true
}

// HasVia returns a boolean if a field has been set.
func (o *GatewayExtraRoute) HasVia() bool {
	if o != nil && !IsNil(o.Via) {
		return true
	}

	return false
}

// SetVia gets a reference to the given string and assigns it to the Via field.
func (o *GatewayExtraRoute) SetVia(v string) {
	o.Via = &v
}

func (o GatewayExtraRoute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayExtraRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Via) {
		toSerialize["via"] = o.Via
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayExtraRoute) UnmarshalJSON(data []byte) (err error) {
	varGatewayExtraRoute := _GatewayExtraRoute{}

	err = json.Unmarshal(data, &varGatewayExtraRoute)

	if err != nil {
		return err
	}

	*o = GatewayExtraRoute(varGatewayExtraRoute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "via")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayExtraRoute struct {
	value *GatewayExtraRoute
	isSet bool
}

func (v NullableGatewayExtraRoute) Get() *GatewayExtraRoute {
	return v.value
}

func (v *NullableGatewayExtraRoute) Set(val *GatewayExtraRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayExtraRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayExtraRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayExtraRoute(val *GatewayExtraRoute) *NullableGatewayExtraRoute {
	return &NullableGatewayExtraRoute{value: val, isSet: true}
}

func (v NullableGatewayExtraRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayExtraRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



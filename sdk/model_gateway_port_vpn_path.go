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

// checks if the GatewayPortVpnPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayPortVpnPath{}

// GatewayPortVpnPath struct for GatewayPortVpnPath
type GatewayPortVpnPath struct {
	BfdProfile *GatewayPortVpnPathBfdProfile `json:"bfd_profile,omitempty"`
	// whether to use tunnel mode. SSR only
	BfdUseTunnelMode *bool `json:"bfd_use_tunnel_mode,omitempty"`
	// for a given VPN, when `path_selection.strategy`==`simple`, the preference for a path (lower is preferred)
	Preference *int32 `json:"preference,omitempty"`
	Role *GatewayPortVpnPathRole `json:"role,omitempty"`
	TrafficShaping *GatewayTrafficShaping `json:"traffic_shaping,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayPortVpnPath GatewayPortVpnPath

// NewGatewayPortVpnPath instantiates a new GatewayPortVpnPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayPortVpnPath() *GatewayPortVpnPath {
	this := GatewayPortVpnPath{}
	var bfdProfile GatewayPortVpnPathBfdProfile = GATEWAYPORTVPNPATHBFDPROFILE_BROADBAND
	this.BfdProfile = &bfdProfile
	var bfdUseTunnelMode bool = false
	this.BfdUseTunnelMode = &bfdUseTunnelMode
	var role GatewayPortVpnPathRole = GATEWAYPORTVPNPATHROLE_SPOKE
	this.Role = &role
	return &this
}

// NewGatewayPortVpnPathWithDefaults instantiates a new GatewayPortVpnPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayPortVpnPathWithDefaults() *GatewayPortVpnPath {
	this := GatewayPortVpnPath{}
	var bfdProfile GatewayPortVpnPathBfdProfile = GATEWAYPORTVPNPATHBFDPROFILE_BROADBAND
	this.BfdProfile = &bfdProfile
	var bfdUseTunnelMode bool = false
	this.BfdUseTunnelMode = &bfdUseTunnelMode
	var role GatewayPortVpnPathRole = GATEWAYPORTVPNPATHROLE_SPOKE
	this.Role = &role
	return &this
}

// GetBfdProfile returns the BfdProfile field value if set, zero value otherwise.
func (o *GatewayPortVpnPath) GetBfdProfile() GatewayPortVpnPathBfdProfile {
	if o == nil || IsNil(o.BfdProfile) {
		var ret GatewayPortVpnPathBfdProfile
		return ret
	}
	return *o.BfdProfile
}

// GetBfdProfileOk returns a tuple with the BfdProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortVpnPath) GetBfdProfileOk() (*GatewayPortVpnPathBfdProfile, bool) {
	if o == nil || IsNil(o.BfdProfile) {
		return nil, false
	}
	return o.BfdProfile, true
}

// HasBfdProfile returns a boolean if a field has been set.
func (o *GatewayPortVpnPath) HasBfdProfile() bool {
	if o != nil && !IsNil(o.BfdProfile) {
		return true
	}

	return false
}

// SetBfdProfile gets a reference to the given GatewayPortVpnPathBfdProfile and assigns it to the BfdProfile field.
func (o *GatewayPortVpnPath) SetBfdProfile(v GatewayPortVpnPathBfdProfile) {
	o.BfdProfile = &v
}

// GetBfdUseTunnelMode returns the BfdUseTunnelMode field value if set, zero value otherwise.
func (o *GatewayPortVpnPath) GetBfdUseTunnelMode() bool {
	if o == nil || IsNil(o.BfdUseTunnelMode) {
		var ret bool
		return ret
	}
	return *o.BfdUseTunnelMode
}

// GetBfdUseTunnelModeOk returns a tuple with the BfdUseTunnelMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortVpnPath) GetBfdUseTunnelModeOk() (*bool, bool) {
	if o == nil || IsNil(o.BfdUseTunnelMode) {
		return nil, false
	}
	return o.BfdUseTunnelMode, true
}

// HasBfdUseTunnelMode returns a boolean if a field has been set.
func (o *GatewayPortVpnPath) HasBfdUseTunnelMode() bool {
	if o != nil && !IsNil(o.BfdUseTunnelMode) {
		return true
	}

	return false
}

// SetBfdUseTunnelMode gets a reference to the given bool and assigns it to the BfdUseTunnelMode field.
func (o *GatewayPortVpnPath) SetBfdUseTunnelMode(v bool) {
	o.BfdUseTunnelMode = &v
}

// GetPreference returns the Preference field value if set, zero value otherwise.
func (o *GatewayPortVpnPath) GetPreference() int32 {
	if o == nil || IsNil(o.Preference) {
		var ret int32
		return ret
	}
	return *o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortVpnPath) GetPreferenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Preference) {
		return nil, false
	}
	return o.Preference, true
}

// HasPreference returns a boolean if a field has been set.
func (o *GatewayPortVpnPath) HasPreference() bool {
	if o != nil && !IsNil(o.Preference) {
		return true
	}

	return false
}

// SetPreference gets a reference to the given int32 and assigns it to the Preference field.
func (o *GatewayPortVpnPath) SetPreference(v int32) {
	o.Preference = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GatewayPortVpnPath) GetRole() GatewayPortVpnPathRole {
	if o == nil || IsNil(o.Role) {
		var ret GatewayPortVpnPathRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortVpnPath) GetRoleOk() (*GatewayPortVpnPathRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GatewayPortVpnPath) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given GatewayPortVpnPathRole and assigns it to the Role field.
func (o *GatewayPortVpnPath) SetRole(v GatewayPortVpnPathRole) {
	o.Role = &v
}

// GetTrafficShaping returns the TrafficShaping field value if set, zero value otherwise.
func (o *GatewayPortVpnPath) GetTrafficShaping() GatewayTrafficShaping {
	if o == nil || IsNil(o.TrafficShaping) {
		var ret GatewayTrafficShaping
		return ret
	}
	return *o.TrafficShaping
}

// GetTrafficShapingOk returns a tuple with the TrafficShaping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPortVpnPath) GetTrafficShapingOk() (*GatewayTrafficShaping, bool) {
	if o == nil || IsNil(o.TrafficShaping) {
		return nil, false
	}
	return o.TrafficShaping, true
}

// HasTrafficShaping returns a boolean if a field has been set.
func (o *GatewayPortVpnPath) HasTrafficShaping() bool {
	if o != nil && !IsNil(o.TrafficShaping) {
		return true
	}

	return false
}

// SetTrafficShaping gets a reference to the given GatewayTrafficShaping and assigns it to the TrafficShaping field.
func (o *GatewayPortVpnPath) SetTrafficShaping(v GatewayTrafficShaping) {
	o.TrafficShaping = &v
}

func (o GatewayPortVpnPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayPortVpnPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BfdProfile) {
		toSerialize["bfd_profile"] = o.BfdProfile
	}
	if !IsNil(o.BfdUseTunnelMode) {
		toSerialize["bfd_use_tunnel_mode"] = o.BfdUseTunnelMode
	}
	if !IsNil(o.Preference) {
		toSerialize["preference"] = o.Preference
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.TrafficShaping) {
		toSerialize["traffic_shaping"] = o.TrafficShaping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayPortVpnPath) UnmarshalJSON(data []byte) (err error) {
	varGatewayPortVpnPath := _GatewayPortVpnPath{}

	err = json.Unmarshal(data, &varGatewayPortVpnPath)

	if err != nil {
		return err
	}

	*o = GatewayPortVpnPath(varGatewayPortVpnPath)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bfd_profile")
		delete(additionalProperties, "bfd_use_tunnel_mode")
		delete(additionalProperties, "preference")
		delete(additionalProperties, "role")
		delete(additionalProperties, "traffic_shaping")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayPortVpnPath struct {
	value *GatewayPortVpnPath
	isSet bool
}

func (v NullableGatewayPortVpnPath) Get() *GatewayPortVpnPath {
	return v.value
}

func (v *NullableGatewayPortVpnPath) Set(val *GatewayPortVpnPath) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPortVpnPath) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPortVpnPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPortVpnPath(val *GatewayPortVpnPath) *NullableGatewayPortVpnPath {
	return &NullableGatewayPortVpnPath{value: val, isSet: true}
}

func (v NullableGatewayPortVpnPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPortVpnPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



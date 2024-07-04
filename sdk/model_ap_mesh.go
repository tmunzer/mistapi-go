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
)

// checks if the ApMesh type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApMesh{}

// ApMesh Mesh AP settings
type ApMesh struct {
	// whether mesh is enabled on this AP
	Enabled *bool `json:"enabled,omitempty"`
	// mesh group, base AP(s) will only allow remote AP(s) in the same mesh group to join, 1-9, optional
	Group NullableInt32 `json:"group,omitempty"`
	Role *ApMeshRole `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApMesh ApMesh

// NewApMesh instantiates a new ApMesh object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApMesh() *ApMesh {
	this := ApMesh{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewApMeshWithDefaults instantiates a new ApMesh object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApMeshWithDefaults() *ApMesh {
	this := ApMesh{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApMesh) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApMesh) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApMesh) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApMesh) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApMesh) GetGroup() int32 {
	if o == nil || IsNil(o.Group.Get()) {
		var ret int32
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApMesh) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *ApMesh) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableInt32 and assigns it to the Group field.
func (o *ApMesh) SetGroup(v int32) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *ApMesh) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *ApMesh) UnsetGroup() {
	o.Group.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ApMesh) GetRole() ApMeshRole {
	if o == nil || IsNil(o.Role) {
		var ret ApMeshRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApMesh) GetRoleOk() (*ApMeshRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ApMesh) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given ApMeshRole and assigns it to the Role field.
func (o *ApMesh) SetRole(v ApMeshRole) {
	o.Role = &v
}

func (o ApMesh) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApMesh) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApMesh) UnmarshalJSON(data []byte) (err error) {
	varApMesh := _ApMesh{}

	err = json.Unmarshal(data, &varApMesh)

	if err != nil {
		return err
	}

	*o = ApMesh(varApMesh)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "group")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApMesh struct {
	value *ApMesh
	isSet bool
}

func (v NullableApMesh) Get() *ApMesh {
	return v.value
}

func (v *NullableApMesh) Set(val *ApMesh) {
	v.value = val
	v.isSet = true
}

func (v NullableApMesh) IsSet() bool {
	return v.isSet
}

func (v *NullableApMesh) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApMesh(val *ApMesh) *NullableApMesh {
	return &NullableApMesh{value: val, isSet: true}
}

func (v NullableApMesh) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApMesh) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



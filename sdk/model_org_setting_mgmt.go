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

// checks if the OrgSettingMgmt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingMgmt{}

// OrgSettingMgmt management-related properties
type OrgSettingMgmt struct {
	// list of Mist Tunnels
	MxtunnelIds []string `json:"mxtunnel_ids,omitempty"`
	// whether to use Mist Tunnel for mgmt connectivity, this takes precedence over use_wxtunnel
	UseMxtunnel *bool `json:"use_mxtunnel,omitempty"`
	// whether to use wxtunnel for mgmt connectivity
	UseWxtunnel *bool `json:"use_wxtunnel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingMgmt OrgSettingMgmt

// NewOrgSettingMgmt instantiates a new OrgSettingMgmt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingMgmt() *OrgSettingMgmt {
	this := OrgSettingMgmt{}
	var useMxtunnel bool = false
	this.UseMxtunnel = &useMxtunnel
	var useWxtunnel bool = false
	this.UseWxtunnel = &useWxtunnel
	return &this
}

// NewOrgSettingMgmtWithDefaults instantiates a new OrgSettingMgmt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingMgmtWithDefaults() *OrgSettingMgmt {
	this := OrgSettingMgmt{}
	var useMxtunnel bool = false
	this.UseMxtunnel = &useMxtunnel
	var useWxtunnel bool = false
	this.UseWxtunnel = &useWxtunnel
	return &this
}

// GetMxtunnelIds returns the MxtunnelIds field value if set, zero value otherwise.
func (o *OrgSettingMgmt) GetMxtunnelIds() []string {
	if o == nil || IsNil(o.MxtunnelIds) {
		var ret []string
		return ret
	}
	return o.MxtunnelIds
}

// GetMxtunnelIdsOk returns a tuple with the MxtunnelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingMgmt) GetMxtunnelIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MxtunnelIds) {
		return nil, false
	}
	return o.MxtunnelIds, true
}

// HasMxtunnelIds returns a boolean if a field has been set.
func (o *OrgSettingMgmt) HasMxtunnelIds() bool {
	if o != nil && !IsNil(o.MxtunnelIds) {
		return true
	}

	return false
}

// SetMxtunnelIds gets a reference to the given []string and assigns it to the MxtunnelIds field.
func (o *OrgSettingMgmt) SetMxtunnelIds(v []string) {
	o.MxtunnelIds = v
}

// GetUseMxtunnel returns the UseMxtunnel field value if set, zero value otherwise.
func (o *OrgSettingMgmt) GetUseMxtunnel() bool {
	if o == nil || IsNil(o.UseMxtunnel) {
		var ret bool
		return ret
	}
	return *o.UseMxtunnel
}

// GetUseMxtunnelOk returns a tuple with the UseMxtunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingMgmt) GetUseMxtunnelOk() (*bool, bool) {
	if o == nil || IsNil(o.UseMxtunnel) {
		return nil, false
	}
	return o.UseMxtunnel, true
}

// HasUseMxtunnel returns a boolean if a field has been set.
func (o *OrgSettingMgmt) HasUseMxtunnel() bool {
	if o != nil && !IsNil(o.UseMxtunnel) {
		return true
	}

	return false
}

// SetUseMxtunnel gets a reference to the given bool and assigns it to the UseMxtunnel field.
func (o *OrgSettingMgmt) SetUseMxtunnel(v bool) {
	o.UseMxtunnel = &v
}

// GetUseWxtunnel returns the UseWxtunnel field value if set, zero value otherwise.
func (o *OrgSettingMgmt) GetUseWxtunnel() bool {
	if o == nil || IsNil(o.UseWxtunnel) {
		var ret bool
		return ret
	}
	return *o.UseWxtunnel
}

// GetUseWxtunnelOk returns a tuple with the UseWxtunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingMgmt) GetUseWxtunnelOk() (*bool, bool) {
	if o == nil || IsNil(o.UseWxtunnel) {
		return nil, false
	}
	return o.UseWxtunnel, true
}

// HasUseWxtunnel returns a boolean if a field has been set.
func (o *OrgSettingMgmt) HasUseWxtunnel() bool {
	if o != nil && !IsNil(o.UseWxtunnel) {
		return true
	}

	return false
}

// SetUseWxtunnel gets a reference to the given bool and assigns it to the UseWxtunnel field.
func (o *OrgSettingMgmt) SetUseWxtunnel(v bool) {
	o.UseWxtunnel = &v
}

func (o OrgSettingMgmt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingMgmt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MxtunnelIds) {
		toSerialize["mxtunnel_ids"] = o.MxtunnelIds
	}
	if !IsNil(o.UseMxtunnel) {
		toSerialize["use_mxtunnel"] = o.UseMxtunnel
	}
	if !IsNil(o.UseWxtunnel) {
		toSerialize["use_wxtunnel"] = o.UseWxtunnel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingMgmt) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingMgmt := _OrgSettingMgmt{}

	err = json.Unmarshal(data, &varOrgSettingMgmt)

	if err != nil {
		return err
	}

	*o = OrgSettingMgmt(varOrgSettingMgmt)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mxtunnel_ids")
		delete(additionalProperties, "use_mxtunnel")
		delete(additionalProperties, "use_wxtunnel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingMgmt struct {
	value *OrgSettingMgmt
	isSet bool
}

func (v NullableOrgSettingMgmt) Get() *OrgSettingMgmt {
	return v.value
}

func (v *NullableOrgSettingMgmt) Set(val *OrgSettingMgmt) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingMgmt) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingMgmt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingMgmt(val *OrgSettingMgmt) *NullableOrgSettingMgmt {
	return &NullableOrgSettingMgmt{value: val, isSet: true}
}

func (v NullableOrgSettingMgmt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingMgmt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



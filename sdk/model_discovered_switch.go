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

// checks if the DiscoveredSwitch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveredSwitch{}

// DiscoveredSwitch struct for DiscoveredSwitch
type DiscoveredSwitch struct {
	Adopted *bool `json:"adopted,omitempty"`
	ApRedundancy *ApRedundancy `json:"ap_redundancy,omitempty"`
	Aps []DiscoveredSwitchAp `json:"aps,omitempty"`
	ChassisId []string `json:"chassis_id,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	Model *string `json:"model,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	SystemDesc *string `json:"system_desc,omitempty"`
	SystemName *string `json:"system_name,omitempty"`
	Timestamp *float32 `json:"timestamp,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveredSwitch DiscoveredSwitch

// NewDiscoveredSwitch instantiates a new DiscoveredSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredSwitch() *DiscoveredSwitch {
	this := DiscoveredSwitch{}
	return &this
}

// NewDiscoveredSwitchWithDefaults instantiates a new DiscoveredSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredSwitchWithDefaults() *DiscoveredSwitch {
	this := DiscoveredSwitch{}
	return &this
}

// GetAdopted returns the Adopted field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetAdopted() bool {
	if o == nil || IsNil(o.Adopted) {
		var ret bool
		return ret
	}
	return *o.Adopted
}

// GetAdoptedOk returns a tuple with the Adopted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetAdoptedOk() (*bool, bool) {
	if o == nil || IsNil(o.Adopted) {
		return nil, false
	}
	return o.Adopted, true
}

// HasAdopted returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasAdopted() bool {
	if o != nil && !IsNil(o.Adopted) {
		return true
	}

	return false
}

// SetAdopted gets a reference to the given bool and assigns it to the Adopted field.
func (o *DiscoveredSwitch) SetAdopted(v bool) {
	o.Adopted = &v
}

// GetApRedundancy returns the ApRedundancy field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetApRedundancy() ApRedundancy {
	if o == nil || IsNil(o.ApRedundancy) {
		var ret ApRedundancy
		return ret
	}
	return *o.ApRedundancy
}

// GetApRedundancyOk returns a tuple with the ApRedundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetApRedundancyOk() (*ApRedundancy, bool) {
	if o == nil || IsNil(o.ApRedundancy) {
		return nil, false
	}
	return o.ApRedundancy, true
}

// HasApRedundancy returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasApRedundancy() bool {
	if o != nil && !IsNil(o.ApRedundancy) {
		return true
	}

	return false
}

// SetApRedundancy gets a reference to the given ApRedundancy and assigns it to the ApRedundancy field.
func (o *DiscoveredSwitch) SetApRedundancy(v ApRedundancy) {
	o.ApRedundancy = &v
}

// GetAps returns the Aps field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetAps() []DiscoveredSwitchAp {
	if o == nil || IsNil(o.Aps) {
		var ret []DiscoveredSwitchAp
		return ret
	}
	return o.Aps
}

// GetApsOk returns a tuple with the Aps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetApsOk() ([]DiscoveredSwitchAp, bool) {
	if o == nil || IsNil(o.Aps) {
		return nil, false
	}
	return o.Aps, true
}

// HasAps returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasAps() bool {
	if o != nil && !IsNil(o.Aps) {
		return true
	}

	return false
}

// SetAps gets a reference to the given []DiscoveredSwitchAp and assigns it to the Aps field.
func (o *DiscoveredSwitch) SetAps(v []DiscoveredSwitchAp) {
	o.Aps = v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetChassisId() []string {
	if o == nil || IsNil(o.ChassisId) {
		var ret []string
		return ret
	}
	return o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetChassisIdOk() ([]string, bool) {
	if o == nil || IsNil(o.ChassisId) {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasChassisId() bool {
	if o != nil && !IsNil(o.ChassisId) {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given []string and assigns it to the ChassisId field.
func (o *DiscoveredSwitch) SetChassisId(v []string) {
	o.ChassisId = v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *DiscoveredSwitch) SetForSite(v bool) {
	o.ForSite = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *DiscoveredSwitch) SetModel(v string) {
	o.Model = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *DiscoveredSwitch) SetOrgId(v string) {
	o.OrgId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *DiscoveredSwitch) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSystemDesc returns the SystemDesc field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetSystemDesc() string {
	if o == nil || IsNil(o.SystemDesc) {
		var ret string
		return ret
	}
	return *o.SystemDesc
}

// GetSystemDescOk returns a tuple with the SystemDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetSystemDescOk() (*string, bool) {
	if o == nil || IsNil(o.SystemDesc) {
		return nil, false
	}
	return o.SystemDesc, true
}

// HasSystemDesc returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasSystemDesc() bool {
	if o != nil && !IsNil(o.SystemDesc) {
		return true
	}

	return false
}

// SetSystemDesc gets a reference to the given string and assigns it to the SystemDesc field.
func (o *DiscoveredSwitch) SetSystemDesc(v string) {
	o.SystemDesc = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetSystemName() string {
	if o == nil || IsNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetSystemNameOk() (*string, bool) {
	if o == nil || IsNil(o.SystemName) {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasSystemName() bool {
	if o != nil && !IsNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *DiscoveredSwitch) SetSystemName(v string) {
	o.SystemName = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetTimestamp() float32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret float32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetTimestampOk() (*float32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given float32 and assigns it to the Timestamp field.
func (o *DiscoveredSwitch) SetTimestamp(v float32) {
	o.Timestamp = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *DiscoveredSwitch) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DiscoveredSwitch) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitch) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DiscoveredSwitch) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DiscoveredSwitch) SetVersion(v string) {
	o.Version = &v
}

func (o DiscoveredSwitch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveredSwitch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adopted) {
		toSerialize["adopted"] = o.Adopted
	}
	if !IsNil(o.ApRedundancy) {
		toSerialize["ap_redundancy"] = o.ApRedundancy
	}
	if !IsNil(o.Aps) {
		toSerialize["aps"] = o.Aps
	}
	if !IsNil(o.ChassisId) {
		toSerialize["chassis_id"] = o.ChassisId
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.SystemDesc) {
		toSerialize["system_desc"] = o.SystemDesc
	}
	if !IsNil(o.SystemName) {
		toSerialize["system_name"] = o.SystemName
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveredSwitch) UnmarshalJSON(data []byte) (err error) {
	varDiscoveredSwitch := _DiscoveredSwitch{}

	err = json.Unmarshal(data, &varDiscoveredSwitch)

	if err != nil {
		return err
	}

	*o = DiscoveredSwitch(varDiscoveredSwitch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "adopted")
		delete(additionalProperties, "ap_redundancy")
		delete(additionalProperties, "aps")
		delete(additionalProperties, "chassis_id")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "model")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "system_desc")
		delete(additionalProperties, "system_name")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveredSwitch struct {
	value *DiscoveredSwitch
	isSet bool
}

func (v NullableDiscoveredSwitch) Get() *DiscoveredSwitch {
	return v.value
}

func (v *NullableDiscoveredSwitch) Set(val *DiscoveredSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveredSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveredSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveredSwitch(val *DiscoveredSwitch) *NullableDiscoveredSwitch {
	return &NullableDiscoveredSwitch{value: val, isSet: true}
}

func (v NullableDiscoveredSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveredSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



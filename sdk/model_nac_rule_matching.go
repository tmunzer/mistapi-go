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

// checks if the NacRuleMatching type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NacRuleMatching{}

// NacRuleMatching struct for NacRuleMatching
type NacRuleMatching struct {
	AuthType *NacRuleMatchingAuthType `json:"auth_type,omitempty"`
	Nactags []string `json:"nactags,omitempty"`
	PortTypes []NacRuleMatchingPortType `json:"port_types,omitempty"`
	// list of site ids to match
	SiteIds []string `json:"site_ids,omitempty"`
	// list of sitegroup ids to match
	SitegroupIds []string `json:"sitegroup_ids,omitempty"`
	// list of vendors to match
	Vendor []string `json:"vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NacRuleMatching NacRuleMatching

// NewNacRuleMatching instantiates a new NacRuleMatching object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNacRuleMatching() *NacRuleMatching {
	this := NacRuleMatching{}
	return &this
}

// NewNacRuleMatchingWithDefaults instantiates a new NacRuleMatching object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNacRuleMatchingWithDefaults() *NacRuleMatching {
	this := NacRuleMatching{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *NacRuleMatching) GetAuthType() NacRuleMatchingAuthType {
	if o == nil || IsNil(o.AuthType) {
		var ret NacRuleMatchingAuthType
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacRuleMatching) GetAuthTypeOk() (*NacRuleMatchingAuthType, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *NacRuleMatching) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given NacRuleMatchingAuthType and assigns it to the AuthType field.
func (o *NacRuleMatching) SetAuthType(v NacRuleMatchingAuthType) {
	o.AuthType = &v
}

// GetNactags returns the Nactags field value if set, zero value otherwise.
func (o *NacRuleMatching) GetNactags() []string {
	if o == nil || IsNil(o.Nactags) {
		var ret []string
		return ret
	}
	return o.Nactags
}

// GetNactagsOk returns a tuple with the Nactags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacRuleMatching) GetNactagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Nactags) {
		return nil, false
	}
	return o.Nactags, true
}

// HasNactags returns a boolean if a field has been set.
func (o *NacRuleMatching) HasNactags() bool {
	if o != nil && !IsNil(o.Nactags) {
		return true
	}

	return false
}

// SetNactags gets a reference to the given []string and assigns it to the Nactags field.
func (o *NacRuleMatching) SetNactags(v []string) {
	o.Nactags = v
}

// GetPortTypes returns the PortTypes field value if set, zero value otherwise.
func (o *NacRuleMatching) GetPortTypes() []NacRuleMatchingPortType {
	if o == nil || IsNil(o.PortTypes) {
		var ret []NacRuleMatchingPortType
		return ret
	}
	return o.PortTypes
}

// GetPortTypesOk returns a tuple with the PortTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacRuleMatching) GetPortTypesOk() ([]NacRuleMatchingPortType, bool) {
	if o == nil || IsNil(o.PortTypes) {
		return nil, false
	}
	return o.PortTypes, true
}

// HasPortTypes returns a boolean if a field has been set.
func (o *NacRuleMatching) HasPortTypes() bool {
	if o != nil && !IsNil(o.PortTypes) {
		return true
	}

	return false
}

// SetPortTypes gets a reference to the given []NacRuleMatchingPortType and assigns it to the PortTypes field.
func (o *NacRuleMatching) SetPortTypes(v []NacRuleMatchingPortType) {
	o.PortTypes = v
}

// GetSiteIds returns the SiteIds field value if set, zero value otherwise.
func (o *NacRuleMatching) GetSiteIds() []string {
	if o == nil || IsNil(o.SiteIds) {
		var ret []string
		return ret
	}
	return o.SiteIds
}

// GetSiteIdsOk returns a tuple with the SiteIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacRuleMatching) GetSiteIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SiteIds) {
		return nil, false
	}
	return o.SiteIds, true
}

// HasSiteIds returns a boolean if a field has been set.
func (o *NacRuleMatching) HasSiteIds() bool {
	if o != nil && !IsNil(o.SiteIds) {
		return true
	}

	return false
}

// SetSiteIds gets a reference to the given []string and assigns it to the SiteIds field.
func (o *NacRuleMatching) SetSiteIds(v []string) {
	o.SiteIds = v
}

// GetSitegroupIds returns the SitegroupIds field value if set, zero value otherwise.
func (o *NacRuleMatching) GetSitegroupIds() []string {
	if o == nil || IsNil(o.SitegroupIds) {
		var ret []string
		return ret
	}
	return o.SitegroupIds
}

// GetSitegroupIdsOk returns a tuple with the SitegroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacRuleMatching) GetSitegroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SitegroupIds) {
		return nil, false
	}
	return o.SitegroupIds, true
}

// HasSitegroupIds returns a boolean if a field has been set.
func (o *NacRuleMatching) HasSitegroupIds() bool {
	if o != nil && !IsNil(o.SitegroupIds) {
		return true
	}

	return false
}

// SetSitegroupIds gets a reference to the given []string and assigns it to the SitegroupIds field.
func (o *NacRuleMatching) SetSitegroupIds(v []string) {
	o.SitegroupIds = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *NacRuleMatching) GetVendor() []string {
	if o == nil || IsNil(o.Vendor) {
		var ret []string
		return ret
	}
	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacRuleMatching) GetVendorOk() ([]string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *NacRuleMatching) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given []string and assigns it to the Vendor field.
func (o *NacRuleMatching) SetVendor(v []string) {
	o.Vendor = v
}

func (o NacRuleMatching) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NacRuleMatching) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.Nactags) {
		toSerialize["nactags"] = o.Nactags
	}
	if !IsNil(o.PortTypes) {
		toSerialize["port_types"] = o.PortTypes
	}
	if !IsNil(o.SiteIds) {
		toSerialize["site_ids"] = o.SiteIds
	}
	if !IsNil(o.SitegroupIds) {
		toSerialize["sitegroup_ids"] = o.SitegroupIds
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NacRuleMatching) UnmarshalJSON(data []byte) (err error) {
	varNacRuleMatching := _NacRuleMatching{}

	err = json.Unmarshal(data, &varNacRuleMatching)

	if err != nil {
		return err
	}

	*o = NacRuleMatching(varNacRuleMatching)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "nactags")
		delete(additionalProperties, "port_types")
		delete(additionalProperties, "site_ids")
		delete(additionalProperties, "sitegroup_ids")
		delete(additionalProperties, "vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNacRuleMatching struct {
	value *NacRuleMatching
	isSet bool
}

func (v NullableNacRuleMatching) Get() *NacRuleMatching {
	return v.value
}

func (v *NullableNacRuleMatching) Set(val *NacRuleMatching) {
	v.value = val
	v.isSet = true
}

func (v NullableNacRuleMatching) IsSet() bool {
	return v.isSet
}

func (v *NullableNacRuleMatching) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNacRuleMatching(val *NacRuleMatching) *NullableNacRuleMatching {
	return &NullableNacRuleMatching{value: val, isSet: true}
}

func (v NullableNacRuleMatching) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNacRuleMatching) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



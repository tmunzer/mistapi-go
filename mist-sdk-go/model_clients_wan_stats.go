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

// checks if the ClientsWanStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientsWanStats{}

// ClientsWanStats struct for ClientsWanStats
type ClientsWanStats struct {
	When *string `json:"When,omitempty"`
	Hostname []string `json:"hostname,omitempty"`
	Ip []string `json:"ip,omitempty"`
	LastHostname *string `json:"last_hostname,omitempty"`
	LastIp *string `json:"last_ip,omitempty"`
	Mfg *string `json:"mfg,omitempty"`
	Network *string `json:"network,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	Wcid *string `json:"wcid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientsWanStats ClientsWanStats

// NewClientsWanStats instantiates a new ClientsWanStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientsWanStats() *ClientsWanStats {
	this := ClientsWanStats{}
	return &this
}

// NewClientsWanStatsWithDefaults instantiates a new ClientsWanStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientsWanStatsWithDefaults() *ClientsWanStats {
	this := ClientsWanStats{}
	return &this
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *ClientsWanStats) GetWhen() string {
	if o == nil || IsNil(o.When) {
		var ret string
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetWhenOk() (*string, bool) {
	if o == nil || IsNil(o.When) {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *ClientsWanStats) HasWhen() bool {
	if o != nil && !IsNil(o.When) {
		return true
	}

	return false
}

// SetWhen gets a reference to the given string and assigns it to the When field.
func (o *ClientsWanStats) SetWhen(v string) {
	o.When = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ClientsWanStats) GetHostname() []string {
	if o == nil || IsNil(o.Hostname) {
		var ret []string
		return ret
	}
	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetHostnameOk() ([]string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ClientsWanStats) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given []string and assigns it to the Hostname field.
func (o *ClientsWanStats) SetHostname(v []string) {
	o.Hostname = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ClientsWanStats) GetIp() []string {
	if o == nil || IsNil(o.Ip) {
		var ret []string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetIpOk() ([]string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ClientsWanStats) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given []string and assigns it to the Ip field.
func (o *ClientsWanStats) SetIp(v []string) {
	o.Ip = v
}

// GetLastHostname returns the LastHostname field value if set, zero value otherwise.
func (o *ClientsWanStats) GetLastHostname() string {
	if o == nil || IsNil(o.LastHostname) {
		var ret string
		return ret
	}
	return *o.LastHostname
}

// GetLastHostnameOk returns a tuple with the LastHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetLastHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.LastHostname) {
		return nil, false
	}
	return o.LastHostname, true
}

// HasLastHostname returns a boolean if a field has been set.
func (o *ClientsWanStats) HasLastHostname() bool {
	if o != nil && !IsNil(o.LastHostname) {
		return true
	}

	return false
}

// SetLastHostname gets a reference to the given string and assigns it to the LastHostname field.
func (o *ClientsWanStats) SetLastHostname(v string) {
	o.LastHostname = &v
}

// GetLastIp returns the LastIp field value if set, zero value otherwise.
func (o *ClientsWanStats) GetLastIp() string {
	if o == nil || IsNil(o.LastIp) {
		var ret string
		return ret
	}
	return *o.LastIp
}

// GetLastIpOk returns a tuple with the LastIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetLastIpOk() (*string, bool) {
	if o == nil || IsNil(o.LastIp) {
		return nil, false
	}
	return o.LastIp, true
}

// HasLastIp returns a boolean if a field has been set.
func (o *ClientsWanStats) HasLastIp() bool {
	if o != nil && !IsNil(o.LastIp) {
		return true
	}

	return false
}

// SetLastIp gets a reference to the given string and assigns it to the LastIp field.
func (o *ClientsWanStats) SetLastIp(v string) {
	o.LastIp = &v
}

// GetMfg returns the Mfg field value if set, zero value otherwise.
func (o *ClientsWanStats) GetMfg() string {
	if o == nil || IsNil(o.Mfg) {
		var ret string
		return ret
	}
	return *o.Mfg
}

// GetMfgOk returns a tuple with the Mfg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetMfgOk() (*string, bool) {
	if o == nil || IsNil(o.Mfg) {
		return nil, false
	}
	return o.Mfg, true
}

// HasMfg returns a boolean if a field has been set.
func (o *ClientsWanStats) HasMfg() bool {
	if o != nil && !IsNil(o.Mfg) {
		return true
	}

	return false
}

// SetMfg gets a reference to the given string and assigns it to the Mfg field.
func (o *ClientsWanStats) SetMfg(v string) {
	o.Mfg = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ClientsWanStats) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ClientsWanStats) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *ClientsWanStats) SetNetwork(v string) {
	o.Network = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ClientsWanStats) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ClientsWanStats) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ClientsWanStats) SetOrgId(v string) {
	o.OrgId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *ClientsWanStats) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *ClientsWanStats) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *ClientsWanStats) SetSiteId(v string) {
	o.SiteId = &v
}

// GetWcid returns the Wcid field value if set, zero value otherwise.
func (o *ClientsWanStats) GetWcid() string {
	if o == nil || IsNil(o.Wcid) {
		var ret string
		return ret
	}
	return *o.Wcid
}

// GetWcidOk returns a tuple with the Wcid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientsWanStats) GetWcidOk() (*string, bool) {
	if o == nil || IsNil(o.Wcid) {
		return nil, false
	}
	return o.Wcid, true
}

// HasWcid returns a boolean if a field has been set.
func (o *ClientsWanStats) HasWcid() bool {
	if o != nil && !IsNil(o.Wcid) {
		return true
	}

	return false
}

// SetWcid gets a reference to the given string and assigns it to the Wcid field.
func (o *ClientsWanStats) SetWcid(v string) {
	o.Wcid = &v
}

func (o ClientsWanStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientsWanStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.When) {
		toSerialize["When"] = o.When
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.LastHostname) {
		toSerialize["last_hostname"] = o.LastHostname
	}
	if !IsNil(o.LastIp) {
		toSerialize["last_ip"] = o.LastIp
	}
	if !IsNil(o.Mfg) {
		toSerialize["mfg"] = o.Mfg
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Wcid) {
		toSerialize["wcid"] = o.Wcid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientsWanStats) UnmarshalJSON(data []byte) (err error) {
	varClientsWanStats := _ClientsWanStats{}

	err = json.Unmarshal(data, &varClientsWanStats)

	if err != nil {
		return err
	}

	*o = ClientsWanStats(varClientsWanStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "When")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "last_hostname")
		delete(additionalProperties, "last_ip")
		delete(additionalProperties, "mfg")
		delete(additionalProperties, "network")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "wcid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientsWanStats struct {
	value *ClientsWanStats
	isSet bool
}

func (v NullableClientsWanStats) Get() *ClientsWanStats {
	return v.value
}

func (v *NullableClientsWanStats) Set(val *ClientsWanStats) {
	v.value = val
	v.isSet = true
}

func (v NullableClientsWanStats) IsSet() bool {
	return v.isSet
}

func (v *NullableClientsWanStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientsWanStats(val *ClientsWanStats) *NullableClientsWanStats {
	return &NullableClientsWanStats{value: val, isSet: true}
}

func (v NullableClientsWanStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientsWanStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



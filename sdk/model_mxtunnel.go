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

// checks if the Mxtunnel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mxtunnel{}

// Mxtunnel MxTunnel
type Mxtunnel struct {
	// list of anchor mxtunnels used for forming edge to edge tunnels
	AnchorMxtunnelIds []string `json:"anchor_mxtunnel_ids,omitempty"`
	AutoPreemption *AutoPreemption `json:"auto_preemption,omitempty"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	// in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by `hello_retries`.
	HelloInterval NullableInt32 `json:"hello_interval,omitempty"`
	HelloRetries NullableInt32 `json:"hello_retries,omitempty"`
	Id *string `json:"id,omitempty"`
	Ipsec *MxtunnelIpsec `json:"ipsec,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	// 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU
	Mtu *int32 `json:"mtu,omitempty"`
	// list of mxclusters to deploy this tunnel to
	MxclusterIds []string `json:"mxcluster_ids,omitempty"`
	Name NullableString `json:"name,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	Protocol *MxtunnelProtocol `json:"protocol,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	// list of vlan_ids that will be used
	VlanIds []int32 `json:"vlan_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Mxtunnel Mxtunnel

// NewMxtunnel instantiates a new Mxtunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxtunnel() *Mxtunnel {
	this := Mxtunnel{}
	var helloInterval int32 = 60
	this.HelloInterval = *NewNullableInt32(&helloInterval)
	var helloRetries int32 = 7
	this.HelloRetries = *NewNullableInt32(&helloRetries)
	var mtu int32 = 0
	this.Mtu = &mtu
	var protocol MxtunnelProtocol = MXTUNNELPROTOCOL_UDP
	this.Protocol = &protocol
	return &this
}

// NewMxtunnelWithDefaults instantiates a new Mxtunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxtunnelWithDefaults() *Mxtunnel {
	this := Mxtunnel{}
	var helloInterval int32 = 60
	this.HelloInterval = *NewNullableInt32(&helloInterval)
	var helloRetries int32 = 7
	this.HelloRetries = *NewNullableInt32(&helloRetries)
	var mtu int32 = 0
	this.Mtu = &mtu
	var protocol MxtunnelProtocol = MXTUNNELPROTOCOL_UDP
	this.Protocol = &protocol
	return &this
}

// GetAnchorMxtunnelIds returns the AnchorMxtunnelIds field value if set, zero value otherwise.
func (o *Mxtunnel) GetAnchorMxtunnelIds() []string {
	if o == nil || IsNil(o.AnchorMxtunnelIds) {
		var ret []string
		return ret
	}
	return o.AnchorMxtunnelIds
}

// GetAnchorMxtunnelIdsOk returns a tuple with the AnchorMxtunnelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetAnchorMxtunnelIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AnchorMxtunnelIds) {
		return nil, false
	}
	return o.AnchorMxtunnelIds, true
}

// HasAnchorMxtunnelIds returns a boolean if a field has been set.
func (o *Mxtunnel) HasAnchorMxtunnelIds() bool {
	if o != nil && !IsNil(o.AnchorMxtunnelIds) {
		return true
	}

	return false
}

// SetAnchorMxtunnelIds gets a reference to the given []string and assigns it to the AnchorMxtunnelIds field.
func (o *Mxtunnel) SetAnchorMxtunnelIds(v []string) {
	o.AnchorMxtunnelIds = v
}

// GetAutoPreemption returns the AutoPreemption field value if set, zero value otherwise.
func (o *Mxtunnel) GetAutoPreemption() AutoPreemption {
	if o == nil || IsNil(o.AutoPreemption) {
		var ret AutoPreemption
		return ret
	}
	return *o.AutoPreemption
}

// GetAutoPreemptionOk returns a tuple with the AutoPreemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetAutoPreemptionOk() (*AutoPreemption, bool) {
	if o == nil || IsNil(o.AutoPreemption) {
		return nil, false
	}
	return o.AutoPreemption, true
}

// HasAutoPreemption returns a boolean if a field has been set.
func (o *Mxtunnel) HasAutoPreemption() bool {
	if o != nil && !IsNil(o.AutoPreemption) {
		return true
	}

	return false
}

// SetAutoPreemption gets a reference to the given AutoPreemption and assigns it to the AutoPreemption field.
func (o *Mxtunnel) SetAutoPreemption(v AutoPreemption) {
	o.AutoPreemption = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Mxtunnel) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Mxtunnel) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Mxtunnel) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *Mxtunnel) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *Mxtunnel) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *Mxtunnel) SetForSite(v bool) {
	o.ForSite = &v
}

// GetHelloInterval returns the HelloInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mxtunnel) GetHelloInterval() int32 {
	if o == nil || IsNil(o.HelloInterval.Get()) {
		var ret int32
		return ret
	}
	return *o.HelloInterval.Get()
}

// GetHelloIntervalOk returns a tuple with the HelloInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mxtunnel) GetHelloIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.HelloInterval.Get(), o.HelloInterval.IsSet()
}

// HasHelloInterval returns a boolean if a field has been set.
func (o *Mxtunnel) HasHelloInterval() bool {
	if o != nil && o.HelloInterval.IsSet() {
		return true
	}

	return false
}

// SetHelloInterval gets a reference to the given NullableInt32 and assigns it to the HelloInterval field.
func (o *Mxtunnel) SetHelloInterval(v int32) {
	o.HelloInterval.Set(&v)
}
// SetHelloIntervalNil sets the value for HelloInterval to be an explicit nil
func (o *Mxtunnel) SetHelloIntervalNil() {
	o.HelloInterval.Set(nil)
}

// UnsetHelloInterval ensures that no value is present for HelloInterval, not even an explicit nil
func (o *Mxtunnel) UnsetHelloInterval() {
	o.HelloInterval.Unset()
}

// GetHelloRetries returns the HelloRetries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mxtunnel) GetHelloRetries() int32 {
	if o == nil || IsNil(o.HelloRetries.Get()) {
		var ret int32
		return ret
	}
	return *o.HelloRetries.Get()
}

// GetHelloRetriesOk returns a tuple with the HelloRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mxtunnel) GetHelloRetriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.HelloRetries.Get(), o.HelloRetries.IsSet()
}

// HasHelloRetries returns a boolean if a field has been set.
func (o *Mxtunnel) HasHelloRetries() bool {
	if o != nil && o.HelloRetries.IsSet() {
		return true
	}

	return false
}

// SetHelloRetries gets a reference to the given NullableInt32 and assigns it to the HelloRetries field.
func (o *Mxtunnel) SetHelloRetries(v int32) {
	o.HelloRetries.Set(&v)
}
// SetHelloRetriesNil sets the value for HelloRetries to be an explicit nil
func (o *Mxtunnel) SetHelloRetriesNil() {
	o.HelloRetries.Set(nil)
}

// UnsetHelloRetries ensures that no value is present for HelloRetries, not even an explicit nil
func (o *Mxtunnel) UnsetHelloRetries() {
	o.HelloRetries.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Mxtunnel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Mxtunnel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Mxtunnel) SetId(v string) {
	o.Id = &v
}

// GetIpsec returns the Ipsec field value if set, zero value otherwise.
func (o *Mxtunnel) GetIpsec() MxtunnelIpsec {
	if o == nil || IsNil(o.Ipsec) {
		var ret MxtunnelIpsec
		return ret
	}
	return *o.Ipsec
}

// GetIpsecOk returns a tuple with the Ipsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetIpsecOk() (*MxtunnelIpsec, bool) {
	if o == nil || IsNil(o.Ipsec) {
		return nil, false
	}
	return o.Ipsec, true
}

// HasIpsec returns a boolean if a field has been set.
func (o *Mxtunnel) HasIpsec() bool {
	if o != nil && !IsNil(o.Ipsec) {
		return true
	}

	return false
}

// SetIpsec gets a reference to the given MxtunnelIpsec and assigns it to the Ipsec field.
func (o *Mxtunnel) SetIpsec(v MxtunnelIpsec) {
	o.Ipsec = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Mxtunnel) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Mxtunnel) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *Mxtunnel) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *Mxtunnel) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu) {
		var ret int32
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetMtuOk() (*int32, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *Mxtunnel) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int32 and assigns it to the Mtu field.
func (o *Mxtunnel) SetMtu(v int32) {
	o.Mtu = &v
}

// GetMxclusterIds returns the MxclusterIds field value if set, zero value otherwise.
func (o *Mxtunnel) GetMxclusterIds() []string {
	if o == nil || IsNil(o.MxclusterIds) {
		var ret []string
		return ret
	}
	return o.MxclusterIds
}

// GetMxclusterIdsOk returns a tuple with the MxclusterIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetMxclusterIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MxclusterIds) {
		return nil, false
	}
	return o.MxclusterIds, true
}

// HasMxclusterIds returns a boolean if a field has been set.
func (o *Mxtunnel) HasMxclusterIds() bool {
	if o != nil && !IsNil(o.MxclusterIds) {
		return true
	}

	return false
}

// SetMxclusterIds gets a reference to the given []string and assigns it to the MxclusterIds field.
func (o *Mxtunnel) SetMxclusterIds(v []string) {
	o.MxclusterIds = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mxtunnel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mxtunnel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Mxtunnel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Mxtunnel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Mxtunnel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Mxtunnel) UnsetName() {
	o.Name.Unset()
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Mxtunnel) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Mxtunnel) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Mxtunnel) SetOrgId(v string) {
	o.OrgId = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Mxtunnel) GetProtocol() MxtunnelProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret MxtunnelProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetProtocolOk() (*MxtunnelProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Mxtunnel) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given MxtunnelProtocol and assigns it to the Protocol field.
func (o *Mxtunnel) SetProtocol(v MxtunnelProtocol) {
	o.Protocol = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *Mxtunnel) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *Mxtunnel) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *Mxtunnel) SetSiteId(v string) {
	o.SiteId = &v
}

// GetVlanIds returns the VlanIds field value if set, zero value otherwise.
func (o *Mxtunnel) GetVlanIds() []int32 {
	if o == nil || IsNil(o.VlanIds) {
		var ret []int32
		return ret
	}
	return o.VlanIds
}

// GetVlanIdsOk returns a tuple with the VlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxtunnel) GetVlanIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.VlanIds) {
		return nil, false
	}
	return o.VlanIds, true
}

// HasVlanIds returns a boolean if a field has been set.
func (o *Mxtunnel) HasVlanIds() bool {
	if o != nil && !IsNil(o.VlanIds) {
		return true
	}

	return false
}

// SetVlanIds gets a reference to the given []int32 and assigns it to the VlanIds field.
func (o *Mxtunnel) SetVlanIds(v []int32) {
	o.VlanIds = v
}

func (o Mxtunnel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mxtunnel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnchorMxtunnelIds) {
		toSerialize["anchor_mxtunnel_ids"] = o.AnchorMxtunnelIds
	}
	if !IsNil(o.AutoPreemption) {
		toSerialize["auto_preemption"] = o.AutoPreemption
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if o.HelloInterval.IsSet() {
		toSerialize["hello_interval"] = o.HelloInterval.Get()
	}
	if o.HelloRetries.IsSet() {
		toSerialize["hello_retries"] = o.HelloRetries.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ipsec) {
		toSerialize["ipsec"] = o.Ipsec
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	if !IsNil(o.MxclusterIds) {
		toSerialize["mxcluster_ids"] = o.MxclusterIds
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.VlanIds) {
		toSerialize["vlan_ids"] = o.VlanIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Mxtunnel) UnmarshalJSON(data []byte) (err error) {
	varMxtunnel := _Mxtunnel{}

	err = json.Unmarshal(data, &varMxtunnel)

	if err != nil {
		return err
	}

	*o = Mxtunnel(varMxtunnel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "anchor_mxtunnel_ids")
		delete(additionalProperties, "auto_preemption")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "hello_interval")
		delete(additionalProperties, "hello_retries")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ipsec")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "mtu")
		delete(additionalProperties, "mxcluster_ids")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "vlan_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxtunnel struct {
	value *Mxtunnel
	isSet bool
}

func (v NullableMxtunnel) Get() *Mxtunnel {
	return v.value
}

func (v *NullableMxtunnel) Set(val *Mxtunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableMxtunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableMxtunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxtunnel(val *Mxtunnel) *NullableMxtunnel {
	return &NullableMxtunnel{value: val, isSet: true}
}

func (v NullableMxtunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxtunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



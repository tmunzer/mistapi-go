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

// checks if the Mxedge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mxedge{}

// Mxedge MxEdge
type Mxedge struct {
	CreatedTime *float32 `json:"created_time,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	Id *string `json:"id,omitempty"`
	Magic *string `json:"magic,omitempty"`
	Model string `json:"model"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	MxagentRegistered *bool `json:"mxagent_registered,omitempty"`
	// MxCluster this MxEdge belongs to
	MxclusterId *string `json:"mxcluster_id,omitempty"`
	MxedgeMgmt *MxedgeMgmt `json:"mxedge_mgmt,omitempty"`
	Name string `json:"name"`
	Note *string `json:"note,omitempty"`
	NtpServers []string `json:"ntp_servers,omitempty"`
	OobIpConfig *MxedgeOobIpConfig `json:"oob_ip_config,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	Proxy *Proxy `json:"proxy,omitempty"`
	// list of services to run, tunterm only for now
	Services []string `json:"services,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	TuntermDhcpdConfig *MxedgeTuntermDhcpdConfigs `json:"tunterm_dhcpd_config,omitempty"`
	// Property key is a CIDR
	TuntermExtraRoutes *map[string]MxedgeTuntermExtraRoute `json:"tunterm_extra_routes,omitempty"`
	TuntermIgmpSnoopingConfig *MxedgeTuntermIgmpSnoopingConfig `json:"tunterm_igmp_snooping_config,omitempty"`
	TuntermIpConfig *MxedgeTuntermIpConfig `json:"tunterm_ip_config,omitempty"`
	TuntermMonitoring [][]TuntermMonitoringItem `json:"tunterm_monitoring,omitempty"`
	TuntermMulticastConfig *MxedgeTuntermMulticastConfig `json:"tunterm_multicast_config,omitempty"`
	// ip configs by VLAN ID. Property key is the VLAN ID
	TuntermOtherIpConfigs *map[string]MxedgeTuntermOtherIpConfig `json:"tunterm_other_ip_configs,omitempty"`
	TuntermPortConfig *TuntermPortConfig `json:"tunterm_port_config,omitempty"`
	TuntermRegistered *bool `json:"tunterm_registered,omitempty"`
	TuntermSwitchConfig *MxedgeTuntermSwitchConfigs `json:"tunterm_switch_config,omitempty"`
	Versions *MxedgeVersions `json:"versions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Mxedge Mxedge

// NewMxedge instantiates a new Mxedge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedge(model string, name string) *Mxedge {
	this := Mxedge{}
	this.Model = model
	this.Name = name
	return &this
}

// NewMxedgeWithDefaults instantiates a new Mxedge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeWithDefaults() *Mxedge {
	this := Mxedge{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Mxedge) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Mxedge) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Mxedge) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *Mxedge) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *Mxedge) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *Mxedge) SetForSite(v bool) {
	o.ForSite = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Mxedge) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Mxedge) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Mxedge) SetId(v string) {
	o.Id = &v
}

// GetMagic returns the Magic field value if set, zero value otherwise.
func (o *Mxedge) GetMagic() string {
	if o == nil || IsNil(o.Magic) {
		var ret string
		return ret
	}
	return *o.Magic
}

// GetMagicOk returns a tuple with the Magic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetMagicOk() (*string, bool) {
	if o == nil || IsNil(o.Magic) {
		return nil, false
	}
	return o.Magic, true
}

// HasMagic returns a boolean if a field has been set.
func (o *Mxedge) HasMagic() bool {
	if o != nil && !IsNil(o.Magic) {
		return true
	}

	return false
}

// SetMagic gets a reference to the given string and assigns it to the Magic field.
func (o *Mxedge) SetMagic(v string) {
	o.Magic = &v
}

// GetModel returns the Model field value
func (o *Mxedge) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *Mxedge) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *Mxedge) SetModel(v string) {
	o.Model = v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Mxedge) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Mxedge) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *Mxedge) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetMxagentRegistered returns the MxagentRegistered field value if set, zero value otherwise.
func (o *Mxedge) GetMxagentRegistered() bool {
	if o == nil || IsNil(o.MxagentRegistered) {
		var ret bool
		return ret
	}
	return *o.MxagentRegistered
}

// GetMxagentRegisteredOk returns a tuple with the MxagentRegistered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetMxagentRegisteredOk() (*bool, bool) {
	if o == nil || IsNil(o.MxagentRegistered) {
		return nil, false
	}
	return o.MxagentRegistered, true
}

// HasMxagentRegistered returns a boolean if a field has been set.
func (o *Mxedge) HasMxagentRegistered() bool {
	if o != nil && !IsNil(o.MxagentRegistered) {
		return true
	}

	return false
}

// SetMxagentRegistered gets a reference to the given bool and assigns it to the MxagentRegistered field.
func (o *Mxedge) SetMxagentRegistered(v bool) {
	o.MxagentRegistered = &v
}

// GetMxclusterId returns the MxclusterId field value if set, zero value otherwise.
func (o *Mxedge) GetMxclusterId() string {
	if o == nil || IsNil(o.MxclusterId) {
		var ret string
		return ret
	}
	return *o.MxclusterId
}

// GetMxclusterIdOk returns a tuple with the MxclusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetMxclusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.MxclusterId) {
		return nil, false
	}
	return o.MxclusterId, true
}

// HasMxclusterId returns a boolean if a field has been set.
func (o *Mxedge) HasMxclusterId() bool {
	if o != nil && !IsNil(o.MxclusterId) {
		return true
	}

	return false
}

// SetMxclusterId gets a reference to the given string and assigns it to the MxclusterId field.
func (o *Mxedge) SetMxclusterId(v string) {
	o.MxclusterId = &v
}

// GetMxedgeMgmt returns the MxedgeMgmt field value if set, zero value otherwise.
func (o *Mxedge) GetMxedgeMgmt() MxedgeMgmt {
	if o == nil || IsNil(o.MxedgeMgmt) {
		var ret MxedgeMgmt
		return ret
	}
	return *o.MxedgeMgmt
}

// GetMxedgeMgmtOk returns a tuple with the MxedgeMgmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetMxedgeMgmtOk() (*MxedgeMgmt, bool) {
	if o == nil || IsNil(o.MxedgeMgmt) {
		return nil, false
	}
	return o.MxedgeMgmt, true
}

// HasMxedgeMgmt returns a boolean if a field has been set.
func (o *Mxedge) HasMxedgeMgmt() bool {
	if o != nil && !IsNil(o.MxedgeMgmt) {
		return true
	}

	return false
}

// SetMxedgeMgmt gets a reference to the given MxedgeMgmt and assigns it to the MxedgeMgmt field.
func (o *Mxedge) SetMxedgeMgmt(v MxedgeMgmt) {
	o.MxedgeMgmt = &v
}

// GetName returns the Name field value
func (o *Mxedge) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Mxedge) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Mxedge) SetName(v string) {
	o.Name = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *Mxedge) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *Mxedge) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *Mxedge) SetNote(v string) {
	o.Note = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *Mxedge) GetNtpServers() []string {
	if o == nil || IsNil(o.NtpServers) {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetNtpServersOk() ([]string, bool) {
	if o == nil || IsNil(o.NtpServers) {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *Mxedge) HasNtpServers() bool {
	if o != nil && !IsNil(o.NtpServers) {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *Mxedge) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetOobIpConfig returns the OobIpConfig field value if set, zero value otherwise.
func (o *Mxedge) GetOobIpConfig() MxedgeOobIpConfig {
	if o == nil || IsNil(o.OobIpConfig) {
		var ret MxedgeOobIpConfig
		return ret
	}
	return *o.OobIpConfig
}

// GetOobIpConfigOk returns a tuple with the OobIpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetOobIpConfigOk() (*MxedgeOobIpConfig, bool) {
	if o == nil || IsNil(o.OobIpConfig) {
		return nil, false
	}
	return o.OobIpConfig, true
}

// HasOobIpConfig returns a boolean if a field has been set.
func (o *Mxedge) HasOobIpConfig() bool {
	if o != nil && !IsNil(o.OobIpConfig) {
		return true
	}

	return false
}

// SetOobIpConfig gets a reference to the given MxedgeOobIpConfig and assigns it to the OobIpConfig field.
func (o *Mxedge) SetOobIpConfig(v MxedgeOobIpConfig) {
	o.OobIpConfig = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Mxedge) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Mxedge) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Mxedge) SetOrgId(v string) {
	o.OrgId = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *Mxedge) GetProxy() Proxy {
	if o == nil || IsNil(o.Proxy) {
		var ret Proxy
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetProxyOk() (*Proxy, bool) {
	if o == nil || IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *Mxedge) HasProxy() bool {
	if o != nil && !IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given Proxy and assigns it to the Proxy field.
func (o *Mxedge) SetProxy(v Proxy) {
	o.Proxy = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *Mxedge) GetServices() []string {
	if o == nil || IsNil(o.Services) {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *Mxedge) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *Mxedge) SetServices(v []string) {
	o.Services = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *Mxedge) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *Mxedge) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *Mxedge) SetSiteId(v string) {
	o.SiteId = &v
}

// GetTuntermDhcpdConfig returns the TuntermDhcpdConfig field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermDhcpdConfig() MxedgeTuntermDhcpdConfigs {
	if o == nil || IsNil(o.TuntermDhcpdConfig) {
		var ret MxedgeTuntermDhcpdConfigs
		return ret
	}
	return *o.TuntermDhcpdConfig
}

// GetTuntermDhcpdConfigOk returns a tuple with the TuntermDhcpdConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermDhcpdConfigOk() (*MxedgeTuntermDhcpdConfigs, bool) {
	if o == nil || IsNil(o.TuntermDhcpdConfig) {
		return nil, false
	}
	return o.TuntermDhcpdConfig, true
}

// HasTuntermDhcpdConfig returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermDhcpdConfig() bool {
	if o != nil && !IsNil(o.TuntermDhcpdConfig) {
		return true
	}

	return false
}

// SetTuntermDhcpdConfig gets a reference to the given MxedgeTuntermDhcpdConfigs and assigns it to the TuntermDhcpdConfig field.
func (o *Mxedge) SetTuntermDhcpdConfig(v MxedgeTuntermDhcpdConfigs) {
	o.TuntermDhcpdConfig = &v
}

// GetTuntermExtraRoutes returns the TuntermExtraRoutes field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermExtraRoutes() map[string]MxedgeTuntermExtraRoute {
	if o == nil || IsNil(o.TuntermExtraRoutes) {
		var ret map[string]MxedgeTuntermExtraRoute
		return ret
	}
	return *o.TuntermExtraRoutes
}

// GetTuntermExtraRoutesOk returns a tuple with the TuntermExtraRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermExtraRoutesOk() (*map[string]MxedgeTuntermExtraRoute, bool) {
	if o == nil || IsNil(o.TuntermExtraRoutes) {
		return nil, false
	}
	return o.TuntermExtraRoutes, true
}

// HasTuntermExtraRoutes returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermExtraRoutes() bool {
	if o != nil && !IsNil(o.TuntermExtraRoutes) {
		return true
	}

	return false
}

// SetTuntermExtraRoutes gets a reference to the given map[string]MxedgeTuntermExtraRoute and assigns it to the TuntermExtraRoutes field.
func (o *Mxedge) SetTuntermExtraRoutes(v map[string]MxedgeTuntermExtraRoute) {
	o.TuntermExtraRoutes = &v
}

// GetTuntermIgmpSnoopingConfig returns the TuntermIgmpSnoopingConfig field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermIgmpSnoopingConfig() MxedgeTuntermIgmpSnoopingConfig {
	if o == nil || IsNil(o.TuntermIgmpSnoopingConfig) {
		var ret MxedgeTuntermIgmpSnoopingConfig
		return ret
	}
	return *o.TuntermIgmpSnoopingConfig
}

// GetTuntermIgmpSnoopingConfigOk returns a tuple with the TuntermIgmpSnoopingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermIgmpSnoopingConfigOk() (*MxedgeTuntermIgmpSnoopingConfig, bool) {
	if o == nil || IsNil(o.TuntermIgmpSnoopingConfig) {
		return nil, false
	}
	return o.TuntermIgmpSnoopingConfig, true
}

// HasTuntermIgmpSnoopingConfig returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermIgmpSnoopingConfig() bool {
	if o != nil && !IsNil(o.TuntermIgmpSnoopingConfig) {
		return true
	}

	return false
}

// SetTuntermIgmpSnoopingConfig gets a reference to the given MxedgeTuntermIgmpSnoopingConfig and assigns it to the TuntermIgmpSnoopingConfig field.
func (o *Mxedge) SetTuntermIgmpSnoopingConfig(v MxedgeTuntermIgmpSnoopingConfig) {
	o.TuntermIgmpSnoopingConfig = &v
}

// GetTuntermIpConfig returns the TuntermIpConfig field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermIpConfig() MxedgeTuntermIpConfig {
	if o == nil || IsNil(o.TuntermIpConfig) {
		var ret MxedgeTuntermIpConfig
		return ret
	}
	return *o.TuntermIpConfig
}

// GetTuntermIpConfigOk returns a tuple with the TuntermIpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermIpConfigOk() (*MxedgeTuntermIpConfig, bool) {
	if o == nil || IsNil(o.TuntermIpConfig) {
		return nil, false
	}
	return o.TuntermIpConfig, true
}

// HasTuntermIpConfig returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermIpConfig() bool {
	if o != nil && !IsNil(o.TuntermIpConfig) {
		return true
	}

	return false
}

// SetTuntermIpConfig gets a reference to the given MxedgeTuntermIpConfig and assigns it to the TuntermIpConfig field.
func (o *Mxedge) SetTuntermIpConfig(v MxedgeTuntermIpConfig) {
	o.TuntermIpConfig = &v
}

// GetTuntermMonitoring returns the TuntermMonitoring field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermMonitoring() [][]TuntermMonitoringItem {
	if o == nil || IsNil(o.TuntermMonitoring) {
		var ret [][]TuntermMonitoringItem
		return ret
	}
	return o.TuntermMonitoring
}

// GetTuntermMonitoringOk returns a tuple with the TuntermMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermMonitoringOk() ([][]TuntermMonitoringItem, bool) {
	if o == nil || IsNil(o.TuntermMonitoring) {
		return nil, false
	}
	return o.TuntermMonitoring, true
}

// HasTuntermMonitoring returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermMonitoring() bool {
	if o != nil && !IsNil(o.TuntermMonitoring) {
		return true
	}

	return false
}

// SetTuntermMonitoring gets a reference to the given [][]TuntermMonitoringItem and assigns it to the TuntermMonitoring field.
func (o *Mxedge) SetTuntermMonitoring(v [][]TuntermMonitoringItem) {
	o.TuntermMonitoring = v
}

// GetTuntermMulticastConfig returns the TuntermMulticastConfig field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermMulticastConfig() MxedgeTuntermMulticastConfig {
	if o == nil || IsNil(o.TuntermMulticastConfig) {
		var ret MxedgeTuntermMulticastConfig
		return ret
	}
	return *o.TuntermMulticastConfig
}

// GetTuntermMulticastConfigOk returns a tuple with the TuntermMulticastConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermMulticastConfigOk() (*MxedgeTuntermMulticastConfig, bool) {
	if o == nil || IsNil(o.TuntermMulticastConfig) {
		return nil, false
	}
	return o.TuntermMulticastConfig, true
}

// HasTuntermMulticastConfig returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermMulticastConfig() bool {
	if o != nil && !IsNil(o.TuntermMulticastConfig) {
		return true
	}

	return false
}

// SetTuntermMulticastConfig gets a reference to the given MxedgeTuntermMulticastConfig and assigns it to the TuntermMulticastConfig field.
func (o *Mxedge) SetTuntermMulticastConfig(v MxedgeTuntermMulticastConfig) {
	o.TuntermMulticastConfig = &v
}

// GetTuntermOtherIpConfigs returns the TuntermOtherIpConfigs field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermOtherIpConfigs() map[string]MxedgeTuntermOtherIpConfig {
	if o == nil || IsNil(o.TuntermOtherIpConfigs) {
		var ret map[string]MxedgeTuntermOtherIpConfig
		return ret
	}
	return *o.TuntermOtherIpConfigs
}

// GetTuntermOtherIpConfigsOk returns a tuple with the TuntermOtherIpConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermOtherIpConfigsOk() (*map[string]MxedgeTuntermOtherIpConfig, bool) {
	if o == nil || IsNil(o.TuntermOtherIpConfigs) {
		return nil, false
	}
	return o.TuntermOtherIpConfigs, true
}

// HasTuntermOtherIpConfigs returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermOtherIpConfigs() bool {
	if o != nil && !IsNil(o.TuntermOtherIpConfigs) {
		return true
	}

	return false
}

// SetTuntermOtherIpConfigs gets a reference to the given map[string]MxedgeTuntermOtherIpConfig and assigns it to the TuntermOtherIpConfigs field.
func (o *Mxedge) SetTuntermOtherIpConfigs(v map[string]MxedgeTuntermOtherIpConfig) {
	o.TuntermOtherIpConfigs = &v
}

// GetTuntermPortConfig returns the TuntermPortConfig field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermPortConfig() TuntermPortConfig {
	if o == nil || IsNil(o.TuntermPortConfig) {
		var ret TuntermPortConfig
		return ret
	}
	return *o.TuntermPortConfig
}

// GetTuntermPortConfigOk returns a tuple with the TuntermPortConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermPortConfigOk() (*TuntermPortConfig, bool) {
	if o == nil || IsNil(o.TuntermPortConfig) {
		return nil, false
	}
	return o.TuntermPortConfig, true
}

// HasTuntermPortConfig returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermPortConfig() bool {
	if o != nil && !IsNil(o.TuntermPortConfig) {
		return true
	}

	return false
}

// SetTuntermPortConfig gets a reference to the given TuntermPortConfig and assigns it to the TuntermPortConfig field.
func (o *Mxedge) SetTuntermPortConfig(v TuntermPortConfig) {
	o.TuntermPortConfig = &v
}

// GetTuntermRegistered returns the TuntermRegistered field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermRegistered() bool {
	if o == nil || IsNil(o.TuntermRegistered) {
		var ret bool
		return ret
	}
	return *o.TuntermRegistered
}

// GetTuntermRegisteredOk returns a tuple with the TuntermRegistered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermRegisteredOk() (*bool, bool) {
	if o == nil || IsNil(o.TuntermRegistered) {
		return nil, false
	}
	return o.TuntermRegistered, true
}

// HasTuntermRegistered returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermRegistered() bool {
	if o != nil && !IsNil(o.TuntermRegistered) {
		return true
	}

	return false
}

// SetTuntermRegistered gets a reference to the given bool and assigns it to the TuntermRegistered field.
func (o *Mxedge) SetTuntermRegistered(v bool) {
	o.TuntermRegistered = &v
}

// GetTuntermSwitchConfig returns the TuntermSwitchConfig field value if set, zero value otherwise.
func (o *Mxedge) GetTuntermSwitchConfig() MxedgeTuntermSwitchConfigs {
	if o == nil || IsNil(o.TuntermSwitchConfig) {
		var ret MxedgeTuntermSwitchConfigs
		return ret
	}
	return *o.TuntermSwitchConfig
}

// GetTuntermSwitchConfigOk returns a tuple with the TuntermSwitchConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetTuntermSwitchConfigOk() (*MxedgeTuntermSwitchConfigs, bool) {
	if o == nil || IsNil(o.TuntermSwitchConfig) {
		return nil, false
	}
	return o.TuntermSwitchConfig, true
}

// HasTuntermSwitchConfig returns a boolean if a field has been set.
func (o *Mxedge) HasTuntermSwitchConfig() bool {
	if o != nil && !IsNil(o.TuntermSwitchConfig) {
		return true
	}

	return false
}

// SetTuntermSwitchConfig gets a reference to the given MxedgeTuntermSwitchConfigs and assigns it to the TuntermSwitchConfig field.
func (o *Mxedge) SetTuntermSwitchConfig(v MxedgeTuntermSwitchConfigs) {
	o.TuntermSwitchConfig = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *Mxedge) GetVersions() MxedgeVersions {
	if o == nil || IsNil(o.Versions) {
		var ret MxedgeVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxedge) GetVersionsOk() (*MxedgeVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *Mxedge) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given MxedgeVersions and assigns it to the Versions field.
func (o *Mxedge) SetVersions(v MxedgeVersions) {
	o.Versions = &v
}

func (o Mxedge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mxedge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Magic) {
		toSerialize["magic"] = o.Magic
	}
	toSerialize["model"] = o.Model
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.MxagentRegistered) {
		toSerialize["mxagent_registered"] = o.MxagentRegistered
	}
	if !IsNil(o.MxclusterId) {
		toSerialize["mxcluster_id"] = o.MxclusterId
	}
	if !IsNil(o.MxedgeMgmt) {
		toSerialize["mxedge_mgmt"] = o.MxedgeMgmt
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.NtpServers) {
		toSerialize["ntp_servers"] = o.NtpServers
	}
	if !IsNil(o.OobIpConfig) {
		toSerialize["oob_ip_config"] = o.OobIpConfig
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.TuntermDhcpdConfig) {
		toSerialize["tunterm_dhcpd_config"] = o.TuntermDhcpdConfig
	}
	if !IsNil(o.TuntermExtraRoutes) {
		toSerialize["tunterm_extra_routes"] = o.TuntermExtraRoutes
	}
	if !IsNil(o.TuntermIgmpSnoopingConfig) {
		toSerialize["tunterm_igmp_snooping_config"] = o.TuntermIgmpSnoopingConfig
	}
	if !IsNil(o.TuntermIpConfig) {
		toSerialize["tunterm_ip_config"] = o.TuntermIpConfig
	}
	if !IsNil(o.TuntermMonitoring) {
		toSerialize["tunterm_monitoring"] = o.TuntermMonitoring
	}
	if !IsNil(o.TuntermMulticastConfig) {
		toSerialize["tunterm_multicast_config"] = o.TuntermMulticastConfig
	}
	if !IsNil(o.TuntermOtherIpConfigs) {
		toSerialize["tunterm_other_ip_configs"] = o.TuntermOtherIpConfigs
	}
	if !IsNil(o.TuntermPortConfig) {
		toSerialize["tunterm_port_config"] = o.TuntermPortConfig
	}
	if !IsNil(o.TuntermRegistered) {
		toSerialize["tunterm_registered"] = o.TuntermRegistered
	}
	if !IsNil(o.TuntermSwitchConfig) {
		toSerialize["tunterm_switch_config"] = o.TuntermSwitchConfig
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Mxedge) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMxedge := _Mxedge{}

	err = json.Unmarshal(data, &varMxedge)

	if err != nil {
		return err
	}

	*o = Mxedge(varMxedge)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "id")
		delete(additionalProperties, "magic")
		delete(additionalProperties, "model")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "mxagent_registered")
		delete(additionalProperties, "mxcluster_id")
		delete(additionalProperties, "mxedge_mgmt")
		delete(additionalProperties, "name")
		delete(additionalProperties, "note")
		delete(additionalProperties, "ntp_servers")
		delete(additionalProperties, "oob_ip_config")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "services")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "tunterm_dhcpd_config")
		delete(additionalProperties, "tunterm_extra_routes")
		delete(additionalProperties, "tunterm_igmp_snooping_config")
		delete(additionalProperties, "tunterm_ip_config")
		delete(additionalProperties, "tunterm_monitoring")
		delete(additionalProperties, "tunterm_multicast_config")
		delete(additionalProperties, "tunterm_other_ip_configs")
		delete(additionalProperties, "tunterm_port_config")
		delete(additionalProperties, "tunterm_registered")
		delete(additionalProperties, "tunterm_switch_config")
		delete(additionalProperties, "versions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedge struct {
	value *Mxedge
	isSet bool
}

func (v NullableMxedge) Get() *Mxedge {
	return v.value
}

func (v *NullableMxedge) Set(val *Mxedge) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedge) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedge(val *Mxedge) *NullableMxedge {
	return &NullableMxedge{value: val, isSet: true}
}

func (v NullableMxedge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



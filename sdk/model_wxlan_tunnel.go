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

// checks if the WxlanTunnel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WxlanTunnel{}

// WxlanTunnel WxLAn Tunnel
type WxlanTunnel struct {
	CreatedTime *float32 `json:"created_time,omitempty"`
	Dmvpn *WxlanTunnelDmvpn `json:"dmvpn,omitempty"`
	// determined during creation time and cannot be toggled. A management tunnel cannot be used by wxlan rule or by wlan
	ForMgmt *bool `json:"for_mgmt,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	// in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries.
	HelloInterval *int32 `json:"hello_interval,omitempty"`
	HelloRetries *int32 `json:"hello_retries,omitempty"`
	// optional, overwrite the hostname in SCCRQ control message, default is “” or null, %H and %M can be used, which will be replace with corresponding values: * %H: name of the ap if provided (and will be stripped so it can be used for hostname) and fallbacks to MAC * %M: MAC (e.g. 5c5b350e0060)
	Hostname *string `json:"hostname,omitempty"`
	Id *string `json:"id,omitempty"`
	Ipsec *WxlanTunnelIpsec `json:"ipsec,omitempty"`
	// whether it’s static/unmanaged (i.e. no control session). As the session configurations are not compatible, cannot be toggled.
	IsStatic *bool `json:"is_static,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	// 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU
	Mtu *int32 `json:"mtu,omitempty"`
	// The name of the tunnel
	Name string `json:"name"`
	OrgId *string `json:"org_id,omitempty"`
	// list of remote peers’ IP or hostname
	Peers []string `json:"peers,omitempty"`
	// optional, overwrite the router-id in SCCRQ control message, default is “0” or null, can also be an IPv4 address
	RouterId *string `json:"router_id,omitempty"`
	// secret, ‘’ if no auth is used
	Secret *string `json:"secret,omitempty"`
	// sessions to be established with the tunnel. Has to be >= 1 in order for this tunnel to be useful. For management tunnel, it can only have 1
	Sessions []WxlanTunnelSession `json:"sessions,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	// udp port if `use_udp`==`true`
	UdpPort *int32 `json:"udp_port,omitempty"`
	// whether to use UDP instead of IP (proto=115, which is default of L2TPv3)
	UseUdp *bool `json:"use_udp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WxlanTunnel WxlanTunnel

// NewWxlanTunnel instantiates a new WxlanTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWxlanTunnel(name string) *WxlanTunnel {
	this := WxlanTunnel{}
	var forMgmt bool = false
	this.ForMgmt = &forMgmt
	var helloInterval int32 = 60
	this.HelloInterval = &helloInterval
	var helloRetries int32 = 7
	this.HelloRetries = &helloRetries
	var isStatic bool = false
	this.IsStatic = &isStatic
	var mtu int32 = 0
	this.Mtu = &mtu
	this.Name = name
	var useUdp bool = false
	this.UseUdp = &useUdp
	return &this
}

// NewWxlanTunnelWithDefaults instantiates a new WxlanTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWxlanTunnelWithDefaults() *WxlanTunnel {
	this := WxlanTunnel{}
	var forMgmt bool = false
	this.ForMgmt = &forMgmt
	var helloInterval int32 = 60
	this.HelloInterval = &helloInterval
	var helloRetries int32 = 7
	this.HelloRetries = &helloRetries
	var isStatic bool = false
	this.IsStatic = &isStatic
	var mtu int32 = 0
	this.Mtu = &mtu
	var useUdp bool = false
	this.UseUdp = &useUdp
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *WxlanTunnel) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *WxlanTunnel) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *WxlanTunnel) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetDmvpn returns the Dmvpn field value if set, zero value otherwise.
func (o *WxlanTunnel) GetDmvpn() WxlanTunnelDmvpn {
	if o == nil || IsNil(o.Dmvpn) {
		var ret WxlanTunnelDmvpn
		return ret
	}
	return *o.Dmvpn
}

// GetDmvpnOk returns a tuple with the Dmvpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetDmvpnOk() (*WxlanTunnelDmvpn, bool) {
	if o == nil || IsNil(o.Dmvpn) {
		return nil, false
	}
	return o.Dmvpn, true
}

// HasDmvpn returns a boolean if a field has been set.
func (o *WxlanTunnel) HasDmvpn() bool {
	if o != nil && !IsNil(o.Dmvpn) {
		return true
	}

	return false
}

// SetDmvpn gets a reference to the given WxlanTunnelDmvpn and assigns it to the Dmvpn field.
func (o *WxlanTunnel) SetDmvpn(v WxlanTunnelDmvpn) {
	o.Dmvpn = &v
}

// GetForMgmt returns the ForMgmt field value if set, zero value otherwise.
func (o *WxlanTunnel) GetForMgmt() bool {
	if o == nil || IsNil(o.ForMgmt) {
		var ret bool
		return ret
	}
	return *o.ForMgmt
}

// GetForMgmtOk returns a tuple with the ForMgmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetForMgmtOk() (*bool, bool) {
	if o == nil || IsNil(o.ForMgmt) {
		return nil, false
	}
	return o.ForMgmt, true
}

// HasForMgmt returns a boolean if a field has been set.
func (o *WxlanTunnel) HasForMgmt() bool {
	if o != nil && !IsNil(o.ForMgmt) {
		return true
	}

	return false
}

// SetForMgmt gets a reference to the given bool and assigns it to the ForMgmt field.
func (o *WxlanTunnel) SetForMgmt(v bool) {
	o.ForMgmt = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *WxlanTunnel) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *WxlanTunnel) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *WxlanTunnel) SetForSite(v bool) {
	o.ForSite = &v
}

// GetHelloInterval returns the HelloInterval field value if set, zero value otherwise.
func (o *WxlanTunnel) GetHelloInterval() int32 {
	if o == nil || IsNil(o.HelloInterval) {
		var ret int32
		return ret
	}
	return *o.HelloInterval
}

// GetHelloIntervalOk returns a tuple with the HelloInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetHelloIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.HelloInterval) {
		return nil, false
	}
	return o.HelloInterval, true
}

// HasHelloInterval returns a boolean if a field has been set.
func (o *WxlanTunnel) HasHelloInterval() bool {
	if o != nil && !IsNil(o.HelloInterval) {
		return true
	}

	return false
}

// SetHelloInterval gets a reference to the given int32 and assigns it to the HelloInterval field.
func (o *WxlanTunnel) SetHelloInterval(v int32) {
	o.HelloInterval = &v
}

// GetHelloRetries returns the HelloRetries field value if set, zero value otherwise.
func (o *WxlanTunnel) GetHelloRetries() int32 {
	if o == nil || IsNil(o.HelloRetries) {
		var ret int32
		return ret
	}
	return *o.HelloRetries
}

// GetHelloRetriesOk returns a tuple with the HelloRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetHelloRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.HelloRetries) {
		return nil, false
	}
	return o.HelloRetries, true
}

// HasHelloRetries returns a boolean if a field has been set.
func (o *WxlanTunnel) HasHelloRetries() bool {
	if o != nil && !IsNil(o.HelloRetries) {
		return true
	}

	return false
}

// SetHelloRetries gets a reference to the given int32 and assigns it to the HelloRetries field.
func (o *WxlanTunnel) SetHelloRetries(v int32) {
	o.HelloRetries = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *WxlanTunnel) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *WxlanTunnel) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *WxlanTunnel) SetHostname(v string) {
	o.Hostname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WxlanTunnel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WxlanTunnel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WxlanTunnel) SetId(v string) {
	o.Id = &v
}

// GetIpsec returns the Ipsec field value if set, zero value otherwise.
func (o *WxlanTunnel) GetIpsec() WxlanTunnelIpsec {
	if o == nil || IsNil(o.Ipsec) {
		var ret WxlanTunnelIpsec
		return ret
	}
	return *o.Ipsec
}

// GetIpsecOk returns a tuple with the Ipsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetIpsecOk() (*WxlanTunnelIpsec, bool) {
	if o == nil || IsNil(o.Ipsec) {
		return nil, false
	}
	return o.Ipsec, true
}

// HasIpsec returns a boolean if a field has been set.
func (o *WxlanTunnel) HasIpsec() bool {
	if o != nil && !IsNil(o.Ipsec) {
		return true
	}

	return false
}

// SetIpsec gets a reference to the given WxlanTunnelIpsec and assigns it to the Ipsec field.
func (o *WxlanTunnel) SetIpsec(v WxlanTunnelIpsec) {
	o.Ipsec = &v
}

// GetIsStatic returns the IsStatic field value if set, zero value otherwise.
func (o *WxlanTunnel) GetIsStatic() bool {
	if o == nil || IsNil(o.IsStatic) {
		var ret bool
		return ret
	}
	return *o.IsStatic
}

// GetIsStaticOk returns a tuple with the IsStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetIsStaticOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStatic) {
		return nil, false
	}
	return o.IsStatic, true
}

// HasIsStatic returns a boolean if a field has been set.
func (o *WxlanTunnel) HasIsStatic() bool {
	if o != nil && !IsNil(o.IsStatic) {
		return true
	}

	return false
}

// SetIsStatic gets a reference to the given bool and assigns it to the IsStatic field.
func (o *WxlanTunnel) SetIsStatic(v bool) {
	o.IsStatic = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *WxlanTunnel) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *WxlanTunnel) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *WxlanTunnel) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *WxlanTunnel) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu) {
		var ret int32
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetMtuOk() (*int32, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *WxlanTunnel) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int32 and assigns it to the Mtu field.
func (o *WxlanTunnel) SetMtu(v int32) {
	o.Mtu = &v
}

// GetName returns the Name field value
func (o *WxlanTunnel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WxlanTunnel) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *WxlanTunnel) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *WxlanTunnel) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *WxlanTunnel) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *WxlanTunnel) GetPeers() []string {
	if o == nil || IsNil(o.Peers) {
		var ret []string
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetPeersOk() ([]string, bool) {
	if o == nil || IsNil(o.Peers) {
		return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *WxlanTunnel) HasPeers() bool {
	if o != nil && !IsNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []string and assigns it to the Peers field.
func (o *WxlanTunnel) SetPeers(v []string) {
	o.Peers = v
}

// GetRouterId returns the RouterId field value if set, zero value otherwise.
func (o *WxlanTunnel) GetRouterId() string {
	if o == nil || IsNil(o.RouterId) {
		var ret string
		return ret
	}
	return *o.RouterId
}

// GetRouterIdOk returns a tuple with the RouterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetRouterIdOk() (*string, bool) {
	if o == nil || IsNil(o.RouterId) {
		return nil, false
	}
	return o.RouterId, true
}

// HasRouterId returns a boolean if a field has been set.
func (o *WxlanTunnel) HasRouterId() bool {
	if o != nil && !IsNil(o.RouterId) {
		return true
	}

	return false
}

// SetRouterId gets a reference to the given string and assigns it to the RouterId field.
func (o *WxlanTunnel) SetRouterId(v string) {
	o.RouterId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WxlanTunnel) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WxlanTunnel) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WxlanTunnel) SetSecret(v string) {
	o.Secret = &v
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *WxlanTunnel) GetSessions() []WxlanTunnelSession {
	if o == nil || IsNil(o.Sessions) {
		var ret []WxlanTunnelSession
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetSessionsOk() ([]WxlanTunnelSession, bool) {
	if o == nil || IsNil(o.Sessions) {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *WxlanTunnel) HasSessions() bool {
	if o != nil && !IsNil(o.Sessions) {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []WxlanTunnelSession and assigns it to the Sessions field.
func (o *WxlanTunnel) SetSessions(v []WxlanTunnelSession) {
	o.Sessions = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *WxlanTunnel) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *WxlanTunnel) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *WxlanTunnel) SetSiteId(v string) {
	o.SiteId = &v
}

// GetUdpPort returns the UdpPort field value if set, zero value otherwise.
func (o *WxlanTunnel) GetUdpPort() int32 {
	if o == nil || IsNil(o.UdpPort) {
		var ret int32
		return ret
	}
	return *o.UdpPort
}

// GetUdpPortOk returns a tuple with the UdpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetUdpPortOk() (*int32, bool) {
	if o == nil || IsNil(o.UdpPort) {
		return nil, false
	}
	return o.UdpPort, true
}

// HasUdpPort returns a boolean if a field has been set.
func (o *WxlanTunnel) HasUdpPort() bool {
	if o != nil && !IsNil(o.UdpPort) {
		return true
	}

	return false
}

// SetUdpPort gets a reference to the given int32 and assigns it to the UdpPort field.
func (o *WxlanTunnel) SetUdpPort(v int32) {
	o.UdpPort = &v
}

// GetUseUdp returns the UseUdp field value if set, zero value otherwise.
func (o *WxlanTunnel) GetUseUdp() bool {
	if o == nil || IsNil(o.UseUdp) {
		var ret bool
		return ret
	}
	return *o.UseUdp
}

// GetUseUdpOk returns a tuple with the UseUdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnel) GetUseUdpOk() (*bool, bool) {
	if o == nil || IsNil(o.UseUdp) {
		return nil, false
	}
	return o.UseUdp, true
}

// HasUseUdp returns a boolean if a field has been set.
func (o *WxlanTunnel) HasUseUdp() bool {
	if o != nil && !IsNil(o.UseUdp) {
		return true
	}

	return false
}

// SetUseUdp gets a reference to the given bool and assigns it to the UseUdp field.
func (o *WxlanTunnel) SetUseUdp(v bool) {
	o.UseUdp = &v
}

func (o WxlanTunnel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WxlanTunnel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Dmvpn) {
		toSerialize["dmvpn"] = o.Dmvpn
	}
	if !IsNil(o.ForMgmt) {
		toSerialize["for_mgmt"] = o.ForMgmt
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.HelloInterval) {
		toSerialize["hello_interval"] = o.HelloInterval
	}
	if !IsNil(o.HelloRetries) {
		toSerialize["hello_retries"] = o.HelloRetries
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ipsec) {
		toSerialize["ipsec"] = o.Ipsec
	}
	if !IsNil(o.IsStatic) {
		toSerialize["is_static"] = o.IsStatic
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	if !IsNil(o.RouterId) {
		toSerialize["router_id"] = o.RouterId
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.Sessions) {
		toSerialize["sessions"] = o.Sessions
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.UdpPort) {
		toSerialize["udp_port"] = o.UdpPort
	}
	if !IsNil(o.UseUdp) {
		toSerialize["use_udp"] = o.UseUdp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WxlanTunnel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varWxlanTunnel := _WxlanTunnel{}

	err = json.Unmarshal(data, &varWxlanTunnel)

	if err != nil {
		return err
	}

	*o = WxlanTunnel(varWxlanTunnel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "dmvpn")
		delete(additionalProperties, "for_mgmt")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "hello_interval")
		delete(additionalProperties, "hello_retries")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ipsec")
		delete(additionalProperties, "is_static")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "mtu")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "peers")
		delete(additionalProperties, "router_id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "sessions")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "udp_port")
		delete(additionalProperties, "use_udp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWxlanTunnel struct {
	value *WxlanTunnel
	isSet bool
}

func (v NullableWxlanTunnel) Get() *WxlanTunnel {
	return v.value
}

func (v *NullableWxlanTunnel) Set(val *WxlanTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableWxlanTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableWxlanTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWxlanTunnel(val *WxlanTunnel) *NullableWxlanTunnel {
	return &NullableWxlanTunnel{value: val, isSet: true}
}

func (v NullableWxlanTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWxlanTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Mist API

> Version: **2406.1.17** > > Date: **July 5, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.17
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
)

// checks if the Gateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gateway{}

// Gateway device gateway
type Gateway struct {
	// additional CLI commands to append to the generated Junos config  **Note**: no check is done
	AdditionalConfigCmds []string `json:"additional_config_cmds,omitempty"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	DeviceprofileId *string `json:"deviceprofile_id,omitempty"`
	DhcpdConfig *DhcpdConfigs `json:"dhcpd_config,omitempty"`
	ExtraRoutes *map[string]GatewayExtraRoute `json:"extra_routes,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	Id *string `json:"id,omitempty"`
	Image1Url *string `json:"image1_url,omitempty"`
	Image2Url *string `json:"image2_url,omitempty"`
	Image3Url *string `json:"image3_url,omitempty"`
	// Property key is the network name
	IpConfig *map[string]GatewayTemplateIpConfig `json:"ip_config,omitempty"`
	Managed *bool `json:"managed,omitempty"`
	// map where the device belongs to
	MapId *string `json:"map_id,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	MspId *string `json:"msp_id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Property key is the network name or a CIDR
	Networks *map[string]GatewayNetwork `json:"networks,omitempty"`
	NtpServers []string `json:"ntp_servers,omitempty"`
	OobIpConfig *GatewayOobIpConfig `json:"oob_ip_config,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	// Property key is the port name or range (e.g. \"ge-0/0/0-10\")
	PortConfig *map[string]GatewayPortConfig `json:"port_config,omitempty"`
	PortMirroring *GatewayPortMirroring `json:"port_mirroring,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	// a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
	Vars *map[string]string `json:"vars,omitempty"`
	// x in pixel
	X *float32 `json:"x,omitempty"`
	// y in pixel
	Y *float32 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Gateway Gateway

// NewGateway instantiates a new Gateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateway() *Gateway {
	this := Gateway{}
	return &this
}

// NewGatewayWithDefaults instantiates a new Gateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayWithDefaults() *Gateway {
	this := Gateway{}
	return &this
}

// GetAdditionalConfigCmds returns the AdditionalConfigCmds field value if set, zero value otherwise.
func (o *Gateway) GetAdditionalConfigCmds() []string {
	if o == nil || IsNil(o.AdditionalConfigCmds) {
		var ret []string
		return ret
	}
	return o.AdditionalConfigCmds
}

// GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetAdditionalConfigCmdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalConfigCmds) {
		return nil, false
	}
	return o.AdditionalConfigCmds, true
}

// HasAdditionalConfigCmds returns a boolean if a field has been set.
func (o *Gateway) HasAdditionalConfigCmds() bool {
	if o != nil && !IsNil(o.AdditionalConfigCmds) {
		return true
	}

	return false
}

// SetAdditionalConfigCmds gets a reference to the given []string and assigns it to the AdditionalConfigCmds field.
func (o *Gateway) SetAdditionalConfigCmds(v []string) {
	o.AdditionalConfigCmds = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Gateway) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Gateway) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Gateway) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetDeviceprofileId returns the DeviceprofileId field value if set, zero value otherwise.
func (o *Gateway) GetDeviceprofileId() string {
	if o == nil || IsNil(o.DeviceprofileId) {
		var ret string
		return ret
	}
	return *o.DeviceprofileId
}

// GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetDeviceprofileIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceprofileId) {
		return nil, false
	}
	return o.DeviceprofileId, true
}

// HasDeviceprofileId returns a boolean if a field has been set.
func (o *Gateway) HasDeviceprofileId() bool {
	if o != nil && !IsNil(o.DeviceprofileId) {
		return true
	}

	return false
}

// SetDeviceprofileId gets a reference to the given string and assigns it to the DeviceprofileId field.
func (o *Gateway) SetDeviceprofileId(v string) {
	o.DeviceprofileId = &v
}

// GetDhcpdConfig returns the DhcpdConfig field value if set, zero value otherwise.
func (o *Gateway) GetDhcpdConfig() DhcpdConfigs {
	if o == nil || IsNil(o.DhcpdConfig) {
		var ret DhcpdConfigs
		return ret
	}
	return *o.DhcpdConfig
}

// GetDhcpdConfigOk returns a tuple with the DhcpdConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetDhcpdConfigOk() (*DhcpdConfigs, bool) {
	if o == nil || IsNil(o.DhcpdConfig) {
		return nil, false
	}
	return o.DhcpdConfig, true
}

// HasDhcpdConfig returns a boolean if a field has been set.
func (o *Gateway) HasDhcpdConfig() bool {
	if o != nil && !IsNil(o.DhcpdConfig) {
		return true
	}

	return false
}

// SetDhcpdConfig gets a reference to the given DhcpdConfigs and assigns it to the DhcpdConfig field.
func (o *Gateway) SetDhcpdConfig(v DhcpdConfigs) {
	o.DhcpdConfig = &v
}

// GetExtraRoutes returns the ExtraRoutes field value if set, zero value otherwise.
func (o *Gateway) GetExtraRoutes() map[string]GatewayExtraRoute {
	if o == nil || IsNil(o.ExtraRoutes) {
		var ret map[string]GatewayExtraRoute
		return ret
	}
	return *o.ExtraRoutes
}

// GetExtraRoutesOk returns a tuple with the ExtraRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetExtraRoutesOk() (*map[string]GatewayExtraRoute, bool) {
	if o == nil || IsNil(o.ExtraRoutes) {
		return nil, false
	}
	return o.ExtraRoutes, true
}

// HasExtraRoutes returns a boolean if a field has been set.
func (o *Gateway) HasExtraRoutes() bool {
	if o != nil && !IsNil(o.ExtraRoutes) {
		return true
	}

	return false
}

// SetExtraRoutes gets a reference to the given map[string]GatewayExtraRoute and assigns it to the ExtraRoutes field.
func (o *Gateway) SetExtraRoutes(v map[string]GatewayExtraRoute) {
	o.ExtraRoutes = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *Gateway) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *Gateway) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *Gateway) SetForSite(v bool) {
	o.ForSite = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Gateway) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Gateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Gateway) SetId(v string) {
	o.Id = &v
}

// GetImage1Url returns the Image1Url field value if set, zero value otherwise.
func (o *Gateway) GetImage1Url() string {
	if o == nil || IsNil(o.Image1Url) {
		var ret string
		return ret
	}
	return *o.Image1Url
}

// GetImage1UrlOk returns a tuple with the Image1Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetImage1UrlOk() (*string, bool) {
	if o == nil || IsNil(o.Image1Url) {
		return nil, false
	}
	return o.Image1Url, true
}

// HasImage1Url returns a boolean if a field has been set.
func (o *Gateway) HasImage1Url() bool {
	if o != nil && !IsNil(o.Image1Url) {
		return true
	}

	return false
}

// SetImage1Url gets a reference to the given string and assigns it to the Image1Url field.
func (o *Gateway) SetImage1Url(v string) {
	o.Image1Url = &v
}

// GetImage2Url returns the Image2Url field value if set, zero value otherwise.
func (o *Gateway) GetImage2Url() string {
	if o == nil || IsNil(o.Image2Url) {
		var ret string
		return ret
	}
	return *o.Image2Url
}

// GetImage2UrlOk returns a tuple with the Image2Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetImage2UrlOk() (*string, bool) {
	if o == nil || IsNil(o.Image2Url) {
		return nil, false
	}
	return o.Image2Url, true
}

// HasImage2Url returns a boolean if a field has been set.
func (o *Gateway) HasImage2Url() bool {
	if o != nil && !IsNil(o.Image2Url) {
		return true
	}

	return false
}

// SetImage2Url gets a reference to the given string and assigns it to the Image2Url field.
func (o *Gateway) SetImage2Url(v string) {
	o.Image2Url = &v
}

// GetImage3Url returns the Image3Url field value if set, zero value otherwise.
func (o *Gateway) GetImage3Url() string {
	if o == nil || IsNil(o.Image3Url) {
		var ret string
		return ret
	}
	return *o.Image3Url
}

// GetImage3UrlOk returns a tuple with the Image3Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetImage3UrlOk() (*string, bool) {
	if o == nil || IsNil(o.Image3Url) {
		return nil, false
	}
	return o.Image3Url, true
}

// HasImage3Url returns a boolean if a field has been set.
func (o *Gateway) HasImage3Url() bool {
	if o != nil && !IsNil(o.Image3Url) {
		return true
	}

	return false
}

// SetImage3Url gets a reference to the given string and assigns it to the Image3Url field.
func (o *Gateway) SetImage3Url(v string) {
	o.Image3Url = &v
}

// GetIpConfig returns the IpConfig field value if set, zero value otherwise.
func (o *Gateway) GetIpConfig() map[string]GatewayTemplateIpConfig {
	if o == nil || IsNil(o.IpConfig) {
		var ret map[string]GatewayTemplateIpConfig
		return ret
	}
	return *o.IpConfig
}

// GetIpConfigOk returns a tuple with the IpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetIpConfigOk() (*map[string]GatewayTemplateIpConfig, bool) {
	if o == nil || IsNil(o.IpConfig) {
		return nil, false
	}
	return o.IpConfig, true
}

// HasIpConfig returns a boolean if a field has been set.
func (o *Gateway) HasIpConfig() bool {
	if o != nil && !IsNil(o.IpConfig) {
		return true
	}

	return false
}

// SetIpConfig gets a reference to the given map[string]GatewayTemplateIpConfig and assigns it to the IpConfig field.
func (o *Gateway) SetIpConfig(v map[string]GatewayTemplateIpConfig) {
	o.IpConfig = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *Gateway) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *Gateway) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *Gateway) SetManaged(v bool) {
	o.Managed = &v
}

// GetMapId returns the MapId field value if set, zero value otherwise.
func (o *Gateway) GetMapId() string {
	if o == nil || IsNil(o.MapId) {
		var ret string
		return ret
	}
	return *o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetMapIdOk() (*string, bool) {
	if o == nil || IsNil(o.MapId) {
		return nil, false
	}
	return o.MapId, true
}

// HasMapId returns a boolean if a field has been set.
func (o *Gateway) HasMapId() bool {
	if o != nil && !IsNil(o.MapId) {
		return true
	}

	return false
}

// SetMapId gets a reference to the given string and assigns it to the MapId field.
func (o *Gateway) SetMapId(v string) {
	o.MapId = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Gateway) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Gateway) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *Gateway) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetMspId returns the MspId field value if set, zero value otherwise.
func (o *Gateway) GetMspId() string {
	if o == nil || IsNil(o.MspId) {
		var ret string
		return ret
	}
	return *o.MspId
}

// GetMspIdOk returns a tuple with the MspId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetMspIdOk() (*string, bool) {
	if o == nil || IsNil(o.MspId) {
		return nil, false
	}
	return o.MspId, true
}

// HasMspId returns a boolean if a field has been set.
func (o *Gateway) HasMspId() bool {
	if o != nil && !IsNil(o.MspId) {
		return true
	}

	return false
}

// SetMspId gets a reference to the given string and assigns it to the MspId field.
func (o *Gateway) SetMspId(v string) {
	o.MspId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Gateway) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Gateway) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Gateway) SetName(v string) {
	o.Name = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *Gateway) GetNetworks() map[string]GatewayNetwork {
	if o == nil || IsNil(o.Networks) {
		var ret map[string]GatewayNetwork
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetNetworksOk() (*map[string]GatewayNetwork, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *Gateway) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given map[string]GatewayNetwork and assigns it to the Networks field.
func (o *Gateway) SetNetworks(v map[string]GatewayNetwork) {
	o.Networks = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *Gateway) GetNtpServers() []string {
	if o == nil || IsNil(o.NtpServers) {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetNtpServersOk() ([]string, bool) {
	if o == nil || IsNil(o.NtpServers) {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *Gateway) HasNtpServers() bool {
	if o != nil && !IsNil(o.NtpServers) {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *Gateway) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetOobIpConfig returns the OobIpConfig field value if set, zero value otherwise.
func (o *Gateway) GetOobIpConfig() GatewayOobIpConfig {
	if o == nil || IsNil(o.OobIpConfig) {
		var ret GatewayOobIpConfig
		return ret
	}
	return *o.OobIpConfig
}

// GetOobIpConfigOk returns a tuple with the OobIpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetOobIpConfigOk() (*GatewayOobIpConfig, bool) {
	if o == nil || IsNil(o.OobIpConfig) {
		return nil, false
	}
	return o.OobIpConfig, true
}

// HasOobIpConfig returns a boolean if a field has been set.
func (o *Gateway) HasOobIpConfig() bool {
	if o != nil && !IsNil(o.OobIpConfig) {
		return true
	}

	return false
}

// SetOobIpConfig gets a reference to the given GatewayOobIpConfig and assigns it to the OobIpConfig field.
func (o *Gateway) SetOobIpConfig(v GatewayOobIpConfig) {
	o.OobIpConfig = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Gateway) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Gateway) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Gateway) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPortConfig returns the PortConfig field value if set, zero value otherwise.
func (o *Gateway) GetPortConfig() map[string]GatewayPortConfig {
	if o == nil || IsNil(o.PortConfig) {
		var ret map[string]GatewayPortConfig
		return ret
	}
	return *o.PortConfig
}

// GetPortConfigOk returns a tuple with the PortConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetPortConfigOk() (*map[string]GatewayPortConfig, bool) {
	if o == nil || IsNil(o.PortConfig) {
		return nil, false
	}
	return o.PortConfig, true
}

// HasPortConfig returns a boolean if a field has been set.
func (o *Gateway) HasPortConfig() bool {
	if o != nil && !IsNil(o.PortConfig) {
		return true
	}

	return false
}

// SetPortConfig gets a reference to the given map[string]GatewayPortConfig and assigns it to the PortConfig field.
func (o *Gateway) SetPortConfig(v map[string]GatewayPortConfig) {
	o.PortConfig = &v
}

// GetPortMirroring returns the PortMirroring field value if set, zero value otherwise.
func (o *Gateway) GetPortMirroring() GatewayPortMirroring {
	if o == nil || IsNil(o.PortMirroring) {
		var ret GatewayPortMirroring
		return ret
	}
	return *o.PortMirroring
}

// GetPortMirroringOk returns a tuple with the PortMirroring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetPortMirroringOk() (*GatewayPortMirroring, bool) {
	if o == nil || IsNil(o.PortMirroring) {
		return nil, false
	}
	return o.PortMirroring, true
}

// HasPortMirroring returns a boolean if a field has been set.
func (o *Gateway) HasPortMirroring() bool {
	if o != nil && !IsNil(o.PortMirroring) {
		return true
	}

	return false
}

// SetPortMirroring gets a reference to the given GatewayPortMirroring and assigns it to the PortMirroring field.
func (o *Gateway) SetPortMirroring(v GatewayPortMirroring) {
	o.PortMirroring = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *Gateway) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *Gateway) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *Gateway) SetSiteId(v string) {
	o.SiteId = &v
}

// GetVars returns the Vars field value if set, zero value otherwise.
func (o *Gateway) GetVars() map[string]string {
	if o == nil || IsNil(o.Vars) {
		var ret map[string]string
		return ret
	}
	return *o.Vars
}

// GetVarsOk returns a tuple with the Vars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetVarsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Vars) {
		return nil, false
	}
	return o.Vars, true
}

// HasVars returns a boolean if a field has been set.
func (o *Gateway) HasVars() bool {
	if o != nil && !IsNil(o.Vars) {
		return true
	}

	return false
}

// SetVars gets a reference to the given map[string]string and assigns it to the Vars field.
func (o *Gateway) SetVars(v map[string]string) {
	o.Vars = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *Gateway) GetX() float32 {
	if o == nil || IsNil(o.X) {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetXOk() (*float32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *Gateway) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *Gateway) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *Gateway) GetY() float32 {
	if o == nil || IsNil(o.Y) {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetYOk() (*float32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *Gateway) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *Gateway) SetY(v float32) {
	o.Y = &v
}

func (o Gateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalConfigCmds) {
		toSerialize["additional_config_cmds"] = o.AdditionalConfigCmds
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.DeviceprofileId) {
		toSerialize["deviceprofile_id"] = o.DeviceprofileId
	}
	if !IsNil(o.DhcpdConfig) {
		toSerialize["dhcpd_config"] = o.DhcpdConfig
	}
	if !IsNil(o.ExtraRoutes) {
		toSerialize["extra_routes"] = o.ExtraRoutes
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Image1Url) {
		toSerialize["image1_url"] = o.Image1Url
	}
	if !IsNil(o.Image2Url) {
		toSerialize["image2_url"] = o.Image2Url
	}
	if !IsNil(o.Image3Url) {
		toSerialize["image3_url"] = o.Image3Url
	}
	if !IsNil(o.IpConfig) {
		toSerialize["ip_config"] = o.IpConfig
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	if !IsNil(o.MapId) {
		toSerialize["map_id"] = o.MapId
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.MspId) {
		toSerialize["msp_id"] = o.MspId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
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
	if !IsNil(o.PortConfig) {
		toSerialize["port_config"] = o.PortConfig
	}
	if !IsNil(o.PortMirroring) {
		toSerialize["port_mirroring"] = o.PortMirroring
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Vars) {
		toSerialize["vars"] = o.Vars
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Gateway) UnmarshalJSON(data []byte) (err error) {
	varGateway := _Gateway{}

	err = json.Unmarshal(data, &varGateway)

	if err != nil {
		return err
	}

	*o = Gateway(varGateway)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_config_cmds")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "deviceprofile_id")
		delete(additionalProperties, "dhcpd_config")
		delete(additionalProperties, "extra_routes")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "id")
		delete(additionalProperties, "image1_url")
		delete(additionalProperties, "image2_url")
		delete(additionalProperties, "image3_url")
		delete(additionalProperties, "ip_config")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "msp_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "networks")
		delete(additionalProperties, "ntp_servers")
		delete(additionalProperties, "oob_ip_config")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "port_config")
		delete(additionalProperties, "port_mirroring")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "vars")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGateway struct {
	value *Gateway
	isSet bool
}

func (v NullableGateway) Get() *Gateway {
	return v.value
}

func (v *NullableGateway) Set(val *Gateway) {
	v.value = val
	v.isSet = true
}

func (v NullableGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGateway(val *Gateway) *NullableGateway {
	return &NullableGateway{value: val, isSet: true}
}

func (v NullableGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



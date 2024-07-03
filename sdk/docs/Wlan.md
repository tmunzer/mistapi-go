# Wlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctImmediateUpdate** | Pointer to **bool** | enable coa-immediate-update and address-change-immediate-update on the access profile. | [optional] [default to false]
**AcctInterimInterval** | Pointer to **int32** | how frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled | [optional] [default to 0]
**AcctServers** | Pointer to [**[]RadiusAcctServer**](RadiusAcctServer.md) | list of RADIUS accounting servers, optional, order matters where the first one is treated as primary | [optional] 
**Airwatch** | Pointer to [**WlanAirwatch**](WlanAirwatch.md) |  | [optional] 
**AllowIpv6Ndp** | Pointer to **bool** | only applicable when limit_bcast&#x3D;&#x3D;true, which allows or disallows ipv6 Neighbor Discovery packets to go through | [optional] [default to true]
**AllowMdns** | Pointer to **bool** | only applicable when limit_bcast&#x3D;&#x3D;true, which allows mDNS / Bonjour packets to go through | [optional] [default to false]
**AllowSsdp** | Pointer to **bool** | only applicable when &#x60;limit_bcast&#x60;&#x3D;&#x3D;&#x60;tru&#x60;e, which allows SSDP | [optional] [default to false]
**ApIds** | Pointer to **[]string** | list of device ids | [optional] 
**AppLimit** | Pointer to [**WlanAppLimit**](WlanAppLimit.md) |  | [optional] 
**AppQos** | Pointer to [**WlanAppQos**](WlanAppQos.md) |  | [optional] 
**ApplyTo** | Pointer to [**WlanApplyTo**](WlanApplyTo.md) |  | [optional] 
**ArpFilter** | Pointer to **bool** | whether to enable smart arp filter | [optional] [default to false]
**Auth** | Pointer to [**WlanAuth**](WlanAuth.md) |  | [optional] 
**AuthServerSelection** | Pointer to [**WlanAuthServerSelection**](WlanAuthServerSelection.md) |  | [optional] [default to WLANAUTHSERVERSELECTION_ORDERED]
**AuthServers** | Pointer to [**[]RadiusAuthServer**](RadiusAuthServer.md) | list of RADIUS authentication servers, at least one is needed if &#x60;auth type&#x60;&#x3D;&#x3D;&#x60;eap&#x60;, order matters where the first one is treated as primary | [optional] 
**AuthServersNasId** | Pointer to **NullableString** | optional, up to 48 bytes, will be dynamically generated if not provided. used only for authentication servers | [optional] 
**AuthServersNasIp** | Pointer to **NullableString** | optional, NAS-IP-ADDRESS to use | [optional] 
**AuthServersRetries** | Pointer to **int32** | radius auth session retries. Following fast timers are set if “fast_dot1x_timers” knob is enabled. ‘retries’ are set to value of auth_servers_retries. ‘max-requests’ is also set when setting auth_servers_retries and is set to default value to 3. | [optional] [default to 2]
**AuthServersTimeout** | Pointer to **int32** | radius auth session timeout. Following fast timers are set if “fast_dot1x_timers” knob is enabled. ‘quite-period’ and ‘transmit-period’ are set to half the value of auth_servers_timeout. ‘supplicant-timeout’ is also set when setting auth_servers_timeout and is set to default value of 10. | [optional] [default to 5]
**Band** | Pointer to **string** | &#x60;band&#x60; is deprecated and kept for backward compability. Use bands instead | [optional] 
**BandSteer** | Pointer to **bool** | whether to enable band_steering, this works only when band&#x3D;&#x3D;both | [optional] [default to false]
**BandSteerForceBand5** | Pointer to **bool** | force dual_band capable client to connect to 5G | [optional] [default to false]
**Bands** | Pointer to [**[]Dot11Band**](Dot11Band.md) | list of radios that the wlan should apply to | [optional] [default to ["24","5"]]
**BlockBlacklistClients** | Pointer to **bool** | whether to block the clients in the blacklist (up to first 256 macs) | [optional] [default to false]
**Bonjour** | Pointer to [**WlanBonjour**](WlanBonjour.md) |  | [optional] 
**CiscoCwa** | Pointer to [**WlanCiscoCwa**](WlanCiscoCwa.md) |  | [optional] 
**ClientLimitDown** | Pointer to **int32** | kbps | [optional] 
**ClientLimitDownEnabled** | Pointer to **bool** | if downlink limiting per-client is enabled | [optional] [default to false]
**ClientLimitUp** | Pointer to **int32** | kbps | [optional] 
**ClientLimitUpEnabled** | Pointer to **bool** | if uplink limiting per-client is enabled | [optional] [default to false]
**CoaServers** | Pointer to [**[]CoaServer**](CoaServer.md) | list of COA (change of authorization) servers, optional | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Disable11ax** | Pointer to **bool** | some old WLAN drivers may not be compatible | [optional] [default to false]
**DisableHtVhtRates** | Pointer to **bool** | to disable ht or vht rates | [optional] [default to false]
**DisableUapsd** | Pointer to **bool** | whether to disable U-APSD | [optional] [default to false]
**DisableV1RoamNotify** | Pointer to **bool** | disable sending v2 roam notification messages | [optional] [default to false]
**DisableV2RoamNotify** | Pointer to **bool** | disable sending v2 roam notification messages | [optional] [default to false]
**DisableWmm** | Pointer to **bool** | whether to disable WMM | [optional] [default to false]
**DnsServerRewrite** | Pointer to [**NullableWlanDnsServerRewrite**](WlanDnsServerRewrite.md) |  | [optional] 
**Dtim** | Pointer to **int32** |  | [optional] [default to 2]
**DynamicPsk** | Pointer to [**NullableWlanDynamicPsk**](WlanDynamicPsk.md) |  | [optional] 
**DynamicVlan** | Pointer to [**NullableWlanDynamicVlan**](WlanDynamicVlan.md) |  | [optional] 
**EnableLocalKeycaching** | Pointer to **bool** | enable AP-AP keycaching via multicast | [optional] [default to false]
**EnableWirelessBridging** | Pointer to **bool** | by default, we&#39;d inspect all DHCP packets and drop those unrelated to the wireless client itself in the case where client is a wireless bridge (DHCP packets for other MACs will need to be orwarded), wireless_bridging can be enabled | [optional] [default to false]
**EnableWirelessBridgingDhcpTracking** | Pointer to **bool** | if the client bridge is doing DHCP on behalf of other devices (L2-NAT), enable dhcp_tracking will cut down DHCP response packets to be forwarded to wireless | [optional] [default to false]
**Enabled** | Pointer to **bool** | if this wlan is enabled | [optional] [default to true]
**FastDot1xTimers** | Pointer to **bool** | if set to true, sets default fast-timers with values calculated from ‘auth_servers_timeout’ and ‘auth_server_retries’. | [optional] [default to false]
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**HideSsid** | Pointer to **bool** | whether to hide SSID in beacon | [optional] [default to false]
**HostnameIe** | Pointer to **bool** | include hostname inside IE in AP beacons / probe responses | [optional] [default to false]
**Hotspot20** | Pointer to [**WlanHotspot20**](WlanHotspot20.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**InjectDhcpOption82** | Pointer to [**WlanInjectDhcpOption82**](WlanInjectDhcpOption82.md) |  | [optional] 
**Interface** | Pointer to [**WlanInterface**](WlanInterface.md) |  | [optional] [default to WLANINTERFACE_ALL]
**Isolation** | Pointer to **bool** | whether to stop clients to talk to each other | [optional] [default to false]
**L2Isolation** | Pointer to **bool** | if isolation is enabled, whether to deny clients to talk to L2 on the LAN | [optional] [default to false]
**LegacyOverds** | Pointer to **bool** | legacy devices requires the Over-DS (for Fast BSS Transition) bit set (while our chip doesn’t support it). Warning! Enabling this will cause problem for iOS devices. | [optional] [default to false]
**LimitBcast** | Pointer to **bool** | whether to limit broadcast packets going to wireless (i.e. only allow certain bcast packets to go through) | [optional] [default to false]
**LimitProbeResponse** | Pointer to **bool** | limit probe response base on some heuristic rules | [optional] [default to false]
**MaxIdletime** | Pointer to **int32** | max idle time in seconds | [optional] [default to 1800]
**MistNac** | Pointer to [**WlanMistNac**](WlanMistNac.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**MspId** | Pointer to **string** |  | [optional] [readonly] 
**MxtunnelId** | Pointer to **string** | (deprecated, use mxtunnel_ids instead) when &#x60;interface&#x60;&#x3D;&#x3D;&#x60;mxtunnel&#x60;, id of the Mist Tunnel | [optional] 
**MxtunnelIds** | Pointer to **[]string** | when &#x60;interface&#x60;&#x3D;&#x60;mxtunnel&#x60;, id of the Mist Tunnel | [optional] 
**MxtunnelName** | Pointer to **[]string** | when &#x60;interface&#x60;&#x3D;&#x60;site_medge&#x60;, name of the mxtunnel that in mxtunnels under Site Setting | [optional] 
**NoStaticDns** | Pointer to **bool** | whether to only allow client to use DNS that we’ve learned from DHCP response | [optional] [default to false]
**NoStaticIp** | Pointer to **bool** | whether to only allow client that we’ve learned from DHCP exchange to talk | [optional] [default to false]
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Portal** | Pointer to [**WlanPortal**](WlanPortal.md) |  | [optional] 
**PortalAllowedHostnames** | Pointer to **[]string** | list of hostnames without http(s):// (matched by substring) | [optional] [default to []]
**PortalAllowedSubnets** | Pointer to **[]string** | list of CIDRs | [optional] [default to []]
**PortalApiSecret** | Pointer to **NullableString** | api secret (auto-generated) that can be used to sign guest authorization requests | [optional] [default to ""]
**PortalDeniedHostnames** | Pointer to **[]string** | list of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames | [optional] [default to []]
**PortalImage** | Pointer to **NullableString** | Url of portal background image | [optional] [readonly] [default to ""]
**PortalSsoUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**PortalTemplateUrl** | Pointer to **NullableString** | N.B portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template. | [optional] [readonly] [default to ""]
**Qos** | Pointer to [**WlanQos**](WlanQos.md) |  | [optional] 
**Radsec** | Pointer to [**Radsec**](Radsec.md) |  | [optional] 
**Rateset** | Pointer to [**WlanRateset**](WlanRateset.md) |  | [optional] 
**RoamMode** | Pointer to [**WlanRoamMode**](WlanRoamMode.md) |  | [optional] [default to WLANROAMMODE_NONE]
**Schedule** | Pointer to [**WlanSchedule**](WlanSchedule.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SleExcluded** | Pointer to **bool** | whether to exclude this WLAN from SLE metrics | [optional] [default to false]
**Ssid** | **string** | the name of the SSID | 
**TemplateId** | Pointer to **NullableString** |  | [optional] 
**Thumbnail** | Pointer to **NullableString** | Url of portal background image thumbnail | [optional] [readonly] [default to ""]
**UseEapolV1** | Pointer to **bool** | if &#x60;auth.type&#x60;&#x3D;&#x3D;’eap’ or ‘psk’, should only be set for legacy client, such as pre-2004, 802.11b devices | [optional] [default to false]
**VlanEnabled** | Pointer to **bool** | if vlan tagging is enabled | [optional] [default to false]
**VlanId** | Pointer to **NullableInt32** |  | [optional] 
**VlanIds** | Pointer to **[]int32** | vlan_ids to use when there’s no match from RA | [optional] 
**VlanPooling** | Pointer to **bool** | vlan pooling allows AP to place client on different VLAN using a deterministic algorithm | [optional] [default to false]
**WlanLimitDown** | Pointer to **NullableInt32** | kbps | [optional] 
**WlanLimitDownEnabled** | Pointer to **bool** | if downlink limiting for whole wlan is enabled | [optional] [default to false]
**WlanLimitUp** | Pointer to **NullableInt32** | kbps | [optional] 
**WlanLimitUpEnabled** | Pointer to **bool** | if uplink limiting for whole wlan is enabled | [optional] [default to false]
**WxtagIds** | Pointer to **[]string** | list of wxtag_ids | [optional] 
**WxtunnelId** | Pointer to **NullableString** | when &#x60;interface&#x60;&#x3D;&#x60;wxtunnel&#x60;, id of the WXLAN Tunnel | [optional] [default to ""]
**WxtunnelRemoteId** | Pointer to **NullableString** | when &#x60;interface&#x60;&#x3D;&#x60;wxtunnel&#x60;, remote tunnel identifier | [optional] [default to ""]

## Methods

### NewWlan

`func NewWlan(ssid string, ) *Wlan`

NewWlan instantiates a new Wlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanWithDefaults

`func NewWlanWithDefaults() *Wlan`

NewWlanWithDefaults instantiates a new Wlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctImmediateUpdate

`func (o *Wlan) GetAcctImmediateUpdate() bool`

GetAcctImmediateUpdate returns the AcctImmediateUpdate field if non-nil, zero value otherwise.

### GetAcctImmediateUpdateOk

`func (o *Wlan) GetAcctImmediateUpdateOk() (*bool, bool)`

GetAcctImmediateUpdateOk returns a tuple with the AcctImmediateUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctImmediateUpdate

`func (o *Wlan) SetAcctImmediateUpdate(v bool)`

SetAcctImmediateUpdate sets AcctImmediateUpdate field to given value.

### HasAcctImmediateUpdate

`func (o *Wlan) HasAcctImmediateUpdate() bool`

HasAcctImmediateUpdate returns a boolean if a field has been set.

### GetAcctInterimInterval

`func (o *Wlan) GetAcctInterimInterval() int32`

GetAcctInterimInterval returns the AcctInterimInterval field if non-nil, zero value otherwise.

### GetAcctInterimIntervalOk

`func (o *Wlan) GetAcctInterimIntervalOk() (*int32, bool)`

GetAcctInterimIntervalOk returns a tuple with the AcctInterimInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctInterimInterval

`func (o *Wlan) SetAcctInterimInterval(v int32)`

SetAcctInterimInterval sets AcctInterimInterval field to given value.

### HasAcctInterimInterval

`func (o *Wlan) HasAcctInterimInterval() bool`

HasAcctInterimInterval returns a boolean if a field has been set.

### GetAcctServers

`func (o *Wlan) GetAcctServers() []RadiusAcctServer`

GetAcctServers returns the AcctServers field if non-nil, zero value otherwise.

### GetAcctServersOk

`func (o *Wlan) GetAcctServersOk() (*[]RadiusAcctServer, bool)`

GetAcctServersOk returns a tuple with the AcctServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctServers

`func (o *Wlan) SetAcctServers(v []RadiusAcctServer)`

SetAcctServers sets AcctServers field to given value.

### HasAcctServers

`func (o *Wlan) HasAcctServers() bool`

HasAcctServers returns a boolean if a field has been set.

### GetAirwatch

`func (o *Wlan) GetAirwatch() WlanAirwatch`

GetAirwatch returns the Airwatch field if non-nil, zero value otherwise.

### GetAirwatchOk

`func (o *Wlan) GetAirwatchOk() (*WlanAirwatch, bool)`

GetAirwatchOk returns a tuple with the Airwatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirwatch

`func (o *Wlan) SetAirwatch(v WlanAirwatch)`

SetAirwatch sets Airwatch field to given value.

### HasAirwatch

`func (o *Wlan) HasAirwatch() bool`

HasAirwatch returns a boolean if a field has been set.

### GetAllowIpv6Ndp

`func (o *Wlan) GetAllowIpv6Ndp() bool`

GetAllowIpv6Ndp returns the AllowIpv6Ndp field if non-nil, zero value otherwise.

### GetAllowIpv6NdpOk

`func (o *Wlan) GetAllowIpv6NdpOk() (*bool, bool)`

GetAllowIpv6NdpOk returns a tuple with the AllowIpv6Ndp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIpv6Ndp

`func (o *Wlan) SetAllowIpv6Ndp(v bool)`

SetAllowIpv6Ndp sets AllowIpv6Ndp field to given value.

### HasAllowIpv6Ndp

`func (o *Wlan) HasAllowIpv6Ndp() bool`

HasAllowIpv6Ndp returns a boolean if a field has been set.

### GetAllowMdns

`func (o *Wlan) GetAllowMdns() bool`

GetAllowMdns returns the AllowMdns field if non-nil, zero value otherwise.

### GetAllowMdnsOk

`func (o *Wlan) GetAllowMdnsOk() (*bool, bool)`

GetAllowMdnsOk returns a tuple with the AllowMdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMdns

`func (o *Wlan) SetAllowMdns(v bool)`

SetAllowMdns sets AllowMdns field to given value.

### HasAllowMdns

`func (o *Wlan) HasAllowMdns() bool`

HasAllowMdns returns a boolean if a field has been set.

### GetAllowSsdp

`func (o *Wlan) GetAllowSsdp() bool`

GetAllowSsdp returns the AllowSsdp field if non-nil, zero value otherwise.

### GetAllowSsdpOk

`func (o *Wlan) GetAllowSsdpOk() (*bool, bool)`

GetAllowSsdpOk returns a tuple with the AllowSsdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSsdp

`func (o *Wlan) SetAllowSsdp(v bool)`

SetAllowSsdp sets AllowSsdp field to given value.

### HasAllowSsdp

`func (o *Wlan) HasAllowSsdp() bool`

HasAllowSsdp returns a boolean if a field has been set.

### GetApIds

`func (o *Wlan) GetApIds() []string`

GetApIds returns the ApIds field if non-nil, zero value otherwise.

### GetApIdsOk

`func (o *Wlan) GetApIdsOk() (*[]string, bool)`

GetApIdsOk returns a tuple with the ApIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApIds

`func (o *Wlan) SetApIds(v []string)`

SetApIds sets ApIds field to given value.

### HasApIds

`func (o *Wlan) HasApIds() bool`

HasApIds returns a boolean if a field has been set.

### SetApIdsNil

`func (o *Wlan) SetApIdsNil(b bool)`

 SetApIdsNil sets the value for ApIds to be an explicit nil

### UnsetApIds
`func (o *Wlan) UnsetApIds()`

UnsetApIds ensures that no value is present for ApIds, not even an explicit nil
### GetAppLimit

`func (o *Wlan) GetAppLimit() WlanAppLimit`

GetAppLimit returns the AppLimit field if non-nil, zero value otherwise.

### GetAppLimitOk

`func (o *Wlan) GetAppLimitOk() (*WlanAppLimit, bool)`

GetAppLimitOk returns a tuple with the AppLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLimit

`func (o *Wlan) SetAppLimit(v WlanAppLimit)`

SetAppLimit sets AppLimit field to given value.

### HasAppLimit

`func (o *Wlan) HasAppLimit() bool`

HasAppLimit returns a boolean if a field has been set.

### GetAppQos

`func (o *Wlan) GetAppQos() WlanAppQos`

GetAppQos returns the AppQos field if non-nil, zero value otherwise.

### GetAppQosOk

`func (o *Wlan) GetAppQosOk() (*WlanAppQos, bool)`

GetAppQosOk returns a tuple with the AppQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppQos

`func (o *Wlan) SetAppQos(v WlanAppQos)`

SetAppQos sets AppQos field to given value.

### HasAppQos

`func (o *Wlan) HasAppQos() bool`

HasAppQos returns a boolean if a field has been set.

### GetApplyTo

`func (o *Wlan) GetApplyTo() WlanApplyTo`

GetApplyTo returns the ApplyTo field if non-nil, zero value otherwise.

### GetApplyToOk

`func (o *Wlan) GetApplyToOk() (*WlanApplyTo, bool)`

GetApplyToOk returns a tuple with the ApplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTo

`func (o *Wlan) SetApplyTo(v WlanApplyTo)`

SetApplyTo sets ApplyTo field to given value.

### HasApplyTo

`func (o *Wlan) HasApplyTo() bool`

HasApplyTo returns a boolean if a field has been set.

### GetArpFilter

`func (o *Wlan) GetArpFilter() bool`

GetArpFilter returns the ArpFilter field if non-nil, zero value otherwise.

### GetArpFilterOk

`func (o *Wlan) GetArpFilterOk() (*bool, bool)`

GetArpFilterOk returns a tuple with the ArpFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpFilter

`func (o *Wlan) SetArpFilter(v bool)`

SetArpFilter sets ArpFilter field to given value.

### HasArpFilter

`func (o *Wlan) HasArpFilter() bool`

HasArpFilter returns a boolean if a field has been set.

### GetAuth

`func (o *Wlan) GetAuth() WlanAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Wlan) GetAuthOk() (*WlanAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Wlan) SetAuth(v WlanAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *Wlan) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthServerSelection

`func (o *Wlan) GetAuthServerSelection() WlanAuthServerSelection`

GetAuthServerSelection returns the AuthServerSelection field if non-nil, zero value otherwise.

### GetAuthServerSelectionOk

`func (o *Wlan) GetAuthServerSelectionOk() (*WlanAuthServerSelection, bool)`

GetAuthServerSelectionOk returns a tuple with the AuthServerSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServerSelection

`func (o *Wlan) SetAuthServerSelection(v WlanAuthServerSelection)`

SetAuthServerSelection sets AuthServerSelection field to given value.

### HasAuthServerSelection

`func (o *Wlan) HasAuthServerSelection() bool`

HasAuthServerSelection returns a boolean if a field has been set.

### GetAuthServers

`func (o *Wlan) GetAuthServers() []RadiusAuthServer`

GetAuthServers returns the AuthServers field if non-nil, zero value otherwise.

### GetAuthServersOk

`func (o *Wlan) GetAuthServersOk() (*[]RadiusAuthServer, bool)`

GetAuthServersOk returns a tuple with the AuthServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServers

`func (o *Wlan) SetAuthServers(v []RadiusAuthServer)`

SetAuthServers sets AuthServers field to given value.

### HasAuthServers

`func (o *Wlan) HasAuthServers() bool`

HasAuthServers returns a boolean if a field has been set.

### GetAuthServersNasId

`func (o *Wlan) GetAuthServersNasId() string`

GetAuthServersNasId returns the AuthServersNasId field if non-nil, zero value otherwise.

### GetAuthServersNasIdOk

`func (o *Wlan) GetAuthServersNasIdOk() (*string, bool)`

GetAuthServersNasIdOk returns a tuple with the AuthServersNasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServersNasId

`func (o *Wlan) SetAuthServersNasId(v string)`

SetAuthServersNasId sets AuthServersNasId field to given value.

### HasAuthServersNasId

`func (o *Wlan) HasAuthServersNasId() bool`

HasAuthServersNasId returns a boolean if a field has been set.

### SetAuthServersNasIdNil

`func (o *Wlan) SetAuthServersNasIdNil(b bool)`

 SetAuthServersNasIdNil sets the value for AuthServersNasId to be an explicit nil

### UnsetAuthServersNasId
`func (o *Wlan) UnsetAuthServersNasId()`

UnsetAuthServersNasId ensures that no value is present for AuthServersNasId, not even an explicit nil
### GetAuthServersNasIp

`func (o *Wlan) GetAuthServersNasIp() string`

GetAuthServersNasIp returns the AuthServersNasIp field if non-nil, zero value otherwise.

### GetAuthServersNasIpOk

`func (o *Wlan) GetAuthServersNasIpOk() (*string, bool)`

GetAuthServersNasIpOk returns a tuple with the AuthServersNasIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServersNasIp

`func (o *Wlan) SetAuthServersNasIp(v string)`

SetAuthServersNasIp sets AuthServersNasIp field to given value.

### HasAuthServersNasIp

`func (o *Wlan) HasAuthServersNasIp() bool`

HasAuthServersNasIp returns a boolean if a field has been set.

### SetAuthServersNasIpNil

`func (o *Wlan) SetAuthServersNasIpNil(b bool)`

 SetAuthServersNasIpNil sets the value for AuthServersNasIp to be an explicit nil

### UnsetAuthServersNasIp
`func (o *Wlan) UnsetAuthServersNasIp()`

UnsetAuthServersNasIp ensures that no value is present for AuthServersNasIp, not even an explicit nil
### GetAuthServersRetries

`func (o *Wlan) GetAuthServersRetries() int32`

GetAuthServersRetries returns the AuthServersRetries field if non-nil, zero value otherwise.

### GetAuthServersRetriesOk

`func (o *Wlan) GetAuthServersRetriesOk() (*int32, bool)`

GetAuthServersRetriesOk returns a tuple with the AuthServersRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServersRetries

`func (o *Wlan) SetAuthServersRetries(v int32)`

SetAuthServersRetries sets AuthServersRetries field to given value.

### HasAuthServersRetries

`func (o *Wlan) HasAuthServersRetries() bool`

HasAuthServersRetries returns a boolean if a field has been set.

### GetAuthServersTimeout

`func (o *Wlan) GetAuthServersTimeout() int32`

GetAuthServersTimeout returns the AuthServersTimeout field if non-nil, zero value otherwise.

### GetAuthServersTimeoutOk

`func (o *Wlan) GetAuthServersTimeoutOk() (*int32, bool)`

GetAuthServersTimeoutOk returns a tuple with the AuthServersTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServersTimeout

`func (o *Wlan) SetAuthServersTimeout(v int32)`

SetAuthServersTimeout sets AuthServersTimeout field to given value.

### HasAuthServersTimeout

`func (o *Wlan) HasAuthServersTimeout() bool`

HasAuthServersTimeout returns a boolean if a field has been set.

### GetBand

`func (o *Wlan) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *Wlan) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *Wlan) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *Wlan) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetBandSteer

`func (o *Wlan) GetBandSteer() bool`

GetBandSteer returns the BandSteer field if non-nil, zero value otherwise.

### GetBandSteerOk

`func (o *Wlan) GetBandSteerOk() (*bool, bool)`

GetBandSteerOk returns a tuple with the BandSteer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteer

`func (o *Wlan) SetBandSteer(v bool)`

SetBandSteer sets BandSteer field to given value.

### HasBandSteer

`func (o *Wlan) HasBandSteer() bool`

HasBandSteer returns a boolean if a field has been set.

### GetBandSteerForceBand5

`func (o *Wlan) GetBandSteerForceBand5() bool`

GetBandSteerForceBand5 returns the BandSteerForceBand5 field if non-nil, zero value otherwise.

### GetBandSteerForceBand5Ok

`func (o *Wlan) GetBandSteerForceBand5Ok() (*bool, bool)`

GetBandSteerForceBand5Ok returns a tuple with the BandSteerForceBand5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteerForceBand5

`func (o *Wlan) SetBandSteerForceBand5(v bool)`

SetBandSteerForceBand5 sets BandSteerForceBand5 field to given value.

### HasBandSteerForceBand5

`func (o *Wlan) HasBandSteerForceBand5() bool`

HasBandSteerForceBand5 returns a boolean if a field has been set.

### GetBands

`func (o *Wlan) GetBands() []Dot11Band`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *Wlan) GetBandsOk() (*[]Dot11Band, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *Wlan) SetBands(v []Dot11Band)`

SetBands sets Bands field to given value.

### HasBands

`func (o *Wlan) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetBlockBlacklistClients

`func (o *Wlan) GetBlockBlacklistClients() bool`

GetBlockBlacklistClients returns the BlockBlacklistClients field if non-nil, zero value otherwise.

### GetBlockBlacklistClientsOk

`func (o *Wlan) GetBlockBlacklistClientsOk() (*bool, bool)`

GetBlockBlacklistClientsOk returns a tuple with the BlockBlacklistClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockBlacklistClients

`func (o *Wlan) SetBlockBlacklistClients(v bool)`

SetBlockBlacklistClients sets BlockBlacklistClients field to given value.

### HasBlockBlacklistClients

`func (o *Wlan) HasBlockBlacklistClients() bool`

HasBlockBlacklistClients returns a boolean if a field has been set.

### GetBonjour

`func (o *Wlan) GetBonjour() WlanBonjour`

GetBonjour returns the Bonjour field if non-nil, zero value otherwise.

### GetBonjourOk

`func (o *Wlan) GetBonjourOk() (*WlanBonjour, bool)`

GetBonjourOk returns a tuple with the Bonjour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjour

`func (o *Wlan) SetBonjour(v WlanBonjour)`

SetBonjour sets Bonjour field to given value.

### HasBonjour

`func (o *Wlan) HasBonjour() bool`

HasBonjour returns a boolean if a field has been set.

### GetCiscoCwa

`func (o *Wlan) GetCiscoCwa() WlanCiscoCwa`

GetCiscoCwa returns the CiscoCwa field if non-nil, zero value otherwise.

### GetCiscoCwaOk

`func (o *Wlan) GetCiscoCwaOk() (*WlanCiscoCwa, bool)`

GetCiscoCwaOk returns a tuple with the CiscoCwa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoCwa

`func (o *Wlan) SetCiscoCwa(v WlanCiscoCwa)`

SetCiscoCwa sets CiscoCwa field to given value.

### HasCiscoCwa

`func (o *Wlan) HasCiscoCwa() bool`

HasCiscoCwa returns a boolean if a field has been set.

### GetClientLimitDown

`func (o *Wlan) GetClientLimitDown() int32`

GetClientLimitDown returns the ClientLimitDown field if non-nil, zero value otherwise.

### GetClientLimitDownOk

`func (o *Wlan) GetClientLimitDownOk() (*int32, bool)`

GetClientLimitDownOk returns a tuple with the ClientLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientLimitDown

`func (o *Wlan) SetClientLimitDown(v int32)`

SetClientLimitDown sets ClientLimitDown field to given value.

### HasClientLimitDown

`func (o *Wlan) HasClientLimitDown() bool`

HasClientLimitDown returns a boolean if a field has been set.

### GetClientLimitDownEnabled

`func (o *Wlan) GetClientLimitDownEnabled() bool`

GetClientLimitDownEnabled returns the ClientLimitDownEnabled field if non-nil, zero value otherwise.

### GetClientLimitDownEnabledOk

`func (o *Wlan) GetClientLimitDownEnabledOk() (*bool, bool)`

GetClientLimitDownEnabledOk returns a tuple with the ClientLimitDownEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientLimitDownEnabled

`func (o *Wlan) SetClientLimitDownEnabled(v bool)`

SetClientLimitDownEnabled sets ClientLimitDownEnabled field to given value.

### HasClientLimitDownEnabled

`func (o *Wlan) HasClientLimitDownEnabled() bool`

HasClientLimitDownEnabled returns a boolean if a field has been set.

### GetClientLimitUp

`func (o *Wlan) GetClientLimitUp() int32`

GetClientLimitUp returns the ClientLimitUp field if non-nil, zero value otherwise.

### GetClientLimitUpOk

`func (o *Wlan) GetClientLimitUpOk() (*int32, bool)`

GetClientLimitUpOk returns a tuple with the ClientLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientLimitUp

`func (o *Wlan) SetClientLimitUp(v int32)`

SetClientLimitUp sets ClientLimitUp field to given value.

### HasClientLimitUp

`func (o *Wlan) HasClientLimitUp() bool`

HasClientLimitUp returns a boolean if a field has been set.

### GetClientLimitUpEnabled

`func (o *Wlan) GetClientLimitUpEnabled() bool`

GetClientLimitUpEnabled returns the ClientLimitUpEnabled field if non-nil, zero value otherwise.

### GetClientLimitUpEnabledOk

`func (o *Wlan) GetClientLimitUpEnabledOk() (*bool, bool)`

GetClientLimitUpEnabledOk returns a tuple with the ClientLimitUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientLimitUpEnabled

`func (o *Wlan) SetClientLimitUpEnabled(v bool)`

SetClientLimitUpEnabled sets ClientLimitUpEnabled field to given value.

### HasClientLimitUpEnabled

`func (o *Wlan) HasClientLimitUpEnabled() bool`

HasClientLimitUpEnabled returns a boolean if a field has been set.

### GetCoaServers

`func (o *Wlan) GetCoaServers() []CoaServer`

GetCoaServers returns the CoaServers field if non-nil, zero value otherwise.

### GetCoaServersOk

`func (o *Wlan) GetCoaServersOk() (*[]CoaServer, bool)`

GetCoaServersOk returns a tuple with the CoaServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaServers

`func (o *Wlan) SetCoaServers(v []CoaServer)`

SetCoaServers sets CoaServers field to given value.

### HasCoaServers

`func (o *Wlan) HasCoaServers() bool`

HasCoaServers returns a boolean if a field has been set.

### SetCoaServersNil

`func (o *Wlan) SetCoaServersNil(b bool)`

 SetCoaServersNil sets the value for CoaServers to be an explicit nil

### UnsetCoaServers
`func (o *Wlan) UnsetCoaServers()`

UnsetCoaServers ensures that no value is present for CoaServers, not even an explicit nil
### GetCreatedTime

`func (o *Wlan) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Wlan) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Wlan) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Wlan) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisable11ax

`func (o *Wlan) GetDisable11ax() bool`

GetDisable11ax returns the Disable11ax field if non-nil, zero value otherwise.

### GetDisable11axOk

`func (o *Wlan) GetDisable11axOk() (*bool, bool)`

GetDisable11axOk returns a tuple with the Disable11ax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable11ax

`func (o *Wlan) SetDisable11ax(v bool)`

SetDisable11ax sets Disable11ax field to given value.

### HasDisable11ax

`func (o *Wlan) HasDisable11ax() bool`

HasDisable11ax returns a boolean if a field has been set.

### GetDisableHtVhtRates

`func (o *Wlan) GetDisableHtVhtRates() bool`

GetDisableHtVhtRates returns the DisableHtVhtRates field if non-nil, zero value otherwise.

### GetDisableHtVhtRatesOk

`func (o *Wlan) GetDisableHtVhtRatesOk() (*bool, bool)`

GetDisableHtVhtRatesOk returns a tuple with the DisableHtVhtRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHtVhtRates

`func (o *Wlan) SetDisableHtVhtRates(v bool)`

SetDisableHtVhtRates sets DisableHtVhtRates field to given value.

### HasDisableHtVhtRates

`func (o *Wlan) HasDisableHtVhtRates() bool`

HasDisableHtVhtRates returns a boolean if a field has been set.

### GetDisableUapsd

`func (o *Wlan) GetDisableUapsd() bool`

GetDisableUapsd returns the DisableUapsd field if non-nil, zero value otherwise.

### GetDisableUapsdOk

`func (o *Wlan) GetDisableUapsdOk() (*bool, bool)`

GetDisableUapsdOk returns a tuple with the DisableUapsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUapsd

`func (o *Wlan) SetDisableUapsd(v bool)`

SetDisableUapsd sets DisableUapsd field to given value.

### HasDisableUapsd

`func (o *Wlan) HasDisableUapsd() bool`

HasDisableUapsd returns a boolean if a field has been set.

### GetDisableV1RoamNotify

`func (o *Wlan) GetDisableV1RoamNotify() bool`

GetDisableV1RoamNotify returns the DisableV1RoamNotify field if non-nil, zero value otherwise.

### GetDisableV1RoamNotifyOk

`func (o *Wlan) GetDisableV1RoamNotifyOk() (*bool, bool)`

GetDisableV1RoamNotifyOk returns a tuple with the DisableV1RoamNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableV1RoamNotify

`func (o *Wlan) SetDisableV1RoamNotify(v bool)`

SetDisableV1RoamNotify sets DisableV1RoamNotify field to given value.

### HasDisableV1RoamNotify

`func (o *Wlan) HasDisableV1RoamNotify() bool`

HasDisableV1RoamNotify returns a boolean if a field has been set.

### GetDisableV2RoamNotify

`func (o *Wlan) GetDisableV2RoamNotify() bool`

GetDisableV2RoamNotify returns the DisableV2RoamNotify field if non-nil, zero value otherwise.

### GetDisableV2RoamNotifyOk

`func (o *Wlan) GetDisableV2RoamNotifyOk() (*bool, bool)`

GetDisableV2RoamNotifyOk returns a tuple with the DisableV2RoamNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableV2RoamNotify

`func (o *Wlan) SetDisableV2RoamNotify(v bool)`

SetDisableV2RoamNotify sets DisableV2RoamNotify field to given value.

### HasDisableV2RoamNotify

`func (o *Wlan) HasDisableV2RoamNotify() bool`

HasDisableV2RoamNotify returns a boolean if a field has been set.

### GetDisableWmm

`func (o *Wlan) GetDisableWmm() bool`

GetDisableWmm returns the DisableWmm field if non-nil, zero value otherwise.

### GetDisableWmmOk

`func (o *Wlan) GetDisableWmmOk() (*bool, bool)`

GetDisableWmmOk returns a tuple with the DisableWmm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableWmm

`func (o *Wlan) SetDisableWmm(v bool)`

SetDisableWmm sets DisableWmm field to given value.

### HasDisableWmm

`func (o *Wlan) HasDisableWmm() bool`

HasDisableWmm returns a boolean if a field has been set.

### GetDnsServerRewrite

`func (o *Wlan) GetDnsServerRewrite() WlanDnsServerRewrite`

GetDnsServerRewrite returns the DnsServerRewrite field if non-nil, zero value otherwise.

### GetDnsServerRewriteOk

`func (o *Wlan) GetDnsServerRewriteOk() (*WlanDnsServerRewrite, bool)`

GetDnsServerRewriteOk returns a tuple with the DnsServerRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerRewrite

`func (o *Wlan) SetDnsServerRewrite(v WlanDnsServerRewrite)`

SetDnsServerRewrite sets DnsServerRewrite field to given value.

### HasDnsServerRewrite

`func (o *Wlan) HasDnsServerRewrite() bool`

HasDnsServerRewrite returns a boolean if a field has been set.

### SetDnsServerRewriteNil

`func (o *Wlan) SetDnsServerRewriteNil(b bool)`

 SetDnsServerRewriteNil sets the value for DnsServerRewrite to be an explicit nil

### UnsetDnsServerRewrite
`func (o *Wlan) UnsetDnsServerRewrite()`

UnsetDnsServerRewrite ensures that no value is present for DnsServerRewrite, not even an explicit nil
### GetDtim

`func (o *Wlan) GetDtim() int32`

GetDtim returns the Dtim field if non-nil, zero value otherwise.

### GetDtimOk

`func (o *Wlan) GetDtimOk() (*int32, bool)`

GetDtimOk returns a tuple with the Dtim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtim

`func (o *Wlan) SetDtim(v int32)`

SetDtim sets Dtim field to given value.

### HasDtim

`func (o *Wlan) HasDtim() bool`

HasDtim returns a boolean if a field has been set.

### GetDynamicPsk

`func (o *Wlan) GetDynamicPsk() WlanDynamicPsk`

GetDynamicPsk returns the DynamicPsk field if non-nil, zero value otherwise.

### GetDynamicPskOk

`func (o *Wlan) GetDynamicPskOk() (*WlanDynamicPsk, bool)`

GetDynamicPskOk returns a tuple with the DynamicPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPsk

`func (o *Wlan) SetDynamicPsk(v WlanDynamicPsk)`

SetDynamicPsk sets DynamicPsk field to given value.

### HasDynamicPsk

`func (o *Wlan) HasDynamicPsk() bool`

HasDynamicPsk returns a boolean if a field has been set.

### SetDynamicPskNil

`func (o *Wlan) SetDynamicPskNil(b bool)`

 SetDynamicPskNil sets the value for DynamicPsk to be an explicit nil

### UnsetDynamicPsk
`func (o *Wlan) UnsetDynamicPsk()`

UnsetDynamicPsk ensures that no value is present for DynamicPsk, not even an explicit nil
### GetDynamicVlan

`func (o *Wlan) GetDynamicVlan() WlanDynamicVlan`

GetDynamicVlan returns the DynamicVlan field if non-nil, zero value otherwise.

### GetDynamicVlanOk

`func (o *Wlan) GetDynamicVlanOk() (*WlanDynamicVlan, bool)`

GetDynamicVlanOk returns a tuple with the DynamicVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVlan

`func (o *Wlan) SetDynamicVlan(v WlanDynamicVlan)`

SetDynamicVlan sets DynamicVlan field to given value.

### HasDynamicVlan

`func (o *Wlan) HasDynamicVlan() bool`

HasDynamicVlan returns a boolean if a field has been set.

### SetDynamicVlanNil

`func (o *Wlan) SetDynamicVlanNil(b bool)`

 SetDynamicVlanNil sets the value for DynamicVlan to be an explicit nil

### UnsetDynamicVlan
`func (o *Wlan) UnsetDynamicVlan()`

UnsetDynamicVlan ensures that no value is present for DynamicVlan, not even an explicit nil
### GetEnableLocalKeycaching

`func (o *Wlan) GetEnableLocalKeycaching() bool`

GetEnableLocalKeycaching returns the EnableLocalKeycaching field if non-nil, zero value otherwise.

### GetEnableLocalKeycachingOk

`func (o *Wlan) GetEnableLocalKeycachingOk() (*bool, bool)`

GetEnableLocalKeycachingOk returns a tuple with the EnableLocalKeycaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalKeycaching

`func (o *Wlan) SetEnableLocalKeycaching(v bool)`

SetEnableLocalKeycaching sets EnableLocalKeycaching field to given value.

### HasEnableLocalKeycaching

`func (o *Wlan) HasEnableLocalKeycaching() bool`

HasEnableLocalKeycaching returns a boolean if a field has been set.

### GetEnableWirelessBridging

`func (o *Wlan) GetEnableWirelessBridging() bool`

GetEnableWirelessBridging returns the EnableWirelessBridging field if non-nil, zero value otherwise.

### GetEnableWirelessBridgingOk

`func (o *Wlan) GetEnableWirelessBridgingOk() (*bool, bool)`

GetEnableWirelessBridgingOk returns a tuple with the EnableWirelessBridging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWirelessBridging

`func (o *Wlan) SetEnableWirelessBridging(v bool)`

SetEnableWirelessBridging sets EnableWirelessBridging field to given value.

### HasEnableWirelessBridging

`func (o *Wlan) HasEnableWirelessBridging() bool`

HasEnableWirelessBridging returns a boolean if a field has been set.

### GetEnableWirelessBridgingDhcpTracking

`func (o *Wlan) GetEnableWirelessBridgingDhcpTracking() bool`

GetEnableWirelessBridgingDhcpTracking returns the EnableWirelessBridgingDhcpTracking field if non-nil, zero value otherwise.

### GetEnableWirelessBridgingDhcpTrackingOk

`func (o *Wlan) GetEnableWirelessBridgingDhcpTrackingOk() (*bool, bool)`

GetEnableWirelessBridgingDhcpTrackingOk returns a tuple with the EnableWirelessBridgingDhcpTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWirelessBridgingDhcpTracking

`func (o *Wlan) SetEnableWirelessBridgingDhcpTracking(v bool)`

SetEnableWirelessBridgingDhcpTracking sets EnableWirelessBridgingDhcpTracking field to given value.

### HasEnableWirelessBridgingDhcpTracking

`func (o *Wlan) HasEnableWirelessBridgingDhcpTracking() bool`

HasEnableWirelessBridgingDhcpTracking returns a boolean if a field has been set.

### GetEnabled

`func (o *Wlan) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Wlan) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Wlan) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Wlan) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFastDot1xTimers

`func (o *Wlan) GetFastDot1xTimers() bool`

GetFastDot1xTimers returns the FastDot1xTimers field if non-nil, zero value otherwise.

### GetFastDot1xTimersOk

`func (o *Wlan) GetFastDot1xTimersOk() (*bool, bool)`

GetFastDot1xTimersOk returns a tuple with the FastDot1xTimers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastDot1xTimers

`func (o *Wlan) SetFastDot1xTimers(v bool)`

SetFastDot1xTimers sets FastDot1xTimers field to given value.

### HasFastDot1xTimers

`func (o *Wlan) HasFastDot1xTimers() bool`

HasFastDot1xTimers returns a boolean if a field has been set.

### GetForSite

`func (o *Wlan) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Wlan) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Wlan) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Wlan) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHideSsid

`func (o *Wlan) GetHideSsid() bool`

GetHideSsid returns the HideSsid field if non-nil, zero value otherwise.

### GetHideSsidOk

`func (o *Wlan) GetHideSsidOk() (*bool, bool)`

GetHideSsidOk returns a tuple with the HideSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideSsid

`func (o *Wlan) SetHideSsid(v bool)`

SetHideSsid sets HideSsid field to given value.

### HasHideSsid

`func (o *Wlan) HasHideSsid() bool`

HasHideSsid returns a boolean if a field has been set.

### GetHostnameIe

`func (o *Wlan) GetHostnameIe() bool`

GetHostnameIe returns the HostnameIe field if non-nil, zero value otherwise.

### GetHostnameIeOk

`func (o *Wlan) GetHostnameIeOk() (*bool, bool)`

GetHostnameIeOk returns a tuple with the HostnameIe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameIe

`func (o *Wlan) SetHostnameIe(v bool)`

SetHostnameIe sets HostnameIe field to given value.

### HasHostnameIe

`func (o *Wlan) HasHostnameIe() bool`

HasHostnameIe returns a boolean if a field has been set.

### GetHotspot20

`func (o *Wlan) GetHotspot20() WlanHotspot20`

GetHotspot20 returns the Hotspot20 field if non-nil, zero value otherwise.

### GetHotspot20Ok

`func (o *Wlan) GetHotspot20Ok() (*WlanHotspot20, bool)`

GetHotspot20Ok returns a tuple with the Hotspot20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspot20

`func (o *Wlan) SetHotspot20(v WlanHotspot20)`

SetHotspot20 sets Hotspot20 field to given value.

### HasHotspot20

`func (o *Wlan) HasHotspot20() bool`

HasHotspot20 returns a boolean if a field has been set.

### GetId

`func (o *Wlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Wlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Wlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Wlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInjectDhcpOption82

`func (o *Wlan) GetInjectDhcpOption82() WlanInjectDhcpOption82`

GetInjectDhcpOption82 returns the InjectDhcpOption82 field if non-nil, zero value otherwise.

### GetInjectDhcpOption82Ok

`func (o *Wlan) GetInjectDhcpOption82Ok() (*WlanInjectDhcpOption82, bool)`

GetInjectDhcpOption82Ok returns a tuple with the InjectDhcpOption82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectDhcpOption82

`func (o *Wlan) SetInjectDhcpOption82(v WlanInjectDhcpOption82)`

SetInjectDhcpOption82 sets InjectDhcpOption82 field to given value.

### HasInjectDhcpOption82

`func (o *Wlan) HasInjectDhcpOption82() bool`

HasInjectDhcpOption82 returns a boolean if a field has been set.

### GetInterface

`func (o *Wlan) GetInterface() WlanInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *Wlan) GetInterfaceOk() (*WlanInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *Wlan) SetInterface(v WlanInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *Wlan) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIsolation

`func (o *Wlan) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *Wlan) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *Wlan) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *Wlan) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetL2Isolation

`func (o *Wlan) GetL2Isolation() bool`

GetL2Isolation returns the L2Isolation field if non-nil, zero value otherwise.

### GetL2IsolationOk

`func (o *Wlan) GetL2IsolationOk() (*bool, bool)`

GetL2IsolationOk returns a tuple with the L2Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Isolation

`func (o *Wlan) SetL2Isolation(v bool)`

SetL2Isolation sets L2Isolation field to given value.

### HasL2Isolation

`func (o *Wlan) HasL2Isolation() bool`

HasL2Isolation returns a boolean if a field has been set.

### GetLegacyOverds

`func (o *Wlan) GetLegacyOverds() bool`

GetLegacyOverds returns the LegacyOverds field if non-nil, zero value otherwise.

### GetLegacyOverdsOk

`func (o *Wlan) GetLegacyOverdsOk() (*bool, bool)`

GetLegacyOverdsOk returns a tuple with the LegacyOverds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyOverds

`func (o *Wlan) SetLegacyOverds(v bool)`

SetLegacyOverds sets LegacyOverds field to given value.

### HasLegacyOverds

`func (o *Wlan) HasLegacyOverds() bool`

HasLegacyOverds returns a boolean if a field has been set.

### GetLimitBcast

`func (o *Wlan) GetLimitBcast() bool`

GetLimitBcast returns the LimitBcast field if non-nil, zero value otherwise.

### GetLimitBcastOk

`func (o *Wlan) GetLimitBcastOk() (*bool, bool)`

GetLimitBcastOk returns a tuple with the LimitBcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBcast

`func (o *Wlan) SetLimitBcast(v bool)`

SetLimitBcast sets LimitBcast field to given value.

### HasLimitBcast

`func (o *Wlan) HasLimitBcast() bool`

HasLimitBcast returns a boolean if a field has been set.

### GetLimitProbeResponse

`func (o *Wlan) GetLimitProbeResponse() bool`

GetLimitProbeResponse returns the LimitProbeResponse field if non-nil, zero value otherwise.

### GetLimitProbeResponseOk

`func (o *Wlan) GetLimitProbeResponseOk() (*bool, bool)`

GetLimitProbeResponseOk returns a tuple with the LimitProbeResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitProbeResponse

`func (o *Wlan) SetLimitProbeResponse(v bool)`

SetLimitProbeResponse sets LimitProbeResponse field to given value.

### HasLimitProbeResponse

`func (o *Wlan) HasLimitProbeResponse() bool`

HasLimitProbeResponse returns a boolean if a field has been set.

### GetMaxIdletime

`func (o *Wlan) GetMaxIdletime() int32`

GetMaxIdletime returns the MaxIdletime field if non-nil, zero value otherwise.

### GetMaxIdletimeOk

`func (o *Wlan) GetMaxIdletimeOk() (*int32, bool)`

GetMaxIdletimeOk returns a tuple with the MaxIdletime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdletime

`func (o *Wlan) SetMaxIdletime(v int32)`

SetMaxIdletime sets MaxIdletime field to given value.

### HasMaxIdletime

`func (o *Wlan) HasMaxIdletime() bool`

HasMaxIdletime returns a boolean if a field has been set.

### GetMistNac

`func (o *Wlan) GetMistNac() WlanMistNac`

GetMistNac returns the MistNac field if non-nil, zero value otherwise.

### GetMistNacOk

`func (o *Wlan) GetMistNacOk() (*WlanMistNac, bool)`

GetMistNacOk returns a tuple with the MistNac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistNac

`func (o *Wlan) SetMistNac(v WlanMistNac)`

SetMistNac sets MistNac field to given value.

### HasMistNac

`func (o *Wlan) HasMistNac() bool`

HasMistNac returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Wlan) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Wlan) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Wlan) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Wlan) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMspId

`func (o *Wlan) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *Wlan) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *Wlan) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *Wlan) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetMxtunnelId

`func (o *Wlan) GetMxtunnelId() string`

GetMxtunnelId returns the MxtunnelId field if non-nil, zero value otherwise.

### GetMxtunnelIdOk

`func (o *Wlan) GetMxtunnelIdOk() (*string, bool)`

GetMxtunnelIdOk returns a tuple with the MxtunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxtunnelId

`func (o *Wlan) SetMxtunnelId(v string)`

SetMxtunnelId sets MxtunnelId field to given value.

### HasMxtunnelId

`func (o *Wlan) HasMxtunnelId() bool`

HasMxtunnelId returns a boolean if a field has been set.

### GetMxtunnelIds

`func (o *Wlan) GetMxtunnelIds() []string`

GetMxtunnelIds returns the MxtunnelIds field if non-nil, zero value otherwise.

### GetMxtunnelIdsOk

`func (o *Wlan) GetMxtunnelIdsOk() (*[]string, bool)`

GetMxtunnelIdsOk returns a tuple with the MxtunnelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxtunnelIds

`func (o *Wlan) SetMxtunnelIds(v []string)`

SetMxtunnelIds sets MxtunnelIds field to given value.

### HasMxtunnelIds

`func (o *Wlan) HasMxtunnelIds() bool`

HasMxtunnelIds returns a boolean if a field has been set.

### GetMxtunnelName

`func (o *Wlan) GetMxtunnelName() []string`

GetMxtunnelName returns the MxtunnelName field if non-nil, zero value otherwise.

### GetMxtunnelNameOk

`func (o *Wlan) GetMxtunnelNameOk() (*[]string, bool)`

GetMxtunnelNameOk returns a tuple with the MxtunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxtunnelName

`func (o *Wlan) SetMxtunnelName(v []string)`

SetMxtunnelName sets MxtunnelName field to given value.

### HasMxtunnelName

`func (o *Wlan) HasMxtunnelName() bool`

HasMxtunnelName returns a boolean if a field has been set.

### GetNoStaticDns

`func (o *Wlan) GetNoStaticDns() bool`

GetNoStaticDns returns the NoStaticDns field if non-nil, zero value otherwise.

### GetNoStaticDnsOk

`func (o *Wlan) GetNoStaticDnsOk() (*bool, bool)`

GetNoStaticDnsOk returns a tuple with the NoStaticDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoStaticDns

`func (o *Wlan) SetNoStaticDns(v bool)`

SetNoStaticDns sets NoStaticDns field to given value.

### HasNoStaticDns

`func (o *Wlan) HasNoStaticDns() bool`

HasNoStaticDns returns a boolean if a field has been set.

### GetNoStaticIp

`func (o *Wlan) GetNoStaticIp() bool`

GetNoStaticIp returns the NoStaticIp field if non-nil, zero value otherwise.

### GetNoStaticIpOk

`func (o *Wlan) GetNoStaticIpOk() (*bool, bool)`

GetNoStaticIpOk returns a tuple with the NoStaticIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoStaticIp

`func (o *Wlan) SetNoStaticIp(v bool)`

SetNoStaticIp sets NoStaticIp field to given value.

### HasNoStaticIp

`func (o *Wlan) HasNoStaticIp() bool`

HasNoStaticIp returns a boolean if a field has been set.

### GetOrgId

`func (o *Wlan) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Wlan) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Wlan) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Wlan) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPortal

`func (o *Wlan) GetPortal() WlanPortal`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *Wlan) GetPortalOk() (*WlanPortal, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *Wlan) SetPortal(v WlanPortal)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *Wlan) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetPortalAllowedHostnames

`func (o *Wlan) GetPortalAllowedHostnames() []string`

GetPortalAllowedHostnames returns the PortalAllowedHostnames field if non-nil, zero value otherwise.

### GetPortalAllowedHostnamesOk

`func (o *Wlan) GetPortalAllowedHostnamesOk() (*[]string, bool)`

GetPortalAllowedHostnamesOk returns a tuple with the PortalAllowedHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalAllowedHostnames

`func (o *Wlan) SetPortalAllowedHostnames(v []string)`

SetPortalAllowedHostnames sets PortalAllowedHostnames field to given value.

### HasPortalAllowedHostnames

`func (o *Wlan) HasPortalAllowedHostnames() bool`

HasPortalAllowedHostnames returns a boolean if a field has been set.

### GetPortalAllowedSubnets

`func (o *Wlan) GetPortalAllowedSubnets() []string`

GetPortalAllowedSubnets returns the PortalAllowedSubnets field if non-nil, zero value otherwise.

### GetPortalAllowedSubnetsOk

`func (o *Wlan) GetPortalAllowedSubnetsOk() (*[]string, bool)`

GetPortalAllowedSubnetsOk returns a tuple with the PortalAllowedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalAllowedSubnets

`func (o *Wlan) SetPortalAllowedSubnets(v []string)`

SetPortalAllowedSubnets sets PortalAllowedSubnets field to given value.

### HasPortalAllowedSubnets

`func (o *Wlan) HasPortalAllowedSubnets() bool`

HasPortalAllowedSubnets returns a boolean if a field has been set.

### GetPortalApiSecret

`func (o *Wlan) GetPortalApiSecret() string`

GetPortalApiSecret returns the PortalApiSecret field if non-nil, zero value otherwise.

### GetPortalApiSecretOk

`func (o *Wlan) GetPortalApiSecretOk() (*string, bool)`

GetPortalApiSecretOk returns a tuple with the PortalApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalApiSecret

`func (o *Wlan) SetPortalApiSecret(v string)`

SetPortalApiSecret sets PortalApiSecret field to given value.

### HasPortalApiSecret

`func (o *Wlan) HasPortalApiSecret() bool`

HasPortalApiSecret returns a boolean if a field has been set.

### SetPortalApiSecretNil

`func (o *Wlan) SetPortalApiSecretNil(b bool)`

 SetPortalApiSecretNil sets the value for PortalApiSecret to be an explicit nil

### UnsetPortalApiSecret
`func (o *Wlan) UnsetPortalApiSecret()`

UnsetPortalApiSecret ensures that no value is present for PortalApiSecret, not even an explicit nil
### GetPortalDeniedHostnames

`func (o *Wlan) GetPortalDeniedHostnames() []string`

GetPortalDeniedHostnames returns the PortalDeniedHostnames field if non-nil, zero value otherwise.

### GetPortalDeniedHostnamesOk

`func (o *Wlan) GetPortalDeniedHostnamesOk() (*[]string, bool)`

GetPortalDeniedHostnamesOk returns a tuple with the PortalDeniedHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalDeniedHostnames

`func (o *Wlan) SetPortalDeniedHostnames(v []string)`

SetPortalDeniedHostnames sets PortalDeniedHostnames field to given value.

### HasPortalDeniedHostnames

`func (o *Wlan) HasPortalDeniedHostnames() bool`

HasPortalDeniedHostnames returns a boolean if a field has been set.

### GetPortalImage

`func (o *Wlan) GetPortalImage() string`

GetPortalImage returns the PortalImage field if non-nil, zero value otherwise.

### GetPortalImageOk

`func (o *Wlan) GetPortalImageOk() (*string, bool)`

GetPortalImageOk returns a tuple with the PortalImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalImage

`func (o *Wlan) SetPortalImage(v string)`

SetPortalImage sets PortalImage field to given value.

### HasPortalImage

`func (o *Wlan) HasPortalImage() bool`

HasPortalImage returns a boolean if a field has been set.

### SetPortalImageNil

`func (o *Wlan) SetPortalImageNil(b bool)`

 SetPortalImageNil sets the value for PortalImage to be an explicit nil

### UnsetPortalImage
`func (o *Wlan) UnsetPortalImage()`

UnsetPortalImage ensures that no value is present for PortalImage, not even an explicit nil
### GetPortalSsoUrl

`func (o *Wlan) GetPortalSsoUrl() string`

GetPortalSsoUrl returns the PortalSsoUrl field if non-nil, zero value otherwise.

### GetPortalSsoUrlOk

`func (o *Wlan) GetPortalSsoUrlOk() (*string, bool)`

GetPortalSsoUrlOk returns a tuple with the PortalSsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalSsoUrl

`func (o *Wlan) SetPortalSsoUrl(v string)`

SetPortalSsoUrl sets PortalSsoUrl field to given value.

### HasPortalSsoUrl

`func (o *Wlan) HasPortalSsoUrl() bool`

HasPortalSsoUrl returns a boolean if a field has been set.

### SetPortalSsoUrlNil

`func (o *Wlan) SetPortalSsoUrlNil(b bool)`

 SetPortalSsoUrlNil sets the value for PortalSsoUrl to be an explicit nil

### UnsetPortalSsoUrl
`func (o *Wlan) UnsetPortalSsoUrl()`

UnsetPortalSsoUrl ensures that no value is present for PortalSsoUrl, not even an explicit nil
### GetPortalTemplateUrl

`func (o *Wlan) GetPortalTemplateUrl() string`

GetPortalTemplateUrl returns the PortalTemplateUrl field if non-nil, zero value otherwise.

### GetPortalTemplateUrlOk

`func (o *Wlan) GetPortalTemplateUrlOk() (*string, bool)`

GetPortalTemplateUrlOk returns a tuple with the PortalTemplateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalTemplateUrl

`func (o *Wlan) SetPortalTemplateUrl(v string)`

SetPortalTemplateUrl sets PortalTemplateUrl field to given value.

### HasPortalTemplateUrl

`func (o *Wlan) HasPortalTemplateUrl() bool`

HasPortalTemplateUrl returns a boolean if a field has been set.

### SetPortalTemplateUrlNil

`func (o *Wlan) SetPortalTemplateUrlNil(b bool)`

 SetPortalTemplateUrlNil sets the value for PortalTemplateUrl to be an explicit nil

### UnsetPortalTemplateUrl
`func (o *Wlan) UnsetPortalTemplateUrl()`

UnsetPortalTemplateUrl ensures that no value is present for PortalTemplateUrl, not even an explicit nil
### GetQos

`func (o *Wlan) GetQos() WlanQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Wlan) GetQosOk() (*WlanQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Wlan) SetQos(v WlanQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Wlan) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetRadsec

`func (o *Wlan) GetRadsec() Radsec`

GetRadsec returns the Radsec field if non-nil, zero value otherwise.

### GetRadsecOk

`func (o *Wlan) GetRadsecOk() (*Radsec, bool)`

GetRadsecOk returns a tuple with the Radsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadsec

`func (o *Wlan) SetRadsec(v Radsec)`

SetRadsec sets Radsec field to given value.

### HasRadsec

`func (o *Wlan) HasRadsec() bool`

HasRadsec returns a boolean if a field has been set.

### GetRateset

`func (o *Wlan) GetRateset() WlanRateset`

GetRateset returns the Rateset field if non-nil, zero value otherwise.

### GetRatesetOk

`func (o *Wlan) GetRatesetOk() (*WlanRateset, bool)`

GetRatesetOk returns a tuple with the Rateset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateset

`func (o *Wlan) SetRateset(v WlanRateset)`

SetRateset sets Rateset field to given value.

### HasRateset

`func (o *Wlan) HasRateset() bool`

HasRateset returns a boolean if a field has been set.

### GetRoamMode

`func (o *Wlan) GetRoamMode() WlanRoamMode`

GetRoamMode returns the RoamMode field if non-nil, zero value otherwise.

### GetRoamModeOk

`func (o *Wlan) GetRoamModeOk() (*WlanRoamMode, bool)`

GetRoamModeOk returns a tuple with the RoamMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamMode

`func (o *Wlan) SetRoamMode(v WlanRoamMode)`

SetRoamMode sets RoamMode field to given value.

### HasRoamMode

`func (o *Wlan) HasRoamMode() bool`

HasRoamMode returns a boolean if a field has been set.

### GetSchedule

`func (o *Wlan) GetSchedule() WlanSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Wlan) GetScheduleOk() (*WlanSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Wlan) SetSchedule(v WlanSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Wlan) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSiteId

`func (o *Wlan) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Wlan) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Wlan) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Wlan) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSleExcluded

`func (o *Wlan) GetSleExcluded() bool`

GetSleExcluded returns the SleExcluded field if non-nil, zero value otherwise.

### GetSleExcludedOk

`func (o *Wlan) GetSleExcludedOk() (*bool, bool)`

GetSleExcludedOk returns a tuple with the SleExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleExcluded

`func (o *Wlan) SetSleExcluded(v bool)`

SetSleExcluded sets SleExcluded field to given value.

### HasSleExcluded

`func (o *Wlan) HasSleExcluded() bool`

HasSleExcluded returns a boolean if a field has been set.

### GetSsid

`func (o *Wlan) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *Wlan) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *Wlan) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetTemplateId

`func (o *Wlan) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Wlan) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Wlan) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Wlan) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *Wlan) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *Wlan) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetThumbnail

`func (o *Wlan) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Wlan) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Wlan) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Wlan) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### SetThumbnailNil

`func (o *Wlan) SetThumbnailNil(b bool)`

 SetThumbnailNil sets the value for Thumbnail to be an explicit nil

### UnsetThumbnail
`func (o *Wlan) UnsetThumbnail()`

UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
### GetUseEapolV1

`func (o *Wlan) GetUseEapolV1() bool`

GetUseEapolV1 returns the UseEapolV1 field if non-nil, zero value otherwise.

### GetUseEapolV1Ok

`func (o *Wlan) GetUseEapolV1Ok() (*bool, bool)`

GetUseEapolV1Ok returns a tuple with the UseEapolV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEapolV1

`func (o *Wlan) SetUseEapolV1(v bool)`

SetUseEapolV1 sets UseEapolV1 field to given value.

### HasUseEapolV1

`func (o *Wlan) HasUseEapolV1() bool`

HasUseEapolV1 returns a boolean if a field has been set.

### GetVlanEnabled

`func (o *Wlan) GetVlanEnabled() bool`

GetVlanEnabled returns the VlanEnabled field if non-nil, zero value otherwise.

### GetVlanEnabledOk

`func (o *Wlan) GetVlanEnabledOk() (*bool, bool)`

GetVlanEnabledOk returns a tuple with the VlanEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnabled

`func (o *Wlan) SetVlanEnabled(v bool)`

SetVlanEnabled sets VlanEnabled field to given value.

### HasVlanEnabled

`func (o *Wlan) HasVlanEnabled() bool`

HasVlanEnabled returns a boolean if a field has been set.

### GetVlanId

`func (o *Wlan) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *Wlan) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *Wlan) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *Wlan) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *Wlan) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *Wlan) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetVlanIds

`func (o *Wlan) GetVlanIds() []*int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *Wlan) GetVlanIdsOk() (*[]*int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *Wlan) SetVlanIds(v []*int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *Wlan) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.

### GetVlanPooling

`func (o *Wlan) GetVlanPooling() bool`

GetVlanPooling returns the VlanPooling field if non-nil, zero value otherwise.

### GetVlanPoolingOk

`func (o *Wlan) GetVlanPoolingOk() (*bool, bool)`

GetVlanPoolingOk returns a tuple with the VlanPooling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPooling

`func (o *Wlan) SetVlanPooling(v bool)`

SetVlanPooling sets VlanPooling field to given value.

### HasVlanPooling

`func (o *Wlan) HasVlanPooling() bool`

HasVlanPooling returns a boolean if a field has been set.

### GetWlanLimitDown

`func (o *Wlan) GetWlanLimitDown() int32`

GetWlanLimitDown returns the WlanLimitDown field if non-nil, zero value otherwise.

### GetWlanLimitDownOk

`func (o *Wlan) GetWlanLimitDownOk() (*int32, bool)`

GetWlanLimitDownOk returns a tuple with the WlanLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanLimitDown

`func (o *Wlan) SetWlanLimitDown(v int32)`

SetWlanLimitDown sets WlanLimitDown field to given value.

### HasWlanLimitDown

`func (o *Wlan) HasWlanLimitDown() bool`

HasWlanLimitDown returns a boolean if a field has been set.

### SetWlanLimitDownNil

`func (o *Wlan) SetWlanLimitDownNil(b bool)`

 SetWlanLimitDownNil sets the value for WlanLimitDown to be an explicit nil

### UnsetWlanLimitDown
`func (o *Wlan) UnsetWlanLimitDown()`

UnsetWlanLimitDown ensures that no value is present for WlanLimitDown, not even an explicit nil
### GetWlanLimitDownEnabled

`func (o *Wlan) GetWlanLimitDownEnabled() bool`

GetWlanLimitDownEnabled returns the WlanLimitDownEnabled field if non-nil, zero value otherwise.

### GetWlanLimitDownEnabledOk

`func (o *Wlan) GetWlanLimitDownEnabledOk() (*bool, bool)`

GetWlanLimitDownEnabledOk returns a tuple with the WlanLimitDownEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanLimitDownEnabled

`func (o *Wlan) SetWlanLimitDownEnabled(v bool)`

SetWlanLimitDownEnabled sets WlanLimitDownEnabled field to given value.

### HasWlanLimitDownEnabled

`func (o *Wlan) HasWlanLimitDownEnabled() bool`

HasWlanLimitDownEnabled returns a boolean if a field has been set.

### GetWlanLimitUp

`func (o *Wlan) GetWlanLimitUp() int32`

GetWlanLimitUp returns the WlanLimitUp field if non-nil, zero value otherwise.

### GetWlanLimitUpOk

`func (o *Wlan) GetWlanLimitUpOk() (*int32, bool)`

GetWlanLimitUpOk returns a tuple with the WlanLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanLimitUp

`func (o *Wlan) SetWlanLimitUp(v int32)`

SetWlanLimitUp sets WlanLimitUp field to given value.

### HasWlanLimitUp

`func (o *Wlan) HasWlanLimitUp() bool`

HasWlanLimitUp returns a boolean if a field has been set.

### SetWlanLimitUpNil

`func (o *Wlan) SetWlanLimitUpNil(b bool)`

 SetWlanLimitUpNil sets the value for WlanLimitUp to be an explicit nil

### UnsetWlanLimitUp
`func (o *Wlan) UnsetWlanLimitUp()`

UnsetWlanLimitUp ensures that no value is present for WlanLimitUp, not even an explicit nil
### GetWlanLimitUpEnabled

`func (o *Wlan) GetWlanLimitUpEnabled() bool`

GetWlanLimitUpEnabled returns the WlanLimitUpEnabled field if non-nil, zero value otherwise.

### GetWlanLimitUpEnabledOk

`func (o *Wlan) GetWlanLimitUpEnabledOk() (*bool, bool)`

GetWlanLimitUpEnabledOk returns a tuple with the WlanLimitUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanLimitUpEnabled

`func (o *Wlan) SetWlanLimitUpEnabled(v bool)`

SetWlanLimitUpEnabled sets WlanLimitUpEnabled field to given value.

### HasWlanLimitUpEnabled

`func (o *Wlan) HasWlanLimitUpEnabled() bool`

HasWlanLimitUpEnabled returns a boolean if a field has been set.

### GetWxtagIds

`func (o *Wlan) GetWxtagIds() []string`

GetWxtagIds returns the WxtagIds field if non-nil, zero value otherwise.

### GetWxtagIdsOk

`func (o *Wlan) GetWxtagIdsOk() (*[]string, bool)`

GetWxtagIdsOk returns a tuple with the WxtagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxtagIds

`func (o *Wlan) SetWxtagIds(v []string)`

SetWxtagIds sets WxtagIds field to given value.

### HasWxtagIds

`func (o *Wlan) HasWxtagIds() bool`

HasWxtagIds returns a boolean if a field has been set.

### SetWxtagIdsNil

`func (o *Wlan) SetWxtagIdsNil(b bool)`

 SetWxtagIdsNil sets the value for WxtagIds to be an explicit nil

### UnsetWxtagIds
`func (o *Wlan) UnsetWxtagIds()`

UnsetWxtagIds ensures that no value is present for WxtagIds, not even an explicit nil
### GetWxtunnelId

`func (o *Wlan) GetWxtunnelId() string`

GetWxtunnelId returns the WxtunnelId field if non-nil, zero value otherwise.

### GetWxtunnelIdOk

`func (o *Wlan) GetWxtunnelIdOk() (*string, bool)`

GetWxtunnelIdOk returns a tuple with the WxtunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxtunnelId

`func (o *Wlan) SetWxtunnelId(v string)`

SetWxtunnelId sets WxtunnelId field to given value.

### HasWxtunnelId

`func (o *Wlan) HasWxtunnelId() bool`

HasWxtunnelId returns a boolean if a field has been set.

### SetWxtunnelIdNil

`func (o *Wlan) SetWxtunnelIdNil(b bool)`

 SetWxtunnelIdNil sets the value for WxtunnelId to be an explicit nil

### UnsetWxtunnelId
`func (o *Wlan) UnsetWxtunnelId()`

UnsetWxtunnelId ensures that no value is present for WxtunnelId, not even an explicit nil
### GetWxtunnelRemoteId

`func (o *Wlan) GetWxtunnelRemoteId() string`

GetWxtunnelRemoteId returns the WxtunnelRemoteId field if non-nil, zero value otherwise.

### GetWxtunnelRemoteIdOk

`func (o *Wlan) GetWxtunnelRemoteIdOk() (*string, bool)`

GetWxtunnelRemoteIdOk returns a tuple with the WxtunnelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxtunnelRemoteId

`func (o *Wlan) SetWxtunnelRemoteId(v string)`

SetWxtunnelRemoteId sets WxtunnelRemoteId field to given value.

### HasWxtunnelRemoteId

`func (o *Wlan) HasWxtunnelRemoteId() bool`

HasWxtunnelRemoteId returns a boolean if a field has been set.

### SetWxtunnelRemoteIdNil

`func (o *Wlan) SetWxtunnelRemoteIdNil(b bool)`

 SetWxtunnelRemoteIdNil sets the value for WxtunnelRemoteId to be an explicit nil

### UnsetWxtunnelRemoteId
`func (o *Wlan) UnsetWxtunnelRemoteId()`

UnsetWxtunnelRemoteId ensures that no value is present for WxtunnelRemoteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



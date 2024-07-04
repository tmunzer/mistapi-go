# SiteSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclPolicies** | Pointer to [**[]AclPolicy**](AclPolicy.md) |  | [optional] 
**AclTags** | Pointer to [**map[string]AclTag**](AclTag.md) | ACL Tags to identify traffic source or destination. Key name is the tag name | [optional] 
**AdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**Analytic** | Pointer to [**SiteSettingAnalytic**](SiteSettingAnalytic.md) |  | [optional] 
**ApMatching** | Pointer to [**SiteSettingApMatching**](SiteSettingApMatching.md) |  | [optional] 
**ApPortConfig** | Pointer to [**SiteSettingApPortConfig**](SiteSettingApPortConfig.md) |  | [optional] 
**ApUpdownThreshold** | Pointer to **NullableInt32** | enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and &#x60;device_updown_threshold&#x60; is ignored. | [optional] 
**AutoPlacement** | Pointer to [**SiteSettingAutoPlacement**](SiteSettingAutoPlacement.md) |  | [optional] 
**AutoUpgrade** | Pointer to [**SiteSettingAutoUpgrade**](SiteSettingAutoUpgrade.md) |  | [optional] 
**BlacklistUrl** | Pointer to **string** |  | [optional] [readonly] 
**BleConfig** | Pointer to [**BleConfig**](BleConfig.md) |  | [optional] 
**ConfigAutoRevert** | Pointer to **bool** | whether to enable ap auto config revert | [optional] [default to false]
**ConfigPushPolicy** | Pointer to [**SiteSettingConfigPushPolicy**](SiteSettingConfigPushPolicy.md) |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**CriticalUrlMonitoring** | Pointer to [**SiteSettingCriticalUrlMonitoring**](SiteSettingCriticalUrlMonitoring.md) |  | [optional] 
**DeviceUpdownThreshold** | Pointer to **int32** | sending AP_DISCONNECTED event in device-updowns only if AP_CONNECTED is not seen within the threshold, in minutes | [optional] [default to 0]
**DhcpSnooping** | Pointer to [**DhcpSnooping**](DhcpSnooping.md) |  | [optional] 
**DisabledSystemDefinedPortUsages** | Pointer to **[]string** | if some system-default port usages are not desired - namely, ap / iot / uplink | [optional] 
**DnsServers** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**DnsSuffix** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**Engagement** | Pointer to [**SiteEngagement**](SiteEngagement.md) |  | [optional] 
**EvpnOptions** | Pointer to [**EvpnOptions**](EvpnOptions.md) |  | [optional] 
**ExtraRoutes** | Pointer to [**map[string]ExtraRouteProperties**](ExtraRouteProperties.md) |  | [optional] 
**ExtraRoutes6** | Pointer to [**map[string]ExtraRoute6Properties**](ExtraRoute6Properties.md) | Property key is the destination CIDR (e.g. \&quot;2a02:1234:420a:10c9::/64\&quot;) | [optional] 
**Flags** | Pointer to **map[string]string** | name/val pair objects for location engine to use | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Gateway** | Pointer to [**GatewayTemplate**](GatewayTemplate.md) |  | [optional] 
**GatewayAdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**GatewayMgmt** | Pointer to [**SiteSettingGatewayMgmt**](SiteSettingGatewayMgmt.md) |  | [optional] 
**GatewayUpdownThreshold** | Pointer to **NullableInt32** | enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and &#x60;device_updown_threshold&#x60; is ignored. | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Led** | Pointer to [**ApLed**](ApLed.md) |  | [optional] 
**MistNac** | Pointer to [**NetworkTemplateMistNac**](NetworkTemplateMistNac.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Mxedge** | Pointer to [**SiteSettingMxedge**](SiteSettingMxedge.md) |  | [optional] 
**MxedgeMgmt** | Pointer to [**MxedgeMgmt**](MxedgeMgmt.md) |  | [optional] 
**Mxtunnels** | Pointer to [**SiteMxtunnel**](SiteMxtunnel.md) |  | [optional] 
**Networks** | Pointer to [**map[string]SwitchNetwork**](SwitchNetwork.md) | Property key is network name | [optional] 
**NtpServers** | Pointer to **[]string** | list of NTP servers | [optional] 
**Occupancy** | Pointer to [**SiteOccupancyAnalytics**](SiteOccupancyAnalytics.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**OspfAreas** | Pointer to [**map[string]OspfAreas**](OspfAreas.md) | OSPF Areas can be configured both in Site Level (Switch/Gateway Template) or Device (Switch/Gateway) Level | [optional] 
**PaloaltoNetworks** | Pointer to [**SiteSettingPaloaltoNetworks**](SiteSettingPaloaltoNetworks.md) |  | [optional] 
**PersistConfigOnDevice** | Pointer to **bool** | whether to store the config on AP | [optional] [default to false]
**PortMirrorings** | Pointer to [**map[string]SwitchPortMirroring**](SwitchPortMirroring.md) | Property key is the port mirroring instance name port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. | [optional] 
**PortUsages** | Pointer to [**map[string]SwitchPortUsage**](SwitchPortUsage.md) |  | [optional] 
**Proxy** | Pointer to [**Proxy**](Proxy.md) |  | [optional] 
**RadioConfig** | Pointer to [**ApRadio**](ApRadio.md) |  | [optional] 
**RadiusConfig** | Pointer to [**RadiusConfig**](RadiusConfig.md) |  | [optional] 
**RemoteSyslog** | Pointer to [**RemoteSyslog**](RemoteSyslog.md) |  | [optional] 
**ReportGatt** | Pointer to **bool** | whether AP should periodically connect to BLE devices and report GATT device info (device name, manufacturer name, serial number, battery %, temperature, humidity) | [optional] [default to false]
**Rogue** | Pointer to [**SiteRogue**](SiteRogue.md) |  | [optional] 
**Rtsa** | Pointer to [**SiteSettingRtsa**](SiteSettingRtsa.md) |  | [optional] 
**SimpleAlert** | Pointer to [**SimpleAlert**](SimpleAlert.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Skyatp** | Pointer to [**SiteSettingSkyatp**](SiteSettingSkyatp.md) |  | [optional] 
**SnmpConfig** | Pointer to [**SnmpConfig**](SnmpConfig.md) |  | [optional] 
**SrxApp** | Pointer to [**SiteSettingSrxApp**](SiteSettingSrxApp.md) |  | [optional] 
**SshKeys** | Pointer to **[]string** | when limit_ssh_access &#x3D; true in Org Setting, list of SSH public keys provided by Mist Support to install onto APs (see Org:Setting) | [optional] 
**Ssr** | Pointer to [**SiteSettingSsr**](SiteSettingSsr.md) |  | [optional] 
**StatusPortal** | Pointer to [**SiteSettingStatusPortal**](SiteSettingStatusPortal.md) |  | [optional] 
**Switch** | Pointer to [**NetworkTemplate**](NetworkTemplate.md) |  | [optional] 
**SwitchMatching** | Pointer to [**SwitchMatching**](SwitchMatching.md) |  | [optional] 
**SwitchMgmt** | Pointer to [**SwitchMgmt**](SwitchMgmt.md) |  | [optional] 
**SwitchUpdownThreshold** | Pointer to **NullableInt32** | enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and &#x60;device_updown_threshold&#x60; is ignored. | [optional] 
**SyntheticTest** | Pointer to [**SyntheticTestConfig**](SyntheticTestConfig.md) |  | [optional] 
**TrackAnonymousDevices** | Pointer to **bool** | whether to track anonymous BLE assets (requires ‘track_asset’ enabled) | [optional] [default to false]
**TuntermMonitoring** | Pointer to [**[]TuntermMonitoringItem**](TuntermMonitoringItem.md) |  | [optional] 
**TuntermMonitoringDisabled** | Pointer to **bool** |  | [optional] [default to false]
**TuntermMulticastConfig** | Pointer to [**SiteSettingTuntermMulticastConfig**](SiteSettingTuntermMulticastConfig.md) |  | [optional] 
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 
**Vna** | Pointer to [**SiteSettingVna**](SiteSettingVna.md) |  | [optional] 
**VrfConfig** | Pointer to [**VrfConfig**](VrfConfig.md) |  | [optional] 
**VrfInstances** | Pointer to [**map[string]VrfInstance**](VrfInstance.md) | Property key is the network name | [optional] 
**VrrpGroups** | Pointer to [**map[string]VrrpGroup**](VrrpGroup.md) | Property key is the vrrp group | [optional] 
**WanVna** | Pointer to [**SiteSettingWanVna**](SiteSettingWanVna.md) |  | [optional] 
**WatchedStationUrl** | Pointer to **string** |  | [optional] [readonly] 
**WhitelistUrl** | Pointer to **string** |  | [optional] [readonly] 
**Wids** | Pointer to [**SiteWids**](SiteWids.md) |  | [optional] 
**Wifi** | Pointer to [**SiteWifi**](SiteWifi.md) |  | [optional] 
**WiredVna** | Pointer to [**SiteSettingWiredVna**](SiteSettingWiredVna.md) |  | [optional] 
**ZoneOccupancyAlert** | Pointer to [**SiteZoneOccupancyAlert**](SiteZoneOccupancyAlert.md) |  | [optional] 

## Methods

### NewSiteSetting

`func NewSiteSetting() *SiteSetting`

NewSiteSetting instantiates a new SiteSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingWithDefaults

`func NewSiteSettingWithDefaults() *SiteSetting`

NewSiteSettingWithDefaults instantiates a new SiteSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclPolicies

`func (o *SiteSetting) GetAclPolicies() []AclPolicy`

GetAclPolicies returns the AclPolicies field if non-nil, zero value otherwise.

### GetAclPoliciesOk

`func (o *SiteSetting) GetAclPoliciesOk() (*[]AclPolicy, bool)`

GetAclPoliciesOk returns a tuple with the AclPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclPolicies

`func (o *SiteSetting) SetAclPolicies(v []AclPolicy)`

SetAclPolicies sets AclPolicies field to given value.

### HasAclPolicies

`func (o *SiteSetting) HasAclPolicies() bool`

HasAclPolicies returns a boolean if a field has been set.

### GetAclTags

`func (o *SiteSetting) GetAclTags() map[string]AclTag`

GetAclTags returns the AclTags field if non-nil, zero value otherwise.

### GetAclTagsOk

`func (o *SiteSetting) GetAclTagsOk() (*map[string]AclTag, bool)`

GetAclTagsOk returns a tuple with the AclTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclTags

`func (o *SiteSetting) SetAclTags(v map[string]AclTag)`

SetAclTags sets AclTags field to given value.

### HasAclTags

`func (o *SiteSetting) HasAclTags() bool`

HasAclTags returns a boolean if a field has been set.

### GetAdditionalConfigCmds

`func (o *SiteSetting) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *SiteSetting) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *SiteSetting) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *SiteSetting) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetAnalytic

`func (o *SiteSetting) GetAnalytic() SiteSettingAnalytic`

GetAnalytic returns the Analytic field if non-nil, zero value otherwise.

### GetAnalyticOk

`func (o *SiteSetting) GetAnalyticOk() (*SiteSettingAnalytic, bool)`

GetAnalyticOk returns a tuple with the Analytic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytic

`func (o *SiteSetting) SetAnalytic(v SiteSettingAnalytic)`

SetAnalytic sets Analytic field to given value.

### HasAnalytic

`func (o *SiteSetting) HasAnalytic() bool`

HasAnalytic returns a boolean if a field has been set.

### GetApMatching

`func (o *SiteSetting) GetApMatching() SiteSettingApMatching`

GetApMatching returns the ApMatching field if non-nil, zero value otherwise.

### GetApMatchingOk

`func (o *SiteSetting) GetApMatchingOk() (*SiteSettingApMatching, bool)`

GetApMatchingOk returns a tuple with the ApMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMatching

`func (o *SiteSetting) SetApMatching(v SiteSettingApMatching)`

SetApMatching sets ApMatching field to given value.

### HasApMatching

`func (o *SiteSetting) HasApMatching() bool`

HasApMatching returns a boolean if a field has been set.

### GetApPortConfig

`func (o *SiteSetting) GetApPortConfig() SiteSettingApPortConfig`

GetApPortConfig returns the ApPortConfig field if non-nil, zero value otherwise.

### GetApPortConfigOk

`func (o *SiteSetting) GetApPortConfigOk() (*SiteSettingApPortConfig, bool)`

GetApPortConfigOk returns a tuple with the ApPortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApPortConfig

`func (o *SiteSetting) SetApPortConfig(v SiteSettingApPortConfig)`

SetApPortConfig sets ApPortConfig field to given value.

### HasApPortConfig

`func (o *SiteSetting) HasApPortConfig() bool`

HasApPortConfig returns a boolean if a field has been set.

### GetApUpdownThreshold

`func (o *SiteSetting) GetApUpdownThreshold() int32`

GetApUpdownThreshold returns the ApUpdownThreshold field if non-nil, zero value otherwise.

### GetApUpdownThresholdOk

`func (o *SiteSetting) GetApUpdownThresholdOk() (*int32, bool)`

GetApUpdownThresholdOk returns a tuple with the ApUpdownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApUpdownThreshold

`func (o *SiteSetting) SetApUpdownThreshold(v int32)`

SetApUpdownThreshold sets ApUpdownThreshold field to given value.

### HasApUpdownThreshold

`func (o *SiteSetting) HasApUpdownThreshold() bool`

HasApUpdownThreshold returns a boolean if a field has been set.

### SetApUpdownThresholdNil

`func (o *SiteSetting) SetApUpdownThresholdNil(b bool)`

 SetApUpdownThresholdNil sets the value for ApUpdownThreshold to be an explicit nil

### UnsetApUpdownThreshold
`func (o *SiteSetting) UnsetApUpdownThreshold()`

UnsetApUpdownThreshold ensures that no value is present for ApUpdownThreshold, not even an explicit nil
### GetAutoPlacement

`func (o *SiteSetting) GetAutoPlacement() SiteSettingAutoPlacement`

GetAutoPlacement returns the AutoPlacement field if non-nil, zero value otherwise.

### GetAutoPlacementOk

`func (o *SiteSetting) GetAutoPlacementOk() (*SiteSettingAutoPlacement, bool)`

GetAutoPlacementOk returns a tuple with the AutoPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlacement

`func (o *SiteSetting) SetAutoPlacement(v SiteSettingAutoPlacement)`

SetAutoPlacement sets AutoPlacement field to given value.

### HasAutoPlacement

`func (o *SiteSetting) HasAutoPlacement() bool`

HasAutoPlacement returns a boolean if a field has been set.

### GetAutoUpgrade

`func (o *SiteSetting) GetAutoUpgrade() SiteSettingAutoUpgrade`

GetAutoUpgrade returns the AutoUpgrade field if non-nil, zero value otherwise.

### GetAutoUpgradeOk

`func (o *SiteSetting) GetAutoUpgradeOk() (*SiteSettingAutoUpgrade, bool)`

GetAutoUpgradeOk returns a tuple with the AutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgrade

`func (o *SiteSetting) SetAutoUpgrade(v SiteSettingAutoUpgrade)`

SetAutoUpgrade sets AutoUpgrade field to given value.

### HasAutoUpgrade

`func (o *SiteSetting) HasAutoUpgrade() bool`

HasAutoUpgrade returns a boolean if a field has been set.

### GetBlacklistUrl

`func (o *SiteSetting) GetBlacklistUrl() string`

GetBlacklistUrl returns the BlacklistUrl field if non-nil, zero value otherwise.

### GetBlacklistUrlOk

`func (o *SiteSetting) GetBlacklistUrlOk() (*string, bool)`

GetBlacklistUrlOk returns a tuple with the BlacklistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistUrl

`func (o *SiteSetting) SetBlacklistUrl(v string)`

SetBlacklistUrl sets BlacklistUrl field to given value.

### HasBlacklistUrl

`func (o *SiteSetting) HasBlacklistUrl() bool`

HasBlacklistUrl returns a boolean if a field has been set.

### GetBleConfig

`func (o *SiteSetting) GetBleConfig() BleConfig`

GetBleConfig returns the BleConfig field if non-nil, zero value otherwise.

### GetBleConfigOk

`func (o *SiteSetting) GetBleConfigOk() (*BleConfig, bool)`

GetBleConfigOk returns a tuple with the BleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleConfig

`func (o *SiteSetting) SetBleConfig(v BleConfig)`

SetBleConfig sets BleConfig field to given value.

### HasBleConfig

`func (o *SiteSetting) HasBleConfig() bool`

HasBleConfig returns a boolean if a field has been set.

### GetConfigAutoRevert

`func (o *SiteSetting) GetConfigAutoRevert() bool`

GetConfigAutoRevert returns the ConfigAutoRevert field if non-nil, zero value otherwise.

### GetConfigAutoRevertOk

`func (o *SiteSetting) GetConfigAutoRevertOk() (*bool, bool)`

GetConfigAutoRevertOk returns a tuple with the ConfigAutoRevert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAutoRevert

`func (o *SiteSetting) SetConfigAutoRevert(v bool)`

SetConfigAutoRevert sets ConfigAutoRevert field to given value.

### HasConfigAutoRevert

`func (o *SiteSetting) HasConfigAutoRevert() bool`

HasConfigAutoRevert returns a boolean if a field has been set.

### GetConfigPushPolicy

`func (o *SiteSetting) GetConfigPushPolicy() SiteSettingConfigPushPolicy`

GetConfigPushPolicy returns the ConfigPushPolicy field if non-nil, zero value otherwise.

### GetConfigPushPolicyOk

`func (o *SiteSetting) GetConfigPushPolicyOk() (*SiteSettingConfigPushPolicy, bool)`

GetConfigPushPolicyOk returns a tuple with the ConfigPushPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigPushPolicy

`func (o *SiteSetting) SetConfigPushPolicy(v SiteSettingConfigPushPolicy)`

SetConfigPushPolicy sets ConfigPushPolicy field to given value.

### HasConfigPushPolicy

`func (o *SiteSetting) HasConfigPushPolicy() bool`

HasConfigPushPolicy returns a boolean if a field has been set.

### GetCreatedTime

`func (o *SiteSetting) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SiteSetting) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SiteSetting) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SiteSetting) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCriticalUrlMonitoring

`func (o *SiteSetting) GetCriticalUrlMonitoring() SiteSettingCriticalUrlMonitoring`

GetCriticalUrlMonitoring returns the CriticalUrlMonitoring field if non-nil, zero value otherwise.

### GetCriticalUrlMonitoringOk

`func (o *SiteSetting) GetCriticalUrlMonitoringOk() (*SiteSettingCriticalUrlMonitoring, bool)`

GetCriticalUrlMonitoringOk returns a tuple with the CriticalUrlMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalUrlMonitoring

`func (o *SiteSetting) SetCriticalUrlMonitoring(v SiteSettingCriticalUrlMonitoring)`

SetCriticalUrlMonitoring sets CriticalUrlMonitoring field to given value.

### HasCriticalUrlMonitoring

`func (o *SiteSetting) HasCriticalUrlMonitoring() bool`

HasCriticalUrlMonitoring returns a boolean if a field has been set.

### GetDeviceUpdownThreshold

`func (o *SiteSetting) GetDeviceUpdownThreshold() int32`

GetDeviceUpdownThreshold returns the DeviceUpdownThreshold field if non-nil, zero value otherwise.

### GetDeviceUpdownThresholdOk

`func (o *SiteSetting) GetDeviceUpdownThresholdOk() (*int32, bool)`

GetDeviceUpdownThresholdOk returns a tuple with the DeviceUpdownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUpdownThreshold

`func (o *SiteSetting) SetDeviceUpdownThreshold(v int32)`

SetDeviceUpdownThreshold sets DeviceUpdownThreshold field to given value.

### HasDeviceUpdownThreshold

`func (o *SiteSetting) HasDeviceUpdownThreshold() bool`

HasDeviceUpdownThreshold returns a boolean if a field has been set.

### GetDhcpSnooping

`func (o *SiteSetting) GetDhcpSnooping() DhcpSnooping`

GetDhcpSnooping returns the DhcpSnooping field if non-nil, zero value otherwise.

### GetDhcpSnoopingOk

`func (o *SiteSetting) GetDhcpSnoopingOk() (*DhcpSnooping, bool)`

GetDhcpSnoopingOk returns a tuple with the DhcpSnooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnooping

`func (o *SiteSetting) SetDhcpSnooping(v DhcpSnooping)`

SetDhcpSnooping sets DhcpSnooping field to given value.

### HasDhcpSnooping

`func (o *SiteSetting) HasDhcpSnooping() bool`

HasDhcpSnooping returns a boolean if a field has been set.

### GetDisabledSystemDefinedPortUsages

`func (o *SiteSetting) GetDisabledSystemDefinedPortUsages() []string`

GetDisabledSystemDefinedPortUsages returns the DisabledSystemDefinedPortUsages field if non-nil, zero value otherwise.

### GetDisabledSystemDefinedPortUsagesOk

`func (o *SiteSetting) GetDisabledSystemDefinedPortUsagesOk() (*[]string, bool)`

GetDisabledSystemDefinedPortUsagesOk returns a tuple with the DisabledSystemDefinedPortUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSystemDefinedPortUsages

`func (o *SiteSetting) SetDisabledSystemDefinedPortUsages(v []string)`

SetDisabledSystemDefinedPortUsages sets DisabledSystemDefinedPortUsages field to given value.

### HasDisabledSystemDefinedPortUsages

`func (o *SiteSetting) HasDisabledSystemDefinedPortUsages() bool`

HasDisabledSystemDefinedPortUsages returns a boolean if a field has been set.

### GetDnsServers

`func (o *SiteSetting) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *SiteSetting) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *SiteSetting) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *SiteSetting) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *SiteSetting) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *SiteSetting) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *SiteSetting) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *SiteSetting) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetEngagement

`func (o *SiteSetting) GetEngagement() SiteEngagement`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *SiteSetting) GetEngagementOk() (*SiteEngagement, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *SiteSetting) SetEngagement(v SiteEngagement)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *SiteSetting) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### GetEvpnOptions

`func (o *SiteSetting) GetEvpnOptions() EvpnOptions`

GetEvpnOptions returns the EvpnOptions field if non-nil, zero value otherwise.

### GetEvpnOptionsOk

`func (o *SiteSetting) GetEvpnOptionsOk() (*EvpnOptions, bool)`

GetEvpnOptionsOk returns a tuple with the EvpnOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnOptions

`func (o *SiteSetting) SetEvpnOptions(v EvpnOptions)`

SetEvpnOptions sets EvpnOptions field to given value.

### HasEvpnOptions

`func (o *SiteSetting) HasEvpnOptions() bool`

HasEvpnOptions returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *SiteSetting) GetExtraRoutes() map[string]ExtraRouteProperties`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *SiteSetting) GetExtraRoutesOk() (*map[string]ExtraRouteProperties, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *SiteSetting) SetExtraRoutes(v map[string]ExtraRouteProperties)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *SiteSetting) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetExtraRoutes6

`func (o *SiteSetting) GetExtraRoutes6() map[string]ExtraRoute6Properties`

GetExtraRoutes6 returns the ExtraRoutes6 field if non-nil, zero value otherwise.

### GetExtraRoutes6Ok

`func (o *SiteSetting) GetExtraRoutes6Ok() (*map[string]ExtraRoute6Properties, bool)`

GetExtraRoutes6Ok returns a tuple with the ExtraRoutes6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes6

`func (o *SiteSetting) SetExtraRoutes6(v map[string]ExtraRoute6Properties)`

SetExtraRoutes6 sets ExtraRoutes6 field to given value.

### HasExtraRoutes6

`func (o *SiteSetting) HasExtraRoutes6() bool`

HasExtraRoutes6 returns a boolean if a field has been set.

### GetFlags

`func (o *SiteSetting) GetFlags() map[string]string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SiteSetting) GetFlagsOk() (*map[string]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SiteSetting) SetFlags(v map[string]string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *SiteSetting) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetForSite

`func (o *SiteSetting) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *SiteSetting) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *SiteSetting) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *SiteSetting) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetGateway

`func (o *SiteSetting) GetGateway() GatewayTemplate`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *SiteSetting) GetGatewayOk() (*GatewayTemplate, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *SiteSetting) SetGateway(v GatewayTemplate)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *SiteSetting) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayAdditionalConfigCmds

`func (o *SiteSetting) GetGatewayAdditionalConfigCmds() []string`

GetGatewayAdditionalConfigCmds returns the GatewayAdditionalConfigCmds field if non-nil, zero value otherwise.

### GetGatewayAdditionalConfigCmdsOk

`func (o *SiteSetting) GetGatewayAdditionalConfigCmdsOk() (*[]string, bool)`

GetGatewayAdditionalConfigCmdsOk returns a tuple with the GatewayAdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAdditionalConfigCmds

`func (o *SiteSetting) SetGatewayAdditionalConfigCmds(v []string)`

SetGatewayAdditionalConfigCmds sets GatewayAdditionalConfigCmds field to given value.

### HasGatewayAdditionalConfigCmds

`func (o *SiteSetting) HasGatewayAdditionalConfigCmds() bool`

HasGatewayAdditionalConfigCmds returns a boolean if a field has been set.

### GetGatewayMgmt

`func (o *SiteSetting) GetGatewayMgmt() SiteSettingGatewayMgmt`

GetGatewayMgmt returns the GatewayMgmt field if non-nil, zero value otherwise.

### GetGatewayMgmtOk

`func (o *SiteSetting) GetGatewayMgmtOk() (*SiteSettingGatewayMgmt, bool)`

GetGatewayMgmtOk returns a tuple with the GatewayMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMgmt

`func (o *SiteSetting) SetGatewayMgmt(v SiteSettingGatewayMgmt)`

SetGatewayMgmt sets GatewayMgmt field to given value.

### HasGatewayMgmt

`func (o *SiteSetting) HasGatewayMgmt() bool`

HasGatewayMgmt returns a boolean if a field has been set.

### GetGatewayUpdownThreshold

`func (o *SiteSetting) GetGatewayUpdownThreshold() int32`

GetGatewayUpdownThreshold returns the GatewayUpdownThreshold field if non-nil, zero value otherwise.

### GetGatewayUpdownThresholdOk

`func (o *SiteSetting) GetGatewayUpdownThresholdOk() (*int32, bool)`

GetGatewayUpdownThresholdOk returns a tuple with the GatewayUpdownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUpdownThreshold

`func (o *SiteSetting) SetGatewayUpdownThreshold(v int32)`

SetGatewayUpdownThreshold sets GatewayUpdownThreshold field to given value.

### HasGatewayUpdownThreshold

`func (o *SiteSetting) HasGatewayUpdownThreshold() bool`

HasGatewayUpdownThreshold returns a boolean if a field has been set.

### SetGatewayUpdownThresholdNil

`func (o *SiteSetting) SetGatewayUpdownThresholdNil(b bool)`

 SetGatewayUpdownThresholdNil sets the value for GatewayUpdownThreshold to be an explicit nil

### UnsetGatewayUpdownThreshold
`func (o *SiteSetting) UnsetGatewayUpdownThreshold()`

UnsetGatewayUpdownThreshold ensures that no value is present for GatewayUpdownThreshold, not even an explicit nil
### GetId

`func (o *SiteSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteSetting) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SiteSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLed

`func (o *SiteSetting) GetLed() ApLed`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *SiteSetting) GetLedOk() (*ApLed, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *SiteSetting) SetLed(v ApLed)`

SetLed sets Led field to given value.

### HasLed

`func (o *SiteSetting) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetMistNac

`func (o *SiteSetting) GetMistNac() NetworkTemplateMistNac`

GetMistNac returns the MistNac field if non-nil, zero value otherwise.

### GetMistNacOk

`func (o *SiteSetting) GetMistNacOk() (*NetworkTemplateMistNac, bool)`

GetMistNacOk returns a tuple with the MistNac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistNac

`func (o *SiteSetting) SetMistNac(v NetworkTemplateMistNac)`

SetMistNac sets MistNac field to given value.

### HasMistNac

`func (o *SiteSetting) HasMistNac() bool`

HasMistNac returns a boolean if a field has been set.

### GetModifiedTime

`func (o *SiteSetting) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *SiteSetting) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *SiteSetting) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *SiteSetting) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMxedge

`func (o *SiteSetting) GetMxedge() SiteSettingMxedge`

GetMxedge returns the Mxedge field if non-nil, zero value otherwise.

### GetMxedgeOk

`func (o *SiteSetting) GetMxedgeOk() (*SiteSettingMxedge, bool)`

GetMxedgeOk returns a tuple with the Mxedge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedge

`func (o *SiteSetting) SetMxedge(v SiteSettingMxedge)`

SetMxedge sets Mxedge field to given value.

### HasMxedge

`func (o *SiteSetting) HasMxedge() bool`

HasMxedge returns a boolean if a field has been set.

### GetMxedgeMgmt

`func (o *SiteSetting) GetMxedgeMgmt() MxedgeMgmt`

GetMxedgeMgmt returns the MxedgeMgmt field if non-nil, zero value otherwise.

### GetMxedgeMgmtOk

`func (o *SiteSetting) GetMxedgeMgmtOk() (*MxedgeMgmt, bool)`

GetMxedgeMgmtOk returns a tuple with the MxedgeMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeMgmt

`func (o *SiteSetting) SetMxedgeMgmt(v MxedgeMgmt)`

SetMxedgeMgmt sets MxedgeMgmt field to given value.

### HasMxedgeMgmt

`func (o *SiteSetting) HasMxedgeMgmt() bool`

HasMxedgeMgmt returns a boolean if a field has been set.

### GetMxtunnels

`func (o *SiteSetting) GetMxtunnels() SiteMxtunnel`

GetMxtunnels returns the Mxtunnels field if non-nil, zero value otherwise.

### GetMxtunnelsOk

`func (o *SiteSetting) GetMxtunnelsOk() (*SiteMxtunnel, bool)`

GetMxtunnelsOk returns a tuple with the Mxtunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxtunnels

`func (o *SiteSetting) SetMxtunnels(v SiteMxtunnel)`

SetMxtunnels sets Mxtunnels field to given value.

### HasMxtunnels

`func (o *SiteSetting) HasMxtunnels() bool`

HasMxtunnels returns a boolean if a field has been set.

### GetNetworks

`func (o *SiteSetting) GetNetworks() map[string]SwitchNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *SiteSetting) GetNetworksOk() (*map[string]SwitchNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *SiteSetting) SetNetworks(v map[string]SwitchNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *SiteSetting) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNtpServers

`func (o *SiteSetting) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *SiteSetting) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *SiteSetting) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *SiteSetting) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOccupancy

`func (o *SiteSetting) GetOccupancy() SiteOccupancyAnalytics`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *SiteSetting) GetOccupancyOk() (*SiteOccupancyAnalytics, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *SiteSetting) SetOccupancy(v SiteOccupancyAnalytics)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *SiteSetting) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.

### GetOrgId

`func (o *SiteSetting) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SiteSetting) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SiteSetting) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SiteSetting) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOspfAreas

`func (o *SiteSetting) GetOspfAreas() map[string]OspfAreas`

GetOspfAreas returns the OspfAreas field if non-nil, zero value otherwise.

### GetOspfAreasOk

`func (o *SiteSetting) GetOspfAreasOk() (*map[string]OspfAreas, bool)`

GetOspfAreasOk returns a tuple with the OspfAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfAreas

`func (o *SiteSetting) SetOspfAreas(v map[string]OspfAreas)`

SetOspfAreas sets OspfAreas field to given value.

### HasOspfAreas

`func (o *SiteSetting) HasOspfAreas() bool`

HasOspfAreas returns a boolean if a field has been set.

### GetPaloaltoNetworks

`func (o *SiteSetting) GetPaloaltoNetworks() SiteSettingPaloaltoNetworks`

GetPaloaltoNetworks returns the PaloaltoNetworks field if non-nil, zero value otherwise.

### GetPaloaltoNetworksOk

`func (o *SiteSetting) GetPaloaltoNetworksOk() (*SiteSettingPaloaltoNetworks, bool)`

GetPaloaltoNetworksOk returns a tuple with the PaloaltoNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaloaltoNetworks

`func (o *SiteSetting) SetPaloaltoNetworks(v SiteSettingPaloaltoNetworks)`

SetPaloaltoNetworks sets PaloaltoNetworks field to given value.

### HasPaloaltoNetworks

`func (o *SiteSetting) HasPaloaltoNetworks() bool`

HasPaloaltoNetworks returns a boolean if a field has been set.

### GetPersistConfigOnDevice

`func (o *SiteSetting) GetPersistConfigOnDevice() bool`

GetPersistConfigOnDevice returns the PersistConfigOnDevice field if non-nil, zero value otherwise.

### GetPersistConfigOnDeviceOk

`func (o *SiteSetting) GetPersistConfigOnDeviceOk() (*bool, bool)`

GetPersistConfigOnDeviceOk returns a tuple with the PersistConfigOnDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistConfigOnDevice

`func (o *SiteSetting) SetPersistConfigOnDevice(v bool)`

SetPersistConfigOnDevice sets PersistConfigOnDevice field to given value.

### HasPersistConfigOnDevice

`func (o *SiteSetting) HasPersistConfigOnDevice() bool`

HasPersistConfigOnDevice returns a boolean if a field has been set.

### GetPortMirrorings

`func (o *SiteSetting) GetPortMirrorings() map[string]SwitchPortMirroring`

GetPortMirrorings returns the PortMirrorings field if non-nil, zero value otherwise.

### GetPortMirroringsOk

`func (o *SiteSetting) GetPortMirroringsOk() (*map[string]SwitchPortMirroring, bool)`

GetPortMirroringsOk returns a tuple with the PortMirrorings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMirrorings

`func (o *SiteSetting) SetPortMirrorings(v map[string]SwitchPortMirroring)`

SetPortMirrorings sets PortMirrorings field to given value.

### HasPortMirrorings

`func (o *SiteSetting) HasPortMirrorings() bool`

HasPortMirrorings returns a boolean if a field has been set.

### GetPortUsages

`func (o *SiteSetting) GetPortUsages() map[string]SwitchPortUsage`

GetPortUsages returns the PortUsages field if non-nil, zero value otherwise.

### GetPortUsagesOk

`func (o *SiteSetting) GetPortUsagesOk() (*map[string]SwitchPortUsage, bool)`

GetPortUsagesOk returns a tuple with the PortUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUsages

`func (o *SiteSetting) SetPortUsages(v map[string]SwitchPortUsage)`

SetPortUsages sets PortUsages field to given value.

### HasPortUsages

`func (o *SiteSetting) HasPortUsages() bool`

HasPortUsages returns a boolean if a field has been set.

### GetProxy

`func (o *SiteSetting) GetProxy() Proxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SiteSetting) GetProxyOk() (*Proxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SiteSetting) SetProxy(v Proxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SiteSetting) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetRadioConfig

`func (o *SiteSetting) GetRadioConfig() ApRadio`

GetRadioConfig returns the RadioConfig field if non-nil, zero value otherwise.

### GetRadioConfigOk

`func (o *SiteSetting) GetRadioConfigOk() (*ApRadio, bool)`

GetRadioConfigOk returns a tuple with the RadioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioConfig

`func (o *SiteSetting) SetRadioConfig(v ApRadio)`

SetRadioConfig sets RadioConfig field to given value.

### HasRadioConfig

`func (o *SiteSetting) HasRadioConfig() bool`

HasRadioConfig returns a boolean if a field has been set.

### GetRadiusConfig

`func (o *SiteSetting) GetRadiusConfig() RadiusConfig`

GetRadiusConfig returns the RadiusConfig field if non-nil, zero value otherwise.

### GetRadiusConfigOk

`func (o *SiteSetting) GetRadiusConfigOk() (*RadiusConfig, bool)`

GetRadiusConfigOk returns a tuple with the RadiusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusConfig

`func (o *SiteSetting) SetRadiusConfig(v RadiusConfig)`

SetRadiusConfig sets RadiusConfig field to given value.

### HasRadiusConfig

`func (o *SiteSetting) HasRadiusConfig() bool`

HasRadiusConfig returns a boolean if a field has been set.

### GetRemoteSyslog

`func (o *SiteSetting) GetRemoteSyslog() RemoteSyslog`

GetRemoteSyslog returns the RemoteSyslog field if non-nil, zero value otherwise.

### GetRemoteSyslogOk

`func (o *SiteSetting) GetRemoteSyslogOk() (*RemoteSyslog, bool)`

GetRemoteSyslogOk returns a tuple with the RemoteSyslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSyslog

`func (o *SiteSetting) SetRemoteSyslog(v RemoteSyslog)`

SetRemoteSyslog sets RemoteSyslog field to given value.

### HasRemoteSyslog

`func (o *SiteSetting) HasRemoteSyslog() bool`

HasRemoteSyslog returns a boolean if a field has been set.

### GetReportGatt

`func (o *SiteSetting) GetReportGatt() bool`

GetReportGatt returns the ReportGatt field if non-nil, zero value otherwise.

### GetReportGattOk

`func (o *SiteSetting) GetReportGattOk() (*bool, bool)`

GetReportGattOk returns a tuple with the ReportGatt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportGatt

`func (o *SiteSetting) SetReportGatt(v bool)`

SetReportGatt sets ReportGatt field to given value.

### HasReportGatt

`func (o *SiteSetting) HasReportGatt() bool`

HasReportGatt returns a boolean if a field has been set.

### GetRogue

`func (o *SiteSetting) GetRogue() SiteRogue`

GetRogue returns the Rogue field if non-nil, zero value otherwise.

### GetRogueOk

`func (o *SiteSetting) GetRogueOk() (*SiteRogue, bool)`

GetRogueOk returns a tuple with the Rogue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRogue

`func (o *SiteSetting) SetRogue(v SiteRogue)`

SetRogue sets Rogue field to given value.

### HasRogue

`func (o *SiteSetting) HasRogue() bool`

HasRogue returns a boolean if a field has been set.

### GetRtsa

`func (o *SiteSetting) GetRtsa() SiteSettingRtsa`

GetRtsa returns the Rtsa field if non-nil, zero value otherwise.

### GetRtsaOk

`func (o *SiteSetting) GetRtsaOk() (*SiteSettingRtsa, bool)`

GetRtsaOk returns a tuple with the Rtsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtsa

`func (o *SiteSetting) SetRtsa(v SiteSettingRtsa)`

SetRtsa sets Rtsa field to given value.

### HasRtsa

`func (o *SiteSetting) HasRtsa() bool`

HasRtsa returns a boolean if a field has been set.

### GetSimpleAlert

`func (o *SiteSetting) GetSimpleAlert() SimpleAlert`

GetSimpleAlert returns the SimpleAlert field if non-nil, zero value otherwise.

### GetSimpleAlertOk

`func (o *SiteSetting) GetSimpleAlertOk() (*SimpleAlert, bool)`

GetSimpleAlertOk returns a tuple with the SimpleAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAlert

`func (o *SiteSetting) SetSimpleAlert(v SimpleAlert)`

SetSimpleAlert sets SimpleAlert field to given value.

### HasSimpleAlert

`func (o *SiteSetting) HasSimpleAlert() bool`

HasSimpleAlert returns a boolean if a field has been set.

### GetSiteId

`func (o *SiteSetting) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteSetting) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteSetting) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteSetting) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSkyatp

`func (o *SiteSetting) GetSkyatp() SiteSettingSkyatp`

GetSkyatp returns the Skyatp field if non-nil, zero value otherwise.

### GetSkyatpOk

`func (o *SiteSetting) GetSkyatpOk() (*SiteSettingSkyatp, bool)`

GetSkyatpOk returns a tuple with the Skyatp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkyatp

`func (o *SiteSetting) SetSkyatp(v SiteSettingSkyatp)`

SetSkyatp sets Skyatp field to given value.

### HasSkyatp

`func (o *SiteSetting) HasSkyatp() bool`

HasSkyatp returns a boolean if a field has been set.

### GetSnmpConfig

`func (o *SiteSetting) GetSnmpConfig() SnmpConfig`

GetSnmpConfig returns the SnmpConfig field if non-nil, zero value otherwise.

### GetSnmpConfigOk

`func (o *SiteSetting) GetSnmpConfigOk() (*SnmpConfig, bool)`

GetSnmpConfigOk returns a tuple with the SnmpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpConfig

`func (o *SiteSetting) SetSnmpConfig(v SnmpConfig)`

SetSnmpConfig sets SnmpConfig field to given value.

### HasSnmpConfig

`func (o *SiteSetting) HasSnmpConfig() bool`

HasSnmpConfig returns a boolean if a field has been set.

### GetSrxApp

`func (o *SiteSetting) GetSrxApp() SiteSettingSrxApp`

GetSrxApp returns the SrxApp field if non-nil, zero value otherwise.

### GetSrxAppOk

`func (o *SiteSetting) GetSrxAppOk() (*SiteSettingSrxApp, bool)`

GetSrxAppOk returns a tuple with the SrxApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrxApp

`func (o *SiteSetting) SetSrxApp(v SiteSettingSrxApp)`

SetSrxApp sets SrxApp field to given value.

### HasSrxApp

`func (o *SiteSetting) HasSrxApp() bool`

HasSrxApp returns a boolean if a field has been set.

### GetSshKeys

`func (o *SiteSetting) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *SiteSetting) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *SiteSetting) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *SiteSetting) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetSsr

`func (o *SiteSetting) GetSsr() SiteSettingSsr`

GetSsr returns the Ssr field if non-nil, zero value otherwise.

### GetSsrOk

`func (o *SiteSetting) GetSsrOk() (*SiteSettingSsr, bool)`

GetSsrOk returns a tuple with the Ssr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsr

`func (o *SiteSetting) SetSsr(v SiteSettingSsr)`

SetSsr sets Ssr field to given value.

### HasSsr

`func (o *SiteSetting) HasSsr() bool`

HasSsr returns a boolean if a field has been set.

### GetStatusPortal

`func (o *SiteSetting) GetStatusPortal() SiteSettingStatusPortal`

GetStatusPortal returns the StatusPortal field if non-nil, zero value otherwise.

### GetStatusPortalOk

`func (o *SiteSetting) GetStatusPortalOk() (*SiteSettingStatusPortal, bool)`

GetStatusPortalOk returns a tuple with the StatusPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPortal

`func (o *SiteSetting) SetStatusPortal(v SiteSettingStatusPortal)`

SetStatusPortal sets StatusPortal field to given value.

### HasStatusPortal

`func (o *SiteSetting) HasStatusPortal() bool`

HasStatusPortal returns a boolean if a field has been set.

### GetSwitch

`func (o *SiteSetting) GetSwitch() NetworkTemplate`

GetSwitch returns the Switch field if non-nil, zero value otherwise.

### GetSwitchOk

`func (o *SiteSetting) GetSwitchOk() (*NetworkTemplate, bool)`

GetSwitchOk returns a tuple with the Switch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitch

`func (o *SiteSetting) SetSwitch(v NetworkTemplate)`

SetSwitch sets Switch field to given value.

### HasSwitch

`func (o *SiteSetting) HasSwitch() bool`

HasSwitch returns a boolean if a field has been set.

### GetSwitchMatching

`func (o *SiteSetting) GetSwitchMatching() SwitchMatching`

GetSwitchMatching returns the SwitchMatching field if non-nil, zero value otherwise.

### GetSwitchMatchingOk

`func (o *SiteSetting) GetSwitchMatchingOk() (*SwitchMatching, bool)`

GetSwitchMatchingOk returns a tuple with the SwitchMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMatching

`func (o *SiteSetting) SetSwitchMatching(v SwitchMatching)`

SetSwitchMatching sets SwitchMatching field to given value.

### HasSwitchMatching

`func (o *SiteSetting) HasSwitchMatching() bool`

HasSwitchMatching returns a boolean if a field has been set.

### GetSwitchMgmt

`func (o *SiteSetting) GetSwitchMgmt() SwitchMgmt`

GetSwitchMgmt returns the SwitchMgmt field if non-nil, zero value otherwise.

### GetSwitchMgmtOk

`func (o *SiteSetting) GetSwitchMgmtOk() (*SwitchMgmt, bool)`

GetSwitchMgmtOk returns a tuple with the SwitchMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMgmt

`func (o *SiteSetting) SetSwitchMgmt(v SwitchMgmt)`

SetSwitchMgmt sets SwitchMgmt field to given value.

### HasSwitchMgmt

`func (o *SiteSetting) HasSwitchMgmt() bool`

HasSwitchMgmt returns a boolean if a field has been set.

### GetSwitchUpdownThreshold

`func (o *SiteSetting) GetSwitchUpdownThreshold() int32`

GetSwitchUpdownThreshold returns the SwitchUpdownThreshold field if non-nil, zero value otherwise.

### GetSwitchUpdownThresholdOk

`func (o *SiteSetting) GetSwitchUpdownThresholdOk() (*int32, bool)`

GetSwitchUpdownThresholdOk returns a tuple with the SwitchUpdownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchUpdownThreshold

`func (o *SiteSetting) SetSwitchUpdownThreshold(v int32)`

SetSwitchUpdownThreshold sets SwitchUpdownThreshold field to given value.

### HasSwitchUpdownThreshold

`func (o *SiteSetting) HasSwitchUpdownThreshold() bool`

HasSwitchUpdownThreshold returns a boolean if a field has been set.

### SetSwitchUpdownThresholdNil

`func (o *SiteSetting) SetSwitchUpdownThresholdNil(b bool)`

 SetSwitchUpdownThresholdNil sets the value for SwitchUpdownThreshold to be an explicit nil

### UnsetSwitchUpdownThreshold
`func (o *SiteSetting) UnsetSwitchUpdownThreshold()`

UnsetSwitchUpdownThreshold ensures that no value is present for SwitchUpdownThreshold, not even an explicit nil
### GetSyntheticTest

`func (o *SiteSetting) GetSyntheticTest() SyntheticTestConfig`

GetSyntheticTest returns the SyntheticTest field if non-nil, zero value otherwise.

### GetSyntheticTestOk

`func (o *SiteSetting) GetSyntheticTestOk() (*SyntheticTestConfig, bool)`

GetSyntheticTestOk returns a tuple with the SyntheticTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticTest

`func (o *SiteSetting) SetSyntheticTest(v SyntheticTestConfig)`

SetSyntheticTest sets SyntheticTest field to given value.

### HasSyntheticTest

`func (o *SiteSetting) HasSyntheticTest() bool`

HasSyntheticTest returns a boolean if a field has been set.

### GetTrackAnonymousDevices

`func (o *SiteSetting) GetTrackAnonymousDevices() bool`

GetTrackAnonymousDevices returns the TrackAnonymousDevices field if non-nil, zero value otherwise.

### GetTrackAnonymousDevicesOk

`func (o *SiteSetting) GetTrackAnonymousDevicesOk() (*bool, bool)`

GetTrackAnonymousDevicesOk returns a tuple with the TrackAnonymousDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackAnonymousDevices

`func (o *SiteSetting) SetTrackAnonymousDevices(v bool)`

SetTrackAnonymousDevices sets TrackAnonymousDevices field to given value.

### HasTrackAnonymousDevices

`func (o *SiteSetting) HasTrackAnonymousDevices() bool`

HasTrackAnonymousDevices returns a boolean if a field has been set.

### GetTuntermMonitoring

`func (o *SiteSetting) GetTuntermMonitoring() []TuntermMonitoringItem`

GetTuntermMonitoring returns the TuntermMonitoring field if non-nil, zero value otherwise.

### GetTuntermMonitoringOk

`func (o *SiteSetting) GetTuntermMonitoringOk() (*[]TuntermMonitoringItem, bool)`

GetTuntermMonitoringOk returns a tuple with the TuntermMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermMonitoring

`func (o *SiteSetting) SetTuntermMonitoring(v []TuntermMonitoringItem)`

SetTuntermMonitoring sets TuntermMonitoring field to given value.

### HasTuntermMonitoring

`func (o *SiteSetting) HasTuntermMonitoring() bool`

HasTuntermMonitoring returns a boolean if a field has been set.

### GetTuntermMonitoringDisabled

`func (o *SiteSetting) GetTuntermMonitoringDisabled() bool`

GetTuntermMonitoringDisabled returns the TuntermMonitoringDisabled field if non-nil, zero value otherwise.

### GetTuntermMonitoringDisabledOk

`func (o *SiteSetting) GetTuntermMonitoringDisabledOk() (*bool, bool)`

GetTuntermMonitoringDisabledOk returns a tuple with the TuntermMonitoringDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermMonitoringDisabled

`func (o *SiteSetting) SetTuntermMonitoringDisabled(v bool)`

SetTuntermMonitoringDisabled sets TuntermMonitoringDisabled field to given value.

### HasTuntermMonitoringDisabled

`func (o *SiteSetting) HasTuntermMonitoringDisabled() bool`

HasTuntermMonitoringDisabled returns a boolean if a field has been set.

### GetTuntermMulticastConfig

`func (o *SiteSetting) GetTuntermMulticastConfig() SiteSettingTuntermMulticastConfig`

GetTuntermMulticastConfig returns the TuntermMulticastConfig field if non-nil, zero value otherwise.

### GetTuntermMulticastConfigOk

`func (o *SiteSetting) GetTuntermMulticastConfigOk() (*SiteSettingTuntermMulticastConfig, bool)`

GetTuntermMulticastConfigOk returns a tuple with the TuntermMulticastConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermMulticastConfig

`func (o *SiteSetting) SetTuntermMulticastConfig(v SiteSettingTuntermMulticastConfig)`

SetTuntermMulticastConfig sets TuntermMulticastConfig field to given value.

### HasTuntermMulticastConfig

`func (o *SiteSetting) HasTuntermMulticastConfig() bool`

HasTuntermMulticastConfig returns a boolean if a field has been set.

### GetVars

`func (o *SiteSetting) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *SiteSetting) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *SiteSetting) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *SiteSetting) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetVna

`func (o *SiteSetting) GetVna() SiteSettingVna`

GetVna returns the Vna field if non-nil, zero value otherwise.

### GetVnaOk

`func (o *SiteSetting) GetVnaOk() (*SiteSettingVna, bool)`

GetVnaOk returns a tuple with the Vna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVna

`func (o *SiteSetting) SetVna(v SiteSettingVna)`

SetVna sets Vna field to given value.

### HasVna

`func (o *SiteSetting) HasVna() bool`

HasVna returns a boolean if a field has been set.

### GetVrfConfig

`func (o *SiteSetting) GetVrfConfig() VrfConfig`

GetVrfConfig returns the VrfConfig field if non-nil, zero value otherwise.

### GetVrfConfigOk

`func (o *SiteSetting) GetVrfConfigOk() (*VrfConfig, bool)`

GetVrfConfigOk returns a tuple with the VrfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfConfig

`func (o *SiteSetting) SetVrfConfig(v VrfConfig)`

SetVrfConfig sets VrfConfig field to given value.

### HasVrfConfig

`func (o *SiteSetting) HasVrfConfig() bool`

HasVrfConfig returns a boolean if a field has been set.

### GetVrfInstances

`func (o *SiteSetting) GetVrfInstances() map[string]VrfInstance`

GetVrfInstances returns the VrfInstances field if non-nil, zero value otherwise.

### GetVrfInstancesOk

`func (o *SiteSetting) GetVrfInstancesOk() (*map[string]VrfInstance, bool)`

GetVrfInstancesOk returns a tuple with the VrfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfInstances

`func (o *SiteSetting) SetVrfInstances(v map[string]VrfInstance)`

SetVrfInstances sets VrfInstances field to given value.

### HasVrfInstances

`func (o *SiteSetting) HasVrfInstances() bool`

HasVrfInstances returns a boolean if a field has been set.

### GetVrrpGroups

`func (o *SiteSetting) GetVrrpGroups() map[string]VrrpGroup`

GetVrrpGroups returns the VrrpGroups field if non-nil, zero value otherwise.

### GetVrrpGroupsOk

`func (o *SiteSetting) GetVrrpGroupsOk() (*map[string]VrrpGroup, bool)`

GetVrrpGroupsOk returns a tuple with the VrrpGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpGroups

`func (o *SiteSetting) SetVrrpGroups(v map[string]VrrpGroup)`

SetVrrpGroups sets VrrpGroups field to given value.

### HasVrrpGroups

`func (o *SiteSetting) HasVrrpGroups() bool`

HasVrrpGroups returns a boolean if a field has been set.

### GetWanVna

`func (o *SiteSetting) GetWanVna() SiteSettingWanVna`

GetWanVna returns the WanVna field if non-nil, zero value otherwise.

### GetWanVnaOk

`func (o *SiteSetting) GetWanVnaOk() (*SiteSettingWanVna, bool)`

GetWanVnaOk returns a tuple with the WanVna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanVna

`func (o *SiteSetting) SetWanVna(v SiteSettingWanVna)`

SetWanVna sets WanVna field to given value.

### HasWanVna

`func (o *SiteSetting) HasWanVna() bool`

HasWanVna returns a boolean if a field has been set.

### GetWatchedStationUrl

`func (o *SiteSetting) GetWatchedStationUrl() string`

GetWatchedStationUrl returns the WatchedStationUrl field if non-nil, zero value otherwise.

### GetWatchedStationUrlOk

`func (o *SiteSetting) GetWatchedStationUrlOk() (*string, bool)`

GetWatchedStationUrlOk returns a tuple with the WatchedStationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchedStationUrl

`func (o *SiteSetting) SetWatchedStationUrl(v string)`

SetWatchedStationUrl sets WatchedStationUrl field to given value.

### HasWatchedStationUrl

`func (o *SiteSetting) HasWatchedStationUrl() bool`

HasWatchedStationUrl returns a boolean if a field has been set.

### GetWhitelistUrl

`func (o *SiteSetting) GetWhitelistUrl() string`

GetWhitelistUrl returns the WhitelistUrl field if non-nil, zero value otherwise.

### GetWhitelistUrlOk

`func (o *SiteSetting) GetWhitelistUrlOk() (*string, bool)`

GetWhitelistUrlOk returns a tuple with the WhitelistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistUrl

`func (o *SiteSetting) SetWhitelistUrl(v string)`

SetWhitelistUrl sets WhitelistUrl field to given value.

### HasWhitelistUrl

`func (o *SiteSetting) HasWhitelistUrl() bool`

HasWhitelistUrl returns a boolean if a field has been set.

### GetWids

`func (o *SiteSetting) GetWids() SiteWids`

GetWids returns the Wids field if non-nil, zero value otherwise.

### GetWidsOk

`func (o *SiteSetting) GetWidsOk() (*SiteWids, bool)`

GetWidsOk returns a tuple with the Wids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWids

`func (o *SiteSetting) SetWids(v SiteWids)`

SetWids sets Wids field to given value.

### HasWids

`func (o *SiteSetting) HasWids() bool`

HasWids returns a boolean if a field has been set.

### GetWifi

`func (o *SiteSetting) GetWifi() SiteWifi`

GetWifi returns the Wifi field if non-nil, zero value otherwise.

### GetWifiOk

`func (o *SiteSetting) GetWifiOk() (*SiteWifi, bool)`

GetWifiOk returns a tuple with the Wifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi

`func (o *SiteSetting) SetWifi(v SiteWifi)`

SetWifi sets Wifi field to given value.

### HasWifi

`func (o *SiteSetting) HasWifi() bool`

HasWifi returns a boolean if a field has been set.

### GetWiredVna

`func (o *SiteSetting) GetWiredVna() SiteSettingWiredVna`

GetWiredVna returns the WiredVna field if non-nil, zero value otherwise.

### GetWiredVnaOk

`func (o *SiteSetting) GetWiredVnaOk() (*SiteSettingWiredVna, bool)`

GetWiredVnaOk returns a tuple with the WiredVna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredVna

`func (o *SiteSetting) SetWiredVna(v SiteSettingWiredVna)`

SetWiredVna sets WiredVna field to given value.

### HasWiredVna

`func (o *SiteSetting) HasWiredVna() bool`

HasWiredVna returns a boolean if a field has been set.

### GetZoneOccupancyAlert

`func (o *SiteSetting) GetZoneOccupancyAlert() SiteZoneOccupancyAlert`

GetZoneOccupancyAlert returns the ZoneOccupancyAlert field if non-nil, zero value otherwise.

### GetZoneOccupancyAlertOk

`func (o *SiteSetting) GetZoneOccupancyAlertOk() (*SiteZoneOccupancyAlert, bool)`

GetZoneOccupancyAlertOk returns a tuple with the ZoneOccupancyAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneOccupancyAlert

`func (o *SiteSetting) SetZoneOccupancyAlert(v SiteZoneOccupancyAlert)`

SetZoneOccupancyAlert sets ZoneOccupancyAlert field to given value.

### HasZoneOccupancyAlert

`func (o *SiteSetting) HasZoneOccupancyAlert() bool`

HasZoneOccupancyAlert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



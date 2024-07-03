# GatewayPortConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisableAutoneg** | Pointer to **bool** |  | [optional] [default to false]
**Disabled** | Pointer to **bool** | port admin up (true) / down (false) | [optional] [default to false]
**DslType** | Pointer to [**GatewayPortDslType**](GatewayPortDslType.md) |  | [optional] [default to GATEWAYPORTDSLTYPE_VDSL]
**DslVci** | Pointer to **int32** | if &#x60;wan_type&#x60;&#x3D;&#x3D;&#x60;dsl&#x60; 16 bit int | [optional] [default to 35]
**DslVpi** | Pointer to **int32** | if &#x60;wan_type&#x60;&#x3D;&#x3D;&#x60;dsl&#x60; 8 bit int | [optional] [default to 0]
**Duplex** | Pointer to [**GatewayPortDuplex**](GatewayPortDuplex.md) |  | [optional] [default to GATEWAYPORTDUPLEX_AUTO]
**IpConfig** | Pointer to [**GatewayIpConfig**](GatewayIpConfig.md) |  | [optional] 
**LteApn** | Pointer to **string** | if &#x60;wan_type&#x60;&#x3D;&#x3D;&#x60;lte&#x60; | [optional] 
**LteAuth** | Pointer to [**GatewayPortLteAuth**](GatewayPortLteAuth.md) |  | [optional] [default to GATEWAYPORTLTEAUTH_NONE]
**LteBackup** | Pointer to **bool** |  | [optional] 
**LtePassword** | Pointer to **string** | if &#x60;wan_type&#x60;&#x3D;&#x3D;&#x60;lte&#x60; | [optional] 
**LteUsername** | Pointer to **string** | if &#x60;wan_type&#x60;&#x3D;&#x3D;&#x60;lte&#x60; | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | name that we&#39;ll use to derive config | [optional] 
**Networks** | Pointer to **[]string** | if &#x60;usage&#x60;&#x3D;&#x3D;&#x60;lan&#x60; | [optional] 
**OuterVlanId** | Pointer to **int32** | for Q-in-Q | [optional] 
**PoeDisabled** | Pointer to **bool** |  | [optional] [default to false]
**PortNetwork** | Pointer to **string** | if &#x60;usage&#x60;&#x3D;&#x3D;&#x60;lan&#x60; | [optional] 
**PreserveDscp** | Pointer to **bool** | whether to preserve dscp when sending traffic over VPN (SSR-only) | [optional] [default to true]
**Redundant** | Pointer to **bool** | if HA mode | [optional] 
**RethIdx** | Pointer to **int32** | if HA mode | [optional] 
**RethNode** | Pointer to **string** | if HA mode | [optional] 
**RethNodes** | Pointer to **[]string** | SSR only - supporting vlan-based redundancy (matching the size of &#x60;networks&#x60;) | [optional] 
**Speed** | Pointer to **string** |  | [optional] [default to "auto"]
**SsrNoVirtualMac** | Pointer to **bool** | when SSR is running as VM, this is required on certain hosting platforms | [optional] [default to false]
**SvrPortRange** | Pointer to **string** | for SSR only | [optional] [default to "none"]
**TrafficShaping** | Pointer to [**GatewayTrafficShaping**](GatewayTrafficShaping.md) |  | [optional] 
**Usage** | [**GatewayPortUsage**](GatewayPortUsage.md) |  | 
**VlanId** | Pointer to **int32** | if WAN interface is on a VLAN | [optional] 
**VpnPaths** | Pointer to [**map[string]GatewayPortVpnPath**](GatewayPortVpnPath.md) |  | [optional] 
**WanArpPolicer** | Pointer to [**GatewayPortWanArpPolicer**](GatewayPortWanArpPolicer.md) |  | [optional] [default to GATEWAYPORTWANARPPOLICER_RECOMMENDED]
**WanExtIp** | Pointer to **string** | optional, if spoke should reach this port by a different IP | [optional] 
**WanSourceNat** | Pointer to [**GatewayPortWanSourceNat**](GatewayPortWanSourceNat.md) |  | [optional] 
**WanType** | Pointer to [**GatewayPortWanType**](GatewayPortWanType.md) |  | [optional] [default to GATEWAYPORTWANTYPE_BROADBAND]

## Methods

### NewGatewayPortConfig

`func NewGatewayPortConfig(usage GatewayPortUsage, ) *GatewayPortConfig`

NewGatewayPortConfig instantiates a new GatewayPortConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortConfigWithDefaults

`func NewGatewayPortConfigWithDefaults() *GatewayPortConfig`

NewGatewayPortConfigWithDefaults instantiates a new GatewayPortConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GatewayPortConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayPortConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayPortConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayPortConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableAutoneg

`func (o *GatewayPortConfig) GetDisableAutoneg() bool`

GetDisableAutoneg returns the DisableAutoneg field if non-nil, zero value otherwise.

### GetDisableAutonegOk

`func (o *GatewayPortConfig) GetDisableAutonegOk() (*bool, bool)`

GetDisableAutonegOk returns a tuple with the DisableAutoneg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoneg

`func (o *GatewayPortConfig) SetDisableAutoneg(v bool)`

SetDisableAutoneg sets DisableAutoneg field to given value.

### HasDisableAutoneg

`func (o *GatewayPortConfig) HasDisableAutoneg() bool`

HasDisableAutoneg returns a boolean if a field has been set.

### GetDisabled

`func (o *GatewayPortConfig) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GatewayPortConfig) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GatewayPortConfig) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GatewayPortConfig) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDslType

`func (o *GatewayPortConfig) GetDslType() GatewayPortDslType`

GetDslType returns the DslType field if non-nil, zero value otherwise.

### GetDslTypeOk

`func (o *GatewayPortConfig) GetDslTypeOk() (*GatewayPortDslType, bool)`

GetDslTypeOk returns a tuple with the DslType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslType

`func (o *GatewayPortConfig) SetDslType(v GatewayPortDslType)`

SetDslType sets DslType field to given value.

### HasDslType

`func (o *GatewayPortConfig) HasDslType() bool`

HasDslType returns a boolean if a field has been set.

### GetDslVci

`func (o *GatewayPortConfig) GetDslVci() int32`

GetDslVci returns the DslVci field if non-nil, zero value otherwise.

### GetDslVciOk

`func (o *GatewayPortConfig) GetDslVciOk() (*int32, bool)`

GetDslVciOk returns a tuple with the DslVci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslVci

`func (o *GatewayPortConfig) SetDslVci(v int32)`

SetDslVci sets DslVci field to given value.

### HasDslVci

`func (o *GatewayPortConfig) HasDslVci() bool`

HasDslVci returns a boolean if a field has been set.

### GetDslVpi

`func (o *GatewayPortConfig) GetDslVpi() int32`

GetDslVpi returns the DslVpi field if non-nil, zero value otherwise.

### GetDslVpiOk

`func (o *GatewayPortConfig) GetDslVpiOk() (*int32, bool)`

GetDslVpiOk returns a tuple with the DslVpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslVpi

`func (o *GatewayPortConfig) SetDslVpi(v int32)`

SetDslVpi sets DslVpi field to given value.

### HasDslVpi

`func (o *GatewayPortConfig) HasDslVpi() bool`

HasDslVpi returns a boolean if a field has been set.

### GetDuplex

`func (o *GatewayPortConfig) GetDuplex() GatewayPortDuplex`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *GatewayPortConfig) GetDuplexOk() (*GatewayPortDuplex, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *GatewayPortConfig) SetDuplex(v GatewayPortDuplex)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *GatewayPortConfig) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetIpConfig

`func (o *GatewayPortConfig) GetIpConfig() GatewayIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *GatewayPortConfig) GetIpConfigOk() (*GatewayIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *GatewayPortConfig) SetIpConfig(v GatewayIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *GatewayPortConfig) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetLteApn

`func (o *GatewayPortConfig) GetLteApn() string`

GetLteApn returns the LteApn field if non-nil, zero value otherwise.

### GetLteApnOk

`func (o *GatewayPortConfig) GetLteApnOk() (*string, bool)`

GetLteApnOk returns a tuple with the LteApn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteApn

`func (o *GatewayPortConfig) SetLteApn(v string)`

SetLteApn sets LteApn field to given value.

### HasLteApn

`func (o *GatewayPortConfig) HasLteApn() bool`

HasLteApn returns a boolean if a field has been set.

### GetLteAuth

`func (o *GatewayPortConfig) GetLteAuth() GatewayPortLteAuth`

GetLteAuth returns the LteAuth field if non-nil, zero value otherwise.

### GetLteAuthOk

`func (o *GatewayPortConfig) GetLteAuthOk() (*GatewayPortLteAuth, bool)`

GetLteAuthOk returns a tuple with the LteAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteAuth

`func (o *GatewayPortConfig) SetLteAuth(v GatewayPortLteAuth)`

SetLteAuth sets LteAuth field to given value.

### HasLteAuth

`func (o *GatewayPortConfig) HasLteAuth() bool`

HasLteAuth returns a boolean if a field has been set.

### GetLteBackup

`func (o *GatewayPortConfig) GetLteBackup() bool`

GetLteBackup returns the LteBackup field if non-nil, zero value otherwise.

### GetLteBackupOk

`func (o *GatewayPortConfig) GetLteBackupOk() (*bool, bool)`

GetLteBackupOk returns a tuple with the LteBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteBackup

`func (o *GatewayPortConfig) SetLteBackup(v bool)`

SetLteBackup sets LteBackup field to given value.

### HasLteBackup

`func (o *GatewayPortConfig) HasLteBackup() bool`

HasLteBackup returns a boolean if a field has been set.

### GetLtePassword

`func (o *GatewayPortConfig) GetLtePassword() string`

GetLtePassword returns the LtePassword field if non-nil, zero value otherwise.

### GetLtePasswordOk

`func (o *GatewayPortConfig) GetLtePasswordOk() (*string, bool)`

GetLtePasswordOk returns a tuple with the LtePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtePassword

`func (o *GatewayPortConfig) SetLtePassword(v string)`

SetLtePassword sets LtePassword field to given value.

### HasLtePassword

`func (o *GatewayPortConfig) HasLtePassword() bool`

HasLtePassword returns a boolean if a field has been set.

### GetLteUsername

`func (o *GatewayPortConfig) GetLteUsername() string`

GetLteUsername returns the LteUsername field if non-nil, zero value otherwise.

### GetLteUsernameOk

`func (o *GatewayPortConfig) GetLteUsernameOk() (*string, bool)`

GetLteUsernameOk returns a tuple with the LteUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteUsername

`func (o *GatewayPortConfig) SetLteUsername(v string)`

SetLteUsername sets LteUsername field to given value.

### HasLteUsername

`func (o *GatewayPortConfig) HasLteUsername() bool`

HasLteUsername returns a boolean if a field has been set.

### GetMtu

`func (o *GatewayPortConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *GatewayPortConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *GatewayPortConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *GatewayPortConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *GatewayPortConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayPortConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayPortConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayPortConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *GatewayPortConfig) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GatewayPortConfig) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GatewayPortConfig) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GatewayPortConfig) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOuterVlanId

`func (o *GatewayPortConfig) GetOuterVlanId() int32`

GetOuterVlanId returns the OuterVlanId field if non-nil, zero value otherwise.

### GetOuterVlanIdOk

`func (o *GatewayPortConfig) GetOuterVlanIdOk() (*int32, bool)`

GetOuterVlanIdOk returns a tuple with the OuterVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterVlanId

`func (o *GatewayPortConfig) SetOuterVlanId(v int32)`

SetOuterVlanId sets OuterVlanId field to given value.

### HasOuterVlanId

`func (o *GatewayPortConfig) HasOuterVlanId() bool`

HasOuterVlanId returns a boolean if a field has been set.

### GetPoeDisabled

`func (o *GatewayPortConfig) GetPoeDisabled() bool`

GetPoeDisabled returns the PoeDisabled field if non-nil, zero value otherwise.

### GetPoeDisabledOk

`func (o *GatewayPortConfig) GetPoeDisabledOk() (*bool, bool)`

GetPoeDisabledOk returns a tuple with the PoeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisabled

`func (o *GatewayPortConfig) SetPoeDisabled(v bool)`

SetPoeDisabled sets PoeDisabled field to given value.

### HasPoeDisabled

`func (o *GatewayPortConfig) HasPoeDisabled() bool`

HasPoeDisabled returns a boolean if a field has been set.

### GetPortNetwork

`func (o *GatewayPortConfig) GetPortNetwork() string`

GetPortNetwork returns the PortNetwork field if non-nil, zero value otherwise.

### GetPortNetworkOk

`func (o *GatewayPortConfig) GetPortNetworkOk() (*string, bool)`

GetPortNetworkOk returns a tuple with the PortNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNetwork

`func (o *GatewayPortConfig) SetPortNetwork(v string)`

SetPortNetwork sets PortNetwork field to given value.

### HasPortNetwork

`func (o *GatewayPortConfig) HasPortNetwork() bool`

HasPortNetwork returns a boolean if a field has been set.

### GetPreserveDscp

`func (o *GatewayPortConfig) GetPreserveDscp() bool`

GetPreserveDscp returns the PreserveDscp field if non-nil, zero value otherwise.

### GetPreserveDscpOk

`func (o *GatewayPortConfig) GetPreserveDscpOk() (*bool, bool)`

GetPreserveDscpOk returns a tuple with the PreserveDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveDscp

`func (o *GatewayPortConfig) SetPreserveDscp(v bool)`

SetPreserveDscp sets PreserveDscp field to given value.

### HasPreserveDscp

`func (o *GatewayPortConfig) HasPreserveDscp() bool`

HasPreserveDscp returns a boolean if a field has been set.

### GetRedundant

`func (o *GatewayPortConfig) GetRedundant() bool`

GetRedundant returns the Redundant field if non-nil, zero value otherwise.

### GetRedundantOk

`func (o *GatewayPortConfig) GetRedundantOk() (*bool, bool)`

GetRedundantOk returns a tuple with the Redundant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundant

`func (o *GatewayPortConfig) SetRedundant(v bool)`

SetRedundant sets Redundant field to given value.

### HasRedundant

`func (o *GatewayPortConfig) HasRedundant() bool`

HasRedundant returns a boolean if a field has been set.

### GetRethIdx

`func (o *GatewayPortConfig) GetRethIdx() int32`

GetRethIdx returns the RethIdx field if non-nil, zero value otherwise.

### GetRethIdxOk

`func (o *GatewayPortConfig) GetRethIdxOk() (*int32, bool)`

GetRethIdxOk returns a tuple with the RethIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRethIdx

`func (o *GatewayPortConfig) SetRethIdx(v int32)`

SetRethIdx sets RethIdx field to given value.

### HasRethIdx

`func (o *GatewayPortConfig) HasRethIdx() bool`

HasRethIdx returns a boolean if a field has been set.

### GetRethNode

`func (o *GatewayPortConfig) GetRethNode() string`

GetRethNode returns the RethNode field if non-nil, zero value otherwise.

### GetRethNodeOk

`func (o *GatewayPortConfig) GetRethNodeOk() (*string, bool)`

GetRethNodeOk returns a tuple with the RethNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRethNode

`func (o *GatewayPortConfig) SetRethNode(v string)`

SetRethNode sets RethNode field to given value.

### HasRethNode

`func (o *GatewayPortConfig) HasRethNode() bool`

HasRethNode returns a boolean if a field has been set.

### GetRethNodes

`func (o *GatewayPortConfig) GetRethNodes() []string`

GetRethNodes returns the RethNodes field if non-nil, zero value otherwise.

### GetRethNodesOk

`func (o *GatewayPortConfig) GetRethNodesOk() (*[]string, bool)`

GetRethNodesOk returns a tuple with the RethNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRethNodes

`func (o *GatewayPortConfig) SetRethNodes(v []string)`

SetRethNodes sets RethNodes field to given value.

### HasRethNodes

`func (o *GatewayPortConfig) HasRethNodes() bool`

HasRethNodes returns a boolean if a field has been set.

### GetSpeed

`func (o *GatewayPortConfig) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GatewayPortConfig) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GatewayPortConfig) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GatewayPortConfig) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSsrNoVirtualMac

`func (o *GatewayPortConfig) GetSsrNoVirtualMac() bool`

GetSsrNoVirtualMac returns the SsrNoVirtualMac field if non-nil, zero value otherwise.

### GetSsrNoVirtualMacOk

`func (o *GatewayPortConfig) GetSsrNoVirtualMacOk() (*bool, bool)`

GetSsrNoVirtualMacOk returns a tuple with the SsrNoVirtualMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsrNoVirtualMac

`func (o *GatewayPortConfig) SetSsrNoVirtualMac(v bool)`

SetSsrNoVirtualMac sets SsrNoVirtualMac field to given value.

### HasSsrNoVirtualMac

`func (o *GatewayPortConfig) HasSsrNoVirtualMac() bool`

HasSsrNoVirtualMac returns a boolean if a field has been set.

### GetSvrPortRange

`func (o *GatewayPortConfig) GetSvrPortRange() string`

GetSvrPortRange returns the SvrPortRange field if non-nil, zero value otherwise.

### GetSvrPortRangeOk

`func (o *GatewayPortConfig) GetSvrPortRangeOk() (*string, bool)`

GetSvrPortRangeOk returns a tuple with the SvrPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvrPortRange

`func (o *GatewayPortConfig) SetSvrPortRange(v string)`

SetSvrPortRange sets SvrPortRange field to given value.

### HasSvrPortRange

`func (o *GatewayPortConfig) HasSvrPortRange() bool`

HasSvrPortRange returns a boolean if a field has been set.

### GetTrafficShaping

`func (o *GatewayPortConfig) GetTrafficShaping() GatewayTrafficShaping`

GetTrafficShaping returns the TrafficShaping field if non-nil, zero value otherwise.

### GetTrafficShapingOk

`func (o *GatewayPortConfig) GetTrafficShapingOk() (*GatewayTrafficShaping, bool)`

GetTrafficShapingOk returns a tuple with the TrafficShaping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShaping

`func (o *GatewayPortConfig) SetTrafficShaping(v GatewayTrafficShaping)`

SetTrafficShaping sets TrafficShaping field to given value.

### HasTrafficShaping

`func (o *GatewayPortConfig) HasTrafficShaping() bool`

HasTrafficShaping returns a boolean if a field has been set.

### GetUsage

`func (o *GatewayPortConfig) GetUsage() GatewayPortUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GatewayPortConfig) GetUsageOk() (*GatewayPortUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GatewayPortConfig) SetUsage(v GatewayPortUsage)`

SetUsage sets Usage field to given value.


### GetVlanId

`func (o *GatewayPortConfig) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GatewayPortConfig) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GatewayPortConfig) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *GatewayPortConfig) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVpnPaths

`func (o *GatewayPortConfig) GetVpnPaths() map[string]GatewayPortVpnPath`

GetVpnPaths returns the VpnPaths field if non-nil, zero value otherwise.

### GetVpnPathsOk

`func (o *GatewayPortConfig) GetVpnPathsOk() (*map[string]GatewayPortVpnPath, bool)`

GetVpnPathsOk returns a tuple with the VpnPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPaths

`func (o *GatewayPortConfig) SetVpnPaths(v map[string]GatewayPortVpnPath)`

SetVpnPaths sets VpnPaths field to given value.

### HasVpnPaths

`func (o *GatewayPortConfig) HasVpnPaths() bool`

HasVpnPaths returns a boolean if a field has been set.

### GetWanArpPolicer

`func (o *GatewayPortConfig) GetWanArpPolicer() GatewayPortWanArpPolicer`

GetWanArpPolicer returns the WanArpPolicer field if non-nil, zero value otherwise.

### GetWanArpPolicerOk

`func (o *GatewayPortConfig) GetWanArpPolicerOk() (*GatewayPortWanArpPolicer, bool)`

GetWanArpPolicerOk returns a tuple with the WanArpPolicer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanArpPolicer

`func (o *GatewayPortConfig) SetWanArpPolicer(v GatewayPortWanArpPolicer)`

SetWanArpPolicer sets WanArpPolicer field to given value.

### HasWanArpPolicer

`func (o *GatewayPortConfig) HasWanArpPolicer() bool`

HasWanArpPolicer returns a boolean if a field has been set.

### GetWanExtIp

`func (o *GatewayPortConfig) GetWanExtIp() string`

GetWanExtIp returns the WanExtIp field if non-nil, zero value otherwise.

### GetWanExtIpOk

`func (o *GatewayPortConfig) GetWanExtIpOk() (*string, bool)`

GetWanExtIpOk returns a tuple with the WanExtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanExtIp

`func (o *GatewayPortConfig) SetWanExtIp(v string)`

SetWanExtIp sets WanExtIp field to given value.

### HasWanExtIp

`func (o *GatewayPortConfig) HasWanExtIp() bool`

HasWanExtIp returns a boolean if a field has been set.

### GetWanSourceNat

`func (o *GatewayPortConfig) GetWanSourceNat() GatewayPortWanSourceNat`

GetWanSourceNat returns the WanSourceNat field if non-nil, zero value otherwise.

### GetWanSourceNatOk

`func (o *GatewayPortConfig) GetWanSourceNatOk() (*GatewayPortWanSourceNat, bool)`

GetWanSourceNatOk returns a tuple with the WanSourceNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanSourceNat

`func (o *GatewayPortConfig) SetWanSourceNat(v GatewayPortWanSourceNat)`

SetWanSourceNat sets WanSourceNat field to given value.

### HasWanSourceNat

`func (o *GatewayPortConfig) HasWanSourceNat() bool`

HasWanSourceNat returns a boolean if a field has been set.

### GetWanType

`func (o *GatewayPortConfig) GetWanType() GatewayPortWanType`

GetWanType returns the WanType field if non-nil, zero value otherwise.

### GetWanTypeOk

`func (o *GatewayPortConfig) GetWanTypeOk() (*GatewayPortWanType, bool)`

GetWanTypeOk returns a tuple with the WanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanType

`func (o *GatewayPortConfig) SetWanType(v GatewayPortWanType)`

SetWanType sets WanType field to given value.

### HasWanType

`func (o *GatewayPortConfig) HasWanType() bool`

HasWanType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



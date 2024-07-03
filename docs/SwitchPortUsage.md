# SwitchPortUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllNetworks** | Pointer to **bool** | Only if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;trunk&#x60; whether to trunk all network/vlans | [optional] [default to false]
**AllowDhcpd** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; if DHCP snooping is enabled, whether DHCP server is allowed on the interfaces with. All the interfaces from port configs using this port usage are effected. Please notice that allow_dhcpd is a tri_state.  When it is not defined, it means using the system’s default setting which depends on whether the port is a access or trunk port. | [optional] 
**AllowMultipleSupplicants** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; | [optional] [default to false]
**BypassAuthWhenServerDown** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60; bypass auth for known clients if set to true when RADIUS server is down | [optional] [default to false]
**BypassAuthWhenServerDownForUnkonwnClient** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x60;dot1x&#x60; bypass auth for all (including unknown clients) if set to true when RADIUS server is down | [optional] [default to false]
**Description** | Pointer to **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; | [optional] 
**DisableAutoneg** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; if speed and duplex are specified, whether to disable autonegotiation | [optional] [default to false]
**Disabled** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; whether the port is disabled | [optional] [default to false]
**Duplex** | Pointer to [**SwitchPortUsageDuplex**](SwitchPortUsageDuplex.md) |  | [optional] [default to SWITCHPORTUSAGEDUPLEX_AUTO]
**DynamicVlanNetworks** | Pointer to **[]string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; if dynamic vlan is used, specify the possible networks/vlans RADIUS can return | [optional] 
**EnableMacAuth** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60; whether to enable MAC Auth | [optional] [default to false]
**EnableQos** | Pointer to **bool** |  | [optional] [default to false]
**GuestNetwork** | Pointer to **NullableString** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60; which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed) | [optional] 
**InterSwitchLink** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; inter_switch_link is used together with \&quot;isolation\&quot; under networks NOTE: inter_switch_link works only between Juniper device. This has to be applied to both ports connected together | [optional] [default to false]
**MacAuthOnly** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;enable_mac_auth&#x60;&#x3D;&#x3D;&#x60;true&#x60; | [optional] 
**MacAuthProtocol** | Pointer to [**SwitchPortUsageMacAuthProtocol**](SwitchPortUsageMacAuthProtocol.md) |  | [optional] [default to SWITCHPORTUSAGEMACAUTHPROTOCOL_EAP_MD5]
**MacLimit** | Pointer to **int32** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; max number of mac addresses, default is 0 for unlimited, otherwise range is 1 or higher, with upper bound constrained by platform | [optional] [default to 0]
**Mode** | Pointer to [**SwitchPortUsageMode**](SwitchPortUsageMode.md) |  | [optional] 
**Mtu** | Pointer to **int32** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514. | [optional] 
**Networks** | Pointer to **[]string** | Only if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;trunk&#x60;, the list of network/vlans | [optional] 
**PersistMac** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;mode&#x60;&#x3D;&#x3D;&#x60;access&#x60; and &#x60;port_auth&#x60;!&#x3D;&#x60;dot1x&#x60; whether the port should retain dynamically learned MAC addresses | [optional] [default to false]
**PoeDisabled** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; whether PoE capabilities are disabled for a port | [optional] [default to false]
**PortAuth** | Pointer to **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; if dot1x is desired, set to dot1x | [optional] 
**PortNetwork** | Pointer to **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; native network/vlan for untagged traffic | [optional] 
**ReauthInterval** | Pointer to **int32** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x60;dot1x&#x60; reauthentication interval range | [optional] [default to 3600]
**RejectedNetwork** | Pointer to **NullableString** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; and &#x60;port_auth&#x60;&#x3D;&#x3D;&#x60;dot1x&#x60; when radius server reject / fails | [optional] 
**ResetDefaultWhen** | Pointer to [**SwitchPortUsageDynamicResetDefaultWhen**](SwitchPortUsageDynamicResetDefaultWhen.md) |  | [optional] [default to SWITCHPORTUSAGEDYNAMICRESETDEFAULTWHEN_LINK_DOWN]
**Rules** | Pointer to [**[]SwitchPortUsageDynamicRule**](SwitchPortUsageDynamicRule.md) | Only if &#x60;mode&#x60;&#x3D;&#x3D;&#x60;dynamic&#x60; | [optional] 
**Speed** | Pointer to **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; speed, default is auto to automatically negotiate speed | [optional] 
**StormControl** | Pointer to [**SwitchPortUsageStormControl**](SwitchPortUsageStormControl.md) |  | [optional] 
**StpEdge** | Pointer to **bool** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; when enabled, the port is not expected to receive BPDU frames | [optional] [default to false]
**VoipNetwork** | Pointer to **string** | Only if &#x60;mode&#x60;!&#x3D;&#x60;dynamic&#x60; network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth | [optional] 

## Methods

### NewSwitchPortUsage

`func NewSwitchPortUsage() *SwitchPortUsage`

NewSwitchPortUsage instantiates a new SwitchPortUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortUsageWithDefaults

`func NewSwitchPortUsageWithDefaults() *SwitchPortUsage`

NewSwitchPortUsageWithDefaults instantiates a new SwitchPortUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllNetworks

`func (o *SwitchPortUsage) GetAllNetworks() bool`

GetAllNetworks returns the AllNetworks field if non-nil, zero value otherwise.

### GetAllNetworksOk

`func (o *SwitchPortUsage) GetAllNetworksOk() (*bool, bool)`

GetAllNetworksOk returns a tuple with the AllNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllNetworks

`func (o *SwitchPortUsage) SetAllNetworks(v bool)`

SetAllNetworks sets AllNetworks field to given value.

### HasAllNetworks

`func (o *SwitchPortUsage) HasAllNetworks() bool`

HasAllNetworks returns a boolean if a field has been set.

### GetAllowDhcpd

`func (o *SwitchPortUsage) GetAllowDhcpd() bool`

GetAllowDhcpd returns the AllowDhcpd field if non-nil, zero value otherwise.

### GetAllowDhcpdOk

`func (o *SwitchPortUsage) GetAllowDhcpdOk() (*bool, bool)`

GetAllowDhcpdOk returns a tuple with the AllowDhcpd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDhcpd

`func (o *SwitchPortUsage) SetAllowDhcpd(v bool)`

SetAllowDhcpd sets AllowDhcpd field to given value.

### HasAllowDhcpd

`func (o *SwitchPortUsage) HasAllowDhcpd() bool`

HasAllowDhcpd returns a boolean if a field has been set.

### GetAllowMultipleSupplicants

`func (o *SwitchPortUsage) GetAllowMultipleSupplicants() bool`

GetAllowMultipleSupplicants returns the AllowMultipleSupplicants field if non-nil, zero value otherwise.

### GetAllowMultipleSupplicantsOk

`func (o *SwitchPortUsage) GetAllowMultipleSupplicantsOk() (*bool, bool)`

GetAllowMultipleSupplicantsOk returns a tuple with the AllowMultipleSupplicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSupplicants

`func (o *SwitchPortUsage) SetAllowMultipleSupplicants(v bool)`

SetAllowMultipleSupplicants sets AllowMultipleSupplicants field to given value.

### HasAllowMultipleSupplicants

`func (o *SwitchPortUsage) HasAllowMultipleSupplicants() bool`

HasAllowMultipleSupplicants returns a boolean if a field has been set.

### GetBypassAuthWhenServerDown

`func (o *SwitchPortUsage) GetBypassAuthWhenServerDown() bool`

GetBypassAuthWhenServerDown returns the BypassAuthWhenServerDown field if non-nil, zero value otherwise.

### GetBypassAuthWhenServerDownOk

`func (o *SwitchPortUsage) GetBypassAuthWhenServerDownOk() (*bool, bool)`

GetBypassAuthWhenServerDownOk returns a tuple with the BypassAuthWhenServerDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAuthWhenServerDown

`func (o *SwitchPortUsage) SetBypassAuthWhenServerDown(v bool)`

SetBypassAuthWhenServerDown sets BypassAuthWhenServerDown field to given value.

### HasBypassAuthWhenServerDown

`func (o *SwitchPortUsage) HasBypassAuthWhenServerDown() bool`

HasBypassAuthWhenServerDown returns a boolean if a field has been set.

### GetBypassAuthWhenServerDownForUnkonwnClient

`func (o *SwitchPortUsage) GetBypassAuthWhenServerDownForUnkonwnClient() bool`

GetBypassAuthWhenServerDownForUnkonwnClient returns the BypassAuthWhenServerDownForUnkonwnClient field if non-nil, zero value otherwise.

### GetBypassAuthWhenServerDownForUnkonwnClientOk

`func (o *SwitchPortUsage) GetBypassAuthWhenServerDownForUnkonwnClientOk() (*bool, bool)`

GetBypassAuthWhenServerDownForUnkonwnClientOk returns a tuple with the BypassAuthWhenServerDownForUnkonwnClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAuthWhenServerDownForUnkonwnClient

`func (o *SwitchPortUsage) SetBypassAuthWhenServerDownForUnkonwnClient(v bool)`

SetBypassAuthWhenServerDownForUnkonwnClient sets BypassAuthWhenServerDownForUnkonwnClient field to given value.

### HasBypassAuthWhenServerDownForUnkonwnClient

`func (o *SwitchPortUsage) HasBypassAuthWhenServerDownForUnkonwnClient() bool`

HasBypassAuthWhenServerDownForUnkonwnClient returns a boolean if a field has been set.

### GetDescription

`func (o *SwitchPortUsage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SwitchPortUsage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SwitchPortUsage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SwitchPortUsage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableAutoneg

`func (o *SwitchPortUsage) GetDisableAutoneg() bool`

GetDisableAutoneg returns the DisableAutoneg field if non-nil, zero value otherwise.

### GetDisableAutonegOk

`func (o *SwitchPortUsage) GetDisableAutonegOk() (*bool, bool)`

GetDisableAutonegOk returns a tuple with the DisableAutoneg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoneg

`func (o *SwitchPortUsage) SetDisableAutoneg(v bool)`

SetDisableAutoneg sets DisableAutoneg field to given value.

### HasDisableAutoneg

`func (o *SwitchPortUsage) HasDisableAutoneg() bool`

HasDisableAutoneg returns a boolean if a field has been set.

### GetDisabled

`func (o *SwitchPortUsage) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SwitchPortUsage) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SwitchPortUsage) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SwitchPortUsage) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDuplex

`func (o *SwitchPortUsage) GetDuplex() SwitchPortUsageDuplex`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *SwitchPortUsage) GetDuplexOk() (*SwitchPortUsageDuplex, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *SwitchPortUsage) SetDuplex(v SwitchPortUsageDuplex)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *SwitchPortUsage) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetDynamicVlanNetworks

`func (o *SwitchPortUsage) GetDynamicVlanNetworks() []string`

GetDynamicVlanNetworks returns the DynamicVlanNetworks field if non-nil, zero value otherwise.

### GetDynamicVlanNetworksOk

`func (o *SwitchPortUsage) GetDynamicVlanNetworksOk() (*[]string, bool)`

GetDynamicVlanNetworksOk returns a tuple with the DynamicVlanNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVlanNetworks

`func (o *SwitchPortUsage) SetDynamicVlanNetworks(v []string)`

SetDynamicVlanNetworks sets DynamicVlanNetworks field to given value.

### HasDynamicVlanNetworks

`func (o *SwitchPortUsage) HasDynamicVlanNetworks() bool`

HasDynamicVlanNetworks returns a boolean if a field has been set.

### GetEnableMacAuth

`func (o *SwitchPortUsage) GetEnableMacAuth() bool`

GetEnableMacAuth returns the EnableMacAuth field if non-nil, zero value otherwise.

### GetEnableMacAuthOk

`func (o *SwitchPortUsage) GetEnableMacAuthOk() (*bool, bool)`

GetEnableMacAuthOk returns a tuple with the EnableMacAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMacAuth

`func (o *SwitchPortUsage) SetEnableMacAuth(v bool)`

SetEnableMacAuth sets EnableMacAuth field to given value.

### HasEnableMacAuth

`func (o *SwitchPortUsage) HasEnableMacAuth() bool`

HasEnableMacAuth returns a boolean if a field has been set.

### GetEnableQos

`func (o *SwitchPortUsage) GetEnableQos() bool`

GetEnableQos returns the EnableQos field if non-nil, zero value otherwise.

### GetEnableQosOk

`func (o *SwitchPortUsage) GetEnableQosOk() (*bool, bool)`

GetEnableQosOk returns a tuple with the EnableQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQos

`func (o *SwitchPortUsage) SetEnableQos(v bool)`

SetEnableQos sets EnableQos field to given value.

### HasEnableQos

`func (o *SwitchPortUsage) HasEnableQos() bool`

HasEnableQos returns a boolean if a field has been set.

### GetGuestNetwork

`func (o *SwitchPortUsage) GetGuestNetwork() string`

GetGuestNetwork returns the GuestNetwork field if non-nil, zero value otherwise.

### GetGuestNetworkOk

`func (o *SwitchPortUsage) GetGuestNetworkOk() (*string, bool)`

GetGuestNetworkOk returns a tuple with the GuestNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNetwork

`func (o *SwitchPortUsage) SetGuestNetwork(v string)`

SetGuestNetwork sets GuestNetwork field to given value.

### HasGuestNetwork

`func (o *SwitchPortUsage) HasGuestNetwork() bool`

HasGuestNetwork returns a boolean if a field has been set.

### SetGuestNetworkNil

`func (o *SwitchPortUsage) SetGuestNetworkNil(b bool)`

 SetGuestNetworkNil sets the value for GuestNetwork to be an explicit nil

### UnsetGuestNetwork
`func (o *SwitchPortUsage) UnsetGuestNetwork()`

UnsetGuestNetwork ensures that no value is present for GuestNetwork, not even an explicit nil
### GetInterSwitchLink

`func (o *SwitchPortUsage) GetInterSwitchLink() bool`

GetInterSwitchLink returns the InterSwitchLink field if non-nil, zero value otherwise.

### GetInterSwitchLinkOk

`func (o *SwitchPortUsage) GetInterSwitchLinkOk() (*bool, bool)`

GetInterSwitchLinkOk returns a tuple with the InterSwitchLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterSwitchLink

`func (o *SwitchPortUsage) SetInterSwitchLink(v bool)`

SetInterSwitchLink sets InterSwitchLink field to given value.

### HasInterSwitchLink

`func (o *SwitchPortUsage) HasInterSwitchLink() bool`

HasInterSwitchLink returns a boolean if a field has been set.

### GetMacAuthOnly

`func (o *SwitchPortUsage) GetMacAuthOnly() bool`

GetMacAuthOnly returns the MacAuthOnly field if non-nil, zero value otherwise.

### GetMacAuthOnlyOk

`func (o *SwitchPortUsage) GetMacAuthOnlyOk() (*bool, bool)`

GetMacAuthOnlyOk returns a tuple with the MacAuthOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAuthOnly

`func (o *SwitchPortUsage) SetMacAuthOnly(v bool)`

SetMacAuthOnly sets MacAuthOnly field to given value.

### HasMacAuthOnly

`func (o *SwitchPortUsage) HasMacAuthOnly() bool`

HasMacAuthOnly returns a boolean if a field has been set.

### GetMacAuthProtocol

`func (o *SwitchPortUsage) GetMacAuthProtocol() SwitchPortUsageMacAuthProtocol`

GetMacAuthProtocol returns the MacAuthProtocol field if non-nil, zero value otherwise.

### GetMacAuthProtocolOk

`func (o *SwitchPortUsage) GetMacAuthProtocolOk() (*SwitchPortUsageMacAuthProtocol, bool)`

GetMacAuthProtocolOk returns a tuple with the MacAuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAuthProtocol

`func (o *SwitchPortUsage) SetMacAuthProtocol(v SwitchPortUsageMacAuthProtocol)`

SetMacAuthProtocol sets MacAuthProtocol field to given value.

### HasMacAuthProtocol

`func (o *SwitchPortUsage) HasMacAuthProtocol() bool`

HasMacAuthProtocol returns a boolean if a field has been set.

### GetMacLimit

`func (o *SwitchPortUsage) GetMacLimit() int32`

GetMacLimit returns the MacLimit field if non-nil, zero value otherwise.

### GetMacLimitOk

`func (o *SwitchPortUsage) GetMacLimitOk() (*int32, bool)`

GetMacLimitOk returns a tuple with the MacLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLimit

`func (o *SwitchPortUsage) SetMacLimit(v int32)`

SetMacLimit sets MacLimit field to given value.

### HasMacLimit

`func (o *SwitchPortUsage) HasMacLimit() bool`

HasMacLimit returns a boolean if a field has been set.

### GetMode

`func (o *SwitchPortUsage) GetMode() SwitchPortUsageMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SwitchPortUsage) GetModeOk() (*SwitchPortUsageMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SwitchPortUsage) SetMode(v SwitchPortUsageMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SwitchPortUsage) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMtu

`func (o *SwitchPortUsage) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *SwitchPortUsage) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *SwitchPortUsage) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *SwitchPortUsage) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetworks

`func (o *SwitchPortUsage) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *SwitchPortUsage) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *SwitchPortUsage) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *SwitchPortUsage) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetPersistMac

`func (o *SwitchPortUsage) GetPersistMac() bool`

GetPersistMac returns the PersistMac field if non-nil, zero value otherwise.

### GetPersistMacOk

`func (o *SwitchPortUsage) GetPersistMacOk() (*bool, bool)`

GetPersistMacOk returns a tuple with the PersistMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMac

`func (o *SwitchPortUsage) SetPersistMac(v bool)`

SetPersistMac sets PersistMac field to given value.

### HasPersistMac

`func (o *SwitchPortUsage) HasPersistMac() bool`

HasPersistMac returns a boolean if a field has been set.

### GetPoeDisabled

`func (o *SwitchPortUsage) GetPoeDisabled() bool`

GetPoeDisabled returns the PoeDisabled field if non-nil, zero value otherwise.

### GetPoeDisabledOk

`func (o *SwitchPortUsage) GetPoeDisabledOk() (*bool, bool)`

GetPoeDisabledOk returns a tuple with the PoeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisabled

`func (o *SwitchPortUsage) SetPoeDisabled(v bool)`

SetPoeDisabled sets PoeDisabled field to given value.

### HasPoeDisabled

`func (o *SwitchPortUsage) HasPoeDisabled() bool`

HasPoeDisabled returns a boolean if a field has been set.

### GetPortAuth

`func (o *SwitchPortUsage) GetPortAuth() string`

GetPortAuth returns the PortAuth field if non-nil, zero value otherwise.

### GetPortAuthOk

`func (o *SwitchPortUsage) GetPortAuthOk() (*string, bool)`

GetPortAuthOk returns a tuple with the PortAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAuth

`func (o *SwitchPortUsage) SetPortAuth(v string)`

SetPortAuth sets PortAuth field to given value.

### HasPortAuth

`func (o *SwitchPortUsage) HasPortAuth() bool`

HasPortAuth returns a boolean if a field has been set.

### GetPortNetwork

`func (o *SwitchPortUsage) GetPortNetwork() string`

GetPortNetwork returns the PortNetwork field if non-nil, zero value otherwise.

### GetPortNetworkOk

`func (o *SwitchPortUsage) GetPortNetworkOk() (*string, bool)`

GetPortNetworkOk returns a tuple with the PortNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNetwork

`func (o *SwitchPortUsage) SetPortNetwork(v string)`

SetPortNetwork sets PortNetwork field to given value.

### HasPortNetwork

`func (o *SwitchPortUsage) HasPortNetwork() bool`

HasPortNetwork returns a boolean if a field has been set.

### GetReauthInterval

`func (o *SwitchPortUsage) GetReauthInterval() int32`

GetReauthInterval returns the ReauthInterval field if non-nil, zero value otherwise.

### GetReauthIntervalOk

`func (o *SwitchPortUsage) GetReauthIntervalOk() (*int32, bool)`

GetReauthIntervalOk returns a tuple with the ReauthInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthInterval

`func (o *SwitchPortUsage) SetReauthInterval(v int32)`

SetReauthInterval sets ReauthInterval field to given value.

### HasReauthInterval

`func (o *SwitchPortUsage) HasReauthInterval() bool`

HasReauthInterval returns a boolean if a field has been set.

### GetRejectedNetwork

`func (o *SwitchPortUsage) GetRejectedNetwork() string`

GetRejectedNetwork returns the RejectedNetwork field if non-nil, zero value otherwise.

### GetRejectedNetworkOk

`func (o *SwitchPortUsage) GetRejectedNetworkOk() (*string, bool)`

GetRejectedNetworkOk returns a tuple with the RejectedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNetwork

`func (o *SwitchPortUsage) SetRejectedNetwork(v string)`

SetRejectedNetwork sets RejectedNetwork field to given value.

### HasRejectedNetwork

`func (o *SwitchPortUsage) HasRejectedNetwork() bool`

HasRejectedNetwork returns a boolean if a field has been set.

### SetRejectedNetworkNil

`func (o *SwitchPortUsage) SetRejectedNetworkNil(b bool)`

 SetRejectedNetworkNil sets the value for RejectedNetwork to be an explicit nil

### UnsetRejectedNetwork
`func (o *SwitchPortUsage) UnsetRejectedNetwork()`

UnsetRejectedNetwork ensures that no value is present for RejectedNetwork, not even an explicit nil
### GetResetDefaultWhen

`func (o *SwitchPortUsage) GetResetDefaultWhen() SwitchPortUsageDynamicResetDefaultWhen`

GetResetDefaultWhen returns the ResetDefaultWhen field if non-nil, zero value otherwise.

### GetResetDefaultWhenOk

`func (o *SwitchPortUsage) GetResetDefaultWhenOk() (*SwitchPortUsageDynamicResetDefaultWhen, bool)`

GetResetDefaultWhenOk returns a tuple with the ResetDefaultWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetDefaultWhen

`func (o *SwitchPortUsage) SetResetDefaultWhen(v SwitchPortUsageDynamicResetDefaultWhen)`

SetResetDefaultWhen sets ResetDefaultWhen field to given value.

### HasResetDefaultWhen

`func (o *SwitchPortUsage) HasResetDefaultWhen() bool`

HasResetDefaultWhen returns a boolean if a field has been set.

### GetRules

`func (o *SwitchPortUsage) GetRules() []SwitchPortUsageDynamicRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SwitchPortUsage) GetRulesOk() (*[]SwitchPortUsageDynamicRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SwitchPortUsage) SetRules(v []SwitchPortUsageDynamicRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SwitchPortUsage) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSpeed

`func (o *SwitchPortUsage) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *SwitchPortUsage) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *SwitchPortUsage) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *SwitchPortUsage) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStormControl

`func (o *SwitchPortUsage) GetStormControl() SwitchPortUsageStormControl`

GetStormControl returns the StormControl field if non-nil, zero value otherwise.

### GetStormControlOk

`func (o *SwitchPortUsage) GetStormControlOk() (*SwitchPortUsageStormControl, bool)`

GetStormControlOk returns a tuple with the StormControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormControl

`func (o *SwitchPortUsage) SetStormControl(v SwitchPortUsageStormControl)`

SetStormControl sets StormControl field to given value.

### HasStormControl

`func (o *SwitchPortUsage) HasStormControl() bool`

HasStormControl returns a boolean if a field has been set.

### GetStpEdge

`func (o *SwitchPortUsage) GetStpEdge() bool`

GetStpEdge returns the StpEdge field if non-nil, zero value otherwise.

### GetStpEdgeOk

`func (o *SwitchPortUsage) GetStpEdgeOk() (*bool, bool)`

GetStpEdgeOk returns a tuple with the StpEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpEdge

`func (o *SwitchPortUsage) SetStpEdge(v bool)`

SetStpEdge sets StpEdge field to given value.

### HasStpEdge

`func (o *SwitchPortUsage) HasStpEdge() bool`

HasStpEdge returns a boolean if a field has been set.

### GetVoipNetwork

`func (o *SwitchPortUsage) GetVoipNetwork() string`

GetVoipNetwork returns the VoipNetwork field if non-nil, zero value otherwise.

### GetVoipNetworkOk

`func (o *SwitchPortUsage) GetVoipNetworkOk() (*string, bool)`

GetVoipNetworkOk returns a tuple with the VoipNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipNetwork

`func (o *SwitchPortUsage) SetVoipNetwork(v string)`

SetVoipNetwork sets VoipNetwork field to given value.

### HasVoipNetwork

`func (o *SwitchPortUsage) HasVoipNetwork() bool`

HasVoipNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



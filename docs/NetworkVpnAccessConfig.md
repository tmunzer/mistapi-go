# NetworkVpnAccessConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertisedSubnet** | Pointer to **string** | if &#x60;routed&#x60;&#x3D;&#x3D;&#x60;true&#x60;, whether to advertise an aggregated subnet toward HUB this is useful when there are multiple networks on SPOKE&#39;s side | [optional] 
**AllowPing** | Pointer to **bool** | whether to allow ping from vpn into this routed network | [optional] 
**DestinationNat** | Pointer to [**map[string]NetworkDestinationNatProperty**](NetworkDestinationNatProperty.md) | Property key may be an IP/Port (i.e. \&quot;63.16.0.3:443\&quot;), or a port (i.e. \&quot;:2222\&quot;) | [optional] 
**NatPool** | Pointer to **string** | if &#x60;routed&#x60;&#x3D;&#x3D;&#x60;false&#x60; (usually at Spoke), but some hosts needs to be reachable from Hub, a subnet is required to create and advertise the route to Hub | [optional] 
**NoReadvertiseToLanBgp** | Pointer to **bool** | toward LAN-side BGP peers | [optional] [default to false]
**NoReadvertiseToLanOspf** | Pointer to **bool** | toward LAN-side OSPF peers | [optional] [default to false]
**NoReadvertiseToOverlay** | Pointer to **bool** | toward overlay how HUB should deal with routes it received from Spokes | [optional] 
**OtherVrfs** | Pointer to **[]string** | by default, the routes are only readvertised toward the same vrf on spoke to allow it to be leaked to other vrfs | [optional] 
**Routed** | Pointer to **bool** | whether this network is routable | [optional] 
**SourceNat** | Pointer to [**NetworkSourceNat**](NetworkSourceNat.md) |  | [optional] 
**StaticNat** | Pointer to [**map[string]NetworkStaticNatProperty**](NetworkStaticNatProperty.md) | Property key may be an IP Address (i.e. \&quot;172.16.0.1\&quot;), and IP Address and Port (i.e. \&quot;172.16.0.1:8443\&quot;) or a CIDR (i.e. \&quot;172.16.0.12/20\&quot;) | [optional] 
**SummarizedSubnet** | Pointer to **string** | toward overlay how HUB should deal with routes it received from Spokes | [optional] 
**SummarizedSubnetToLanBgp** | Pointer to **string** | toward LAN-side BGP peers | [optional] 
**SummarizedSubnetToLanOspf** | Pointer to **string** | toward LAN-side OSPF peers | [optional] 

## Methods

### NewNetworkVpnAccessConfig

`func NewNetworkVpnAccessConfig() *NetworkVpnAccessConfig`

NewNetworkVpnAccessConfig instantiates a new NetworkVpnAccessConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVpnAccessConfigWithDefaults

`func NewNetworkVpnAccessConfigWithDefaults() *NetworkVpnAccessConfig`

NewNetworkVpnAccessConfigWithDefaults instantiates a new NetworkVpnAccessConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertisedSubnet

`func (o *NetworkVpnAccessConfig) GetAdvertisedSubnet() string`

GetAdvertisedSubnet returns the AdvertisedSubnet field if non-nil, zero value otherwise.

### GetAdvertisedSubnetOk

`func (o *NetworkVpnAccessConfig) GetAdvertisedSubnetOk() (*string, bool)`

GetAdvertisedSubnetOk returns a tuple with the AdvertisedSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisedSubnet

`func (o *NetworkVpnAccessConfig) SetAdvertisedSubnet(v string)`

SetAdvertisedSubnet sets AdvertisedSubnet field to given value.

### HasAdvertisedSubnet

`func (o *NetworkVpnAccessConfig) HasAdvertisedSubnet() bool`

HasAdvertisedSubnet returns a boolean if a field has been set.

### GetAllowPing

`func (o *NetworkVpnAccessConfig) GetAllowPing() bool`

GetAllowPing returns the AllowPing field if non-nil, zero value otherwise.

### GetAllowPingOk

`func (o *NetworkVpnAccessConfig) GetAllowPingOk() (*bool, bool)`

GetAllowPingOk returns a tuple with the AllowPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPing

`func (o *NetworkVpnAccessConfig) SetAllowPing(v bool)`

SetAllowPing sets AllowPing field to given value.

### HasAllowPing

`func (o *NetworkVpnAccessConfig) HasAllowPing() bool`

HasAllowPing returns a boolean if a field has been set.

### GetDestinationNat

`func (o *NetworkVpnAccessConfig) GetDestinationNat() map[string]NetworkDestinationNatProperty`

GetDestinationNat returns the DestinationNat field if non-nil, zero value otherwise.

### GetDestinationNatOk

`func (o *NetworkVpnAccessConfig) GetDestinationNatOk() (*map[string]NetworkDestinationNatProperty, bool)`

GetDestinationNatOk returns a tuple with the DestinationNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNat

`func (o *NetworkVpnAccessConfig) SetDestinationNat(v map[string]NetworkDestinationNatProperty)`

SetDestinationNat sets DestinationNat field to given value.

### HasDestinationNat

`func (o *NetworkVpnAccessConfig) HasDestinationNat() bool`

HasDestinationNat returns a boolean if a field has been set.

### GetNatPool

`func (o *NetworkVpnAccessConfig) GetNatPool() string`

GetNatPool returns the NatPool field if non-nil, zero value otherwise.

### GetNatPoolOk

`func (o *NetworkVpnAccessConfig) GetNatPoolOk() (*string, bool)`

GetNatPoolOk returns a tuple with the NatPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPool

`func (o *NetworkVpnAccessConfig) SetNatPool(v string)`

SetNatPool sets NatPool field to given value.

### HasNatPool

`func (o *NetworkVpnAccessConfig) HasNatPool() bool`

HasNatPool returns a boolean if a field has been set.

### GetNoReadvertiseToLanBgp

`func (o *NetworkVpnAccessConfig) GetNoReadvertiseToLanBgp() bool`

GetNoReadvertiseToLanBgp returns the NoReadvertiseToLanBgp field if non-nil, zero value otherwise.

### GetNoReadvertiseToLanBgpOk

`func (o *NetworkVpnAccessConfig) GetNoReadvertiseToLanBgpOk() (*bool, bool)`

GetNoReadvertiseToLanBgpOk returns a tuple with the NoReadvertiseToLanBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoReadvertiseToLanBgp

`func (o *NetworkVpnAccessConfig) SetNoReadvertiseToLanBgp(v bool)`

SetNoReadvertiseToLanBgp sets NoReadvertiseToLanBgp field to given value.

### HasNoReadvertiseToLanBgp

`func (o *NetworkVpnAccessConfig) HasNoReadvertiseToLanBgp() bool`

HasNoReadvertiseToLanBgp returns a boolean if a field has been set.

### GetNoReadvertiseToLanOspf

`func (o *NetworkVpnAccessConfig) GetNoReadvertiseToLanOspf() bool`

GetNoReadvertiseToLanOspf returns the NoReadvertiseToLanOspf field if non-nil, zero value otherwise.

### GetNoReadvertiseToLanOspfOk

`func (o *NetworkVpnAccessConfig) GetNoReadvertiseToLanOspfOk() (*bool, bool)`

GetNoReadvertiseToLanOspfOk returns a tuple with the NoReadvertiseToLanOspf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoReadvertiseToLanOspf

`func (o *NetworkVpnAccessConfig) SetNoReadvertiseToLanOspf(v bool)`

SetNoReadvertiseToLanOspf sets NoReadvertiseToLanOspf field to given value.

### HasNoReadvertiseToLanOspf

`func (o *NetworkVpnAccessConfig) HasNoReadvertiseToLanOspf() bool`

HasNoReadvertiseToLanOspf returns a boolean if a field has been set.

### GetNoReadvertiseToOverlay

`func (o *NetworkVpnAccessConfig) GetNoReadvertiseToOverlay() bool`

GetNoReadvertiseToOverlay returns the NoReadvertiseToOverlay field if non-nil, zero value otherwise.

### GetNoReadvertiseToOverlayOk

`func (o *NetworkVpnAccessConfig) GetNoReadvertiseToOverlayOk() (*bool, bool)`

GetNoReadvertiseToOverlayOk returns a tuple with the NoReadvertiseToOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoReadvertiseToOverlay

`func (o *NetworkVpnAccessConfig) SetNoReadvertiseToOverlay(v bool)`

SetNoReadvertiseToOverlay sets NoReadvertiseToOverlay field to given value.

### HasNoReadvertiseToOverlay

`func (o *NetworkVpnAccessConfig) HasNoReadvertiseToOverlay() bool`

HasNoReadvertiseToOverlay returns a boolean if a field has been set.

### GetOtherVrfs

`func (o *NetworkVpnAccessConfig) GetOtherVrfs() []string`

GetOtherVrfs returns the OtherVrfs field if non-nil, zero value otherwise.

### GetOtherVrfsOk

`func (o *NetworkVpnAccessConfig) GetOtherVrfsOk() (*[]string, bool)`

GetOtherVrfsOk returns a tuple with the OtherVrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherVrfs

`func (o *NetworkVpnAccessConfig) SetOtherVrfs(v []string)`

SetOtherVrfs sets OtherVrfs field to given value.

### HasOtherVrfs

`func (o *NetworkVpnAccessConfig) HasOtherVrfs() bool`

HasOtherVrfs returns a boolean if a field has been set.

### GetRouted

`func (o *NetworkVpnAccessConfig) GetRouted() bool`

GetRouted returns the Routed field if non-nil, zero value otherwise.

### GetRoutedOk

`func (o *NetworkVpnAccessConfig) GetRoutedOk() (*bool, bool)`

GetRoutedOk returns a tuple with the Routed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouted

`func (o *NetworkVpnAccessConfig) SetRouted(v bool)`

SetRouted sets Routed field to given value.

### HasRouted

`func (o *NetworkVpnAccessConfig) HasRouted() bool`

HasRouted returns a boolean if a field has been set.

### GetSourceNat

`func (o *NetworkVpnAccessConfig) GetSourceNat() NetworkSourceNat`

GetSourceNat returns the SourceNat field if non-nil, zero value otherwise.

### GetSourceNatOk

`func (o *NetworkVpnAccessConfig) GetSourceNatOk() (*NetworkSourceNat, bool)`

GetSourceNatOk returns a tuple with the SourceNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNat

`func (o *NetworkVpnAccessConfig) SetSourceNat(v NetworkSourceNat)`

SetSourceNat sets SourceNat field to given value.

### HasSourceNat

`func (o *NetworkVpnAccessConfig) HasSourceNat() bool`

HasSourceNat returns a boolean if a field has been set.

### GetStaticNat

`func (o *NetworkVpnAccessConfig) GetStaticNat() map[string]NetworkStaticNatProperty`

GetStaticNat returns the StaticNat field if non-nil, zero value otherwise.

### GetStaticNatOk

`func (o *NetworkVpnAccessConfig) GetStaticNatOk() (*map[string]NetworkStaticNatProperty, bool)`

GetStaticNatOk returns a tuple with the StaticNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticNat

`func (o *NetworkVpnAccessConfig) SetStaticNat(v map[string]NetworkStaticNatProperty)`

SetStaticNat sets StaticNat field to given value.

### HasStaticNat

`func (o *NetworkVpnAccessConfig) HasStaticNat() bool`

HasStaticNat returns a boolean if a field has been set.

### GetSummarizedSubnet

`func (o *NetworkVpnAccessConfig) GetSummarizedSubnet() string`

GetSummarizedSubnet returns the SummarizedSubnet field if non-nil, zero value otherwise.

### GetSummarizedSubnetOk

`func (o *NetworkVpnAccessConfig) GetSummarizedSubnetOk() (*string, bool)`

GetSummarizedSubnetOk returns a tuple with the SummarizedSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarizedSubnet

`func (o *NetworkVpnAccessConfig) SetSummarizedSubnet(v string)`

SetSummarizedSubnet sets SummarizedSubnet field to given value.

### HasSummarizedSubnet

`func (o *NetworkVpnAccessConfig) HasSummarizedSubnet() bool`

HasSummarizedSubnet returns a boolean if a field has been set.

### GetSummarizedSubnetToLanBgp

`func (o *NetworkVpnAccessConfig) GetSummarizedSubnetToLanBgp() string`

GetSummarizedSubnetToLanBgp returns the SummarizedSubnetToLanBgp field if non-nil, zero value otherwise.

### GetSummarizedSubnetToLanBgpOk

`func (o *NetworkVpnAccessConfig) GetSummarizedSubnetToLanBgpOk() (*string, bool)`

GetSummarizedSubnetToLanBgpOk returns a tuple with the SummarizedSubnetToLanBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarizedSubnetToLanBgp

`func (o *NetworkVpnAccessConfig) SetSummarizedSubnetToLanBgp(v string)`

SetSummarizedSubnetToLanBgp sets SummarizedSubnetToLanBgp field to given value.

### HasSummarizedSubnetToLanBgp

`func (o *NetworkVpnAccessConfig) HasSummarizedSubnetToLanBgp() bool`

HasSummarizedSubnetToLanBgp returns a boolean if a field has been set.

### GetSummarizedSubnetToLanOspf

`func (o *NetworkVpnAccessConfig) GetSummarizedSubnetToLanOspf() string`

GetSummarizedSubnetToLanOspf returns the SummarizedSubnetToLanOspf field if non-nil, zero value otherwise.

### GetSummarizedSubnetToLanOspfOk

`func (o *NetworkVpnAccessConfig) GetSummarizedSubnetToLanOspfOk() (*string, bool)`

GetSummarizedSubnetToLanOspfOk returns a tuple with the SummarizedSubnetToLanOspf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarizedSubnetToLanOspf

`func (o *NetworkVpnAccessConfig) SetSummarizedSubnetToLanOspf(v string)`

SetSummarizedSubnetToLanOspf sets SummarizedSubnetToLanOspf field to given value.

### HasSummarizedSubnetToLanOspf

`func (o *NetworkVpnAccessConfig) HasSummarizedSubnetToLanOspf() bool`

HasSummarizedSubnetToLanOspf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



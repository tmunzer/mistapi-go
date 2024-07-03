# RoutingPolicyTermMatching

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsPath** | Pointer to **[]string** | takes regular expression | [optional] 
**Community** | Pointer to **[]string** |  | [optional] 
**Network** | Pointer to **[]string** |  | [optional] 
**Prefix** | Pointer to **[]string** | zero or more criteria/filter can be specified to match the term, all criteria have to be met | [optional] 
**Protocol** | Pointer to **[]string** | &#x60;direct&#x60;, &#x60;bgp&#x60;, &#x60;osp&#x60;, ... | [optional] 
**RouteExists** | Pointer to [**RoutingPolicyTermMatchingRouteExists**](RoutingPolicyTermMatchingRouteExists.md) |  | [optional] 
**VpnNeighborMac** | Pointer to **[]string** | overlay-facing criteria (used for bgp_config where via&#x3D;vpn) | [optional] 
**VpnPath** | Pointer to **[]string** | overlay-facing criteria (used for bgp_config where via&#x3D;vpn) ordered- | [optional] 
**VpnPathSla** | Pointer to [**RoutingPolicyTermMatchingVpnPathSla**](RoutingPolicyTermMatchingVpnPathSla.md) |  | [optional] 

## Methods

### NewRoutingPolicyTermMatching

`func NewRoutingPolicyTermMatching() *RoutingPolicyTermMatching`

NewRoutingPolicyTermMatching instantiates a new RoutingPolicyTermMatching object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyTermMatchingWithDefaults

`func NewRoutingPolicyTermMatchingWithDefaults() *RoutingPolicyTermMatching`

NewRoutingPolicyTermMatchingWithDefaults instantiates a new RoutingPolicyTermMatching object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsPath

`func (o *RoutingPolicyTermMatching) GetAsPath() []string`

GetAsPath returns the AsPath field if non-nil, zero value otherwise.

### GetAsPathOk

`func (o *RoutingPolicyTermMatching) GetAsPathOk() (*[]string, bool)`

GetAsPathOk returns a tuple with the AsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsPath

`func (o *RoutingPolicyTermMatching) SetAsPath(v []string)`

SetAsPath sets AsPath field to given value.

### HasAsPath

`func (o *RoutingPolicyTermMatching) HasAsPath() bool`

HasAsPath returns a boolean if a field has been set.

### GetCommunity

`func (o *RoutingPolicyTermMatching) GetCommunity() []string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *RoutingPolicyTermMatching) GetCommunityOk() (*[]string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *RoutingPolicyTermMatching) SetCommunity(v []string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *RoutingPolicyTermMatching) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetNetwork

`func (o *RoutingPolicyTermMatching) GetNetwork() []string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RoutingPolicyTermMatching) GetNetworkOk() (*[]string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RoutingPolicyTermMatching) SetNetwork(v []string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RoutingPolicyTermMatching) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPrefix

`func (o *RoutingPolicyTermMatching) GetPrefix() []string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RoutingPolicyTermMatching) GetPrefixOk() (*[]string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RoutingPolicyTermMatching) SetPrefix(v []string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RoutingPolicyTermMatching) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProtocol

`func (o *RoutingPolicyTermMatching) GetProtocol() []string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RoutingPolicyTermMatching) GetProtocolOk() (*[]string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RoutingPolicyTermMatching) SetProtocol(v []string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RoutingPolicyTermMatching) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRouteExists

`func (o *RoutingPolicyTermMatching) GetRouteExists() RoutingPolicyTermMatchingRouteExists`

GetRouteExists returns the RouteExists field if non-nil, zero value otherwise.

### GetRouteExistsOk

`func (o *RoutingPolicyTermMatching) GetRouteExistsOk() (*RoutingPolicyTermMatchingRouteExists, bool)`

GetRouteExistsOk returns a tuple with the RouteExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteExists

`func (o *RoutingPolicyTermMatching) SetRouteExists(v RoutingPolicyTermMatchingRouteExists)`

SetRouteExists sets RouteExists field to given value.

### HasRouteExists

`func (o *RoutingPolicyTermMatching) HasRouteExists() bool`

HasRouteExists returns a boolean if a field has been set.

### GetVpnNeighborMac

`func (o *RoutingPolicyTermMatching) GetVpnNeighborMac() []string`

GetVpnNeighborMac returns the VpnNeighborMac field if non-nil, zero value otherwise.

### GetVpnNeighborMacOk

`func (o *RoutingPolicyTermMatching) GetVpnNeighborMacOk() (*[]string, bool)`

GetVpnNeighborMacOk returns a tuple with the VpnNeighborMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnNeighborMac

`func (o *RoutingPolicyTermMatching) SetVpnNeighborMac(v []string)`

SetVpnNeighborMac sets VpnNeighborMac field to given value.

### HasVpnNeighborMac

`func (o *RoutingPolicyTermMatching) HasVpnNeighborMac() bool`

HasVpnNeighborMac returns a boolean if a field has been set.

### GetVpnPath

`func (o *RoutingPolicyTermMatching) GetVpnPath() []string`

GetVpnPath returns the VpnPath field if non-nil, zero value otherwise.

### GetVpnPathOk

`func (o *RoutingPolicyTermMatching) GetVpnPathOk() (*[]string, bool)`

GetVpnPathOk returns a tuple with the VpnPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPath

`func (o *RoutingPolicyTermMatching) SetVpnPath(v []string)`

SetVpnPath sets VpnPath field to given value.

### HasVpnPath

`func (o *RoutingPolicyTermMatching) HasVpnPath() bool`

HasVpnPath returns a boolean if a field has been set.

### GetVpnPathSla

`func (o *RoutingPolicyTermMatching) GetVpnPathSla() RoutingPolicyTermMatchingVpnPathSla`

GetVpnPathSla returns the VpnPathSla field if non-nil, zero value otherwise.

### GetVpnPathSlaOk

`func (o *RoutingPolicyTermMatching) GetVpnPathSlaOk() (*RoutingPolicyTermMatchingVpnPathSla, bool)`

GetVpnPathSlaOk returns a tuple with the VpnPathSla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPathSla

`func (o *RoutingPolicyTermMatching) SetVpnPathSla(v RoutingPolicyTermMatchingVpnPathSla)`

SetVpnPathSla sets VpnPathSla field to given value.

### HasVpnPathSla

`func (o *RoutingPolicyTermMatching) HasVpnPathSla() bool`

HasVpnPathSla returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



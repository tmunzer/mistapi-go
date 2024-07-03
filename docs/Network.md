# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DisallowMistServices** | Pointer to **bool** | whether to disallow Mist Devices in the network | [optional] [default to false]
**Gateway** | Pointer to **string** |  | [optional] 
**Gateway6** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**InternalAccess** | Pointer to [**NetworkInternalAccess**](NetworkInternalAccess.md) |  | [optional] 
**InternetAccess** | Pointer to [**NetworkInternetAccess**](NetworkInternetAccess.md) |  | [optional] 
**Isolation** | Pointer to **bool** | whether to allow clients in the network to talk to each other | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**RoutedForNetworks** | Pointer to **[]string** | for a Network (usually LAN), it can be routable to other networks (e.g. OSPF) | [optional] 
**Subnet** | Pointer to **string** |  | [optional] 
**Subnet6** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**map[string]NetworkTenant**](NetworkTenant.md) |  | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 
**VpnAccess** | Pointer to [**map[string]NetworkVpnAccessConfig**](NetworkVpnAccessConfig.md) | Property key is the VPN name. Whether this network can be accessed from vpn | [optional] 

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *Network) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Network) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Network) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Network) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisallowMistServices

`func (o *Network) GetDisallowMistServices() bool`

GetDisallowMistServices returns the DisallowMistServices field if non-nil, zero value otherwise.

### GetDisallowMistServicesOk

`func (o *Network) GetDisallowMistServicesOk() (*bool, bool)`

GetDisallowMistServicesOk returns a tuple with the DisallowMistServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowMistServices

`func (o *Network) SetDisallowMistServices(v bool)`

SetDisallowMistServices sets DisallowMistServices field to given value.

### HasDisallowMistServices

`func (o *Network) HasDisallowMistServices() bool`

HasDisallowMistServices returns a boolean if a field has been set.

### GetGateway

`func (o *Network) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Network) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Network) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Network) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGateway6

`func (o *Network) GetGateway6() string`

GetGateway6 returns the Gateway6 field if non-nil, zero value otherwise.

### GetGateway6Ok

`func (o *Network) GetGateway6Ok() (*string, bool)`

GetGateway6Ok returns a tuple with the Gateway6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway6

`func (o *Network) SetGateway6(v string)`

SetGateway6 sets Gateway6 field to given value.

### HasGateway6

`func (o *Network) HasGateway6() bool`

HasGateway6 returns a boolean if a field has been set.

### GetId

`func (o *Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Network) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Network) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalAccess

`func (o *Network) GetInternalAccess() NetworkInternalAccess`

GetInternalAccess returns the InternalAccess field if non-nil, zero value otherwise.

### GetInternalAccessOk

`func (o *Network) GetInternalAccessOk() (*NetworkInternalAccess, bool)`

GetInternalAccessOk returns a tuple with the InternalAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccess

`func (o *Network) SetInternalAccess(v NetworkInternalAccess)`

SetInternalAccess sets InternalAccess field to given value.

### HasInternalAccess

`func (o *Network) HasInternalAccess() bool`

HasInternalAccess returns a boolean if a field has been set.

### GetInternetAccess

`func (o *Network) GetInternetAccess() NetworkInternetAccess`

GetInternetAccess returns the InternetAccess field if non-nil, zero value otherwise.

### GetInternetAccessOk

`func (o *Network) GetInternetAccessOk() (*NetworkInternetAccess, bool)`

GetInternetAccessOk returns a tuple with the InternetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetAccess

`func (o *Network) SetInternetAccess(v NetworkInternetAccess)`

SetInternetAccess sets InternetAccess field to given value.

### HasInternetAccess

`func (o *Network) HasInternetAccess() bool`

HasInternetAccess returns a boolean if a field has been set.

### GetIsolation

`func (o *Network) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *Network) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *Network) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *Network) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Network) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Network) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Network) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Network) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Network) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Network) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Network) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Network) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Network) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Network) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRoutedForNetworks

`func (o *Network) GetRoutedForNetworks() []string`

GetRoutedForNetworks returns the RoutedForNetworks field if non-nil, zero value otherwise.

### GetRoutedForNetworksOk

`func (o *Network) GetRoutedForNetworksOk() (*[]string, bool)`

GetRoutedForNetworksOk returns a tuple with the RoutedForNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutedForNetworks

`func (o *Network) SetRoutedForNetworks(v []string)`

SetRoutedForNetworks sets RoutedForNetworks field to given value.

### HasRoutedForNetworks

`func (o *Network) HasRoutedForNetworks() bool`

HasRoutedForNetworks returns a boolean if a field has been set.

### GetSubnet

`func (o *Network) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *Network) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *Network) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *Network) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSubnet6

`func (o *Network) GetSubnet6() string`

GetSubnet6 returns the Subnet6 field if non-nil, zero value otherwise.

### GetSubnet6Ok

`func (o *Network) GetSubnet6Ok() (*string, bool)`

GetSubnet6Ok returns a tuple with the Subnet6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet6

`func (o *Network) SetSubnet6(v string)`

SetSubnet6 sets Subnet6 field to given value.

### HasSubnet6

`func (o *Network) HasSubnet6() bool`

HasSubnet6 returns a boolean if a field has been set.

### GetTenants

`func (o *Network) GetTenants() map[string]NetworkTenant`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *Network) GetTenantsOk() (*map[string]NetworkTenant, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *Network) SetTenants(v map[string]NetworkTenant)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *Network) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetVlanId

`func (o *Network) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *Network) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *Network) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *Network) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVpnAccess

`func (o *Network) GetVpnAccess() map[string]NetworkVpnAccessConfig`

GetVpnAccess returns the VpnAccess field if non-nil, zero value otherwise.

### GetVpnAccessOk

`func (o *Network) GetVpnAccessOk() (*map[string]NetworkVpnAccessConfig, bool)`

GetVpnAccessOk returns a tuple with the VpnAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnAccess

`func (o *Network) SetVpnAccess(v map[string]NetworkVpnAccessConfig)`

SetVpnAccess sets VpnAccess field to given value.

### HasVpnAccess

`func (o *Network) HasVpnAccess() bool`

HasVpnAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



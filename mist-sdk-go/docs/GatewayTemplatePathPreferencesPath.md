# GatewayTemplatePathPreferencesPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **int32** |  | [optional] 
**Disabled** | Pointer to **bool** | For SSR Only. &#x60;true&#x60;&#x60; if this specific path is undesired | [optional] 
**GatewayIp** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60;, if a different gateway is desired | [optional] 
**InternetAccess** | Pointer to **bool** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;vpn&#x60;, if this vpn path can be used for internet | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; | [optional] 
**TargetIps** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60;, if destination IP is to be replaced | [optional] 
**Type** | Pointer to [**GatewayPathType**](GatewayPathType.md) |  | [optional] 
**WanName** | Pointer to **string** | Spoke&#39;s outgoing wan | [optional] 

## Methods

### NewGatewayTemplatePathPreferencesPath

`func NewGatewayTemplatePathPreferencesPath() *GatewayTemplatePathPreferencesPath`

NewGatewayTemplatePathPreferencesPath instantiates a new GatewayTemplatePathPreferencesPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplatePathPreferencesPathWithDefaults

`func NewGatewayTemplatePathPreferencesPathWithDefaults() *GatewayTemplatePathPreferencesPath`

NewGatewayTemplatePathPreferencesPathWithDefaults instantiates a new GatewayTemplatePathPreferencesPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *GatewayTemplatePathPreferencesPath) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GatewayTemplatePathPreferencesPath) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GatewayTemplatePathPreferencesPath) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GatewayTemplatePathPreferencesPath) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDisabled

`func (o *GatewayTemplatePathPreferencesPath) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GatewayTemplatePathPreferencesPath) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GatewayTemplatePathPreferencesPath) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GatewayTemplatePathPreferencesPath) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGatewayIp

`func (o *GatewayTemplatePathPreferencesPath) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *GatewayTemplatePathPreferencesPath) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *GatewayTemplatePathPreferencesPath) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *GatewayTemplatePathPreferencesPath) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetInternetAccess

`func (o *GatewayTemplatePathPreferencesPath) GetInternetAccess() bool`

GetInternetAccess returns the InternetAccess field if non-nil, zero value otherwise.

### GetInternetAccessOk

`func (o *GatewayTemplatePathPreferencesPath) GetInternetAccessOk() (*bool, bool)`

GetInternetAccessOk returns a tuple with the InternetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetAccess

`func (o *GatewayTemplatePathPreferencesPath) SetInternetAccess(v bool)`

SetInternetAccess sets InternetAccess field to given value.

### HasInternetAccess

`func (o *GatewayTemplatePathPreferencesPath) HasInternetAccess() bool`

HasInternetAccess returns a boolean if a field has been set.

### GetName

`func (o *GatewayTemplatePathPreferencesPath) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayTemplatePathPreferencesPath) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayTemplatePathPreferencesPath) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayTemplatePathPreferencesPath) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *GatewayTemplatePathPreferencesPath) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GatewayTemplatePathPreferencesPath) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GatewayTemplatePathPreferencesPath) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GatewayTemplatePathPreferencesPath) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetTargetIps

`func (o *GatewayTemplatePathPreferencesPath) GetTargetIps() []string`

GetTargetIps returns the TargetIps field if non-nil, zero value otherwise.

### GetTargetIpsOk

`func (o *GatewayTemplatePathPreferencesPath) GetTargetIpsOk() (*[]string, bool)`

GetTargetIpsOk returns a tuple with the TargetIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIps

`func (o *GatewayTemplatePathPreferencesPath) SetTargetIps(v []string)`

SetTargetIps sets TargetIps field to given value.

### HasTargetIps

`func (o *GatewayTemplatePathPreferencesPath) HasTargetIps() bool`

HasTargetIps returns a boolean if a field has been set.

### GetType

`func (o *GatewayTemplatePathPreferencesPath) GetType() GatewayPathType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayTemplatePathPreferencesPath) GetTypeOk() (*GatewayPathType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayTemplatePathPreferencesPath) SetType(v GatewayPathType)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayTemplatePathPreferencesPath) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWanName

`func (o *GatewayTemplatePathPreferencesPath) GetWanName() string`

GetWanName returns the WanName field if non-nil, zero value otherwise.

### GetWanNameOk

`func (o *GatewayTemplatePathPreferencesPath) GetWanNameOk() (*string, bool)`

GetWanNameOk returns a tuple with the WanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanName

`func (o *GatewayTemplatePathPreferencesPath) SetWanName(v string)`

SetWanName sets WanName field to given value.

### HasWanName

`func (o *GatewayTemplatePathPreferencesPath) HasWanName() bool`

HasWanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



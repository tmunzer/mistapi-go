# BgpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | Pointer to **string** |  | [optional] 
**BfdMinimumInterval** | Pointer to **NullableInt32** | when bfd_multiplier is configured alone default: * 1000 if &#x60;type&#x60;&#x3D;&#x3D;&#x60;external&#x60;&#x60; * 350 &#x60;type&#x60;&#x3D;&#x3D;&#x60;internal&#x60; | [optional] [default to 350]
**BfdMultiplier** | Pointer to **NullableInt32** | when bfd_minimum_interval_is_configured alone | [optional] [default to 3]
**Communities** | Pointer to [**[]BgpConfigCommunity**](BgpConfigCommunity.md) |  | [optional] 
**DisableBfd** | Pointer to **bool** | BFD provides faster path failure detection and is enabled by default | [optional] [default to false]
**Export** | Pointer to **string** |  | [optional] 
**ExportPolicy** | Pointer to **string** | default export policies if no per-neighbor policies defined | [optional] 
**ExtendedV4Nexthop** | Pointer to **bool** | by default, either inet/net6 unicast depending on neighbor IP family (v4 or v6) for v6 neighbors, to exchange v4 nexthop, which allows dual-stack support, enable this | [optional] 
**GracefulRestartTime** | Pointer to **int32** | &#x60;0&#x60; means disable | [optional] [default to 0]
**HoldTime** | Pointer to **int32** |  | [optional] [default to 90]
**Import** | Pointer to **string** |  | [optional] 
**ImportPolicy** | Pointer to **string** | default import policies if no per-neighbor policies defined | [optional] 
**LocalAs** | Pointer to **int32** |  | [optional] 
**NeighborAs** | Pointer to **int32** |  | [optional] 
**Neighbors** | Pointer to [**map[string]BgpConfigNeighbors**](BgpConfigNeighbors.md) | if per-neighbor as is desired. Property key is the neighbor address | [optional] 
**Networks** | Pointer to **[]string** | if &#x60;type&#x60;!&#x3D;&#x60;external&#x60;or &#x60;via&#x60;&#x3D;&#x3D;&#x60;wan&#x60;networks where we expect BGP neighbor to connect to/from | [optional] 
**NoReadvertiseToOverlay** | Pointer to **bool** | by default, we&#39;ll re-advertise all learned BGP routers toward overlay | [optional] [default to false]
**Type** | Pointer to [**BgpConfigType**](BgpConfigType.md) |  | [optional] 
**Via** | Pointer to [**BgpConfigVia**](BgpConfigVia.md) |  | [optional] [default to BGPCONFIGVIA_LAN]
**VpnName** | Pointer to **string** |  | [optional] 
**WanName** | Pointer to **string** | if &#x60;via&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | [optional] 

## Methods

### NewBgpConfig

`func NewBgpConfig() *BgpConfig`

NewBgpConfig instantiates a new BgpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfigWithDefaults

`func NewBgpConfigWithDefaults() *BgpConfig`

NewBgpConfigWithDefaults instantiates a new BgpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKey

`func (o *BgpConfig) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *BgpConfig) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *BgpConfig) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *BgpConfig) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetBfdMinimumInterval

`func (o *BgpConfig) GetBfdMinimumInterval() int32`

GetBfdMinimumInterval returns the BfdMinimumInterval field if non-nil, zero value otherwise.

### GetBfdMinimumIntervalOk

`func (o *BgpConfig) GetBfdMinimumIntervalOk() (*int32, bool)`

GetBfdMinimumIntervalOk returns a tuple with the BfdMinimumInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMinimumInterval

`func (o *BgpConfig) SetBfdMinimumInterval(v int32)`

SetBfdMinimumInterval sets BfdMinimumInterval field to given value.

### HasBfdMinimumInterval

`func (o *BgpConfig) HasBfdMinimumInterval() bool`

HasBfdMinimumInterval returns a boolean if a field has been set.

### SetBfdMinimumIntervalNil

`func (o *BgpConfig) SetBfdMinimumIntervalNil(b bool)`

 SetBfdMinimumIntervalNil sets the value for BfdMinimumInterval to be an explicit nil

### UnsetBfdMinimumInterval
`func (o *BgpConfig) UnsetBfdMinimumInterval()`

UnsetBfdMinimumInterval ensures that no value is present for BfdMinimumInterval, not even an explicit nil
### GetBfdMultiplier

`func (o *BgpConfig) GetBfdMultiplier() int32`

GetBfdMultiplier returns the BfdMultiplier field if non-nil, zero value otherwise.

### GetBfdMultiplierOk

`func (o *BgpConfig) GetBfdMultiplierOk() (*int32, bool)`

GetBfdMultiplierOk returns a tuple with the BfdMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMultiplier

`func (o *BgpConfig) SetBfdMultiplier(v int32)`

SetBfdMultiplier sets BfdMultiplier field to given value.

### HasBfdMultiplier

`func (o *BgpConfig) HasBfdMultiplier() bool`

HasBfdMultiplier returns a boolean if a field has been set.

### SetBfdMultiplierNil

`func (o *BgpConfig) SetBfdMultiplierNil(b bool)`

 SetBfdMultiplierNil sets the value for BfdMultiplier to be an explicit nil

### UnsetBfdMultiplier
`func (o *BgpConfig) UnsetBfdMultiplier()`

UnsetBfdMultiplier ensures that no value is present for BfdMultiplier, not even an explicit nil
### GetCommunities

`func (o *BgpConfig) GetCommunities() []BgpConfigCommunity`

GetCommunities returns the Communities field if non-nil, zero value otherwise.

### GetCommunitiesOk

`func (o *BgpConfig) GetCommunitiesOk() (*[]BgpConfigCommunity, bool)`

GetCommunitiesOk returns a tuple with the Communities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunities

`func (o *BgpConfig) SetCommunities(v []BgpConfigCommunity)`

SetCommunities sets Communities field to given value.

### HasCommunities

`func (o *BgpConfig) HasCommunities() bool`

HasCommunities returns a boolean if a field has been set.

### GetDisableBfd

`func (o *BgpConfig) GetDisableBfd() bool`

GetDisableBfd returns the DisableBfd field if non-nil, zero value otherwise.

### GetDisableBfdOk

`func (o *BgpConfig) GetDisableBfdOk() (*bool, bool)`

GetDisableBfdOk returns a tuple with the DisableBfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBfd

`func (o *BgpConfig) SetDisableBfd(v bool)`

SetDisableBfd sets DisableBfd field to given value.

### HasDisableBfd

`func (o *BgpConfig) HasDisableBfd() bool`

HasDisableBfd returns a boolean if a field has been set.

### GetExport

`func (o *BgpConfig) GetExport() string`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *BgpConfig) GetExportOk() (*string, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *BgpConfig) SetExport(v string)`

SetExport sets Export field to given value.

### HasExport

`func (o *BgpConfig) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetExportPolicy

`func (o *BgpConfig) GetExportPolicy() string`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *BgpConfig) GetExportPolicyOk() (*string, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *BgpConfig) SetExportPolicy(v string)`

SetExportPolicy sets ExportPolicy field to given value.

### HasExportPolicy

`func (o *BgpConfig) HasExportPolicy() bool`

HasExportPolicy returns a boolean if a field has been set.

### GetExtendedV4Nexthop

`func (o *BgpConfig) GetExtendedV4Nexthop() bool`

GetExtendedV4Nexthop returns the ExtendedV4Nexthop field if non-nil, zero value otherwise.

### GetExtendedV4NexthopOk

`func (o *BgpConfig) GetExtendedV4NexthopOk() (*bool, bool)`

GetExtendedV4NexthopOk returns a tuple with the ExtendedV4Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedV4Nexthop

`func (o *BgpConfig) SetExtendedV4Nexthop(v bool)`

SetExtendedV4Nexthop sets ExtendedV4Nexthop field to given value.

### HasExtendedV4Nexthop

`func (o *BgpConfig) HasExtendedV4Nexthop() bool`

HasExtendedV4Nexthop returns a boolean if a field has been set.

### GetGracefulRestartTime

`func (o *BgpConfig) GetGracefulRestartTime() int32`

GetGracefulRestartTime returns the GracefulRestartTime field if non-nil, zero value otherwise.

### GetGracefulRestartTimeOk

`func (o *BgpConfig) GetGracefulRestartTimeOk() (*int32, bool)`

GetGracefulRestartTimeOk returns a tuple with the GracefulRestartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracefulRestartTime

`func (o *BgpConfig) SetGracefulRestartTime(v int32)`

SetGracefulRestartTime sets GracefulRestartTime field to given value.

### HasGracefulRestartTime

`func (o *BgpConfig) HasGracefulRestartTime() bool`

HasGracefulRestartTime returns a boolean if a field has been set.

### GetHoldTime

`func (o *BgpConfig) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *BgpConfig) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *BgpConfig) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *BgpConfig) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### GetImport

`func (o *BgpConfig) GetImport() string`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *BgpConfig) GetImportOk() (*string, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *BgpConfig) SetImport(v string)`

SetImport sets Import field to given value.

### HasImport

`func (o *BgpConfig) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetImportPolicy

`func (o *BgpConfig) GetImportPolicy() string`

GetImportPolicy returns the ImportPolicy field if non-nil, zero value otherwise.

### GetImportPolicyOk

`func (o *BgpConfig) GetImportPolicyOk() (*string, bool)`

GetImportPolicyOk returns a tuple with the ImportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicy

`func (o *BgpConfig) SetImportPolicy(v string)`

SetImportPolicy sets ImportPolicy field to given value.

### HasImportPolicy

`func (o *BgpConfig) HasImportPolicy() bool`

HasImportPolicy returns a boolean if a field has been set.

### GetLocalAs

`func (o *BgpConfig) GetLocalAs() int32`

GetLocalAs returns the LocalAs field if non-nil, zero value otherwise.

### GetLocalAsOk

`func (o *BgpConfig) GetLocalAsOk() (*int32, bool)`

GetLocalAsOk returns a tuple with the LocalAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAs

`func (o *BgpConfig) SetLocalAs(v int32)`

SetLocalAs sets LocalAs field to given value.

### HasLocalAs

`func (o *BgpConfig) HasLocalAs() bool`

HasLocalAs returns a boolean if a field has been set.

### GetNeighborAs

`func (o *BgpConfig) GetNeighborAs() int32`

GetNeighborAs returns the NeighborAs field if non-nil, zero value otherwise.

### GetNeighborAsOk

`func (o *BgpConfig) GetNeighborAsOk() (*int32, bool)`

GetNeighborAsOk returns a tuple with the NeighborAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborAs

`func (o *BgpConfig) SetNeighborAs(v int32)`

SetNeighborAs sets NeighborAs field to given value.

### HasNeighborAs

`func (o *BgpConfig) HasNeighborAs() bool`

HasNeighborAs returns a boolean if a field has been set.

### GetNeighbors

`func (o *BgpConfig) GetNeighbors() map[string]BgpConfigNeighbors`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *BgpConfig) GetNeighborsOk() (*map[string]BgpConfigNeighbors, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *BgpConfig) SetNeighbors(v map[string]BgpConfigNeighbors)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *BgpConfig) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.

### GetNetworks

`func (o *BgpConfig) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *BgpConfig) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *BgpConfig) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *BgpConfig) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNoReadvertiseToOverlay

`func (o *BgpConfig) GetNoReadvertiseToOverlay() bool`

GetNoReadvertiseToOverlay returns the NoReadvertiseToOverlay field if non-nil, zero value otherwise.

### GetNoReadvertiseToOverlayOk

`func (o *BgpConfig) GetNoReadvertiseToOverlayOk() (*bool, bool)`

GetNoReadvertiseToOverlayOk returns a tuple with the NoReadvertiseToOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoReadvertiseToOverlay

`func (o *BgpConfig) SetNoReadvertiseToOverlay(v bool)`

SetNoReadvertiseToOverlay sets NoReadvertiseToOverlay field to given value.

### HasNoReadvertiseToOverlay

`func (o *BgpConfig) HasNoReadvertiseToOverlay() bool`

HasNoReadvertiseToOverlay returns a boolean if a field has been set.

### GetType

`func (o *BgpConfig) GetType() BgpConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BgpConfig) GetTypeOk() (*BgpConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BgpConfig) SetType(v BgpConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *BgpConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVia

`func (o *BgpConfig) GetVia() BgpConfigVia`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *BgpConfig) GetViaOk() (*BgpConfigVia, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *BgpConfig) SetVia(v BgpConfigVia)`

SetVia sets Via field to given value.

### HasVia

`func (o *BgpConfig) HasVia() bool`

HasVia returns a boolean if a field has been set.

### GetVpnName

`func (o *BgpConfig) GetVpnName() string`

GetVpnName returns the VpnName field if non-nil, zero value otherwise.

### GetVpnNameOk

`func (o *BgpConfig) GetVpnNameOk() (*string, bool)`

GetVpnNameOk returns a tuple with the VpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnName

`func (o *BgpConfig) SetVpnName(v string)`

SetVpnName sets VpnName field to given value.

### HasVpnName

`func (o *BgpConfig) HasVpnName() bool`

HasVpnName returns a boolean if a field has been set.

### GetWanName

`func (o *BgpConfig) GetWanName() string`

GetWanName returns the WanName field if non-nil, zero value otherwise.

### GetWanNameOk

`func (o *BgpConfig) GetWanNameOk() (*string, bool)`

GetWanNameOk returns a tuple with the WanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanName

`func (o *BgpConfig) SetWanName(v string)`

SetWanName sets WanName field to given value.

### HasWanName

`func (o *BgpConfig) HasWanName() bool`

HasWanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



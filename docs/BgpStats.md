# BgpStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvpnOverlay** | Pointer to **bool** | if this is created for evpn overlay | [optional] 
**ForOverlay** | Pointer to **bool** | if this is created for overlay | [optional] 
**LocalAs** | Pointer to **int32** | AS | [optional] 
**Mac** | Pointer to **string** | router mac address | [optional] 
**Neighbor** | Pointer to **string** |  | [optional] 
**NeighborAs** | Pointer to **int32** |  | [optional] 
**NeighborMac** | Pointer to **string** | if it&#39;s another device in the same org | [optional] 
**Node** | Pointer to **string** | node0/node1 | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**RxPkts** | Pointer to **int32** |  | [optional] 
**RxRoutes** | Pointer to **int32** | number of received routes | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to [**BgpStatsState**](BgpStatsState.md) |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**TxPkts** | Pointer to **int32** |  | [optional] 
**TxRoutes** | Pointer to **int32** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**Uptime** | Pointer to **int32** |  | [optional] 
**VrfName** | Pointer to **string** |  | [optional] 

## Methods

### NewBgpStats

`func NewBgpStats() *BgpStats`

NewBgpStats instantiates a new BgpStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpStatsWithDefaults

`func NewBgpStatsWithDefaults() *BgpStats`

NewBgpStatsWithDefaults instantiates a new BgpStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvpnOverlay

`func (o *BgpStats) GetEvpnOverlay() bool`

GetEvpnOverlay returns the EvpnOverlay field if non-nil, zero value otherwise.

### GetEvpnOverlayOk

`func (o *BgpStats) GetEvpnOverlayOk() (*bool, bool)`

GetEvpnOverlayOk returns a tuple with the EvpnOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnOverlay

`func (o *BgpStats) SetEvpnOverlay(v bool)`

SetEvpnOverlay sets EvpnOverlay field to given value.

### HasEvpnOverlay

`func (o *BgpStats) HasEvpnOverlay() bool`

HasEvpnOverlay returns a boolean if a field has been set.

### GetForOverlay

`func (o *BgpStats) GetForOverlay() bool`

GetForOverlay returns the ForOverlay field if non-nil, zero value otherwise.

### GetForOverlayOk

`func (o *BgpStats) GetForOverlayOk() (*bool, bool)`

GetForOverlayOk returns a tuple with the ForOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForOverlay

`func (o *BgpStats) SetForOverlay(v bool)`

SetForOverlay sets ForOverlay field to given value.

### HasForOverlay

`func (o *BgpStats) HasForOverlay() bool`

HasForOverlay returns a boolean if a field has been set.

### GetLocalAs

`func (o *BgpStats) GetLocalAs() int32`

GetLocalAs returns the LocalAs field if non-nil, zero value otherwise.

### GetLocalAsOk

`func (o *BgpStats) GetLocalAsOk() (*int32, bool)`

GetLocalAsOk returns a tuple with the LocalAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAs

`func (o *BgpStats) SetLocalAs(v int32)`

SetLocalAs sets LocalAs field to given value.

### HasLocalAs

`func (o *BgpStats) HasLocalAs() bool`

HasLocalAs returns a boolean if a field has been set.

### GetMac

`func (o *BgpStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *BgpStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *BgpStats) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *BgpStats) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNeighbor

`func (o *BgpStats) GetNeighbor() string`

GetNeighbor returns the Neighbor field if non-nil, zero value otherwise.

### GetNeighborOk

`func (o *BgpStats) GetNeighborOk() (*string, bool)`

GetNeighborOk returns a tuple with the Neighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbor

`func (o *BgpStats) SetNeighbor(v string)`

SetNeighbor sets Neighbor field to given value.

### HasNeighbor

`func (o *BgpStats) HasNeighbor() bool`

HasNeighbor returns a boolean if a field has been set.

### GetNeighborAs

`func (o *BgpStats) GetNeighborAs() int32`

GetNeighborAs returns the NeighborAs field if non-nil, zero value otherwise.

### GetNeighborAsOk

`func (o *BgpStats) GetNeighborAsOk() (*int32, bool)`

GetNeighborAsOk returns a tuple with the NeighborAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborAs

`func (o *BgpStats) SetNeighborAs(v int32)`

SetNeighborAs sets NeighborAs field to given value.

### HasNeighborAs

`func (o *BgpStats) HasNeighborAs() bool`

HasNeighborAs returns a boolean if a field has been set.

### GetNeighborMac

`func (o *BgpStats) GetNeighborMac() string`

GetNeighborMac returns the NeighborMac field if non-nil, zero value otherwise.

### GetNeighborMacOk

`func (o *BgpStats) GetNeighborMacOk() (*string, bool)`

GetNeighborMacOk returns a tuple with the NeighborMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborMac

`func (o *BgpStats) SetNeighborMac(v string)`

SetNeighborMac sets NeighborMac field to given value.

### HasNeighborMac

`func (o *BgpStats) HasNeighborMac() bool`

HasNeighborMac returns a boolean if a field has been set.

### GetNode

`func (o *BgpStats) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *BgpStats) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *BgpStats) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *BgpStats) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetOrgId

`func (o *BgpStats) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *BgpStats) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *BgpStats) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *BgpStats) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRxPkts

`func (o *BgpStats) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *BgpStats) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *BgpStats) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *BgpStats) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetRxRoutes

`func (o *BgpStats) GetRxRoutes() int32`

GetRxRoutes returns the RxRoutes field if non-nil, zero value otherwise.

### GetRxRoutesOk

`func (o *BgpStats) GetRxRoutesOk() (*int32, bool)`

GetRxRoutesOk returns a tuple with the RxRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRoutes

`func (o *BgpStats) SetRxRoutes(v int32)`

SetRxRoutes sets RxRoutes field to given value.

### HasRxRoutes

`func (o *BgpStats) HasRxRoutes() bool`

HasRxRoutes returns a boolean if a field has been set.

### GetSiteId

`func (o *BgpStats) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *BgpStats) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *BgpStats) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *BgpStats) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetState

`func (o *BgpStats) GetState() BgpStatsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BgpStats) GetStateOk() (*BgpStatsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BgpStats) SetState(v BgpStatsState)`

SetState sets State field to given value.

### HasState

`func (o *BgpStats) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimestamp

`func (o *BgpStats) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BgpStats) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BgpStats) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BgpStats) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTxPkts

`func (o *BgpStats) GetTxPkts() int32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *BgpStats) GetTxPktsOk() (*int32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *BgpStats) SetTxPkts(v int32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *BgpStats) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetTxRoutes

`func (o *BgpStats) GetTxRoutes() int32`

GetTxRoutes returns the TxRoutes field if non-nil, zero value otherwise.

### GetTxRoutesOk

`func (o *BgpStats) GetTxRoutesOk() (*int32, bool)`

GetTxRoutesOk returns a tuple with the TxRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRoutes

`func (o *BgpStats) SetTxRoutes(v int32)`

SetTxRoutes sets TxRoutes field to given value.

### HasTxRoutes

`func (o *BgpStats) HasTxRoutes() bool`

HasTxRoutes returns a boolean if a field has been set.

### GetUp

`func (o *BgpStats) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *BgpStats) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *BgpStats) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *BgpStats) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUptime

`func (o *BgpStats) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *BgpStats) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *BgpStats) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *BgpStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVrfName

`func (o *BgpStats) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *BgpStats) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *BgpStats) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *BgpStats) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



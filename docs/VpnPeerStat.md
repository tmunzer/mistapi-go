# VpnPeerStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** | Redundancy status of the associated interface | [optional] 
**LastSeen** | Pointer to **float32** |  | [optional] 
**Latency** | Pointer to **float32** |  | [optional] 
**Mac** | Pointer to **string** | router mac address | [optional] 
**Mos** | Pointer to **float32** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PeerMac** | Pointer to **string** | peer router mac address | [optional] 
**PeerPortId** | Pointer to **string** | peer router device interface | [optional] 
**PeerRouterName** | Pointer to **string** |  | [optional] 
**PeerSiteId** | Pointer to **string** |  | [optional] [readonly] 
**PortId** | Pointer to **string** | router device interface | [optional] 
**RouterName** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** | &#x60;ipsec&#x60;for SRX, &#x60;svr&#x60; for 128T | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**Uptime** | Pointer to **int32** |  | [optional] 

## Methods

### NewVpnPeerStat

`func NewVpnPeerStat() *VpnPeerStat`

NewVpnPeerStat instantiates a new VpnPeerStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPeerStatWithDefaults

`func NewVpnPeerStatWithDefaults() *VpnPeerStat`

NewVpnPeerStatWithDefaults instantiates a new VpnPeerStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *VpnPeerStat) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *VpnPeerStat) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *VpnPeerStat) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *VpnPeerStat) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastSeen

`func (o *VpnPeerStat) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *VpnPeerStat) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *VpnPeerStat) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *VpnPeerStat) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLatency

`func (o *VpnPeerStat) GetLatency() float32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *VpnPeerStat) GetLatencyOk() (*float32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *VpnPeerStat) SetLatency(v float32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *VpnPeerStat) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetMac

`func (o *VpnPeerStat) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VpnPeerStat) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VpnPeerStat) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VpnPeerStat) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMos

`func (o *VpnPeerStat) GetMos() float32`

GetMos returns the Mos field if non-nil, zero value otherwise.

### GetMosOk

`func (o *VpnPeerStat) GetMosOk() (*float32, bool)`

GetMosOk returns a tuple with the Mos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMos

`func (o *VpnPeerStat) SetMos(v float32)`

SetMos sets Mos field to given value.

### HasMos

`func (o *VpnPeerStat) HasMos() bool`

HasMos returns a boolean if a field has been set.

### GetMtu

`func (o *VpnPeerStat) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VpnPeerStat) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VpnPeerStat) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VpnPeerStat) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetOrgId

`func (o *VpnPeerStat) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *VpnPeerStat) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *VpnPeerStat) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *VpnPeerStat) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPeerMac

`func (o *VpnPeerStat) GetPeerMac() string`

GetPeerMac returns the PeerMac field if non-nil, zero value otherwise.

### GetPeerMacOk

`func (o *VpnPeerStat) GetPeerMacOk() (*string, bool)`

GetPeerMacOk returns a tuple with the PeerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerMac

`func (o *VpnPeerStat) SetPeerMac(v string)`

SetPeerMac sets PeerMac field to given value.

### HasPeerMac

`func (o *VpnPeerStat) HasPeerMac() bool`

HasPeerMac returns a boolean if a field has been set.

### GetPeerPortId

`func (o *VpnPeerStat) GetPeerPortId() string`

GetPeerPortId returns the PeerPortId field if non-nil, zero value otherwise.

### GetPeerPortIdOk

`func (o *VpnPeerStat) GetPeerPortIdOk() (*string, bool)`

GetPeerPortIdOk returns a tuple with the PeerPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerPortId

`func (o *VpnPeerStat) SetPeerPortId(v string)`

SetPeerPortId sets PeerPortId field to given value.

### HasPeerPortId

`func (o *VpnPeerStat) HasPeerPortId() bool`

HasPeerPortId returns a boolean if a field has been set.

### GetPeerRouterName

`func (o *VpnPeerStat) GetPeerRouterName() string`

GetPeerRouterName returns the PeerRouterName field if non-nil, zero value otherwise.

### GetPeerRouterNameOk

`func (o *VpnPeerStat) GetPeerRouterNameOk() (*string, bool)`

GetPeerRouterNameOk returns a tuple with the PeerRouterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerRouterName

`func (o *VpnPeerStat) SetPeerRouterName(v string)`

SetPeerRouterName sets PeerRouterName field to given value.

### HasPeerRouterName

`func (o *VpnPeerStat) HasPeerRouterName() bool`

HasPeerRouterName returns a boolean if a field has been set.

### GetPeerSiteId

`func (o *VpnPeerStat) GetPeerSiteId() string`

GetPeerSiteId returns the PeerSiteId field if non-nil, zero value otherwise.

### GetPeerSiteIdOk

`func (o *VpnPeerStat) GetPeerSiteIdOk() (*string, bool)`

GetPeerSiteIdOk returns a tuple with the PeerSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSiteId

`func (o *VpnPeerStat) SetPeerSiteId(v string)`

SetPeerSiteId sets PeerSiteId field to given value.

### HasPeerSiteId

`func (o *VpnPeerStat) HasPeerSiteId() bool`

HasPeerSiteId returns a boolean if a field has been set.

### GetPortId

`func (o *VpnPeerStat) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *VpnPeerStat) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *VpnPeerStat) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *VpnPeerStat) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRouterName

`func (o *VpnPeerStat) GetRouterName() string`

GetRouterName returns the RouterName field if non-nil, zero value otherwise.

### GetRouterNameOk

`func (o *VpnPeerStat) GetRouterNameOk() (*string, bool)`

GetRouterNameOk returns a tuple with the RouterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterName

`func (o *VpnPeerStat) SetRouterName(v string)`

SetRouterName sets RouterName field to given value.

### HasRouterName

`func (o *VpnPeerStat) HasRouterName() bool`

HasRouterName returns a boolean if a field has been set.

### GetSiteId

`func (o *VpnPeerStat) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VpnPeerStat) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VpnPeerStat) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VpnPeerStat) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetType

`func (o *VpnPeerStat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VpnPeerStat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VpnPeerStat) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VpnPeerStat) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUp

`func (o *VpnPeerStat) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *VpnPeerStat) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *VpnPeerStat) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *VpnPeerStat) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUptime

`func (o *VpnPeerStat) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *VpnPeerStat) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *VpnPeerStat) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *VpnPeerStat) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Mxtunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnchorMxtunnelIds** | Pointer to **[]string** | list of anchor mxtunnels used for forming edge to edge tunnels | [optional] 
**AutoPreemption** | Pointer to [**AutoPreemption**](AutoPreemption.md) |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**HelloInterval** | Pointer to **NullableInt32** | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by &#x60;hello_retries&#x60;. | [optional] [default to 60]
**HelloRetries** | Pointer to **NullableInt32** |  | [optional] [default to 7]
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Ipsec** | Pointer to [**MxtunnelIpsec**](MxtunnelIpsec.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Mtu** | Pointer to **int32** | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU | [optional] [default to 0]
**MxclusterIds** | Pointer to **[]string** | list of mxclusters to deploy this tunnel to | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Protocol** | Pointer to [**MxtunnelProtocol**](MxtunnelProtocol.md) |  | [optional] [default to MXTUNNELPROTOCOL_UDP]
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**VlanIds** | Pointer to **[]int32** | list of vlan_ids that will be used | [optional] 

## Methods

### NewMxtunnel

`func NewMxtunnel() *Mxtunnel`

NewMxtunnel instantiates a new Mxtunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxtunnelWithDefaults

`func NewMxtunnelWithDefaults() *Mxtunnel`

NewMxtunnelWithDefaults instantiates a new Mxtunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchorMxtunnelIds

`func (o *Mxtunnel) GetAnchorMxtunnelIds() []string`

GetAnchorMxtunnelIds returns the AnchorMxtunnelIds field if non-nil, zero value otherwise.

### GetAnchorMxtunnelIdsOk

`func (o *Mxtunnel) GetAnchorMxtunnelIdsOk() (*[]string, bool)`

GetAnchorMxtunnelIdsOk returns a tuple with the AnchorMxtunnelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMxtunnelIds

`func (o *Mxtunnel) SetAnchorMxtunnelIds(v []string)`

SetAnchorMxtunnelIds sets AnchorMxtunnelIds field to given value.

### HasAnchorMxtunnelIds

`func (o *Mxtunnel) HasAnchorMxtunnelIds() bool`

HasAnchorMxtunnelIds returns a boolean if a field has been set.

### GetAutoPreemption

`func (o *Mxtunnel) GetAutoPreemption() AutoPreemption`

GetAutoPreemption returns the AutoPreemption field if non-nil, zero value otherwise.

### GetAutoPreemptionOk

`func (o *Mxtunnel) GetAutoPreemptionOk() (*AutoPreemption, bool)`

GetAutoPreemptionOk returns a tuple with the AutoPreemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreemption

`func (o *Mxtunnel) SetAutoPreemption(v AutoPreemption)`

SetAutoPreemption sets AutoPreemption field to given value.

### HasAutoPreemption

`func (o *Mxtunnel) HasAutoPreemption() bool`

HasAutoPreemption returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Mxtunnel) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Mxtunnel) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Mxtunnel) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Mxtunnel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetForSite

`func (o *Mxtunnel) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Mxtunnel) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Mxtunnel) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Mxtunnel) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHelloInterval

`func (o *Mxtunnel) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *Mxtunnel) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *Mxtunnel) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *Mxtunnel) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### SetHelloIntervalNil

`func (o *Mxtunnel) SetHelloIntervalNil(b bool)`

 SetHelloIntervalNil sets the value for HelloInterval to be an explicit nil

### UnsetHelloInterval
`func (o *Mxtunnel) UnsetHelloInterval()`

UnsetHelloInterval ensures that no value is present for HelloInterval, not even an explicit nil
### GetHelloRetries

`func (o *Mxtunnel) GetHelloRetries() int32`

GetHelloRetries returns the HelloRetries field if non-nil, zero value otherwise.

### GetHelloRetriesOk

`func (o *Mxtunnel) GetHelloRetriesOk() (*int32, bool)`

GetHelloRetriesOk returns a tuple with the HelloRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloRetries

`func (o *Mxtunnel) SetHelloRetries(v int32)`

SetHelloRetries sets HelloRetries field to given value.

### HasHelloRetries

`func (o *Mxtunnel) HasHelloRetries() bool`

HasHelloRetries returns a boolean if a field has been set.

### SetHelloRetriesNil

`func (o *Mxtunnel) SetHelloRetriesNil(b bool)`

 SetHelloRetriesNil sets the value for HelloRetries to be an explicit nil

### UnsetHelloRetries
`func (o *Mxtunnel) UnsetHelloRetries()`

UnsetHelloRetries ensures that no value is present for HelloRetries, not even an explicit nil
### GetId

`func (o *Mxtunnel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mxtunnel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mxtunnel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Mxtunnel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpsec

`func (o *Mxtunnel) GetIpsec() MxtunnelIpsec`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *Mxtunnel) GetIpsecOk() (*MxtunnelIpsec, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *Mxtunnel) SetIpsec(v MxtunnelIpsec)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *Mxtunnel) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Mxtunnel) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Mxtunnel) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Mxtunnel) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Mxtunnel) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMtu

`func (o *Mxtunnel) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Mxtunnel) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Mxtunnel) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *Mxtunnel) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetMxclusterIds

`func (o *Mxtunnel) GetMxclusterIds() []string`

GetMxclusterIds returns the MxclusterIds field if non-nil, zero value otherwise.

### GetMxclusterIdsOk

`func (o *Mxtunnel) GetMxclusterIdsOk() (*[]string, bool)`

GetMxclusterIdsOk returns a tuple with the MxclusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxclusterIds

`func (o *Mxtunnel) SetMxclusterIds(v []string)`

SetMxclusterIds sets MxclusterIds field to given value.

### HasMxclusterIds

`func (o *Mxtunnel) HasMxclusterIds() bool`

HasMxclusterIds returns a boolean if a field has been set.

### GetName

`func (o *Mxtunnel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mxtunnel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mxtunnel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mxtunnel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Mxtunnel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Mxtunnel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrgId

`func (o *Mxtunnel) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Mxtunnel) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Mxtunnel) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Mxtunnel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProtocol

`func (o *Mxtunnel) GetProtocol() MxtunnelProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Mxtunnel) GetProtocolOk() (*MxtunnelProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Mxtunnel) SetProtocol(v MxtunnelProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Mxtunnel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSiteId

`func (o *Mxtunnel) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Mxtunnel) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Mxtunnel) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Mxtunnel) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVlanIds

`func (o *Mxtunnel) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *Mxtunnel) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *Mxtunnel) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *Mxtunnel) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SiteMxtunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalMxtunnels** | Pointer to [**map[string]SiteMxtunnelAdditionalMxtunnel**](SiteMxtunnelAdditionalMxtunnel.md) |  | [optional] 
**ApSubnets** | Pointer to **[]string** | list of subnets where we allow AP to establish Mist Tunnels from | [optional] 
**AutoPreemption** | Pointer to [**AutoPreemption**](AutoPreemption.md) |  | [optional] 
**Clusters** | Pointer to [**[]SiteMxtunnelCluster**](SiteMxtunnelCluster.md) | for AP, how to connect to tunterm or radsecproxy | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**HelloInterval** | Pointer to **int32** | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries | [optional] [default to 60]
**HelloRetries** | Pointer to **int32** |  | [optional] [default to 7]
**Hosts** | Pointer to **[]string** | hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP) | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Mtu** | Pointer to **int32** | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU | [optional] [default to 0]
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Protocol** | Pointer to [**MxtunnelProtocol**](MxtunnelProtocol.md) |  | [optional] [default to MXTUNNELPROTOCOL_UDP]
**Radsec** | Pointer to [**SiteMxtunnelRadsec**](SiteMxtunnelRadsec.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**VlanIds** | Pointer to **[]int32** | list of vlan_ids that will be used | [optional] 

## Methods

### NewSiteMxtunnel

`func NewSiteMxtunnel() *SiteMxtunnel`

NewSiteMxtunnel instantiates a new SiteMxtunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteMxtunnelWithDefaults

`func NewSiteMxtunnelWithDefaults() *SiteMxtunnel`

NewSiteMxtunnelWithDefaults instantiates a new SiteMxtunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalMxtunnels

`func (o *SiteMxtunnel) GetAdditionalMxtunnels() map[string]SiteMxtunnelAdditionalMxtunnel`

GetAdditionalMxtunnels returns the AdditionalMxtunnels field if non-nil, zero value otherwise.

### GetAdditionalMxtunnelsOk

`func (o *SiteMxtunnel) GetAdditionalMxtunnelsOk() (*map[string]SiteMxtunnelAdditionalMxtunnel, bool)`

GetAdditionalMxtunnelsOk returns a tuple with the AdditionalMxtunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMxtunnels

`func (o *SiteMxtunnel) SetAdditionalMxtunnels(v map[string]SiteMxtunnelAdditionalMxtunnel)`

SetAdditionalMxtunnels sets AdditionalMxtunnels field to given value.

### HasAdditionalMxtunnels

`func (o *SiteMxtunnel) HasAdditionalMxtunnels() bool`

HasAdditionalMxtunnels returns a boolean if a field has been set.

### GetApSubnets

`func (o *SiteMxtunnel) GetApSubnets() []string`

GetApSubnets returns the ApSubnets field if non-nil, zero value otherwise.

### GetApSubnetsOk

`func (o *SiteMxtunnel) GetApSubnetsOk() (*[]string, bool)`

GetApSubnetsOk returns a tuple with the ApSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApSubnets

`func (o *SiteMxtunnel) SetApSubnets(v []string)`

SetApSubnets sets ApSubnets field to given value.

### HasApSubnets

`func (o *SiteMxtunnel) HasApSubnets() bool`

HasApSubnets returns a boolean if a field has been set.

### GetAutoPreemption

`func (o *SiteMxtunnel) GetAutoPreemption() AutoPreemption`

GetAutoPreemption returns the AutoPreemption field if non-nil, zero value otherwise.

### GetAutoPreemptionOk

`func (o *SiteMxtunnel) GetAutoPreemptionOk() (*AutoPreemption, bool)`

GetAutoPreemptionOk returns a tuple with the AutoPreemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreemption

`func (o *SiteMxtunnel) SetAutoPreemption(v AutoPreemption)`

SetAutoPreemption sets AutoPreemption field to given value.

### HasAutoPreemption

`func (o *SiteMxtunnel) HasAutoPreemption() bool`

HasAutoPreemption returns a boolean if a field has been set.

### GetClusters

`func (o *SiteMxtunnel) GetClusters() []SiteMxtunnelCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SiteMxtunnel) GetClustersOk() (*[]SiteMxtunnelCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SiteMxtunnel) SetClusters(v []SiteMxtunnelCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SiteMxtunnel) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetCreatedTime

`func (o *SiteMxtunnel) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SiteMxtunnel) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SiteMxtunnel) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SiteMxtunnel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEnabled

`func (o *SiteMxtunnel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SiteMxtunnel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SiteMxtunnel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SiteMxtunnel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForSite

`func (o *SiteMxtunnel) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *SiteMxtunnel) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *SiteMxtunnel) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *SiteMxtunnel) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHelloInterval

`func (o *SiteMxtunnel) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *SiteMxtunnel) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *SiteMxtunnel) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *SiteMxtunnel) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetHelloRetries

`func (o *SiteMxtunnel) GetHelloRetries() int32`

GetHelloRetries returns the HelloRetries field if non-nil, zero value otherwise.

### GetHelloRetriesOk

`func (o *SiteMxtunnel) GetHelloRetriesOk() (*int32, bool)`

GetHelloRetriesOk returns a tuple with the HelloRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloRetries

`func (o *SiteMxtunnel) SetHelloRetries(v int32)`

SetHelloRetries sets HelloRetries field to given value.

### HasHelloRetries

`func (o *SiteMxtunnel) HasHelloRetries() bool`

HasHelloRetries returns a boolean if a field has been set.

### GetHosts

`func (o *SiteMxtunnel) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *SiteMxtunnel) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *SiteMxtunnel) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *SiteMxtunnel) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *SiteMxtunnel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteMxtunnel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteMxtunnel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SiteMxtunnel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *SiteMxtunnel) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *SiteMxtunnel) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *SiteMxtunnel) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *SiteMxtunnel) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMtu

`func (o *SiteMxtunnel) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *SiteMxtunnel) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *SiteMxtunnel) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *SiteMxtunnel) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetOrgId

`func (o *SiteMxtunnel) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SiteMxtunnel) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SiteMxtunnel) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SiteMxtunnel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProtocol

`func (o *SiteMxtunnel) GetProtocol() MxtunnelProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SiteMxtunnel) GetProtocolOk() (*MxtunnelProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SiteMxtunnel) SetProtocol(v MxtunnelProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SiteMxtunnel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRadsec

`func (o *SiteMxtunnel) GetRadsec() SiteMxtunnelRadsec`

GetRadsec returns the Radsec field if non-nil, zero value otherwise.

### GetRadsecOk

`func (o *SiteMxtunnel) GetRadsecOk() (*SiteMxtunnelRadsec, bool)`

GetRadsecOk returns a tuple with the Radsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadsec

`func (o *SiteMxtunnel) SetRadsec(v SiteMxtunnelRadsec)`

SetRadsec sets Radsec field to given value.

### HasRadsec

`func (o *SiteMxtunnel) HasRadsec() bool`

HasRadsec returns a boolean if a field has been set.

### GetSiteId

`func (o *SiteMxtunnel) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteMxtunnel) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteMxtunnel) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteMxtunnel) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVlanIds

`func (o *SiteMxtunnel) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *SiteMxtunnel) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *SiteMxtunnel) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *SiteMxtunnel) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



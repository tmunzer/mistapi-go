# SiteMxtunnelAdditionalMxtunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]SiteMxtunnelCluster**](SiteMxtunnelCluster.md) | for AP, how to connect to tunterm or radsecproxy | [optional] 
**HelloInterval** | Pointer to **int32** | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries | [optional] [default to 60]
**HelloRetries** | Pointer to **int32** |  | [optional] [default to 7]
**Protocol** | Pointer to [**SiteMxtunnelProtocol**](SiteMxtunnelProtocol.md) |  | [optional] 
**VlanIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewSiteMxtunnelAdditionalMxtunnel

`func NewSiteMxtunnelAdditionalMxtunnel() *SiteMxtunnelAdditionalMxtunnel`

NewSiteMxtunnelAdditionalMxtunnel instantiates a new SiteMxtunnelAdditionalMxtunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteMxtunnelAdditionalMxtunnelWithDefaults

`func NewSiteMxtunnelAdditionalMxtunnelWithDefaults() *SiteMxtunnelAdditionalMxtunnel`

NewSiteMxtunnelAdditionalMxtunnelWithDefaults instantiates a new SiteMxtunnelAdditionalMxtunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *SiteMxtunnelAdditionalMxtunnel) GetClusters() []SiteMxtunnelCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SiteMxtunnelAdditionalMxtunnel) GetClustersOk() (*[]SiteMxtunnelCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SiteMxtunnelAdditionalMxtunnel) SetClusters(v []SiteMxtunnelCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SiteMxtunnelAdditionalMxtunnel) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetHelloInterval

`func (o *SiteMxtunnelAdditionalMxtunnel) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *SiteMxtunnelAdditionalMxtunnel) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *SiteMxtunnelAdditionalMxtunnel) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *SiteMxtunnelAdditionalMxtunnel) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetHelloRetries

`func (o *SiteMxtunnelAdditionalMxtunnel) GetHelloRetries() int32`

GetHelloRetries returns the HelloRetries field if non-nil, zero value otherwise.

### GetHelloRetriesOk

`func (o *SiteMxtunnelAdditionalMxtunnel) GetHelloRetriesOk() (*int32, bool)`

GetHelloRetriesOk returns a tuple with the HelloRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloRetries

`func (o *SiteMxtunnelAdditionalMxtunnel) SetHelloRetries(v int32)`

SetHelloRetries sets HelloRetries field to given value.

### HasHelloRetries

`func (o *SiteMxtunnelAdditionalMxtunnel) HasHelloRetries() bool`

HasHelloRetries returns a boolean if a field has been set.

### GetProtocol

`func (o *SiteMxtunnelAdditionalMxtunnel) GetProtocol() SiteMxtunnelProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SiteMxtunnelAdditionalMxtunnel) GetProtocolOk() (*SiteMxtunnelProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SiteMxtunnelAdditionalMxtunnel) SetProtocol(v SiteMxtunnelProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SiteMxtunnelAdditionalMxtunnel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetVlanIds

`func (o *SiteMxtunnelAdditionalMxtunnel) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *SiteMxtunnelAdditionalMxtunnel) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *SiteMxtunnelAdditionalMxtunnel) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *SiteMxtunnelAdditionalMxtunnel) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



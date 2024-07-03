# UtilsShowRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value. | [optional] [default to 0]
**Interval** | Pointer to **int32** | rate at which output will refresh | [optional] [default to 0]
**Neighbor** | Pointer to **string** | IP of the neighbor | [optional] 
**Node** | Pointer to [**HaClusterNode**](HaClusterNode.md) |  | [optional] 
**Prefix** | Pointer to **string** | route prefix | [optional] 
**Protocol** | Pointer to [**UtilsShowRouteProtocol**](UtilsShowRouteProtocol.md) |  | [optional] [default to UTILSSHOWROUTEPROTOCOL_BGP]
**Route** | Pointer to **string** | if specified, dump bot received and advertised, if not specified, both will be shown * for SSR, show bgp neighbors 10.250.18.202 received-routes/advertised-routes * for SRX and Switches, show route receive_protocol/advertise_protocol bgp 192.168.255.12 | [optional] 
**Vrf** | Pointer to **string** | VRF name | [optional] 

## Methods

### NewUtilsShowRoute

`func NewUtilsShowRoute() *UtilsShowRoute`

NewUtilsShowRoute instantiates a new UtilsShowRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsShowRouteWithDefaults

`func NewUtilsShowRouteWithDefaults() *UtilsShowRoute`

NewUtilsShowRouteWithDefaults instantiates a new UtilsShowRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *UtilsShowRoute) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UtilsShowRoute) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UtilsShowRoute) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UtilsShowRoute) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetInterval

`func (o *UtilsShowRoute) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UtilsShowRoute) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UtilsShowRoute) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UtilsShowRoute) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetNeighbor

`func (o *UtilsShowRoute) GetNeighbor() string`

GetNeighbor returns the Neighbor field if non-nil, zero value otherwise.

### GetNeighborOk

`func (o *UtilsShowRoute) GetNeighborOk() (*string, bool)`

GetNeighborOk returns a tuple with the Neighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbor

`func (o *UtilsShowRoute) SetNeighbor(v string)`

SetNeighbor sets Neighbor field to given value.

### HasNeighbor

`func (o *UtilsShowRoute) HasNeighbor() bool`

HasNeighbor returns a boolean if a field has been set.

### GetNode

`func (o *UtilsShowRoute) GetNode() HaClusterNode`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *UtilsShowRoute) GetNodeOk() (*HaClusterNode, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *UtilsShowRoute) SetNode(v HaClusterNode)`

SetNode sets Node field to given value.

### HasNode

`func (o *UtilsShowRoute) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPrefix

`func (o *UtilsShowRoute) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UtilsShowRoute) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UtilsShowRoute) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UtilsShowRoute) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProtocol

`func (o *UtilsShowRoute) GetProtocol() UtilsShowRouteProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UtilsShowRoute) GetProtocolOk() (*UtilsShowRouteProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UtilsShowRoute) SetProtocol(v UtilsShowRouteProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UtilsShowRoute) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRoute

`func (o *UtilsShowRoute) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *UtilsShowRoute) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *UtilsShowRoute) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *UtilsShowRoute) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetVrf

`func (o *UtilsShowRoute) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *UtilsShowRoute) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *UtilsShowRoute) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *UtilsShowRoute) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



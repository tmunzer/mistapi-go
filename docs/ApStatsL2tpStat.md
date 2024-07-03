# ApStatsL2tpStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | Pointer to [**[]ApStatsL2tpStatSession**](ApStatsL2tpStatSession.md) | list of sessions | [optional] 
**State** | Pointer to [**L2tpState**](L2tpState.md) |  | [optional] 
**Uptime** | Pointer to **int32** | uptime | [optional] 
**WxtunnelId** | Pointer to **string** | WxlanTunnel ID | [optional] 

## Methods

### NewApStatsL2tpStat

`func NewApStatsL2tpStat() *ApStatsL2tpStat`

NewApStatsL2tpStat instantiates a new ApStatsL2tpStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApStatsL2tpStatWithDefaults

`func NewApStatsL2tpStatWithDefaults() *ApStatsL2tpStat`

NewApStatsL2tpStatWithDefaults instantiates a new ApStatsL2tpStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *ApStatsL2tpStat) GetSessions() []ApStatsL2tpStatSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ApStatsL2tpStat) GetSessionsOk() (*[]ApStatsL2tpStatSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ApStatsL2tpStat) SetSessions(v []ApStatsL2tpStatSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ApStatsL2tpStat) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetState

`func (o *ApStatsL2tpStat) GetState() L2tpState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApStatsL2tpStat) GetStateOk() (*L2tpState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApStatsL2tpStat) SetState(v L2tpState)`

SetState sets State field to given value.

### HasState

`func (o *ApStatsL2tpStat) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUptime

`func (o *ApStatsL2tpStat) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ApStatsL2tpStat) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ApStatsL2tpStat) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ApStatsL2tpStat) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetWxtunnelId

`func (o *ApStatsL2tpStat) GetWxtunnelId() string`

GetWxtunnelId returns the WxtunnelId field if non-nil, zero value otherwise.

### GetWxtunnelIdOk

`func (o *ApStatsL2tpStat) GetWxtunnelIdOk() (*string, bool)`

GetWxtunnelIdOk returns a tuple with the WxtunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxtunnelId

`func (o *ApStatsL2tpStat) SetWxtunnelId(v string)`

SetWxtunnelId sets WxtunnelId field to given value.

### HasWxtunnelId

`func (o *ApStatsL2tpStat) HasWxtunnelId() bool`

HasWxtunnelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



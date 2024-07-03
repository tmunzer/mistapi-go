# MxedgeStatsServiceStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtIp** | Pointer to **string** | external IP from ep-terminator’s point of view. valid only for service having its own cloud connection | [optional] 
**LastSeen** | Pointer to **int32** | timestamp when the last stats is seen (cloud unix time, in second). valid only for service having its own stats or whole system (last among last_seen of all services) | [optional] 
**PackageState** | Pointer to **string** | package/service installation state. | [optional] 
**PackageVersion** | Pointer to **string** | package/service installation state. | [optional] 
**RunningState** | Pointer to **string** | service running state. | [optional] 
**Uptime** | Pointer to **int32** | service uptime. | [optional] 

## Methods

### NewMxedgeStatsServiceStat

`func NewMxedgeStatsServiceStat() *MxedgeStatsServiceStat`

NewMxedgeStatsServiceStat instantiates a new MxedgeStatsServiceStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeStatsServiceStatWithDefaults

`func NewMxedgeStatsServiceStatWithDefaults() *MxedgeStatsServiceStat`

NewMxedgeStatsServiceStatWithDefaults instantiates a new MxedgeStatsServiceStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtIp

`func (o *MxedgeStatsServiceStat) GetExtIp() string`

GetExtIp returns the ExtIp field if non-nil, zero value otherwise.

### GetExtIpOk

`func (o *MxedgeStatsServiceStat) GetExtIpOk() (*string, bool)`

GetExtIpOk returns a tuple with the ExtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtIp

`func (o *MxedgeStatsServiceStat) SetExtIp(v string)`

SetExtIp sets ExtIp field to given value.

### HasExtIp

`func (o *MxedgeStatsServiceStat) HasExtIp() bool`

HasExtIp returns a boolean if a field has been set.

### GetLastSeen

`func (o *MxedgeStatsServiceStat) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *MxedgeStatsServiceStat) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *MxedgeStatsServiceStat) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *MxedgeStatsServiceStat) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetPackageState

`func (o *MxedgeStatsServiceStat) GetPackageState() string`

GetPackageState returns the PackageState field if non-nil, zero value otherwise.

### GetPackageStateOk

`func (o *MxedgeStatsServiceStat) GetPackageStateOk() (*string, bool)`

GetPackageStateOk returns a tuple with the PackageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageState

`func (o *MxedgeStatsServiceStat) SetPackageState(v string)`

SetPackageState sets PackageState field to given value.

### HasPackageState

`func (o *MxedgeStatsServiceStat) HasPackageState() bool`

HasPackageState returns a boolean if a field has been set.

### GetPackageVersion

`func (o *MxedgeStatsServiceStat) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *MxedgeStatsServiceStat) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *MxedgeStatsServiceStat) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *MxedgeStatsServiceStat) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetRunningState

`func (o *MxedgeStatsServiceStat) GetRunningState() string`

GetRunningState returns the RunningState field if non-nil, zero value otherwise.

### GetRunningStateOk

`func (o *MxedgeStatsServiceStat) GetRunningStateOk() (*string, bool)`

GetRunningStateOk returns a tuple with the RunningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningState

`func (o *MxedgeStatsServiceStat) SetRunningState(v string)`

SetRunningState sets RunningState field to given value.

### HasRunningState

`func (o *MxedgeStatsServiceStat) HasRunningState() bool`

HasRunningState returns a boolean if a field has been set.

### GetUptime

`func (o *MxedgeStatsServiceStat) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MxedgeStatsServiceStat) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MxedgeStatsServiceStat) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MxedgeStatsServiceStat) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



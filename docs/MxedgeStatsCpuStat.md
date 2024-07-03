# MxedgeStatsCpuStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | Pointer to [**map[string]CpuStat**](CpuStat.md) |  | [optional] 
**Idle** | Pointer to **int32** | percentage of Idle, Idle/(Idle + Busy) since last sampling | [optional] 
**Interrupt** | Pointer to **int32** | percentage of Interrupt, (Irq + SoftIrq)/(Idle + Busy) since last sampling | [optional] 
**System** | Pointer to **int32** | percentage of System, System/(Idle + Busy) since last sampling | [optional] 
**Usage** | Pointer to **int32** | percentage of load, Busy/(Idle + Busy) since last sampling | [optional] 
**User** | Pointer to **int32** | percentage of User, User/(Idle + Busy) since last sampling | [optional] 

## Methods

### NewMxedgeStatsCpuStat

`func NewMxedgeStatsCpuStat() *MxedgeStatsCpuStat`

NewMxedgeStatsCpuStat instantiates a new MxedgeStatsCpuStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeStatsCpuStatWithDefaults

`func NewMxedgeStatsCpuStatWithDefaults() *MxedgeStatsCpuStat`

NewMxedgeStatsCpuStatWithDefaults instantiates a new MxedgeStatsCpuStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpus

`func (o *MxedgeStatsCpuStat) GetCpus() map[string]CpuStat`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *MxedgeStatsCpuStat) GetCpusOk() (*map[string]CpuStat, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *MxedgeStatsCpuStat) SetCpus(v map[string]CpuStat)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *MxedgeStatsCpuStat) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetIdle

`func (o *MxedgeStatsCpuStat) GetIdle() int32`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *MxedgeStatsCpuStat) GetIdleOk() (*int32, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *MxedgeStatsCpuStat) SetIdle(v int32)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *MxedgeStatsCpuStat) HasIdle() bool`

HasIdle returns a boolean if a field has been set.

### GetInterrupt

`func (o *MxedgeStatsCpuStat) GetInterrupt() int32`

GetInterrupt returns the Interrupt field if non-nil, zero value otherwise.

### GetInterruptOk

`func (o *MxedgeStatsCpuStat) GetInterruptOk() (*int32, bool)`

GetInterruptOk returns a tuple with the Interrupt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterrupt

`func (o *MxedgeStatsCpuStat) SetInterrupt(v int32)`

SetInterrupt sets Interrupt field to given value.

### HasInterrupt

`func (o *MxedgeStatsCpuStat) HasInterrupt() bool`

HasInterrupt returns a boolean if a field has been set.

### GetSystem

`func (o *MxedgeStatsCpuStat) GetSystem() int32`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *MxedgeStatsCpuStat) GetSystemOk() (*int32, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *MxedgeStatsCpuStat) SetSystem(v int32)`

SetSystem sets System field to given value.

### HasSystem

`func (o *MxedgeStatsCpuStat) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetUsage

`func (o *MxedgeStatsCpuStat) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *MxedgeStatsCpuStat) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *MxedgeStatsCpuStat) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *MxedgeStatsCpuStat) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUser

`func (o *MxedgeStatsCpuStat) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MxedgeStatsCpuStat) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MxedgeStatsCpuStat) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *MxedgeStatsCpuStat) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



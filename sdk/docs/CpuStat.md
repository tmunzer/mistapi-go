# CpuStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idle** | Pointer to **float32** | Percentage of CPU time that is idle | [optional] 
**Interrupt** | Pointer to **float32** | Percentage of CPU time being used by interrupts | [optional] 
**LoadAvg** | Pointer to **[]float32** | Load averages for the last 1, 5, and 15 minutes | [optional] 
**System** | Pointer to **float32** | Percentage of CPU time being used by system processes | [optional] 
**User** | Pointer to **float32** | Percentage of CPU time being used by user processe | [optional] 

## Methods

### NewCpuStat

`func NewCpuStat() *CpuStat`

NewCpuStat instantiates a new CpuStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuStatWithDefaults

`func NewCpuStatWithDefaults() *CpuStat`

NewCpuStatWithDefaults instantiates a new CpuStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdle

`func (o *CpuStat) GetIdle() float32`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *CpuStat) GetIdleOk() (*float32, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *CpuStat) SetIdle(v float32)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *CpuStat) HasIdle() bool`

HasIdle returns a boolean if a field has been set.

### GetInterrupt

`func (o *CpuStat) GetInterrupt() float32`

GetInterrupt returns the Interrupt field if non-nil, zero value otherwise.

### GetInterruptOk

`func (o *CpuStat) GetInterruptOk() (*float32, bool)`

GetInterruptOk returns a tuple with the Interrupt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterrupt

`func (o *CpuStat) SetInterrupt(v float32)`

SetInterrupt sets Interrupt field to given value.

### HasInterrupt

`func (o *CpuStat) HasInterrupt() bool`

HasInterrupt returns a boolean if a field has been set.

### GetLoadAvg

`func (o *CpuStat) GetLoadAvg() []float32`

GetLoadAvg returns the LoadAvg field if non-nil, zero value otherwise.

### GetLoadAvgOk

`func (o *CpuStat) GetLoadAvgOk() (*[]float32, bool)`

GetLoadAvgOk returns a tuple with the LoadAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadAvg

`func (o *CpuStat) SetLoadAvg(v []float32)`

SetLoadAvg sets LoadAvg field to given value.

### HasLoadAvg

`func (o *CpuStat) HasLoadAvg() bool`

HasLoadAvg returns a boolean if a field has been set.

### GetSystem

`func (o *CpuStat) GetSystem() float32`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *CpuStat) GetSystemOk() (*float32, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *CpuStat) SetSystem(v float32)`

SetSystem sets System field to given value.

### HasSystem

`func (o *CpuStat) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetUser

`func (o *CpuStat) GetUser() float32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CpuStat) GetUserOk() (*float32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CpuStat) SetUser(v float32)`

SetUser sets User field to given value.

### HasUser

`func (o *CpuStat) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MxedgeStatsMemoryStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | The amount of memory, in kibibytes, that has been used more recently and is usually not reclaimed unless absolutely necessary. | [optional] 
**Available** | Pointer to **int32** | An estimate of how much memory is available for starting new applications, without swapping. | [optional] 
**Buffers** | Pointer to **int32** | The amount, in kibibytes, of temporary storage for raw disk blocks. | [optional] 
**Cached** | Pointer to **int32** | The amount of physical RAM, in kibibytes, used as cache memory. | [optional] 
**Free** | Pointer to **int32** | The amount of physical RAM, in kibibytes, left unused by the system | [optional] 
**Inactive** | Pointer to **int32** | The amount of memory, in kibibytes, that has been used less recently and is more eligible to be reclaimed for other purposes. | [optional] 
**SwapCached** | Pointer to **int32** | The amount of memory, in kibibytes, that has once been moved into swap, then back into the main memory, but still also remains in the swapfile. | [optional] 
**SwapFree** | Pointer to **int32** | The total amount of swap free, in kibibytes. | [optional] 
**SwapTotal** | Pointer to **int32** | The total amount of swap available, in kibibytes. | [optional] 
**Total** | Pointer to **int32** | Total amount of usable RAM, in kibibytes, which is physical RAM minus a number of reserved bits and the kernel binary code | [optional] 
**Usage** | Pointer to **int32** |  | [optional] 

## Methods

### NewMxedgeStatsMemoryStat

`func NewMxedgeStatsMemoryStat() *MxedgeStatsMemoryStat`

NewMxedgeStatsMemoryStat instantiates a new MxedgeStatsMemoryStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeStatsMemoryStatWithDefaults

`func NewMxedgeStatsMemoryStatWithDefaults() *MxedgeStatsMemoryStat`

NewMxedgeStatsMemoryStatWithDefaults instantiates a new MxedgeStatsMemoryStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MxedgeStatsMemoryStat) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MxedgeStatsMemoryStat) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MxedgeStatsMemoryStat) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *MxedgeStatsMemoryStat) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAvailable

`func (o *MxedgeStatsMemoryStat) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *MxedgeStatsMemoryStat) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *MxedgeStatsMemoryStat) SetAvailable(v int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *MxedgeStatsMemoryStat) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBuffers

`func (o *MxedgeStatsMemoryStat) GetBuffers() int32`

GetBuffers returns the Buffers field if non-nil, zero value otherwise.

### GetBuffersOk

`func (o *MxedgeStatsMemoryStat) GetBuffersOk() (*int32, bool)`

GetBuffersOk returns a tuple with the Buffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuffers

`func (o *MxedgeStatsMemoryStat) SetBuffers(v int32)`

SetBuffers sets Buffers field to given value.

### HasBuffers

`func (o *MxedgeStatsMemoryStat) HasBuffers() bool`

HasBuffers returns a boolean if a field has been set.

### GetCached

`func (o *MxedgeStatsMemoryStat) GetCached() int32`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *MxedgeStatsMemoryStat) GetCachedOk() (*int32, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *MxedgeStatsMemoryStat) SetCached(v int32)`

SetCached sets Cached field to given value.

### HasCached

`func (o *MxedgeStatsMemoryStat) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetFree

`func (o *MxedgeStatsMemoryStat) GetFree() int32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *MxedgeStatsMemoryStat) GetFreeOk() (*int32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *MxedgeStatsMemoryStat) SetFree(v int32)`

SetFree sets Free field to given value.

### HasFree

`func (o *MxedgeStatsMemoryStat) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetInactive

`func (o *MxedgeStatsMemoryStat) GetInactive() int32`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *MxedgeStatsMemoryStat) GetInactiveOk() (*int32, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *MxedgeStatsMemoryStat) SetInactive(v int32)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *MxedgeStatsMemoryStat) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetSwapCached

`func (o *MxedgeStatsMemoryStat) GetSwapCached() int32`

GetSwapCached returns the SwapCached field if non-nil, zero value otherwise.

### GetSwapCachedOk

`func (o *MxedgeStatsMemoryStat) GetSwapCachedOk() (*int32, bool)`

GetSwapCachedOk returns a tuple with the SwapCached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapCached

`func (o *MxedgeStatsMemoryStat) SetSwapCached(v int32)`

SetSwapCached sets SwapCached field to given value.

### HasSwapCached

`func (o *MxedgeStatsMemoryStat) HasSwapCached() bool`

HasSwapCached returns a boolean if a field has been set.

### GetSwapFree

`func (o *MxedgeStatsMemoryStat) GetSwapFree() int32`

GetSwapFree returns the SwapFree field if non-nil, zero value otherwise.

### GetSwapFreeOk

`func (o *MxedgeStatsMemoryStat) GetSwapFreeOk() (*int32, bool)`

GetSwapFreeOk returns a tuple with the SwapFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapFree

`func (o *MxedgeStatsMemoryStat) SetSwapFree(v int32)`

SetSwapFree sets SwapFree field to given value.

### HasSwapFree

`func (o *MxedgeStatsMemoryStat) HasSwapFree() bool`

HasSwapFree returns a boolean if a field has been set.

### GetSwapTotal

`func (o *MxedgeStatsMemoryStat) GetSwapTotal() int32`

GetSwapTotal returns the SwapTotal field if non-nil, zero value otherwise.

### GetSwapTotalOk

`func (o *MxedgeStatsMemoryStat) GetSwapTotalOk() (*int32, bool)`

GetSwapTotalOk returns a tuple with the SwapTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapTotal

`func (o *MxedgeStatsMemoryStat) SetSwapTotal(v int32)`

SetSwapTotal sets SwapTotal field to given value.

### HasSwapTotal

`func (o *MxedgeStatsMemoryStat) HasSwapTotal() bool`

HasSwapTotal returns a boolean if a field has been set.

### GetTotal

`func (o *MxedgeStatsMemoryStat) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MxedgeStatsMemoryStat) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MxedgeStatsMemoryStat) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MxedgeStatsMemoryStat) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsage

`func (o *MxedgeStatsMemoryStat) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *MxedgeStatsMemoryStat) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *MxedgeStatsMemoryStat) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *MxedgeStatsMemoryStat) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpgradeSiteDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanaryPhases** | Pointer to **[]int32** | phases for canary deployment. Each phase represents percentage of AP’s that need to be upgraded. default is [1, 10, 50, 100] | [optional] [default to [1,10,50,100]]
**DeviceIds** | Pointer to **[]string** |  | [optional] 
**EnableP2p** | Pointer to **bool** | whether to allow local AP-to-AP FW upgrade | [optional] 
**Force** | Pointer to **bool** | true will force upgrade when requested version is same as running version | [optional] [default to false]
**MaxFailurePercentage** | Pointer to **float32** | percentage of failures allowed across the entire upgrade(not applicable for &#x60;big_bang&#x60;) | [optional] [default to 5]
**MaxFailures** | Pointer to **[]int32** | number of failures allowed within each phase(applicable for &#x60;canary&#x60; or &#x60;rrm&#x60;). Will be used if provided, else max_failure_percentage will be used | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**P2pClusterSize** | Pointer to **int32** |  | [optional] [default to 10]
**P2pParallelism** | Pointer to **int32** | number of parallel p2p download batches to creat | [optional] 
**Reboot** | Pointer to **bool** | Reboot device immediately after upgrade is completed (Available on Junos OS devices) | [optional] [default to false]
**RebootAt** | Pointer to **float32** | reboot start time in epoch seconds, default is &#x60;start_time&#x60; | [optional] 
**RrmFirstBatchPercentage** | Pointer to **int32** | percentage of AP’s that need to be present in the first rrm batch | [optional] 
**RrmMaxBatchPercentage** | Pointer to **int32** | max percentage of AP’s that need to be present in each rrm batch | [optional] 
**RrmMeshUpgrade** | Pointer to **string** | sequential or parallel (default parallel). Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade | [optional] 
**RrmNodeOrder** | Pointer to [**UpgradeSiteDevicesRrmNodeOrder**](UpgradeSiteDevicesRrmNodeOrder.md) |  | [optional] [default to UPGRADESITEDEVICESRRMNODEORDER_FRINGE_TO_CENTER]
**RrmSlowRamp** | Pointer to **bool** | true will make rrm batch sizes slowly ramp up | [optional] 
**Snapshot** | Pointer to **bool** | Perform recovery snapshot after device is rebooted (Available on Junos OS devices) | [optional] [default to false]
**StartTime** | Pointer to **float32** | upgrade start time in epoch seconds, default is now | [optional] 
**Strategy** | Pointer to [**DeviceUpgradeStrategy**](DeviceUpgradeStrategy.md) |  | [optional] [default to DEVICEUPGRADESTRATEGY_BIG_BANG]
**Version** | Pointer to **string** | specific version / stable | [optional] [default to "latest"]

## Methods

### NewUpgradeSiteDevices

`func NewUpgradeSiteDevices() *UpgradeSiteDevices`

NewUpgradeSiteDevices instantiates a new UpgradeSiteDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeSiteDevicesWithDefaults

`func NewUpgradeSiteDevicesWithDefaults() *UpgradeSiteDevices`

NewUpgradeSiteDevicesWithDefaults instantiates a new UpgradeSiteDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanaryPhases

`func (o *UpgradeSiteDevices) GetCanaryPhases() []int32`

GetCanaryPhases returns the CanaryPhases field if non-nil, zero value otherwise.

### GetCanaryPhasesOk

`func (o *UpgradeSiteDevices) GetCanaryPhasesOk() (*[]int32, bool)`

GetCanaryPhasesOk returns a tuple with the CanaryPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryPhases

`func (o *UpgradeSiteDevices) SetCanaryPhases(v []int32)`

SetCanaryPhases sets CanaryPhases field to given value.

### HasCanaryPhases

`func (o *UpgradeSiteDevices) HasCanaryPhases() bool`

HasCanaryPhases returns a boolean if a field has been set.

### GetDeviceIds

`func (o *UpgradeSiteDevices) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *UpgradeSiteDevices) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *UpgradeSiteDevices) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *UpgradeSiteDevices) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetEnableP2p

`func (o *UpgradeSiteDevices) GetEnableP2p() bool`

GetEnableP2p returns the EnableP2p field if non-nil, zero value otherwise.

### GetEnableP2pOk

`func (o *UpgradeSiteDevices) GetEnableP2pOk() (*bool, bool)`

GetEnableP2pOk returns a tuple with the EnableP2p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableP2p

`func (o *UpgradeSiteDevices) SetEnableP2p(v bool)`

SetEnableP2p sets EnableP2p field to given value.

### HasEnableP2p

`func (o *UpgradeSiteDevices) HasEnableP2p() bool`

HasEnableP2p returns a boolean if a field has been set.

### GetForce

`func (o *UpgradeSiteDevices) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *UpgradeSiteDevices) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *UpgradeSiteDevices) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *UpgradeSiteDevices) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetMaxFailurePercentage

`func (o *UpgradeSiteDevices) GetMaxFailurePercentage() float32`

GetMaxFailurePercentage returns the MaxFailurePercentage field if non-nil, zero value otherwise.

### GetMaxFailurePercentageOk

`func (o *UpgradeSiteDevices) GetMaxFailurePercentageOk() (*float32, bool)`

GetMaxFailurePercentageOk returns a tuple with the MaxFailurePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailurePercentage

`func (o *UpgradeSiteDevices) SetMaxFailurePercentage(v float32)`

SetMaxFailurePercentage sets MaxFailurePercentage field to given value.

### HasMaxFailurePercentage

`func (o *UpgradeSiteDevices) HasMaxFailurePercentage() bool`

HasMaxFailurePercentage returns a boolean if a field has been set.

### GetMaxFailures

`func (o *UpgradeSiteDevices) GetMaxFailures() []int32`

GetMaxFailures returns the MaxFailures field if non-nil, zero value otherwise.

### GetMaxFailuresOk

`func (o *UpgradeSiteDevices) GetMaxFailuresOk() (*[]int32, bool)`

GetMaxFailuresOk returns a tuple with the MaxFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailures

`func (o *UpgradeSiteDevices) SetMaxFailures(v []int32)`

SetMaxFailures sets MaxFailures field to given value.

### HasMaxFailures

`func (o *UpgradeSiteDevices) HasMaxFailures() bool`

HasMaxFailures returns a boolean if a field has been set.

### GetModels

`func (o *UpgradeSiteDevices) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *UpgradeSiteDevices) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *UpgradeSiteDevices) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *UpgradeSiteDevices) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetP2pClusterSize

`func (o *UpgradeSiteDevices) GetP2pClusterSize() int32`

GetP2pClusterSize returns the P2pClusterSize field if non-nil, zero value otherwise.

### GetP2pClusterSizeOk

`func (o *UpgradeSiteDevices) GetP2pClusterSizeOk() (*int32, bool)`

GetP2pClusterSizeOk returns a tuple with the P2pClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pClusterSize

`func (o *UpgradeSiteDevices) SetP2pClusterSize(v int32)`

SetP2pClusterSize sets P2pClusterSize field to given value.

### HasP2pClusterSize

`func (o *UpgradeSiteDevices) HasP2pClusterSize() bool`

HasP2pClusterSize returns a boolean if a field has been set.

### GetP2pParallelism

`func (o *UpgradeSiteDevices) GetP2pParallelism() int32`

GetP2pParallelism returns the P2pParallelism field if non-nil, zero value otherwise.

### GetP2pParallelismOk

`func (o *UpgradeSiteDevices) GetP2pParallelismOk() (*int32, bool)`

GetP2pParallelismOk returns a tuple with the P2pParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pParallelism

`func (o *UpgradeSiteDevices) SetP2pParallelism(v int32)`

SetP2pParallelism sets P2pParallelism field to given value.

### HasP2pParallelism

`func (o *UpgradeSiteDevices) HasP2pParallelism() bool`

HasP2pParallelism returns a boolean if a field has been set.

### GetReboot

`func (o *UpgradeSiteDevices) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *UpgradeSiteDevices) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *UpgradeSiteDevices) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *UpgradeSiteDevices) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRebootAt

`func (o *UpgradeSiteDevices) GetRebootAt() float32`

GetRebootAt returns the RebootAt field if non-nil, zero value otherwise.

### GetRebootAtOk

`func (o *UpgradeSiteDevices) GetRebootAtOk() (*float32, bool)`

GetRebootAtOk returns a tuple with the RebootAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootAt

`func (o *UpgradeSiteDevices) SetRebootAt(v float32)`

SetRebootAt sets RebootAt field to given value.

### HasRebootAt

`func (o *UpgradeSiteDevices) HasRebootAt() bool`

HasRebootAt returns a boolean if a field has been set.

### GetRrmFirstBatchPercentage

`func (o *UpgradeSiteDevices) GetRrmFirstBatchPercentage() int32`

GetRrmFirstBatchPercentage returns the RrmFirstBatchPercentage field if non-nil, zero value otherwise.

### GetRrmFirstBatchPercentageOk

`func (o *UpgradeSiteDevices) GetRrmFirstBatchPercentageOk() (*int32, bool)`

GetRrmFirstBatchPercentageOk returns a tuple with the RrmFirstBatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmFirstBatchPercentage

`func (o *UpgradeSiteDevices) SetRrmFirstBatchPercentage(v int32)`

SetRrmFirstBatchPercentage sets RrmFirstBatchPercentage field to given value.

### HasRrmFirstBatchPercentage

`func (o *UpgradeSiteDevices) HasRrmFirstBatchPercentage() bool`

HasRrmFirstBatchPercentage returns a boolean if a field has been set.

### GetRrmMaxBatchPercentage

`func (o *UpgradeSiteDevices) GetRrmMaxBatchPercentage() int32`

GetRrmMaxBatchPercentage returns the RrmMaxBatchPercentage field if non-nil, zero value otherwise.

### GetRrmMaxBatchPercentageOk

`func (o *UpgradeSiteDevices) GetRrmMaxBatchPercentageOk() (*int32, bool)`

GetRrmMaxBatchPercentageOk returns a tuple with the RrmMaxBatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmMaxBatchPercentage

`func (o *UpgradeSiteDevices) SetRrmMaxBatchPercentage(v int32)`

SetRrmMaxBatchPercentage sets RrmMaxBatchPercentage field to given value.

### HasRrmMaxBatchPercentage

`func (o *UpgradeSiteDevices) HasRrmMaxBatchPercentage() bool`

HasRrmMaxBatchPercentage returns a boolean if a field has been set.

### GetRrmMeshUpgrade

`func (o *UpgradeSiteDevices) GetRrmMeshUpgrade() string`

GetRrmMeshUpgrade returns the RrmMeshUpgrade field if non-nil, zero value otherwise.

### GetRrmMeshUpgradeOk

`func (o *UpgradeSiteDevices) GetRrmMeshUpgradeOk() (*string, bool)`

GetRrmMeshUpgradeOk returns a tuple with the RrmMeshUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmMeshUpgrade

`func (o *UpgradeSiteDevices) SetRrmMeshUpgrade(v string)`

SetRrmMeshUpgrade sets RrmMeshUpgrade field to given value.

### HasRrmMeshUpgrade

`func (o *UpgradeSiteDevices) HasRrmMeshUpgrade() bool`

HasRrmMeshUpgrade returns a boolean if a field has been set.

### GetRrmNodeOrder

`func (o *UpgradeSiteDevices) GetRrmNodeOrder() UpgradeSiteDevicesRrmNodeOrder`

GetRrmNodeOrder returns the RrmNodeOrder field if non-nil, zero value otherwise.

### GetRrmNodeOrderOk

`func (o *UpgradeSiteDevices) GetRrmNodeOrderOk() (*UpgradeSiteDevicesRrmNodeOrder, bool)`

GetRrmNodeOrderOk returns a tuple with the RrmNodeOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmNodeOrder

`func (o *UpgradeSiteDevices) SetRrmNodeOrder(v UpgradeSiteDevicesRrmNodeOrder)`

SetRrmNodeOrder sets RrmNodeOrder field to given value.

### HasRrmNodeOrder

`func (o *UpgradeSiteDevices) HasRrmNodeOrder() bool`

HasRrmNodeOrder returns a boolean if a field has been set.

### GetRrmSlowRamp

`func (o *UpgradeSiteDevices) GetRrmSlowRamp() bool`

GetRrmSlowRamp returns the RrmSlowRamp field if non-nil, zero value otherwise.

### GetRrmSlowRampOk

`func (o *UpgradeSiteDevices) GetRrmSlowRampOk() (*bool, bool)`

GetRrmSlowRampOk returns a tuple with the RrmSlowRamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmSlowRamp

`func (o *UpgradeSiteDevices) SetRrmSlowRamp(v bool)`

SetRrmSlowRamp sets RrmSlowRamp field to given value.

### HasRrmSlowRamp

`func (o *UpgradeSiteDevices) HasRrmSlowRamp() bool`

HasRrmSlowRamp returns a boolean if a field has been set.

### GetSnapshot

`func (o *UpgradeSiteDevices) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *UpgradeSiteDevices) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *UpgradeSiteDevices) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *UpgradeSiteDevices) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetStartTime

`func (o *UpgradeSiteDevices) GetStartTime() float32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpgradeSiteDevices) GetStartTimeOk() (*float32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpgradeSiteDevices) SetStartTime(v float32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpgradeSiteDevices) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStrategy

`func (o *UpgradeSiteDevices) GetStrategy() DeviceUpgradeStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *UpgradeSiteDevices) GetStrategyOk() (*DeviceUpgradeStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *UpgradeSiteDevices) SetStrategy(v DeviceUpgradeStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *UpgradeSiteDevices) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetVersion

`func (o *UpgradeSiteDevices) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpgradeSiteDevices) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpgradeSiteDevices) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpgradeSiteDevices) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



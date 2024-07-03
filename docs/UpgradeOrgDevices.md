# UpgradeOrgDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanaryPhases** | Pointer to **[]int32** | phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. | [optional] [default to [1,10,50,100]]
**EnableP2p** | Pointer to **bool** | whether to allow local AP-to-AP FW upgrade | [optional] 
**Force** | Pointer to **bool** | true will force upgrade when requested version is same as running version | [optional] [default to false]
**MaxFailurePercentage** | Pointer to **float32** | percentage of failures allowed across the entire upgrade(not applicable for &#x60;big_bang&#x60;) | [optional] [default to 5]
**MaxFailures** | Pointer to **[]int32** | number of failures allowed within each phase. Only applicable for canary. Array length should be same as &#x60;canary_phases&#x60;&#x60;. Will be used if provided, else max_failure_percentage will be used. | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**P2pClusterSize** | Pointer to **int32** |  | [optional] [default to 10]
**P2pParallelism** | Pointer to **int32** | number of parallel p2p download batches to creat | [optional] 
**Reboot** | Pointer to **bool** | Reboot device immediately after upgrade is completed (Available on Junos OS devices) | [optional] [default to false]
**RebootAt** | Pointer to **float32** | reboot start time in epoch seconds, default is &#x60;start_time&#x60; | [optional] 
**RrmFirstBatchPercentage** | Pointer to **int32** | percentage of AP’s that need to be present in the first rrm batch | [optional] 
**RrmMaxBatchPercentage** | Pointer to **int32** | max percentage of AP’s that need to be present in each rrm batch | [optional] 
**RrmMeshUpgrade** | Pointer to [**DeviceUpgradeRrmMeshUpgrade**](DeviceUpgradeRrmMeshUpgrade.md) |  | [optional] [default to DEVICEUPGRADERRMMESHUPGRADE_SEQUENTIAL]
**RrmNodeOrder** | Pointer to [**DeviceUpgradeRrmNodeOrder**](DeviceUpgradeRrmNodeOrder.md) |  | [optional] [default to DEVICEUPGRADERRMNODEORDER_FRINGE_TO_CENTER]
**RrmSlowRamp** | Pointer to **bool** | true will make rrm batch sizes slowly ramp up | [optional] 
**SiteIds** | Pointer to **[]string** |  | [optional] 
**Snapshot** | Pointer to **bool** | Perform recovery snapshot after device is rebooted (Available on Junos OS devices) | [optional] [default to false]
**StartTime** | Pointer to **float32** | upgrade start time in epoch seconds, default is now | [optional] 
**Strategy** | Pointer to [**DeviceUpgradeStrategy**](DeviceUpgradeStrategy.md) |  | [optional] [default to DEVICEUPGRADESTRATEGY_BIG_BANG]
**Version** | Pointer to **string** | specific version / stable | [optional] [default to "latest"]

## Methods

### NewUpgradeOrgDevices

`func NewUpgradeOrgDevices() *UpgradeOrgDevices`

NewUpgradeOrgDevices instantiates a new UpgradeOrgDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeOrgDevicesWithDefaults

`func NewUpgradeOrgDevicesWithDefaults() *UpgradeOrgDevices`

NewUpgradeOrgDevicesWithDefaults instantiates a new UpgradeOrgDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanaryPhases

`func (o *UpgradeOrgDevices) GetCanaryPhases() []int32`

GetCanaryPhases returns the CanaryPhases field if non-nil, zero value otherwise.

### GetCanaryPhasesOk

`func (o *UpgradeOrgDevices) GetCanaryPhasesOk() (*[]int32, bool)`

GetCanaryPhasesOk returns a tuple with the CanaryPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryPhases

`func (o *UpgradeOrgDevices) SetCanaryPhases(v []int32)`

SetCanaryPhases sets CanaryPhases field to given value.

### HasCanaryPhases

`func (o *UpgradeOrgDevices) HasCanaryPhases() bool`

HasCanaryPhases returns a boolean if a field has been set.

### GetEnableP2p

`func (o *UpgradeOrgDevices) GetEnableP2p() bool`

GetEnableP2p returns the EnableP2p field if non-nil, zero value otherwise.

### GetEnableP2pOk

`func (o *UpgradeOrgDevices) GetEnableP2pOk() (*bool, bool)`

GetEnableP2pOk returns a tuple with the EnableP2p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableP2p

`func (o *UpgradeOrgDevices) SetEnableP2p(v bool)`

SetEnableP2p sets EnableP2p field to given value.

### HasEnableP2p

`func (o *UpgradeOrgDevices) HasEnableP2p() bool`

HasEnableP2p returns a boolean if a field has been set.

### GetForce

`func (o *UpgradeOrgDevices) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *UpgradeOrgDevices) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *UpgradeOrgDevices) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *UpgradeOrgDevices) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetMaxFailurePercentage

`func (o *UpgradeOrgDevices) GetMaxFailurePercentage() float32`

GetMaxFailurePercentage returns the MaxFailurePercentage field if non-nil, zero value otherwise.

### GetMaxFailurePercentageOk

`func (o *UpgradeOrgDevices) GetMaxFailurePercentageOk() (*float32, bool)`

GetMaxFailurePercentageOk returns a tuple with the MaxFailurePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailurePercentage

`func (o *UpgradeOrgDevices) SetMaxFailurePercentage(v float32)`

SetMaxFailurePercentage sets MaxFailurePercentage field to given value.

### HasMaxFailurePercentage

`func (o *UpgradeOrgDevices) HasMaxFailurePercentage() bool`

HasMaxFailurePercentage returns a boolean if a field has been set.

### GetMaxFailures

`func (o *UpgradeOrgDevices) GetMaxFailures() []int32`

GetMaxFailures returns the MaxFailures field if non-nil, zero value otherwise.

### GetMaxFailuresOk

`func (o *UpgradeOrgDevices) GetMaxFailuresOk() (*[]int32, bool)`

GetMaxFailuresOk returns a tuple with the MaxFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailures

`func (o *UpgradeOrgDevices) SetMaxFailures(v []int32)`

SetMaxFailures sets MaxFailures field to given value.

### HasMaxFailures

`func (o *UpgradeOrgDevices) HasMaxFailures() bool`

HasMaxFailures returns a boolean if a field has been set.

### GetModels

`func (o *UpgradeOrgDevices) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *UpgradeOrgDevices) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *UpgradeOrgDevices) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *UpgradeOrgDevices) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetP2pClusterSize

`func (o *UpgradeOrgDevices) GetP2pClusterSize() int32`

GetP2pClusterSize returns the P2pClusterSize field if non-nil, zero value otherwise.

### GetP2pClusterSizeOk

`func (o *UpgradeOrgDevices) GetP2pClusterSizeOk() (*int32, bool)`

GetP2pClusterSizeOk returns a tuple with the P2pClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pClusterSize

`func (o *UpgradeOrgDevices) SetP2pClusterSize(v int32)`

SetP2pClusterSize sets P2pClusterSize field to given value.

### HasP2pClusterSize

`func (o *UpgradeOrgDevices) HasP2pClusterSize() bool`

HasP2pClusterSize returns a boolean if a field has been set.

### GetP2pParallelism

`func (o *UpgradeOrgDevices) GetP2pParallelism() int32`

GetP2pParallelism returns the P2pParallelism field if non-nil, zero value otherwise.

### GetP2pParallelismOk

`func (o *UpgradeOrgDevices) GetP2pParallelismOk() (*int32, bool)`

GetP2pParallelismOk returns a tuple with the P2pParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pParallelism

`func (o *UpgradeOrgDevices) SetP2pParallelism(v int32)`

SetP2pParallelism sets P2pParallelism field to given value.

### HasP2pParallelism

`func (o *UpgradeOrgDevices) HasP2pParallelism() bool`

HasP2pParallelism returns a boolean if a field has been set.

### GetReboot

`func (o *UpgradeOrgDevices) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *UpgradeOrgDevices) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *UpgradeOrgDevices) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *UpgradeOrgDevices) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRebootAt

`func (o *UpgradeOrgDevices) GetRebootAt() float32`

GetRebootAt returns the RebootAt field if non-nil, zero value otherwise.

### GetRebootAtOk

`func (o *UpgradeOrgDevices) GetRebootAtOk() (*float32, bool)`

GetRebootAtOk returns a tuple with the RebootAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootAt

`func (o *UpgradeOrgDevices) SetRebootAt(v float32)`

SetRebootAt sets RebootAt field to given value.

### HasRebootAt

`func (o *UpgradeOrgDevices) HasRebootAt() bool`

HasRebootAt returns a boolean if a field has been set.

### GetRrmFirstBatchPercentage

`func (o *UpgradeOrgDevices) GetRrmFirstBatchPercentage() int32`

GetRrmFirstBatchPercentage returns the RrmFirstBatchPercentage field if non-nil, zero value otherwise.

### GetRrmFirstBatchPercentageOk

`func (o *UpgradeOrgDevices) GetRrmFirstBatchPercentageOk() (*int32, bool)`

GetRrmFirstBatchPercentageOk returns a tuple with the RrmFirstBatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmFirstBatchPercentage

`func (o *UpgradeOrgDevices) SetRrmFirstBatchPercentage(v int32)`

SetRrmFirstBatchPercentage sets RrmFirstBatchPercentage field to given value.

### HasRrmFirstBatchPercentage

`func (o *UpgradeOrgDevices) HasRrmFirstBatchPercentage() bool`

HasRrmFirstBatchPercentage returns a boolean if a field has been set.

### GetRrmMaxBatchPercentage

`func (o *UpgradeOrgDevices) GetRrmMaxBatchPercentage() int32`

GetRrmMaxBatchPercentage returns the RrmMaxBatchPercentage field if non-nil, zero value otherwise.

### GetRrmMaxBatchPercentageOk

`func (o *UpgradeOrgDevices) GetRrmMaxBatchPercentageOk() (*int32, bool)`

GetRrmMaxBatchPercentageOk returns a tuple with the RrmMaxBatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmMaxBatchPercentage

`func (o *UpgradeOrgDevices) SetRrmMaxBatchPercentage(v int32)`

SetRrmMaxBatchPercentage sets RrmMaxBatchPercentage field to given value.

### HasRrmMaxBatchPercentage

`func (o *UpgradeOrgDevices) HasRrmMaxBatchPercentage() bool`

HasRrmMaxBatchPercentage returns a boolean if a field has been set.

### GetRrmMeshUpgrade

`func (o *UpgradeOrgDevices) GetRrmMeshUpgrade() DeviceUpgradeRrmMeshUpgrade`

GetRrmMeshUpgrade returns the RrmMeshUpgrade field if non-nil, zero value otherwise.

### GetRrmMeshUpgradeOk

`func (o *UpgradeOrgDevices) GetRrmMeshUpgradeOk() (*DeviceUpgradeRrmMeshUpgrade, bool)`

GetRrmMeshUpgradeOk returns a tuple with the RrmMeshUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmMeshUpgrade

`func (o *UpgradeOrgDevices) SetRrmMeshUpgrade(v DeviceUpgradeRrmMeshUpgrade)`

SetRrmMeshUpgrade sets RrmMeshUpgrade field to given value.

### HasRrmMeshUpgrade

`func (o *UpgradeOrgDevices) HasRrmMeshUpgrade() bool`

HasRrmMeshUpgrade returns a boolean if a field has been set.

### GetRrmNodeOrder

`func (o *UpgradeOrgDevices) GetRrmNodeOrder() DeviceUpgradeRrmNodeOrder`

GetRrmNodeOrder returns the RrmNodeOrder field if non-nil, zero value otherwise.

### GetRrmNodeOrderOk

`func (o *UpgradeOrgDevices) GetRrmNodeOrderOk() (*DeviceUpgradeRrmNodeOrder, bool)`

GetRrmNodeOrderOk returns a tuple with the RrmNodeOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmNodeOrder

`func (o *UpgradeOrgDevices) SetRrmNodeOrder(v DeviceUpgradeRrmNodeOrder)`

SetRrmNodeOrder sets RrmNodeOrder field to given value.

### HasRrmNodeOrder

`func (o *UpgradeOrgDevices) HasRrmNodeOrder() bool`

HasRrmNodeOrder returns a boolean if a field has been set.

### GetRrmSlowRamp

`func (o *UpgradeOrgDevices) GetRrmSlowRamp() bool`

GetRrmSlowRamp returns the RrmSlowRamp field if non-nil, zero value otherwise.

### GetRrmSlowRampOk

`func (o *UpgradeOrgDevices) GetRrmSlowRampOk() (*bool, bool)`

GetRrmSlowRampOk returns a tuple with the RrmSlowRamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrmSlowRamp

`func (o *UpgradeOrgDevices) SetRrmSlowRamp(v bool)`

SetRrmSlowRamp sets RrmSlowRamp field to given value.

### HasRrmSlowRamp

`func (o *UpgradeOrgDevices) HasRrmSlowRamp() bool`

HasRrmSlowRamp returns a boolean if a field has been set.

### GetSiteIds

`func (o *UpgradeOrgDevices) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *UpgradeOrgDevices) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *UpgradeOrgDevices) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *UpgradeOrgDevices) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### GetSnapshot

`func (o *UpgradeOrgDevices) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *UpgradeOrgDevices) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *UpgradeOrgDevices) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *UpgradeOrgDevices) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetStartTime

`func (o *UpgradeOrgDevices) GetStartTime() float32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpgradeOrgDevices) GetStartTimeOk() (*float32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpgradeOrgDevices) SetStartTime(v float32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpgradeOrgDevices) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStrategy

`func (o *UpgradeOrgDevices) GetStrategy() DeviceUpgradeStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *UpgradeOrgDevices) GetStrategyOk() (*DeviceUpgradeStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *UpgradeOrgDevices) SetStrategy(v DeviceUpgradeStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *UpgradeOrgDevices) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetVersion

`func (o *UpgradeOrgDevices) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpgradeOrgDevices) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpgradeOrgDevices) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpgradeOrgDevices) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



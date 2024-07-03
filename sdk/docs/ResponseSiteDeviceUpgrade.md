# ResponseSiteDeviceUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counts** | Pointer to [**ResponseSiteDeviceUpgradeCounts**](ResponseSiteDeviceUpgradeCounts.md) |  | [optional] 
**CurrentPhase** | Pointer to **int32** | current canary or rrm phase in progress | [optional] 
**EnableP2p** | Pointer to **bool** | whether to allow local AP-to-AP FW upgrade | [optional] 
**Force** | Pointer to **bool** | whether to force upgrade when requested version is same as running version | [optional] 
**Id** | **string** | unique id for the upgrade | [readonly] 
**MaxFailurePercentage** | Pointer to **int32** | percentage of failures allowed | [optional] 
**MaxFailures** | Pointer to **[]int32** | number of failures allowed within a canary phase or serial rollout | [optional] 
**RebootAt** | Pointer to **int32** | reboot start time in epoch | [optional] 
**StartTime** | Pointer to **float32** | firmware download start time in epoch | [optional] 
**Status** | Pointer to [**DeviceUpgradeStatus**](DeviceUpgradeStatus.md) |  | [optional] 
**Strategy** | Pointer to [**DeviceUpgradeStrategy**](DeviceUpgradeStrategy.md) |  | [optional] [default to DEVICEUPGRADESTRATEGY_BIG_BANG]
**TargetVersion** | Pointer to **string** | version to upgrade to | [optional] 
**UpgradePlan** | Pointer to **map[string]interface{}** | a dictionary of rrm phase number to devices part of that phase | [optional] 

## Methods

### NewResponseSiteDeviceUpgrade

`func NewResponseSiteDeviceUpgrade(id string, ) *ResponseSiteDeviceUpgrade`

NewResponseSiteDeviceUpgrade instantiates a new ResponseSiteDeviceUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSiteDeviceUpgradeWithDefaults

`func NewResponseSiteDeviceUpgradeWithDefaults() *ResponseSiteDeviceUpgrade`

NewResponseSiteDeviceUpgradeWithDefaults instantiates a new ResponseSiteDeviceUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounts

`func (o *ResponseSiteDeviceUpgrade) GetCounts() ResponseSiteDeviceUpgradeCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *ResponseSiteDeviceUpgrade) GetCountsOk() (*ResponseSiteDeviceUpgradeCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *ResponseSiteDeviceUpgrade) SetCounts(v ResponseSiteDeviceUpgradeCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *ResponseSiteDeviceUpgrade) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetCurrentPhase

`func (o *ResponseSiteDeviceUpgrade) GetCurrentPhase() int32`

GetCurrentPhase returns the CurrentPhase field if non-nil, zero value otherwise.

### GetCurrentPhaseOk

`func (o *ResponseSiteDeviceUpgrade) GetCurrentPhaseOk() (*int32, bool)`

GetCurrentPhaseOk returns a tuple with the CurrentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPhase

`func (o *ResponseSiteDeviceUpgrade) SetCurrentPhase(v int32)`

SetCurrentPhase sets CurrentPhase field to given value.

### HasCurrentPhase

`func (o *ResponseSiteDeviceUpgrade) HasCurrentPhase() bool`

HasCurrentPhase returns a boolean if a field has been set.

### GetEnableP2p

`func (o *ResponseSiteDeviceUpgrade) GetEnableP2p() bool`

GetEnableP2p returns the EnableP2p field if non-nil, zero value otherwise.

### GetEnableP2pOk

`func (o *ResponseSiteDeviceUpgrade) GetEnableP2pOk() (*bool, bool)`

GetEnableP2pOk returns a tuple with the EnableP2p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableP2p

`func (o *ResponseSiteDeviceUpgrade) SetEnableP2p(v bool)`

SetEnableP2p sets EnableP2p field to given value.

### HasEnableP2p

`func (o *ResponseSiteDeviceUpgrade) HasEnableP2p() bool`

HasEnableP2p returns a boolean if a field has been set.

### GetForce

`func (o *ResponseSiteDeviceUpgrade) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *ResponseSiteDeviceUpgrade) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *ResponseSiteDeviceUpgrade) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *ResponseSiteDeviceUpgrade) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetId

`func (o *ResponseSiteDeviceUpgrade) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseSiteDeviceUpgrade) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseSiteDeviceUpgrade) SetId(v string)`

SetId sets Id field to given value.


### GetMaxFailurePercentage

`func (o *ResponseSiteDeviceUpgrade) GetMaxFailurePercentage() int32`

GetMaxFailurePercentage returns the MaxFailurePercentage field if non-nil, zero value otherwise.

### GetMaxFailurePercentageOk

`func (o *ResponseSiteDeviceUpgrade) GetMaxFailurePercentageOk() (*int32, bool)`

GetMaxFailurePercentageOk returns a tuple with the MaxFailurePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailurePercentage

`func (o *ResponseSiteDeviceUpgrade) SetMaxFailurePercentage(v int32)`

SetMaxFailurePercentage sets MaxFailurePercentage field to given value.

### HasMaxFailurePercentage

`func (o *ResponseSiteDeviceUpgrade) HasMaxFailurePercentage() bool`

HasMaxFailurePercentage returns a boolean if a field has been set.

### GetMaxFailures

`func (o *ResponseSiteDeviceUpgrade) GetMaxFailures() []int32`

GetMaxFailures returns the MaxFailures field if non-nil, zero value otherwise.

### GetMaxFailuresOk

`func (o *ResponseSiteDeviceUpgrade) GetMaxFailuresOk() (*[]int32, bool)`

GetMaxFailuresOk returns a tuple with the MaxFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailures

`func (o *ResponseSiteDeviceUpgrade) SetMaxFailures(v []int32)`

SetMaxFailures sets MaxFailures field to given value.

### HasMaxFailures

`func (o *ResponseSiteDeviceUpgrade) HasMaxFailures() bool`

HasMaxFailures returns a boolean if a field has been set.

### GetRebootAt

`func (o *ResponseSiteDeviceUpgrade) GetRebootAt() int32`

GetRebootAt returns the RebootAt field if non-nil, zero value otherwise.

### GetRebootAtOk

`func (o *ResponseSiteDeviceUpgrade) GetRebootAtOk() (*int32, bool)`

GetRebootAtOk returns a tuple with the RebootAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootAt

`func (o *ResponseSiteDeviceUpgrade) SetRebootAt(v int32)`

SetRebootAt sets RebootAt field to given value.

### HasRebootAt

`func (o *ResponseSiteDeviceUpgrade) HasRebootAt() bool`

HasRebootAt returns a boolean if a field has been set.

### GetStartTime

`func (o *ResponseSiteDeviceUpgrade) GetStartTime() float32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ResponseSiteDeviceUpgrade) GetStartTimeOk() (*float32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ResponseSiteDeviceUpgrade) SetStartTime(v float32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ResponseSiteDeviceUpgrade) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseSiteDeviceUpgrade) GetStatus() DeviceUpgradeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseSiteDeviceUpgrade) GetStatusOk() (*DeviceUpgradeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseSiteDeviceUpgrade) SetStatus(v DeviceUpgradeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseSiteDeviceUpgrade) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStrategy

`func (o *ResponseSiteDeviceUpgrade) GetStrategy() DeviceUpgradeStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *ResponseSiteDeviceUpgrade) GetStrategyOk() (*DeviceUpgradeStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *ResponseSiteDeviceUpgrade) SetStrategy(v DeviceUpgradeStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *ResponseSiteDeviceUpgrade) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetTargetVersion

`func (o *ResponseSiteDeviceUpgrade) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *ResponseSiteDeviceUpgrade) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *ResponseSiteDeviceUpgrade) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *ResponseSiteDeviceUpgrade) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetUpgradePlan

`func (o *ResponseSiteDeviceUpgrade) GetUpgradePlan() map[string]interface{}`

GetUpgradePlan returns the UpgradePlan field if non-nil, zero value otherwise.

### GetUpgradePlanOk

`func (o *ResponseSiteDeviceUpgrade) GetUpgradePlanOk() (*map[string]interface{}, bool)`

GetUpgradePlanOk returns a tuple with the UpgradePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePlan

`func (o *ResponseSiteDeviceUpgrade) SetUpgradePlan(v map[string]interface{})`

SetUpgradePlan sets UpgradePlan field to given value.

### HasUpgradePlan

`func (o *ResponseSiteDeviceUpgrade) HasUpgradePlan() bool`

HasUpgradePlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



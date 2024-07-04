# UpgradeOrgDeviceTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadRequested** | Pointer to **[]string** | list of devices MAC Addresses which cloud has requested to download firmware | [optional] 
**Downloaded** | Pointer to **[]string** | list of devices MAC Addresses which have the firmware downloaded | [optional] 
**Failed** | Pointer to **[]string** | list of devices MAC Addresses which have failed to upgrade | [optional] 
**RebootInProgress** | Pointer to **[]string** | list of devices MAC Addresses which are rebooting | [optional] 
**Rebooted** | Pointer to **[]string** | list of devices MAC Addresses which have rebooted successfully | [optional] 
**Scheduled** | Pointer to **[]string** | list of devices MAC Addresses which cloud has scheduled an upgrade for | [optional] 
**Skipped** | Pointer to **[]string** | list of devices MAC Addresses which skipped upgrade since requested version was same as running version. Use force to always upgrade | [optional] 
**Total** | Pointer to **int32** | count of devices part of this upgrade | [optional] 
**Upgraded** | Pointer to **[]string** | count of devices which have upgraded successfully | [optional] 

## Methods

### NewUpgradeOrgDeviceTargets

`func NewUpgradeOrgDeviceTargets() *UpgradeOrgDeviceTargets`

NewUpgradeOrgDeviceTargets instantiates a new UpgradeOrgDeviceTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeOrgDeviceTargetsWithDefaults

`func NewUpgradeOrgDeviceTargetsWithDefaults() *UpgradeOrgDeviceTargets`

NewUpgradeOrgDeviceTargetsWithDefaults instantiates a new UpgradeOrgDeviceTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadRequested

`func (o *UpgradeOrgDeviceTargets) GetDownloadRequested() []string`

GetDownloadRequested returns the DownloadRequested field if non-nil, zero value otherwise.

### GetDownloadRequestedOk

`func (o *UpgradeOrgDeviceTargets) GetDownloadRequestedOk() (*[]string, bool)`

GetDownloadRequestedOk returns a tuple with the DownloadRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadRequested

`func (o *UpgradeOrgDeviceTargets) SetDownloadRequested(v []string)`

SetDownloadRequested sets DownloadRequested field to given value.

### HasDownloadRequested

`func (o *UpgradeOrgDeviceTargets) HasDownloadRequested() bool`

HasDownloadRequested returns a boolean if a field has been set.

### GetDownloaded

`func (o *UpgradeOrgDeviceTargets) GetDownloaded() []string`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *UpgradeOrgDeviceTargets) GetDownloadedOk() (*[]string, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *UpgradeOrgDeviceTargets) SetDownloaded(v []string)`

SetDownloaded sets Downloaded field to given value.

### HasDownloaded

`func (o *UpgradeOrgDeviceTargets) HasDownloaded() bool`

HasDownloaded returns a boolean if a field has been set.

### GetFailed

`func (o *UpgradeOrgDeviceTargets) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *UpgradeOrgDeviceTargets) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *UpgradeOrgDeviceTargets) SetFailed(v []string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *UpgradeOrgDeviceTargets) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetRebootInProgress

`func (o *UpgradeOrgDeviceTargets) GetRebootInProgress() []string`

GetRebootInProgress returns the RebootInProgress field if non-nil, zero value otherwise.

### GetRebootInProgressOk

`func (o *UpgradeOrgDeviceTargets) GetRebootInProgressOk() (*[]string, bool)`

GetRebootInProgressOk returns a tuple with the RebootInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootInProgress

`func (o *UpgradeOrgDeviceTargets) SetRebootInProgress(v []string)`

SetRebootInProgress sets RebootInProgress field to given value.

### HasRebootInProgress

`func (o *UpgradeOrgDeviceTargets) HasRebootInProgress() bool`

HasRebootInProgress returns a boolean if a field has been set.

### GetRebooted

`func (o *UpgradeOrgDeviceTargets) GetRebooted() []string`

GetRebooted returns the Rebooted field if non-nil, zero value otherwise.

### GetRebootedOk

`func (o *UpgradeOrgDeviceTargets) GetRebootedOk() (*[]string, bool)`

GetRebootedOk returns a tuple with the Rebooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebooted

`func (o *UpgradeOrgDeviceTargets) SetRebooted(v []string)`

SetRebooted sets Rebooted field to given value.

### HasRebooted

`func (o *UpgradeOrgDeviceTargets) HasRebooted() bool`

HasRebooted returns a boolean if a field has been set.

### GetScheduled

`func (o *UpgradeOrgDeviceTargets) GetScheduled() []string`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *UpgradeOrgDeviceTargets) GetScheduledOk() (*[]string, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *UpgradeOrgDeviceTargets) SetScheduled(v []string)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *UpgradeOrgDeviceTargets) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetSkipped

`func (o *UpgradeOrgDeviceTargets) GetSkipped() []string`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *UpgradeOrgDeviceTargets) GetSkippedOk() (*[]string, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *UpgradeOrgDeviceTargets) SetSkipped(v []string)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *UpgradeOrgDeviceTargets) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetTotal

`func (o *UpgradeOrgDeviceTargets) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UpgradeOrgDeviceTargets) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UpgradeOrgDeviceTargets) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UpgradeOrgDeviceTargets) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpgraded

`func (o *UpgradeOrgDeviceTargets) GetUpgraded() []string`

GetUpgraded returns the Upgraded field if non-nil, zero value otherwise.

### GetUpgradedOk

`func (o *UpgradeOrgDeviceTargets) GetUpgradedOk() (*[]string, bool)`

GetUpgradedOk returns a tuple with the Upgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgraded

`func (o *UpgradeOrgDeviceTargets) SetUpgraded(v []string)`

SetUpgraded sets Upgraded field to given value.

### HasUpgraded

`func (o *UpgradeOrgDeviceTargets) HasUpgraded() bool`

HasUpgraded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



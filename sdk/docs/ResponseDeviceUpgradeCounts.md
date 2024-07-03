# ResponseDeviceUpgradeCounts

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

### NewResponseDeviceUpgradeCounts

`func NewResponseDeviceUpgradeCounts() *ResponseDeviceUpgradeCounts`

NewResponseDeviceUpgradeCounts instantiates a new ResponseDeviceUpgradeCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeviceUpgradeCountsWithDefaults

`func NewResponseDeviceUpgradeCountsWithDefaults() *ResponseDeviceUpgradeCounts`

NewResponseDeviceUpgradeCountsWithDefaults instantiates a new ResponseDeviceUpgradeCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadRequested

`func (o *ResponseDeviceUpgradeCounts) GetDownloadRequested() []string`

GetDownloadRequested returns the DownloadRequested field if non-nil, zero value otherwise.

### GetDownloadRequestedOk

`func (o *ResponseDeviceUpgradeCounts) GetDownloadRequestedOk() (*[]string, bool)`

GetDownloadRequestedOk returns a tuple with the DownloadRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadRequested

`func (o *ResponseDeviceUpgradeCounts) SetDownloadRequested(v []string)`

SetDownloadRequested sets DownloadRequested field to given value.

### HasDownloadRequested

`func (o *ResponseDeviceUpgradeCounts) HasDownloadRequested() bool`

HasDownloadRequested returns a boolean if a field has been set.

### GetDownloaded

`func (o *ResponseDeviceUpgradeCounts) GetDownloaded() []string`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *ResponseDeviceUpgradeCounts) GetDownloadedOk() (*[]string, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *ResponseDeviceUpgradeCounts) SetDownloaded(v []string)`

SetDownloaded sets Downloaded field to given value.

### HasDownloaded

`func (o *ResponseDeviceUpgradeCounts) HasDownloaded() bool`

HasDownloaded returns a boolean if a field has been set.

### GetFailed

`func (o *ResponseDeviceUpgradeCounts) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ResponseDeviceUpgradeCounts) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ResponseDeviceUpgradeCounts) SetFailed(v []string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ResponseDeviceUpgradeCounts) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetRebootInProgress

`func (o *ResponseDeviceUpgradeCounts) GetRebootInProgress() []string`

GetRebootInProgress returns the RebootInProgress field if non-nil, zero value otherwise.

### GetRebootInProgressOk

`func (o *ResponseDeviceUpgradeCounts) GetRebootInProgressOk() (*[]string, bool)`

GetRebootInProgressOk returns a tuple with the RebootInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootInProgress

`func (o *ResponseDeviceUpgradeCounts) SetRebootInProgress(v []string)`

SetRebootInProgress sets RebootInProgress field to given value.

### HasRebootInProgress

`func (o *ResponseDeviceUpgradeCounts) HasRebootInProgress() bool`

HasRebootInProgress returns a boolean if a field has been set.

### GetRebooted

`func (o *ResponseDeviceUpgradeCounts) GetRebooted() []string`

GetRebooted returns the Rebooted field if non-nil, zero value otherwise.

### GetRebootedOk

`func (o *ResponseDeviceUpgradeCounts) GetRebootedOk() (*[]string, bool)`

GetRebootedOk returns a tuple with the Rebooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebooted

`func (o *ResponseDeviceUpgradeCounts) SetRebooted(v []string)`

SetRebooted sets Rebooted field to given value.

### HasRebooted

`func (o *ResponseDeviceUpgradeCounts) HasRebooted() bool`

HasRebooted returns a boolean if a field has been set.

### GetScheduled

`func (o *ResponseDeviceUpgradeCounts) GetScheduled() []string`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *ResponseDeviceUpgradeCounts) GetScheduledOk() (*[]string, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *ResponseDeviceUpgradeCounts) SetScheduled(v []string)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *ResponseDeviceUpgradeCounts) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetSkipped

`func (o *ResponseDeviceUpgradeCounts) GetSkipped() []string`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *ResponseDeviceUpgradeCounts) GetSkippedOk() (*[]string, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *ResponseDeviceUpgradeCounts) SetSkipped(v []string)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *ResponseDeviceUpgradeCounts) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetTotal

`func (o *ResponseDeviceUpgradeCounts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseDeviceUpgradeCounts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseDeviceUpgradeCounts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ResponseDeviceUpgradeCounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpgraded

`func (o *ResponseDeviceUpgradeCounts) GetUpgraded() []string`

GetUpgraded returns the Upgraded field if non-nil, zero value otherwise.

### GetUpgradedOk

`func (o *ResponseDeviceUpgradeCounts) GetUpgradedOk() (*[]string, bool)`

GetUpgradedOk returns a tuple with the Upgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgraded

`func (o *ResponseDeviceUpgradeCounts) SetUpgraded(v []string)`

SetUpgraded sets Upgraded field to given value.

### HasUpgraded

`func (o *ResponseDeviceUpgradeCounts) HasUpgraded() bool`

HasUpgraded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



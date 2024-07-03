# ResponseSiteDeviceUpgradeCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadRequested** | Pointer to **int32** | count of devices which cloud has requested to download firmware | [optional] 
**Downloaded** | Pointer to **int32** | count of ap’s which have the firmware downloaded | [optional] 
**Failed** | Pointer to **int32** | count of devices which have failed to upgrade | [optional] 
**RebootInProgress** | Pointer to **int32** | count of devices which are rebooting | [optional] 
**Rebooted** | Pointer to **int32** | count of devices which have rebooted successfully | [optional] 
**Scheduled** | Pointer to **int32** | count of devices which cloud has scheduled an upgrade for | [optional] 
**Skipped** | Pointer to **int32** | count of devices which skipped upgrade since requested version was same as running version. Use force to always upgrade | [optional] 
**Total** | Pointer to **int32** | count of devices part of this upgrade | [optional] 
**Upgraded** | Pointer to **int32** | count of devices which have upgraded successfully | [optional] 

## Methods

### NewResponseSiteDeviceUpgradeCounts

`func NewResponseSiteDeviceUpgradeCounts() *ResponseSiteDeviceUpgradeCounts`

NewResponseSiteDeviceUpgradeCounts instantiates a new ResponseSiteDeviceUpgradeCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSiteDeviceUpgradeCountsWithDefaults

`func NewResponseSiteDeviceUpgradeCountsWithDefaults() *ResponseSiteDeviceUpgradeCounts`

NewResponseSiteDeviceUpgradeCountsWithDefaults instantiates a new ResponseSiteDeviceUpgradeCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadRequested

`func (o *ResponseSiteDeviceUpgradeCounts) GetDownloadRequested() int32`

GetDownloadRequested returns the DownloadRequested field if non-nil, zero value otherwise.

### GetDownloadRequestedOk

`func (o *ResponseSiteDeviceUpgradeCounts) GetDownloadRequestedOk() (*int32, bool)`

GetDownloadRequestedOk returns a tuple with the DownloadRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadRequested

`func (o *ResponseSiteDeviceUpgradeCounts) SetDownloadRequested(v int32)`

SetDownloadRequested sets DownloadRequested field to given value.

### HasDownloadRequested

`func (o *ResponseSiteDeviceUpgradeCounts) HasDownloadRequested() bool`

HasDownloadRequested returns a boolean if a field has been set.

### GetDownloaded

`func (o *ResponseSiteDeviceUpgradeCounts) GetDownloaded() int32`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *ResponseSiteDeviceUpgradeCounts) GetDownloadedOk() (*int32, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *ResponseSiteDeviceUpgradeCounts) SetDownloaded(v int32)`

SetDownloaded sets Downloaded field to given value.

### HasDownloaded

`func (o *ResponseSiteDeviceUpgradeCounts) HasDownloaded() bool`

HasDownloaded returns a boolean if a field has been set.

### GetFailed

`func (o *ResponseSiteDeviceUpgradeCounts) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ResponseSiteDeviceUpgradeCounts) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ResponseSiteDeviceUpgradeCounts) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ResponseSiteDeviceUpgradeCounts) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetRebootInProgress

`func (o *ResponseSiteDeviceUpgradeCounts) GetRebootInProgress() int32`

GetRebootInProgress returns the RebootInProgress field if non-nil, zero value otherwise.

### GetRebootInProgressOk

`func (o *ResponseSiteDeviceUpgradeCounts) GetRebootInProgressOk() (*int32, bool)`

GetRebootInProgressOk returns a tuple with the RebootInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootInProgress

`func (o *ResponseSiteDeviceUpgradeCounts) SetRebootInProgress(v int32)`

SetRebootInProgress sets RebootInProgress field to given value.

### HasRebootInProgress

`func (o *ResponseSiteDeviceUpgradeCounts) HasRebootInProgress() bool`

HasRebootInProgress returns a boolean if a field has been set.

### GetRebooted

`func (o *ResponseSiteDeviceUpgradeCounts) GetRebooted() int32`

GetRebooted returns the Rebooted field if non-nil, zero value otherwise.

### GetRebootedOk

`func (o *ResponseSiteDeviceUpgradeCounts) GetRebootedOk() (*int32, bool)`

GetRebootedOk returns a tuple with the Rebooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebooted

`func (o *ResponseSiteDeviceUpgradeCounts) SetRebooted(v int32)`

SetRebooted sets Rebooted field to given value.

### HasRebooted

`func (o *ResponseSiteDeviceUpgradeCounts) HasRebooted() bool`

HasRebooted returns a boolean if a field has been set.

### GetScheduled

`func (o *ResponseSiteDeviceUpgradeCounts) GetScheduled() int32`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *ResponseSiteDeviceUpgradeCounts) GetScheduledOk() (*int32, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *ResponseSiteDeviceUpgradeCounts) SetScheduled(v int32)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *ResponseSiteDeviceUpgradeCounts) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetSkipped

`func (o *ResponseSiteDeviceUpgradeCounts) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *ResponseSiteDeviceUpgradeCounts) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *ResponseSiteDeviceUpgradeCounts) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *ResponseSiteDeviceUpgradeCounts) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetTotal

`func (o *ResponseSiteDeviceUpgradeCounts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseSiteDeviceUpgradeCounts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseSiteDeviceUpgradeCounts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ResponseSiteDeviceUpgradeCounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpgraded

`func (o *ResponseSiteDeviceUpgradeCounts) GetUpgraded() int32`

GetUpgraded returns the Upgraded field if non-nil, zero value otherwise.

### GetUpgradedOk

`func (o *ResponseSiteDeviceUpgradeCounts) GetUpgradedOk() (*int32, bool)`

GetUpgradedOk returns a tuple with the Upgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgraded

`func (o *ResponseSiteDeviceUpgradeCounts) SetUpgraded(v int32)`

SetUpgraded sets Upgraded field to given value.

### HasUpgraded

`func (o *ResponseSiteDeviceUpgradeCounts) HasUpgraded() bool`

HasUpgraded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



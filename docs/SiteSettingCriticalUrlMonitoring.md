# SiteSettingCriticalUrlMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Monitors** | Pointer to [**[]SiteSettingCriticalUrlMonitoringMonitor**](SiteSettingCriticalUrlMonitoringMonitor.md) |  | [optional] 

## Methods

### NewSiteSettingCriticalUrlMonitoring

`func NewSiteSettingCriticalUrlMonitoring() *SiteSettingCriticalUrlMonitoring`

NewSiteSettingCriticalUrlMonitoring instantiates a new SiteSettingCriticalUrlMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingCriticalUrlMonitoringWithDefaults

`func NewSiteSettingCriticalUrlMonitoringWithDefaults() *SiteSettingCriticalUrlMonitoring`

NewSiteSettingCriticalUrlMonitoringWithDefaults instantiates a new SiteSettingCriticalUrlMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SiteSettingCriticalUrlMonitoring) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SiteSettingCriticalUrlMonitoring) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SiteSettingCriticalUrlMonitoring) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SiteSettingCriticalUrlMonitoring) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMonitors

`func (o *SiteSettingCriticalUrlMonitoring) GetMonitors() []SiteSettingCriticalUrlMonitoringMonitor`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *SiteSettingCriticalUrlMonitoring) GetMonitorsOk() (*[]SiteSettingCriticalUrlMonitoringMonitor, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *SiteSettingCriticalUrlMonitoring) SetMonitors(v []SiteSettingCriticalUrlMonitoringMonitor)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *SiteSettingCriticalUrlMonitoring) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SiteOccupancyAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetsEnabled** | Pointer to **bool** | indicate whether named BLE assets should be included in the zone occupancy calculation | [optional] [default to false]
**ClientsEnabled** | Pointer to **bool** | indicate whether connected WiFi clients should be included in the zone occupancy calculation | [optional] [default to true]
**MinDuration** | Pointer to **int32** | minimum duration | [optional] [default to 3000]
**SdkclientsEnabled** | Pointer to **bool** | indicate whether SDK clients should be included in the zone occupancy calculation | [optional] [default to false]
**UnconnectedClientsEnabled** | Pointer to **bool** | indicate whether unconnected WiFi clients should be included in the zone occupancy calculation | [optional] [default to false]

## Methods

### NewSiteOccupancyAnalytics

`func NewSiteOccupancyAnalytics() *SiteOccupancyAnalytics`

NewSiteOccupancyAnalytics instantiates a new SiteOccupancyAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteOccupancyAnalyticsWithDefaults

`func NewSiteOccupancyAnalyticsWithDefaults() *SiteOccupancyAnalytics`

NewSiteOccupancyAnalyticsWithDefaults instantiates a new SiteOccupancyAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetsEnabled

`func (o *SiteOccupancyAnalytics) GetAssetsEnabled() bool`

GetAssetsEnabled returns the AssetsEnabled field if non-nil, zero value otherwise.

### GetAssetsEnabledOk

`func (o *SiteOccupancyAnalytics) GetAssetsEnabledOk() (*bool, bool)`

GetAssetsEnabledOk returns a tuple with the AssetsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsEnabled

`func (o *SiteOccupancyAnalytics) SetAssetsEnabled(v bool)`

SetAssetsEnabled sets AssetsEnabled field to given value.

### HasAssetsEnabled

`func (o *SiteOccupancyAnalytics) HasAssetsEnabled() bool`

HasAssetsEnabled returns a boolean if a field has been set.

### GetClientsEnabled

`func (o *SiteOccupancyAnalytics) GetClientsEnabled() bool`

GetClientsEnabled returns the ClientsEnabled field if non-nil, zero value otherwise.

### GetClientsEnabledOk

`func (o *SiteOccupancyAnalytics) GetClientsEnabledOk() (*bool, bool)`

GetClientsEnabledOk returns a tuple with the ClientsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsEnabled

`func (o *SiteOccupancyAnalytics) SetClientsEnabled(v bool)`

SetClientsEnabled sets ClientsEnabled field to given value.

### HasClientsEnabled

`func (o *SiteOccupancyAnalytics) HasClientsEnabled() bool`

HasClientsEnabled returns a boolean if a field has been set.

### GetMinDuration

`func (o *SiteOccupancyAnalytics) GetMinDuration() int32`

GetMinDuration returns the MinDuration field if non-nil, zero value otherwise.

### GetMinDurationOk

`func (o *SiteOccupancyAnalytics) GetMinDurationOk() (*int32, bool)`

GetMinDurationOk returns a tuple with the MinDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDuration

`func (o *SiteOccupancyAnalytics) SetMinDuration(v int32)`

SetMinDuration sets MinDuration field to given value.

### HasMinDuration

`func (o *SiteOccupancyAnalytics) HasMinDuration() bool`

HasMinDuration returns a boolean if a field has been set.

### GetSdkclientsEnabled

`func (o *SiteOccupancyAnalytics) GetSdkclientsEnabled() bool`

GetSdkclientsEnabled returns the SdkclientsEnabled field if non-nil, zero value otherwise.

### GetSdkclientsEnabledOk

`func (o *SiteOccupancyAnalytics) GetSdkclientsEnabledOk() (*bool, bool)`

GetSdkclientsEnabledOk returns a tuple with the SdkclientsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkclientsEnabled

`func (o *SiteOccupancyAnalytics) SetSdkclientsEnabled(v bool)`

SetSdkclientsEnabled sets SdkclientsEnabled field to given value.

### HasSdkclientsEnabled

`func (o *SiteOccupancyAnalytics) HasSdkclientsEnabled() bool`

HasSdkclientsEnabled returns a boolean if a field has been set.

### GetUnconnectedClientsEnabled

`func (o *SiteOccupancyAnalytics) GetUnconnectedClientsEnabled() bool`

GetUnconnectedClientsEnabled returns the UnconnectedClientsEnabled field if non-nil, zero value otherwise.

### GetUnconnectedClientsEnabledOk

`func (o *SiteOccupancyAnalytics) GetUnconnectedClientsEnabledOk() (*bool, bool)`

GetUnconnectedClientsEnabledOk returns a tuple with the UnconnectedClientsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconnectedClientsEnabled

`func (o *SiteOccupancyAnalytics) SetUnconnectedClientsEnabled(v bool)`

SetUnconnectedClientsEnabled sets UnconnectedClientsEnabled field to given value.

### HasUnconnectedClientsEnabled

`func (o *SiteOccupancyAnalytics) HasUnconnectedClientsEnabled() bool`

HasUnconnectedClientsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SiteZoneOccupancyAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailNotifiers** | Pointer to **[]string** | list of email addresses to send email notifications when the alert threshold is reached | [optional] 
**Enabled** | Pointer to **bool** | indicate whether zone occupancy alert is enabled for the site | [optional] [default to false]
**Threshold** | Pointer to **int32** | sending zone-occupancy-alert webhook message only if a zone stays non-compliant (i.e. actual occupancy &gt; occupancy_limit) for a minimum duration specified in the threshold, in minutes | [optional] [default to 5]

## Methods

### NewSiteZoneOccupancyAlert

`func NewSiteZoneOccupancyAlert() *SiteZoneOccupancyAlert`

NewSiteZoneOccupancyAlert instantiates a new SiteZoneOccupancyAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteZoneOccupancyAlertWithDefaults

`func NewSiteZoneOccupancyAlertWithDefaults() *SiteZoneOccupancyAlert`

NewSiteZoneOccupancyAlertWithDefaults instantiates a new SiteZoneOccupancyAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailNotifiers

`func (o *SiteZoneOccupancyAlert) GetEmailNotifiers() []string`

GetEmailNotifiers returns the EmailNotifiers field if non-nil, zero value otherwise.

### GetEmailNotifiersOk

`func (o *SiteZoneOccupancyAlert) GetEmailNotifiersOk() (*[]string, bool)`

GetEmailNotifiersOk returns a tuple with the EmailNotifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifiers

`func (o *SiteZoneOccupancyAlert) SetEmailNotifiers(v []string)`

SetEmailNotifiers sets EmailNotifiers field to given value.

### HasEmailNotifiers

`func (o *SiteZoneOccupancyAlert) HasEmailNotifiers() bool`

HasEmailNotifiers returns a boolean if a field has been set.

### GetEnabled

`func (o *SiteZoneOccupancyAlert) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SiteZoneOccupancyAlert) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SiteZoneOccupancyAlert) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SiteZoneOccupancyAlert) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetThreshold

`func (o *SiteZoneOccupancyAlert) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SiteZoneOccupancyAlert) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SiteZoneOccupancyAlert) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *SiteZoneOccupancyAlert) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



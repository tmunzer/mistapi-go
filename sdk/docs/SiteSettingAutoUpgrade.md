# SiteSettingAutoUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomVersions** | Pointer to **map[string]string** | custom versions for different models. Property key is the model name (e.g. \&quot;AP41\&quot;) | [optional] 
**DayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) |  | [optional] 
**Enabled** | Pointer to **bool** | whether auto upgrade should happen (Note that Mist may auto-upgrade if the version is not supported) | [optional] [default to false]
**TimeOfDay** | Pointer to **string** | any / HH:MM (24-hour format), upgrade will happen within up to 1-hour from this time | [optional] 
**Version** | Pointer to [**SiteAutoUpgradeVersion**](SiteAutoUpgradeVersion.md) |  | [optional] [default to SITEAUTOUPGRADEVERSION_STABLE]

## Methods

### NewSiteSettingAutoUpgrade

`func NewSiteSettingAutoUpgrade() *SiteSettingAutoUpgrade`

NewSiteSettingAutoUpgrade instantiates a new SiteSettingAutoUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingAutoUpgradeWithDefaults

`func NewSiteSettingAutoUpgradeWithDefaults() *SiteSettingAutoUpgrade`

NewSiteSettingAutoUpgradeWithDefaults instantiates a new SiteSettingAutoUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomVersions

`func (o *SiteSettingAutoUpgrade) GetCustomVersions() map[string]string`

GetCustomVersions returns the CustomVersions field if non-nil, zero value otherwise.

### GetCustomVersionsOk

`func (o *SiteSettingAutoUpgrade) GetCustomVersionsOk() (*map[string]string, bool)`

GetCustomVersionsOk returns a tuple with the CustomVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVersions

`func (o *SiteSettingAutoUpgrade) SetCustomVersions(v map[string]string)`

SetCustomVersions sets CustomVersions field to given value.

### HasCustomVersions

`func (o *SiteSettingAutoUpgrade) HasCustomVersions() bool`

HasCustomVersions returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *SiteSettingAutoUpgrade) GetDayOfWeek() DayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SiteSettingAutoUpgrade) GetDayOfWeekOk() (*DayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SiteSettingAutoUpgrade) SetDayOfWeek(v DayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SiteSettingAutoUpgrade) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetEnabled

`func (o *SiteSettingAutoUpgrade) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SiteSettingAutoUpgrade) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SiteSettingAutoUpgrade) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SiteSettingAutoUpgrade) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *SiteSettingAutoUpgrade) GetTimeOfDay() string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *SiteSettingAutoUpgrade) GetTimeOfDayOk() (*string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *SiteSettingAutoUpgrade) SetTimeOfDay(v string)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *SiteSettingAutoUpgrade) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetVersion

`func (o *SiteSettingAutoUpgrade) GetVersion() SiteAutoUpgradeVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SiteSettingAutoUpgrade) GetVersionOk() (*SiteAutoUpgradeVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SiteSettingAutoUpgrade) SetVersion(v SiteAutoUpgradeVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SiteSettingAutoUpgrade) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



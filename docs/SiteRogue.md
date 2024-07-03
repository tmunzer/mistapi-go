# SiteRogue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | whether or not rogue detection is enabled | [optional] [default to false]
**HoneypotEnabled** | Pointer to **bool** | whether or not honeypot detection is enabled | [optional] [default to false]
**MinDuration** | Pointer to **int32** | minimum duration for a bssid to be considered rogue | [optional] [default to 10]
**MinRssi** | Pointer to **int32** | minimum RSSI for an AP to be considered rogue (ignoring APs that’s far away) | [optional] [default to -80]
**WhitelistedBssids** | Pointer to **[]string** | list of BSSIDs to whitelist. Ex: \&quot;cc-:8e-:6f-:d4-:bf-:16\&quot;, \&quot;cc-8e-6f-d4-bf-16\&quot;, \&quot;cc-73-*\&quot;, \&quot;cc:82:*\&quot; | [optional] 
**WhitelistedSsids** | Pointer to **[]string** | list of SSIDs to whitelist | [optional] 

## Methods

### NewSiteRogue

`func NewSiteRogue() *SiteRogue`

NewSiteRogue instantiates a new SiteRogue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteRogueWithDefaults

`func NewSiteRogueWithDefaults() *SiteRogue`

NewSiteRogueWithDefaults instantiates a new SiteRogue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SiteRogue) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SiteRogue) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SiteRogue) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SiteRogue) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHoneypotEnabled

`func (o *SiteRogue) GetHoneypotEnabled() bool`

GetHoneypotEnabled returns the HoneypotEnabled field if non-nil, zero value otherwise.

### GetHoneypotEnabledOk

`func (o *SiteRogue) GetHoneypotEnabledOk() (*bool, bool)`

GetHoneypotEnabledOk returns a tuple with the HoneypotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoneypotEnabled

`func (o *SiteRogue) SetHoneypotEnabled(v bool)`

SetHoneypotEnabled sets HoneypotEnabled field to given value.

### HasHoneypotEnabled

`func (o *SiteRogue) HasHoneypotEnabled() bool`

HasHoneypotEnabled returns a boolean if a field has been set.

### GetMinDuration

`func (o *SiteRogue) GetMinDuration() int32`

GetMinDuration returns the MinDuration field if non-nil, zero value otherwise.

### GetMinDurationOk

`func (o *SiteRogue) GetMinDurationOk() (*int32, bool)`

GetMinDurationOk returns a tuple with the MinDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDuration

`func (o *SiteRogue) SetMinDuration(v int32)`

SetMinDuration sets MinDuration field to given value.

### HasMinDuration

`func (o *SiteRogue) HasMinDuration() bool`

HasMinDuration returns a boolean if a field has been set.

### GetMinRssi

`func (o *SiteRogue) GetMinRssi() int32`

GetMinRssi returns the MinRssi field if non-nil, zero value otherwise.

### GetMinRssiOk

`func (o *SiteRogue) GetMinRssiOk() (*int32, bool)`

GetMinRssiOk returns a tuple with the MinRssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRssi

`func (o *SiteRogue) SetMinRssi(v int32)`

SetMinRssi sets MinRssi field to given value.

### HasMinRssi

`func (o *SiteRogue) HasMinRssi() bool`

HasMinRssi returns a boolean if a field has been set.

### GetWhitelistedBssids

`func (o *SiteRogue) GetWhitelistedBssids() []string`

GetWhitelistedBssids returns the WhitelistedBssids field if non-nil, zero value otherwise.

### GetWhitelistedBssidsOk

`func (o *SiteRogue) GetWhitelistedBssidsOk() (*[]string, bool)`

GetWhitelistedBssidsOk returns a tuple with the WhitelistedBssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedBssids

`func (o *SiteRogue) SetWhitelistedBssids(v []string)`

SetWhitelistedBssids sets WhitelistedBssids field to given value.

### HasWhitelistedBssids

`func (o *SiteRogue) HasWhitelistedBssids() bool`

HasWhitelistedBssids returns a boolean if a field has been set.

### GetWhitelistedSsids

`func (o *SiteRogue) GetWhitelistedSsids() []string`

GetWhitelistedSsids returns the WhitelistedSsids field if non-nil, zero value otherwise.

### GetWhitelistedSsidsOk

`func (o *SiteRogue) GetWhitelistedSsidsOk() (*[]string, bool)`

GetWhitelistedSsidsOk returns a tuple with the WhitelistedSsids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedSsids

`func (o *SiteRogue) SetWhitelistedSsids(v []string)`

SetWhitelistedSsids sets WhitelistedSsids field to given value.

### HasWhitelistedSsids

`func (o *SiteRogue) HasWhitelistedSsids() bool`

HasWhitelistedSsids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



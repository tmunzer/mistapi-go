# SuppressedAlarmApplies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** | UUID of the current org (if provided, the alarms will be suppressed at org level) | [optional] 
**SiteIds** | Pointer to **[]string** | List of UUID of the sites within the org (if provided, the alarms will be suppressed for all the mentioned sites under the org) | [optional] 
**SitegroupIds** | Pointer to **[]string** | List of UUID of the site groups within the org (if provided, the alarms will be suppressed for all the sites under those site groups) | [optional] 

## Methods

### NewSuppressedAlarmApplies

`func NewSuppressedAlarmApplies() *SuppressedAlarmApplies`

NewSuppressedAlarmApplies instantiates a new SuppressedAlarmApplies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressedAlarmAppliesWithDefaults

`func NewSuppressedAlarmAppliesWithDefaults() *SuppressedAlarmApplies`

NewSuppressedAlarmAppliesWithDefaults instantiates a new SuppressedAlarmApplies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *SuppressedAlarmApplies) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SuppressedAlarmApplies) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SuppressedAlarmApplies) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SuppressedAlarmApplies) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSiteIds

`func (o *SuppressedAlarmApplies) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *SuppressedAlarmApplies) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *SuppressedAlarmApplies) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *SuppressedAlarmApplies) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### GetSitegroupIds

`func (o *SuppressedAlarmApplies) GetSitegroupIds() []string`

GetSitegroupIds returns the SitegroupIds field if non-nil, zero value otherwise.

### GetSitegroupIdsOk

`func (o *SuppressedAlarmApplies) GetSitegroupIdsOk() (*[]string, bool)`

GetSitegroupIdsOk returns a tuple with the SitegroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitegroupIds

`func (o *SuppressedAlarmApplies) SetSitegroupIds(v []string)`

SetSitegroupIds sets SitegroupIds field to given value.

### HasSitegroupIds

`func (o *SuppressedAlarmApplies) HasSitegroupIds() bool`

HasSitegroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UiSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DefaultScopeId** | Pointer to **string** |  | [optional] 
**DefaultScopeType** | Pointer to **string** |  | [optional] 
**DefaultTimeRange** | Pointer to [**UiSettingsDefaultTimeRange**](UiSettingsDefaultTimeRange.md) |  | [optional] 
**Description** | **string** |  | 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IsCustomDataboard** | Pointer to **bool** |  | [optional] 
**IsScopeLinked** | Pointer to **bool** |  | [optional] 
**IsTimeRangeLinked** | Pointer to **bool** |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Purpose** | **string** |  | 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Tiles** | Pointer to [**[]UiSettingsTile**](UiSettingsTile.md) |  | [optional] 

## Methods

### NewUiSettings

`func NewUiSettings(description string, purpose string, ) *UiSettings`

NewUiSettings instantiates a new UiSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiSettingsWithDefaults

`func NewUiSettingsWithDefaults() *UiSettings`

NewUiSettingsWithDefaults instantiates a new UiSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *UiSettings) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *UiSettings) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *UiSettings) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *UiSettings) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDefaultScopeId

`func (o *UiSettings) GetDefaultScopeId() string`

GetDefaultScopeId returns the DefaultScopeId field if non-nil, zero value otherwise.

### GetDefaultScopeIdOk

`func (o *UiSettings) GetDefaultScopeIdOk() (*string, bool)`

GetDefaultScopeIdOk returns a tuple with the DefaultScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultScopeId

`func (o *UiSettings) SetDefaultScopeId(v string)`

SetDefaultScopeId sets DefaultScopeId field to given value.

### HasDefaultScopeId

`func (o *UiSettings) HasDefaultScopeId() bool`

HasDefaultScopeId returns a boolean if a field has been set.

### GetDefaultScopeType

`func (o *UiSettings) GetDefaultScopeType() string`

GetDefaultScopeType returns the DefaultScopeType field if non-nil, zero value otherwise.

### GetDefaultScopeTypeOk

`func (o *UiSettings) GetDefaultScopeTypeOk() (*string, bool)`

GetDefaultScopeTypeOk returns a tuple with the DefaultScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultScopeType

`func (o *UiSettings) SetDefaultScopeType(v string)`

SetDefaultScopeType sets DefaultScopeType field to given value.

### HasDefaultScopeType

`func (o *UiSettings) HasDefaultScopeType() bool`

HasDefaultScopeType returns a boolean if a field has been set.

### GetDefaultTimeRange

`func (o *UiSettings) GetDefaultTimeRange() UiSettingsDefaultTimeRange`

GetDefaultTimeRange returns the DefaultTimeRange field if non-nil, zero value otherwise.

### GetDefaultTimeRangeOk

`func (o *UiSettings) GetDefaultTimeRangeOk() (*UiSettingsDefaultTimeRange, bool)`

GetDefaultTimeRangeOk returns a tuple with the DefaultTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeRange

`func (o *UiSettings) SetDefaultTimeRange(v UiSettingsDefaultTimeRange)`

SetDefaultTimeRange sets DefaultTimeRange field to given value.

### HasDefaultTimeRange

`func (o *UiSettings) HasDefaultTimeRange() bool`

HasDefaultTimeRange returns a boolean if a field has been set.

### GetDescription

`func (o *UiSettings) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UiSettings) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UiSettings) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetForSite

`func (o *UiSettings) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *UiSettings) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *UiSettings) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *UiSettings) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *UiSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UiSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UiSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UiSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCustomDataboard

`func (o *UiSettings) GetIsCustomDataboard() bool`

GetIsCustomDataboard returns the IsCustomDataboard field if non-nil, zero value otherwise.

### GetIsCustomDataboardOk

`func (o *UiSettings) GetIsCustomDataboardOk() (*bool, bool)`

GetIsCustomDataboardOk returns a tuple with the IsCustomDataboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomDataboard

`func (o *UiSettings) SetIsCustomDataboard(v bool)`

SetIsCustomDataboard sets IsCustomDataboard field to given value.

### HasIsCustomDataboard

`func (o *UiSettings) HasIsCustomDataboard() bool`

HasIsCustomDataboard returns a boolean if a field has been set.

### GetIsScopeLinked

`func (o *UiSettings) GetIsScopeLinked() bool`

GetIsScopeLinked returns the IsScopeLinked field if non-nil, zero value otherwise.

### GetIsScopeLinkedOk

`func (o *UiSettings) GetIsScopeLinkedOk() (*bool, bool)`

GetIsScopeLinkedOk returns a tuple with the IsScopeLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScopeLinked

`func (o *UiSettings) SetIsScopeLinked(v bool)`

SetIsScopeLinked sets IsScopeLinked field to given value.

### HasIsScopeLinked

`func (o *UiSettings) HasIsScopeLinked() bool`

HasIsScopeLinked returns a boolean if a field has been set.

### GetIsTimeRangeLinked

`func (o *UiSettings) GetIsTimeRangeLinked() bool`

GetIsTimeRangeLinked returns the IsTimeRangeLinked field if non-nil, zero value otherwise.

### GetIsTimeRangeLinkedOk

`func (o *UiSettings) GetIsTimeRangeLinkedOk() (*bool, bool)`

GetIsTimeRangeLinkedOk returns a tuple with the IsTimeRangeLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeRangeLinked

`func (o *UiSettings) SetIsTimeRangeLinked(v bool)`

SetIsTimeRangeLinked sets IsTimeRangeLinked field to given value.

### HasIsTimeRangeLinked

`func (o *UiSettings) HasIsTimeRangeLinked() bool`

HasIsTimeRangeLinked returns a boolean if a field has been set.

### GetModifiedTime

`func (o *UiSettings) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *UiSettings) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *UiSettings) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *UiSettings) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *UiSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UiSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UiSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UiSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *UiSettings) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *UiSettings) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *UiSettings) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *UiSettings) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPurpose

`func (o *UiSettings) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *UiSettings) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *UiSettings) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetSiteId

`func (o *UiSettings) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UiSettings) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UiSettings) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UiSettings) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTiles

`func (o *UiSettings) GetTiles() []UiSettingsTile`

GetTiles returns the Tiles field if non-nil, zero value otherwise.

### GetTilesOk

`func (o *UiSettings) GetTilesOk() (*[]UiSettingsTile, bool)`

GetTilesOk returns a tuple with the Tiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiles

`func (o *UiSettings) SetTiles(v []UiSettingsTile)`

SetTiles sets Tiles field to given value.

### HasTiles

`func (o *UiSettings) HasTiles() bool`

HasTiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



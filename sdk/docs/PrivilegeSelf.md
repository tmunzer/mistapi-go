# PrivilegeSelf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MspId** | Pointer to **string** |  | [optional] 
**MspLogoUrl** | Pointer to **string** | logo of the MSP (if the MSP belongs to an Advanced tier) | [optional] [readonly] 
**MspName** | Pointer to **NullableString** | name of the MSP (if the org belongs to an MSP) | [optional] [readonly] 
**MspUrl** | Pointer to **string** | custom url of the MSP (if the MSP belongs to an Advanced tier) | [optional] [readonly] 
**Name** | Pointer to **string** | name of the org/site/MSP depending on object scope | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**OrgName** | Pointer to **string** | name of the org (for a site belonging to org) | [optional] 
**OrggroupIds** | Pointer to **[]string** | if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;orggroup&#x60; | [optional] 
**Role** | [**PrivilegeSelfRole**](PrivilegeSelfRole.md) |  | 
**Scope** | [**PrivilegeSelfScope**](PrivilegeSelfScope.md) |  | 
**SiteId** | Pointer to **string** |  | [optional] 
**SitegroupIds** | Pointer to **[]string** |  | [optional] 
**Views** | Pointer to [**PrivilegeSelfViews**](PrivilegeSelfViews.md) |  | [optional] 

## Methods

### NewPrivilegeSelf

`func NewPrivilegeSelf(role PrivilegeSelfRole, scope PrivilegeSelfScope, ) *PrivilegeSelf`

NewPrivilegeSelf instantiates a new PrivilegeSelf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeSelfWithDefaults

`func NewPrivilegeSelfWithDefaults() *PrivilegeSelf`

NewPrivilegeSelfWithDefaults instantiates a new PrivilegeSelf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMspId

`func (o *PrivilegeSelf) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *PrivilegeSelf) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *PrivilegeSelf) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *PrivilegeSelf) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetMspLogoUrl

`func (o *PrivilegeSelf) GetMspLogoUrl() string`

GetMspLogoUrl returns the MspLogoUrl field if non-nil, zero value otherwise.

### GetMspLogoUrlOk

`func (o *PrivilegeSelf) GetMspLogoUrlOk() (*string, bool)`

GetMspLogoUrlOk returns a tuple with the MspLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspLogoUrl

`func (o *PrivilegeSelf) SetMspLogoUrl(v string)`

SetMspLogoUrl sets MspLogoUrl field to given value.

### HasMspLogoUrl

`func (o *PrivilegeSelf) HasMspLogoUrl() bool`

HasMspLogoUrl returns a boolean if a field has been set.

### GetMspName

`func (o *PrivilegeSelf) GetMspName() string`

GetMspName returns the MspName field if non-nil, zero value otherwise.

### GetMspNameOk

`func (o *PrivilegeSelf) GetMspNameOk() (*string, bool)`

GetMspNameOk returns a tuple with the MspName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspName

`func (o *PrivilegeSelf) SetMspName(v string)`

SetMspName sets MspName field to given value.

### HasMspName

`func (o *PrivilegeSelf) HasMspName() bool`

HasMspName returns a boolean if a field has been set.

### SetMspNameNil

`func (o *PrivilegeSelf) SetMspNameNil(b bool)`

 SetMspNameNil sets the value for MspName to be an explicit nil

### UnsetMspName
`func (o *PrivilegeSelf) UnsetMspName()`

UnsetMspName ensures that no value is present for MspName, not even an explicit nil
### GetMspUrl

`func (o *PrivilegeSelf) GetMspUrl() string`

GetMspUrl returns the MspUrl field if non-nil, zero value otherwise.

### GetMspUrlOk

`func (o *PrivilegeSelf) GetMspUrlOk() (*string, bool)`

GetMspUrlOk returns a tuple with the MspUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspUrl

`func (o *PrivilegeSelf) SetMspUrl(v string)`

SetMspUrl sets MspUrl field to given value.

### HasMspUrl

`func (o *PrivilegeSelf) HasMspUrl() bool`

HasMspUrl returns a boolean if a field has been set.

### GetName

`func (o *PrivilegeSelf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivilegeSelf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivilegeSelf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivilegeSelf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *PrivilegeSelf) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PrivilegeSelf) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PrivilegeSelf) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PrivilegeSelf) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *PrivilegeSelf) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *PrivilegeSelf) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *PrivilegeSelf) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *PrivilegeSelf) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrggroupIds

`func (o *PrivilegeSelf) GetOrggroupIds() []string`

GetOrggroupIds returns the OrggroupIds field if non-nil, zero value otherwise.

### GetOrggroupIdsOk

`func (o *PrivilegeSelf) GetOrggroupIdsOk() (*[]string, bool)`

GetOrggroupIdsOk returns a tuple with the OrggroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrggroupIds

`func (o *PrivilegeSelf) SetOrggroupIds(v []string)`

SetOrggroupIds sets OrggroupIds field to given value.

### HasOrggroupIds

`func (o *PrivilegeSelf) HasOrggroupIds() bool`

HasOrggroupIds returns a boolean if a field has been set.

### GetRole

`func (o *PrivilegeSelf) GetRole() PrivilegeSelfRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PrivilegeSelf) GetRoleOk() (*PrivilegeSelfRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PrivilegeSelf) SetRole(v PrivilegeSelfRole)`

SetRole sets Role field to given value.


### GetScope

`func (o *PrivilegeSelf) GetScope() PrivilegeSelfScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PrivilegeSelf) GetScopeOk() (*PrivilegeSelfScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PrivilegeSelf) SetScope(v PrivilegeSelfScope)`

SetScope sets Scope field to given value.


### GetSiteId

`func (o *PrivilegeSelf) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PrivilegeSelf) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PrivilegeSelf) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PrivilegeSelf) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSitegroupIds

`func (o *PrivilegeSelf) GetSitegroupIds() []string`

GetSitegroupIds returns the SitegroupIds field if non-nil, zero value otherwise.

### GetSitegroupIdsOk

`func (o *PrivilegeSelf) GetSitegroupIdsOk() (*[]string, bool)`

GetSitegroupIdsOk returns a tuple with the SitegroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitegroupIds

`func (o *PrivilegeSelf) SetSitegroupIds(v []string)`

SetSitegroupIds sets SitegroupIds field to given value.

### HasSitegroupIds

`func (o *PrivilegeSelf) HasSitegroupIds() bool`

HasSitegroupIds returns a boolean if a field has been set.

### GetViews

`func (o *PrivilegeSelf) GetViews() PrivilegeSelfViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *PrivilegeSelf) GetViewsOk() (*PrivilegeSelfViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *PrivilegeSelf) SetViews(v PrivilegeSelfViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *PrivilegeSelf) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



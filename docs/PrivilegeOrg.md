# PrivilegeOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**PrivilegeOrgRole**](PrivilegeOrgRole.md) |  | 
**Scope** | [**PrivilegeOrgScope**](PrivilegeOrgScope.md) |  | 
**SiteId** | Pointer to **string** | if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; | [optional] 
**SitegroupId** | Pointer to **string** | if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;sitegroup&#x60; | [optional] 
**Views** | Pointer to [**PrivilegeOrgViews**](PrivilegeOrgViews.md) |  | [optional] 

## Methods

### NewPrivilegeOrg

`func NewPrivilegeOrg(role PrivilegeOrgRole, scope PrivilegeOrgScope, ) *PrivilegeOrg`

NewPrivilegeOrg instantiates a new PrivilegeOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeOrgWithDefaults

`func NewPrivilegeOrgWithDefaults() *PrivilegeOrg`

NewPrivilegeOrgWithDefaults instantiates a new PrivilegeOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *PrivilegeOrg) GetRole() PrivilegeOrgRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PrivilegeOrg) GetRoleOk() (*PrivilegeOrgRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PrivilegeOrg) SetRole(v PrivilegeOrgRole)`

SetRole sets Role field to given value.


### GetScope

`func (o *PrivilegeOrg) GetScope() PrivilegeOrgScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PrivilegeOrg) GetScopeOk() (*PrivilegeOrgScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PrivilegeOrg) SetScope(v PrivilegeOrgScope)`

SetScope sets Scope field to given value.


### GetSiteId

`func (o *PrivilegeOrg) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PrivilegeOrg) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PrivilegeOrg) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PrivilegeOrg) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSitegroupId

`func (o *PrivilegeOrg) GetSitegroupId() string`

GetSitegroupId returns the SitegroupId field if non-nil, zero value otherwise.

### GetSitegroupIdOk

`func (o *PrivilegeOrg) GetSitegroupIdOk() (*string, bool)`

GetSitegroupIdOk returns a tuple with the SitegroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitegroupId

`func (o *PrivilegeOrg) SetSitegroupId(v string)`

SetSitegroupId sets SitegroupId field to given value.

### HasSitegroupId

`func (o *PrivilegeOrg) HasSitegroupId() bool`

HasSitegroupId returns a boolean if a field has been set.

### GetViews

`func (o *PrivilegeOrg) GetViews() PrivilegeOrgViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *PrivilegeOrg) GetViewsOk() (*PrivilegeOrgViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *PrivilegeOrg) SetViews(v PrivilegeOrgViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *PrivilegeOrg) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



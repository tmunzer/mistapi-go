# PrivilegeMsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** | if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;org&#x60; | [optional] 
**OrgName** | Pointer to **string** | name of the org (for a site belonging to org) | [optional] 
**OrggroupId** | Pointer to **string** | if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;orggroup&#x60; | [optional] 
**Role** | [**PrivilegeMspRole**](PrivilegeMspRole.md) |  | 
**Scope** | [**PrivilegeMspScope**](PrivilegeMspScope.md) |  | 
**Views** | Pointer to [**PrivilegeMspView**](PrivilegeMspView.md) |  | [optional] 

## Methods

### NewPrivilegeMsp

`func NewPrivilegeMsp(role PrivilegeMspRole, scope PrivilegeMspScope, ) *PrivilegeMsp`

NewPrivilegeMsp instantiates a new PrivilegeMsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeMspWithDefaults

`func NewPrivilegeMspWithDefaults() *PrivilegeMsp`

NewPrivilegeMspWithDefaults instantiates a new PrivilegeMsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *PrivilegeMsp) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PrivilegeMsp) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PrivilegeMsp) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PrivilegeMsp) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *PrivilegeMsp) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *PrivilegeMsp) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *PrivilegeMsp) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *PrivilegeMsp) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrggroupId

`func (o *PrivilegeMsp) GetOrggroupId() string`

GetOrggroupId returns the OrggroupId field if non-nil, zero value otherwise.

### GetOrggroupIdOk

`func (o *PrivilegeMsp) GetOrggroupIdOk() (*string, bool)`

GetOrggroupIdOk returns a tuple with the OrggroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrggroupId

`func (o *PrivilegeMsp) SetOrggroupId(v string)`

SetOrggroupId sets OrggroupId field to given value.

### HasOrggroupId

`func (o *PrivilegeMsp) HasOrggroupId() bool`

HasOrggroupId returns a boolean if a field has been set.

### GetRole

`func (o *PrivilegeMsp) GetRole() PrivilegeMspRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PrivilegeMsp) GetRoleOk() (*PrivilegeMspRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PrivilegeMsp) SetRole(v PrivilegeMspRole)`

SetRole sets Role field to given value.


### GetScope

`func (o *PrivilegeMsp) GetScope() PrivilegeMspScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PrivilegeMsp) GetScopeOk() (*PrivilegeMspScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PrivilegeMsp) SetScope(v PrivilegeMspScope)`

SetScope sets Scope field to given value.


### GetViews

`func (o *PrivilegeMsp) GetViews() PrivilegeMspView`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *PrivilegeMsp) GetViewsOk() (*PrivilegeMspView, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *PrivilegeMsp) SetViews(v PrivilegeMspView)`

SetViews sets Views field to given value.

### HasViews

`func (o *PrivilegeMsp) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



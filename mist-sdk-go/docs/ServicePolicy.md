# ServicePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AllowDeny**](AllowDeny.md) |  | [optional] [default to ALLOWDENY_ALLOW]
**Appqoe** | Pointer to [**ServicePolicyAppqoe**](ServicePolicyAppqoe.md) |  | [optional] 
**Ewf** | Pointer to [**[]ServicePolicyEwfRule**](ServicePolicyEwfRule.md) |  | [optional] 
**Idp** | Pointer to [**IdpConfig**](IdpConfig.md) |  | [optional] 
**LocalRouting** | Pointer to **bool** | access within the same VRF | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 
**PathPreferences** | Pointer to **string** | by default, we derive all paths available and use them optionally, you can customize by using &#x60;path_preference&#x60; | [optional] 
**ServicepolicyId** | Pointer to **string** | used to link servicepolicy defined at org level and overwrite some attributes | [optional] 
**Services** | Pointer to **[]string** |  | [optional] 
**Tenants** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServicePolicy

`func NewServicePolicy() *ServicePolicy`

NewServicePolicy instantiates a new ServicePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePolicyWithDefaults

`func NewServicePolicyWithDefaults() *ServicePolicy`

NewServicePolicyWithDefaults instantiates a new ServicePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ServicePolicy) GetAction() AllowDeny`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ServicePolicy) GetActionOk() (*AllowDeny, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ServicePolicy) SetAction(v AllowDeny)`

SetAction sets Action field to given value.

### HasAction

`func (o *ServicePolicy) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAppqoe

`func (o *ServicePolicy) GetAppqoe() ServicePolicyAppqoe`

GetAppqoe returns the Appqoe field if non-nil, zero value otherwise.

### GetAppqoeOk

`func (o *ServicePolicy) GetAppqoeOk() (*ServicePolicyAppqoe, bool)`

GetAppqoeOk returns a tuple with the Appqoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppqoe

`func (o *ServicePolicy) SetAppqoe(v ServicePolicyAppqoe)`

SetAppqoe sets Appqoe field to given value.

### HasAppqoe

`func (o *ServicePolicy) HasAppqoe() bool`

HasAppqoe returns a boolean if a field has been set.

### GetEwf

`func (o *ServicePolicy) GetEwf() []ServicePolicyEwfRule`

GetEwf returns the Ewf field if non-nil, zero value otherwise.

### GetEwfOk

`func (o *ServicePolicy) GetEwfOk() (*[]ServicePolicyEwfRule, bool)`

GetEwfOk returns a tuple with the Ewf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwf

`func (o *ServicePolicy) SetEwf(v []ServicePolicyEwfRule)`

SetEwf sets Ewf field to given value.

### HasEwf

`func (o *ServicePolicy) HasEwf() bool`

HasEwf returns a boolean if a field has been set.

### GetIdp

`func (o *ServicePolicy) GetIdp() IdpConfig`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *ServicePolicy) GetIdpOk() (*IdpConfig, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *ServicePolicy) SetIdp(v IdpConfig)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *ServicePolicy) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetLocalRouting

`func (o *ServicePolicy) GetLocalRouting() bool`

GetLocalRouting returns the LocalRouting field if non-nil, zero value otherwise.

### GetLocalRoutingOk

`func (o *ServicePolicy) GetLocalRoutingOk() (*bool, bool)`

GetLocalRoutingOk returns a tuple with the LocalRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRouting

`func (o *ServicePolicy) SetLocalRouting(v bool)`

SetLocalRouting sets LocalRouting field to given value.

### HasLocalRouting

`func (o *ServicePolicy) HasLocalRouting() bool`

HasLocalRouting returns a boolean if a field has been set.

### GetName

`func (o *ServicePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPathPreferences

`func (o *ServicePolicy) GetPathPreferences() string`

GetPathPreferences returns the PathPreferences field if non-nil, zero value otherwise.

### GetPathPreferencesOk

`func (o *ServicePolicy) GetPathPreferencesOk() (*string, bool)`

GetPathPreferencesOk returns a tuple with the PathPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPreferences

`func (o *ServicePolicy) SetPathPreferences(v string)`

SetPathPreferences sets PathPreferences field to given value.

### HasPathPreferences

`func (o *ServicePolicy) HasPathPreferences() bool`

HasPathPreferences returns a boolean if a field has been set.

### GetServicepolicyId

`func (o *ServicePolicy) GetServicepolicyId() string`

GetServicepolicyId returns the ServicepolicyId field if non-nil, zero value otherwise.

### GetServicepolicyIdOk

`func (o *ServicePolicy) GetServicepolicyIdOk() (*string, bool)`

GetServicepolicyIdOk returns a tuple with the ServicepolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicepolicyId

`func (o *ServicePolicy) SetServicepolicyId(v string)`

SetServicepolicyId sets ServicepolicyId field to given value.

### HasServicepolicyId

`func (o *ServicePolicy) HasServicepolicyId() bool`

HasServicepolicyId returns a boolean if a field has been set.

### GetServices

`func (o *ServicePolicy) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServicePolicy) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServicePolicy) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *ServicePolicy) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTenants

`func (o *ServicePolicy) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ServicePolicy) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ServicePolicy) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ServicePolicy) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PskPortalSso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRoles** | Pointer to **[]string** | // allowed roles for accessing psk portal, if none, any role is permitted | [optional] 
**IdpCert** | Pointer to **string** |  | [optional] 
**IdpSignAlgo** | Pointer to **string** |  | [optional] 
**IdpSsoUrl** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**NameidFormat** | Pointer to **string** |  | [optional] 
**RoleMapping** | Pointer to **map[string]string** | Property key is the role name, property value is the SSO Attribute | [optional] 
**UseSsoRoleForPskRole** | Pointer to **bool** | if enabled, the &#x60;role&#x60; above will be ignored | [optional] 

## Methods

### NewPskPortalSso

`func NewPskPortalSso() *PskPortalSso`

NewPskPortalSso instantiates a new PskPortalSso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPskPortalSsoWithDefaults

`func NewPskPortalSsoWithDefaults() *PskPortalSso`

NewPskPortalSsoWithDefaults instantiates a new PskPortalSso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedRoles

`func (o *PskPortalSso) GetAllowedRoles() []string`

GetAllowedRoles returns the AllowedRoles field if non-nil, zero value otherwise.

### GetAllowedRolesOk

`func (o *PskPortalSso) GetAllowedRolesOk() (*[]string, bool)`

GetAllowedRolesOk returns a tuple with the AllowedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRoles

`func (o *PskPortalSso) SetAllowedRoles(v []string)`

SetAllowedRoles sets AllowedRoles field to given value.

### HasAllowedRoles

`func (o *PskPortalSso) HasAllowedRoles() bool`

HasAllowedRoles returns a boolean if a field has been set.

### GetIdpCert

`func (o *PskPortalSso) GetIdpCert() string`

GetIdpCert returns the IdpCert field if non-nil, zero value otherwise.

### GetIdpCertOk

`func (o *PskPortalSso) GetIdpCertOk() (*string, bool)`

GetIdpCertOk returns a tuple with the IdpCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCert

`func (o *PskPortalSso) SetIdpCert(v string)`

SetIdpCert sets IdpCert field to given value.

### HasIdpCert

`func (o *PskPortalSso) HasIdpCert() bool`

HasIdpCert returns a boolean if a field has been set.

### GetIdpSignAlgo

`func (o *PskPortalSso) GetIdpSignAlgo() string`

GetIdpSignAlgo returns the IdpSignAlgo field if non-nil, zero value otherwise.

### GetIdpSignAlgoOk

`func (o *PskPortalSso) GetIdpSignAlgoOk() (*string, bool)`

GetIdpSignAlgoOk returns a tuple with the IdpSignAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSignAlgo

`func (o *PskPortalSso) SetIdpSignAlgo(v string)`

SetIdpSignAlgo sets IdpSignAlgo field to given value.

### HasIdpSignAlgo

`func (o *PskPortalSso) HasIdpSignAlgo() bool`

HasIdpSignAlgo returns a boolean if a field has been set.

### GetIdpSsoUrl

`func (o *PskPortalSso) GetIdpSsoUrl() string`

GetIdpSsoUrl returns the IdpSsoUrl field if non-nil, zero value otherwise.

### GetIdpSsoUrlOk

`func (o *PskPortalSso) GetIdpSsoUrlOk() (*string, bool)`

GetIdpSsoUrlOk returns a tuple with the IdpSsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSsoUrl

`func (o *PskPortalSso) SetIdpSsoUrl(v string)`

SetIdpSsoUrl sets IdpSsoUrl field to given value.

### HasIdpSsoUrl

`func (o *PskPortalSso) HasIdpSsoUrl() bool`

HasIdpSsoUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *PskPortalSso) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PskPortalSso) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PskPortalSso) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PskPortalSso) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetNameidFormat

`func (o *PskPortalSso) GetNameidFormat() string`

GetNameidFormat returns the NameidFormat field if non-nil, zero value otherwise.

### GetNameidFormatOk

`func (o *PskPortalSso) GetNameidFormatOk() (*string, bool)`

GetNameidFormatOk returns a tuple with the NameidFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameidFormat

`func (o *PskPortalSso) SetNameidFormat(v string)`

SetNameidFormat sets NameidFormat field to given value.

### HasNameidFormat

`func (o *PskPortalSso) HasNameidFormat() bool`

HasNameidFormat returns a boolean if a field has been set.

### GetRoleMapping

`func (o *PskPortalSso) GetRoleMapping() map[string]string`

GetRoleMapping returns the RoleMapping field if non-nil, zero value otherwise.

### GetRoleMappingOk

`func (o *PskPortalSso) GetRoleMappingOk() (*map[string]string, bool)`

GetRoleMappingOk returns a tuple with the RoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMapping

`func (o *PskPortalSso) SetRoleMapping(v map[string]string)`

SetRoleMapping sets RoleMapping field to given value.

### HasRoleMapping

`func (o *PskPortalSso) HasRoleMapping() bool`

HasRoleMapping returns a boolean if a field has been set.

### GetUseSsoRoleForPskRole

`func (o *PskPortalSso) GetUseSsoRoleForPskRole() bool`

GetUseSsoRoleForPskRole returns the UseSsoRoleForPskRole field if non-nil, zero value otherwise.

### GetUseSsoRoleForPskRoleOk

`func (o *PskPortalSso) GetUseSsoRoleForPskRoleOk() (*bool, bool)`

GetUseSsoRoleForPskRoleOk returns a tuple with the UseSsoRoleForPskRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsoRoleForPskRole

`func (o *PskPortalSso) SetUseSsoRoleForPskRole(v bool)`

SetUseSsoRoleForPskRole sets UseSsoRoleForPskRole field to given value.

### HasUseSsoRoleForPskRole

`func (o *PskPortalSso) HasUseSsoRoleForPskRole() bool`

HasUseSsoRoleForPskRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



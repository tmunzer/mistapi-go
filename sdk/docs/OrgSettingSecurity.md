# OrgSettingSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableLocalSsh** | Pointer to **bool** | whether to disable local SSH (by default, local SSH is enabled with allow_mist in Org is enabled | [optional] 
**FipsZeroizePassword** | Pointer to **string** | password required to zeroize devices (FIPS) on site level | [optional] 
**LimitSshAccess** | Pointer to **bool** | whether to allow certain SSH keys to SSH into the AP (see Site:Setting) | [optional] [default to false]

## Methods

### NewOrgSettingSecurity

`func NewOrgSettingSecurity() *OrgSettingSecurity`

NewOrgSettingSecurity instantiates a new OrgSettingSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingSecurityWithDefaults

`func NewOrgSettingSecurityWithDefaults() *OrgSettingSecurity`

NewOrgSettingSecurityWithDefaults instantiates a new OrgSettingSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableLocalSsh

`func (o *OrgSettingSecurity) GetDisableLocalSsh() bool`

GetDisableLocalSsh returns the DisableLocalSsh field if non-nil, zero value otherwise.

### GetDisableLocalSshOk

`func (o *OrgSettingSecurity) GetDisableLocalSshOk() (*bool, bool)`

GetDisableLocalSshOk returns a tuple with the DisableLocalSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLocalSsh

`func (o *OrgSettingSecurity) SetDisableLocalSsh(v bool)`

SetDisableLocalSsh sets DisableLocalSsh field to given value.

### HasDisableLocalSsh

`func (o *OrgSettingSecurity) HasDisableLocalSsh() bool`

HasDisableLocalSsh returns a boolean if a field has been set.

### GetFipsZeroizePassword

`func (o *OrgSettingSecurity) GetFipsZeroizePassword() string`

GetFipsZeroizePassword returns the FipsZeroizePassword field if non-nil, zero value otherwise.

### GetFipsZeroizePasswordOk

`func (o *OrgSettingSecurity) GetFipsZeroizePasswordOk() (*string, bool)`

GetFipsZeroizePasswordOk returns a tuple with the FipsZeroizePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsZeroizePassword

`func (o *OrgSettingSecurity) SetFipsZeroizePassword(v string)`

SetFipsZeroizePassword sets FipsZeroizePassword field to given value.

### HasFipsZeroizePassword

`func (o *OrgSettingSecurity) HasFipsZeroizePassword() bool`

HasFipsZeroizePassword returns a boolean if a field has been set.

### GetLimitSshAccess

`func (o *OrgSettingSecurity) GetLimitSshAccess() bool`

GetLimitSshAccess returns the LimitSshAccess field if non-nil, zero value otherwise.

### GetLimitSshAccessOk

`func (o *OrgSettingSecurity) GetLimitSshAccessOk() (*bool, bool)`

GetLimitSshAccessOk returns a tuple with the LimitSshAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSshAccess

`func (o *OrgSettingSecurity) SetLimitSshAccess(v bool)`

SetLimitSshAccess sets LimitSshAccess field to given value.

### HasLimitSshAccess

`func (o *OrgSettingSecurity) HasLimitSshAccess() bool`

HasLimitSshAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



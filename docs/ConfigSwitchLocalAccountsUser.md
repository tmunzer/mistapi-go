# ConfigSwitchLocalAccountsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**ConfigSwitchLocalAccountsUserRole**](ConfigSwitchLocalAccountsUserRole.md) |  | [optional] [default to CONFIGSWITCHLOCALACCOUNTSUSERROLE_NONE]

## Methods

### NewConfigSwitchLocalAccountsUser

`func NewConfigSwitchLocalAccountsUser() *ConfigSwitchLocalAccountsUser`

NewConfigSwitchLocalAccountsUser instantiates a new ConfigSwitchLocalAccountsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSwitchLocalAccountsUserWithDefaults

`func NewConfigSwitchLocalAccountsUserWithDefaults() *ConfigSwitchLocalAccountsUser`

NewConfigSwitchLocalAccountsUserWithDefaults instantiates a new ConfigSwitchLocalAccountsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ConfigSwitchLocalAccountsUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConfigSwitchLocalAccountsUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConfigSwitchLocalAccountsUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ConfigSwitchLocalAccountsUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRole

`func (o *ConfigSwitchLocalAccountsUser) GetRole() ConfigSwitchLocalAccountsUserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ConfigSwitchLocalAccountsUser) GetRoleOk() (*ConfigSwitchLocalAccountsUserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ConfigSwitchLocalAccountsUser) SetRole(v ConfigSwitchLocalAccountsUserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ConfigSwitchLocalAccountsUser) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



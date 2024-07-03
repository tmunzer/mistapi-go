# OrgSettingPasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | whether the policy is enabled | [optional] [default to false]
**Freshness** | Pointer to **int32** | days, required if password policy is enabled | [optional] 
**MinLength** | Pointer to **int32** | required password length | [optional] [default to 8]
**RequiresSpecialChar** | Pointer to **bool** | whether to require special character | [optional] [default to false]
**RequiresTwoFactorAuth** | Pointer to **bool** | whether to require two-factor auth | [optional] [default to false]

## Methods

### NewOrgSettingPasswordPolicy

`func NewOrgSettingPasswordPolicy() *OrgSettingPasswordPolicy`

NewOrgSettingPasswordPolicy instantiates a new OrgSettingPasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingPasswordPolicyWithDefaults

`func NewOrgSettingPasswordPolicyWithDefaults() *OrgSettingPasswordPolicy`

NewOrgSettingPasswordPolicyWithDefaults instantiates a new OrgSettingPasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OrgSettingPasswordPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrgSettingPasswordPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrgSettingPasswordPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrgSettingPasswordPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFreshness

`func (o *OrgSettingPasswordPolicy) GetFreshness() int32`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *OrgSettingPasswordPolicy) GetFreshnessOk() (*int32, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *OrgSettingPasswordPolicy) SetFreshness(v int32)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *OrgSettingPasswordPolicy) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetMinLength

`func (o *OrgSettingPasswordPolicy) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *OrgSettingPasswordPolicy) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *OrgSettingPasswordPolicy) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *OrgSettingPasswordPolicy) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetRequiresSpecialChar

`func (o *OrgSettingPasswordPolicy) GetRequiresSpecialChar() bool`

GetRequiresSpecialChar returns the RequiresSpecialChar field if non-nil, zero value otherwise.

### GetRequiresSpecialCharOk

`func (o *OrgSettingPasswordPolicy) GetRequiresSpecialCharOk() (*bool, bool)`

GetRequiresSpecialCharOk returns a tuple with the RequiresSpecialChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSpecialChar

`func (o *OrgSettingPasswordPolicy) SetRequiresSpecialChar(v bool)`

SetRequiresSpecialChar sets RequiresSpecialChar field to given value.

### HasRequiresSpecialChar

`func (o *OrgSettingPasswordPolicy) HasRequiresSpecialChar() bool`

HasRequiresSpecialChar returns a boolean if a field has been set.

### GetRequiresTwoFactorAuth

`func (o *OrgSettingPasswordPolicy) GetRequiresTwoFactorAuth() bool`

GetRequiresTwoFactorAuth returns the RequiresTwoFactorAuth field if non-nil, zero value otherwise.

### GetRequiresTwoFactorAuthOk

`func (o *OrgSettingPasswordPolicy) GetRequiresTwoFactorAuthOk() (*bool, bool)`

GetRequiresTwoFactorAuthOk returns a tuple with the RequiresTwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTwoFactorAuth

`func (o *OrgSettingPasswordPolicy) SetRequiresTwoFactorAuth(v bool)`

SetRequiresTwoFactorAuth sets RequiresTwoFactorAuth field to given value.

### HasRequiresTwoFactorAuth

`func (o *OrgSettingPasswordPolicy) HasRequiresTwoFactorAuth() bool`

HasRequiresTwoFactorAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Admin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | Pointer to **string** |  | [optional] [readonly] 
**ComplianceStatus** | Pointer to [**AdminComplianceStatus**](AdminComplianceStatus.md) |  | [optional] 
**Email** | **string** |  | 
**EnableTwoFactor** | Pointer to **bool** |  | [optional] [readonly] 
**ExpireTime** | Pointer to **int32** |  | [optional] 
**FirstName** | **string** | for an invite, this is the original first name used | 
**Hours** | Pointer to **int32** | how long the invite should be valid | [optional] [default to 24]
**LastName** | **string** | for an invite, this is the original last name used | 
**OauthGoogle** | Pointer to **bool** |  | [optional] [readonly] 
**Phone** | Pointer to **string** | phone number (numbers only, including country code) | [optional] 
**Phone2** | Pointer to **string** | secondary phone number (numbers only, including country code) | [optional] 
**Privileges** | Pointer to [**[]PrivilegeSelf**](PrivilegeSelf.md) | list of privileges the admin has | [optional] 
**SessionExpiry** | Pointer to **int32** |  | [optional] [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] [readonly] 
**TwoFactorVerified** | Pointer to **bool** | two factor status | [optional] [readonly] 
**ViaSso** | Pointer to **bool** | an admin alogin via_sso is more restircted. (password and email cannot be changed) | [optional] [readonly] 

## Methods

### NewAdmin

`func NewAdmin(email string, firstName string, lastName string, ) *Admin`

NewAdmin instantiates a new Admin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminWithDefaults

`func NewAdminWithDefaults() *Admin`

NewAdminWithDefaults instantiates a new Admin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *Admin) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *Admin) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *Admin) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *Admin) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetComplianceStatus

`func (o *Admin) GetComplianceStatus() AdminComplianceStatus`

GetComplianceStatus returns the ComplianceStatus field if non-nil, zero value otherwise.

### GetComplianceStatusOk

`func (o *Admin) GetComplianceStatusOk() (*AdminComplianceStatus, bool)`

GetComplianceStatusOk returns a tuple with the ComplianceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceStatus

`func (o *Admin) SetComplianceStatus(v AdminComplianceStatus)`

SetComplianceStatus sets ComplianceStatus field to given value.

### HasComplianceStatus

`func (o *Admin) HasComplianceStatus() bool`

HasComplianceStatus returns a boolean if a field has been set.

### GetEmail

`func (o *Admin) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Admin) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Admin) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnableTwoFactor

`func (o *Admin) GetEnableTwoFactor() bool`

GetEnableTwoFactor returns the EnableTwoFactor field if non-nil, zero value otherwise.

### GetEnableTwoFactorOk

`func (o *Admin) GetEnableTwoFactorOk() (*bool, bool)`

GetEnableTwoFactorOk returns a tuple with the EnableTwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTwoFactor

`func (o *Admin) SetEnableTwoFactor(v bool)`

SetEnableTwoFactor sets EnableTwoFactor field to given value.

### HasEnableTwoFactor

`func (o *Admin) HasEnableTwoFactor() bool`

HasEnableTwoFactor returns a boolean if a field has been set.

### GetExpireTime

`func (o *Admin) GetExpireTime() int32`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *Admin) GetExpireTimeOk() (*int32, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *Admin) SetExpireTime(v int32)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *Admin) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetFirstName

`func (o *Admin) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Admin) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Admin) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetHours

`func (o *Admin) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *Admin) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *Admin) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *Admin) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetLastName

`func (o *Admin) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Admin) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Admin) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetOauthGoogle

`func (o *Admin) GetOauthGoogle() bool`

GetOauthGoogle returns the OauthGoogle field if non-nil, zero value otherwise.

### GetOauthGoogleOk

`func (o *Admin) GetOauthGoogleOk() (*bool, bool)`

GetOauthGoogleOk returns a tuple with the OauthGoogle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthGoogle

`func (o *Admin) SetOauthGoogle(v bool)`

SetOauthGoogle sets OauthGoogle field to given value.

### HasOauthGoogle

`func (o *Admin) HasOauthGoogle() bool`

HasOauthGoogle returns a boolean if a field has been set.

### GetPhone

`func (o *Admin) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Admin) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Admin) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Admin) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhone2

`func (o *Admin) GetPhone2() string`

GetPhone2 returns the Phone2 field if non-nil, zero value otherwise.

### GetPhone2Ok

`func (o *Admin) GetPhone2Ok() (*string, bool)`

GetPhone2Ok returns a tuple with the Phone2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone2

`func (o *Admin) SetPhone2(v string)`

SetPhone2 sets Phone2 field to given value.

### HasPhone2

`func (o *Admin) HasPhone2() bool`

HasPhone2 returns a boolean if a field has been set.

### GetPrivileges

`func (o *Admin) GetPrivileges() []PrivilegeSelf`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *Admin) GetPrivilegesOk() (*[]PrivilegeSelf, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *Admin) SetPrivileges(v []PrivilegeSelf)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *Admin) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetSessionExpiry

`func (o *Admin) GetSessionExpiry() int32`

GetSessionExpiry returns the SessionExpiry field if non-nil, zero value otherwise.

### GetSessionExpiryOk

`func (o *Admin) GetSessionExpiryOk() (*int32, bool)`

GetSessionExpiryOk returns a tuple with the SessionExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExpiry

`func (o *Admin) SetSessionExpiry(v int32)`

SetSessionExpiry sets SessionExpiry field to given value.

### HasSessionExpiry

`func (o *Admin) HasSessionExpiry() bool`

HasSessionExpiry returns a boolean if a field has been set.

### GetTags

`func (o *Admin) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Admin) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Admin) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Admin) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTwoFactorVerified

`func (o *Admin) GetTwoFactorVerified() bool`

GetTwoFactorVerified returns the TwoFactorVerified field if non-nil, zero value otherwise.

### GetTwoFactorVerifiedOk

`func (o *Admin) GetTwoFactorVerifiedOk() (*bool, bool)`

GetTwoFactorVerifiedOk returns a tuple with the TwoFactorVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorVerified

`func (o *Admin) SetTwoFactorVerified(v bool)`

SetTwoFactorVerified sets TwoFactorVerified field to given value.

### HasTwoFactorVerified

`func (o *Admin) HasTwoFactorVerified() bool`

HasTwoFactorVerified returns a boolean if a field has been set.

### GetViaSso

`func (o *Admin) GetViaSso() bool`

GetViaSso returns the ViaSso field if non-nil, zero value otherwise.

### GetViaSsoOk

`func (o *Admin) GetViaSsoOk() (*bool, bool)`

GetViaSsoOk returns a tuple with the ViaSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaSso

`func (o *Admin) SetViaSso(v bool)`

SetViaSso sets ViaSso field to given value.

### HasViaSso

`func (o *Admin) HasViaSso() bool`

HasViaSso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



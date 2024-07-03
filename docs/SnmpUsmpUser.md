# SnmpUsmpUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationPassword** | Pointer to **string** | Not required if &#x60;authentication_type&#x60;&#x3D;&#x3D;&#x60;authentication_none&#x60; include alphabetic, numeric, and special characters, but it cannot include control characters. | [optional] 
**AuthenticationType** | Pointer to [**SnmpUsmpUserAuthenticationType**](SnmpUsmpUserAuthenticationType.md) |  | [optional] 
**EncryptionPassword** | Pointer to **string** | Not required if &#x60;encryption_type&#x60;&#x3D;&#x3D;&#x60;privacy-none&#x60; include alphabetic, numeric, and special characters, but it cannot include control characters | [optional] 
**EncryptionType** | Pointer to [**SnmpUsmpUserEncryptionType**](SnmpUsmpUserEncryptionType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewSnmpUsmpUser

`func NewSnmpUsmpUser() *SnmpUsmpUser`

NewSnmpUsmpUser instantiates a new SnmpUsmpUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpUsmpUserWithDefaults

`func NewSnmpUsmpUserWithDefaults() *SnmpUsmpUser`

NewSnmpUsmpUserWithDefaults instantiates a new SnmpUsmpUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationPassword

`func (o *SnmpUsmpUser) GetAuthenticationPassword() string`

GetAuthenticationPassword returns the AuthenticationPassword field if non-nil, zero value otherwise.

### GetAuthenticationPasswordOk

`func (o *SnmpUsmpUser) GetAuthenticationPasswordOk() (*string, bool)`

GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPassword

`func (o *SnmpUsmpUser) SetAuthenticationPassword(v string)`

SetAuthenticationPassword sets AuthenticationPassword field to given value.

### HasAuthenticationPassword

`func (o *SnmpUsmpUser) HasAuthenticationPassword() bool`

HasAuthenticationPassword returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *SnmpUsmpUser) GetAuthenticationType() SnmpUsmpUserAuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *SnmpUsmpUser) GetAuthenticationTypeOk() (*SnmpUsmpUserAuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *SnmpUsmpUser) SetAuthenticationType(v SnmpUsmpUserAuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *SnmpUsmpUser) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetEncryptionPassword

`func (o *SnmpUsmpUser) GetEncryptionPassword() string`

GetEncryptionPassword returns the EncryptionPassword field if non-nil, zero value otherwise.

### GetEncryptionPasswordOk

`func (o *SnmpUsmpUser) GetEncryptionPasswordOk() (*string, bool)`

GetEncryptionPasswordOk returns a tuple with the EncryptionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassword

`func (o *SnmpUsmpUser) SetEncryptionPassword(v string)`

SetEncryptionPassword sets EncryptionPassword field to given value.

### HasEncryptionPassword

`func (o *SnmpUsmpUser) HasEncryptionPassword() bool`

HasEncryptionPassword returns a boolean if a field has been set.

### GetEncryptionType

`func (o *SnmpUsmpUser) GetEncryptionType() SnmpUsmpUserEncryptionType`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *SnmpUsmpUser) GetEncryptionTypeOk() (*SnmpUsmpUserEncryptionType, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *SnmpUsmpUser) SetEncryptionType(v SnmpUsmpUserEncryptionType)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *SnmpUsmpUser) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### GetName

`func (o *SnmpUsmpUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnmpUsmpUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnmpUsmpUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnmpUsmpUser) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



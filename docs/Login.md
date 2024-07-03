# Login

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**TwoFactor** | Pointer to **string** |  | [optional] 

## Methods

### NewLogin

`func NewLogin(email string, password string, ) *Login`

NewLogin instantiates a new Login object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWithDefaults

`func NewLoginWithDefaults() *Login`

NewLoginWithDefaults instantiates a new Login object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Login) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Login) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Login) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *Login) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Login) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Login) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTwoFactor

`func (o *Login) GetTwoFactor() string`

GetTwoFactor returns the TwoFactor field if non-nil, zero value otherwise.

### GetTwoFactorOk

`func (o *Login) GetTwoFactorOk() (*string, bool)`

GetTwoFactorOk returns a tuple with the TwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactor

`func (o *Login) SetTwoFactor(v string)`

SetTwoFactor sets TwoFactor field to given value.

### HasTwoFactor

`func (o *Login) HasTwoFactor() bool`

HasTwoFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



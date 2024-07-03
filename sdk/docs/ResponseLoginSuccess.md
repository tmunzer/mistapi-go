# ResponseLoginSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**TwoFactorPassed** | Pointer to **bool** |  | [optional] 
**TwoFactorRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewResponseLoginSuccess

`func NewResponseLoginSuccess() *ResponseLoginSuccess`

NewResponseLoginSuccess instantiates a new ResponseLoginSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLoginSuccessWithDefaults

`func NewResponseLoginSuccessWithDefaults() *ResponseLoginSuccess`

NewResponseLoginSuccessWithDefaults instantiates a new ResponseLoginSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ResponseLoginSuccess) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResponseLoginSuccess) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResponseLoginSuccess) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ResponseLoginSuccess) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTwoFactorPassed

`func (o *ResponseLoginSuccess) GetTwoFactorPassed() bool`

GetTwoFactorPassed returns the TwoFactorPassed field if non-nil, zero value otherwise.

### GetTwoFactorPassedOk

`func (o *ResponseLoginSuccess) GetTwoFactorPassedOk() (*bool, bool)`

GetTwoFactorPassedOk returns a tuple with the TwoFactorPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorPassed

`func (o *ResponseLoginSuccess) SetTwoFactorPassed(v bool)`

SetTwoFactorPassed sets TwoFactorPassed field to given value.

### HasTwoFactorPassed

`func (o *ResponseLoginSuccess) HasTwoFactorPassed() bool`

HasTwoFactorPassed returns a boolean if a field has been set.

### GetTwoFactorRequired

`func (o *ResponseLoginSuccess) GetTwoFactorRequired() bool`

GetTwoFactorRequired returns the TwoFactorRequired field if non-nil, zero value otherwise.

### GetTwoFactorRequiredOk

`func (o *ResponseLoginSuccess) GetTwoFactorRequiredOk() (*bool, bool)`

GetTwoFactorRequiredOk returns a tuple with the TwoFactorRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorRequired

`func (o *ResponseLoginSuccess) SetTwoFactorRequired(v bool)`

SetTwoFactorRequired sets TwoFactorRequired field to given value.

### HasTwoFactorRequired

`func (o *ResponseLoginSuccess) HasTwoFactorRequired() bool`

HasTwoFactorRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



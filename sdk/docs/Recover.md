# Recover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Recaptcha** | Pointer to **string** | see https://www.google.com/recaptcha/ | [optional] 
**RecaptchaFlavor** | Pointer to [**RecaptchaFlavor**](RecaptchaFlavor.md) |  | [optional] [default to RECAPTCHAFLAVOR_GOOGLE]

## Methods

### NewRecover

`func NewRecover(email string, ) *Recover`

NewRecover instantiates a new Recover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverWithDefaults

`func NewRecoverWithDefaults() *Recover`

NewRecoverWithDefaults instantiates a new Recover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Recover) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Recover) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Recover) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRecaptcha

`func (o *Recover) GetRecaptcha() string`

GetRecaptcha returns the Recaptcha field if non-nil, zero value otherwise.

### GetRecaptchaOk

`func (o *Recover) GetRecaptchaOk() (*string, bool)`

GetRecaptchaOk returns a tuple with the Recaptcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptcha

`func (o *Recover) SetRecaptcha(v string)`

SetRecaptcha sets Recaptcha field to given value.

### HasRecaptcha

`func (o *Recover) HasRecaptcha() bool`

HasRecaptcha returns a boolean if a field has been set.

### GetRecaptchaFlavor

`func (o *Recover) GetRecaptchaFlavor() RecaptchaFlavor`

GetRecaptchaFlavor returns the RecaptchaFlavor field if non-nil, zero value otherwise.

### GetRecaptchaFlavorOk

`func (o *Recover) GetRecaptchaFlavorOk() (*RecaptchaFlavor, bool)`

GetRecaptchaFlavorOk returns a tuple with the RecaptchaFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaFlavor

`func (o *Recover) SetRecaptchaFlavor(v RecaptchaFlavor)`

SetRecaptchaFlavor sets RecaptchaFlavor field to given value.

### HasRecaptchaFlavor

`func (o *Recover) HasRecaptchaFlavor() bool`

HasRecaptchaFlavor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Recaptcha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavor** | Pointer to [**RecaptchaFlavor**](RecaptchaFlavor.md) |  | [optional] [default to RECAPTCHAFLAVOR_GOOGLE]
**Required** | Pointer to **bool** |  | [optional] 
**Sitekey** | Pointer to **string** |  | [optional] 

## Methods

### NewRecaptcha

`func NewRecaptcha() *Recaptcha`

NewRecaptcha instantiates a new Recaptcha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecaptchaWithDefaults

`func NewRecaptchaWithDefaults() *Recaptcha`

NewRecaptchaWithDefaults instantiates a new Recaptcha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavor

`func (o *Recaptcha) GetFlavor() RecaptchaFlavor`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *Recaptcha) GetFlavorOk() (*RecaptchaFlavor, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *Recaptcha) SetFlavor(v RecaptchaFlavor)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *Recaptcha) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetRequired

`func (o *Recaptcha) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Recaptcha) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Recaptcha) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Recaptcha) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSitekey

`func (o *Recaptcha) GetSitekey() string`

GetSitekey returns the Sitekey field if non-nil, zero value otherwise.

### GetSitekeyOk

`func (o *Recaptcha) GetSitekeyOk() (*string, bool)`

GetSitekeyOk returns a tuple with the Sitekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitekey

`func (o *Recaptcha) SetSitekey(v string)`

SetSitekey sets Sitekey field to given value.

### HasSitekey

`func (o *Recaptcha) HasSitekey() bool`

HasSitekey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



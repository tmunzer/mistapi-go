# AdminInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountOnly** | Pointer to **bool** | skip creating initial setup if true | [optional] [default to false]
**AllowMist** | Pointer to **bool** | whether to allow Mist to look at this org | [optional] [default to false]
**City** | Pointer to **string** | city of registering user | [optional] 
**Country** | Pointer to **string** | country/region of registering user | [optional] 
**Email** | **string** |  | 
**FirstName** | **string** |  | 
**InviteCode** | Pointer to **string** | required initially | [optional] 
**LastName** | **string** |  | 
**OrgName** | **string** |  | 
**Password** | **string** |  | 
**Recaptcha** | **string** | reCAPTCHA , see https://www.google.com/recaptcha/ | 
**RecaptchaFlavor** | Pointer to [**RecaptchaFlavor**](RecaptchaFlavor.md) |  | [optional] [default to RECAPTCHAFLAVOR_GOOGLE]
**RefererInviteToken** | Pointer to **string** | the invite token to apply after account creation | [optional] 
**ReturnTo** | Pointer to **string** | the url the user should be redirected back to | [optional] 
**State** | Pointer to **string** | state of registering user, optional (depends on country/region) | [optional] 
**StreetAddress** | Pointer to **string** | street address of registering user | [optional] 
**StreetAddress2** | Pointer to **string** | street address 2 of registering user | [optional] 
**Zipcode** | Pointer to **string** | zipcode of registering user | [optional] 

## Methods

### NewAdminInvite

`func NewAdminInvite(email string, firstName string, lastName string, orgName string, password string, recaptcha string, ) *AdminInvite`

NewAdminInvite instantiates a new AdminInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminInviteWithDefaults

`func NewAdminInviteWithDefaults() *AdminInvite`

NewAdminInviteWithDefaults instantiates a new AdminInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountOnly

`func (o *AdminInvite) GetAccountOnly() bool`

GetAccountOnly returns the AccountOnly field if non-nil, zero value otherwise.

### GetAccountOnlyOk

`func (o *AdminInvite) GetAccountOnlyOk() (*bool, bool)`

GetAccountOnlyOk returns a tuple with the AccountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOnly

`func (o *AdminInvite) SetAccountOnly(v bool)`

SetAccountOnly sets AccountOnly field to given value.

### HasAccountOnly

`func (o *AdminInvite) HasAccountOnly() bool`

HasAccountOnly returns a boolean if a field has been set.

### GetAllowMist

`func (o *AdminInvite) GetAllowMist() bool`

GetAllowMist returns the AllowMist field if non-nil, zero value otherwise.

### GetAllowMistOk

`func (o *AdminInvite) GetAllowMistOk() (*bool, bool)`

GetAllowMistOk returns a tuple with the AllowMist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMist

`func (o *AdminInvite) SetAllowMist(v bool)`

SetAllowMist sets AllowMist field to given value.

### HasAllowMist

`func (o *AdminInvite) HasAllowMist() bool`

HasAllowMist returns a boolean if a field has been set.

### GetCity

`func (o *AdminInvite) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AdminInvite) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AdminInvite) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AdminInvite) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *AdminInvite) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AdminInvite) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AdminInvite) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AdminInvite) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *AdminInvite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminInvite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminInvite) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *AdminInvite) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AdminInvite) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AdminInvite) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetInviteCode

`func (o *AdminInvite) GetInviteCode() string`

GetInviteCode returns the InviteCode field if non-nil, zero value otherwise.

### GetInviteCodeOk

`func (o *AdminInvite) GetInviteCodeOk() (*string, bool)`

GetInviteCodeOk returns a tuple with the InviteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteCode

`func (o *AdminInvite) SetInviteCode(v string)`

SetInviteCode sets InviteCode field to given value.

### HasInviteCode

`func (o *AdminInvite) HasInviteCode() bool`

HasInviteCode returns a boolean if a field has been set.

### GetLastName

`func (o *AdminInvite) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AdminInvite) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AdminInvite) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetOrgName

`func (o *AdminInvite) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *AdminInvite) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *AdminInvite) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetPassword

`func (o *AdminInvite) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AdminInvite) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AdminInvite) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRecaptcha

`func (o *AdminInvite) GetRecaptcha() string`

GetRecaptcha returns the Recaptcha field if non-nil, zero value otherwise.

### GetRecaptchaOk

`func (o *AdminInvite) GetRecaptchaOk() (*string, bool)`

GetRecaptchaOk returns a tuple with the Recaptcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptcha

`func (o *AdminInvite) SetRecaptcha(v string)`

SetRecaptcha sets Recaptcha field to given value.


### GetRecaptchaFlavor

`func (o *AdminInvite) GetRecaptchaFlavor() RecaptchaFlavor`

GetRecaptchaFlavor returns the RecaptchaFlavor field if non-nil, zero value otherwise.

### GetRecaptchaFlavorOk

`func (o *AdminInvite) GetRecaptchaFlavorOk() (*RecaptchaFlavor, bool)`

GetRecaptchaFlavorOk returns a tuple with the RecaptchaFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaFlavor

`func (o *AdminInvite) SetRecaptchaFlavor(v RecaptchaFlavor)`

SetRecaptchaFlavor sets RecaptchaFlavor field to given value.

### HasRecaptchaFlavor

`func (o *AdminInvite) HasRecaptchaFlavor() bool`

HasRecaptchaFlavor returns a boolean if a field has been set.

### GetRefererInviteToken

`func (o *AdminInvite) GetRefererInviteToken() string`

GetRefererInviteToken returns the RefererInviteToken field if non-nil, zero value otherwise.

### GetRefererInviteTokenOk

`func (o *AdminInvite) GetRefererInviteTokenOk() (*string, bool)`

GetRefererInviteTokenOk returns a tuple with the RefererInviteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefererInviteToken

`func (o *AdminInvite) SetRefererInviteToken(v string)`

SetRefererInviteToken sets RefererInviteToken field to given value.

### HasRefererInviteToken

`func (o *AdminInvite) HasRefererInviteToken() bool`

HasRefererInviteToken returns a boolean if a field has been set.

### GetReturnTo

`func (o *AdminInvite) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *AdminInvite) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *AdminInvite) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.

### HasReturnTo

`func (o *AdminInvite) HasReturnTo() bool`

HasReturnTo returns a boolean if a field has been set.

### GetState

`func (o *AdminInvite) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AdminInvite) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AdminInvite) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AdminInvite) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreetAddress

`func (o *AdminInvite) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *AdminInvite) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *AdminInvite) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *AdminInvite) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetStreetAddress2

`func (o *AdminInvite) GetStreetAddress2() string`

GetStreetAddress2 returns the StreetAddress2 field if non-nil, zero value otherwise.

### GetStreetAddress2Ok

`func (o *AdminInvite) GetStreetAddress2Ok() (*string, bool)`

GetStreetAddress2Ok returns a tuple with the StreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress2

`func (o *AdminInvite) SetStreetAddress2(v string)`

SetStreetAddress2 sets StreetAddress2 field to given value.

### HasStreetAddress2

`func (o *AdminInvite) HasStreetAddress2() bool`

HasStreetAddress2 returns a boolean if a field has been set.

### GetZipcode

`func (o *AdminInvite) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *AdminInvite) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *AdminInvite) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *AdminInvite) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NacPortalSso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpCert** | Pointer to **string** |  | [optional] 
**IdpSignAlgo** | Pointer to **string** |  | [optional] 
**IdpSsoUrl** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**NameidFormat** | Pointer to **string** |  | [optional] 
**SsoRoleMatching** | Pointer to [**[]NacPortalSsoRoleMatching**](NacPortalSsoRoleMatching.md) |  | [optional] 
**UseSsoRoleForCert** | Pointer to **bool** | if it&#39;s desired to inject a role into Cert&#39;s Subject (so it can be used later on in policy) | [optional] 

## Methods

### NewNacPortalSso

`func NewNacPortalSso() *NacPortalSso`

NewNacPortalSso instantiates a new NacPortalSso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNacPortalSsoWithDefaults

`func NewNacPortalSsoWithDefaults() *NacPortalSso`

NewNacPortalSsoWithDefaults instantiates a new NacPortalSso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpCert

`func (o *NacPortalSso) GetIdpCert() string`

GetIdpCert returns the IdpCert field if non-nil, zero value otherwise.

### GetIdpCertOk

`func (o *NacPortalSso) GetIdpCertOk() (*string, bool)`

GetIdpCertOk returns a tuple with the IdpCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCert

`func (o *NacPortalSso) SetIdpCert(v string)`

SetIdpCert sets IdpCert field to given value.

### HasIdpCert

`func (o *NacPortalSso) HasIdpCert() bool`

HasIdpCert returns a boolean if a field has been set.

### GetIdpSignAlgo

`func (o *NacPortalSso) GetIdpSignAlgo() string`

GetIdpSignAlgo returns the IdpSignAlgo field if non-nil, zero value otherwise.

### GetIdpSignAlgoOk

`func (o *NacPortalSso) GetIdpSignAlgoOk() (*string, bool)`

GetIdpSignAlgoOk returns a tuple with the IdpSignAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSignAlgo

`func (o *NacPortalSso) SetIdpSignAlgo(v string)`

SetIdpSignAlgo sets IdpSignAlgo field to given value.

### HasIdpSignAlgo

`func (o *NacPortalSso) HasIdpSignAlgo() bool`

HasIdpSignAlgo returns a boolean if a field has been set.

### GetIdpSsoUrl

`func (o *NacPortalSso) GetIdpSsoUrl() string`

GetIdpSsoUrl returns the IdpSsoUrl field if non-nil, zero value otherwise.

### GetIdpSsoUrlOk

`func (o *NacPortalSso) GetIdpSsoUrlOk() (*string, bool)`

GetIdpSsoUrlOk returns a tuple with the IdpSsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSsoUrl

`func (o *NacPortalSso) SetIdpSsoUrl(v string)`

SetIdpSsoUrl sets IdpSsoUrl field to given value.

### HasIdpSsoUrl

`func (o *NacPortalSso) HasIdpSsoUrl() bool`

HasIdpSsoUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *NacPortalSso) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *NacPortalSso) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *NacPortalSso) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *NacPortalSso) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetNameidFormat

`func (o *NacPortalSso) GetNameidFormat() string`

GetNameidFormat returns the NameidFormat field if non-nil, zero value otherwise.

### GetNameidFormatOk

`func (o *NacPortalSso) GetNameidFormatOk() (*string, bool)`

GetNameidFormatOk returns a tuple with the NameidFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameidFormat

`func (o *NacPortalSso) SetNameidFormat(v string)`

SetNameidFormat sets NameidFormat field to given value.

### HasNameidFormat

`func (o *NacPortalSso) HasNameidFormat() bool`

HasNameidFormat returns a boolean if a field has been set.

### GetSsoRoleMatching

`func (o *NacPortalSso) GetSsoRoleMatching() []NacPortalSsoRoleMatching`

GetSsoRoleMatching returns the SsoRoleMatching field if non-nil, zero value otherwise.

### GetSsoRoleMatchingOk

`func (o *NacPortalSso) GetSsoRoleMatchingOk() (*[]NacPortalSsoRoleMatching, bool)`

GetSsoRoleMatchingOk returns a tuple with the SsoRoleMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoRoleMatching

`func (o *NacPortalSso) SetSsoRoleMatching(v []NacPortalSsoRoleMatching)`

SetSsoRoleMatching sets SsoRoleMatching field to given value.

### HasSsoRoleMatching

`func (o *NacPortalSso) HasSsoRoleMatching() bool`

HasSsoRoleMatching returns a boolean if a field has been set.

### GetUseSsoRoleForCert

`func (o *NacPortalSso) GetUseSsoRoleForCert() bool`

GetUseSsoRoleForCert returns the UseSsoRoleForCert field if non-nil, zero value otherwise.

### GetUseSsoRoleForCertOk

`func (o *NacPortalSso) GetUseSsoRoleForCertOk() (*bool, bool)`

GetUseSsoRoleForCertOk returns a tuple with the UseSsoRoleForCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsoRoleForCert

`func (o *NacPortalSso) SetUseSsoRoleForCert(v bool)`

SetUseSsoRoleForCert sets UseSsoRoleForCert field to given value.

### HasUseSsoRoleForCert

`func (o *NacPortalSso) HasUseSsoRoleForCert() bool`

HasUseSsoRoleForCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



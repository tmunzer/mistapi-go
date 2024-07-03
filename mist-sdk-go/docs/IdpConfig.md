# IdpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertOnly** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**IdpprofileId** | Pointer to **string** | org_level IDP Profile can be used, this takes precedence over &#x60;profile&#x60; | [optional] 
**Profile** | Pointer to **string** | &#x60;strict&#x60; (default) / &#x60;standard&#x60; / or keys from from idp_profiles | [optional] [default to "strict"]

## Methods

### NewIdpConfig

`func NewIdpConfig() *IdpConfig`

NewIdpConfig instantiates a new IdpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpConfigWithDefaults

`func NewIdpConfigWithDefaults() *IdpConfig`

NewIdpConfigWithDefaults instantiates a new IdpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertOnly

`func (o *IdpConfig) GetAlertOnly() bool`

GetAlertOnly returns the AlertOnly field if non-nil, zero value otherwise.

### GetAlertOnlyOk

`func (o *IdpConfig) GetAlertOnlyOk() (*bool, bool)`

GetAlertOnlyOk returns a tuple with the AlertOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnly

`func (o *IdpConfig) SetAlertOnly(v bool)`

SetAlertOnly sets AlertOnly field to given value.

### HasAlertOnly

`func (o *IdpConfig) HasAlertOnly() bool`

HasAlertOnly returns a boolean if a field has been set.

### GetEnabled

`func (o *IdpConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdpConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdpConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdpConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIdpprofileId

`func (o *IdpConfig) GetIdpprofileId() string`

GetIdpprofileId returns the IdpprofileId field if non-nil, zero value otherwise.

### GetIdpprofileIdOk

`func (o *IdpConfig) GetIdpprofileIdOk() (*string, bool)`

GetIdpprofileIdOk returns a tuple with the IdpprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpprofileId

`func (o *IdpConfig) SetIdpprofileId(v string)`

SetIdpprofileId sets IdpprofileId field to given value.

### HasIdpprofileId

`func (o *IdpConfig) HasIdpprofileId() bool`

HasIdpprofileId returns a boolean if a field has been set.

### GetProfile

`func (o *IdpConfig) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *IdpConfig) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *IdpConfig) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *IdpConfig) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



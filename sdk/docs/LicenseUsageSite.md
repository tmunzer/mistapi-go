# LicenseUsageSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgEntitled** | **map[string]int32** | license entitlement for the entire org | 
**SvnaEnabled** | **bool** | eligibility for the Switch SLE | 
**TrialEnabled** | **bool** |  | 
**Usages** | **map[string]int32** | subscriptions and their quantities | 
**VnaEligible** | **bool** | eligibility for the AP/Client SLE | 
**VnaUi** | **bool** | if True, Conversational Assistant and Marvis Action available | 
**WvnaEligible** | **bool** | eligibility for the WAN SLE | 

## Methods

### NewLicenseUsageSite

`func NewLicenseUsageSite(orgEntitled map[string]int32, svnaEnabled bool, trialEnabled bool, usages map[string]int32, vnaEligible bool, vnaUi bool, wvnaEligible bool, ) *LicenseUsageSite`

NewLicenseUsageSite instantiates a new LicenseUsageSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseUsageSiteWithDefaults

`func NewLicenseUsageSiteWithDefaults() *LicenseUsageSite`

NewLicenseUsageSiteWithDefaults instantiates a new LicenseUsageSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgEntitled

`func (o *LicenseUsageSite) GetOrgEntitled() map[string]int32`

GetOrgEntitled returns the OrgEntitled field if non-nil, zero value otherwise.

### GetOrgEntitledOk

`func (o *LicenseUsageSite) GetOrgEntitledOk() (*map[string]int32, bool)`

GetOrgEntitledOk returns a tuple with the OrgEntitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgEntitled

`func (o *LicenseUsageSite) SetOrgEntitled(v map[string]int32)`

SetOrgEntitled sets OrgEntitled field to given value.


### GetSvnaEnabled

`func (o *LicenseUsageSite) GetSvnaEnabled() bool`

GetSvnaEnabled returns the SvnaEnabled field if non-nil, zero value otherwise.

### GetSvnaEnabledOk

`func (o *LicenseUsageSite) GetSvnaEnabledOk() (*bool, bool)`

GetSvnaEnabledOk returns a tuple with the SvnaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvnaEnabled

`func (o *LicenseUsageSite) SetSvnaEnabled(v bool)`

SetSvnaEnabled sets SvnaEnabled field to given value.


### GetTrialEnabled

`func (o *LicenseUsageSite) GetTrialEnabled() bool`

GetTrialEnabled returns the TrialEnabled field if non-nil, zero value otherwise.

### GetTrialEnabledOk

`func (o *LicenseUsageSite) GetTrialEnabledOk() (*bool, bool)`

GetTrialEnabledOk returns a tuple with the TrialEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnabled

`func (o *LicenseUsageSite) SetTrialEnabled(v bool)`

SetTrialEnabled sets TrialEnabled field to given value.


### GetUsages

`func (o *LicenseUsageSite) GetUsages() map[string]int32`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *LicenseUsageSite) GetUsagesOk() (*map[string]int32, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *LicenseUsageSite) SetUsages(v map[string]int32)`

SetUsages sets Usages field to given value.


### GetVnaEligible

`func (o *LicenseUsageSite) GetVnaEligible() bool`

GetVnaEligible returns the VnaEligible field if non-nil, zero value otherwise.

### GetVnaEligibleOk

`func (o *LicenseUsageSite) GetVnaEligibleOk() (*bool, bool)`

GetVnaEligibleOk returns a tuple with the VnaEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnaEligible

`func (o *LicenseUsageSite) SetVnaEligible(v bool)`

SetVnaEligible sets VnaEligible field to given value.


### GetVnaUi

`func (o *LicenseUsageSite) GetVnaUi() bool`

GetVnaUi returns the VnaUi field if non-nil, zero value otherwise.

### GetVnaUiOk

`func (o *LicenseUsageSite) GetVnaUiOk() (*bool, bool)`

GetVnaUiOk returns a tuple with the VnaUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnaUi

`func (o *LicenseUsageSite) SetVnaUi(v bool)`

SetVnaUi sets VnaUi field to given value.


### GetWvnaEligible

`func (o *LicenseUsageSite) GetWvnaEligible() bool`

GetWvnaEligible returns the WvnaEligible field if non-nil, zero value otherwise.

### GetWvnaEligibleOk

`func (o *LicenseUsageSite) GetWvnaEligibleOk() (*bool, bool)`

GetWvnaEligibleOk returns a tuple with the WvnaEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWvnaEligible

`func (o *LicenseUsageSite) SetWvnaEligible(v bool)`

SetWvnaEligible sets WvnaEligible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



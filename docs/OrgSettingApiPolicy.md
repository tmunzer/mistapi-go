# OrgSettingApiPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoReveal** | Pointer to **bool** | by default, API hides password/secrets when the user doesn&#39;t have write access * &#x60;true&#x60;: API will hide passwords/secrets for all users * &#x60;false&#x60;: API will hide passwords/secrests for read-only users | [optional] [default to false]

## Methods

### NewOrgSettingApiPolicy

`func NewOrgSettingApiPolicy() *OrgSettingApiPolicy`

NewOrgSettingApiPolicy instantiates a new OrgSettingApiPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingApiPolicyWithDefaults

`func NewOrgSettingApiPolicyWithDefaults() *OrgSettingApiPolicy`

NewOrgSettingApiPolicyWithDefaults instantiates a new OrgSettingApiPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoReveal

`func (o *OrgSettingApiPolicy) GetNoReveal() bool`

GetNoReveal returns the NoReveal field if non-nil, zero value otherwise.

### GetNoRevealOk

`func (o *OrgSettingApiPolicy) GetNoRevealOk() (*bool, bool)`

GetNoRevealOk returns a tuple with the NoReveal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoReveal

`func (o *OrgSettingApiPolicy) SetNoReveal(v bool)`

SetNoReveal sets NoReveal field to given value.

### HasNoReveal

`func (o *OrgSettingApiPolicy) HasNoReveal() bool`

HasNoReveal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



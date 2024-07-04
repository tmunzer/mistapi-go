# OrgSettingMistNacIdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeRealms** | Pointer to **[]string** | when the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**UserRealms** | Pointer to **[]string** | which realm should trigger this IDP. we extract user realm from 1. Username-AVP (&#x60;mist.com&#x60; from john@mist.com) 2. Cert CN | [optional] 

## Methods

### NewOrgSettingMistNacIdp

`func NewOrgSettingMistNacIdp() *OrgSettingMistNacIdp`

NewOrgSettingMistNacIdp instantiates a new OrgSettingMistNacIdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingMistNacIdpWithDefaults

`func NewOrgSettingMistNacIdpWithDefaults() *OrgSettingMistNacIdp`

NewOrgSettingMistNacIdpWithDefaults instantiates a new OrgSettingMistNacIdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeRealms

`func (o *OrgSettingMistNacIdp) GetExcludeRealms() []string`

GetExcludeRealms returns the ExcludeRealms field if non-nil, zero value otherwise.

### GetExcludeRealmsOk

`func (o *OrgSettingMistNacIdp) GetExcludeRealmsOk() (*[]string, bool)`

GetExcludeRealmsOk returns a tuple with the ExcludeRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRealms

`func (o *OrgSettingMistNacIdp) SetExcludeRealms(v []string)`

SetExcludeRealms sets ExcludeRealms field to given value.

### HasExcludeRealms

`func (o *OrgSettingMistNacIdp) HasExcludeRealms() bool`

HasExcludeRealms returns a boolean if a field has been set.

### GetId

`func (o *OrgSettingMistNacIdp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgSettingMistNacIdp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgSettingMistNacIdp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgSettingMistNacIdp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserRealms

`func (o *OrgSettingMistNacIdp) GetUserRealms() []string`

GetUserRealms returns the UserRealms field if non-nil, zero value otherwise.

### GetUserRealmsOk

`func (o *OrgSettingMistNacIdp) GetUserRealmsOk() (*[]string, bool)`

GetUserRealmsOk returns a tuple with the UserRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRealms

`func (o *OrgSettingMistNacIdp) SetUserRealms(v []string)`

SetUserRealms sets UserRealms field to given value.

### HasUserRealms

`func (o *OrgSettingMistNacIdp) HasUserRealms() bool`

HasUserRealms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



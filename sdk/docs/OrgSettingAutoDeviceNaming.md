# OrgSettingAutoDeviceNaming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]OrgAutoRules**](OrgAutoRules.md) |  | [optional] 

## Methods

### NewOrgSettingAutoDeviceNaming

`func NewOrgSettingAutoDeviceNaming() *OrgSettingAutoDeviceNaming`

NewOrgSettingAutoDeviceNaming instantiates a new OrgSettingAutoDeviceNaming object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingAutoDeviceNamingWithDefaults

`func NewOrgSettingAutoDeviceNamingWithDefaults() *OrgSettingAutoDeviceNaming`

NewOrgSettingAutoDeviceNamingWithDefaults instantiates a new OrgSettingAutoDeviceNaming object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *OrgSettingAutoDeviceNaming) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OrgSettingAutoDeviceNaming) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OrgSettingAutoDeviceNaming) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *OrgSettingAutoDeviceNaming) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetRules

`func (o *OrgSettingAutoDeviceNaming) GetRules() []OrgAutoRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *OrgSettingAutoDeviceNaming) GetRulesOk() (*[]OrgAutoRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *OrgSettingAutoDeviceNaming) SetRules(v []OrgAutoRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *OrgSettingAutoDeviceNaming) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *OrgSettingAutoDeviceNaming) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *OrgSettingAutoDeviceNaming) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



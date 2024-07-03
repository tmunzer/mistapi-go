# OrgSettingAutoDeviceprofileAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]OrgAutoRules**](OrgAutoRules.md) |  | [optional] 

## Methods

### NewOrgSettingAutoDeviceprofileAssignment

`func NewOrgSettingAutoDeviceprofileAssignment() *OrgSettingAutoDeviceprofileAssignment`

NewOrgSettingAutoDeviceprofileAssignment instantiates a new OrgSettingAutoDeviceprofileAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingAutoDeviceprofileAssignmentWithDefaults

`func NewOrgSettingAutoDeviceprofileAssignmentWithDefaults() *OrgSettingAutoDeviceprofileAssignment`

NewOrgSettingAutoDeviceprofileAssignmentWithDefaults instantiates a new OrgSettingAutoDeviceprofileAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *OrgSettingAutoDeviceprofileAssignment) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OrgSettingAutoDeviceprofileAssignment) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OrgSettingAutoDeviceprofileAssignment) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *OrgSettingAutoDeviceprofileAssignment) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetRules

`func (o *OrgSettingAutoDeviceprofileAssignment) GetRules() []OrgAutoRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *OrgSettingAutoDeviceprofileAssignment) GetRulesOk() (*[]OrgAutoRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *OrgSettingAutoDeviceprofileAssignment) SetRules(v []OrgAutoRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *OrgSettingAutoDeviceprofileAssignment) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *OrgSettingAutoDeviceprofileAssignment) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *OrgSettingAutoDeviceprofileAssignment) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



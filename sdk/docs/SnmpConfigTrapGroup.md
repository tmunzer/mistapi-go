# SnmpConfigTrapGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** |  | [optional] 
**GroupName** | Pointer to **string** | Categories list can refer to https://www.juniper.net/documentation/software/topics/task/configuration/snmp_trap-groups-configuring-junos-nm.html | [optional] 
**Targets** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to [**SnmpConfigTrapVerion**](SnmpConfigTrapVerion.md) |  | [optional] [default to SNMPCONFIGTRAPVERION_V2]

## Methods

### NewSnmpConfigTrapGroup

`func NewSnmpConfigTrapGroup() *SnmpConfigTrapGroup`

NewSnmpConfigTrapGroup instantiates a new SnmpConfigTrapGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpConfigTrapGroupWithDefaults

`func NewSnmpConfigTrapGroupWithDefaults() *SnmpConfigTrapGroup`

NewSnmpConfigTrapGroupWithDefaults instantiates a new SnmpConfigTrapGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *SnmpConfigTrapGroup) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *SnmpConfigTrapGroup) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *SnmpConfigTrapGroup) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *SnmpConfigTrapGroup) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetGroupName

`func (o *SnmpConfigTrapGroup) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *SnmpConfigTrapGroup) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *SnmpConfigTrapGroup) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *SnmpConfigTrapGroup) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetTargets

`func (o *SnmpConfigTrapGroup) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SnmpConfigTrapGroup) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SnmpConfigTrapGroup) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *SnmpConfigTrapGroup) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetVersion

`func (o *SnmpConfigTrapGroup) GetVersion() SnmpConfigTrapVerion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnmpConfigTrapGroup) GetVersionOk() (*SnmpConfigTrapVerion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnmpConfigTrapGroup) SetVersion(v SnmpConfigTrapVerion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnmpConfigTrapGroup) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OrgSettingMgmt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MxtunnelIds** | Pointer to **[]string** | list of Mist Tunnels | [optional] 
**UseMxtunnel** | Pointer to **bool** | whether to use Mist Tunnel for mgmt connectivity, this takes precedence over use_wxtunnel | [optional] [default to false]
**UseWxtunnel** | Pointer to **bool** | whether to use wxtunnel for mgmt connectivity | [optional] [default to false]

## Methods

### NewOrgSettingMgmt

`func NewOrgSettingMgmt() *OrgSettingMgmt`

NewOrgSettingMgmt instantiates a new OrgSettingMgmt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingMgmtWithDefaults

`func NewOrgSettingMgmtWithDefaults() *OrgSettingMgmt`

NewOrgSettingMgmtWithDefaults instantiates a new OrgSettingMgmt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMxtunnelIds

`func (o *OrgSettingMgmt) GetMxtunnelIds() []string`

GetMxtunnelIds returns the MxtunnelIds field if non-nil, zero value otherwise.

### GetMxtunnelIdsOk

`func (o *OrgSettingMgmt) GetMxtunnelIdsOk() (*[]string, bool)`

GetMxtunnelIdsOk returns a tuple with the MxtunnelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxtunnelIds

`func (o *OrgSettingMgmt) SetMxtunnelIds(v []string)`

SetMxtunnelIds sets MxtunnelIds field to given value.

### HasMxtunnelIds

`func (o *OrgSettingMgmt) HasMxtunnelIds() bool`

HasMxtunnelIds returns a boolean if a field has been set.

### GetUseMxtunnel

`func (o *OrgSettingMgmt) GetUseMxtunnel() bool`

GetUseMxtunnel returns the UseMxtunnel field if non-nil, zero value otherwise.

### GetUseMxtunnelOk

`func (o *OrgSettingMgmt) GetUseMxtunnelOk() (*bool, bool)`

GetUseMxtunnelOk returns a tuple with the UseMxtunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMxtunnel

`func (o *OrgSettingMgmt) SetUseMxtunnel(v bool)`

SetUseMxtunnel sets UseMxtunnel field to given value.

### HasUseMxtunnel

`func (o *OrgSettingMgmt) HasUseMxtunnel() bool`

HasUseMxtunnel returns a boolean if a field has been set.

### GetUseWxtunnel

`func (o *OrgSettingMgmt) GetUseWxtunnel() bool`

GetUseWxtunnel returns the UseWxtunnel field if non-nil, zero value otherwise.

### GetUseWxtunnelOk

`func (o *OrgSettingMgmt) GetUseWxtunnelOk() (*bool, bool)`

GetUseWxtunnelOk returns a tuple with the UseWxtunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWxtunnel

`func (o *OrgSettingMgmt) SetUseWxtunnel(v bool)`

SetUseWxtunnel sets UseWxtunnel field to given value.

### HasUseWxtunnel

`func (o *OrgSettingMgmt) HasUseWxtunnel() bool`

HasUseWxtunnel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



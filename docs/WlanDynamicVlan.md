# WlanDynamicVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVlanId** | Pointer to **NullableInt32** | vlan_id to use when there’s no match from RADIUS | [optional] [default to 999]
**Enabled** | Pointer to **bool** | whether to enable dynamic vlan | [optional] [default to false]
**LocalVlanIds** | Pointer to **[]int32** | vlan_ids to be locally bridged | [optional] 
**Type** | Pointer to [**WlanDynamicVlanType**](WlanDynamicVlanType.md) |  | [optional] [default to WLANDYNAMICVLANTYPE_STANDARD]
**Vlans** | Pointer to **map[string]string** | map between vlan_id (as string) to airespace interface names (comma-separated) or null for stndard mapping   * if &#x60;dynamic_vlan.type&#x60;&#x3D;&#x3D;&#x60;standard&#x60;, property key is the Vlan ID and property value is \&quot;\&quot;   * if &#x60;dynamic_vlan.type&#x60;&#x3D;&#x3D;&#x60;airespace-interface-name&#x60;, property key is the Vlan ID and property value is the Airespace Interface Name | [optional] 

## Methods

### NewWlanDynamicVlan

`func NewWlanDynamicVlan() *WlanDynamicVlan`

NewWlanDynamicVlan instantiates a new WlanDynamicVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanDynamicVlanWithDefaults

`func NewWlanDynamicVlanWithDefaults() *WlanDynamicVlan`

NewWlanDynamicVlanWithDefaults instantiates a new WlanDynamicVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVlanId

`func (o *WlanDynamicVlan) GetDefaultVlanId() int32`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *WlanDynamicVlan) GetDefaultVlanIdOk() (*int32, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *WlanDynamicVlan) SetDefaultVlanId(v int32)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *WlanDynamicVlan) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### SetDefaultVlanIdNil

`func (o *WlanDynamicVlan) SetDefaultVlanIdNil(b bool)`

 SetDefaultVlanIdNil sets the value for DefaultVlanId to be an explicit nil

### UnsetDefaultVlanId
`func (o *WlanDynamicVlan) UnsetDefaultVlanId()`

UnsetDefaultVlanId ensures that no value is present for DefaultVlanId, not even an explicit nil
### GetEnabled

`func (o *WlanDynamicVlan) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanDynamicVlan) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanDynamicVlan) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanDynamicVlan) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocalVlanIds

`func (o *WlanDynamicVlan) GetLocalVlanIds() []*int32`

GetLocalVlanIds returns the LocalVlanIds field if non-nil, zero value otherwise.

### GetLocalVlanIdsOk

`func (o *WlanDynamicVlan) GetLocalVlanIdsOk() (*[]*int32, bool)`

GetLocalVlanIdsOk returns a tuple with the LocalVlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanIds

`func (o *WlanDynamicVlan) SetLocalVlanIds(v []*int32)`

SetLocalVlanIds sets LocalVlanIds field to given value.

### HasLocalVlanIds

`func (o *WlanDynamicVlan) HasLocalVlanIds() bool`

HasLocalVlanIds returns a boolean if a field has been set.

### GetType

`func (o *WlanDynamicVlan) GetType() WlanDynamicVlanType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WlanDynamicVlan) GetTypeOk() (*WlanDynamicVlanType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WlanDynamicVlan) SetType(v WlanDynamicVlanType)`

SetType sets Type field to given value.

### HasType

`func (o *WlanDynamicVlan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlans

`func (o *WlanDynamicVlan) GetVlans() map[string]string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *WlanDynamicVlan) GetVlansOk() (*map[string]string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *WlanDynamicVlan) SetVlans(v map[string]string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *WlanDynamicVlan) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



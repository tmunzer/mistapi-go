# WlanAppQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**map[string]WlanAppQosAppsProperties**](WlanAppQosAppsProperties.md) |  | [optional] [default to {}]
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Others** | Pointer to [**[]WlanAppQosOthersItem**](WlanAppQosOthersItem.md) |  | [optional] 

## Methods

### NewWlanAppQos

`func NewWlanAppQos() *WlanAppQos`

NewWlanAppQos instantiates a new WlanAppQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanAppQosWithDefaults

`func NewWlanAppQosWithDefaults() *WlanAppQos`

NewWlanAppQosWithDefaults instantiates a new WlanAppQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *WlanAppQos) GetApps() map[string]WlanAppQosAppsProperties`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *WlanAppQos) GetAppsOk() (*map[string]WlanAppQosAppsProperties, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *WlanAppQos) SetApps(v map[string]WlanAppQosAppsProperties)`

SetApps sets Apps field to given value.

### HasApps

`func (o *WlanAppQos) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetEnabled

`func (o *WlanAppQos) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanAppQos) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanAppQos) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanAppQos) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOthers

`func (o *WlanAppQos) GetOthers() []WlanAppQosOthersItem`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *WlanAppQos) GetOthersOk() (*[]WlanAppQosOthersItem, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *WlanAppQos) SetOthers(v []WlanAppQosOthersItem)`

SetOthers sets Others field to given value.

### HasOthers

`func (o *WlanAppQos) HasOthers() bool`

HasOthers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



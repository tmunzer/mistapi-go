# WlanAppLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to **map[string]int32** | Map from app key to bandwidth in kbps.  Property key is the app key, defined in Get Application List | [optional] [default to {}]
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**WxtagIds** | Pointer to **map[string]int32** | Map from wxtag_id of Hostname Wxlan Tags to bandwidth in kbps Property key is the wxtag id | [optional] [default to {}]

## Methods

### NewWlanAppLimit

`func NewWlanAppLimit() *WlanAppLimit`

NewWlanAppLimit instantiates a new WlanAppLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanAppLimitWithDefaults

`func NewWlanAppLimitWithDefaults() *WlanAppLimit`

NewWlanAppLimitWithDefaults instantiates a new WlanAppLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *WlanAppLimit) GetApps() map[string]int32`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *WlanAppLimit) GetAppsOk() (*map[string]int32, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *WlanAppLimit) SetApps(v map[string]int32)`

SetApps sets Apps field to given value.

### HasApps

`func (o *WlanAppLimit) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetEnabled

`func (o *WlanAppLimit) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanAppLimit) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanAppLimit) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanAppLimit) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWxtagIds

`func (o *WlanAppLimit) GetWxtagIds() map[string]int32`

GetWxtagIds returns the WxtagIds field if non-nil, zero value otherwise.

### GetWxtagIdsOk

`func (o *WlanAppLimit) GetWxtagIdsOk() (*map[string]int32, bool)`

GetWxtagIdsOk returns a tuple with the WxtagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxtagIds

`func (o *WlanAppLimit) SetWxtagIds(v map[string]int32)`

SetWxtagIds sets WxtagIds field to given value.

### HasWxtagIds

`func (o *WlanAppLimit) HasWxtagIds() bool`

HasWxtagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



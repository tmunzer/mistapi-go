# WlanBonjour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalVlanIds** | **[]int32** | additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses | [default to []]
**Enabled** | Pointer to **bool** | whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false | [optional] [default to false]
**Services** | [**map[string]WlanBonjourServiceProperties**](WlanBonjourServiceProperties.md) | what services are allowed.  Property key is the service name | [default to {}]

## Methods

### NewWlanBonjour

`func NewWlanBonjour(additionalVlanIds []int32, services map[string]WlanBonjourServiceProperties, ) *WlanBonjour`

NewWlanBonjour instantiates a new WlanBonjour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanBonjourWithDefaults

`func NewWlanBonjourWithDefaults() *WlanBonjour`

NewWlanBonjourWithDefaults instantiates a new WlanBonjour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalVlanIds

`func (o *WlanBonjour) GetAdditionalVlanIds() []int32`

GetAdditionalVlanIds returns the AdditionalVlanIds field if non-nil, zero value otherwise.

### GetAdditionalVlanIdsOk

`func (o *WlanBonjour) GetAdditionalVlanIdsOk() (*[]int32, bool)`

GetAdditionalVlanIdsOk returns a tuple with the AdditionalVlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalVlanIds

`func (o *WlanBonjour) SetAdditionalVlanIds(v []int32)`

SetAdditionalVlanIds sets AdditionalVlanIds field to given value.


### GetEnabled

`func (o *WlanBonjour) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanBonjour) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanBonjour) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanBonjour) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServices

`func (o *WlanBonjour) GetServices() map[string]WlanBonjourServiceProperties`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *WlanBonjour) GetServicesOk() (*map[string]WlanBonjourServiceProperties, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *WlanBonjour) SetServices(v map[string]WlanBonjourServiceProperties)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



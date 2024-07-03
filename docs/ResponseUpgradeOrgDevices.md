# ResponseUpgradeOrgDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableP2p** | Pointer to **bool** | whether to allow local AP-to-AP FW upgrade | [optional] 
**Force** | Pointer to **bool** | whether to force upgrade when requested version is same as running version | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to [**DeviceUpgradeStrategy**](DeviceUpgradeStrategy.md) |  | [optional] [default to DEVICEUPGRADESTRATEGY_BIG_BANG]
**TargetVersion** | Pointer to **string** | version to upgrade to | [optional] 
**Upgrades** | Pointer to [**[]ResponseUpgradeOrgDevice**](ResponseUpgradeOrgDevice.md) |  | [optional] 

## Methods

### NewResponseUpgradeOrgDevices

`func NewResponseUpgradeOrgDevices() *ResponseUpgradeOrgDevices`

NewResponseUpgradeOrgDevices instantiates a new ResponseUpgradeOrgDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseUpgradeOrgDevicesWithDefaults

`func NewResponseUpgradeOrgDevicesWithDefaults() *ResponseUpgradeOrgDevices`

NewResponseUpgradeOrgDevicesWithDefaults instantiates a new ResponseUpgradeOrgDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableP2p

`func (o *ResponseUpgradeOrgDevices) GetEnableP2p() bool`

GetEnableP2p returns the EnableP2p field if non-nil, zero value otherwise.

### GetEnableP2pOk

`func (o *ResponseUpgradeOrgDevices) GetEnableP2pOk() (*bool, bool)`

GetEnableP2pOk returns a tuple with the EnableP2p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableP2p

`func (o *ResponseUpgradeOrgDevices) SetEnableP2p(v bool)`

SetEnableP2p sets EnableP2p field to given value.

### HasEnableP2p

`func (o *ResponseUpgradeOrgDevices) HasEnableP2p() bool`

HasEnableP2p returns a boolean if a field has been set.

### GetForce

`func (o *ResponseUpgradeOrgDevices) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *ResponseUpgradeOrgDevices) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *ResponseUpgradeOrgDevices) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *ResponseUpgradeOrgDevices) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetId

`func (o *ResponseUpgradeOrgDevices) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseUpgradeOrgDevices) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseUpgradeOrgDevices) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseUpgradeOrgDevices) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStrategy

`func (o *ResponseUpgradeOrgDevices) GetStrategy() DeviceUpgradeStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *ResponseUpgradeOrgDevices) GetStrategyOk() (*DeviceUpgradeStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *ResponseUpgradeOrgDevices) SetStrategy(v DeviceUpgradeStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *ResponseUpgradeOrgDevices) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetTargetVersion

`func (o *ResponseUpgradeOrgDevices) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *ResponseUpgradeOrgDevices) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *ResponseUpgradeOrgDevices) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *ResponseUpgradeOrgDevices) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetUpgrades

`func (o *ResponseUpgradeOrgDevices) GetUpgrades() []ResponseUpgradeOrgDevice`

GetUpgrades returns the Upgrades field if non-nil, zero value otherwise.

### GetUpgradesOk

`func (o *ResponseUpgradeOrgDevices) GetUpgradesOk() (*[]ResponseUpgradeOrgDevice, bool)`

GetUpgradesOk returns a tuple with the Upgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrades

`func (o *ResponseUpgradeOrgDevices) SetUpgrades(v []ResponseUpgradeOrgDevice)`

SetUpgrades sets Upgrades field to given value.

### HasUpgrades

`func (o *ResponseUpgradeOrgDevices) HasUpgrades() bool`

HasUpgrades returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



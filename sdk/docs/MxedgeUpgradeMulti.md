# MxedgeUpgradeMulti

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDowngrades** | Pointer to [**MxedgeUpgradeMultiAllowDowngrades**](MxedgeUpgradeMultiAllowDowngrades.md) |  | [optional] 
**Channel** | Pointer to [**MxedgeUpgradeChannel**](MxedgeUpgradeChannel.md) |  | [optional] [default to MXEDGEUPGRADECHANNEL_STABLE]
**Distro** | Pointer to **string** | distro upgrade, optional, to specific codename (e.g. bullseye) with highest qualified versions | [optional] 
**MxedgeIds** | **[]string** | list of mxedge IDs to upgrade. If not specified, it means all the org mxedges. | 
**Strategy** | Pointer to [**MxedgeUpgradeStrategy**](MxedgeUpgradeStrategy.md) |  | [optional] [default to MXEDGEUPGRADESTRATEGY_BIG_BANG]
**Versions** | Pointer to [**MxedgeUpgradeVersion**](MxedgeUpgradeVersion.md) |  | [optional] 

## Methods

### NewMxedgeUpgradeMulti

`func NewMxedgeUpgradeMulti(mxedgeIds []string, ) *MxedgeUpgradeMulti`

NewMxedgeUpgradeMulti instantiates a new MxedgeUpgradeMulti object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeUpgradeMultiWithDefaults

`func NewMxedgeUpgradeMultiWithDefaults() *MxedgeUpgradeMulti`

NewMxedgeUpgradeMultiWithDefaults instantiates a new MxedgeUpgradeMulti object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDowngrades

`func (o *MxedgeUpgradeMulti) GetAllowDowngrades() MxedgeUpgradeMultiAllowDowngrades`

GetAllowDowngrades returns the AllowDowngrades field if non-nil, zero value otherwise.

### GetAllowDowngradesOk

`func (o *MxedgeUpgradeMulti) GetAllowDowngradesOk() (*MxedgeUpgradeMultiAllowDowngrades, bool)`

GetAllowDowngradesOk returns a tuple with the AllowDowngrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDowngrades

`func (o *MxedgeUpgradeMulti) SetAllowDowngrades(v MxedgeUpgradeMultiAllowDowngrades)`

SetAllowDowngrades sets AllowDowngrades field to given value.

### HasAllowDowngrades

`func (o *MxedgeUpgradeMulti) HasAllowDowngrades() bool`

HasAllowDowngrades returns a boolean if a field has been set.

### GetChannel

`func (o *MxedgeUpgradeMulti) GetChannel() MxedgeUpgradeChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MxedgeUpgradeMulti) GetChannelOk() (*MxedgeUpgradeChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MxedgeUpgradeMulti) SetChannel(v MxedgeUpgradeChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *MxedgeUpgradeMulti) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDistro

`func (o *MxedgeUpgradeMulti) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *MxedgeUpgradeMulti) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *MxedgeUpgradeMulti) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *MxedgeUpgradeMulti) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetMxedgeIds

`func (o *MxedgeUpgradeMulti) GetMxedgeIds() []string`

GetMxedgeIds returns the MxedgeIds field if non-nil, zero value otherwise.

### GetMxedgeIdsOk

`func (o *MxedgeUpgradeMulti) GetMxedgeIdsOk() (*[]string, bool)`

GetMxedgeIdsOk returns a tuple with the MxedgeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeIds

`func (o *MxedgeUpgradeMulti) SetMxedgeIds(v []string)`

SetMxedgeIds sets MxedgeIds field to given value.


### GetStrategy

`func (o *MxedgeUpgradeMulti) GetStrategy() MxedgeUpgradeStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *MxedgeUpgradeMulti) GetStrategyOk() (*MxedgeUpgradeStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *MxedgeUpgradeMulti) SetStrategy(v MxedgeUpgradeStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *MxedgeUpgradeMulti) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetVersions

`func (o *MxedgeUpgradeMulti) GetVersions() MxedgeUpgradeVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MxedgeUpgradeMulti) GetVersionsOk() (*MxedgeUpgradeVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *MxedgeUpgradeMulti) SetVersions(v MxedgeUpgradeVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *MxedgeUpgradeMulti) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SnmpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientList** | Pointer to [**[]SnmpConfigClientList**](SnmpConfigClientList.md) |  | [optional] 
**Contact** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**EngineId** | Pointer to [**SnmpConfigEngineId**](SnmpConfigEngineId.md) |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] [default to "default"]
**TrapGroups** | Pointer to [**[]SnmpConfigTrapGroup**](SnmpConfigTrapGroup.md) |  | [optional] 
**V2cConfig** | Pointer to [**[]SnmpConfigV2cConfig**](SnmpConfigV2cConfig.md) |  | [optional] 
**V3Config** | Pointer to [**Snmpv3Config**](Snmpv3Config.md) |  | [optional] 
**Views** | Pointer to [**SnmpConfigViews**](SnmpConfigViews.md) |  | [optional] 

## Methods

### NewSnmpConfig

`func NewSnmpConfig() *SnmpConfig`

NewSnmpConfig instantiates a new SnmpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpConfigWithDefaults

`func NewSnmpConfigWithDefaults() *SnmpConfig`

NewSnmpConfigWithDefaults instantiates a new SnmpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientList

`func (o *SnmpConfig) GetClientList() []SnmpConfigClientList`

GetClientList returns the ClientList field if non-nil, zero value otherwise.

### GetClientListOk

`func (o *SnmpConfig) GetClientListOk() (*[]SnmpConfigClientList, bool)`

GetClientListOk returns a tuple with the ClientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientList

`func (o *SnmpConfig) SetClientList(v []SnmpConfigClientList)`

SetClientList sets ClientList field to given value.

### HasClientList

`func (o *SnmpConfig) HasClientList() bool`

HasClientList returns a boolean if a field has been set.

### GetContact

`func (o *SnmpConfig) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SnmpConfig) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SnmpConfig) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *SnmpConfig) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDescription

`func (o *SnmpConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnmpConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnmpConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnmpConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SnmpConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnmpConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEngineId

`func (o *SnmpConfig) GetEngineId() SnmpConfigEngineId`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *SnmpConfig) GetEngineIdOk() (*SnmpConfigEngineId, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *SnmpConfig) SetEngineId(v SnmpConfigEngineId)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *SnmpConfig) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetLocation

`func (o *SnmpConfig) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SnmpConfig) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SnmpConfig) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SnmpConfig) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *SnmpConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnmpConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnmpConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnmpConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *SnmpConfig) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SnmpConfig) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SnmpConfig) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *SnmpConfig) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTrapGroups

`func (o *SnmpConfig) GetTrapGroups() []SnmpConfigTrapGroup`

GetTrapGroups returns the TrapGroups field if non-nil, zero value otherwise.

### GetTrapGroupsOk

`func (o *SnmpConfig) GetTrapGroupsOk() (*[]SnmpConfigTrapGroup, bool)`

GetTrapGroupsOk returns a tuple with the TrapGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapGroups

`func (o *SnmpConfig) SetTrapGroups(v []SnmpConfigTrapGroup)`

SetTrapGroups sets TrapGroups field to given value.

### HasTrapGroups

`func (o *SnmpConfig) HasTrapGroups() bool`

HasTrapGroups returns a boolean if a field has been set.

### GetV2cConfig

`func (o *SnmpConfig) GetV2cConfig() []SnmpConfigV2cConfig`

GetV2cConfig returns the V2cConfig field if non-nil, zero value otherwise.

### GetV2cConfigOk

`func (o *SnmpConfig) GetV2cConfigOk() (*[]SnmpConfigV2cConfig, bool)`

GetV2cConfigOk returns a tuple with the V2cConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2cConfig

`func (o *SnmpConfig) SetV2cConfig(v []SnmpConfigV2cConfig)`

SetV2cConfig sets V2cConfig field to given value.

### HasV2cConfig

`func (o *SnmpConfig) HasV2cConfig() bool`

HasV2cConfig returns a boolean if a field has been set.

### GetV3Config

`func (o *SnmpConfig) GetV3Config() Snmpv3Config`

GetV3Config returns the V3Config field if non-nil, zero value otherwise.

### GetV3ConfigOk

`func (o *SnmpConfig) GetV3ConfigOk() (*Snmpv3Config, bool)`

GetV3ConfigOk returns a tuple with the V3Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Config

`func (o *SnmpConfig) SetV3Config(v Snmpv3Config)`

SetV3Config sets V3Config field to given value.

### HasV3Config

`func (o *SnmpConfig) HasV3Config() bool`

HasV3Config returns a boolean if a field has been set.

### GetViews

`func (o *SnmpConfig) GetViews() SnmpConfigViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *SnmpConfig) GetViewsOk() (*SnmpConfigViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *SnmpConfig) SetViews(v SnmpConfigViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *SnmpConfig) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



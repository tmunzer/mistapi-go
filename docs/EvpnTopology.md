# EvpnTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ConfigSwitch**](ConfigSwitch.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Overwrite** | Pointer to **bool** |  | [optional] 
**PodNames** | Pointer to **map[string]string** | Property key is the pod number | [optional] 
**Switches** | [**[]EvpnTopologySwitch**](EvpnTopologySwitch.md) |  | 

## Methods

### NewEvpnTopology

`func NewEvpnTopology(switches []EvpnTopologySwitch, ) *EvpnTopology`

NewEvpnTopology instantiates a new EvpnTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvpnTopologyWithDefaults

`func NewEvpnTopologyWithDefaults() *EvpnTopology`

NewEvpnTopologyWithDefaults instantiates a new EvpnTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *EvpnTopology) GetConfig() ConfigSwitch`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EvpnTopology) GetConfigOk() (*ConfigSwitch, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EvpnTopology) SetConfig(v ConfigSwitch)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *EvpnTopology) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetId

`func (o *EvpnTopology) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvpnTopology) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvpnTopology) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EvpnTopology) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EvpnTopology) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvpnTopology) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvpnTopology) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EvpnTopology) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverwrite

`func (o *EvpnTopology) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *EvpnTopology) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *EvpnTopology) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *EvpnTopology) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetPodNames

`func (o *EvpnTopology) GetPodNames() map[string]string`

GetPodNames returns the PodNames field if non-nil, zero value otherwise.

### GetPodNamesOk

`func (o *EvpnTopology) GetPodNamesOk() (*map[string]string, bool)`

GetPodNamesOk returns a tuple with the PodNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodNames

`func (o *EvpnTopology) SetPodNames(v map[string]string)`

SetPodNames sets PodNames field to given value.

### HasPodNames

`func (o *EvpnTopology) HasPodNames() bool`

HasPodNames returns a boolean if a field has been set.

### GetSwitches

`func (o *EvpnTopology) GetSwitches() []EvpnTopologySwitch`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *EvpnTopology) GetSwitchesOk() (*[]EvpnTopologySwitch, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *EvpnTopology) SetSwitches(v []EvpnTopologySwitch)`

SetSwitches sets Switches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



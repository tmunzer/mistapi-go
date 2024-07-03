# SwitchStatsApRedundancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modules** | Pointer to [**map[string]SwitchStatsApRedundancyModule**](SwitchStatsApRedundancyModule.md) | for a VC / stacked switches. | [optional] 
**NumAps** | Pointer to **int32** |  | [optional] 
**NumApsWithSwitchRedundancy** | Pointer to **int32** |  | [optional] 

## Methods

### NewSwitchStatsApRedundancy

`func NewSwitchStatsApRedundancy() *SwitchStatsApRedundancy`

NewSwitchStatsApRedundancy instantiates a new SwitchStatsApRedundancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchStatsApRedundancyWithDefaults

`func NewSwitchStatsApRedundancyWithDefaults() *SwitchStatsApRedundancy`

NewSwitchStatsApRedundancyWithDefaults instantiates a new SwitchStatsApRedundancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModules

`func (o *SwitchStatsApRedundancy) GetModules() map[string]SwitchStatsApRedundancyModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *SwitchStatsApRedundancy) GetModulesOk() (*map[string]SwitchStatsApRedundancyModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *SwitchStatsApRedundancy) SetModules(v map[string]SwitchStatsApRedundancyModule)`

SetModules sets Modules field to given value.

### HasModules

`func (o *SwitchStatsApRedundancy) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetNumAps

`func (o *SwitchStatsApRedundancy) GetNumAps() int32`

GetNumAps returns the NumAps field if non-nil, zero value otherwise.

### GetNumApsOk

`func (o *SwitchStatsApRedundancy) GetNumApsOk() (*int32, bool)`

GetNumApsOk returns a tuple with the NumAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAps

`func (o *SwitchStatsApRedundancy) SetNumAps(v int32)`

SetNumAps sets NumAps field to given value.

### HasNumAps

`func (o *SwitchStatsApRedundancy) HasNumAps() bool`

HasNumAps returns a boolean if a field has been set.

### GetNumApsWithSwitchRedundancy

`func (o *SwitchStatsApRedundancy) GetNumApsWithSwitchRedundancy() int32`

GetNumApsWithSwitchRedundancy returns the NumApsWithSwitchRedundancy field if non-nil, zero value otherwise.

### GetNumApsWithSwitchRedundancyOk

`func (o *SwitchStatsApRedundancy) GetNumApsWithSwitchRedundancyOk() (*int32, bool)`

GetNumApsWithSwitchRedundancyOk returns a tuple with the NumApsWithSwitchRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApsWithSwitchRedundancy

`func (o *SwitchStatsApRedundancy) SetNumApsWithSwitchRedundancy(v int32)`

SetNumApsWithSwitchRedundancy sets NumApsWithSwitchRedundancy field to given value.

### HasNumApsWithSwitchRedundancy

`func (o *SwitchStatsApRedundancy) HasNumApsWithSwitchRedundancy() bool`

HasNumApsWithSwitchRedundancy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ApRedundancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modules** | Pointer to [**map[string]ApRedundancyModule**](ApRedundancyModule.md) | Property key is the node id | [optional] 
**NumAps** | Pointer to **int32** |  | [optional] 
**NumApsWithSwitchRedundancy** | Pointer to **int32** |  | [optional] 

## Methods

### NewApRedundancy

`func NewApRedundancy() *ApRedundancy`

NewApRedundancy instantiates a new ApRedundancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRedundancyWithDefaults

`func NewApRedundancyWithDefaults() *ApRedundancy`

NewApRedundancyWithDefaults instantiates a new ApRedundancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModules

`func (o *ApRedundancy) GetModules() map[string]ApRedundancyModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ApRedundancy) GetModulesOk() (*map[string]ApRedundancyModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ApRedundancy) SetModules(v map[string]ApRedundancyModule)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ApRedundancy) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetNumAps

`func (o *ApRedundancy) GetNumAps() int32`

GetNumAps returns the NumAps field if non-nil, zero value otherwise.

### GetNumApsOk

`func (o *ApRedundancy) GetNumApsOk() (*int32, bool)`

GetNumApsOk returns a tuple with the NumAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAps

`func (o *ApRedundancy) SetNumAps(v int32)`

SetNumAps sets NumAps field to given value.

### HasNumAps

`func (o *ApRedundancy) HasNumAps() bool`

HasNumAps returns a boolean if a field has been set.

### GetNumApsWithSwitchRedundancy

`func (o *ApRedundancy) GetNumApsWithSwitchRedundancy() int32`

GetNumApsWithSwitchRedundancy returns the NumApsWithSwitchRedundancy field if non-nil, zero value otherwise.

### GetNumApsWithSwitchRedundancyOk

`func (o *ApRedundancy) GetNumApsWithSwitchRedundancyOk() (*int32, bool)`

GetNumApsWithSwitchRedundancyOk returns a tuple with the NumApsWithSwitchRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApsWithSwitchRedundancy

`func (o *ApRedundancy) SetNumApsWithSwitchRedundancy(v int32)`

SetNumApsWithSwitchRedundancy sets NumApsWithSwitchRedundancy field to given value.

### HasNumApsWithSwitchRedundancy

`func (o *ApRedundancy) HasNumApsWithSwitchRedundancy() bool`

HasNumApsWithSwitchRedundancy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



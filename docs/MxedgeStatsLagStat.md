# MxedgeStatsLagStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePorts** | Pointer to **[]string** | list of ports active on the LAG defined by the LACP | [optional] 

## Methods

### NewMxedgeStatsLagStat

`func NewMxedgeStatsLagStat() *MxedgeStatsLagStat`

NewMxedgeStatsLagStat instantiates a new MxedgeStatsLagStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeStatsLagStatWithDefaults

`func NewMxedgeStatsLagStatWithDefaults() *MxedgeStatsLagStat`

NewMxedgeStatsLagStatWithDefaults instantiates a new MxedgeStatsLagStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePorts

`func (o *MxedgeStatsLagStat) GetActivePorts() []string`

GetActivePorts returns the ActivePorts field if non-nil, zero value otherwise.

### GetActivePortsOk

`func (o *MxedgeStatsLagStat) GetActivePortsOk() (*[]string, bool)`

GetActivePortsOk returns a tuple with the ActivePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePorts

`func (o *MxedgeStatsLagStat) SetActivePorts(v []string)`

SetActivePorts sets ActivePorts field to given value.

### HasActivePorts

`func (o *MxedgeStatsLagStat) HasActivePorts() bool`

HasActivePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



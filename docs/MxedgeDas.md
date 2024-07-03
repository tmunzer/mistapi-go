# MxedgeDas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoaServers** | Pointer to [**[]MxedgeDasCoaServer**](MxedgeDasCoaServer.md) | dynamic authorization clients configured to send CoA|DM to mist edges on port 3799 | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewMxedgeDas

`func NewMxedgeDas() *MxedgeDas`

NewMxedgeDas instantiates a new MxedgeDas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeDasWithDefaults

`func NewMxedgeDasWithDefaults() *MxedgeDas`

NewMxedgeDasWithDefaults instantiates a new MxedgeDas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoaServers

`func (o *MxedgeDas) GetCoaServers() []MxedgeDasCoaServer`

GetCoaServers returns the CoaServers field if non-nil, zero value otherwise.

### GetCoaServersOk

`func (o *MxedgeDas) GetCoaServersOk() (*[]MxedgeDasCoaServer, bool)`

GetCoaServersOk returns a tuple with the CoaServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaServers

`func (o *MxedgeDas) SetCoaServers(v []MxedgeDasCoaServer)`

SetCoaServers sets CoaServers field to given value.

### HasCoaServers

`func (o *MxedgeDas) HasCoaServers() bool`

HasCoaServers returns a boolean if a field has been set.

### GetEnabled

`func (o *MxedgeDas) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MxedgeDas) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MxedgeDas) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MxedgeDas) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



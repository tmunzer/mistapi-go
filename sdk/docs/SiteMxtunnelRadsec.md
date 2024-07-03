# SiteMxtunnelRadsec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctServers** | Pointer to [**[]RadiusAcctServer**](RadiusAcctServer.md) |  | [optional] 
**AuthServers** | Pointer to [**[]RadiusAuthServer**](RadiusAuthServer.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**UseMxedge** | Pointer to **bool** |  | [optional] 

## Methods

### NewSiteMxtunnelRadsec

`func NewSiteMxtunnelRadsec() *SiteMxtunnelRadsec`

NewSiteMxtunnelRadsec instantiates a new SiteMxtunnelRadsec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteMxtunnelRadsecWithDefaults

`func NewSiteMxtunnelRadsecWithDefaults() *SiteMxtunnelRadsec`

NewSiteMxtunnelRadsecWithDefaults instantiates a new SiteMxtunnelRadsec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctServers

`func (o *SiteMxtunnelRadsec) GetAcctServers() []RadiusAcctServer`

GetAcctServers returns the AcctServers field if non-nil, zero value otherwise.

### GetAcctServersOk

`func (o *SiteMxtunnelRadsec) GetAcctServersOk() (*[]RadiusAcctServer, bool)`

GetAcctServersOk returns a tuple with the AcctServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctServers

`func (o *SiteMxtunnelRadsec) SetAcctServers(v []RadiusAcctServer)`

SetAcctServers sets AcctServers field to given value.

### HasAcctServers

`func (o *SiteMxtunnelRadsec) HasAcctServers() bool`

HasAcctServers returns a boolean if a field has been set.

### GetAuthServers

`func (o *SiteMxtunnelRadsec) GetAuthServers() []RadiusAuthServer`

GetAuthServers returns the AuthServers field if non-nil, zero value otherwise.

### GetAuthServersOk

`func (o *SiteMxtunnelRadsec) GetAuthServersOk() (*[]RadiusAuthServer, bool)`

GetAuthServersOk returns a tuple with the AuthServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServers

`func (o *SiteMxtunnelRadsec) SetAuthServers(v []RadiusAuthServer)`

SetAuthServers sets AuthServers field to given value.

### HasAuthServers

`func (o *SiteMxtunnelRadsec) HasAuthServers() bool`

HasAuthServers returns a boolean if a field has been set.

### GetEnabled

`func (o *SiteMxtunnelRadsec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SiteMxtunnelRadsec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SiteMxtunnelRadsec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SiteMxtunnelRadsec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUseMxedge

`func (o *SiteMxtunnelRadsec) GetUseMxedge() bool`

GetUseMxedge returns the UseMxedge field if non-nil, zero value otherwise.

### GetUseMxedgeOk

`func (o *SiteMxtunnelRadsec) GetUseMxedgeOk() (*bool, bool)`

GetUseMxedgeOk returns a tuple with the UseMxedge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMxedge

`func (o *SiteMxtunnelRadsec) SetUseMxedge(v bool)`

SetUseMxedge sets UseMxedge field to given value.

### HasUseMxedge

`func (o *SiteMxtunnelRadsec) HasUseMxedge() bool`

HasUseMxedge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InventoryUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableAutoConfig** | Pointer to **bool** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, a **cloud-ready** switch/gateway will be managed/configured by Mist by default, this disabled the behavior | [optional] [default to false]
**Macs** | Pointer to **[]string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, &#x60;op&#x60;&#x3D;&#x3D;&#x60;unassign&#x60;, &#x60;op&#x60;&#x3D;&#x3D;&#x60;upgrade_to_mist&#x60;or &#x60;op&#x60;&#x3D;&#x3D;&#x60;downgrade_to_jsi&#x60; , list of MAC, e.g. [\&quot;5c5b350e0001\&quot;] | [optional] 
**Managed** | Pointer to **bool** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, an **adopted** switch/gateway will not be managed/configured by Mist by default, this enables the behavior | [optional] [default to false]
**NoReassign** | Pointer to **bool** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, if true, treat site assignment against an already assigned AP as error | [optional] 
**Op** | [**InventoryUpdateOperation**](InventoryUpdateOperation.md) |  | 
**Serials** | Pointer to **[]string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;delete&#x60;, list of serial numbers, e.g. [\&quot;FXLH2015150025\&quot;] | [optional] 
**SiteId** | Pointer to **string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, target site id | [optional] 

## Methods

### NewInventoryUpdate

`func NewInventoryUpdate(op InventoryUpdateOperation, ) *InventoryUpdate`

NewInventoryUpdate instantiates a new InventoryUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryUpdateWithDefaults

`func NewInventoryUpdateWithDefaults() *InventoryUpdate`

NewInventoryUpdateWithDefaults instantiates a new InventoryUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableAutoConfig

`func (o *InventoryUpdate) GetDisableAutoConfig() bool`

GetDisableAutoConfig returns the DisableAutoConfig field if non-nil, zero value otherwise.

### GetDisableAutoConfigOk

`func (o *InventoryUpdate) GetDisableAutoConfigOk() (*bool, bool)`

GetDisableAutoConfigOk returns a tuple with the DisableAutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoConfig

`func (o *InventoryUpdate) SetDisableAutoConfig(v bool)`

SetDisableAutoConfig sets DisableAutoConfig field to given value.

### HasDisableAutoConfig

`func (o *InventoryUpdate) HasDisableAutoConfig() bool`

HasDisableAutoConfig returns a boolean if a field has been set.

### GetMacs

`func (o *InventoryUpdate) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *InventoryUpdate) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *InventoryUpdate) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *InventoryUpdate) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetManaged

`func (o *InventoryUpdate) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *InventoryUpdate) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *InventoryUpdate) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *InventoryUpdate) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetNoReassign

`func (o *InventoryUpdate) GetNoReassign() bool`

GetNoReassign returns the NoReassign field if non-nil, zero value otherwise.

### GetNoReassignOk

`func (o *InventoryUpdate) GetNoReassignOk() (*bool, bool)`

GetNoReassignOk returns a tuple with the NoReassign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoReassign

`func (o *InventoryUpdate) SetNoReassign(v bool)`

SetNoReassign sets NoReassign field to given value.

### HasNoReassign

`func (o *InventoryUpdate) HasNoReassign() bool`

HasNoReassign returns a boolean if a field has been set.

### GetOp

`func (o *InventoryUpdate) GetOp() InventoryUpdateOperation`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *InventoryUpdate) GetOpOk() (*InventoryUpdateOperation, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *InventoryUpdate) SetOp(v InventoryUpdateOperation)`

SetOp sets Op field to given value.


### GetSerials

`func (o *InventoryUpdate) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InventoryUpdate) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InventoryUpdate) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InventoryUpdate) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetSiteId

`func (o *InventoryUpdate) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InventoryUpdate) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InventoryUpdate) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InventoryUpdate) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



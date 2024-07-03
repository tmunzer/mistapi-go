# ReplaceDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discard** | Pointer to **[]string** | attributes that you don’t want to copy | [optional] 
**InventoryMac** | Pointer to **string** | MAC Address of the inventory that will be replacing the old one. It has to be claimed and unassigned | [optional] 
**Mac** | Pointer to **string** | MAC Address of the device to replace | [optional] 
**SiteId** | Pointer to **string** | the site_id of the device to be replaced | [optional] 
**TuntermPortConfig** | Pointer to [**TuntermPortConfig**](TuntermPortConfig.md) |  | [optional] 

## Methods

### NewReplaceDevice

`func NewReplaceDevice() *ReplaceDevice`

NewReplaceDevice instantiates a new ReplaceDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceDeviceWithDefaults

`func NewReplaceDeviceWithDefaults() *ReplaceDevice`

NewReplaceDeviceWithDefaults instantiates a new ReplaceDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscard

`func (o *ReplaceDevice) GetDiscard() []string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *ReplaceDevice) GetDiscardOk() (*[]string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *ReplaceDevice) SetDiscard(v []string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *ReplaceDevice) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetInventoryMac

`func (o *ReplaceDevice) GetInventoryMac() string`

GetInventoryMac returns the InventoryMac field if non-nil, zero value otherwise.

### GetInventoryMacOk

`func (o *ReplaceDevice) GetInventoryMacOk() (*string, bool)`

GetInventoryMacOk returns a tuple with the InventoryMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryMac

`func (o *ReplaceDevice) SetInventoryMac(v string)`

SetInventoryMac sets InventoryMac field to given value.

### HasInventoryMac

`func (o *ReplaceDevice) HasInventoryMac() bool`

HasInventoryMac returns a boolean if a field has been set.

### GetMac

`func (o *ReplaceDevice) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ReplaceDevice) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ReplaceDevice) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ReplaceDevice) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetSiteId

`func (o *ReplaceDevice) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ReplaceDevice) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ReplaceDevice) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ReplaceDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTuntermPortConfig

`func (o *ReplaceDevice) GetTuntermPortConfig() TuntermPortConfig`

GetTuntermPortConfig returns the TuntermPortConfig field if non-nil, zero value otherwise.

### GetTuntermPortConfigOk

`func (o *ReplaceDevice) GetTuntermPortConfigOk() (*TuntermPortConfig, bool)`

GetTuntermPortConfigOk returns a tuple with the TuntermPortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermPortConfig

`func (o *ReplaceDevice) SetTuntermPortConfig(v TuntermPortConfig)`

SetTuntermPortConfig sets TuntermPortConfig field to given value.

### HasTuntermPortConfig

`func (o *ReplaceDevice) HasTuntermPortConfig() bool`

HasTuntermPortConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



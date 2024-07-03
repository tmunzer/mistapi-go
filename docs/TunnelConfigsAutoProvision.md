# TunnelConfigsAutoProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Latlng** | Pointer to [**LatLng**](LatLng.md) |  | [optional] 
**Primary** | Pointer to [**TunnelConfigsAutoProvisionNode**](TunnelConfigsAutoProvisionNode.md) |  | [optional] 
**Region** | Pointer to [**TunnelConfigsAutoProvisionRegion**](TunnelConfigsAutoProvisionRegion.md) |  | [optional] [default to TUNNELCONFIGSAUTOPROVISIONREGION_AUTO]
**Secondary** | Pointer to [**TunnelConfigsAutoProvisionNode**](TunnelConfigsAutoProvisionNode.md) |  | [optional] 

## Methods

### NewTunnelConfigsAutoProvision

`func NewTunnelConfigsAutoProvision() *TunnelConfigsAutoProvision`

NewTunnelConfigsAutoProvision instantiates a new TunnelConfigsAutoProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelConfigsAutoProvisionWithDefaults

`func NewTunnelConfigsAutoProvisionWithDefaults() *TunnelConfigsAutoProvision`

NewTunnelConfigsAutoProvisionWithDefaults instantiates a new TunnelConfigsAutoProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *TunnelConfigsAutoProvision) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *TunnelConfigsAutoProvision) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *TunnelConfigsAutoProvision) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *TunnelConfigsAutoProvision) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetLatlng

`func (o *TunnelConfigsAutoProvision) GetLatlng() LatLng`

GetLatlng returns the Latlng field if non-nil, zero value otherwise.

### GetLatlngOk

`func (o *TunnelConfigsAutoProvision) GetLatlngOk() (*LatLng, bool)`

GetLatlngOk returns a tuple with the Latlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatlng

`func (o *TunnelConfigsAutoProvision) SetLatlng(v LatLng)`

SetLatlng sets Latlng field to given value.

### HasLatlng

`func (o *TunnelConfigsAutoProvision) HasLatlng() bool`

HasLatlng returns a boolean if a field has been set.

### GetPrimary

`func (o *TunnelConfigsAutoProvision) GetPrimary() TunnelConfigsAutoProvisionNode`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *TunnelConfigsAutoProvision) GetPrimaryOk() (*TunnelConfigsAutoProvisionNode, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *TunnelConfigsAutoProvision) SetPrimary(v TunnelConfigsAutoProvisionNode)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *TunnelConfigsAutoProvision) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetRegion

`func (o *TunnelConfigsAutoProvision) GetRegion() TunnelConfigsAutoProvisionRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TunnelConfigsAutoProvision) GetRegionOk() (*TunnelConfigsAutoProvisionRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TunnelConfigsAutoProvision) SetRegion(v TunnelConfigsAutoProvisionRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TunnelConfigsAutoProvision) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecondary

`func (o *TunnelConfigsAutoProvision) GetSecondary() TunnelConfigsAutoProvisionNode`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *TunnelConfigsAutoProvision) GetSecondaryOk() (*TunnelConfigsAutoProvisionNode, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *TunnelConfigsAutoProvision) SetSecondary(v TunnelConfigsAutoProvisionNode)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *TunnelConfigsAutoProvision) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WiredClientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | Pointer to **[]string** |  | [optional] [readonly] 
**DeviceMacPort** | Pointer to [**[]WiredClientResponseDeviceMacPortItem**](WiredClientResponseDeviceMacPortItem.md) |  | [optional] [readonly] 
**Ip** | Pointer to **[]string** |  | [optional] [readonly] 
**Mac** | Pointer to **string** |  | [optional] [readonly] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PortId** | Pointer to **[]string** |  | [optional] [readonly] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **float32** |  | [optional] [readonly] 
**Vlan** | Pointer to **[]int32** |  | [optional] [readonly] 

## Methods

### NewWiredClientResponse

`func NewWiredClientResponse() *WiredClientResponse`

NewWiredClientResponse instantiates a new WiredClientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWiredClientResponseWithDefaults

`func NewWiredClientResponseWithDefaults() *WiredClientResponse`

NewWiredClientResponseWithDefaults instantiates a new WiredClientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *WiredClientResponse) GetDeviceMac() []string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *WiredClientResponse) GetDeviceMacOk() (*[]string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *WiredClientResponse) SetDeviceMac(v []string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *WiredClientResponse) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceMacPort

`func (o *WiredClientResponse) GetDeviceMacPort() []WiredClientResponseDeviceMacPortItem`

GetDeviceMacPort returns the DeviceMacPort field if non-nil, zero value otherwise.

### GetDeviceMacPortOk

`func (o *WiredClientResponse) GetDeviceMacPortOk() (*[]WiredClientResponseDeviceMacPortItem, bool)`

GetDeviceMacPortOk returns a tuple with the DeviceMacPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacPort

`func (o *WiredClientResponse) SetDeviceMacPort(v []WiredClientResponseDeviceMacPortItem)`

SetDeviceMacPort sets DeviceMacPort field to given value.

### HasDeviceMacPort

`func (o *WiredClientResponse) HasDeviceMacPort() bool`

HasDeviceMacPort returns a boolean if a field has been set.

### GetIp

`func (o *WiredClientResponse) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *WiredClientResponse) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *WiredClientResponse) SetIp(v []string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *WiredClientResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *WiredClientResponse) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WiredClientResponse) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WiredClientResponse) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WiredClientResponse) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetOrgId

`func (o *WiredClientResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WiredClientResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WiredClientResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WiredClientResponse) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPortId

`func (o *WiredClientResponse) GetPortId() []string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *WiredClientResponse) GetPortIdOk() (*[]string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *WiredClientResponse) SetPortId(v []string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *WiredClientResponse) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSiteId

`func (o *WiredClientResponse) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WiredClientResponse) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WiredClientResponse) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WiredClientResponse) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WiredClientResponse) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WiredClientResponse) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WiredClientResponse) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WiredClientResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVlan

`func (o *WiredClientResponse) GetVlan() []int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *WiredClientResponse) GetVlanOk() (*[]int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *WiredClientResponse) SetVlan(v []int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *WiredClientResponse) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



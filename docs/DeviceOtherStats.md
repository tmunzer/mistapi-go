# DeviceOtherStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigStatus** | Pointer to **string** |  | [optional] 
**LastConfig** | Pointer to **int32** |  | [optional] 
**LastSeen** | Pointer to **int32** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **int32** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**VendorSpecific** | Pointer to [**DeviceOtherStatsVendorSpecific**](DeviceOtherStatsVendorSpecific.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceOtherStats

`func NewDeviceOtherStats() *DeviceOtherStats`

NewDeviceOtherStats instantiates a new DeviceOtherStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceOtherStatsWithDefaults

`func NewDeviceOtherStatsWithDefaults() *DeviceOtherStats`

NewDeviceOtherStatsWithDefaults instantiates a new DeviceOtherStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigStatus

`func (o *DeviceOtherStats) GetConfigStatus() string`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *DeviceOtherStats) GetConfigStatusOk() (*string, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *DeviceOtherStats) SetConfigStatus(v string)`

SetConfigStatus sets ConfigStatus field to given value.

### HasConfigStatus

`func (o *DeviceOtherStats) HasConfigStatus() bool`

HasConfigStatus returns a boolean if a field has been set.

### GetLastConfig

`func (o *DeviceOtherStats) GetLastConfig() int32`

GetLastConfig returns the LastConfig field if non-nil, zero value otherwise.

### GetLastConfigOk

`func (o *DeviceOtherStats) GetLastConfigOk() (*int32, bool)`

GetLastConfigOk returns a tuple with the LastConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfig

`func (o *DeviceOtherStats) SetLastConfig(v int32)`

SetLastConfig sets LastConfig field to given value.

### HasLastConfig

`func (o *DeviceOtherStats) HasLastConfig() bool`

HasLastConfig returns a boolean if a field has been set.

### GetLastSeen

`func (o *DeviceOtherStats) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *DeviceOtherStats) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *DeviceOtherStats) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *DeviceOtherStats) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *DeviceOtherStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceOtherStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceOtherStats) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceOtherStats) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceOtherStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceOtherStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceOtherStats) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceOtherStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUptime

`func (o *DeviceOtherStats) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *DeviceOtherStats) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *DeviceOtherStats) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *DeviceOtherStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVendor

`func (o *DeviceOtherStats) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DeviceOtherStats) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DeviceOtherStats) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DeviceOtherStats) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorSpecific

`func (o *DeviceOtherStats) GetVendorSpecific() DeviceOtherStatsVendorSpecific`

GetVendorSpecific returns the VendorSpecific field if non-nil, zero value otherwise.

### GetVendorSpecificOk

`func (o *DeviceOtherStats) GetVendorSpecificOk() (*DeviceOtherStatsVendorSpecific, bool)`

GetVendorSpecificOk returns a tuple with the VendorSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSpecific

`func (o *DeviceOtherStats) SetVendorSpecific(v DeviceOtherStatsVendorSpecific)`

SetVendorSpecific sets VendorSpecific field to given value.

### HasVendorSpecific

`func (o *DeviceOtherStats) HasVendorSpecific() bool`

HasVendorSpecific returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceOtherStats) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceOtherStats) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceOtherStats) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceOtherStats) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



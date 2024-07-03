# Inventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adopted** | Pointer to **bool** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; whether the switch/gateway is adopted | [optional] 
**Connected** | Pointer to **bool** | whether the device is connected | [optional] 
**CreatedTime** | Pointer to **int32** | inventory created time, in epoch | [optional] 
**DeviceprofileId** | Pointer to **NullableString** | deviceprofile id if assigned, null if not assigned | [optional] 
**Hostname** | Pointer to **string** | hostname reported by the device | [optional] 
**HwRev** | Pointer to **string** | device hardware revision number | [optional] 
**Id** | Pointer to **string** | device id | [optional] 
**Jsi** | Pointer to **bool** |  | [optional] 
**Mac** | Pointer to **string** | device MAC address | [optional] 
**Magic** | Pointer to **string** | device claim code | [optional] 
**Model** | Pointer to **string** | device model | [optional] 
**ModifiedTime** | Pointer to **int32** | inventory last modified time, in epoch | [optional] 
**Name** | Pointer to **string** | device name if configured | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Serial** | Pointer to **string** | device serial | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Sku** | Pointer to **string** | device stock keeping unit | [optional] 
**Type** | Pointer to [**DeviceType**](DeviceType.md) |  | [optional] [default to DEVICETYPE_AP]
**VcMac** | Pointer to **string** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;switch&#x60;, MAC Address of the Virtual Chassis | [optional] 

## Methods

### NewInventory

`func NewInventory() *Inventory`

NewInventory instantiates a new Inventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryWithDefaults

`func NewInventoryWithDefaults() *Inventory`

NewInventoryWithDefaults instantiates a new Inventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdopted

`func (o *Inventory) GetAdopted() bool`

GetAdopted returns the Adopted field if non-nil, zero value otherwise.

### GetAdoptedOk

`func (o *Inventory) GetAdoptedOk() (*bool, bool)`

GetAdoptedOk returns a tuple with the Adopted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdopted

`func (o *Inventory) SetAdopted(v bool)`

SetAdopted sets Adopted field to given value.

### HasAdopted

`func (o *Inventory) HasAdopted() bool`

HasAdopted returns a boolean if a field has been set.

### GetConnected

`func (o *Inventory) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *Inventory) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *Inventory) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *Inventory) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Inventory) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Inventory) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Inventory) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Inventory) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceprofileId

`func (o *Inventory) GetDeviceprofileId() string`

GetDeviceprofileId returns the DeviceprofileId field if non-nil, zero value otherwise.

### GetDeviceprofileIdOk

`func (o *Inventory) GetDeviceprofileIdOk() (*string, bool)`

GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileId

`func (o *Inventory) SetDeviceprofileId(v string)`

SetDeviceprofileId sets DeviceprofileId field to given value.

### HasDeviceprofileId

`func (o *Inventory) HasDeviceprofileId() bool`

HasDeviceprofileId returns a boolean if a field has been set.

### SetDeviceprofileIdNil

`func (o *Inventory) SetDeviceprofileIdNil(b bool)`

 SetDeviceprofileIdNil sets the value for DeviceprofileId to be an explicit nil

### UnsetDeviceprofileId
`func (o *Inventory) UnsetDeviceprofileId()`

UnsetDeviceprofileId ensures that no value is present for DeviceprofileId, not even an explicit nil
### GetHostname

`func (o *Inventory) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Inventory) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Inventory) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Inventory) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHwRev

`func (o *Inventory) GetHwRev() string`

GetHwRev returns the HwRev field if non-nil, zero value otherwise.

### GetHwRevOk

`func (o *Inventory) GetHwRevOk() (*string, bool)`

GetHwRevOk returns a tuple with the HwRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwRev

`func (o *Inventory) SetHwRev(v string)`

SetHwRev sets HwRev field to given value.

### HasHwRev

`func (o *Inventory) HasHwRev() bool`

HasHwRev returns a boolean if a field has been set.

### GetId

`func (o *Inventory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Inventory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Inventory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Inventory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJsi

`func (o *Inventory) GetJsi() bool`

GetJsi returns the Jsi field if non-nil, zero value otherwise.

### GetJsiOk

`func (o *Inventory) GetJsiOk() (*bool, bool)`

GetJsiOk returns a tuple with the Jsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsi

`func (o *Inventory) SetJsi(v bool)`

SetJsi sets Jsi field to given value.

### HasJsi

`func (o *Inventory) HasJsi() bool`

HasJsi returns a boolean if a field has been set.

### GetMac

`func (o *Inventory) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Inventory) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Inventory) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Inventory) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMagic

`func (o *Inventory) GetMagic() string`

GetMagic returns the Magic field if non-nil, zero value otherwise.

### GetMagicOk

`func (o *Inventory) GetMagicOk() (*string, bool)`

GetMagicOk returns a tuple with the Magic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagic

`func (o *Inventory) SetMagic(v string)`

SetMagic sets Magic field to given value.

### HasMagic

`func (o *Inventory) HasMagic() bool`

HasMagic returns a boolean if a field has been set.

### GetModel

`func (o *Inventory) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Inventory) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Inventory) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Inventory) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Inventory) GetModifiedTime() int32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Inventory) GetModifiedTimeOk() (*int32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Inventory) SetModifiedTime(v int32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Inventory) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Inventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Inventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Inventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Inventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Inventory) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Inventory) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Inventory) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Inventory) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSerial

`func (o *Inventory) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Inventory) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Inventory) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Inventory) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteId

`func (o *Inventory) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Inventory) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Inventory) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Inventory) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSku

`func (o *Inventory) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Inventory) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Inventory) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Inventory) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetType

`func (o *Inventory) GetType() DeviceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Inventory) GetTypeOk() (*DeviceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Inventory) SetType(v DeviceType)`

SetType sets Type field to given value.

### HasType

`func (o *Inventory) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVcMac

`func (o *Inventory) GetVcMac() string`

GetVcMac returns the VcMac field if non-nil, zero value otherwise.

### GetVcMacOk

`func (o *Inventory) GetVcMacOk() (*string, bool)`

GetVcMacOk returns a tuple with the VcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcMac

`func (o *Inventory) SetVcMac(v string)`

SetVcMac sets VcMac field to given value.

### HasVcMac

`func (o *Inventory) HasVcMac() bool`

HasVcMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



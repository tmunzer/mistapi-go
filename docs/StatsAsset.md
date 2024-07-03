# StatsAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatteryVoltage** | Pointer to **float32** | battery voltage, in mV | [optional] 
**Beam** | Pointer to **int32** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**EddystoneUidInstance** | Pointer to **string** |  | [optional] 
**EddystoneUidNamespace** | Pointer to **string** |  | [optional] 
**EddystoneUrlUrl** | Pointer to **string** |  | [optional] 
**IbeaconMajor** | Pointer to **int32** |  | [optional] 
**IbeaconMinor** | Pointer to **int32** |  | [optional] 
**IbeaconUuid** | Pointer to **string** |  | [optional] 
**LastSeen** | Pointer to **float32** | last seen timestamp | [optional] 
**Mac** | **string** | bluetooth MAC | 
**MapId** | Pointer to **string** | map where the device belongs to | [optional] 
**Name** | Pointer to **string** | name / label of the device | [optional] 
**Rssi** | Pointer to **int32** |  | [optional] 
**Rssizones** | Pointer to [**[]AssetRssiZone**](AssetRssiZone.md) | only send this for individual asset stat | [optional] 
**Temperatur** | Pointer to **float32** |  | [optional] 
**X** | Pointer to **float32** | x in pixel | [optional] 
**Y** | Pointer to **float32** | y in pixel | [optional] 
**Zones** | Pointer to [**[]AssetZone**](AssetZone.md) | only send this for individual asset stat | [optional] 

## Methods

### NewStatsAsset

`func NewStatsAsset(mac string, ) *StatsAsset`

NewStatsAsset instantiates a new StatsAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsAssetWithDefaults

`func NewStatsAssetWithDefaults() *StatsAsset`

NewStatsAssetWithDefaults instantiates a new StatsAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatteryVoltage

`func (o *StatsAsset) GetBatteryVoltage() float32`

GetBatteryVoltage returns the BatteryVoltage field if non-nil, zero value otherwise.

### GetBatteryVoltageOk

`func (o *StatsAsset) GetBatteryVoltageOk() (*float32, bool)`

GetBatteryVoltageOk returns a tuple with the BatteryVoltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryVoltage

`func (o *StatsAsset) SetBatteryVoltage(v float32)`

SetBatteryVoltage sets BatteryVoltage field to given value.

### HasBatteryVoltage

`func (o *StatsAsset) HasBatteryVoltage() bool`

HasBatteryVoltage returns a boolean if a field has been set.

### GetBeam

`func (o *StatsAsset) GetBeam() int32`

GetBeam returns the Beam field if non-nil, zero value otherwise.

### GetBeamOk

`func (o *StatsAsset) GetBeamOk() (*int32, bool)`

GetBeamOk returns a tuple with the Beam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeam

`func (o *StatsAsset) SetBeam(v int32)`

SetBeam sets Beam field to given value.

### HasBeam

`func (o *StatsAsset) HasBeam() bool`

HasBeam returns a boolean if a field has been set.

### GetDeviceName

`func (o *StatsAsset) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *StatsAsset) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *StatsAsset) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *StatsAsset) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDuration

`func (o *StatsAsset) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *StatsAsset) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *StatsAsset) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *StatsAsset) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEddystoneUidInstance

`func (o *StatsAsset) GetEddystoneUidInstance() string`

GetEddystoneUidInstance returns the EddystoneUidInstance field if non-nil, zero value otherwise.

### GetEddystoneUidInstanceOk

`func (o *StatsAsset) GetEddystoneUidInstanceOk() (*string, bool)`

GetEddystoneUidInstanceOk returns a tuple with the EddystoneUidInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidInstance

`func (o *StatsAsset) SetEddystoneUidInstance(v string)`

SetEddystoneUidInstance sets EddystoneUidInstance field to given value.

### HasEddystoneUidInstance

`func (o *StatsAsset) HasEddystoneUidInstance() bool`

HasEddystoneUidInstance returns a boolean if a field has been set.

### GetEddystoneUidNamespace

`func (o *StatsAsset) GetEddystoneUidNamespace() string`

GetEddystoneUidNamespace returns the EddystoneUidNamespace field if non-nil, zero value otherwise.

### GetEddystoneUidNamespaceOk

`func (o *StatsAsset) GetEddystoneUidNamespaceOk() (*string, bool)`

GetEddystoneUidNamespaceOk returns a tuple with the EddystoneUidNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidNamespace

`func (o *StatsAsset) SetEddystoneUidNamespace(v string)`

SetEddystoneUidNamespace sets EddystoneUidNamespace field to given value.

### HasEddystoneUidNamespace

`func (o *StatsAsset) HasEddystoneUidNamespace() bool`

HasEddystoneUidNamespace returns a boolean if a field has been set.

### GetEddystoneUrlUrl

`func (o *StatsAsset) GetEddystoneUrlUrl() string`

GetEddystoneUrlUrl returns the EddystoneUrlUrl field if non-nil, zero value otherwise.

### GetEddystoneUrlUrlOk

`func (o *StatsAsset) GetEddystoneUrlUrlOk() (*string, bool)`

GetEddystoneUrlUrlOk returns a tuple with the EddystoneUrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlUrl

`func (o *StatsAsset) SetEddystoneUrlUrl(v string)`

SetEddystoneUrlUrl sets EddystoneUrlUrl field to given value.

### HasEddystoneUrlUrl

`func (o *StatsAsset) HasEddystoneUrlUrl() bool`

HasEddystoneUrlUrl returns a boolean if a field has been set.

### GetIbeaconMajor

`func (o *StatsAsset) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *StatsAsset) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *StatsAsset) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *StatsAsset) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconMinor

`func (o *StatsAsset) GetIbeaconMinor() int32`

GetIbeaconMinor returns the IbeaconMinor field if non-nil, zero value otherwise.

### GetIbeaconMinorOk

`func (o *StatsAsset) GetIbeaconMinorOk() (*int32, bool)`

GetIbeaconMinorOk returns a tuple with the IbeaconMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMinor

`func (o *StatsAsset) SetIbeaconMinor(v int32)`

SetIbeaconMinor sets IbeaconMinor field to given value.

### HasIbeaconMinor

`func (o *StatsAsset) HasIbeaconMinor() bool`

HasIbeaconMinor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *StatsAsset) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *StatsAsset) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *StatsAsset) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *StatsAsset) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetLastSeen

`func (o *StatsAsset) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *StatsAsset) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *StatsAsset) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *StatsAsset) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *StatsAsset) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *StatsAsset) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *StatsAsset) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMapId

`func (o *StatsAsset) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *StatsAsset) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *StatsAsset) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *StatsAsset) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetName

`func (o *StatsAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatsAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatsAsset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatsAsset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRssi

`func (o *StatsAsset) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *StatsAsset) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *StatsAsset) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *StatsAsset) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRssizones

`func (o *StatsAsset) GetRssizones() []AssetRssiZone`

GetRssizones returns the Rssizones field if non-nil, zero value otherwise.

### GetRssizonesOk

`func (o *StatsAsset) GetRssizonesOk() (*[]AssetRssiZone, bool)`

GetRssizonesOk returns a tuple with the Rssizones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssizones

`func (o *StatsAsset) SetRssizones(v []AssetRssiZone)`

SetRssizones sets Rssizones field to given value.

### HasRssizones

`func (o *StatsAsset) HasRssizones() bool`

HasRssizones returns a boolean if a field has been set.

### GetTemperatur

`func (o *StatsAsset) GetTemperatur() float32`

GetTemperatur returns the Temperatur field if non-nil, zero value otherwise.

### GetTemperaturOk

`func (o *StatsAsset) GetTemperaturOk() (*float32, bool)`

GetTemperaturOk returns a tuple with the Temperatur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatur

`func (o *StatsAsset) SetTemperatur(v float32)`

SetTemperatur sets Temperatur field to given value.

### HasTemperatur

`func (o *StatsAsset) HasTemperatur() bool`

HasTemperatur returns a boolean if a field has been set.

### GetX

`func (o *StatsAsset) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *StatsAsset) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *StatsAsset) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *StatsAsset) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *StatsAsset) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *StatsAsset) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *StatsAsset) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *StatsAsset) HasY() bool`

HasY returns a boolean if a field has been set.

### GetZones

`func (o *StatsAsset) GetZones() []AssetZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *StatsAsset) GetZonesOk() (*[]AssetZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *StatsAsset) SetZones(v []AssetZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *StatsAsset) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Beacon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**EddystoneInstance** | Pointer to **string** | Eddystone-UID instance (6 bytes) in hexstring format | [optional] 
**EddystoneNamespace** | Pointer to **string** | Eddystone-UID namespace (10 bytes) in hexstring format | [optional] 
**EddystoneUrl** | Pointer to **string** | Eddystone-URL url | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**IbeaconMajor** | Pointer to **int32** | bluetooth tag major | [optional] 
**IbeaconMinor** | Pointer to **int32** | bluetooth tag minor | [optional] 
**IbeaconUuid** | Pointer to **string** | bluetooth tag UUID | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Mac** | Pointer to **string** | optiona, MAC of the beacon, currently used only to identify battery voltage | [optional] 
**MapId** | Pointer to **string** | map where the device belongs to | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** | name / label of the device | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Power** | Pointer to **int32** | in dBm | [optional] [default to -12]
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to [**BeaconType**](BeaconType.md) |  | [optional] [default to BEACONTYPE_EDDYSTONE_UID]
**X** | Pointer to **float32** | x in pixel | [optional] 
**Y** | Pointer to **float32** | y in pixel | [optional] 

## Methods

### NewBeacon

`func NewBeacon() *Beacon`

NewBeacon instantiates a new Beacon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeaconWithDefaults

`func NewBeaconWithDefaults() *Beacon`

NewBeaconWithDefaults instantiates a new Beacon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *Beacon) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Beacon) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Beacon) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Beacon) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEddystoneInstance

`func (o *Beacon) GetEddystoneInstance() string`

GetEddystoneInstance returns the EddystoneInstance field if non-nil, zero value otherwise.

### GetEddystoneInstanceOk

`func (o *Beacon) GetEddystoneInstanceOk() (*string, bool)`

GetEddystoneInstanceOk returns a tuple with the EddystoneInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneInstance

`func (o *Beacon) SetEddystoneInstance(v string)`

SetEddystoneInstance sets EddystoneInstance field to given value.

### HasEddystoneInstance

`func (o *Beacon) HasEddystoneInstance() bool`

HasEddystoneInstance returns a boolean if a field has been set.

### GetEddystoneNamespace

`func (o *Beacon) GetEddystoneNamespace() string`

GetEddystoneNamespace returns the EddystoneNamespace field if non-nil, zero value otherwise.

### GetEddystoneNamespaceOk

`func (o *Beacon) GetEddystoneNamespaceOk() (*string, bool)`

GetEddystoneNamespaceOk returns a tuple with the EddystoneNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneNamespace

`func (o *Beacon) SetEddystoneNamespace(v string)`

SetEddystoneNamespace sets EddystoneNamespace field to given value.

### HasEddystoneNamespace

`func (o *Beacon) HasEddystoneNamespace() bool`

HasEddystoneNamespace returns a boolean if a field has been set.

### GetEddystoneUrl

`func (o *Beacon) GetEddystoneUrl() string`

GetEddystoneUrl returns the EddystoneUrl field if non-nil, zero value otherwise.

### GetEddystoneUrlOk

`func (o *Beacon) GetEddystoneUrlOk() (*string, bool)`

GetEddystoneUrlOk returns a tuple with the EddystoneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrl

`func (o *Beacon) SetEddystoneUrl(v string)`

SetEddystoneUrl sets EddystoneUrl field to given value.

### HasEddystoneUrl

`func (o *Beacon) HasEddystoneUrl() bool`

HasEddystoneUrl returns a boolean if a field has been set.

### GetForSite

`func (o *Beacon) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Beacon) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Beacon) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Beacon) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetIbeaconMajor

`func (o *Beacon) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *Beacon) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *Beacon) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *Beacon) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconMinor

`func (o *Beacon) GetIbeaconMinor() int32`

GetIbeaconMinor returns the IbeaconMinor field if non-nil, zero value otherwise.

### GetIbeaconMinorOk

`func (o *Beacon) GetIbeaconMinorOk() (*int32, bool)`

GetIbeaconMinorOk returns a tuple with the IbeaconMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMinor

`func (o *Beacon) SetIbeaconMinor(v int32)`

SetIbeaconMinor sets IbeaconMinor field to given value.

### HasIbeaconMinor

`func (o *Beacon) HasIbeaconMinor() bool`

HasIbeaconMinor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *Beacon) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *Beacon) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *Beacon) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *Beacon) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetId

`func (o *Beacon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Beacon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Beacon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Beacon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *Beacon) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Beacon) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Beacon) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Beacon) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMapId

`func (o *Beacon) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *Beacon) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *Beacon) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *Beacon) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Beacon) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Beacon) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Beacon) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Beacon) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Beacon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Beacon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Beacon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Beacon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Beacon) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Beacon) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Beacon) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Beacon) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPower

`func (o *Beacon) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *Beacon) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *Beacon) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *Beacon) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetSiteId

`func (o *Beacon) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Beacon) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Beacon) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Beacon) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetType

`func (o *Beacon) GetType() BeaconType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Beacon) GetTypeOk() (*BeaconType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Beacon) SetType(v BeaconType)`

SetType sets Type field to given value.

### HasType

`func (o *Beacon) HasType() bool`

HasType returns a boolean if a field has been set.

### GetX

`func (o *Beacon) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Beacon) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Beacon) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *Beacon) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *Beacon) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Beacon) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Beacon) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *Beacon) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AssetFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **string** |  | [optional] 
**Beam** | Pointer to **int32** |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Disabled** | Pointer to **bool** |  | [optional] [default to false]
**Disasbled** | Pointer to **bool** | whether the asset filter is disabled | [optional] 
**EddystoneUidNamespace** | Pointer to **string** | eddystone uid namespace used to filter assets | [optional] 
**EddystoneUrl** | Pointer to **string** | eddystone url used to filter assets | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**IbeaconMajor** | Pointer to **int32** | ibeacon major value used to filter assets | [optional] 
**IbeaconUuid** | Pointer to **string** | ibeacon uuid used to filter assets | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**MfgCompanyId** | Pointer to **int32** | ble manufacturing-specific company-id used to filter assets | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Rssi** | Pointer to **int32** |  | [optional] 
**ServiceUuid** | Pointer to **string** | ble service data uuid used to filter assets | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAssetFilter

`func NewAssetFilter(name string, ) *AssetFilter`

NewAssetFilter instantiates a new AssetFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetFilterWithDefaults

`func NewAssetFilterWithDefaults() *AssetFilter`

NewAssetFilterWithDefaults instantiates a new AssetFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *AssetFilter) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *AssetFilter) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *AssetFilter) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *AssetFilter) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### GetBeam

`func (o *AssetFilter) GetBeam() int32`

GetBeam returns the Beam field if non-nil, zero value otherwise.

### GetBeamOk

`func (o *AssetFilter) GetBeamOk() (*int32, bool)`

GetBeamOk returns a tuple with the Beam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeam

`func (o *AssetFilter) SetBeam(v int32)`

SetBeam sets Beam field to given value.

### HasBeam

`func (o *AssetFilter) HasBeam() bool`

HasBeam returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AssetFilter) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AssetFilter) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AssetFilter) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AssetFilter) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisabled

`func (o *AssetFilter) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AssetFilter) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AssetFilter) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AssetFilter) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisasbled

`func (o *AssetFilter) GetDisasbled() bool`

GetDisasbled returns the Disasbled field if non-nil, zero value otherwise.

### GetDisasbledOk

`func (o *AssetFilter) GetDisasbledOk() (*bool, bool)`

GetDisasbledOk returns a tuple with the Disasbled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasbled

`func (o *AssetFilter) SetDisasbled(v bool)`

SetDisasbled sets Disasbled field to given value.

### HasDisasbled

`func (o *AssetFilter) HasDisasbled() bool`

HasDisasbled returns a boolean if a field has been set.

### GetEddystoneUidNamespace

`func (o *AssetFilter) GetEddystoneUidNamespace() string`

GetEddystoneUidNamespace returns the EddystoneUidNamespace field if non-nil, zero value otherwise.

### GetEddystoneUidNamespaceOk

`func (o *AssetFilter) GetEddystoneUidNamespaceOk() (*string, bool)`

GetEddystoneUidNamespaceOk returns a tuple with the EddystoneUidNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidNamespace

`func (o *AssetFilter) SetEddystoneUidNamespace(v string)`

SetEddystoneUidNamespace sets EddystoneUidNamespace field to given value.

### HasEddystoneUidNamespace

`func (o *AssetFilter) HasEddystoneUidNamespace() bool`

HasEddystoneUidNamespace returns a boolean if a field has been set.

### GetEddystoneUrl

`func (o *AssetFilter) GetEddystoneUrl() string`

GetEddystoneUrl returns the EddystoneUrl field if non-nil, zero value otherwise.

### GetEddystoneUrlOk

`func (o *AssetFilter) GetEddystoneUrlOk() (*string, bool)`

GetEddystoneUrlOk returns a tuple with the EddystoneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrl

`func (o *AssetFilter) SetEddystoneUrl(v string)`

SetEddystoneUrl sets EddystoneUrl field to given value.

### HasEddystoneUrl

`func (o *AssetFilter) HasEddystoneUrl() bool`

HasEddystoneUrl returns a boolean if a field has been set.

### GetForSite

`func (o *AssetFilter) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *AssetFilter) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *AssetFilter) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *AssetFilter) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetIbeaconMajor

`func (o *AssetFilter) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *AssetFilter) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *AssetFilter) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *AssetFilter) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *AssetFilter) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *AssetFilter) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *AssetFilter) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *AssetFilter) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetId

`func (o *AssetFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetFilter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMfgCompanyId

`func (o *AssetFilter) GetMfgCompanyId() int32`

GetMfgCompanyId returns the MfgCompanyId field if non-nil, zero value otherwise.

### GetMfgCompanyIdOk

`func (o *AssetFilter) GetMfgCompanyIdOk() (*int32, bool)`

GetMfgCompanyIdOk returns a tuple with the MfgCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgCompanyId

`func (o *AssetFilter) SetMfgCompanyId(v int32)`

SetMfgCompanyId sets MfgCompanyId field to given value.

### HasMfgCompanyId

`func (o *AssetFilter) HasMfgCompanyId() bool`

HasMfgCompanyId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *AssetFilter) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *AssetFilter) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *AssetFilter) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *AssetFilter) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *AssetFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetFilter) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *AssetFilter) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AssetFilter) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AssetFilter) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AssetFilter) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRssi

`func (o *AssetFilter) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *AssetFilter) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *AssetFilter) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *AssetFilter) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetServiceUuid

`func (o *AssetFilter) GetServiceUuid() string`

GetServiceUuid returns the ServiceUuid field if non-nil, zero value otherwise.

### GetServiceUuidOk

`func (o *AssetFilter) GetServiceUuidOk() (*string, bool)`

GetServiceUuidOk returns a tuple with the ServiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUuid

`func (o *AssetFilter) SetServiceUuid(v string)`

SetServiceUuid sets ServiceUuid field to given value.

### HasServiceUuid

`func (o *AssetFilter) HasServiceUuid() bool`

HasServiceUuid returns a boolean if a field has been set.

### GetSiteId

`func (o *AssetFilter) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AssetFilter) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AssetFilter) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AssetFilter) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



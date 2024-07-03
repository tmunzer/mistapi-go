# Vbeacon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Major** | Pointer to **int32** | bluetooth tag major | [optional] 
**MapId** | Pointer to **string** | map where the device belongs to | [optional] 
**Message** | Pointer to **string** | a message that can be displayed when the sdkclient gets near the vbeacon | [optional] 
**Minor** | Pointer to **int32** | bluetooth tag minor | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** | name / label of the device | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Power** | Pointer to **int32** | required if &#x60;power_mode&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, -30 - 100, in dBm. For default power_mode, power &#x3D; 4 dBm. | [optional] [default to 4]
**PowerMode** | Pointer to [**BleConfigPowerMode**](BleConfigPowerMode.md) |  | [optional] [default to BLECONFIGPOWERMODE_DEFAULT]
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** | URL to show, optional | [optional] 
**Uuid** | Pointer to **string** | bluetooth tag UUID | [optional] 
**WayfindingNodename** | Pointer to **string** | the name to be used in wayfinding_path or wayfinding_grid blob | [optional] 
**X** | Pointer to **float32** | x in pixel | [optional] 
**Y** | Pointer to **float32** | y in pixel | [optional] 

## Methods

### NewVbeacon

`func NewVbeacon() *Vbeacon`

NewVbeacon instantiates a new Vbeacon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVbeaconWithDefaults

`func NewVbeaconWithDefaults() *Vbeacon`

NewVbeaconWithDefaults instantiates a new Vbeacon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *Vbeacon) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Vbeacon) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Vbeacon) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Vbeacon) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetForSite

`func (o *Vbeacon) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Vbeacon) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Vbeacon) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Vbeacon) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *Vbeacon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vbeacon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vbeacon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Vbeacon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMajor

`func (o *Vbeacon) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *Vbeacon) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *Vbeacon) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *Vbeacon) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMapId

`func (o *Vbeacon) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *Vbeacon) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *Vbeacon) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *Vbeacon) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMessage

`func (o *Vbeacon) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Vbeacon) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Vbeacon) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Vbeacon) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMinor

`func (o *Vbeacon) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *Vbeacon) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *Vbeacon) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *Vbeacon) HasMinor() bool`

HasMinor returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Vbeacon) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Vbeacon) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Vbeacon) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Vbeacon) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Vbeacon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vbeacon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vbeacon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vbeacon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Vbeacon) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Vbeacon) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Vbeacon) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Vbeacon) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPower

`func (o *Vbeacon) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *Vbeacon) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *Vbeacon) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *Vbeacon) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPowerMode

`func (o *Vbeacon) GetPowerMode() BleConfigPowerMode`

GetPowerMode returns the PowerMode field if non-nil, zero value otherwise.

### GetPowerModeOk

`func (o *Vbeacon) GetPowerModeOk() (*BleConfigPowerMode, bool)`

GetPowerModeOk returns a tuple with the PowerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMode

`func (o *Vbeacon) SetPowerMode(v BleConfigPowerMode)`

SetPowerMode sets PowerMode field to given value.

### HasPowerMode

`func (o *Vbeacon) HasPowerMode() bool`

HasPowerMode returns a boolean if a field has been set.

### GetSiteId

`func (o *Vbeacon) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Vbeacon) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Vbeacon) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Vbeacon) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUrl

`func (o *Vbeacon) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Vbeacon) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Vbeacon) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Vbeacon) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUuid

`func (o *Vbeacon) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Vbeacon) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Vbeacon) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Vbeacon) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWayfindingNodename

`func (o *Vbeacon) GetWayfindingNodename() string`

GetWayfindingNodename returns the WayfindingNodename field if non-nil, zero value otherwise.

### GetWayfindingNodenameOk

`func (o *Vbeacon) GetWayfindingNodenameOk() (*string, bool)`

GetWayfindingNodenameOk returns a tuple with the WayfindingNodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayfindingNodename

`func (o *Vbeacon) SetWayfindingNodename(v string)`

SetWayfindingNodename sets WayfindingNodename field to given value.

### HasWayfindingNodename

`func (o *Vbeacon) HasWayfindingNodename() bool`

HasWayfindingNodename returns a boolean if a field has been set.

### GetX

`func (o *Vbeacon) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Vbeacon) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Vbeacon) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *Vbeacon) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *Vbeacon) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Vbeacon) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Vbeacon) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *Vbeacon) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



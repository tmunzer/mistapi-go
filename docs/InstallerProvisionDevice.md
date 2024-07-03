# InstallerProvisionDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceprofileName** | Pointer to **string** |  | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Height** | Pointer to **float32** |  | [optional] 
**MapId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Orientation** | Pointer to **float32** |  | [optional] 
**ReplacingMac** | Pointer to **string** | Onlif this is to replace an existing device | [optional] 
**Role** | Pointer to **string** | optional role for switch / gateway | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float32** |  | [optional] 
**Y** | Pointer to **float32** |  | [optional] 

## Methods

### NewInstallerProvisionDevice

`func NewInstallerProvisionDevice(name string, ) *InstallerProvisionDevice`

NewInstallerProvisionDevice instantiates a new InstallerProvisionDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallerProvisionDeviceWithDefaults

`func NewInstallerProvisionDeviceWithDefaults() *InstallerProvisionDevice`

NewInstallerProvisionDeviceWithDefaults instantiates a new InstallerProvisionDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceprofileName

`func (o *InstallerProvisionDevice) GetDeviceprofileName() string`

GetDeviceprofileName returns the DeviceprofileName field if non-nil, zero value otherwise.

### GetDeviceprofileNameOk

`func (o *InstallerProvisionDevice) GetDeviceprofileNameOk() (*string, bool)`

GetDeviceprofileNameOk returns a tuple with the DeviceprofileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileName

`func (o *InstallerProvisionDevice) SetDeviceprofileName(v string)`

SetDeviceprofileName sets DeviceprofileName field to given value.

### HasDeviceprofileName

`func (o *InstallerProvisionDevice) HasDeviceprofileName() bool`

HasDeviceprofileName returns a boolean if a field has been set.

### GetForSite

`func (o *InstallerProvisionDevice) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *InstallerProvisionDevice) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *InstallerProvisionDevice) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *InstallerProvisionDevice) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHeight

`func (o *InstallerProvisionDevice) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *InstallerProvisionDevice) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *InstallerProvisionDevice) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *InstallerProvisionDevice) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMapId

`func (o *InstallerProvisionDevice) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *InstallerProvisionDevice) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *InstallerProvisionDevice) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *InstallerProvisionDevice) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetName

`func (o *InstallerProvisionDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstallerProvisionDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstallerProvisionDevice) SetName(v string)`

SetName sets Name field to given value.


### GetOrientation

`func (o *InstallerProvisionDevice) GetOrientation() float32`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *InstallerProvisionDevice) GetOrientationOk() (*float32, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *InstallerProvisionDevice) SetOrientation(v float32)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *InstallerProvisionDevice) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetReplacingMac

`func (o *InstallerProvisionDevice) GetReplacingMac() string`

GetReplacingMac returns the ReplacingMac field if non-nil, zero value otherwise.

### GetReplacingMacOk

`func (o *InstallerProvisionDevice) GetReplacingMacOk() (*string, bool)`

GetReplacingMacOk returns a tuple with the ReplacingMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacingMac

`func (o *InstallerProvisionDevice) SetReplacingMac(v string)`

SetReplacingMac sets ReplacingMac field to given value.

### HasReplacingMac

`func (o *InstallerProvisionDevice) HasReplacingMac() bool`

HasReplacingMac returns a boolean if a field has been set.

### GetRole

`func (o *InstallerProvisionDevice) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InstallerProvisionDevice) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InstallerProvisionDevice) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *InstallerProvisionDevice) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSiteId

`func (o *InstallerProvisionDevice) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InstallerProvisionDevice) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InstallerProvisionDevice) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InstallerProvisionDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *InstallerProvisionDevice) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *InstallerProvisionDevice) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *InstallerProvisionDevice) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *InstallerProvisionDevice) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetX

`func (o *InstallerProvisionDevice) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *InstallerProvisionDevice) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *InstallerProvisionDevice) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *InstallerProvisionDevice) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *InstallerProvisionDevice) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *InstallerProvisionDevice) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *InstallerProvisionDevice) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *InstallerProvisionDevice) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



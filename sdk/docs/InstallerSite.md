# InstallerSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**CountryCode** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Latlng** | [**LatLng**](LatLng.md) |  | 
**Name** | **string** |  | 
**RftemplateName** | Pointer to **string** |  | [optional] 
**SitegroupNames** | Pointer to **[]string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewInstallerSite

`func NewInstallerSite(address string, countryCode string, latlng LatLng, name string, ) *InstallerSite`

NewInstallerSite instantiates a new InstallerSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallerSiteWithDefaults

`func NewInstallerSiteWithDefaults() *InstallerSite`

NewInstallerSiteWithDefaults instantiates a new InstallerSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *InstallerSite) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InstallerSite) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InstallerSite) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCountryCode

`func (o *InstallerSite) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *InstallerSite) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *InstallerSite) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetId

`func (o *InstallerSite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstallerSite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstallerSite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstallerSite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatlng

`func (o *InstallerSite) GetLatlng() LatLng`

GetLatlng returns the Latlng field if non-nil, zero value otherwise.

### GetLatlngOk

`func (o *InstallerSite) GetLatlngOk() (*LatLng, bool)`

GetLatlngOk returns a tuple with the Latlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatlng

`func (o *InstallerSite) SetLatlng(v LatLng)`

SetLatlng sets Latlng field to given value.


### GetName

`func (o *InstallerSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstallerSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstallerSite) SetName(v string)`

SetName sets Name field to given value.


### GetRftemplateName

`func (o *InstallerSite) GetRftemplateName() string`

GetRftemplateName returns the RftemplateName field if non-nil, zero value otherwise.

### GetRftemplateNameOk

`func (o *InstallerSite) GetRftemplateNameOk() (*string, bool)`

GetRftemplateNameOk returns a tuple with the RftemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRftemplateName

`func (o *InstallerSite) SetRftemplateName(v string)`

SetRftemplateName sets RftemplateName field to given value.

### HasRftemplateName

`func (o *InstallerSite) HasRftemplateName() bool`

HasRftemplateName returns a boolean if a field has been set.

### GetSitegroupNames

`func (o *InstallerSite) GetSitegroupNames() []string`

GetSitegroupNames returns the SitegroupNames field if non-nil, zero value otherwise.

### GetSitegroupNamesOk

`func (o *InstallerSite) GetSitegroupNamesOk() (*[]string, bool)`

GetSitegroupNamesOk returns a tuple with the SitegroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitegroupNames

`func (o *InstallerSite) SetSitegroupNames(v []string)`

SetSitegroupNames sets SitegroupNames field to given value.

### HasSitegroupNames

`func (o *InstallerSite) HasSitegroupNames() bool`

HasSitegroupNames returns a boolean if a field has been set.

### GetTimezone

`func (o *InstallerSite) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InstallerSite) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InstallerSite) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InstallerSite) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



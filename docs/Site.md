# Site

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | full address of the site | [optional] 
**AlarmtemplateId** | Pointer to **NullableString** | Alarm Template ID, this takes precedence over the Org-level alarmtemplate_id | [optional] 
**AptemplateId** | Pointer to **string** | AP Template ID, used by APs | [optional] 
**CountryCode** | Pointer to **string** | country code for the site (for AP config generation), in two-character | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**GatewaytemplateId** | Pointer to **NullableString** | Gateway Template ID, used by gateways | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Latlng** | Pointer to [**LatLng**](LatLng.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**NetworktemplateId** | Pointer to **NullableString** | Network Template ID, this takes precedence over Site Settings | [optional] 
**Notes** | Pointer to **string** | optional, any notes about the site | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**RftemplateId** | Pointer to **NullableString** | RF Template ID, this takes precedence over Site Settings | [optional] 
**SecpolicyId** | Pointer to **NullableString** | SecPolicy ID | [optional] 
**SitegroupIds** | Pointer to **[]string** | sitegroups this site belongs to | [optional] 
**SitetemplateId** | Pointer to **string** | Site Template ID | [optional] 
**Timezone** | Pointer to **string** | Timezone the site is at | [optional] 

## Methods

### NewSite

`func NewSite(name string, ) *Site`

NewSite instantiates a new Site object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteWithDefaults

`func NewSiteWithDefaults() *Site`

NewSiteWithDefaults instantiates a new Site object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Site) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Site) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Site) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Site) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAlarmtemplateId

`func (o *Site) GetAlarmtemplateId() string`

GetAlarmtemplateId returns the AlarmtemplateId field if non-nil, zero value otherwise.

### GetAlarmtemplateIdOk

`func (o *Site) GetAlarmtemplateIdOk() (*string, bool)`

GetAlarmtemplateIdOk returns a tuple with the AlarmtemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmtemplateId

`func (o *Site) SetAlarmtemplateId(v string)`

SetAlarmtemplateId sets AlarmtemplateId field to given value.

### HasAlarmtemplateId

`func (o *Site) HasAlarmtemplateId() bool`

HasAlarmtemplateId returns a boolean if a field has been set.

### SetAlarmtemplateIdNil

`func (o *Site) SetAlarmtemplateIdNil(b bool)`

 SetAlarmtemplateIdNil sets the value for AlarmtemplateId to be an explicit nil

### UnsetAlarmtemplateId
`func (o *Site) UnsetAlarmtemplateId()`

UnsetAlarmtemplateId ensures that no value is present for AlarmtemplateId, not even an explicit nil
### GetAptemplateId

`func (o *Site) GetAptemplateId() string`

GetAptemplateId returns the AptemplateId field if non-nil, zero value otherwise.

### GetAptemplateIdOk

`func (o *Site) GetAptemplateIdOk() (*string, bool)`

GetAptemplateIdOk returns a tuple with the AptemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptemplateId

`func (o *Site) SetAptemplateId(v string)`

SetAptemplateId sets AptemplateId field to given value.

### HasAptemplateId

`func (o *Site) HasAptemplateId() bool`

HasAptemplateId returns a boolean if a field has been set.

### GetCountryCode

`func (o *Site) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Site) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Site) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Site) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Site) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Site) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Site) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Site) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetGatewaytemplateId

`func (o *Site) GetGatewaytemplateId() string`

GetGatewaytemplateId returns the GatewaytemplateId field if non-nil, zero value otherwise.

### GetGatewaytemplateIdOk

`func (o *Site) GetGatewaytemplateIdOk() (*string, bool)`

GetGatewaytemplateIdOk returns a tuple with the GatewaytemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaytemplateId

`func (o *Site) SetGatewaytemplateId(v string)`

SetGatewaytemplateId sets GatewaytemplateId field to given value.

### HasGatewaytemplateId

`func (o *Site) HasGatewaytemplateId() bool`

HasGatewaytemplateId returns a boolean if a field has been set.

### SetGatewaytemplateIdNil

`func (o *Site) SetGatewaytemplateIdNil(b bool)`

 SetGatewaytemplateIdNil sets the value for GatewaytemplateId to be an explicit nil

### UnsetGatewaytemplateId
`func (o *Site) UnsetGatewaytemplateId()`

UnsetGatewaytemplateId ensures that no value is present for GatewaytemplateId, not even an explicit nil
### GetId

`func (o *Site) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Site) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Site) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Site) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatlng

`func (o *Site) GetLatlng() LatLng`

GetLatlng returns the Latlng field if non-nil, zero value otherwise.

### GetLatlngOk

`func (o *Site) GetLatlngOk() (*LatLng, bool)`

GetLatlngOk returns a tuple with the Latlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatlng

`func (o *Site) SetLatlng(v LatLng)`

SetLatlng sets Latlng field to given value.

### HasLatlng

`func (o *Site) HasLatlng() bool`

HasLatlng returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Site) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Site) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Site) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Site) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Site) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Site) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Site) SetName(v string)`

SetName sets Name field to given value.


### GetNetworktemplateId

`func (o *Site) GetNetworktemplateId() string`

GetNetworktemplateId returns the NetworktemplateId field if non-nil, zero value otherwise.

### GetNetworktemplateIdOk

`func (o *Site) GetNetworktemplateIdOk() (*string, bool)`

GetNetworktemplateIdOk returns a tuple with the NetworktemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworktemplateId

`func (o *Site) SetNetworktemplateId(v string)`

SetNetworktemplateId sets NetworktemplateId field to given value.

### HasNetworktemplateId

`func (o *Site) HasNetworktemplateId() bool`

HasNetworktemplateId returns a boolean if a field has been set.

### SetNetworktemplateIdNil

`func (o *Site) SetNetworktemplateIdNil(b bool)`

 SetNetworktemplateIdNil sets the value for NetworktemplateId to be an explicit nil

### UnsetNetworktemplateId
`func (o *Site) UnsetNetworktemplateId()`

UnsetNetworktemplateId ensures that no value is present for NetworktemplateId, not even an explicit nil
### GetNotes

`func (o *Site) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Site) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Site) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Site) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOrgId

`func (o *Site) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Site) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Site) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Site) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRftemplateId

`func (o *Site) GetRftemplateId() string`

GetRftemplateId returns the RftemplateId field if non-nil, zero value otherwise.

### GetRftemplateIdOk

`func (o *Site) GetRftemplateIdOk() (*string, bool)`

GetRftemplateIdOk returns a tuple with the RftemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRftemplateId

`func (o *Site) SetRftemplateId(v string)`

SetRftemplateId sets RftemplateId field to given value.

### HasRftemplateId

`func (o *Site) HasRftemplateId() bool`

HasRftemplateId returns a boolean if a field has been set.

### SetRftemplateIdNil

`func (o *Site) SetRftemplateIdNil(b bool)`

 SetRftemplateIdNil sets the value for RftemplateId to be an explicit nil

### UnsetRftemplateId
`func (o *Site) UnsetRftemplateId()`

UnsetRftemplateId ensures that no value is present for RftemplateId, not even an explicit nil
### GetSecpolicyId

`func (o *Site) GetSecpolicyId() string`

GetSecpolicyId returns the SecpolicyId field if non-nil, zero value otherwise.

### GetSecpolicyIdOk

`func (o *Site) GetSecpolicyIdOk() (*string, bool)`

GetSecpolicyIdOk returns a tuple with the SecpolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecpolicyId

`func (o *Site) SetSecpolicyId(v string)`

SetSecpolicyId sets SecpolicyId field to given value.

### HasSecpolicyId

`func (o *Site) HasSecpolicyId() bool`

HasSecpolicyId returns a boolean if a field has been set.

### SetSecpolicyIdNil

`func (o *Site) SetSecpolicyIdNil(b bool)`

 SetSecpolicyIdNil sets the value for SecpolicyId to be an explicit nil

### UnsetSecpolicyId
`func (o *Site) UnsetSecpolicyId()`

UnsetSecpolicyId ensures that no value is present for SecpolicyId, not even an explicit nil
### GetSitegroupIds

`func (o *Site) GetSitegroupIds() []string`

GetSitegroupIds returns the SitegroupIds field if non-nil, zero value otherwise.

### GetSitegroupIdsOk

`func (o *Site) GetSitegroupIdsOk() (*[]string, bool)`

GetSitegroupIdsOk returns a tuple with the SitegroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitegroupIds

`func (o *Site) SetSitegroupIds(v []string)`

SetSitegroupIds sets SitegroupIds field to given value.

### HasSitegroupIds

`func (o *Site) HasSitegroupIds() bool`

HasSitegroupIds returns a boolean if a field has been set.

### GetSitetemplateId

`func (o *Site) GetSitetemplateId() string`

GetSitetemplateId returns the SitetemplateId field if non-nil, zero value otherwise.

### GetSitetemplateIdOk

`func (o *Site) GetSitetemplateIdOk() (*string, bool)`

GetSitetemplateIdOk returns a tuple with the SitetemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitetemplateId

`func (o *Site) SetSitetemplateId(v string)`

SetSitetemplateId sets SitetemplateId field to given value.

### HasSitetemplateId

`func (o *Site) HasSitetemplateId() bool`

HasSitetemplateId returns a boolean if a field has been set.

### GetTimezone

`func (o *Site) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Site) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Site) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Site) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



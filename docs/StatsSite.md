# StatsSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**AlarmtemplateId** | **string** |  | 
**CountryCode** | **string** |  | 
**CreatedTime** | **float32** |  | 
**Id** | **string** |  | 
**Lat** | **float32** |  | 
**Latlng** | [**LatLng**](LatLng.md) |  | 
**Lng** | **float32** |  | 
**ModifiedTime** | **float32** |  | 
**MspId** | **string** |  | 
**Name** | **string** |  | 
**NetworktemplateId** | **string** |  | 
**NumAp** | **int32** |  | 
**NumApConnected** | **int32** |  | 
**NumClients** | **int32** |  | 
**NumDevices** | **int32** |  | 
**NumDevicesConnected** | **int32** |  | 
**NumGateway** | **int32** |  | 
**NumGatewayConnected** | **int32** |  | 
**NumSwitch** | **int32** |  | 
**NumSwitchConnected** | **int32** |  | 
**OrgId** | **string** |  | [readonly] 
**RftemplateId** | **string** |  | 
**SecpolicyId** | Pointer to **interface{}** |  | [optional] 
**SitegroupIds** | **[]string** |  | 
**Timezone** | **string** |  | 
**Tzoffset** | **int32** |  | 

## Methods

### NewStatsSite

`func NewStatsSite(address string, alarmtemplateId string, countryCode string, createdTime float32, id string, lat float32, latlng LatLng, lng float32, modifiedTime float32, mspId string, name string, networktemplateId string, numAp int32, numApConnected int32, numClients int32, numDevices int32, numDevicesConnected int32, numGateway int32, numGatewayConnected int32, numSwitch int32, numSwitchConnected int32, orgId string, rftemplateId string, sitegroupIds []string, timezone string, tzoffset int32, ) *StatsSite`

NewStatsSite instantiates a new StatsSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsSiteWithDefaults

`func NewStatsSiteWithDefaults() *StatsSite`

NewStatsSiteWithDefaults instantiates a new StatsSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *StatsSite) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StatsSite) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StatsSite) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAlarmtemplateId

`func (o *StatsSite) GetAlarmtemplateId() string`

GetAlarmtemplateId returns the AlarmtemplateId field if non-nil, zero value otherwise.

### GetAlarmtemplateIdOk

`func (o *StatsSite) GetAlarmtemplateIdOk() (*string, bool)`

GetAlarmtemplateIdOk returns a tuple with the AlarmtemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmtemplateId

`func (o *StatsSite) SetAlarmtemplateId(v string)`

SetAlarmtemplateId sets AlarmtemplateId field to given value.


### GetCountryCode

`func (o *StatsSite) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *StatsSite) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *StatsSite) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCreatedTime

`func (o *StatsSite) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StatsSite) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StatsSite) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetId

`func (o *StatsSite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatsSite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatsSite) SetId(v string)`

SetId sets Id field to given value.


### GetLat

`func (o *StatsSite) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *StatsSite) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *StatsSite) SetLat(v float32)`

SetLat sets Lat field to given value.


### GetLatlng

`func (o *StatsSite) GetLatlng() LatLng`

GetLatlng returns the Latlng field if non-nil, zero value otherwise.

### GetLatlngOk

`func (o *StatsSite) GetLatlngOk() (*LatLng, bool)`

GetLatlngOk returns a tuple with the Latlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatlng

`func (o *StatsSite) SetLatlng(v LatLng)`

SetLatlng sets Latlng field to given value.


### GetLng

`func (o *StatsSite) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *StatsSite) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *StatsSite) SetLng(v float32)`

SetLng sets Lng field to given value.


### GetModifiedTime

`func (o *StatsSite) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *StatsSite) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *StatsSite) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.


### GetMspId

`func (o *StatsSite) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *StatsSite) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *StatsSite) SetMspId(v string)`

SetMspId sets MspId field to given value.


### GetName

`func (o *StatsSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatsSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatsSite) SetName(v string)`

SetName sets Name field to given value.


### GetNetworktemplateId

`func (o *StatsSite) GetNetworktemplateId() string`

GetNetworktemplateId returns the NetworktemplateId field if non-nil, zero value otherwise.

### GetNetworktemplateIdOk

`func (o *StatsSite) GetNetworktemplateIdOk() (*string, bool)`

GetNetworktemplateIdOk returns a tuple with the NetworktemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworktemplateId

`func (o *StatsSite) SetNetworktemplateId(v string)`

SetNetworktemplateId sets NetworktemplateId field to given value.


### GetNumAp

`func (o *StatsSite) GetNumAp() int32`

GetNumAp returns the NumAp field if non-nil, zero value otherwise.

### GetNumApOk

`func (o *StatsSite) GetNumApOk() (*int32, bool)`

GetNumApOk returns a tuple with the NumAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAp

`func (o *StatsSite) SetNumAp(v int32)`

SetNumAp sets NumAp field to given value.


### GetNumApConnected

`func (o *StatsSite) GetNumApConnected() int32`

GetNumApConnected returns the NumApConnected field if non-nil, zero value otherwise.

### GetNumApConnectedOk

`func (o *StatsSite) GetNumApConnectedOk() (*int32, bool)`

GetNumApConnectedOk returns a tuple with the NumApConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApConnected

`func (o *StatsSite) SetNumApConnected(v int32)`

SetNumApConnected sets NumApConnected field to given value.


### GetNumClients

`func (o *StatsSite) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *StatsSite) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *StatsSite) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.


### GetNumDevices

`func (o *StatsSite) GetNumDevices() int32`

GetNumDevices returns the NumDevices field if non-nil, zero value otherwise.

### GetNumDevicesOk

`func (o *StatsSite) GetNumDevicesOk() (*int32, bool)`

GetNumDevicesOk returns a tuple with the NumDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDevices

`func (o *StatsSite) SetNumDevices(v int32)`

SetNumDevices sets NumDevices field to given value.


### GetNumDevicesConnected

`func (o *StatsSite) GetNumDevicesConnected() int32`

GetNumDevicesConnected returns the NumDevicesConnected field if non-nil, zero value otherwise.

### GetNumDevicesConnectedOk

`func (o *StatsSite) GetNumDevicesConnectedOk() (*int32, bool)`

GetNumDevicesConnectedOk returns a tuple with the NumDevicesConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDevicesConnected

`func (o *StatsSite) SetNumDevicesConnected(v int32)`

SetNumDevicesConnected sets NumDevicesConnected field to given value.


### GetNumGateway

`func (o *StatsSite) GetNumGateway() int32`

GetNumGateway returns the NumGateway field if non-nil, zero value otherwise.

### GetNumGatewayOk

`func (o *StatsSite) GetNumGatewayOk() (*int32, bool)`

GetNumGatewayOk returns a tuple with the NumGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGateway

`func (o *StatsSite) SetNumGateway(v int32)`

SetNumGateway sets NumGateway field to given value.


### GetNumGatewayConnected

`func (o *StatsSite) GetNumGatewayConnected() int32`

GetNumGatewayConnected returns the NumGatewayConnected field if non-nil, zero value otherwise.

### GetNumGatewayConnectedOk

`func (o *StatsSite) GetNumGatewayConnectedOk() (*int32, bool)`

GetNumGatewayConnectedOk returns a tuple with the NumGatewayConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGatewayConnected

`func (o *StatsSite) SetNumGatewayConnected(v int32)`

SetNumGatewayConnected sets NumGatewayConnected field to given value.


### GetNumSwitch

`func (o *StatsSite) GetNumSwitch() int32`

GetNumSwitch returns the NumSwitch field if non-nil, zero value otherwise.

### GetNumSwitchOk

`func (o *StatsSite) GetNumSwitchOk() (*int32, bool)`

GetNumSwitchOk returns a tuple with the NumSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSwitch

`func (o *StatsSite) SetNumSwitch(v int32)`

SetNumSwitch sets NumSwitch field to given value.


### GetNumSwitchConnected

`func (o *StatsSite) GetNumSwitchConnected() int32`

GetNumSwitchConnected returns the NumSwitchConnected field if non-nil, zero value otherwise.

### GetNumSwitchConnectedOk

`func (o *StatsSite) GetNumSwitchConnectedOk() (*int32, bool)`

GetNumSwitchConnectedOk returns a tuple with the NumSwitchConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSwitchConnected

`func (o *StatsSite) SetNumSwitchConnected(v int32)`

SetNumSwitchConnected sets NumSwitchConnected field to given value.


### GetOrgId

`func (o *StatsSite) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *StatsSite) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *StatsSite) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRftemplateId

`func (o *StatsSite) GetRftemplateId() string`

GetRftemplateId returns the RftemplateId field if non-nil, zero value otherwise.

### GetRftemplateIdOk

`func (o *StatsSite) GetRftemplateIdOk() (*string, bool)`

GetRftemplateIdOk returns a tuple with the RftemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRftemplateId

`func (o *StatsSite) SetRftemplateId(v string)`

SetRftemplateId sets RftemplateId field to given value.


### GetSecpolicyId

`func (o *StatsSite) GetSecpolicyId() interface{}`

GetSecpolicyId returns the SecpolicyId field if non-nil, zero value otherwise.

### GetSecpolicyIdOk

`func (o *StatsSite) GetSecpolicyIdOk() (*interface{}, bool)`

GetSecpolicyIdOk returns a tuple with the SecpolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecpolicyId

`func (o *StatsSite) SetSecpolicyId(v interface{})`

SetSecpolicyId sets SecpolicyId field to given value.

### HasSecpolicyId

`func (o *StatsSite) HasSecpolicyId() bool`

HasSecpolicyId returns a boolean if a field has been set.

### SetSecpolicyIdNil

`func (o *StatsSite) SetSecpolicyIdNil(b bool)`

 SetSecpolicyIdNil sets the value for SecpolicyId to be an explicit nil

### UnsetSecpolicyId
`func (o *StatsSite) UnsetSecpolicyId()`

UnsetSecpolicyId ensures that no value is present for SecpolicyId, not even an explicit nil
### GetSitegroupIds

`func (o *StatsSite) GetSitegroupIds() []string`

GetSitegroupIds returns the SitegroupIds field if non-nil, zero value otherwise.

### GetSitegroupIdsOk

`func (o *StatsSite) GetSitegroupIdsOk() (*[]string, bool)`

GetSitegroupIdsOk returns a tuple with the SitegroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitegroupIds

`func (o *StatsSite) SetSitegroupIds(v []string)`

SetSitegroupIds sets SitegroupIds field to given value.


### GetTimezone

`func (o *StatsSite) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *StatsSite) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *StatsSite) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetTzoffset

`func (o *StatsSite) GetTzoffset() int32`

GetTzoffset returns the Tzoffset field if non-nil, zero value otherwise.

### GetTzoffsetOk

`func (o *StatsSite) GetTzoffsetOk() (*int32, bool)`

GetTzoffsetOk returns a tuple with the Tzoffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTzoffset

`func (o *StatsSite) SetTzoffset(v int32)`

SetTzoffset sets Tzoffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



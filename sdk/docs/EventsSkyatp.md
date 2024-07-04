# EventsSkyatp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | **string** |  | [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Ip** | **string** |  | [readonly] 
**Mac** | **string** |  | [readonly] 
**OrgId** | **string** |  | [readonly] 
**SiteId** | **string** |  | [readonly] 
**ThreatLevel** | **int32** |  | [readonly] 
**Timestamp** | **float32** |  | [readonly] 
**Type** | **string** |  | [readonly] 

## Methods

### NewEventsSkyatp

`func NewEventsSkyatp(deviceMac string, ip string, mac string, orgId string, siteId string, threatLevel int32, timestamp float32, type_ string, ) *EventsSkyatp`

NewEventsSkyatp instantiates a new EventsSkyatp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsSkyatpWithDefaults

`func NewEventsSkyatpWithDefaults() *EventsSkyatp`

NewEventsSkyatpWithDefaults instantiates a new EventsSkyatp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *EventsSkyatp) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *EventsSkyatp) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *EventsSkyatp) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.


### GetForSite

`func (o *EventsSkyatp) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *EventsSkyatp) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *EventsSkyatp) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *EventsSkyatp) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetIp

`func (o *EventsSkyatp) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *EventsSkyatp) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *EventsSkyatp) SetIp(v string)`

SetIp sets Ip field to given value.


### GetMac

`func (o *EventsSkyatp) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *EventsSkyatp) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *EventsSkyatp) SetMac(v string)`

SetMac sets Mac field to given value.


### GetOrgId

`func (o *EventsSkyatp) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EventsSkyatp) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EventsSkyatp) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSiteId

`func (o *EventsSkyatp) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *EventsSkyatp) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *EventsSkyatp) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetThreatLevel

`func (o *EventsSkyatp) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *EventsSkyatp) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *EventsSkyatp) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.


### GetTimestamp

`func (o *EventsSkyatp) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventsSkyatp) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventsSkyatp) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *EventsSkyatp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventsSkyatp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventsSkyatp) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



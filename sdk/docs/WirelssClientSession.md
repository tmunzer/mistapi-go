# WirelssClientSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** |  | [readonly] 
**Band** | **string** |  | [readonly] 
**ClientManufacture** | **NullableString** |  | [readonly] 
**Connect** | **int32** |  | [readonly] 
**Disconnect** | **int32** |  | [readonly] 
**Duration** | **float32** |  | [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Mac** | **string** |  | [readonly] 
**OrgId** | **string** |  | [readonly] 
**SiteId** | **string** |  | [readonly] 
**Ssid** | **string** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] [readonly] 
**Timestamp** | **float32** |  | [readonly] 
**WlanId** | **string** |  | [readonly] 

## Methods

### NewWirelssClientSession

`func NewWirelssClientSession(ap string, band string, clientManufacture NullableString, connect int32, disconnect int32, duration float32, mac string, orgId string, siteId string, ssid string, timestamp float32, wlanId string, ) *WirelssClientSession`

NewWirelssClientSession instantiates a new WirelssClientSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelssClientSessionWithDefaults

`func NewWirelssClientSessionWithDefaults() *WirelssClientSession`

NewWirelssClientSessionWithDefaults instantiates a new WirelssClientSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *WirelssClientSession) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *WirelssClientSession) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *WirelssClientSession) SetAp(v string)`

SetAp sets Ap field to given value.


### GetBand

`func (o *WirelssClientSession) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *WirelssClientSession) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *WirelssClientSession) SetBand(v string)`

SetBand sets Band field to given value.


### GetClientManufacture

`func (o *WirelssClientSession) GetClientManufacture() string`

GetClientManufacture returns the ClientManufacture field if non-nil, zero value otherwise.

### GetClientManufactureOk

`func (o *WirelssClientSession) GetClientManufactureOk() (*string, bool)`

GetClientManufactureOk returns a tuple with the ClientManufacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManufacture

`func (o *WirelssClientSession) SetClientManufacture(v string)`

SetClientManufacture sets ClientManufacture field to given value.


### SetClientManufactureNil

`func (o *WirelssClientSession) SetClientManufactureNil(b bool)`

 SetClientManufactureNil sets the value for ClientManufacture to be an explicit nil

### UnsetClientManufacture
`func (o *WirelssClientSession) UnsetClientManufacture()`

UnsetClientManufacture ensures that no value is present for ClientManufacture, not even an explicit nil
### GetConnect

`func (o *WirelssClientSession) GetConnect() int32`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *WirelssClientSession) GetConnectOk() (*int32, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *WirelssClientSession) SetConnect(v int32)`

SetConnect sets Connect field to given value.


### GetDisconnect

`func (o *WirelssClientSession) GetDisconnect() int32`

GetDisconnect returns the Disconnect field if non-nil, zero value otherwise.

### GetDisconnectOk

`func (o *WirelssClientSession) GetDisconnectOk() (*int32, bool)`

GetDisconnectOk returns a tuple with the Disconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnect

`func (o *WirelssClientSession) SetDisconnect(v int32)`

SetDisconnect sets Disconnect field to given value.


### GetDuration

`func (o *WirelssClientSession) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WirelssClientSession) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WirelssClientSession) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetForSite

`func (o *WirelssClientSession) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *WirelssClientSession) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *WirelssClientSession) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *WirelssClientSession) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetMac

`func (o *WirelssClientSession) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WirelssClientSession) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WirelssClientSession) SetMac(v string)`

SetMac sets Mac field to given value.


### GetOrgId

`func (o *WirelssClientSession) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WirelssClientSession) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WirelssClientSession) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSiteId

`func (o *WirelssClientSession) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WirelssClientSession) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WirelssClientSession) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSsid

`func (o *WirelssClientSession) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WirelssClientSession) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WirelssClientSession) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetTags

`func (o *WirelssClientSession) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WirelssClientSession) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WirelssClientSession) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WirelssClientSession) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *WirelssClientSession) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WirelssClientSession) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WirelssClientSession) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetWlanId

`func (o *WirelssClientSession) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *WirelssClientSession) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *WirelssClientSession) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



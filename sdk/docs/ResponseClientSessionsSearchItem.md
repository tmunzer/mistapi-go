# ResponseClientSessionsSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** |  | [readonly] 
**Band** | **string** |  | [readonly] 
**ClientManufacture** | **string** |  | [readonly] 
**Connect** | **float32** |  | [readonly] 
**Disconnect** | **float32** |  | [readonly] 
**Duration** | **float32** |  | [readonly] 
**Mac** | **string** |  | [readonly] 
**OrgId** | **string** |  | [readonly] 
**SiteId** | **string** |  | [readonly] 
**Ssid** | **string** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] [readonly] 
**Timestamp** | **float32** |  | 
**WlanId** | **string** |  | 

## Methods

### NewResponseClientSessionsSearchItem

`func NewResponseClientSessionsSearchItem(ap string, band string, clientManufacture string, connect float32, disconnect float32, duration float32, mac string, orgId string, siteId string, ssid string, timestamp float32, wlanId string, ) *ResponseClientSessionsSearchItem`

NewResponseClientSessionsSearchItem instantiates a new ResponseClientSessionsSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseClientSessionsSearchItemWithDefaults

`func NewResponseClientSessionsSearchItemWithDefaults() *ResponseClientSessionsSearchItem`

NewResponseClientSessionsSearchItemWithDefaults instantiates a new ResponseClientSessionsSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *ResponseClientSessionsSearchItem) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *ResponseClientSessionsSearchItem) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *ResponseClientSessionsSearchItem) SetAp(v string)`

SetAp sets Ap field to given value.


### GetBand

`func (o *ResponseClientSessionsSearchItem) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *ResponseClientSessionsSearchItem) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *ResponseClientSessionsSearchItem) SetBand(v string)`

SetBand sets Band field to given value.


### GetClientManufacture

`func (o *ResponseClientSessionsSearchItem) GetClientManufacture() string`

GetClientManufacture returns the ClientManufacture field if non-nil, zero value otherwise.

### GetClientManufactureOk

`func (o *ResponseClientSessionsSearchItem) GetClientManufactureOk() (*string, bool)`

GetClientManufactureOk returns a tuple with the ClientManufacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManufacture

`func (o *ResponseClientSessionsSearchItem) SetClientManufacture(v string)`

SetClientManufacture sets ClientManufacture field to given value.


### GetConnect

`func (o *ResponseClientSessionsSearchItem) GetConnect() float32`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *ResponseClientSessionsSearchItem) GetConnectOk() (*float32, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *ResponseClientSessionsSearchItem) SetConnect(v float32)`

SetConnect sets Connect field to given value.


### GetDisconnect

`func (o *ResponseClientSessionsSearchItem) GetDisconnect() float32`

GetDisconnect returns the Disconnect field if non-nil, zero value otherwise.

### GetDisconnectOk

`func (o *ResponseClientSessionsSearchItem) GetDisconnectOk() (*float32, bool)`

GetDisconnectOk returns a tuple with the Disconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnect

`func (o *ResponseClientSessionsSearchItem) SetDisconnect(v float32)`

SetDisconnect sets Disconnect field to given value.


### GetDuration

`func (o *ResponseClientSessionsSearchItem) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ResponseClientSessionsSearchItem) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ResponseClientSessionsSearchItem) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetMac

`func (o *ResponseClientSessionsSearchItem) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ResponseClientSessionsSearchItem) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ResponseClientSessionsSearchItem) SetMac(v string)`

SetMac sets Mac field to given value.


### GetOrgId

`func (o *ResponseClientSessionsSearchItem) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResponseClientSessionsSearchItem) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResponseClientSessionsSearchItem) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSiteId

`func (o *ResponseClientSessionsSearchItem) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponseClientSessionsSearchItem) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponseClientSessionsSearchItem) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSsid

`func (o *ResponseClientSessionsSearchItem) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ResponseClientSessionsSearchItem) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ResponseClientSessionsSearchItem) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetTags

`func (o *ResponseClientSessionsSearchItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponseClientSessionsSearchItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponseClientSessionsSearchItem) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResponseClientSessionsSearchItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *ResponseClientSessionsSearchItem) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseClientSessionsSearchItem) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseClientSessionsSearchItem) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetWlanId

`func (o *ResponseClientSessionsSearchItem) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *ResponseClientSessionsSearchItem) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *ResponseClientSessionsSearchItem) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



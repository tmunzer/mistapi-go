# EventsDeviceAp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **string** |  | [optional] 
**Apfw** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PortId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Text** | Pointer to **string** |  | [optional] 
**Timestamp** | **float32** |  | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsDeviceAp

`func NewEventsDeviceAp(timestamp float32, ) *EventsDeviceAp`

NewEventsDeviceAp instantiates a new EventsDeviceAp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsDeviceApWithDefaults

`func NewEventsDeviceApWithDefaults() *EventsDeviceAp`

NewEventsDeviceApWithDefaults instantiates a new EventsDeviceAp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *EventsDeviceAp) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *EventsDeviceAp) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *EventsDeviceAp) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *EventsDeviceAp) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetApfw

`func (o *EventsDeviceAp) GetApfw() string`

GetApfw returns the Apfw field if non-nil, zero value otherwise.

### GetApfwOk

`func (o *EventsDeviceAp) GetApfwOk() (*string, bool)`

GetApfwOk returns a tuple with the Apfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApfw

`func (o *EventsDeviceAp) SetApfw(v string)`

SetApfw sets Apfw field to given value.

### HasApfw

`func (o *EventsDeviceAp) HasApfw() bool`

HasApfw returns a boolean if a field has been set.

### GetCount

`func (o *EventsDeviceAp) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventsDeviceAp) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventsDeviceAp) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *EventsDeviceAp) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDeviceType

`func (o *EventsDeviceAp) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *EventsDeviceAp) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *EventsDeviceAp) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *EventsDeviceAp) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetMac

`func (o *EventsDeviceAp) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *EventsDeviceAp) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *EventsDeviceAp) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *EventsDeviceAp) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetOrgId

`func (o *EventsDeviceAp) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EventsDeviceAp) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EventsDeviceAp) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *EventsDeviceAp) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPortId

`func (o *EventsDeviceAp) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *EventsDeviceAp) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *EventsDeviceAp) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *EventsDeviceAp) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSiteId

`func (o *EventsDeviceAp) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *EventsDeviceAp) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *EventsDeviceAp) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *EventsDeviceAp) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetText

`func (o *EventsDeviceAp) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EventsDeviceAp) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EventsDeviceAp) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EventsDeviceAp) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTimestamp

`func (o *EventsDeviceAp) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventsDeviceAp) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventsDeviceAp) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *EventsDeviceAp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventsDeviceAp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventsDeviceAp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventsDeviceAp) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



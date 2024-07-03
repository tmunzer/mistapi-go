# ResponsePcapSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMacs** | Pointer to **[]string** |  | [optional] 
**Aps** | Pointer to **[]string** |  | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxNumPackets** | Pointer to **float32** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PcapAps** | Pointer to [**map[string]ResponsePcapSearchItemPcapApsItem**](ResponsePcapSearchItemPcapApsItem.md) |  | [optional] 
**PcapUrl** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**TerminationReason** | Pointer to **string** |  | [optional] 
**Timestamp** | **float32** |  | 
**Type** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewResponsePcapSearchItem

`func NewResponsePcapSearchItem(timestamp float32, type_ string, url string, ) *ResponsePcapSearchItem`

NewResponsePcapSearchItem instantiates a new ResponsePcapSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePcapSearchItemWithDefaults

`func NewResponsePcapSearchItemWithDefaults() *ResponsePcapSearchItem`

NewResponsePcapSearchItemWithDefaults instantiates a new ResponsePcapSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMacs

`func (o *ResponsePcapSearchItem) GetApMacs() []string`

GetApMacs returns the ApMacs field if non-nil, zero value otherwise.

### GetApMacsOk

`func (o *ResponsePcapSearchItem) GetApMacsOk() (*[]string, bool)`

GetApMacsOk returns a tuple with the ApMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMacs

`func (o *ResponsePcapSearchItem) SetApMacs(v []string)`

SetApMacs sets ApMacs field to given value.

### HasApMacs

`func (o *ResponsePcapSearchItem) HasApMacs() bool`

HasApMacs returns a boolean if a field has been set.

### GetAps

`func (o *ResponsePcapSearchItem) GetAps() []string`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *ResponsePcapSearchItem) GetApsOk() (*[]string, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *ResponsePcapSearchItem) SetAps(v []string)`

SetAps sets Aps field to given value.

### HasAps

`func (o *ResponsePcapSearchItem) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetDuration

`func (o *ResponsePcapSearchItem) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ResponsePcapSearchItem) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ResponsePcapSearchItem) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ResponsePcapSearchItem) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFormat

`func (o *ResponsePcapSearchItem) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ResponsePcapSearchItem) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ResponsePcapSearchItem) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ResponsePcapSearchItem) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetId

`func (o *ResponsePcapSearchItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePcapSearchItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePcapSearchItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsePcapSearchItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxNumPackets

`func (o *ResponsePcapSearchItem) GetMaxNumPackets() float32`

GetMaxNumPackets returns the MaxNumPackets field if non-nil, zero value otherwise.

### GetMaxNumPacketsOk

`func (o *ResponsePcapSearchItem) GetMaxNumPacketsOk() (*float32, bool)`

GetMaxNumPacketsOk returns a tuple with the MaxNumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumPackets

`func (o *ResponsePcapSearchItem) SetMaxNumPackets(v float32)`

SetMaxNumPackets sets MaxNumPackets field to given value.

### HasMaxNumPackets

`func (o *ResponsePcapSearchItem) HasMaxNumPackets() bool`

HasMaxNumPackets returns a boolean if a field has been set.

### GetOrgId

`func (o *ResponsePcapSearchItem) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResponsePcapSearchItem) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResponsePcapSearchItem) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ResponsePcapSearchItem) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPcapAps

`func (o *ResponsePcapSearchItem) GetPcapAps() map[string]ResponsePcapSearchItemPcapApsItem`

GetPcapAps returns the PcapAps field if non-nil, zero value otherwise.

### GetPcapApsOk

`func (o *ResponsePcapSearchItem) GetPcapApsOk() (*map[string]ResponsePcapSearchItemPcapApsItem, bool)`

GetPcapApsOk returns a tuple with the PcapAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapAps

`func (o *ResponsePcapSearchItem) SetPcapAps(v map[string]ResponsePcapSearchItemPcapApsItem)`

SetPcapAps sets PcapAps field to given value.

### HasPcapAps

`func (o *ResponsePcapSearchItem) HasPcapAps() bool`

HasPcapAps returns a boolean if a field has been set.

### GetPcapUrl

`func (o *ResponsePcapSearchItem) GetPcapUrl() string`

GetPcapUrl returns the PcapUrl field if non-nil, zero value otherwise.

### GetPcapUrlOk

`func (o *ResponsePcapSearchItem) GetPcapUrlOk() (*string, bool)`

GetPcapUrlOk returns a tuple with the PcapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapUrl

`func (o *ResponsePcapSearchItem) SetPcapUrl(v string)`

SetPcapUrl sets PcapUrl field to given value.

### HasPcapUrl

`func (o *ResponsePcapSearchItem) HasPcapUrl() bool`

HasPcapUrl returns a boolean if a field has been set.

### GetSiteId

`func (o *ResponsePcapSearchItem) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponsePcapSearchItem) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponsePcapSearchItem) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ResponsePcapSearchItem) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTerminationReason

`func (o *ResponsePcapSearchItem) GetTerminationReason() string`

GetTerminationReason returns the TerminationReason field if non-nil, zero value otherwise.

### GetTerminationReasonOk

`func (o *ResponsePcapSearchItem) GetTerminationReasonOk() (*string, bool)`

GetTerminationReasonOk returns a tuple with the TerminationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationReason

`func (o *ResponsePcapSearchItem) SetTerminationReason(v string)`

SetTerminationReason sets TerminationReason field to given value.

### HasTerminationReason

`func (o *ResponsePcapSearchItem) HasTerminationReason() bool`

HasTerminationReason returns a boolean if a field has been set.

### GetTimestamp

`func (o *ResponsePcapSearchItem) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponsePcapSearchItem) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponsePcapSearchItem) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *ResponsePcapSearchItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePcapSearchItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePcapSearchItem) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *ResponsePcapSearchItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResponsePcapSearchItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResponsePcapSearchItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



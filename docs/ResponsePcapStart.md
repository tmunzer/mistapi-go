# ResponsePcapStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApCount** | Pointer to **int32** |  | [optional] 
**Aps** | Pointer to **[]string** |  | [optional] 
**ClientMac** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Expiry** | Pointer to **float32** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**IncludeMcast** | Pointer to **bool** |  | [optional] 
**MaxPktLen** | Pointer to **int32** |  | [optional] 
**NumPackets** | Pointer to **int32** |  | [optional] 
**OrgId** | **string** |  | [readonly] 
**Raw** | Pointer to **bool** |  | [optional] 
**SiteId** | **string** |  | [readonly] 
**Ssid** | Pointer to **NullableString** |  | [optional] 
**TcpdumpParserExpression** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | **float32** |  | 
**Type** | **string** |  | 

## Methods

### NewResponsePcapStart

`func NewResponsePcapStart(id string, orgId string, siteId string, timestamp float32, type_ string, ) *ResponsePcapStart`

NewResponsePcapStart instantiates a new ResponsePcapStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePcapStartWithDefaults

`func NewResponsePcapStartWithDefaults() *ResponsePcapStart`

NewResponsePcapStartWithDefaults instantiates a new ResponsePcapStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApCount

`func (o *ResponsePcapStart) GetApCount() int32`

GetApCount returns the ApCount field if non-nil, zero value otherwise.

### GetApCountOk

`func (o *ResponsePcapStart) GetApCountOk() (*int32, bool)`

GetApCountOk returns a tuple with the ApCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApCount

`func (o *ResponsePcapStart) SetApCount(v int32)`

SetApCount sets ApCount field to given value.

### HasApCount

`func (o *ResponsePcapStart) HasApCount() bool`

HasApCount returns a boolean if a field has been set.

### GetAps

`func (o *ResponsePcapStart) GetAps() []string`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *ResponsePcapStart) GetApsOk() (*[]string, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *ResponsePcapStart) SetAps(v []string)`

SetAps sets Aps field to given value.

### HasAps

`func (o *ResponsePcapStart) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetClientMac

`func (o *ResponsePcapStart) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *ResponsePcapStart) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *ResponsePcapStart) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *ResponsePcapStart) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### SetClientMacNil

`func (o *ResponsePcapStart) SetClientMacNil(b bool)`

 SetClientMacNil sets the value for ClientMac to be an explicit nil

### UnsetClientMac
`func (o *ResponsePcapStart) UnsetClientMac()`

UnsetClientMac ensures that no value is present for ClientMac, not even an explicit nil
### GetDuration

`func (o *ResponsePcapStart) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ResponsePcapStart) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ResponsePcapStart) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ResponsePcapStart) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnabled

`func (o *ResponsePcapStart) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ResponsePcapStart) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ResponsePcapStart) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ResponsePcapStart) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpiry

`func (o *ResponsePcapStart) GetExpiry() float32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ResponsePcapStart) GetExpiryOk() (*float32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ResponsePcapStart) SetExpiry(v float32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ResponsePcapStart) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetFormat

`func (o *ResponsePcapStart) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ResponsePcapStart) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ResponsePcapStart) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ResponsePcapStart) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetId

`func (o *ResponsePcapStart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePcapStart) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePcapStart) SetId(v string)`

SetId sets Id field to given value.


### GetIncludeMcast

`func (o *ResponsePcapStart) GetIncludeMcast() bool`

GetIncludeMcast returns the IncludeMcast field if non-nil, zero value otherwise.

### GetIncludeMcastOk

`func (o *ResponsePcapStart) GetIncludeMcastOk() (*bool, bool)`

GetIncludeMcastOk returns a tuple with the IncludeMcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMcast

`func (o *ResponsePcapStart) SetIncludeMcast(v bool)`

SetIncludeMcast sets IncludeMcast field to given value.

### HasIncludeMcast

`func (o *ResponsePcapStart) HasIncludeMcast() bool`

HasIncludeMcast returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *ResponsePcapStart) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *ResponsePcapStart) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *ResponsePcapStart) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *ResponsePcapStart) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *ResponsePcapStart) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *ResponsePcapStart) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *ResponsePcapStart) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *ResponsePcapStart) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetOrgId

`func (o *ResponsePcapStart) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResponsePcapStart) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResponsePcapStart) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRaw

`func (o *ResponsePcapStart) GetRaw() bool`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *ResponsePcapStart) GetRawOk() (*bool, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *ResponsePcapStart) SetRaw(v bool)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *ResponsePcapStart) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetSiteId

`func (o *ResponsePcapStart) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponsePcapStart) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponsePcapStart) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSsid

`func (o *ResponsePcapStart) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ResponsePcapStart) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ResponsePcapStart) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ResponsePcapStart) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### SetSsidNil

`func (o *ResponsePcapStart) SetSsidNil(b bool)`

 SetSsidNil sets the value for Ssid to be an explicit nil

### UnsetSsid
`func (o *ResponsePcapStart) UnsetSsid()`

UnsetSsid ensures that no value is present for Ssid, not even an explicit nil
### GetTcpdumpParserExpression

`func (o *ResponsePcapStart) GetTcpdumpParserExpression() string`

GetTcpdumpParserExpression returns the TcpdumpParserExpression field if non-nil, zero value otherwise.

### GetTcpdumpParserExpressionOk

`func (o *ResponsePcapStart) GetTcpdumpParserExpressionOk() (*string, bool)`

GetTcpdumpParserExpressionOk returns a tuple with the TcpdumpParserExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdumpParserExpression

`func (o *ResponsePcapStart) SetTcpdumpParserExpression(v string)`

SetTcpdumpParserExpression sets TcpdumpParserExpression field to given value.

### HasTcpdumpParserExpression

`func (o *ResponsePcapStart) HasTcpdumpParserExpression() bool`

HasTcpdumpParserExpression returns a boolean if a field has been set.

### SetTcpdumpParserExpressionNil

`func (o *ResponsePcapStart) SetTcpdumpParserExpressionNil(b bool)`

 SetTcpdumpParserExpressionNil sets the value for TcpdumpParserExpression to be an explicit nil

### UnsetTcpdumpParserExpression
`func (o *ResponsePcapStart) UnsetTcpdumpParserExpression()`

UnsetTcpdumpParserExpression ensures that no value is present for TcpdumpParserExpression, not even an explicit nil
### GetTimestamp

`func (o *ResponsePcapStart) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponsePcapStart) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponsePcapStart) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *ResponsePcapStart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePcapStart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePcapStart) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



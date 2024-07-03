# CaptureRadiotapwired

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **NullableString** |  | [optional] 
**Band** | Pointer to [**CaptureRadiotapwiredBand**](CaptureRadiotapwiredBand.md) |  | [optional] [default to CAPTURERADIOTAPWIREDBAND__24]
**ClientMac** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | Pointer to [**CaptureRadiotapwiredFormat**](CaptureRadiotapwiredFormat.md) |  | [optional] [default to CAPTURERADIOTAPWIREDFORMAT_PCAP]
**MaxPktLen** | Pointer to **int32** | max_len of each packet to capture | [optional] [default to 128]
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**RadiotapTcpdumpExpression** | Pointer to **string** | tcpdump expression for radiotap interface (802.11 + radio headers) | [optional] 
**Ssid** | Pointer to **NullableString** |  | [optional] 
**TcpdumpExpression** | Pointer to **string** | tcpdump expression common for wired,radiotap | [optional] 
**Type** | [**CaptureRadiotapwiredType**](CaptureRadiotapwiredType.md) |  | 
**WiredTcpdumpExpression** | Pointer to **string** | tcpdump expression for wired | [optional] 
**WirelessTcpdumpExpression** | Pointer to **string** | tcpdump expression for radiotap interface (802.11) | [optional] 
**WlanId** | Pointer to **NullableString** | wlan id associated with the respective ssid. | [optional] 

## Methods

### NewCaptureRadiotapwired

`func NewCaptureRadiotapwired(type_ CaptureRadiotapwiredType, ) *CaptureRadiotapwired`

NewCaptureRadiotapwired instantiates a new CaptureRadiotapwired object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureRadiotapwiredWithDefaults

`func NewCaptureRadiotapwiredWithDefaults() *CaptureRadiotapwired`

NewCaptureRadiotapwiredWithDefaults instantiates a new CaptureRadiotapwired object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *CaptureRadiotapwired) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *CaptureRadiotapwired) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *CaptureRadiotapwired) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *CaptureRadiotapwired) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### SetApMacNil

`func (o *CaptureRadiotapwired) SetApMacNil(b bool)`

 SetApMacNil sets the value for ApMac to be an explicit nil

### UnsetApMac
`func (o *CaptureRadiotapwired) UnsetApMac()`

UnsetApMac ensures that no value is present for ApMac, not even an explicit nil
### GetBand

`func (o *CaptureRadiotapwired) GetBand() CaptureRadiotapwiredBand`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *CaptureRadiotapwired) GetBandOk() (*CaptureRadiotapwiredBand, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *CaptureRadiotapwired) SetBand(v CaptureRadiotapwiredBand)`

SetBand sets Band field to given value.

### HasBand

`func (o *CaptureRadiotapwired) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetClientMac

`func (o *CaptureRadiotapwired) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *CaptureRadiotapwired) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *CaptureRadiotapwired) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *CaptureRadiotapwired) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### SetClientMacNil

`func (o *CaptureRadiotapwired) SetClientMacNil(b bool)`

 SetClientMacNil sets the value for ClientMac to be an explicit nil

### UnsetClientMac
`func (o *CaptureRadiotapwired) UnsetClientMac()`

UnsetClientMac ensures that no value is present for ClientMac, not even an explicit nil
### GetDuration

`func (o *CaptureRadiotapwired) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureRadiotapwired) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureRadiotapwired) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureRadiotapwired) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFormat

`func (o *CaptureRadiotapwired) GetFormat() CaptureRadiotapwiredFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureRadiotapwired) GetFormatOk() (*CaptureRadiotapwiredFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureRadiotapwired) SetFormat(v CaptureRadiotapwiredFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureRadiotapwired) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureRadiotapwired) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureRadiotapwired) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureRadiotapwired) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureRadiotapwired) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureRadiotapwired) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureRadiotapwired) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureRadiotapwired) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureRadiotapwired) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetRadiotapTcpdumpExpression

`func (o *CaptureRadiotapwired) GetRadiotapTcpdumpExpression() string`

GetRadiotapTcpdumpExpression returns the RadiotapTcpdumpExpression field if non-nil, zero value otherwise.

### GetRadiotapTcpdumpExpressionOk

`func (o *CaptureRadiotapwired) GetRadiotapTcpdumpExpressionOk() (*string, bool)`

GetRadiotapTcpdumpExpressionOk returns a tuple with the RadiotapTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiotapTcpdumpExpression

`func (o *CaptureRadiotapwired) SetRadiotapTcpdumpExpression(v string)`

SetRadiotapTcpdumpExpression sets RadiotapTcpdumpExpression field to given value.

### HasRadiotapTcpdumpExpression

`func (o *CaptureRadiotapwired) HasRadiotapTcpdumpExpression() bool`

HasRadiotapTcpdumpExpression returns a boolean if a field has been set.

### GetSsid

`func (o *CaptureRadiotapwired) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *CaptureRadiotapwired) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *CaptureRadiotapwired) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *CaptureRadiotapwired) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### SetSsidNil

`func (o *CaptureRadiotapwired) SetSsidNil(b bool)`

 SetSsidNil sets the value for Ssid to be an explicit nil

### UnsetSsid
`func (o *CaptureRadiotapwired) UnsetSsid()`

UnsetSsid ensures that no value is present for Ssid, not even an explicit nil
### GetTcpdumpExpression

`func (o *CaptureRadiotapwired) GetTcpdumpExpression() string`

GetTcpdumpExpression returns the TcpdumpExpression field if non-nil, zero value otherwise.

### GetTcpdumpExpressionOk

`func (o *CaptureRadiotapwired) GetTcpdumpExpressionOk() (*string, bool)`

GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdumpExpression

`func (o *CaptureRadiotapwired) SetTcpdumpExpression(v string)`

SetTcpdumpExpression sets TcpdumpExpression field to given value.

### HasTcpdumpExpression

`func (o *CaptureRadiotapwired) HasTcpdumpExpression() bool`

HasTcpdumpExpression returns a boolean if a field has been set.

### GetType

`func (o *CaptureRadiotapwired) GetType() CaptureRadiotapwiredType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureRadiotapwired) GetTypeOk() (*CaptureRadiotapwiredType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureRadiotapwired) SetType(v CaptureRadiotapwiredType)`

SetType sets Type field to given value.


### GetWiredTcpdumpExpression

`func (o *CaptureRadiotapwired) GetWiredTcpdumpExpression() string`

GetWiredTcpdumpExpression returns the WiredTcpdumpExpression field if non-nil, zero value otherwise.

### GetWiredTcpdumpExpressionOk

`func (o *CaptureRadiotapwired) GetWiredTcpdumpExpressionOk() (*string, bool)`

GetWiredTcpdumpExpressionOk returns a tuple with the WiredTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredTcpdumpExpression

`func (o *CaptureRadiotapwired) SetWiredTcpdumpExpression(v string)`

SetWiredTcpdumpExpression sets WiredTcpdumpExpression field to given value.

### HasWiredTcpdumpExpression

`func (o *CaptureRadiotapwired) HasWiredTcpdumpExpression() bool`

HasWiredTcpdumpExpression returns a boolean if a field has been set.

### GetWirelessTcpdumpExpression

`func (o *CaptureRadiotapwired) GetWirelessTcpdumpExpression() string`

GetWirelessTcpdumpExpression returns the WirelessTcpdumpExpression field if non-nil, zero value otherwise.

### GetWirelessTcpdumpExpressionOk

`func (o *CaptureRadiotapwired) GetWirelessTcpdumpExpressionOk() (*string, bool)`

GetWirelessTcpdumpExpressionOk returns a tuple with the WirelessTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessTcpdumpExpression

`func (o *CaptureRadiotapwired) SetWirelessTcpdumpExpression(v string)`

SetWirelessTcpdumpExpression sets WirelessTcpdumpExpression field to given value.

### HasWirelessTcpdumpExpression

`func (o *CaptureRadiotapwired) HasWirelessTcpdumpExpression() bool`

HasWirelessTcpdumpExpression returns a boolean if a field has been set.

### GetWlanId

`func (o *CaptureRadiotapwired) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *CaptureRadiotapwired) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *CaptureRadiotapwired) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *CaptureRadiotapwired) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.

### SetWlanIdNil

`func (o *CaptureRadiotapwired) SetWlanIdNil(b bool)`

 SetWlanIdNil sets the value for WlanId to be an explicit nil

### UnsetWlanId
`func (o *CaptureRadiotapwired) UnsetWlanId()`

UnsetWlanId ensures that no value is present for WlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CaptureRadiotap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **string** |  | [optional] 
**Band** | Pointer to [**CaptureRadiotapBand**](CaptureRadiotapBand.md) |  | [optional] [default to CAPTURERADIOTAPBAND__24]
**ClientMac** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | Pointer to [**CaptureRadiotapFormat**](CaptureRadiotapFormat.md) |  | [optional] [default to CAPTURERADIOTAPFORMAT_PCAP]
**MaxPktLen** | Pointer to **int32** | max_len of each packet to capture | [optional] [default to 128]
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**Ssid** | Pointer to **string** |  | [optional] 
**TcpdumpExpression** | Pointer to **string** | tcpdump expression specific to radiotap | [optional] 
**Type** | [**CaptureRadiotapType**](CaptureRadiotapType.md) |  | 
**WlanId** | Pointer to **string** | wlan id associated with the respective ssid. | [optional] 

## Methods

### NewCaptureRadiotap

`func NewCaptureRadiotap(type_ CaptureRadiotapType, ) *CaptureRadiotap`

NewCaptureRadiotap instantiates a new CaptureRadiotap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureRadiotapWithDefaults

`func NewCaptureRadiotapWithDefaults() *CaptureRadiotap`

NewCaptureRadiotapWithDefaults instantiates a new CaptureRadiotap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *CaptureRadiotap) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *CaptureRadiotap) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *CaptureRadiotap) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *CaptureRadiotap) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### GetBand

`func (o *CaptureRadiotap) GetBand() CaptureRadiotapBand`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *CaptureRadiotap) GetBandOk() (*CaptureRadiotapBand, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *CaptureRadiotap) SetBand(v CaptureRadiotapBand)`

SetBand sets Band field to given value.

### HasBand

`func (o *CaptureRadiotap) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetClientMac

`func (o *CaptureRadiotap) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *CaptureRadiotap) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *CaptureRadiotap) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *CaptureRadiotap) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetDuration

`func (o *CaptureRadiotap) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureRadiotap) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureRadiotap) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureRadiotap) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFormat

`func (o *CaptureRadiotap) GetFormat() CaptureRadiotapFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureRadiotap) GetFormatOk() (*CaptureRadiotapFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureRadiotap) SetFormat(v CaptureRadiotapFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureRadiotap) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureRadiotap) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureRadiotap) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureRadiotap) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureRadiotap) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureRadiotap) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureRadiotap) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureRadiotap) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureRadiotap) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetSsid

`func (o *CaptureRadiotap) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *CaptureRadiotap) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *CaptureRadiotap) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *CaptureRadiotap) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetTcpdumpExpression

`func (o *CaptureRadiotap) GetTcpdumpExpression() string`

GetTcpdumpExpression returns the TcpdumpExpression field if non-nil, zero value otherwise.

### GetTcpdumpExpressionOk

`func (o *CaptureRadiotap) GetTcpdumpExpressionOk() (*string, bool)`

GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdumpExpression

`func (o *CaptureRadiotap) SetTcpdumpExpression(v string)`

SetTcpdumpExpression sets TcpdumpExpression field to given value.

### HasTcpdumpExpression

`func (o *CaptureRadiotap) HasTcpdumpExpression() bool`

HasTcpdumpExpression returns a boolean if a field has been set.

### GetType

`func (o *CaptureRadiotap) GetType() CaptureRadiotapType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureRadiotap) GetTypeOk() (*CaptureRadiotapType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureRadiotap) SetType(v CaptureRadiotapType)`

SetType sets Type field to given value.


### GetWlanId

`func (o *CaptureRadiotap) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *CaptureRadiotap) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *CaptureRadiotap) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *CaptureRadiotap) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



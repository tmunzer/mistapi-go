# CaptureWireless

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **NullableString** |  | [optional] 
**Band** | Pointer to [**CaptureWirelessBand**](CaptureWirelessBand.md) |  | [optional] [default to CAPTUREWIRELESSBAND__24]
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | Pointer to [**CaptureWirelessFormat**](CaptureWirelessFormat.md) |  | [optional] [default to CAPTUREWIRELESSFORMAT_PCAP]
**MaxPktLen** | Pointer to **int32** | max_len of each packet to capture | [optional] [default to 128]
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**Ssid** | Pointer to **string** |  | [optional] 
**Type** | [**CaptureWirelessType**](CaptureWirelessType.md) |  | 
**WlanId** | Pointer to **string** | WLAN ID | [optional] 

## Methods

### NewCaptureWireless

`func NewCaptureWireless(type_ CaptureWirelessType, ) *CaptureWireless`

NewCaptureWireless instantiates a new CaptureWireless object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureWirelessWithDefaults

`func NewCaptureWirelessWithDefaults() *CaptureWireless`

NewCaptureWirelessWithDefaults instantiates a new CaptureWireless object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *CaptureWireless) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *CaptureWireless) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *CaptureWireless) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *CaptureWireless) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### SetApMacNil

`func (o *CaptureWireless) SetApMacNil(b bool)`

 SetApMacNil sets the value for ApMac to be an explicit nil

### UnsetApMac
`func (o *CaptureWireless) UnsetApMac()`

UnsetApMac ensures that no value is present for ApMac, not even an explicit nil
### GetBand

`func (o *CaptureWireless) GetBand() CaptureWirelessBand`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *CaptureWireless) GetBandOk() (*CaptureWirelessBand, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *CaptureWireless) SetBand(v CaptureWirelessBand)`

SetBand sets Band field to given value.

### HasBand

`func (o *CaptureWireless) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetDuration

`func (o *CaptureWireless) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureWireless) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureWireless) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureWireless) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFormat

`func (o *CaptureWireless) GetFormat() CaptureWirelessFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureWireless) GetFormatOk() (*CaptureWirelessFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureWireless) SetFormat(v CaptureWirelessFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureWireless) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureWireless) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureWireless) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureWireless) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureWireless) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureWireless) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureWireless) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureWireless) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureWireless) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetSsid

`func (o *CaptureWireless) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *CaptureWireless) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *CaptureWireless) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *CaptureWireless) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetType

`func (o *CaptureWireless) GetType() CaptureWirelessType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureWireless) GetTypeOk() (*CaptureWirelessType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureWireless) SetType(v CaptureWirelessType)`

SetType sets Type field to given value.


### GetWlanId

`func (o *CaptureWireless) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *CaptureWireless) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *CaptureWireless) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *CaptureWireless) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CaptureScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **NullableString** | filter by ap_mac | [optional] 
**Aps** | Pointer to [**map[string]CaptureScanAps**](CaptureScanAps.md) | dictionary key is AP mac and value is a dictionary which contains key “band”, “bandwidth”, “channel” and “tcpdump_expression”. In case keys are missed we will take parent value if parent values are not set we will use default value | [optional] 
**Band** | Pointer to [**NullableCaptureScanBand**](CaptureScanBand.md) |  | [optional] [default to CAPTURESCANBAND__5]
**Bandwidth** | Pointer to [**Dot11Bandwidth**](Dot11Bandwidth.md) |  | [optional] 
**Channel** | Pointer to **int32** | specify the channel value where scan PCAP has to be started, default value gets applied when user provides wrong values | [optional] [default to 1]
**ClientMac** | Pointer to **NullableString** | filter by client mac | [optional] 
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | Pointer to [**CaptureScanFormat**](CaptureScanFormat.md) |  | [optional] [default to CAPTURESCANFORMAT_PCAP]
**MaxPktLen** | Pointer to **int32** | max_len of each packet to capture | [optional] [default to 512]
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**TcpdumpExpression** | Pointer to **string** | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. | [optional] 
**Type** | [**CaptureScanType**](CaptureScanType.md) |  | 
**Width** | Pointer to **string** | specify the bandwidth value with respect to the channel. | [optional] 

## Methods

### NewCaptureScan

`func NewCaptureScan(type_ CaptureScanType, ) *CaptureScan`

NewCaptureScan instantiates a new CaptureScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureScanWithDefaults

`func NewCaptureScanWithDefaults() *CaptureScan`

NewCaptureScanWithDefaults instantiates a new CaptureScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *CaptureScan) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *CaptureScan) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *CaptureScan) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *CaptureScan) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### SetApMacNil

`func (o *CaptureScan) SetApMacNil(b bool)`

 SetApMacNil sets the value for ApMac to be an explicit nil

### UnsetApMac
`func (o *CaptureScan) UnsetApMac()`

UnsetApMac ensures that no value is present for ApMac, not even an explicit nil
### GetAps

`func (o *CaptureScan) GetAps() map[string]CaptureScanAps`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *CaptureScan) GetApsOk() (*map[string]CaptureScanAps, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *CaptureScan) SetAps(v map[string]CaptureScanAps)`

SetAps sets Aps field to given value.

### HasAps

`func (o *CaptureScan) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetBand

`func (o *CaptureScan) GetBand() CaptureScanBand`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *CaptureScan) GetBandOk() (*CaptureScanBand, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *CaptureScan) SetBand(v CaptureScanBand)`

SetBand sets Band field to given value.

### HasBand

`func (o *CaptureScan) HasBand() bool`

HasBand returns a boolean if a field has been set.

### SetBandNil

`func (o *CaptureScan) SetBandNil(b bool)`

 SetBandNil sets the value for Band to be an explicit nil

### UnsetBand
`func (o *CaptureScan) UnsetBand()`

UnsetBand ensures that no value is present for Band, not even an explicit nil
### GetBandwidth

`func (o *CaptureScan) GetBandwidth() Dot11Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *CaptureScan) GetBandwidthOk() (*Dot11Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *CaptureScan) SetBandwidth(v Dot11Bandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *CaptureScan) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetChannel

`func (o *CaptureScan) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CaptureScan) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CaptureScan) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CaptureScan) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetClientMac

`func (o *CaptureScan) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *CaptureScan) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *CaptureScan) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *CaptureScan) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### SetClientMacNil

`func (o *CaptureScan) SetClientMacNil(b bool)`

 SetClientMacNil sets the value for ClientMac to be an explicit nil

### UnsetClientMac
`func (o *CaptureScan) UnsetClientMac()`

UnsetClientMac ensures that no value is present for ClientMac, not even an explicit nil
### GetDuration

`func (o *CaptureScan) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureScan) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureScan) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureScan) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFormat

`func (o *CaptureScan) GetFormat() CaptureScanFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureScan) GetFormatOk() (*CaptureScanFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureScan) SetFormat(v CaptureScanFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureScan) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureScan) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureScan) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureScan) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureScan) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureScan) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureScan) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureScan) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureScan) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetTcpdumpExpression

`func (o *CaptureScan) GetTcpdumpExpression() string`

GetTcpdumpExpression returns the TcpdumpExpression field if non-nil, zero value otherwise.

### GetTcpdumpExpressionOk

`func (o *CaptureScan) GetTcpdumpExpressionOk() (*string, bool)`

GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdumpExpression

`func (o *CaptureScan) SetTcpdumpExpression(v string)`

SetTcpdumpExpression sets TcpdumpExpression field to given value.

### HasTcpdumpExpression

`func (o *CaptureScan) HasTcpdumpExpression() bool`

HasTcpdumpExpression returns a boolean if a field has been set.

### GetType

`func (o *CaptureScan) GetType() CaptureScanType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureScan) GetTypeOk() (*CaptureScanType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureScan) SetType(v CaptureScanType)`

SetType sets Type field to given value.


### GetWidth

`func (o *CaptureScan) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CaptureScan) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CaptureScan) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *CaptureScan) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



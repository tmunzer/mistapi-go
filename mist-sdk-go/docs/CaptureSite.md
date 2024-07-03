# CaptureSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **NullableString** |  | [optional] 
**ClientMac** | Pointer to **NullableString** | filter by client mac | [optional] 
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**IncludesMcast** | Pointer to **bool** |  | [optional] [default to false]
**MaxPktLen** | Pointer to **int32** | max_len of each packet to capture | [optional] [default to 128]
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**Ssid** | Pointer to **string** |  | [optional] 
**Type** | [**CaptureWirelessType**](CaptureWirelessType.md) |  | 
**Format** | Pointer to [**CaptureWirelessFormat**](CaptureWirelessFormat.md) |  | [optional] [default to CAPTUREWIRELESSFORMAT_PCAP]
**Gateways** | Pointer to [**map[string]CaptureGatewayGateways**](CaptureGatewayGateways.md) | List of SSRs. Property key is the SSR MAC | [optional] 
**Ports** | Pointer to **[]string** | dict of port which uses port id as the key | [optional] 
**Band** | Pointer to [**CaptureWirelessBand**](CaptureWirelessBand.md) |  | [optional] [default to CAPTUREWIRELESSBAND__24]
**TcpdumpExpression** | Pointer to **NullableString** | tcpdump expression | [optional] 
**WlanId** | Pointer to **string** | WLAN ID | [optional] 
**RadiotapTcpdumpExpression** | Pointer to **string** | tcpdump expression for radiotap interface (802.11 + radio headers) | [optional] 
**WiredTcpdumpExpression** | Pointer to **string** | tcpdump expression for wired | [optional] 
**WirelessTcpdumpExpression** | Pointer to **string** | tcpdump expression for radiotap interface (802.11) | [optional] 
**Aps** | Pointer to [**map[string]CaptureScanAps**](CaptureScanAps.md) | dictionary key is AP mac and value is a dictionary which contains key “band”, “bandwidth”, “channel” and “tcpdump_expression”. In case keys are missed we will take parent value if parent values are not set we will use default value | [optional] 
**Bandwidth** | Pointer to [**Dot11Bandwidth**](Dot11Bandwidth.md) |  | [optional] 
**Channel** | Pointer to **int32** | specify the channel value where scan PCAP has to be started, default value gets applied when user provides wrong values | [optional] [default to 1]
**Width** | Pointer to **string** | specify the bandwidth value with respect to the channel. | [optional] 
**Switches** | Pointer to [**map[string]CaptureSwitchSwitches**](CaptureSwitchSwitches.md) | Property key is the switch mac | [optional] 

## Methods

### NewCaptureSite

`func NewCaptureSite(type_ CaptureWirelessType, ) *CaptureSite`

NewCaptureSite instantiates a new CaptureSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureSiteWithDefaults

`func NewCaptureSiteWithDefaults() *CaptureSite`

NewCaptureSiteWithDefaults instantiates a new CaptureSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *CaptureSite) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *CaptureSite) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *CaptureSite) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *CaptureSite) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### SetApMacNil

`func (o *CaptureSite) SetApMacNil(b bool)`

 SetApMacNil sets the value for ApMac to be an explicit nil

### UnsetApMac
`func (o *CaptureSite) UnsetApMac()`

UnsetApMac ensures that no value is present for ApMac, not even an explicit nil
### GetClientMac

`func (o *CaptureSite) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *CaptureSite) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *CaptureSite) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *CaptureSite) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### SetClientMacNil

`func (o *CaptureSite) SetClientMacNil(b bool)`

 SetClientMacNil sets the value for ClientMac to be an explicit nil

### UnsetClientMac
`func (o *CaptureSite) UnsetClientMac()`

UnsetClientMac ensures that no value is present for ClientMac, not even an explicit nil
### GetDuration

`func (o *CaptureSite) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureSite) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureSite) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureSite) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetIncludesMcast

`func (o *CaptureSite) GetIncludesMcast() bool`

GetIncludesMcast returns the IncludesMcast field if non-nil, zero value otherwise.

### GetIncludesMcastOk

`func (o *CaptureSite) GetIncludesMcastOk() (*bool, bool)`

GetIncludesMcastOk returns a tuple with the IncludesMcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesMcast

`func (o *CaptureSite) SetIncludesMcast(v bool)`

SetIncludesMcast sets IncludesMcast field to given value.

### HasIncludesMcast

`func (o *CaptureSite) HasIncludesMcast() bool`

HasIncludesMcast returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureSite) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureSite) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureSite) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureSite) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureSite) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureSite) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureSite) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureSite) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetSsid

`func (o *CaptureSite) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *CaptureSite) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *CaptureSite) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *CaptureSite) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetType

`func (o *CaptureSite) GetType() CaptureWirelessType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureSite) GetTypeOk() (*CaptureWirelessType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureSite) SetType(v CaptureWirelessType)`

SetType sets Type field to given value.


### GetFormat

`func (o *CaptureSite) GetFormat() CaptureWirelessFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureSite) GetFormatOk() (*CaptureWirelessFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureSite) SetFormat(v CaptureWirelessFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureSite) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGateways

`func (o *CaptureSite) GetGateways() map[string]CaptureGatewayGateways`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *CaptureSite) GetGatewaysOk() (*map[string]CaptureGatewayGateways, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *CaptureSite) SetGateways(v map[string]CaptureGatewayGateways)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *CaptureSite) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetPorts

`func (o *CaptureSite) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CaptureSite) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CaptureSite) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *CaptureSite) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetBand

`func (o *CaptureSite) GetBand() CaptureWirelessBand`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *CaptureSite) GetBandOk() (*CaptureWirelessBand, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *CaptureSite) SetBand(v CaptureWirelessBand)`

SetBand sets Band field to given value.

### HasBand

`func (o *CaptureSite) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetTcpdumpExpression

`func (o *CaptureSite) GetTcpdumpExpression() string`

GetTcpdumpExpression returns the TcpdumpExpression field if non-nil, zero value otherwise.

### GetTcpdumpExpressionOk

`func (o *CaptureSite) GetTcpdumpExpressionOk() (*string, bool)`

GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdumpExpression

`func (o *CaptureSite) SetTcpdumpExpression(v string)`

SetTcpdumpExpression sets TcpdumpExpression field to given value.

### HasTcpdumpExpression

`func (o *CaptureSite) HasTcpdumpExpression() bool`

HasTcpdumpExpression returns a boolean if a field has been set.

### SetTcpdumpExpressionNil

`func (o *CaptureSite) SetTcpdumpExpressionNil(b bool)`

 SetTcpdumpExpressionNil sets the value for TcpdumpExpression to be an explicit nil

### UnsetTcpdumpExpression
`func (o *CaptureSite) UnsetTcpdumpExpression()`

UnsetTcpdumpExpression ensures that no value is present for TcpdumpExpression, not even an explicit nil
### GetWlanId

`func (o *CaptureSite) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *CaptureSite) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *CaptureSite) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *CaptureSite) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.

### GetRadiotapTcpdumpExpression

`func (o *CaptureSite) GetRadiotapTcpdumpExpression() string`

GetRadiotapTcpdumpExpression returns the RadiotapTcpdumpExpression field if non-nil, zero value otherwise.

### GetRadiotapTcpdumpExpressionOk

`func (o *CaptureSite) GetRadiotapTcpdumpExpressionOk() (*string, bool)`

GetRadiotapTcpdumpExpressionOk returns a tuple with the RadiotapTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiotapTcpdumpExpression

`func (o *CaptureSite) SetRadiotapTcpdumpExpression(v string)`

SetRadiotapTcpdumpExpression sets RadiotapTcpdumpExpression field to given value.

### HasRadiotapTcpdumpExpression

`func (o *CaptureSite) HasRadiotapTcpdumpExpression() bool`

HasRadiotapTcpdumpExpression returns a boolean if a field has been set.

### GetWiredTcpdumpExpression

`func (o *CaptureSite) GetWiredTcpdumpExpression() string`

GetWiredTcpdumpExpression returns the WiredTcpdumpExpression field if non-nil, zero value otherwise.

### GetWiredTcpdumpExpressionOk

`func (o *CaptureSite) GetWiredTcpdumpExpressionOk() (*string, bool)`

GetWiredTcpdumpExpressionOk returns a tuple with the WiredTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredTcpdumpExpression

`func (o *CaptureSite) SetWiredTcpdumpExpression(v string)`

SetWiredTcpdumpExpression sets WiredTcpdumpExpression field to given value.

### HasWiredTcpdumpExpression

`func (o *CaptureSite) HasWiredTcpdumpExpression() bool`

HasWiredTcpdumpExpression returns a boolean if a field has been set.

### GetWirelessTcpdumpExpression

`func (o *CaptureSite) GetWirelessTcpdumpExpression() string`

GetWirelessTcpdumpExpression returns the WirelessTcpdumpExpression field if non-nil, zero value otherwise.

### GetWirelessTcpdumpExpressionOk

`func (o *CaptureSite) GetWirelessTcpdumpExpressionOk() (*string, bool)`

GetWirelessTcpdumpExpressionOk returns a tuple with the WirelessTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessTcpdumpExpression

`func (o *CaptureSite) SetWirelessTcpdumpExpression(v string)`

SetWirelessTcpdumpExpression sets WirelessTcpdumpExpression field to given value.

### HasWirelessTcpdumpExpression

`func (o *CaptureSite) HasWirelessTcpdumpExpression() bool`

HasWirelessTcpdumpExpression returns a boolean if a field has been set.

### GetAps

`func (o *CaptureSite) GetAps() map[string]CaptureScanAps`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *CaptureSite) GetApsOk() (*map[string]CaptureScanAps, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *CaptureSite) SetAps(v map[string]CaptureScanAps)`

SetAps sets Aps field to given value.

### HasAps

`func (o *CaptureSite) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetBandwidth

`func (o *CaptureSite) GetBandwidth() Dot11Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *CaptureSite) GetBandwidthOk() (*Dot11Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *CaptureSite) SetBandwidth(v Dot11Bandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *CaptureSite) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetChannel

`func (o *CaptureSite) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CaptureSite) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CaptureSite) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CaptureSite) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetWidth

`func (o *CaptureSite) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CaptureSite) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CaptureSite) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *CaptureSite) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetSwitches

`func (o *CaptureSite) GetSwitches() map[string]CaptureSwitchSwitches`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *CaptureSite) GetSwitchesOk() (*map[string]CaptureSwitchSwitches, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *CaptureSite) SetSwitches(v map[string]CaptureSwitchSwitches)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *CaptureSite) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



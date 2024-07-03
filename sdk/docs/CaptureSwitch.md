# CaptureSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | Pointer to [**CaptureSwitchFormat**](CaptureSwitchFormat.md) |  | [optional] [default to CAPTURESWITCHFORMAT_STREAM]
**MaxPktLen** | Pointer to **int32** | max_len of each packet to capture | [optional] [default to 512]
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**Ports** | Pointer to **[]string** | dict of port which uses port id as the key | [optional] 
**Switches** | Pointer to [**map[string]CaptureSwitchSwitches**](CaptureSwitchSwitches.md) | Property key is the switch mac | [optional] 
**TcpdumpExpression** | Pointer to **string** | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. | [optional] 
**Type** | [**CaptureSwitchType**](CaptureSwitchType.md) |  | 

## Methods

### NewCaptureSwitch

`func NewCaptureSwitch(type_ CaptureSwitchType, ) *CaptureSwitch`

NewCaptureSwitch instantiates a new CaptureSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureSwitchWithDefaults

`func NewCaptureSwitchWithDefaults() *CaptureSwitch`

NewCaptureSwitchWithDefaults instantiates a new CaptureSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *CaptureSwitch) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureSwitch) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureSwitch) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureSwitch) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFormat

`func (o *CaptureSwitch) GetFormat() CaptureSwitchFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureSwitch) GetFormatOk() (*CaptureSwitchFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureSwitch) SetFormat(v CaptureSwitchFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureSwitch) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureSwitch) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureSwitch) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureSwitch) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureSwitch) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureSwitch) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureSwitch) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureSwitch) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureSwitch) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetPorts

`func (o *CaptureSwitch) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CaptureSwitch) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CaptureSwitch) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *CaptureSwitch) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetSwitches

`func (o *CaptureSwitch) GetSwitches() map[string]CaptureSwitchSwitches`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *CaptureSwitch) GetSwitchesOk() (*map[string]CaptureSwitchSwitches, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *CaptureSwitch) SetSwitches(v map[string]CaptureSwitchSwitches)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *CaptureSwitch) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetTcpdumpExpression

`func (o *CaptureSwitch) GetTcpdumpExpression() string`

GetTcpdumpExpression returns the TcpdumpExpression field if non-nil, zero value otherwise.

### GetTcpdumpExpressionOk

`func (o *CaptureSwitch) GetTcpdumpExpressionOk() (*string, bool)`

GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdumpExpression

`func (o *CaptureSwitch) SetTcpdumpExpression(v string)`

SetTcpdumpExpression sets TcpdumpExpression field to given value.

### HasTcpdumpExpression

`func (o *CaptureSwitch) HasTcpdumpExpression() bool`

HasTcpdumpExpression returns a boolean if a field has been set.

### GetType

`func (o *CaptureSwitch) GetType() CaptureSwitchType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureSwitch) GetTypeOk() (*CaptureSwitchType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureSwitch) SetType(v CaptureSwitchType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



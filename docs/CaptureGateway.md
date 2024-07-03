# CaptureGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | Pointer to [**CaptureGatewayFormat**](CaptureGatewayFormat.md) |  | [optional] [default to CAPTUREGATEWAYFORMAT_STREAM]
**Gateways** | Pointer to [**map[string]CaptureGatewayGateways**](CaptureGatewayGateways.md) | List of SSRs. Property key is the SSR MAC | [optional] 
**MaxPktLen** | Pointer to **int32** | max_len of each packet to capture | [optional] [default to 128]
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**Ports** | Pointer to **[]string** | dict of port which uses port id as the key | [optional] 
**Type** | [**CaptureGatewayType**](CaptureGatewayType.md) |  | 

## Methods

### NewCaptureGateway

`func NewCaptureGateway(type_ CaptureGatewayType, ) *CaptureGateway`

NewCaptureGateway instantiates a new CaptureGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureGatewayWithDefaults

`func NewCaptureGatewayWithDefaults() *CaptureGateway`

NewCaptureGatewayWithDefaults instantiates a new CaptureGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *CaptureGateway) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureGateway) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureGateway) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureGateway) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFormat

`func (o *CaptureGateway) GetFormat() CaptureGatewayFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureGateway) GetFormatOk() (*CaptureGatewayFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureGateway) SetFormat(v CaptureGatewayFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureGateway) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGateways

`func (o *CaptureGateway) GetGateways() map[string]CaptureGatewayGateways`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *CaptureGateway) GetGatewaysOk() (*map[string]CaptureGatewayGateways, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *CaptureGateway) SetGateways(v map[string]CaptureGatewayGateways)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *CaptureGateway) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureGateway) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureGateway) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureGateway) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureGateway) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureGateway) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureGateway) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureGateway) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureGateway) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetPorts

`func (o *CaptureGateway) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CaptureGateway) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CaptureGateway) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *CaptureGateway) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetType

`func (o *CaptureGateway) GetType() CaptureGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureGateway) GetTypeOk() (*CaptureGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureGateway) SetType(v CaptureGatewayType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



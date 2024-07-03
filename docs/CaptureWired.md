# CaptureWired

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | Pointer to [**CaptureWiredFormat**](CaptureWiredFormat.md) |  | [optional] [default to CAPTUREWIREDFORMAT_PCAP]
**MaxPktLen** | Pointer to **int32** | max_len of each packet to capture | [optional] [default to 128]
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**TcpdumpExpression** | Pointer to **NullableString** | tcpdump expression | [optional] 
**Type** | [**CaptureWiredType**](CaptureWiredType.md) |  | 

## Methods

### NewCaptureWired

`func NewCaptureWired(type_ CaptureWiredType, ) *CaptureWired`

NewCaptureWired instantiates a new CaptureWired object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureWiredWithDefaults

`func NewCaptureWiredWithDefaults() *CaptureWired`

NewCaptureWiredWithDefaults instantiates a new CaptureWired object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *CaptureWired) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *CaptureWired) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *CaptureWired) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *CaptureWired) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### SetApMacNil

`func (o *CaptureWired) SetApMacNil(b bool)`

 SetApMacNil sets the value for ApMac to be an explicit nil

### UnsetApMac
`func (o *CaptureWired) UnsetApMac()`

UnsetApMac ensures that no value is present for ApMac, not even an explicit nil
### GetDuration

`func (o *CaptureWired) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureWired) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureWired) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureWired) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFormat

`func (o *CaptureWired) GetFormat() CaptureWiredFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureWired) GetFormatOk() (*CaptureWiredFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureWired) SetFormat(v CaptureWiredFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureWired) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureWired) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureWired) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureWired) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureWired) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureWired) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureWired) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureWired) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureWired) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetTcpdumpExpression

`func (o *CaptureWired) GetTcpdumpExpression() string`

GetTcpdumpExpression returns the TcpdumpExpression field if non-nil, zero value otherwise.

### GetTcpdumpExpressionOk

`func (o *CaptureWired) GetTcpdumpExpressionOk() (*string, bool)`

GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdumpExpression

`func (o *CaptureWired) SetTcpdumpExpression(v string)`

SetTcpdumpExpression sets TcpdumpExpression field to given value.

### HasTcpdumpExpression

`func (o *CaptureWired) HasTcpdumpExpression() bool`

HasTcpdumpExpression returns a boolean if a field has been set.

### SetTcpdumpExpressionNil

`func (o *CaptureWired) SetTcpdumpExpressionNil(b bool)`

 SetTcpdumpExpressionNil sets the value for TcpdumpExpression to be an explicit nil

### UnsetTcpdumpExpression
`func (o *CaptureWired) UnsetTcpdumpExpression()`

UnsetTcpdumpExpression ensures that no value is present for TcpdumpExpression, not even an explicit nil
### GetType

`func (o *CaptureWired) GetType() CaptureWiredType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureWired) GetTypeOk() (*CaptureWiredType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureWired) SetType(v CaptureWiredType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



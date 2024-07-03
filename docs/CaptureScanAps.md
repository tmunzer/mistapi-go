# CaptureScanAps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | Pointer to [**CaptureScanApsBand**](CaptureScanApsBand.md) |  | [optional] [default to CAPTURESCANAPSBAND__24]
**Channel** | Pointer to **string** | specify the channel value where scan PCAP has to be started | [optional] 
**TcpdumpExpression** | Pointer to **string** | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. | [optional] 
**Width** | Pointer to **string** | specify the bandwidth value with respect to the channel. | [optional] 

## Methods

### NewCaptureScanAps

`func NewCaptureScanAps() *CaptureScanAps`

NewCaptureScanAps instantiates a new CaptureScanAps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureScanApsWithDefaults

`func NewCaptureScanApsWithDefaults() *CaptureScanAps`

NewCaptureScanApsWithDefaults instantiates a new CaptureScanAps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand

`func (o *CaptureScanAps) GetBand() CaptureScanApsBand`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *CaptureScanAps) GetBandOk() (*CaptureScanApsBand, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *CaptureScanAps) SetBand(v CaptureScanApsBand)`

SetBand sets Band field to given value.

### HasBand

`func (o *CaptureScanAps) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetChannel

`func (o *CaptureScanAps) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CaptureScanAps) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CaptureScanAps) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CaptureScanAps) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTcpdumpExpression

`func (o *CaptureScanAps) GetTcpdumpExpression() string`

GetTcpdumpExpression returns the TcpdumpExpression field if non-nil, zero value otherwise.

### GetTcpdumpExpressionOk

`func (o *CaptureScanAps) GetTcpdumpExpressionOk() (*string, bool)`

GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdumpExpression

`func (o *CaptureScanAps) SetTcpdumpExpression(v string)`

SetTcpdumpExpression sets TcpdumpExpression field to given value.

### HasTcpdumpExpression

`func (o *CaptureScanAps) HasTcpdumpExpression() bool`

HasTcpdumpExpression returns a boolean if a field has been set.

### GetWidth

`func (o *CaptureScanAps) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CaptureScanAps) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CaptureScanAps) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *CaptureScanAps) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



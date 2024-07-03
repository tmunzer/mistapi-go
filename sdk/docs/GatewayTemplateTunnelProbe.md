# GatewayTemplateTunnelProbe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **int32** | how often to trigger the probe | [optional] 
**Threshold** | Pointer to **int32** | number of consecutive misses before declaring the tunnel down | [optional] 
**Timeout** | Pointer to **int32** | time within which to complete the connectivity check | [optional] 
**Type** | Pointer to [**GatewayTemplateProbeType**](GatewayTemplateProbeType.md) |  | [optional] [default to GATEWAYTEMPLATEPROBETYPE_ICMP]

## Methods

### NewGatewayTemplateTunnelProbe

`func NewGatewayTemplateTunnelProbe() *GatewayTemplateTunnelProbe`

NewGatewayTemplateTunnelProbe instantiates a new GatewayTemplateTunnelProbe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelProbeWithDefaults

`func NewGatewayTemplateTunnelProbeWithDefaults() *GatewayTemplateTunnelProbe`

NewGatewayTemplateTunnelProbeWithDefaults instantiates a new GatewayTemplateTunnelProbe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *GatewayTemplateTunnelProbe) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GatewayTemplateTunnelProbe) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GatewayTemplateTunnelProbe) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GatewayTemplateTunnelProbe) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetThreshold

`func (o *GatewayTemplateTunnelProbe) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GatewayTemplateTunnelProbe) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GatewayTemplateTunnelProbe) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *GatewayTemplateTunnelProbe) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTimeout

`func (o *GatewayTemplateTunnelProbe) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GatewayTemplateTunnelProbe) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GatewayTemplateTunnelProbe) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GatewayTemplateTunnelProbe) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *GatewayTemplateTunnelProbe) GetType() GatewayTemplateProbeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayTemplateTunnelProbe) GetTypeOk() (*GatewayTemplateProbeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayTemplateTunnelProbe) SetType(v GatewayTemplateProbeType)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayTemplateTunnelProbe) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



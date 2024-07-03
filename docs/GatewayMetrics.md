# GatewayMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigSuccess** | Pointer to **float32** | config success score | [optional] 
**VersionCompliance** | Pointer to [**GatewayComplianceVersion**](GatewayComplianceVersion.md) |  | [optional] 

## Methods

### NewGatewayMetrics

`func NewGatewayMetrics() *GatewayMetrics`

NewGatewayMetrics instantiates a new GatewayMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayMetricsWithDefaults

`func NewGatewayMetricsWithDefaults() *GatewayMetrics`

NewGatewayMetricsWithDefaults instantiates a new GatewayMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigSuccess

`func (o *GatewayMetrics) GetConfigSuccess() float32`

GetConfigSuccess returns the ConfigSuccess field if non-nil, zero value otherwise.

### GetConfigSuccessOk

`func (o *GatewayMetrics) GetConfigSuccessOk() (*float32, bool)`

GetConfigSuccessOk returns a tuple with the ConfigSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSuccess

`func (o *GatewayMetrics) SetConfigSuccess(v float32)`

SetConfigSuccess sets ConfigSuccess field to given value.

### HasConfigSuccess

`func (o *GatewayMetrics) HasConfigSuccess() bool`

HasConfigSuccess returns a boolean if a field has been set.

### GetVersionCompliance

`func (o *GatewayMetrics) GetVersionCompliance() GatewayComplianceVersion`

GetVersionCompliance returns the VersionCompliance field if non-nil, zero value otherwise.

### GetVersionComplianceOk

`func (o *GatewayMetrics) GetVersionComplianceOk() (*GatewayComplianceVersion, bool)`

GetVersionComplianceOk returns a tuple with the VersionCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCompliance

`func (o *GatewayMetrics) SetVersionCompliance(v GatewayComplianceVersion)`

SetVersionCompliance sets VersionCompliance field to given value.

### HasVersionCompliance

`func (o *GatewayMetrics) HasVersionCompliance() bool`

HasVersionCompliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



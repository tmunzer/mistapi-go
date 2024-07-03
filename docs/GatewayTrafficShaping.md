# GatewayTrafficShaping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassPercentages** | Pointer to **[]int32** | percentages for differet class of traffic: high / medium / low / best-effort sum must be equal to 100 | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGatewayTrafficShaping

`func NewGatewayTrafficShaping() *GatewayTrafficShaping`

NewGatewayTrafficShaping instantiates a new GatewayTrafficShaping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTrafficShapingWithDefaults

`func NewGatewayTrafficShapingWithDefaults() *GatewayTrafficShaping`

NewGatewayTrafficShapingWithDefaults instantiates a new GatewayTrafficShaping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassPercentages

`func (o *GatewayTrafficShaping) GetClassPercentages() []int32`

GetClassPercentages returns the ClassPercentages field if non-nil, zero value otherwise.

### GetClassPercentagesOk

`func (o *GatewayTrafficShaping) GetClassPercentagesOk() (*[]int32, bool)`

GetClassPercentagesOk returns a tuple with the ClassPercentages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassPercentages

`func (o *GatewayTrafficShaping) SetClassPercentages(v []int32)`

SetClassPercentages sets ClassPercentages field to given value.

### HasClassPercentages

`func (o *GatewayTrafficShaping) HasClassPercentages() bool`

HasClassPercentages returns a boolean if a field has been set.

### GetEnabled

`func (o *GatewayTrafficShaping) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayTrafficShaping) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayTrafficShaping) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GatewayTrafficShaping) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GatewayTemplatePathPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to [**[]GatewayTemplatePathPreferencesPath**](GatewayTemplatePathPreferencesPath.md) |  | [optional] 
**Strategy** | Pointer to [**GatewayPathStrategy**](GatewayPathStrategy.md) |  | [optional] [default to GATEWAYPATHSTRATEGY_ORDERED]

## Methods

### NewGatewayTemplatePathPreferences

`func NewGatewayTemplatePathPreferences() *GatewayTemplatePathPreferences`

NewGatewayTemplatePathPreferences instantiates a new GatewayTemplatePathPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplatePathPreferencesWithDefaults

`func NewGatewayTemplatePathPreferencesWithDefaults() *GatewayTemplatePathPreferences`

NewGatewayTemplatePathPreferencesWithDefaults instantiates a new GatewayTemplatePathPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *GatewayTemplatePathPreferences) GetPaths() []GatewayTemplatePathPreferencesPath`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *GatewayTemplatePathPreferences) GetPathsOk() (*[]GatewayTemplatePathPreferencesPath, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *GatewayTemplatePathPreferences) SetPaths(v []GatewayTemplatePathPreferencesPath)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *GatewayTemplatePathPreferences) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetStrategy

`func (o *GatewayTemplatePathPreferences) GetStrategy() GatewayPathStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *GatewayTemplatePathPreferences) GetStrategyOk() (*GatewayPathStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *GatewayTemplatePathPreferences) SetStrategy(v GatewayPathStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *GatewayTemplatePathPreferences) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


